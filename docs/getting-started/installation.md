# Installation

## System Requirements

- **Operating System**: macOS, Linux, Windows (WSL recommended)
- **Architecture**: x86_64, ARM64
- **Go**: 1.21+ (for building from source)

## Installation Methods

### Option 1: Download Binary (Recommended)

Download the latest release for your platform:

```bash
# macOS (Apple Silicon)
curl -LO https://github.com/legacy-ai/floyd/releases/latest/download/floyd-darwin-arm64
chmod +x floyd-darwin-arm64
sudo mv floyd-darwin-arm64 /usr/local/bin/floyd

# macOS (Intel)
curl -LO https://github.com/legacy-ai/floyd/releases/latest/download/floyd-darwin-amd64
chmod +x floyd-darwin-amd64
sudo mv floyd-darwin-amd64 /usr/local/bin/floyd

# Linux (x86_64)
curl -LO https://github.com/legacy-ai/floyd/releases/latest/download/floyd-linux-amd64
chmod +x floyd-linux-amd64
sudo mv floyd-linux-amd64 /usr/local/bin/floyd

# Linux (ARM64)
curl -LO https://github.com/legacy-ai/floyd/releases/latest/download/floyd-linux-arm64
chmod +x floyd-linux-arm64
sudo mv floyd-linux-arm64 /usr/local/bin/floyd
```

### Option 2: Go Install

```bash
go install github.com/legacy-ai/floyd@latest
```

The binary will be installed to `$GOPATH/bin/floyd`. Make sure `$GOPATH/bin` is in your `PATH`.

### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/legacy-ai/floyd.git
cd floyd

# Build
go build -o floyd .

# Install to PATH (optional)
sudo mv floyd /usr/local/bin/
```

### Option 4: goreleaser (for development)

```bash
# Install goreleaser if not already installed
go install github.com/goreleaser/goreleaser/v2@latest

# Build snapshot
goreleaser build --snapshot --clean

# Binary will be in dist/
```

## Verify Installation

```bash
floyd --version
```

Expected output: `floyd version v3.4` (or newer)

## Shell Autocompletion

Floyd can generate autocompletion scripts for your shell:

### Bash

```bash
floyd completion bash > /etc/bash_completion.d/floyd
source ~/.bashrc
```

### Zsh

```bash
floyd completion zsh > "${fpath[1]}/_floyd"
autoload -U compinit && compinit
```

### Fish

```bash
floyd completion fish > ~/.config/fish/completions/floyd.fish
```

### PowerShell

```powershell
floyd completion powershell | Out-String | Invoke-Expression
```

## Data Directories

Floyd stores data in these locations:

| Directory | Purpose |
|-----------|---------|
| `.floyd/` (per-project) | Session data, database, supercache |
| `~/.config/floyd/` | Global configuration, custom commands |
| `~/.local/share/floyd/` | Global data, providers cache |

You can override the data directory with `-D` / `--data-dir`:

```bash
floyd -D /custom/data/path
```

## Next Steps

- [First Session](first-session.md) - Start your first AI conversation
- [Configuration](configuration.md) - Set up AI providers
