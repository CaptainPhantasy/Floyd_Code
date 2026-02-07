# NEWFILES Alignment Audit

Date: 2026-02-07
Scope: /Volumes/Storage/floyd-main/NEWFILES → Floyd Go application
Note: This document lists intended features inferred from the provided NEWFILES folders and the features confirmed as implemented in the Go codebase. It does not include source code.

## AI (NEWFILES/src/ai)
Intended (from NEWFILES):
- Prompt templates and variable substitution
- Completion request structures, message types
- Client wrapper with defaults and streaming

Confirmed added:
- Prompt templates + render/list CLI (`floyd prompt list|render`)
- AI request/response/message types (CompletionRequest/Response, Message)
- AI client scaffold + streaming parser (Go implementation)
- AI dry-run command to emit a completion request JSON without calling providers

Notes:
- Dry-run now uses configured model defaults; errors if no model configured.

## Codebase (NEWFILES/src/codebase)
Intended:
- Codebase analysis (language counts, files, dependencies)
- Content search
- Dependency manifest parsing

Confirmed added:
- `floyd codebase analyze` (file counts, LOC, languages, directory summary)
- `floyd codebase deps` (manifest parsing)
- `floyd codebase search` (content search)
- Default ignore patterns include `.floyd`

## Commands (NEWFILES/src/commands)
Intended:
- CLI command registry for feature exposure

Confirmed added:
- New CLI surface: `exec`, `exec bg`, `file`, `codebase`, `prompt`, `ai dry-run`

## Config (NEWFILES/src/config)
Intended:
- Defaults + schema for new subsystems

Confirmed added:
- `options.file_ops` and `options.execution` config sections with defaults
- Execution allow/deny lists in config

## Errors (NEWFILES/src/errors)
Intended:
- Structured error types, formatting

Confirmed added:
- Error categories/types + formatter in Go

## Execution (NEWFILES/src/execution)
Intended:
- Shell command execution with safety
- Background jobs
- Allow/deny list controls

Confirmed added:
- Safe exec with dangerous-pattern blocking
- Allowlist/denylist via config and flags
- Exec env overrides, stderr capture, max buffer
- Background exec persisted across invocations (start/list/logs/kill)

## Fileops (NEWFILES/src/fileops)
Intended:
- Read/write/apply/diff/rename/copy/delete
- Append + temp files

Confirmed added:
- `floyd file` subcommands: read/info/ls/mkdir/write/append/apply/cp/mv/diff/rm
- `floyd file temp` (temp file creation)
- `floyd file find` (regex file search)

## FS (NEWFILES/src/fs)
Intended:
- Lower-level filesystem utilities (find/copy/rename/etc.)

Confirmed added:
- fsops helpers in Go (find, stream copy, validation)
- Exposed via `floyd file find`, `floyd file cp`, `floyd file mv`

## Telemetry (NEWFILES/src/telemetry)
Intended:
- Event capture and flush

Confirmed added:
- Telemetry manager with endpoint + batching
- CLI command run/success events tracked and flushed

## Scripts (NEWFILES/scripts)
Intended:
- Preinstall and helper scripts

Confirmed added:
- Not ported (no Go equivalent needed). Existing Go build/test flow maintained.

## Auth (NEWFILES/src/auth)
Intended:
- OAuth token management

Confirmed status:
- Not ported (explicitly requested: “no auth”).

---

## Items explicitly not added
- Additional auth beyond existing Floyd Go TUI
- UI/TUI changes for new commands

## Verification summary
- Build passes (`go build .`).
- Smoke tests run for new commands (exec allow/deny/env/stderr/buffer, exec bg, fileops, codebase, ai dry-run).

