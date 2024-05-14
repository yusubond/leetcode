package template

type (
	// UnionFind define
	// 路径压缩 + 秩优化
	UnionFind struct {
		parent []int // 存储每个节点的父结点
		rank   []int //
		count  int   // 存储连通分量
	}
)

// NewUnionFind define
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		count:  n,
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	// 初始化时, 每个节点的根节点都是自身
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// Union define
// merge x, y to the single set
func (u *UnionFind) Union(x, y int) {
	xp, yp := u.Find(x), u.Find(y)
	if xp == yp {
		return
	}
	// depth of tree_x < depth of tree_y
	if u.rank[xp] < u.rank[yp] {
		u.parent[xp] = yp
	} else {
		u.parent[yp] = xp
		// 如果深度相同，那么新的根节点秩加1
		if u.rank[xp] == u.rank[yp] {
			u.rank[xp]++
		}
	}
	u.count--
}

// Find define
// search the root of x
func (u *UnionFind) Find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//compress path
	for x != u.parent[x] {
		temp := u.parent[x]
		u.parent[x] = root
		x = temp
	}
	return root
}

// GetCount define
// the total counts
func (u *UnionFind) GetCount() int {
	return u.count
}

type (
	// UnionFindCount define
	// 计算每个集合中的元素个数，以及最大元素集合个数
	UnionFindCount struct {
		parent   []int
		count    []int
		maxCount int
	}
)

// NewUnionFindCount define
func NewUnionFindCount(n int) *UnionFindCount {
	c := &UnionFindCount{
		parent:   make([]int, n),
		count:    make([]int, n),
		maxCount: 1,
	}
	// init
	// 每个元素都是一个集合，根节点指向自身
	// 每个根节点的集合元素个数都是1
	for i := 0; i < n; i++ {
		c.parent[i] = i
		c.count[i] = 1
	}
	return c
}

// Union define
func (c *UnionFindCount) Union(x, y int) {
	xp, yp := c.Find(x), c.Find(y)
	if xp == yp {
		return
	}
	// always merge to n-1
	root := len(c.parent) - 1
	if xp == root {
		// xp is root
	} else if yp == root {
		xp, yp = yp, xp
	} else if c.count[xp] > c.count[yp] {
		xp, yp = yp, xp
	}
	// merge tree_y to tree_x
	c.parent[yp] = xp
	c.count[xp] += c.count[yp]
	if c.count[xp] > c.maxCount {
		c.maxCount = c.count[xp]
	}
}

// Find define
func (c *UnionFindCount) Find(x int) int {
	root := x
	for root != c.parent[root] {
		root = c.parent[root]
	}
	// compress path
	for x != c.parent[x] {
		ox := c.parent[x]
		c.parent[x] = root
		x = ox
	}
	return root
}

// MaxCount define
func (c *UnionFindCount) MaxCount() int {
	return c.maxCount
}
