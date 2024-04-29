package template

const MaxBit = 10

type TrieKth struct {
	child [MaxBit]*TrieKth
	cnt   int
	isNum bool
}

func NewTrieKth() *TrieKth {
	return &TrieKth{
		child: [MaxBit]*TrieKth{},
		isNum: false,
	}
}

func (t *TrieKth) Add(num string) {
	cur := t
	for i := 0; i < len(num); i++ {
		idx := num[i] - '0'
		if cur.child[idx] == nil {
			cur.child[idx] = NewTrieKth()
		}
		cur.child[idx].cnt++
		cur = cur.child[idx]
	}
	cur.isNum = true
}

func (t *TrieKth) FindTrieKth(k int) int {
	ans := 0
	path := make([]int, 0, 16)
	isFind := false
	var dfsNode func(root *TrieKth, target int)
	dfsNode = func(root *TrieKth, target int) {
		if root == nil {
			return
		}
		if target == 0 {
			// fmt.Printf("path = %d\n", path)
			isFind = true
			for _, v := range path {
				ans = ans*10 + v
			}
			return
		}
		for i := 0; i < MaxBit; i++ {
			if isFind {
				break
			}
			chd := root.child[i]
			if chd != nil {
				// fmt.Printf("child cnt %d, target %d\n", chd.cnt, target)
				if chd.cnt < target {
					target -= chd.cnt
				} else {
					path = append(path, i)
					if chd.isNum {
						target -= 1
					}
					dfsNode(chd, target)
					path = path[:len(path)-1]
				}
			}
		}
	}

	target := k
	dfsNode(t, target)
	return ans
}
