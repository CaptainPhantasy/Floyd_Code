package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

// Connect opens a SQLite database connection and runs migrations.
func Connect(ctx context.Context, dataDir string) (*sql.DB, error) {
	if dataDir == "" {
		return nil, fmt.Errorf("data.dir is not set")
	}
	dbPath := filepath.Join(dataDir, "floyd.db")

	db, err := openDB(dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Back-fill columns that were added to the initial migration after
	// some databases had already been created. This runs before goose
	// so the SQL migrations always see a consistent schema.
	if err := ensureColumns(ctx, db); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ensure columns: %w", err)
	}

	goose.SetBaseFS(FS)

	if err := goose.SetDialect("sqlite3"); err != nil {
		slog.Error("Failed to set dialect", "error", err)
		return nil, fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		slog.Error("Failed to apply migrations", "error", err)
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return db, nil
}

// ensureColumns idempotently adds columns that may be missing from
// databases created before the column was part of the initial
// migration. SQLite does not support IF NOT EXISTS for ALTER TABLE
// ADD COLUMN, so we check pragma_table_info first.
func ensureColumns(ctx context.Context, db *sql.DB) error {
	type col struct {
		Table  string
		Column string
		DDL    string
	}

	backfills := []col{
		{"sessions", "parent_session_id", "ALTER TABLE sessions ADD COLUMN parent_session_id TEXT"},
	}

	for _, c := range backfills {
		// Check if table exists first
		var tableExists int
		err := db.QueryRowContext(ctx,
			"SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?",
			c.Table,
		).Scan(&tableExists)
		if err != nil || tableExists == 0 {
			// Table doesn't exist yet, skip this backfill (will be created by migrations)
			continue
		}

		var count int
		err = db.QueryRowContext(ctx,
			"SELECT COUNT(*) FROM pragma_table_info(?) WHERE name = ?",
			c.Table, c.Column,
		).Scan(&count)
		if err != nil {
			return fmt.Errorf("checking column %s.%s: %w", c.Table, c.Column, err)
		}
		if count == 0 {
			if _, err := db.ExecContext(ctx, c.DDL); err != nil {
				return fmt.Errorf("adding column %s.%s: %w", c.Table, c.Column, err)
			}
			slog.Info("Added missing column", "table", c.Table, "column", c.Column)
		}
	}
	return nil
}
