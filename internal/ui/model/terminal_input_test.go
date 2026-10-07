//go:build !windows

package model

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"github.com/legacy-ai/floyd/internal/ui/dialog"
	"github.com/legacy-ai/floyd/internal/ui/terminal"
	"github.com/stretchr/testify/require"
)

func TestFocusedTerminalReceivesEscape(t *testing.T) {
	dir := t.TempDir()
	shell := filepath.Join(dir, "read-key")
	require.NoError(t, os.WriteFile(shell, []byte(`#!/bin/sh
stty raw -echo
printf 'ready\n'
value=$(dd bs=1 count=1 2>/dev/null | od -An -tu1 | tr -d '[:space:]')
printf 'key:%s\n' "$value"
exec cat
`), 0o700))
	t.Setenv("SHELL", shell)
	term, read, err := terminal.New(24, 80, dir)
	require.NoError(t, err)
	t.Cleanup(term.Close)
	term.Focus()
	m := &UI{
		focus:    uiFocusTerminal,
		term:     term,
		keyMap:   DefaultKeyMap(),
		dialog:   dialog.NewOverlay(),
		textarea: textarea.New(),
	}

	output := make(chan string, 1)
	go func() {
		var received strings.Builder
		for read != nil {
			msg := read()
			if data, ok := msg.(terminal.OutputMsg); ok {
				received.Write([]byte(data))
				output <- received.String()
			}
			_, read = term.Update(msg)
			if _, closed := msg.(terminal.ClosedMsg); closed {
				return
			}
		}
	}()
	waitFor := func(want string) {
		t.Helper()
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		for {
			select {
			case received := <-output:
				if strings.Contains(received, want) {
					return
				}
			case <-timer.C:
				t.Fatalf("terminal did not produce %q", want)
			}
		}
	}
	waitFor("ready")
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	require.Equal(t, uiFocusTerminal, m.focus)
	require.True(t, term.Focused())
	waitFor("key:27")

	m.Update(tea.KeyPressMsg{Code: 't', Mod: tea.ModCtrl})
	require.Equal(t, uiFocusEditor, m.focus)
	require.False(t, term.Focused())
}
