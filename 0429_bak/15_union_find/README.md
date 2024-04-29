## 并查集Union-Find

### 1、基本概念

并查集被认为是最简洁而优雅的数据结构之一，主要用于解决一些元素分组问题。

它管理一系列不相交的集合，并支持两种操作：

- 合并（Union）：把两个不相交的集合合并成一个集合
- 查询（Find）：查询两个元素是否在同一个集合中



并查集的应用场景很多，其中最直接的一个场景是：亲戚问题。

> 题目：
>
> 如果某个家族人员过于庞大，要判断两个人是否有亲戚关系，不是很容易，现在给出某个亲戚关系图，求任意给出的两个人是否具有亲戚关系。
>
> 规定，如果x和y是亲戚，y和z是亲戚，那么x和z也是亲戚。
>
> 规定，如果x和y是亲戚，那么x的亲戚都是y的亲戚，y的亲戚也都是x的亲戚。



### 2、代码设计

通常并查集的数据结构中有两个变量，parent和count。

Parent 表示节点的关系；count 表示连通分量。

初始化时，每个节点指向自己，连通分量就等于节点总数。

合并操作，先判断两个节点的父结点是否相同，如果相同表示已经在同一个集合了，直接返回；否则将其中一个节点的父结点更新为另一个的父结点。

查询操作，如果节点的父结点不是自己，那么一定要继续查找直到找到根节点。

```go
type (
	UnionFind struct {
		parent []int // 存储每个节点的父结点
		count  int   // 存储连通分量
	}
)

// NewUnionFind define
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		count:  n,
		parent: make([]int, n),
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
	u.parent[yp] = xp
	u.count--
}

// Find define
// search the root of x
func (u *UnionFind) Find(x int) int {
  // 包含了路径压缩
  // 递归写法
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

// GetCount define
// the total counts
func (u *UnionFind) GetCount() int {
	return u.count
}
```

### 3、路径压缩

从上面的代码可以看到，`find`函数很重要，即查找集合里的祖宗节点。但实际上，**我们并不关心集合里这棵树的结构长什么样子，只在乎根节点**。

因为无论树长什么样子，树上每个节点的根节点都是相同的，所以有没有进一步压缩每棵树的高度，使其高度始终保持为常数？

可以的，我们只需要在`find()`函数中增加如下代码，即在查找根节点的时候，顺便把整条链路上的节点的根节点都更新到根节点下面。

```go
func (u *UnionFind) Find(x int) int {
   root := x
   for root != u.parent[root] {
      root = u.parent[root]
   }
   //compress path
   // 迭代写法
   for x != u.parent[x] {
      temp := u.parent[x]
      u.parent[x] = root
      x = temp
   }
   return root
}
```



### 4、按秩合并

有了路径压缩，有些人可能会有一个误解，以为路径压缩优化以后，并查集始终都是一个菊花图（只有两层的树，俗称菊花图）。

但实际上，路径压缩只发生在查询时，而且也只压缩一条路径，所以并查集树的结构仍然可能比较复杂。

比如，现在我们有一棵较复杂的树和一个单元素的集合进行合并。

图片

如果我们可以选择的话，是把7的父结点设置为8好呢？还是把8的父结点设置为7好呢？

当然是后者。因为如果把7的父结点设置为8，那么整棵树的深度会增加，相应地每个元素都根节点的距离都变长了，查找操作也会变长。尽管有路径压缩，也会消耗时间。

而如果把8的父结点设置为7，则不会有这样的问题，因为这种操作方式并没有影响其他节点。

这就启发我们：**应该把简单的树往复杂的树上合并，而不是相反。**



我们用数组`rank`记录每棵树根节点对应的树深度（如果不是根节点，那么 rank 就是以它为根节点子树的深度）。

一开始，我们把`rank`（秩）设置为0，合并的时候比较两个根节点的秩，把秩更小的树往秩更大的树上面合并。



**合并的代码（按秩合并）**

```go
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
```

为什么深度相同，新的根节点深度要加1？

因为，本来都是根节点`root_x`和`root_y`，现在`root_x`变成了`root_y`的一个子节点，那么新的根节点`root_y`的深度当然要加1了。



## 题目赏析

第一类：并查集，求连通分量

这类并查集，可有路径压缩和秩序优化的版本，便于快速查找。

- 第 128 题：最长连续序列（也是求最大集合元素个数，思路不好想）
- 第 130 题：被围绕的区域（连通分量，特殊的连通）
- 第 323 题：无向图中连通分量的个数
- 第 547 题：省份的数量（就是求连通分量的个数）
