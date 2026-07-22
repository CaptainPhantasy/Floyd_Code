// Package tools provides tool registry functionality for boot-time discovery.
package tools

import (
	"encoding/json"
	"fmt"
	"iter"
	"log/slog"
	"sort"
	"strings"
	"time"

	"github.com/legacy-ai/floyd/internal/agent/tools/mcp"
)

// RegistryEntry represents a single tool in the registry.
type RegistryEntry struct {
	Name        string `json:"name"`
	Server      string `json:"server"`
	Description string `json:"description"`
	Category    string `json:"category,omitempty"`
}

// ToolRegistry represents the complete tool registry.
type ToolRegistry struct {
	TotalTools   int                 `json:"total_tools"`
	TotalServers int                 `json:"total_servers"`
	Tools        []RegistryEntry     `json:"tools"`
	ByServer     map[string][]string `json:"by_server"`
	ByCategory   map[string][]string `json:"by_category"`
	Version      string              `json:"version"`
	GeneratedAt  string              `json:"generated_at"`
}

// BuildRegistry builds the complete tool registry from all available MCP tools.
func BuildRegistry() *ToolRegistry {
	return buildRegistry(mcp.Tools(), time.Now())
}

func buildRegistry(toolSets iter.Seq2[string, []*mcp.Tool], generatedAt time.Time) *ToolRegistry {
	registry := &ToolRegistry{
		ByServer:    make(map[string][]string),
		ByCategory:  make(map[string][]string),
		Tools:       make([]RegistryEntry, 0),
		Version:     "1.0",
		GeneratedAt: generatedAt.UTC().Format(time.RFC3339),
	}

	serversSeen := make(map[string]bool)

	// Iterate through all MCP tools
	for mcpName, tools := range toolSets {
		serversSeen[mcpName] = true
		var serverTools []string

		for _, tool := range tools {
			entry := RegistryEntry{
				Name:        tool.Name,
				Server:      mcpName,
				Description: tool.Description,
			}

			// Determine category
			entry.Category = categorizeTool(tool.Name, mcpName)

			registry.Tools = append(registry.Tools, entry)
			serverTools = append(serverTools, tool.Name)

			// By category
			registry.ByCategory[entry.Category] = append(
				registry.ByCategory[entry.Category],
				fmt.Sprintf("%s/%s", mcpName, tool.Name),
			)
		}

		registry.ByServer[mcpName] = serverTools
	}

	registry.TotalTools = len(registry.Tools)
	registry.TotalServers = len(serversSeen)
	sort.Slice(registry.Tools, func(i, j int) bool {
		if registry.Tools[i].Server == registry.Tools[j].Server {
			return registry.Tools[i].Name < registry.Tools[j].Name
		}
		return registry.Tools[i].Server < registry.Tools[j].Server
	})
	for _, tools := range registry.ByServer {
		sort.Strings(tools)
	}
	for _, tools := range registry.ByCategory {
		sort.Strings(tools)
	}

	return registry
}

// categorizeTool determines the category of a tool based on its name and server.
func categorizeTool(name, server string) string {
	switch server {
	case "floyd-supercache":
		return "cache"
	case "floyd-terminal":
		return "terminal"
	case "floyd-devtools":
		return "development"
	case "floyd-safe-ops":
		return "safety"
	case "floyd-git":
		return "git"
	case "floyd-runner":
		return "testing"
	case "floyd-patch":
		return "editing"
	case "lab-lead":
		return "coordination"
	}

	// Extract category from server name if possible
	if idx := strings.LastIndex(server, "-"); idx != -1 {
		server = server[idx+1:]
	}

	switch server {
	case "supercache":
		return "cache"
	case "terminal":
		return "terminal"
	case "devtools":
		return "development"
	case "git":
		return "git"
	case "runner":
		return "testing"
	case "patch":
		return "editing"
	}

	// Fallback: categorize by tool name prefix
	prefixes := map[string]string{
		"cache_":  "cache",
		"git_":    "git",
		"edit_":   "editing",
		"run_":    "testing",
		"format_": "formatting",
		"lint_":   "linting",
		"build_":  "build",
		"test_":   "testing",
		"start_":  "terminal",
		"create_": "creation",
		"list_":   "query",
		"get_":    "query",
		"search_": "query",
		"find_":   "query",
	}

	for prefix, category := range prefixes {
		if len(name) > len(prefix) && name[:len(prefix)] == prefix {
			return category
		}
	}

	return "other"
}

