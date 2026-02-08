# Floyd System Architecture

## Overview

Floyd is a Go-based AI coding agent built on the Fantasy framework. It provides session-based AI interactions with tool execution, streaming responses, and multi-provider LLM support.

## Core Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Floyd CLI                          │
├─────────────────────────────────────────────────────────────┤
│  Coordinator                                                │
│  ├── Manages agents (currently: single "coder" agent)       │
│  ├── Builds providers (Anthropic, OpenAI, Hyper, etc.)     │
│  └── Handles OAuth token refresh                            │
├─────────────────────────────────────────────────────────────┤
│  SessionAgent                                               │
│  ├── Message queue for concurrent requests                  │
│  ├── Auto-summarization at context thresholds               │
│  ├── Streaming with OnTextDelta, OnToolCall callbacks       │
│  └── Workaround for provider media limitations              │
├─────────────────────────────────────────────────────────────┤
│  Tools (fantasy.AgentTool)                                  │
│  ├── File ops: view, edit, write, multiedit                │
│  ├── Search: grep, glob, ls, sourcegraph                   │
│  ├── Execution: bash, job_output, job_kill                  │
│  ├── Network: fetch, download, web_fetch, web_search       │
│  └── LSP: diagnostics, references                           │
├─────────────────────────────────────────────────────────────┤
│  Provider Layer (catwalk + fantasy)                         │
│  ├── Anthropic (Claude with thinking support)              │
│  ├── OpenAI (Responses API with reasoning)                  │
│  ├── Google (Gemini with thinking_config)                   │
│  ├── OpenRouter (with exacto support)                       │
│  ├── Hyper (custom proxy provider)                          │
│  └── Bedrock, Azure, Vercel, OpenAI-compatible             │
└─────────────────────────────────────────────────────────────┘
```

## Key Components

### Coordinator (`internal/agent/coordinator.go`)

The central orchestration layer that:
- Creates and manages agents
- Builds LLM providers from configuration
- Handles provider selection and options merging
- Manages OAuth2 token refresh on 401 errors
- Coordinates tool building and filtering

### SessionAgent (`internal/agent/agent.go`)

The core agent implementation that:
- Manages session-based conversations
- Implements message queuing for concurrent requests
- Provides streaming responses with multiple callbacks
- Auto-summarizes when approaching context limits
- Handles tool execution and result processing
- Works around provider limitations (e.g., images in tool results)

### Tools (`internal/agent/tools/`)

All tools implement `fantasy.AgentTool` with:
- Name and description
- Parameter struct with JSON tags
- Handler function returning `fantasy.ToolResponse`
- Optional metadata response

### Prompt System (`internal/agent/prompt/`)

- Template-based system using Go templates
- Supports variable substitution
- Embeds templates at build time
- Generates system prompts from markdown templates

## Data Flow

```
User Prompt → Coordinator.Run()
                ↓
    SessionAgent.Run()
        ↓
    fantasy.Agent.Stream()
        ↓
    ┌───────────┬────────────┬──────────────┐
    │   OnText  │ OnToolCall │ OnReasoning  │
    │   Delta   │            │   Delta      │
    └───────────┴────────────┴──────────────┘
        ↓           ↓            ↓
  Message     Tool        Reasoning
  Update      Execution    Update
```

## Configuration

- **Models**: Large (for reasoning) + Small (for simple tasks)
- **Providers**: Multiple LLM providers with fallback support
- **Agents**: Currently single "coder" agent, extensible for more
- **Tools**: Configurable allow/deny lists per agent
- **MCP**: Model Context Protocol server integration

## Session Management

- Sessions stored in SQLite database
- Messages linked to sessions with role-based content
- Tool calls and results tracked separately
- Summary messages for context compression
- Usage tracking (tokens, cost)
