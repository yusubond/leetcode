package template

import (
	"fmt"
	"testing"
)

func TestTrieKth_FindTrieKth(t *testing.T) {
	n := 13
	k := 2

	tie := NewTrieKth()
	for i := 1; i <= n; i++ {
		str := fmt.Sprintf("%d", i)
		tie.Add(str)
	}
	ans := tie.FindTrieKth(k)
	if ans != 10 {
		t.Fatalf("fail")
	}
}

func TestTrieKth_FindTrieKthBig(t *testing.T) {
	// this test case timeout on leetcode
	n := 7747794
	k := 5857460

	tie := NewTrieKth()
	for i := 1; i <= n; i++ {
		str := fmt.Sprintf("%d", i)
		tie.Add(str)
	}
	ans := tie.FindTrieKth(k)
	fmt.Printf("ans = %d\n", ans)
}
