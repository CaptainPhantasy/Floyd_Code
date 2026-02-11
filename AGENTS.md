# Floyd Development Guide

## Build/Test/Lint Commands

| Command | Description |
| :--- | :--- |
| `go build .` / `go run .` | Build/run the project |
| `task test` / `go test ./...` | Run tests |
| `go test ./internal/agent -run TestName` | Run specific test |
| `go test ./... -update` | Update golden files |
| `task lint:fix` | Run linter with auto-fix |
| `task fmt` / `gofumpt -w .` | Format code |
| `task dev` | Run with profiling enabled |

## Code Style Guidelines

- **Imports**: goimports format (stdlib, external, internal groups)
- **Formatting**: gofumpt (stricter than gofmt)
- **Naming**: PascalCase for exported, camelCase for unexported
- **Errors**: Return errors explicitly, use `fmt.Errorf` for wrapping
- **Context**: Pass `context.Context` as first parameter for operations
- **Testing**: testify's `require` package, `t.Parallel()`, `t.TempDir()`
- **JSON tags**: snake_case for field names
- **File permissions**: octal notation (0o755, 0o644)
- **Log messages**: Start with capital letter ("Failed to save session")
- **Comments**: End with period when on own line

## Architecture Components

| Component | Source File | Status |
| :--- | :--- | :--- |
| Coordinator | `coordinator.go` | Active - Orchestrates agents, tools, providers |
| SessionAgent | `agent.go` | Active - Session-based AI agent with streaming |
| Tool Registry | `tools/tools.go` | Active - Manages all agent tools |
| Prompt Builder | `prompt/prompt.go` | Active - Builds system prompts from templates |
| Hyper Provider | `hyper/provider.go` | Active - Custom LLM provider integration |
| MCP Integration | `tools/mcp/` | Active - Model Context Protocol server tools |

## Agent Tools

| Tool | File | Description |
| :--- | :--- | :--- |
| bash | `bash.go` | Execute shell commands |
| edit | `edit.go` | Edit files with exact string replacement |
| view | `view.go` | Read files with line numbers |
| write | `write.go` | Create/overwrite files |
| multiedit | `multiedit.go` | Multiple edits in one operation |
| grep | `grep.go` | Search file contents with regex |
| glob | `glob.go` | Find files by pattern matching |
| ls | `ls.go` | List directory contents |
| fetch | `fetch.go` | Fetch remote URLs |
| download | `download.go` | Download files from URLs |
| job_output | `job_output.go` | Get output from background jobs |
| job_kill | `job_kill.go` | Kill background jobs |
| todos | `todos.go` | Manage todo list |
| diagnostics | `diagnostics.go` | Get LSP diagnostics |
| references | `references.go` | Find symbol references |
| sourcegraph | `sourcegraph.go` | Search Sourcegraph |
| agentic_fetch | `agentic_fetch_tool.go` | Delegate to sub-agent for searches |

## Testing with Mock Providers

```go
func TestYourFunction(t *testing.T) {
    originalUseMock := config.UseMockProviders
    config.UseMockProviders = true
    defer func() {
        config.UseMockProviders = originalUseMock
        config.ResetProviders()
    }()
    config.ResetProviders()
    // Test logic here
}
```

## Committing

Use semantic commits: `fix:`, `feat:`, `chore:`, `refactor:`, `docs:`, `sec:`
Keep commits to one line where possible.

## Database Schema Validation Before Refactor Completion

**CRITICAL**: Any refactor task that touches database-related code must include a final verification step that the database schema is in working order before being marked ready for audit.

### Required Checklist for Refactors Affecting Database Code

- [ ] **Schema sync**: Verify `internal/db/models.go` struct fields match the actual database schema
- [ ] **Migration completeness**: Ensure all new columns have corresponding migrations in `internal/db/migrations/`
- [ ] **Backfill handling**: Add missing columns to `ensureColumns()` in `internal/db/connect.go` with proper DDL
- [ ] **Test with real database**: Run Floyd binary against an actual `~/.floyd/floyd.db` to confirm no schema errors
- [ ] **Edge cases**: Test both fresh database creation and existing database migration paths

### Common Schema Mismatch Symptoms

```
SQL logic error: table X has no column named Y
```

If you see this error during testing, the `ensureColumns()` function needs updating.

### Quick Verification Command

```bash
# Verify the binary works against your local Floyd database
floyd --help
# If this fails with a schema error, the refactor is NOT complete
```
