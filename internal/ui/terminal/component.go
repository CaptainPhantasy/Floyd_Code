package terminal

import (
	"io"

	tea "charm.land/bubbletea/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/vt"
)

// OutputMsg carries a chunk of PTY output into the bubbletea event
// loop so we can feed it to the VT emulator on the main goroutine.
type OutputMsg []byte

// ClosedMsg signals that the PTY output stream ended (shell exited).
type ClosedMsg struct{}

// Component is the embedded terminal UI component. It owns a PTY
// session and a virtual terminal emulator and can draw itself into an
// ultraviolet screen region.
type Component struct {
	session *Session
	vt      *vt.SafeEmulator
	focused bool
	width   int
	height  int
}

// New creates a terminal component. It spawns the shell and begins
// reading PTY output.
func New(rows, cols int, cwd string) (*Component, tea.Cmd, error) {
	sess, err := NewSession(rows, cols, cwd)
	if err != nil {
		return nil, nil, err
	}

	emu := vt.NewSafeEmulator(cols, rows)

	c := &Component{
		session: sess,
		vt:      emu,
		width:   cols,
		height:  rows,
	}

	// Kick off the reader goroutine that bridges PTY → bubbletea.
	cmd := c.readLoop()

	return c, cmd, nil
}

// readLoop returns a tea.Cmd that reads the next chunk from the PTY
// and sends it as an outputMsg. It re-schedules itself after each
// read.
func (c *Component) readLoop() tea.Cmd {
	return func() tea.Msg {
		buf := make([]byte, 4096)
		n, err := c.session.Read(buf)
		if err != nil {
			if err == io.EOF {
				return ClosedMsg{}
			}
			return ClosedMsg{}
		}
		data := make([]byte, n)
		copy(data, buf[:n])
		return OutputMsg(data)
	}
}

// teaKeyToUV converts a bubbletea KeyPressMsg into an ultraviolet
// KeyPressEvent. Both types share the same field layout; the
// conversion copies each field explicitly.
func teaKeyToUV(msg tea.KeyPressMsg) uv.KeyPressEvent {
	k := tea.Key(msg)
	return uv.KeyPressEvent(uv.Key{
		Text:        k.Text,
		Mod:         uv.KeyMod(k.Mod),
		Code:        k.Code,
		ShiftedCode: k.ShiftedCode,
		BaseCode:    k.BaseCode,
		IsRepeat:    k.IsRepeat,
	})
}

// Update processes messages relevant to the terminal component.
// Returns true if the message was handled, plus any follow-up command.
func (c *Component) Update(msg tea.Msg) (bool, tea.Cmd) {
	switch msg := msg.(type) {
	case OutputMsg:
		// Feed PTY data into the VT emulator.
		c.vt.Write([]byte(msg)) //nolint:errcheck
		// Keep reading.
		return true, c.readLoop()

	case ClosedMsg:
		// Shell exited.
		return true, nil

	case tea.KeyPressMsg:
		if !c.focused {
			return false, nil
		}
		// Forward key events to the VT emulator, which converts them
		// into the correct escape sequences.
		c.vt.SendKey(teaKeyToUV(msg))
		// Flush VT input to the PTY in a background goroutine to
		// avoid blocking the UI thread if the PTY buffer is full.
		return true, c.flushVTInputCmd()

	case tea.PasteMsg:
		if !c.focused {
			return false, nil
		}
		c.vt.Paste(msg.Content)
		return true, c.flushVTInputCmd()

	case tea.WindowSizeMsg:
		if c.focused {
			c.Resize(msg.Width, msg.Height)
			return true, nil
		}
	}

	return false, nil
}

// flushVTInputCmd returns a tea.Cmd that reads pending input from the
// VT emulator and writes it to the PTY in a background goroutine.
// This prevents the main UI thread from blocking when the PTY buffer
// is full.
func (c *Component) flushVTInputCmd() tea.Cmd {
	return func() tea.Msg {
		buf := make([]byte, 4096)
		for {
			n, err := c.vt.Read(buf)
			if n > 0 {
				c.session.Write(buf[:n]) //nolint:errcheck
			}
			if n == 0 || err != nil {
				break
			}
		}
		return nil
	}
}

// Resize changes both the VT emulator and the PTY window size.
func (c *Component) Resize(width, height int) {
	if width == c.width && height == c.height {
		return
	}
	c.width = width
	c.height = height
	c.vt.Resize(width, height)
	c.session.Resize(height, width)
}

// Draw renders the virtual terminal into the given screen area.
func (c *Component) Draw(scr uv.Screen, area uv.Rectangle) {
	c.vt.Draw(scr, area)
}

// Focus gives input focus to the terminal.
func (c *Component) Focus() {
	c.focused = true
	c.vt.Emulator.Focus()
}

// Blur removes input focus from the terminal.
func (c *Component) Blur() {
	c.focused = false
	c.vt.Emulator.Blur()
}

// Focused reports whether the terminal has input focus.
func (c *Component) Focused() bool {
	return c.focused
}

// Close tears down the PTY session and VT emulator.
func (c *Component) Close() {
	if c.session != nil {
		c.session.Close() //nolint:errcheck
	}
	if c.vt != nil {
		c.vt.Close() //nolint:errcheck
	}
}

// Active reports whether the terminal session is still running.
func (c *Component) Active() bool {
	return c.session != nil && !c.session.Closed()
}
