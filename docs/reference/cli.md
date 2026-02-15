# CLI Reference

Complete reference for all Floyd CLI commands.

## Global Commands

### `floyd`

Launch the interactive TUI.

```bash
floyd
```

### `floyd --version`

Print the version.

```bash
floyd -v
floyd --version
```

Output: `floyd version v3.4`

### `floyd --help`

Show help for all commands.

```bash
floyd -h
floyd --help
```

## Global Flags

| Flag | Description |
|------|-------------|
| `-v, --version` | Print version |
| `-h, --help` | Show help |
| `-d, --debug` | Enable debug logging |
| `-c, --cwd <path>` | Set working directory |
| `-D, --data-dir <path>` | Set custom data directory |
| `-y, --yolo` | Auto-accept all permissions |

---

## Run Command

### `floyd run`

Run a single non-interactive prompt.

```bash
# Basic usage
floyd run "Explain the use of context in Go"

# With specific model
floyd run -m glm-5 "Generate a README"

# Quiet mode (no spinner)
floyd run -q "What is this code doing?"

# Verbose mode (show logs)
floyd run -v "Debug this function"

# Pipe input
curl https://example.com | floyd run "Summarize this website"

# Read from file
floyd run "What is this code doing?" <<< main.go
```

#### Run Flags

| Flag | Description |
|------|-------------|
| `-m, --model <model>` | Model to use (e.g., `glm-5`, `openai/gpt-4o`) |
| `--small-model <model>` | Small model for quick tasks |
| `-q, --quiet` | Hide spinner |
| `-v, --verbose` | Show logs |

---

## Model Commands

### `floyd models`

List all available models from configured providers.

```bash
floyd models
```

Output format: `provider/model-id`

Example:
```
openai/gpt-4o
openai/gpt-4o-mini
glm/glm-5
anthropic/claude-sonnet-4-20250514
```

### `floyd update-providers`

Update provider information from remote or local source.

```bash
# Update from default source
floyd update-providers

# Update from custom URL
floyd update-providers https://example.com/providers.json

# Update from local file
floyd update-providers /path/to/providers.json

# Update embedded providers
floyd update-providers embedded

# Update Hyper providers
floyd update-providers --source=hyper
```

#### Update Providers Flags

| Flag | Description |
|------|-------------|
| `--source <source>` | Provider source (`catwalk` or `hyper`) |

---

## Session Commands

### `floyd projects`

List project directories with recent activity.

```bash
floyd projects
```

Output format: `Project Path | Data Directory | Last Activity`

### `floyd logs`

View application logs.

```bash
# View recent logs
floyd logs

# Follow logs (like tail -f)
floyd logs -f

# Show last N lines
floyd logs -t 100
```

#### Logs Flags

| Flag | Description |
|------|-------------|
| `-f, --follow` | Follow log output |
| `-t, --tail <n>` | Show last N lines (default: 1000) |

### `floyd stats`

Show usage statistics.

```bash
floyd stats
```

Generates: `.floyd/stats/index.html`

### `floyd dirs`

Print directories used by Floyd.

```bash
floyd dirs
```

Output:
```
/Users/you/.config/floyd
/Users/you/.local/share/floyd
```

---

## File Operations

### `floyd file`

File operations with safety checks.

```bash
floyd file <command> [args]
```

#### Subcommands

| Command | Description |
|---------|-------------|
| `read <path>` | Read file contents |
| `write <path>` | Write to file |
| `append <path>` | Append to file |
| `apply <path>` | Apply content from stdin |
| `diff <path>` | Generate unified diff |
| `ls <dir>` | List directory |
| `find <dir>` | Find files by pattern |
| `info <path>` | Show file/directory info |
| `mkdir <dir>` | Create directory |
| `cp <src> <dest>` | Copy file |
| `mv <src> <dest>` | Move/rename file |
| `rm <path>` | Delete file |
| `temp` | Create temporary file |

#### Examples

```bash
# Read a file
floyd file read main.go

# Write to file
floyd file write output.txt --content "Hello"

# List directory
floyd file ls ./src

# Find files
floyd file find . --pattern "*.go"

# Create directory
floyd file mkdir ./new-dir

# Copy file
floyd file cp source.go backup.go

# Delete file
floyd file rm old-file.txt
```

---

## Codebase Analysis

### `floyd codebase`

Analyze codebase structure and dependencies.

```bash
floyd codebase <command> [args]
```

#### Subcommands

| Command | Description |
|---------|-------------|
| `analyze [dir]` | Analyze and print summary |
| `deps [dir]` | Analyze dependency manifests |
| `search <term> [dir]` | Search for content in codebase |

#### Examples

```bash
# Analyze current directory
floyd codebase analyze

# Analyze specific directory
floyd codebase analyze ./internal

# Show dependencies
floyd codebase deps

# Search codebase
floyd codebase search "authentication"
```

---

## Prompt Templates

### `floyd prompt`

Manage and render prompt templates.

```bash
floyd prompt <command> [args]
```

#### Subcommands

| Command | Description |
|---------|-------------|
| `list` | List available templates |
| `render <template>` | Render a template |

#### Available Templates

| Template | Description |
|----------|-------------|
| `debugCode` | Debug code issues |
| `documentCode` | Generate documentation |
| `explainCode` | Explain code behavior |
| `generateCode` | Generate code |
| `refactorCode` | Refactor code |
| `reviewCode` | Review code quality |
| `testCode` | Generate tests |

#### Examples

```bash
# List templates
floyd prompt list

# Render a template
floyd prompt render generateCode --input "HTTP client library"
```

---

## Authentication

### `floyd login`

Authenticate with platforms.

```bash
# Login to Hyper
floyd login

# Login to GitHub Copilot
floyd login copilot
```

---

## Command Execution

### `floyd exec`

Execute shell commands with safety checks.

```bash
# Basic execution
floyd exec "go test ./..."

# With timeout
floyd exec "npm install" --timeout 60s

# Background execution
floyd exec bg "long-running-command"
```

#### Exec Flags

| Flag | Description |
|------|-------------|
| `--cwd <dir>` | Working directory |
| `--env <KEY=value>` | Environment variable override |
| `--timeout <duration>` | Execution timeout |
| `--shell <shell>` | Shell to use |
| `--stderr` | Include stderr in output |
| `--max-buffer <bytes>` | Maximum output size |
| `--allow-prefix <prefix>` | Allowed command prefix |
| `--deny-prefix <prefix>` | Denied command prefix |
| `--allow-regex <regex>` | Allowed command regex |
| `--deny-regex <regex>` | Denied command regex |

---

## Shell Completion

### `floyd completion`

Generate autocompletion scripts.

```bash
# Bash
floyd completion bash > /etc/bash_completion.d/floyd

# Zsh
floyd completion zsh > "${fpath[1]}/_floyd"

# Fish
floyd completion fish > ~/.config/fish/completions/floyd.fish

# PowerShell
floyd completion powershell | Out-String | Invoke-Expression
```

---

## AI Helpers

### `floyd ai`

AI-related helper commands.

```bash
# Dry run - render request without calling API
floyd ai dry-run "Test prompt"
```

---

## Command Combinations

```bash
# Run in a specific directory with debug
floyd -c /path/to/project -d run "Analyze this code"

# Use custom data directory
floyd -D /custom/data run "Quick task"

# YOLO mode for automated workflows
floyd -y run "Fix all linting issues"
```
