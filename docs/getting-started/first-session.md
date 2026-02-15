# First Session

This guide walks you through your first Floyd session.

## Launching Floyd

Open a terminal in your project directory and run:

```bash
floyd
```

The TUI will launch with the main chat interface.

## Initial Setup

On first run, Floyd may prompt you to:

1. **Select a model** - Choose from available providers
2. **Configure API keys** - If not already set in environment

## The Interface

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Floyd v3.4                                              glm-5  ●  Connected │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                                                                     │   │
│  │  Welcome to Floyd! I'm ready to help you with coding tasks.        │   │
│  │                                                                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ User                                                                 │   │
│  │ Help me understand the structure of this codebase                   │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│  > Type your message here...                               │  Ctrl+G Help  │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Key Areas

1. **Header** - Shows version, current model, connection status
2. **Chat Area** - Message history between you and the AI
3. **Input Field** - Where you type your prompts
4. **Status Bar** - Keyboard shortcuts hint

## Sending Your First Prompt

1. Type your question in the input field
2. Press `Enter` to send
3. The AI will respond in the chat area

Example prompts to try:

```
What does this codebase do?
Explain the main entry points
Help me find where authentication is handled
Write a test for the config loader
```

## Adding Context with Files

Floyd can see files you attach:

### Method 1: Type `/`

Type `/` in the input field to open a file picker. Navigate to select files.

### Method 2: Mention files with `@`

Type `@` followed by a filename to reference it:

```
@main.go What does this file do?
```

### Method 3: Add images

Press `Ctrl+F` to add image attachments for visual context.

## Navigating the Chat

| Action | Keys |
|--------|------|
| Scroll up/down | `j/k` or Arrow keys |
| Page down/up | `f/b` or PgDn/PgUp |
| Half page | `d/u` |
| Go to top | `g` |
| Go to bottom | `G` |
| Copy selection | `c` or `y` |
| Expand/collapse | `Space` |

## Starting a New Session

Press `Ctrl+N` to start a fresh conversation. Previous sessions are saved automatically.

## Switching Models

Press `Ctrl+L` to open the model switcher. Select a different AI model for this session.

## Sessions Management

Press `Ctrl+S` to:
- View saved sessions
- Resume previous conversations
- Delete old sessions

## Getting Help

Press `Ctrl+G` to see all available keyboard shortcuts and commands.

## Command Palette

Press `Ctrl+P` to open the command palette for quick access to:
- Export session
- Change theme
- View statistics
- Configuration options

## Exiting

Press `Ctrl+C` to quit Floyd. Your session is automatically saved.

## Next Steps

- [Configuration](configuration.md) - Add more AI providers
- [CLI Reference](../reference/cli.md) - Learn all commands
- [Keyboard Shortcuts](../reference/keyboard-shortcuts.md) - Master navigation
