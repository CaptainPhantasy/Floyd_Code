# Configuration

Floyd uses `floyd.json` for configuration. This guide covers all configuration options.

## Configuration Locations

Floyd looks for configuration in this order:

1. `./floyd.json` - Project directory (highest priority)
2. `~/.config/floyd/floyd.json` - Global user config
3. `~/.local/share/floyd/floyd.json` - Global data config

## Minimal Configuration

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

## Full Configuration Schema

```json
{
  "$schema": "https://charm.land/floyd.json",
  "rules": [
    "API_BEHAVIOR: The backend is a pure Coding Model."
  ],
  "providers": { ... },
  "mcp": { ... },
  "lsp": { ... },
  "options": { ... }
}
```

## Providers

### OpenAI

```json
{
  "providers": {
    "openai": {
      "name": "OpenAI",
      "type": "openai",
      "api_key": "$OPENAI_API_KEY",
      "models": [
        { "id": "gpt-4o", "name": "GPT-4o" },
        { "id": "gpt-4o-mini", "name": "GPT-4o Mini" },
        { "id": "o3-mini", "name": "O3 Mini" }
      ]
    }
  }
}
```

### Anthropic

```json
{
  "providers": {
    "anthropic": {
      "name": "Anthropic",
      "type": "anthropic",
      "api_key": "$ANTHROPIC_API_KEY",
      "models": [
        { "id": "claude-sonnet-4-20250514", "name": "Claude Sonnet 4" }
      ]
    }
  }
}
```

### OpenAI-Compatible (GLM, DeepSeek, xAI)

```json
{
  "providers": {
    "glm": {
      "name": "GLM-5",
      "type": "openai-compat",
      "base_url": "https://api.z.ai/api/paas/v4/",
      "api_key": "$ZAI_API_KEY",
      "models": [
        {
          "id": "glm-5",
          "name": "GLM-5",
          "context_window": 200000,
          "default_max_tokens": 131072,
          "options": {
            "temperature": 0.1
          }
        }
      ]
    }
  }
}
```

### With Extra Body Parameters

Some providers support additional parameters:

```json
{
  "providers": {
    "glm": {
      "type": "openai-compat",
      "base_url": "https://api.z.ai/api/coding/paas/v4/",
      "api_key": "$ZAI_API_KEY",
      "extra_body": {
        "thinking": {
          "type": "enabled"
        }
      }
    }
  }
}
```

### Model Options

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Model identifier used in API calls |
| `name` | string | Display name in UI |
| `context_window` | int | Maximum context length |
| `default_max_tokens` | int | Default output token limit |
| `can_reason` | bool | Supports reasoning/thinking mode |
| `options.temperature` | float | Sampling temperature (0-2) |
| `options.top_p` | float | Nucleus sampling (0-1) |

## MCP Servers

Model Context Protocol servers extend Floyd with additional capabilities:

```json
{
  "mcp": {
    "web-search": {
      "type": "http",
      "url": "https://api.example.com/mcp/search",
      "headers": {
        "Authorization": "Bearer $API_KEY"
      }
    },
    "custom-tool": {
      "type": "stdio",
      "command": "/path/to/server",
      "args": ["--port", "8080"],
      "env": {
        "DEBUG": "true"
      }
    }
  }
}
```

### MCP Connection Types

| Type | Use Case |
|------|----------|
| `stdio` | Local command-line servers |
| `http` | Remote HTTP-based servers |
| `sse` | Server-Sent Events streaming |

## LSP Servers

Configure Language Server Protocol servers for code intelligence:

```json
{
  "lsp": {
    "gopls": {
      "options": {
        "gofumpt": true,
        "staticcheck": true
      }
    }
  }
}
```

LSP servers are automatically detected based on project markers (e.g., `go.mod` for Go).

## Options

Global behavior options:

```json
{
  "options": {
    "data_directory": ".floyd",
    "context_paths": ["FLOYD.md", "FLOYD.local.md"],
    "skills_paths": ["./skills"],
    "compact_mode": false,
    "diff_mode": "split",
    "show_completions": true,
    "auto_summarize": true,
    "summarize_threshold": 80000
  }
}
```

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `data_directory` | string | `.floyd` | Where to store session data |
| `context_paths` | []string | `["FLOYD.md"]` | Files to include in context |
| `skills_paths` | []string | `[]` | Directories containing skills |
| `compact_mode` | bool | false | Reduce UI padding |
| `diff_mode` | string | `split` | Diff display style |
| `auto_summarize` | bool | true | Auto-summarize long contexts |
| `summarize_threshold` | int | 80000 | Token threshold for summarization |

## Environment Variables

Store sensitive values in `.env.local`:

```bash
# .env.local
OPENAI_API_KEY=sk-...
ANTHROPIC_API_KEY=sk-ant-...
ZAI_API_KEY=...
GITHUB_TOKEN=...
```

Floyd loads environment variables from:

1. `~/.floyd/.env.local` - Global (loaded first)
2. `./.env.local` - Project directory (overrides global)

Reference environment variables in config with `$VARIABLE_NAME`:

```json
{
  "api_key": "$OPENAI_API_KEY"
}
```

## Rules

Add custom rules that influence AI behavior:

```json
{
  "rules": [
    "API_BEHAVIOR: Output code blocks immediately.",
    "NO_CHAT: Do not provide conversational filler.",
    "TESTING: Run tests using 'go test ./...'",
    "STYLE: Use 'go fmt' standards."
  ]
}
```

## Checking Available Models

```bash
floyd models
```

This lists all models from configured providers.

## Updating Provider Cache

```bash
# Update from default source
floyd update-providers

# Update from custom URL
floyd update-providers https://example.com/providers.json

# Update from local file
floyd update-providers /path/to/providers.json
```

## Next Steps

- [CLI Reference](../reference/cli.md) - All commands and flags
- [GLM-5 Guide](../GLM-5-Guide.md) - GLM-5 specific configuration
