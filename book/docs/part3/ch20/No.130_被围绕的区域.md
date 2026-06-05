## 130 被围绕的区域-中等

题目：

给你一个 `m x n` 的矩阵 `board` ，由若干字符 `'X'` 和 `'O'` ，找到所有被 `'X'` 围绕的区域，并将这些区域里所有的 `'O'` 用 `'X'` 填充。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2021/02/19/xogrid.jpg)
>
> ```
> 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
> 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
> 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
> ```
>
> **示例 2：**
>
> ```
> 输入：board = [["X"]]
> 输出：[["X"]]
> ```



**解题思路**

这道题可以用并查集解决，但思路并不好想。可以这样思考，矩阵中所有的`O`，肯定是一片一片的存在矩阵中。

既然边缘上的`O`不可能被修改，那么跟边缘形成一片的`O`都不可以修改。

如此一来，我们对`O`的坐标进行并查集操作，具体是：

- 先把边缘的`O`都合作到一个特殊集合里
- 然后内部的`O`如果在四个方向有相邻，也合并起来【这一步主要是为了把跟边缘搭界的内部`O`合并到一起】



```go
// date 2024/01/04
func solve(board [][]byte)  {
    n := len(board)
    m := len(board[0])
    uf := newMyUnionFind(n*m+1)
    spec := n*m

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if (i == 0 || i == n-1 || j == 0 || j == m-1) && board[i][j] == 'O' {
                // 边缘上的 O 跟 n*m(特殊点) 合并
                uf.union(i*m+j, spec)
            } else if board[i][j] == 'O' {
                // 内部 O
                // 如果四个方向上有 O 直接合并
                if board[i-1][j] == 'O' {
                    uf.union(i*m+j, (i-1)*m+j)
                }
                if board[i+1][j] == 'O' {
                    uf.union(i*n+j, (i+1)*m+j)
                }
                if board[i][j-1] == 'O' {
                    uf.union(i*m+j, i*m+j-1)
                }
                if board[i][j+1] == 'O' {
                    uf.union(i*m+j, i*m+j+1)
                }
            }
        }
    }
    // 经过上面的合并，只要内部的O和边缘的O有相邻，那么内部的O一定会合并到边缘O的特殊集合里面
    // 剩下的不在特殊集合里面的，都可以被标记为 'X'
    specP := uf.find(spec)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if uf.find(i*m+j) != specP {
                board[i][j] = 'X'
            }
        }
    }
}

type (
	MyUnionFind struct {
		parent []int // 存储每个节点的父结点
	}
)

// n 表示图中一共有多少个节点
func newMyUnionFind(n int) *MyUnionFind {
	u := &MyUnionFind{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *MyUnionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
}

func (u *MyUnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}
```

