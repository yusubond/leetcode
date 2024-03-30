package template

// for 677
type TrieSumNode struct {
	child [26]*TrieSumNode
	sum   int
}

type MapSum struct {
	cache map[string]int
	root  *TrieSumNode
}

func newTrieSumNode() *TrieSumNode {
	t := &TrieSumNode{
		child: [26]*TrieSumNode{},
		sum:   0,
	}
	return t
}

func (t *TrieSumNode) insertKey(key string, delt int) {
	cur, n := t, len(key)
	for i := 0; i < n; i++ {
		idx := key[i] - 'a'
		if cur.child[idx] == nil {
			cur.child[idx] = newTrieSumNode()
		}
		cur.child[idx].sum += delt
		cur = cur.child[idx]
	}
}

func (t *TrieSumNode) searchPrefixSum(s string) int {
	i, n := 0, len(s)
	cur := t
	for i = 0; i < n; i++ {
		idx := s[i] - 'a'
		if cur.child[idx] != nil {
			cur = cur.child[idx]
		} else {
			break
		}
	}
	if i < n || cur == nil {
		return 0
	}
	return cur.sum
}
