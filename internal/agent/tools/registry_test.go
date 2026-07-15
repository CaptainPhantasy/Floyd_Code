package tools

import (
	"iter"
	"strings"
	"testing"
	"time"

	"github.com/legacy-ai/floyd/internal/agent/tools/mcp"
)

func registrySequence(servers map[string][]*mcp.Tool) iter.Seq2[string, []*mcp.Tool] {
	return func(yield func(string, []*mcp.Tool) bool) {
		for server, tools := range servers {
			if !yield(server, tools) {
				return
			}
		}
	}
}

func TestBuildRegistryIsDeterministic(t *testing.T) {
	generatedAt := time.Date(2026, time.July, 15, 12, 0, 0, 0, time.FixedZone("test", -4*60*60))
	registry := buildRegistry(registrySequence(map[string][]*mcp.Tool{
		"floyd-terminal": {
			{Name: "start_process", Description: "Start a process"},
		},
		"floyd-devtools": {
			{Name: "typescript_semantic_analyzer", Description: "Analyze types"},
			{Name: "dependency_analyzer", Description: "Analyze dependencies"},
		},
	}), generatedAt)

	if registry.TotalServers != 2 || registry.TotalTools != 3 {
		t.Fatalf("unexpected totals: servers=%d tools=%d", registry.TotalServers, registry.TotalTools)
	}
	if registry.GeneratedAt != "2026-07-15T16:00:00Z" {
		t.Fatalf("unexpected generated timestamp: %s", registry.GeneratedAt)
	}
	wantOrder := []string{
		"floyd-devtools/dependency_analyzer",
		"floyd-devtools/typescript_semantic_analyzer",
		"floyd-terminal/start_process",
	}
	for index, entry := range registry.Tools {
		got := entry.Server + "/" + entry.Name
		if got != wantOrder[index] {
			t.Fatalf("entry %d: got %q, want %q", index, got, wantOrder[index])
		}
	}
	if got := strings.Join(registry.ByServer["floyd-devtools"], ","); got != "dependency_analyzer,typescript_semantic_analyzer" {
		t.Fatalf("server tools are not sorted: %s", got)
	}
}

func TestRegistryFormattingAndSearch(t *testing.T) {
	registry := &ToolRegistry{
		TotalTools:   3,
		TotalServers: 2,
		Tools: []RegistryEntry{
			{Name: "cache_retrieve", Server: "floyd-supercache", Description: "Retrieve cached data", Category: "cache"},
			{Name: "cache_store", Server: "floyd-supercache", Description: "Store cached data", Category: "cache"},
			{Name: "git_status", Server: "floyd-git", Description: "Read repository status", Category: "git"},
		},
		ByServer: map[string][]string{
			"floyd-supercache": {"cache_retrieve", "cache_store"},
			"floyd-git":        {"git_status"},
		},
	}

	compact := FormatCompact(registry)
	wantCompact := "Tool Registry: 3 tools from 2 servers\n  - floyd-git: 1 tools\n  - floyd-supercache: 2 tools\n"
	if compact != wantCompact {
		t.Fatalf("unexpected compact output:\n%s", compact)
	}

	detailed := FormatDetailed(registry)
	if strings.Index(detailed, "## floyd-git") > strings.Index(detailed, "## floyd-supercache") {
		t.Fatalf("detailed server sections are not sorted:\n%s", detailed)
	}

	results := SearchTools(registry, "CACHED")
	if len(results) != 2 || results[0].Name != "cache_retrieve" || results[1].Name != "cache_store" {
		t.Fatalf("unexpected search results: %#v", results)
	}
	if got := GetToolsByCategory(registry, "git"); len(got) != 1 || got[0].Name != "git_status" {
		t.Fatalf("unexpected category results: %#v", got)
	}
}

func TestCategorizeTool(t *testing.T) {
	tests := map[string]struct {
		name   string
		server string
		want   string
	}{
		"known server": {name: "anything", server: "floyd-safe-ops", want: "safety"},
		"lab lead":     {name: "anything", server: "lab-lead", want: "coordination"},
		"name prefix":  {name: "search_docs", server: "custom", want: "query"},
		"fallback":     {name: "anything", server: "custom", want: "other"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := categorizeTool(test.name, test.server); got != test.want {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}
