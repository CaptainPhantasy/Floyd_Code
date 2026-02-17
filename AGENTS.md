# Floyd Development Guide

Floyd is a Go-based AI coding agent with a BubbleTea TUI, multi-provider LLM support, and MCP integration.

## Quick Reference

| Command | Description |
|---------|-------------|
| `go build .` / `go run .` | Build/run the project |
| `task test` / `go test ./...` | Run all tests |
| `go test ./internal/agent -run TestName` | Run specific test |
| `go test ./... -update` | Update golden files |
| `task lint:fix` | Run linter with auto-fix |
| `task fmt` / `gofumpt -w .` | Format code |
| `task dev` | Run with profiling enabled (pprof at localhost:6060) |
| `task test:record` | Re-record all VCR cassettes |
| `task schema` | Generate JSON schema |
| `task hyper` | Update embedded Hyper provider.json |

## Code Style Guidelines

- **Imports**: goimports format (stdlib, external, internal groups separated by blank lines)
- **Formatting**: gofumpt (stricter than gofmt)
- **Naming**: PascalCase for exported, camelCase for unexported
- **Errors**: Return errors explicitly, use `fmt.Errorf` for wrapping
- **Context**: Pass `context.Context` as first parameter for operations
- **Testing**: testify's `require` and `assert` packages, `t.Parallel()`, `t.TempDir()`
- **JSON tags**: snake_case for field names (e.g., `json:"file_path"`)
- **File permissions**: octal notation (0o755, 0o644)
- **Log messages**: Start with capital letter ("Failed to save session")
- **Comments**: End with period when on own line

## Project Structure

```
/Volumes/Storage/floyd-main/
├── main.go                    # Entry point, pprof setup
├── go.mod                     # Go 1.25.5
├── Taskfile.yaml              # Task runner commands
├── sqlc.yaml                  # SQL code generation config
├── floyd-schema.json          # Configuration JSON schema
├── .golangci.yml              # Linter configuration
├── .goreleaser.yml            # Release configuration
├── internal/
│   ├── agent/                 # Core agent implementation
│   │   ├── agent.go           # SessionAgent with streaming, queuing, summarization
│   │   ├── coordinator.go     # Orchestrates agents, tools, providers
│   │   ├── prompt/            # Template-based prompt builder
│   │   ├── hyper/             # Custom LLM provider
│   │   ├── tools/             # All agent tool implementations
│   │   │   ├── mcp/           # Model Context Protocol integration
│   │   │   ├── edit.go        # File editing with exact string matching
│   │   │   ├── bash.go        # Shell command execution
│   │   │   ├── view.go        # File reading
│   │   │   └── ...            # Other tools
│   │   └── templates/         # Embedded markdown templates
│   ├── ui/                    # BubbleTea TUI
│   │   ├── model/             # Main UI model and components
│   │   ├── chat/              # Chat message rendering
│   │   ├── dialog/            # Dialog implementations
│   │   ├── list/              # Generic list component
│   │   └── styles/            # All style definitions
│   ├── db/                    # Database layer (SQLite + sqlc)
│   │   ├── sql/               # SQL queries
│   │   ├── migrations/        # Goose migrations
│   │   ├── models.go          # Generated models
│   │   └── *.sql.go           # Generated query code
│   ├── session/               # Session management
│   ├── message/               # Message handling
│   ├── config/                # Configuration types
│   ├── permission/            # Permission service
│   ├── lsp/                   # LSP client integration
│   └── ...                    # Other internal packages
└── scripts/                   # Build/dev scripts
```

## Core Architecture

### Coordinator (`internal/agent/coordinator.go`)
- Creates and manages agents
- Builds LLM providers from configuration
- Handles OAuth2 token refresh on 401 errors
- Coordinates tool building and filtering

### SessionAgent (`internal/agent/agent.go`)
- Session-based conversation management
- Message queuing for concurrent requests
- Streaming responses with OnTextDelta, OnToolCall, OnReasoningDelta callbacks
- Auto-summarization at context thresholds (200K tokens for large windows)
- Works around provider limitations (e.g., images in tool results)

### Tool Pattern (`internal/agent/tools/`)
All tools implement `fantasy.AgentTool`:
```go
func NewEditTool(...) fantasy.AgentTool {
    return fantasy.NewAgentTool(
        "edit",
        string(editDescription), // Embedded from .md file
        func(ctx context.Context, params EditParams, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
            // Implementation
        },
    )
}
```

