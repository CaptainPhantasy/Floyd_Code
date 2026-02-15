# Keyboard Shortcuts

Complete reference for all Floyd keyboard shortcuts.

## Global Shortcuts

These work from anywhere in the application.

| Key | Action |
|-----|--------|
| `Ctrl+C` | Quit Floyd |
| `Ctrl+G` | Show help / more shortcuts |
| `Ctrl+P` | Open command palette |
| `Ctrl+L` | Switch model |
| `Ctrl+M` | Switch model (alternate) |
| `Ctrl+S` | Open sessions manager |
| `Ctrl+T` | Open embedded terminal |
| `Ctrl+Z` | Suspend Floyd |
| `Tab` | Change focus between panels |

---

## Editor Shortcuts

These work when the input field is focused.

| Key | Action |
|-----|--------|
| `Enter` | Send message |
| `Shift+Enter` | Newline |
| `Ctrl+J` | Newline (alternate) |
| `/` | Add file or open commands |
| `@` | Mention a file |
| `Ctrl+F` | Add image or attachment |
| `Ctrl+O` | Open in external editor |
| `Ctrl+R` then `{i}` | Delete attachment at index i |
| `Ctrl+R` then `R` | Delete all attachments |
| `Esc` | Cancel delete mode |
| `Up` | Previous prompt in history |
| `Down` | Next prompt in history |

---

## Chat Navigation

These work when viewing the chat history.

### Basic Navigation

| Key | Action |
|-----|--------|
| `j` | Scroll down |
| `k` | Scroll up |
| `Down` | Scroll down |
| `Up` | Scroll up |
| `Ctrl+J` | Scroll down (vim-style) |
| `Ctrl+K` | Scroll up (vim-style) |

### Page Navigation

| Key | Action |
|-----|--------|
| `d` | Half page down |
| `u` | Half page up |
| `f` | Page down |
| `b` | Page up |
| `Space` | Page down (alternate) |
| `PgDn` | Page down |
| `PgUp` | Page up |

### Quick Jump

| Key | Action |
|-----|--------|
| `g` | Jump to top (home) |
| `G` | Jump to bottom (end) |
| `Home` | Jump to top |
| `End` | Jump to bottom |

### Item Navigation

| Key | Action |
|-----|--------|
| `Shift+Up` | Up one message item |
| `Shift+Down` | Down one message item |
| `K` (shift+k) | Up one message item |
| `J` (shift+j) | Down one message item |

### Selection & Actions

| Key | Action |
|-----|--------|
| `Space` | Expand/collapse message |
| `c` | Copy selection |
| `y` | Copy selection (vim-style) |
| `Esc` | Clear selection / cancel |

### Session Controls

| Key | Action |
|-----|--------|
| `Ctrl+N` | Start new session |
| `Ctrl+D` | Toggle details panel |
| `Ctrl+Space` | Toggle tasks/pills |
| `Left` | Switch to previous section |
| `Right` | Switch to next section |

---

## Initialization Dialog

When Floyd starts in a new project.

| Key | Action |
|-----|--------|
| `y` | Yes / Confirm |
| `n` | No / Cancel |
| `Enter` | Select / Confirm |
| `Tab` | Switch option |
| `Left/Right` | Switch option |
| `Esc` | Cancel / No |

---

## Sessions Panel

When viewing saved sessions.

| Key | Action |
|-----|--------|
| `Up/Down` | Navigate sessions |
| `Enter` | Load selected session |
| `d` | Delete selected session |
| `Esc` | Close sessions panel |

---

## Model Switcher

When selecting a model.

| Key | Action |
|-----|--------|
| `Up/Down` | Navigate models |
| `Enter` | Select model |
| `Esc` | Cancel |
| `Tab` | Switch provider |

---

## Command Palette

When the command palette is open.

| Key | Action |
|-----|--------|
| `Up/Down` | Navigate commands |
| `Enter` | Execute command |
| `Esc` | Close palette |
| Type | Filter commands |

---

## Terminal Panel

When the embedded terminal is focused.

| Key | Action |
|-----|--------|
| `Esc` | Exit terminal / return to chat |
| Standard terminal keys | Normal terminal interaction |

---

## Quick Reference Card

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  FLOYD KEYBOARD SHORTCUTS                                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│  GLOBAL         │  EDITOR           │  NAVIGATION                            │
│  ─────────────  │  ───────────────  │  ───────────────                       │
│  Ctrl+C  Quit   │  Enter     Send   │  j/k      Scroll down/up              │
│  Ctrl+G  Help   │  Shift+Enter Newl │  d/u      Half page                   │
│  Ctrl+P  Cmds   │  /         File   │  f/b      Full page                   │
│  Ctrl+L  Model  │  @         Ment.  │  g/G      Top/Bottom                  │
│  Ctrl+S  Sess.  │  Ctrl+F   Image   │  Space    Expand                      │
│  Ctrl+T  Term   │  Ctrl+O   Editor  │  c/y      Copy                        │
│  Ctrl+N  New    │  Up/Down  History │  Esc      Cancel                      │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Tips

1. **Vim users**: `j/k` for scrolling, `d/u` for half-pages, `g/G` for jump
2. **Quick file add**: Type `/` and start typing the filename
3. **History**: Use Up/Down to recall previous prompts
4. **Copy code**: Navigate to a code block and press `c` or `y`
5. **Multi-line**: Use `Shift+Enter` or `Ctrl+J` for newlines in your prompt
6. **Cancel anything**: `Esc` usually cancels the current action
