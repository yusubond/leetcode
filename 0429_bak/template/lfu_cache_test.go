package template

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := NewLFUCache(2)
	fmt.Printf("put node 1, 1\n")
	lfu.Put(1, 1) // cnt(1) = 1
	fmt.Printf("put node 2, 2\n")
	lfu.Put(2, 2) // cnt(2) = 1
	lfu.ShowLFUNode()

	fmt.Printf("get node 1\n")
	res := lfu.Get(1) // cnt(1) = 2
	if res != 1 {
		t.Fatalf("fail")
	}
	lfu.ShowLFUNode()

	// push 3, delete 2
	fmt.Printf("put node 3, 3\n")
	lfu.Put(3, 3) // cnt(3) = 1
	res = lfu.Get(2)
	if res != -1 {
		t.Fatalf("fail")
	}
	lfu.ShowLFUNode()

	fmt.Printf("get node 3\n")
	res = lfu.Get(3) // cnt(3) = 2
	if res != 3 {
		t.Fatalf("fail")
	}
	lfu.ShowLFUNode()

	// push 4, delete 1
	fmt.Printf("put node 4, 4\n")
	lfu.Put(4, 4) // cnt(4) = 1
	res = lfu.Get(1)
	if res != -1 {
		t.Fatalf("fail")
	}
	lfu.ShowLFUNode()

	fmt.Printf("get node 3\n")
	res = lfu.Get(3)
	if res != 3 {
		t.Fatalf("fail")
	}
	fmt.Printf("get node 4\n")
	res = lfu.Get(4)
	if res != 4 {
		t.Fatalf("fail")
	}
	lfu.ShowLFUNode()
}
