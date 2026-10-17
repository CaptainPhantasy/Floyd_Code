package model

import uv "github.com/charmbracelet/ultraviolet"

func splitVertical(area uv.Rectangle, height int) (top uv.Rectangle, bottom uv.Rectangle) {
	height = min(max(height, 0), area.Dy())
	top = uv.Rectangle{Min: area.Min, Max: uv.Position{X: area.Max.X, Y: area.Min.Y + height}}
	bottom = uv.Rectangle{Min: uv.Position{X: area.Min.X, Y: area.Min.Y + height}, Max: area.Max}
	return top, bottom
}

func splitHorizontal(area uv.Rectangle, width int) (left uv.Rectangle, right uv.Rectangle) {
	width = min(max(width, 0), area.Dx())
	left = uv.Rectangle{Min: area.Min, Max: uv.Position{X: area.Min.X + width, Y: area.Max.Y}}
	right = uv.Rectangle{Min: uv.Position{X: area.Min.X + width, Y: area.Min.Y}, Max: area.Max}
	return left, right
}
