package terminal

import (
	"testing"
)

func TestEnqueueVTInputPreservesOrderAndDropsOnOverflow(t *testing.T) {
	component := &Component{
		inputQueue: make(chan []byte, 2),
		done:       make(chan struct{}),
	}
	first := []byte("first")
	if !component.enqueueVTInput(first) {
		t.Fatal("failed to enqueue first input")
	}
	first[0] = 'F'
	if !component.enqueueVTInput([]byte("second")) {
		t.Fatal("failed to enqueue second input")
	}
	if component.enqueueVTInput([]byte("overflow")) {
		t.Fatal("enqueued input into a full queue")
	}
	if got := string(<-component.inputQueue); got != "first" {
		t.Fatalf("first queued input = %q, want %q", got, "first")
	}
	if got := string(<-component.inputQueue); got != "second" {
		t.Fatalf("second queued input = %q, want %q", got, "second")
	}
}

func TestEnqueueVTInputStopsAfterShutdown(t *testing.T) {
	component := &Component{
		inputQueue: make(chan []byte, 1),
		done:       make(chan struct{}),
	}
	close(component.done)

	if component.enqueueVTInput([]byte("ignored")) {
		t.Fatal("enqueued input after shutdown")
	}
	if got := len(component.inputQueue); got != 0 {
		t.Fatalf("queued input count after shutdown = %d, want 0", got)
	}
}
