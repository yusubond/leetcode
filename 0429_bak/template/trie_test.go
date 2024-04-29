package template

import (
	"fmt"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	keys := []string{"the", "this", "these", "answer", "any", "by", "bye", "pl", "plane"}
	tie := NewTrieNode()
	for _, v := range keys {
		tie.InsertKey(v)
	}
	s1 := "there"
	r1 := tie.SearchKey(s1)
	fmt.Printf("search %s in trie, res = %t\n", s1, r1)
	if r1 {
		t.Fatalf("fail")
	}

	s1 = "the"
	r1 = tie.SearchKey(s1)
	fmt.Printf("search %s in trie, res = %t\n", s1, r1)
	if !r1 {
		t.Fatalf("fail")
	}

	s1 = "any"
	r1 = tie.SearchKey(s1)
	fmt.Printf("search %s in trie, res = %t\n", s1, r1)
	if !r1 {
		t.Fatalf("fail")
	}

	s1 = "and"
	r1 = tie.SearchKey(s1)
	fmt.Printf("search %s in trie, res = %t\n", s1, r1)
	if r1 {
		t.Fatalf("fail")
	}

	// test delete
	s1 = "plane"
	r1 = tie.SearchKey(s1)
	if !r1 {
		t.Fatalf("fail search %s\n", s1)
	}
	tie.DeleteKey("plane")
	s1 = "pl"
	r1 = tie.SearchKey(s1)
	if !r1 {
		t.Fatalf("fail delete")
	}
}
