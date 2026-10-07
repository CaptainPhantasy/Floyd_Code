package model

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/legacy-ai/floyd/internal/message"
	"github.com/legacy-ai/floyd/internal/session"
	"github.com/stretchr/testify/require"
)

func TestRenderSessionExportPreservesParts(t *testing.T) {
	exportedAt := time.Date(2026, time.July, 15, 16, 0, 0, 0, time.UTC)
	sess := session.Session{ID: "session-1", Title: "Recovery"}
	msgs := []message.Message{
		{Role: message.User, Parts: []message.ContentPart{message.TextContent{Text: "hello"}}},
		{Role: message.Assistant, Parts: []message.ContentPart{
			message.ReasoningContent{Thinking: "inspect first"},
			message.ToolCall{Name: "view", Input: "{\"path\":\"a.go\"}"},
			message.ToolResult{Name: "view", Content: "contains ``` fence"},
			message.Finish{Reason: message.FinishReasonEndTurn},
		}},
	}

	got := renderSessionExport(sess, msgs, exportedAt)
	for _, want := range []string{
		"# Session Export - Recovery",
		"**Session ID**: session-1",
		"**Messages**: 2",
		"## User",
		"hello",
		"### Thinking",
		"inspect first",
		"### Tool Call: view",
		"### Tool Result: view",
		"contains ``` fence",
		"*Finish reason: end_turn*",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("export missing %q", want)
		}
	}
	if !strings.Contains(got, "````\ncontains ``` fence\n````") {
		t.Error("tool output containing a Markdown fence was not safely enclosed")
	}
}

func TestWriteSessionExportUsesResolvedDataDirectory(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"custom-data", ".floyd"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dataDir := filepath.Join(t.TempDir(), name)
			path, err := writeSessionExport(dataDir, "transcript", time.Now())
			require.NoError(t, err)
			require.Equal(t, filepath.Join(dataDir, "exports"), filepath.Dir(path))
			content, err := os.ReadFile(path)
			require.NoError(t, err)
			require.Equal(t, "transcript", string(content))
		})
	}
}

func TestWriteSessionExportUsesPrivatePermissionsAndNoOverwrite(t *testing.T) {
	exportedAt := time.Date(2026, time.July, 15, 16, 0, 0, 0, time.UTC)
	root := t.TempDir()

	first, err := writeSessionExport(root, "first", exportedAt)
	if err != nil {
		t.Fatal(err)
	}
	second, err := writeSessionExport(root, "second", exportedAt)
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatal("second export overwrote the first")
	}

	info, err := os.Stat(first)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("export permissions = %o, want 600", got)
	}
	dirInfo, err := os.Stat(filepath.Dir(first))
	if err != nil {
		t.Fatal(err)
	}
	if got := dirInfo.Mode().Perm(); got != 0o700 {
		t.Fatalf("export directory permissions = %o, want 700", got)
	}
}
