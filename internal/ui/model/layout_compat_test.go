package model

import (
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/stretchr/testify/require"
)

func TestSplitCompat(t *testing.T) {
	t.Parallel()

	area := uv.Rectangle{Min: uv.Position{X: 10, Y: 20}, Max: uv.Position{X: 110, Y: 70}}

	top, bottom := splitVertical(area, 15)
	require.Equal(t, uv.Rectangle{Min: area.Min, Max: uv.Position{X: 110, Y: 35}}, top)
	require.Equal(t, uv.Rectangle{Min: uv.Position{X: 10, Y: 35}, Max: area.Max}, bottom)

	left, right := splitHorizontal(area, 40)
	require.Equal(t, uv.Rectangle{Min: area.Min, Max: uv.Position{X: 50, Y: 70}}, left)
	require.Equal(t, uv.Rectangle{Min: uv.Position{X: 50, Y: 20}, Max: area.Max}, right)
}

func TestSplitCompatClamps(t *testing.T) {
	t.Parallel()

	area := uv.Rectangle{Min: uv.Position{X: 2, Y: 3}, Max: uv.Position{X: 12, Y: 13}}

	top, bottom := splitVertical(area, -1)
	require.Equal(t, 0, top.Dy())
	require.Equal(t, area, bottom)

	left, right := splitHorizontal(area, 20)
	require.Equal(t, area, left)
	require.Equal(t, 0, right.Dx())
}
