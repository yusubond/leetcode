## 547 省份数量-中等

题目：

有 `n` 个城市，其中一些彼此相连，另一些没有相连。如果城市 `a` 与城市 `b` 直接相连，且城市 `b` 与城市 `c` 直接相连，那么城市 `a` 与城市 `c` 间接相连。

**省份** 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 `n x n` 的矩阵 `isConnected` ，其中 `isConnected[i][j] = 1` 表示第 `i` 个城市和第 `j` 个城市直接相连，而 `isConnected[i][j] = 0` 表示二者不直接相连。

返回矩阵中 **省份** 的数量。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg)
>
> ```
> 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
> 输出：2
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg)
>
> ```
> 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
> 输出：3
> ```



**解题思路：**

这道题可以有多种解法。

- 并查集思想，详见解法一。思路是求连通分量的并查集，把有连通的城市合并，返回连通分量即可。
- DFS 或者 BFS，详见解法二。思路是利用 FloodFill 思想染色，采用逐行扫描的方式。每一行都是一个城市，如果它和其他城市联通，那么继续去扫描其他城市，直到扫描完成；这样和该城市所有的连通就是一个省份。



```go
// date 2024/01/04
// 解法一
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    uf := NewUnionFind(n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i != j && isConnected[i][j] == 1 {
                uf.Union(i, j)
            }
        }
    }

    return uf.GetCount()
}

// 并查集，求联通分量
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

func (u *UnionFind) Union(x, y int) {
    xp, yp := u.Find(x), u.Find(y)
    if xp == yp {
        return
    }
    u.parent[yp] = xp
    u.count--
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

func (u *UnionFind) GetCount() int {
    return u.count
}
```



```go
// date 2024/01/05
// 解法二
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make(map[int]bool, n)

    var dfs func(isConnected [][]int, cur int, vit map[int]bool)
    dfs = func(isConnected [][]int, cur int, vit map[int]bool) {
        vit[cur] = true
        for j := 0; j < len(isConnected[cur]); j++ {
            if !visited[j] && isConnected[cur][j] == 1 {
                dfs(isConnected, j, vit)
            }
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(isConnected, i, visited)
            count++
        }
    }

    return count
}
```