- Tool descriptions are embedded from `.md` files (e.g., `edit.go` → `edit.md`)
- Parameter structs use JSON tags and description tags
- Return `fantasy.NewTextErrorResponse()` for user-facing errors

## Testing Patterns

### VCR Cassettes
Tests use `charm.land/x/vcr` to record/replay HTTP interactions:
```go
func setupAgent(t *testing.T, pair modelPair) (SessionAgent, fakeEnv) {
    r := vcr.NewRecorder(t)
    large, small := getModels(t, r, pair)
    // ...
}
```

- Cassettes stored in `internal/agent/testdata/TestCoderAgent/{provider}/`
- Re-record with: `task test:record` or `go test ./internal/agent -update`
- Tests run against multiple providers: anthropic-sonnet, openai-gpt-5, openrouter-kimi-k2, zai-glm4.6

### Test Structure
```go
func TestSomething(t *testing.T) {
    t.Parallel()
    
    require.NoError(t, err)
    assert.Equal(t, expected, actual)
}
```

### Mock Provider Pattern
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

## Database Layer (sqlc)

### Generating Code
After modifying SQL queries or migrations:
```bash
sqlc generate
```

### Migrations
- Located in `internal/db/migrations/`
- Use Goose format: `YYYYMMDDHHMMSS_description.sql`
- Schema validation is handled in `internal/db/connect.go` via `ensureColumns()`

### Common Schema Error
```
SQL logic error: table X has no column named Y
```
→ Update `ensureColumns()` in `internal/db/connect.go` with proper DDL

## UI Development (BubbleTea v2)

### Key Principles (from `internal/ui/AGENTS.md`)
- Never do IO or expensive work in `Update`; always use a `tea.Cmd`
- Never change model state inside a command - use messages
- Components should be "dumb" - expose methods, return `tea.Cmd` for side effects
- Handle rendering via `Render(width int) string`

### Architecture
- **Main Model** (`model/ui.go`): Message routing, focus management, layout
- **Components**: Don't handle messages directly; expose state mutation methods
- **Chat Logic** (`model/chat.go`): Most chat-related logic
- **Styles**: All in `styles/styles.go`, accessed via `*common.Common`

### Common Gotchas
- Account for padding/borders in width calculations
- Use `tea.Batch()` when returning multiple commands
- Pass `*common.Common` to components needing styles or app access

## LLM Providers

Supported providers (via `charm.land/fantasy` and `charm.land/catwalk`):
- Anthropic (Claude with thinking support)
- OpenAI (Responses API with reasoning)
- Google (Gemini with thinking_config)
- OpenRouter (with exacto support)
- Hyper (custom proxy provider)
- Bedrock, Azure, Vercel, OpenAI-compatible

## Commit Guidelines

Use semantic commits:
- `feat:` - New features
- `fix:` - Bug fixes
- `chore:` - Maintenance tasks
- `refactor:` - Code restructuring
- `docs:` - Documentation
- `sec:` - Security fixes

Keep commits to one line where possible.

## Important Gotchas

### Database Schema Validation
Before completing any refactor touching database code:
1. Verify `internal/db/models.go` matches actual schema
2. Ensure migrations exist in `internal/db/migrations/`
3. Add missing columns to `ensureColumns()` in `internal/db/connect.go`
4. Test against real `~/.floyd/floyd.db`

### Edit Tool Exact Matching
The edit tool requires EXACT string matching including:
- Whitespace and indentation
- Line breaks
- Trailing spaces

### Log Message Capitalization
Log messages must start with capital letters. The linter checks this via `scripts/check_log_capitalization.sh`.

### Context Keys
Tools receive context with:
- `SessionIDContextKey` - Current session ID
- `MessageIDContextKey` - Current message ID
- `SupportsImagesContextKey` - Whether model supports images
- `ModelNameContextKey` - Current model name

## CI/CD

- **Build**: Runs on push/PR across ubuntu, macos, windows
- **Test**: `go test -race -failfast ./...`
- **Release**: Uses goreleaser with multi-platform builds

## Key Dependencies

- `charm.land/fantasy` - Agent framework
- `charm.land/catwalk` - LLM provider abstraction
- `charm.land/bubbletea/v2` - TUI framework
- `charm.land/lipgloss/v2` - Styling
- `github.com/modelcontextprotocol/go-sdk` - MCP support
- `modernc.org/sqlite` / `ncruces/go-sqlite3` - SQLite drivers
- `github.com/pressly/goose/v3` - Migrations
