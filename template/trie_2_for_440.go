package template

import "fmt"

type Trie struct {
	child [MaxBit]*Trie
	isNum bool
}

func NewTrie() *Trie {
	return &Trie{
		child: [MaxBit]*Trie{},
		isNum: false,
	}
}

func (t *Trie) Add(num string) {
	cur := t
	for i := 0; i < len(num); i++ {
		idx := num[i] - '0'
		if cur.child[idx] == nil {
			cur.child[idx] = NewTrie()
		}
		cur = cur.child[idx]
	}
	cur.isNum = true
}

func (t *Trie) FindKth(k int) int {
	ans := 0
	cnt := 0
	path := make([]int, 0, 16)
	var dfsNode func(root *Trie)
	dfsNode = func(root *Trie) {
		if root == nil {
			return
		}
		if root.isNum {
			cnt++
			if cnt == k {
				fmt.Printf("path = %d\n", path)
				for _, v := range path {
					ans = ans*10 + v
				}
				return
			}
		}
		for i := 0; i < 10; i++ {
			if root.child[i] != nil {
				path = append(path, i)
				dfsNode(root.child[i])
				path = path[:len(path)-1]
			}
		}
	}

	dfsNode(t)

	return ans
}
