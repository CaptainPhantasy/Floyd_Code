package model

import (
	"image"
	"testing"
)

func TestSplitVertical(t *testing.T) {
	area := image.Rect(2, 3, 12, 13)
	leading, remainder := splitVertical(area, 4)

	if want := image.Rect(2, 3, 12, 7); leading != want {
		t.Fatalf("leading rectangle = %v, want %v", leading, want)
	}
	if want := image.Rect(2, 7, 12, 13); remainder != want {
		t.Fatalf("remainder rectangle = %v, want %v", remainder, want)
	}
}

func TestSplitHorizontal(t *testing.T) {
	area := image.Rect(2, 3, 12, 13)
	leading, remainder := splitHorizontal(area, 4)

	if want := image.Rect(2, 3, 6, 13); leading != want {
		t.Fatalf("leading rectangle = %v, want %v", leading, want)
	}
	if want := image.Rect(6, 3, 12, 13); remainder != want {
		t.Fatalf("remainder rectangle = %v, want %v", remainder, want)
	}
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
