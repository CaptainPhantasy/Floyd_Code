package main

import (
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/legacy-ai/floyd/internal/cmd"
)

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			cmd.Execute()
			return
		}
	}

	// Load global config from ~/.floyd/.env.local first
	if homeDir, err := os.UserHomeDir(); err == nil {
		globalEnv := filepath.Join(homeDir, ".floyd", ".env.local")
		_ = godotenv.Load(globalEnv)
	}

	// Then load local .env.local (overrides global)
	_ = godotenv.Load(".env.local")

	if os.Getenv("FLOYD_PROFILE") != "" {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			slog.Info("Serving pprof at localhost:6060")
			if httpErr := http.ListenAndServe("localhost:6060", mux); httpErr != nil {
				slog.Error("Failed to pprof listen", "error", httpErr)
			}
		}()
	}

	cmd.Execute()
}
