package template

import "testing"

func TestFixedSizeStack(t *testing.T) {
	sk := NewFixedSizeStack(1024)
	sk.Push(1)
	sk.Push(2)
	sk.Push(3)
	sk.Push(4)

	sk.Show()

	// reverse stack
	sk.Reverse()

	sk.Show()
}