## 684 冗余连接-中等

题目：

树可以看成是一个连通且 **无环** 的 **无向** 图。

给定往一棵 `n` 个节点 (节点值 `1～n`) 的树中添加一条边后的图。添加的边的两个顶点包含在 `1` 到 `n` 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 `n` 的二维数组 `edges` ，`edges[i] = [ai, bi]` 表示图中在 `ai` 和 `bi` 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 `n` 个节点的树。如果有多个答案，则返回数组 `edges` 中最后出现的那个。



> **示例 1：**
>
> ![img](https://pic.leetcode-cn.com/1626676174-hOEVUL-image.png)
>
> ```
> 输入: edges = [[1,2], [1,3], [2,3]]
> 输出: [2,3]
> ```
>
> **示例 2：**
>
> ![img](https://pic.leetcode-cn.com/1626676179-kGxcmu-image.png)
>
> ```
> 输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
> 输出: [1,4]
> ```



**解题思路**

从题意可知，冗余的连接是指本来已经连通的节点，再加一条连接是意义的。

所以，这道题也用并查集解决。思路是把有边的节点合并，并且合并的时候做判断，如果合并失败（即两个节点的父结点相同，表示这两个节点已经通过其他边连通了），则说明这条边是冗余的。

```go
// date 2024/01/05
func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    uf := NewUnionFind(n+1)
    
    res := []int{}
    for _, v := range edges {
        if !uf.Union(v[0], v[1]) {
            res = []int{v[0], v[1]}
            return res
        }
    }

    return res
}

// 并查集，判断是否能够进行合并
type UnionFind struct {
    parent []int
    count int
}

func NewUnionFind(n int) *UnionFind {
    u := &UnionFind{
        parent: make([]int, n),
        count: n,
    }
    for i := 0; i < n; i++ {
        u.parent[i] = i
    }
    return u
}

func (u *UnionFind) Union(x, y int) bool {
    xp, yp := u.Find(x), u.Find(y)
    if xp == yp {
        return false
    }
    u.parent[yp] = xp
    u.count--
    return true
}

func (u *UnionFind) Find(x int) int {
    root := x
    for root != u.parent[root] {
        root = u.parent[root]
    }
    // 路径压缩
    for x != u.parent[x] {
        ox := u.parent[x]
        u.parent[x] = root
        x = ox
    }
    return root
}
```

