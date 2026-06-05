package template

import "testing"

func TestMiddleStack(t *testing.T) {
	ms := NewMiddleStack()
	if !ms.IsEmpty() {
		t.Fatalf("fail expect empty")
	}

	mid, top := -1, -1
	ms.Push(1)
	// now stack: [1]
	mid = 1
	got := ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail, expect mid %d but got %d\n", mid, got)
	}
	top = 1
	got = ms.Pop()
	// now stack: []
	if got != top {
		t.Fatalf("fail expect top %d but got %d\n", top, got)
	}
	if !ms.IsEmpty() {
		t.Fatalf("fail expect empty")
	}

	ms.Push(1)
	ms.Push(2)
	// now stack: [1,2]
	mid = 1
	got = ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail expect mid %d but got %d\n", mid, got)
	}

	top = 2
	got = ms.Pop()
	// now stack: [1]
	if got != top {
		t.Fatalf("fail expect top %d but got %d\n", top, got)
	}

	// push 2,3,4
	ms.Push(2)
	ms.Push(3)
	// now stack: [1,2,3]
	mid = 2
	got = ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail expect mid %d but got %d\n", mid, got)
	}

	// delete mid
	ms.DeleteMiddle()
	// now stack: [1,3]
	mid = 1
	got = ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail expect mid %d but got %d\n", mid, got)
	}

	// now top: [1,3,4]
	ms.Push(4)
	mid = 3
	got = ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail expect mid %d but got %d\n", mid, got)
	}

	ms.Push(5)
	ms.Push(6)
	ms.Push(7)
	ms.DeleteMiddle()
	// now stack: [1,3,5,6,7]
	mid = 5
	got = ms.GetMiddle()
	if got != mid {
		t.Fatalf("fail expect mid %d but got %d\n", mid, got)
	}
	top = 7
	got = ms.Pop()
	if got != top {
		t.Fatalf("fail expect top %d but got %d\n", top, got)
	}
}
