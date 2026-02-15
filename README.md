# Floyd

An AI-powered development terminal with a rich TUI interface.

Floyd brings AI coding assistants directly into your terminal with a full-featured chat interface, file operations, code analysis, and session management.

## Features

- **Interactive TUI** - Full chat interface with syntax highlighting, markdown rendering, and split-view diffs
- **200+ AI Models** - Support for OpenAI, Anthropic, GLM, DeepSeek, Grok, and OpenRouter providers
- **Session Management** - Save, resume, and export coding sessions
- **File Operations** - Built-in commands for reading, editing, and managing files
- **Code Analysis** - Analyze codebases for structure and dependencies
- **Prompt Templates** - Pre-built templates for common coding tasks
- **MCP Integration** - Extensible via Model Context Protocol servers
- **Shell Autocompletion** - Bash, Fish, Zsh, and PowerShell support

## Installation

### Download Binary

Download the latest release for your platform from the [releases page](https://github.com/legacy-ai/floyd/releases).

### Build from Source

```bash
git clone https://github.com/legacy-ai/floyd.git
cd floyd
go build -o floyd .
```

### Go Install

```bash
go install github.com/legacy-ai/floyd@latest
```

## Quick Start

### 1. Launch Floyd

```bash
floyd
```

This opens the interactive TUI. On first run, Floyd will guide you through initial setup.

### 2. Configure a Provider

Create a `floyd.json` in your project directory:

```json
{
  "providers": {
    "openai": {
      "name": "OpenAI",
      "type": "openai",
      "api_key": "$OPENAI_API_KEY",
      "models": [
        { "id": "gpt-4o", "name": "GPT-4o" }
      ]
    }
  }
}
```

Set your API key:

```bash
export OPENAI_API_KEY=your-api-key-here
```

### 3. Start Coding

Type your prompt in the input field and press `Enter` to send. Use `Ctrl+G` to see all keyboard shortcuts.

## CLI Commands

| Command | Description |
|---------|-------------|
| `floyd` | Launch interactive TUI |
| `floyd run "prompt"` | Run single non-interactive prompt |
| `floyd models` | List all available models |
| `floyd projects` | List project directories |
| `floyd logs` | View application logs |
| `floyd stats` | Show usage statistics |
| `floyd --help` | Show all commands |

### Common Flags

| Flag | Description |
|------|-------------|
| `-v, --version` | Print version |
| `-d, --debug` | Enable debug logging |
| `-c, --cwd` | Set working directory |
| `-D, --data-dir` | Set custom data directory |
| `-y, --yolo` | Auto-accept all permissions |

## Keyboard Shortcuts

### Global

| Key | Action |
|-----|--------|
| `Ctrl+C` | Quit |
| `Ctrl+G` | Help |
| `Ctrl+P` | Command palette |
| `Ctrl+L` | Switch model |
| `Ctrl+S` | Sessions |
| `Ctrl+T` | Terminal |
| `Ctrl+Z` | Suspend |

### Editor

| Key | Action |
|-----|--------|
| `Enter` | Send message |
| `Shift+Enter` / `Ctrl+J` | Newline |
| `/` | Add file or command |
| `@` | Mention file |
| `Ctrl+F` | Add image/attachment |
| `Ctrl+O` | Open external editor |
| `Up/Down` | History navigation |

### Chat Navigation

| Key | Action |
|-----|--------|
| `j/k` or `Up/Down` | Scroll |
| `d/u` | Half page down/up |
| `f/b` or `PgDn/PgUp` | Page down/up |
| `g/G` | Home/End |
| `Space` | Expand/collapse |
| `c/y` | Copy selection |
| `Ctrl+N` | New session |
| `Ctrl+D` | Toggle details |

## Configuration

Floyd uses `floyd.json` for configuration. Place it in your project root or in `~/.config/floyd/` for global settings.

### Minimal Configuration

```json
{
  "providers": {
    "zai": {
      "name": "Z.AI",
      "type": "openai-compat",
      "base_url": "https://api.z.ai/api/paas/v4/",
      "api_key": "$ZAI_API_KEY",
      "models": [
        { "id": "glm-5", "name": "GLM-5" }
      ]
    }
  }
}
```

### Environment Variables

Store API keys in `.env.local`:

```bash
OPENAI_API_KEY=sk-...
ANTHROPIC_API_KEY=sk-ant-...
ZAI_API_KEY=...
```

## Documentation

- [Installation Guide](docs/getting-started/installation.md)
- [First Session](docs/getting-started/first-session.md)
- [Configuration Reference](docs/getting-started/configuration.md)
- [CLI Reference](docs/reference/cli.md)
- [Keyboard Shortcuts](docs/reference/keyboard-shortcuts.md)
- [GLM-5 Integration Guide](docs/GLM-5-Guide.md)

## Development

For development and contribution guidelines, see [AGENTS.md](AGENTS.md).

## License

[Functional Source License 1.1 (FSL-1.1-MIT)](LICENSE.md)
