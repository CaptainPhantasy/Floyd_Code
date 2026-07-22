package model

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/legacy-ai/floyd/internal/message"
	"github.com/legacy-ai/floyd/internal/session"
	"github.com/legacy-ai/floyd/internal/ui/util"
)

// exportSession writes a complete local transcript for the selected session.
func (m *UI) exportSession(sessionID string) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		sess, err := m.com.App.Sessions.Get(ctx, sessionID)
		if err != nil {
			return util.ReportError(fmt.Errorf("get session: %w", err))()
		}
		msgs, err := m.com.App.Messages.List(ctx, sessionID)
		if err != nil {
			return util.ReportError(fmt.Errorf("list session messages: %w", err))()
		}

		root := m.com.Config().WorkingDir()
		if root == "" {
			root = "."
		}
		exportedAt := time.Now()
		path, err := writeSessionExport(root, renderSessionExport(sess, msgs, exportedAt), exportedAt)
		if err != nil {
			return util.ReportError(err)()
		}
		return util.NewInfoMsg("Session exported to: " + path)
	}
}

func writeSessionExport(root, content string, exportedAt time.Time) (string, error) {
	exportsDir := filepath.Join(root, ".floyd", "exports")
	if err := os.MkdirAll(exportsDir, 0o700); err != nil {
		return "", fmt.Errorf("create export directory: %w", err)
	}
	if err := os.Chmod(exportsDir, 0o700); err != nil {
		return "", fmt.Errorf("secure export directory: %w", err)
	}

	stem := "session-export-" + exportedAt.Format("2006-01-02-150405")
	for i := 0; i < 1000; i++ {
		name := stem + ".md"
		if i > 0 {
			name = fmt.Sprintf("%s-%d.md", stem, i)
		}
		path := filepath.Join(exportsDir, name)
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
		if errors.Is(err, os.ErrExist) {
			continue
		}
		if err != nil {
			return "", fmt.Errorf("create session export: %w", err)
		}
		if _, err := file.WriteString(content); err != nil {
			_ = file.Close()
			return "", fmt.Errorf("write session export: %w", err)
		}
		if err := file.Close(); err != nil {
			return "", fmt.Errorf("close session export: %w", err)
		}
		return path, nil
	}
	return "", fmt.Errorf("create session export: too many files share timestamp %s", exportedAt.Format(time.RFC3339))
}

func renderSessionExport(sess session.Session, msgs []message.Message, exportedAt time.Time) string {
	var out strings.Builder
	title := sess.Title
	if title == "" {
		title = "Untitled Session"
	}
	fmt.Fprintf(&out, "# Session Export - %s\n\n", title)
	fmt.Fprintf(&out, "**Exported**: %s\n", exportedAt.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(&out, "**Session ID**: %s\n", sess.ID)
	fmt.Fprintf(&out, "**Messages**: %d\n\n---\n\n", len(msgs))

	for _, msg := range msgs {
		fmt.Fprintf(&out, "## %s\n\n", exportRoleTitle(msg.Role))
		for _, part := range msg.Parts {
			renderExportPart(&out, part)
		}
		out.WriteString("---\n\n")
	}
	return out.String()
}

func exportRoleTitle(role message.MessageRole) string {
	switch role {
	case message.User:
		return "User"
	case message.Assistant:
		return "Assistant"
	case message.Tool:
		return "Tool Message"
	case message.System:
		return "System"
	default:
		return string(role)
	}
}

func renderExportPart(out *strings.Builder, part message.ContentPart) {
	switch p := part.(type) {
	case message.TextContent:
		if p.Text != "" {
			out.WriteString(p.Text)
			out.WriteString("\n\n")
		}
	case message.ReasoningContent:
		if p.Thinking != "" {
			out.WriteString("### Thinking\n\n")
			out.WriteString(p.Thinking)
			out.WriteString("\n\n")
		}
	case message.ToolCall:
		fmt.Fprintf(out, "### Tool Call: %s\n\n", p.Name)
		writeExportCodeBlock(out, "json", p.Input)
	case message.ToolResult:
		fmt.Fprintf(out, "### Tool Result: %s\n\n", p.Name)
		if p.IsError {
			out.WriteString("**Error**\n\n")
		}
		if p.Content != "" {
			writeExportCodeBlock(out, "", p.Content)
		}
		if p.Data != "" {
			out.WriteString("**Data**\n\n")
			writeExportCodeBlock(out, "", p.Data)
		}
	case message.ImageURLContent:
		fmt.Fprintf(out, "![Image](%s)\n\n", p.URL)
	case message.BinaryContent:
		fmt.Fprintf(out, "**Binary Content**: %s (%s, %d bytes)\n\n", p.Path, p.MIMEType, len(p.Data))
	case message.Finish:
		if p.Reason != "" {
			fmt.Fprintf(out, "*Finish reason: %s*\n\n", p.Reason)
		}
	}
}

func writeExportCodeBlock(out *strings.Builder, language, content string) {
	fence := "```"
	for strings.Contains(content, fence) {
		fence += "`"
	}
	fmt.Fprintf(out, "%s%s\n%s\n%s\n\n", fence, language, content, fence)
}
