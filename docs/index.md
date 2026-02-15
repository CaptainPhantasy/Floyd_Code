# Floyd Documentation

Welcome to the Floyd documentation. Floyd is an AI-powered development terminal with a rich TUI interface.

## Feature Index

### Core Features

| Feature | Description | Guide |
|---------|-------------|-------|
| **Interactive TUI** | Full chat interface with Bubble Tea | [First Session](getting-started/first-session.md) |
| **200+ AI Models** | Multiple provider support | [Configuration](getting-started/configuration.md) |
| **Session Management** | Save, resume, export sessions | [CLI Reference](reference/cli.md) |
| **File Operations** | Read, edit, create files | [CLI Reference](reference/cli.md#file-operations) |
| **Code Analysis** | Analyze codebase structure | [CLI Reference](reference/cli.md#codebase-analysis) |
| **Prompt Templates** | Pre-built coding templates | [CLI Reference](reference/cli.md#prompt-templates) |
| **MCP Servers** | Extensible via MCP | [Configuration](getting-started/configuration.md#mcp-servers) |

### User Interface

| Feature | Description | Reference |
|---------|-------------|-----------|
| Chat View | Message history with AI | [Keyboard Shortcuts](reference/keyboard-shortcuts.md) |
| Editor | Input field with attachments | [Keyboard Shortcuts](reference/keyboard-shortcuts.md#editor) |
| Sessions | Switch between conversations | [Keyboard Shortcuts](reference/keyboard-shortcuts.md) |
| Model Switcher | Change AI model mid-session | [Keyboard Shortcuts](reference/keyboard-shortcuts.md) |
| Terminal | Embedded terminal | [Keyboard Shortcuts](reference/keyboard-shortcuts.md) |
| Command Palette | Quick actions | [Keyboard Shortcuts](reference/keyboard-shortcuts.md) |

### Supported Providers

| Provider | Type | Notes |
|----------|------|-------|
| OpenAI | Native | GPT-4o, GPT-5, O-series |
| Anthropic | Native | Claude 3.5, 4 series |
| GLM / Z.AI | OpenAI-compatible | GLM-4.5, GLM-5 |
| DeepSeek | OpenAI-compatible | DeepSeek Chat, Reasoner |
| OpenRouter | Proxy | 150+ models |
| xAI | OpenAI-compatible | Grok-3, Grok-4 |

## Getting Started

1. [Installation](getting-started/installation.md) - Get Floyd running on your system
2. [First Session](getting-started/first-session.md) - Your first AI conversation
3. [Configuration](getting-started/configuration.md) - Set up providers and models

## Reference

- [CLI Reference](reference/cli.md) - All commands and flags
- [Keyboard Shortcuts](reference/keyboard-shortcuts.md) - Complete keybinding reference
- [GLM-5 Guide](GLM-5-Guide.md) - GLM-5 integration details
- [GLM-5 Benchmark Report](GLM5_BENCHMARK_REPORT.md) - Performance analysis

## For Developers

- [AGENTS.md](../AGENTS.md) - Development guide and code conventions
- [FLOYD.md](../FLOYD.md) - Agent protocol documentation

## Getting Help

- Press `Ctrl+G` in Floyd to see keyboard shortcuts
- Run `floyd --help` for CLI commands
- Run `floyd [command] --help` for command details
