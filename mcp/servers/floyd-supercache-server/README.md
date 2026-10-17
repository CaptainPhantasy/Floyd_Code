# Floyd Supercache MCP Server

A local MCP server that stores JSON-serializable values in three persistent tiers:

- `project` for short-lived working data
- `reasoning` for durable decision context
- `vault` for archived patterns and completed reasoning

The server exposes twelve cache tools plus MCP resources that describe its tool registry and health.

## Trust boundary

Run this server only as a trusted local STDIO process. Cache tools can create, read, and delete data under the configured cache directory. Do not expose the process directly over HTTP or to untrusted callers.

Do not place credentials, secrets, private customer data, or hidden reasoning in this cache. The server protects local cache files with owner-only permissions, but it is not an encrypted secret store.

## Install and validate

```sh
npm ci
npm test
npm audit --audit-level=low
```

## Configure an MCP client

Build the package, then configure the client to run the compiled server with an absolute path:

```json
{
  "mcpServers": {
    "floyd-supercache": {
      "command": "node",
      "args": ["/absolute/path/to/floyd-supercache-server/dist/index.js"]
    }
  }
}
```

By default, data is stored under `~/.floyd/supercache`. Set `FLOYD_SUPERCACHE_DIR` to use another directory. The integration test uses that setting to keep test data out of the operator's real cache.

## Tools

- `cache_store`, `cache_retrieve`, `cache_delete`, and `cache_clear`
- `cache_list`, `cache_search`, `cache_stats`, and `cache_prune`
- `cache_store_pattern`
- `cache_store_reasoning`, `cache_load_reasoning`, and `cache_archive_reasoning`

Destructive tools require an explicit action: `cache_clear` requires `confirm=true`, and `cache_prune` supports `dryRun=true` for previewing removals.
