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
}