// FormatCompact returns a compact format suitable for inline inclusion in prompts.
func FormatCompact(registry *ToolRegistry) string {
	// Count tools per server
	serverCounts := make(map[string]int)
	for name := range registry.ByServer {
		serverCounts[name] = len(registry.ByServer[name])
	}

	// Build compact summary
	result := fmt.Sprintf("Tool Registry: %d tools from %d servers\n", registry.TotalTools, registry.TotalServers)
	servers := sortedKeys(serverCounts)
	for _, server := range servers {
		count := serverCounts[server]
		result += fmt.Sprintf("  - %s: %d tools\n", server, count)
	}

	return result
}

// FormatDetailed returns a detailed format with all tools listed.
func FormatDetailed(registry *ToolRegistry) string {
	result := fmt.Sprintf("# Tool Registry\n%d tools from %d servers\n\n", registry.TotalTools, registry.TotalServers)

	for _, server := range sortedKeys(registry.ByServer) {
		tools := registry.ByServer[server]
		result += fmt.Sprintf("## %s (%d tools)\n", server, len(tools))
		for _, tool := range tools {
			// Find description
			for _, entry := range registry.Tools {
				if entry.Name == tool && entry.Server == server {
					result += fmt.Sprintf("  - %s: %s\n", entry.Name, entry.Description)
					break
				}
			}
		}
		result += "\n"
	}

	return result
}

// FormatMCPConfig returns MCP_CONFIG.json format.
func FormatMCPConfig(registry *ToolRegistry) string {
	// This would return the MCP server configuration format
	return FormatCompact(registry)
}

// GetToolsByServer returns all tools from a specific server.
func GetToolsByServer(registry *ToolRegistry, server string) []RegistryEntry {
	var result []RegistryEntry
	for _, entry := range registry.Tools {
		if entry.Server == server {
			result = append(result, entry)
		}
	}
	return result
}

// GetToolsByCategory returns all tools in a specific category.
func GetToolsByCategory(registry *ToolRegistry, category string) []RegistryEntry {
	var result []RegistryEntry
	for _, entry := range registry.Tools {
		if entry.Category == category {
			result = append(result, entry)
		}
	}
	return result
}

// SearchTools searches for tools by name or description.
func SearchTools(registry *ToolRegistry, query string) []RegistryEntry {
	query = lower(query)
	var result []RegistryEntry
	for _, entry := range registry.Tools {
		if contains(lower(entry.Name), query) || contains(lower(entry.Description), query) {
			result = append(result, entry)
		}
	}
	return result
}

func lower(s string) string {
	return strings.ToLower(s)
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// BootToolRegistry builds and returns the tool registry at boot time.
// This should be called during agent initialization.
func BootToolRegistry() (*ToolRegistry, error) {
	slog.Info("Building tool registry at boot")
	registry := BuildRegistry()
	slog.Info("Tool registry built",
		"total_tools", registry.TotalTools,
		"total_servers", registry.TotalServers,
	)
	return registry, nil
}

// BootToolRegistryJSON returns the tool registry as JSON.
func BootToolRegistryJSON() (string, error) {
	registry, err := BootToolRegistry()
	if err != nil {
		return "", err
	}
	data, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal registry: %w", err)
	}
	return string(data), nil
}

// BootSummary returns a one-line summary suitable for boot logging.
func BootSummary() string {
	registry, err := BootToolRegistry()
	if err != nil {
		return fmt.Sprintf("Tool registry: unavailable (%v)", err)
	}
	return fmt.Sprintf("Tools: %d from %d servers", registry.TotalTools, registry.TotalServers)
}

// ListServers returns a list of all MCP servers that provided tools.
func ListServers() []string {
	var servers []string
	for mcpName := range mcp.Tools() {
		servers = append(servers, mcpName)
	}
	sort.Strings(servers)
	return servers
}

// CountToolsByServer returns the count of tools for each server.
func CountToolsByServer() map[string]int {
	counts := make(map[string]int)
	for mcpName, tools := range mcp.Tools() {
		counts[mcpName] = len(tools)
	}
	return counts
}

// GetToolNames returns all tool names in the format "server_tool".
func GetToolNames() []string {
	var names []string
	for mcpName, tools := range mcp.Tools() {
		for _, tool := range tools {
			names = append(names, fmt.Sprintf("%s_%s", mcpName, tool.Name))
		}
	}
	sort.Strings(names)
	return names
}

func sortedKeys[V any](values map[string]V) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// IterateTools iterates over all tools, calling the provided function.
func IterateTools(fn func(mcpName string, tool *mcp.Tool) bool) {
	for mcpName, tools := range mcp.Tools() {
		for _, tool := range tools {
			if !fn(mcpName, tool) {
				return
			}
		}
	}
}

// ToolCount returns the total number of available tools.
func ToolCount() int {
	count := 0
	for _, tools := range mcp.Tools() {
		count += len(tools)
	}
	return count
}

// ServerCount returns the number of MCP servers.
func ServerCount() int {
	count := 0
	for range mcp.Tools() {
		count++
	}
	return count
}
