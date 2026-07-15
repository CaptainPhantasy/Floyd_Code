package model

import (
	uv "github.com/charmbracelet/ultraviolet"
	uvlayout "github.com/charmbracelet/ultraviolet/layout"
)

// splitVertical divides area into a fixed-size leading segment and a remainder.
func splitVertical(area uv.Rectangle, leadingSize int) (leading, remainder uv.Rectangle) {
	uvlayout.Vertical(uvlayout.Len(leadingSize), uvlayout.Fill(1)).
		Split(area).
		Assign(&leading, &remainder)
	return leading, remainder
}

// splitHorizontal divides area into a fixed-size leading segment and a remainder.
func splitHorizontal(area uv.Rectangle, leadingSize int) (leading, remainder uv.Rectangle) {
	uvlayout.Horizontal(uvlayout.Len(leadingSize), uvlayout.Fill(1)).
		Split(area).
		Assign(&leading, &remainder)
	return leading, remainder
}
