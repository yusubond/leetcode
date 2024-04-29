## 323 无向图中连通分量的个数-中等

题目：

你有一个包含 `n` 个节点的图。给定一个整数 `n` 和一个数组 `edges` ，其中 `edges[i] = [ai, bi]` 表示图中 `ai` 和 `bi` 之间有一条边。

返回 *图中已连接分量的数目* 。



> **示例 1:**
>
> ![img](https://assets.leetcode.com/uploads/2021/03/14/conn1-graph.jpg)
>
> ```
> 输入: n = 5, edges = [[0, 1], [1, 2], [3, 4]]
> 输出: 2
> ```
>
> **示例 2:**
>
> ![img](https://assets.leetcode.com/uploads/2021/03/14/conn2-graph.jpg)
>
> ```
> 输入: n = 5, edges = [[0,1], [1,2], [2,3], [3,4]]
> 输出:  1
> ```



分析：

算法1：

这道题跟第 200 题【岛屿的数量】类似，可以用 DFS 搜索，具体为：

- 先讲边的关系构造成有出向节点的图
- 开始遍历，通过 visited 标记已经遍历过的，一旦发现没有遍历过的节点，计数增加1

```go
// date 2024/01/04
func countComponents(n int, edges [][]int) int {
    // out
    out := make(map[int][]int, 4)
    visited := make(map[int]bool, n)
    for _, v := range edges {
        v0, v1 := v[0], v[1]
        _, ok := out[v0]
        if !ok {
            out[v0] = []int{}
        }
        out[v0] = append(out[v0], v1)
        _, ok = out[v1]
        if !ok {
            out[v1] = []int{}
        }
        out[v1] = append(out[v1], v0)
    }
    count := 0

    var dfs func(start int)
    dfs = func(start int) {
        if visited[start] {
            return
        }
        visited[start] = true
        if ov, ok := out[start]; ok && len(ov) != 0 {
            for _, v := range ov {
                dfs(v)
            }
        }
    }

    for i := 0; i < n; i++ {
        if !visited[i] {
            count++
            dfs(i)
        }
    }

    return count
}
```



算法2：

这道题也是基本的并查集算法。

```go
// date 2024/01/04
func countComponents(n int, edges [][]int) int {
    uf := newMyUnionFind(n)
    for _, v := range edges {
        uf.union(v[0], v[1])
    }
    return uf.getCount()
}

type (
	MyUnionFind struct {
		parent []int // 存储每个节点的父结点
		count  int   // 存储连通分量
	}
)

// n 表示图中一共有多少个节点
func newMyUnionFind(n int) *MyUnionFind {
	u := &MyUnionFind{
		count:  n,
		parent: make([]int, n),
	}
	// 初始化时, 每个节点的父结点都是自己
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// 将两个节点x y 合并
func (u *MyUnionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
	u.count--
}

// 查找 x 的父结点
func (u *MyUnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *MyUnionFind) getCount() int {
	return u.count
}
```

