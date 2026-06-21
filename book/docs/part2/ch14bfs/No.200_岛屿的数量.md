## 200 岛屿的数量-中等

**题目**：

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



> **示例 1：**
>
> ```
> 输入：grid = [
>   ["1","1","1","1","0"],
>   ["1","1","0","1","0"],
>   ["1","1","0","0","0"],
>   ["0","0","0","0","0"]
> ]
> 输出：1
> ```
>
> **示例 2：**
>
> ```
> 输入：grid = [
>   ["1","1","0","0","0"],
>   ["1","1","0","0","0"],
>   ["0","0","1","0","0"],
>   ["0","0","0","1","1"]
> ]
> 输出：3
> ```



**分析：**

这道题是图的 BFS 算法，基本思路是：

1. 通过 visited 标记已经访问过的点
2. 二层遍历，一旦找到没有访问过，且为`1`的点，结果递增；并且以此点为起点开始四个方向的 BFS

```go
// date 2023/12/28
func numIslands(grid [][]byte) int {
    n := len(grid)
    if n == 0 {
        return 0
    }
    m := len(grid[0])
    if m == 0 {
        return 0
    }

    dir := [][]int{
        []int{1,0},  // x = x+1, y = y
        []int{-1,0}, // x = x-1, y = y
        []int{0,1},  // x = x, y = y-1
        []int{0,-1}, // x = x, y = y+1
    }
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }

    var isInBoard func(x, y int) bool
    isInBoard = func(x, y int) bool {
        return x >= 0 && x < n && y >= 0 && y < m
    }

    // search land
    var bfs func(x, y int)
    bfs = func(x, y int) {
        visited[x][y] = true
        for i := 0; i < 4; i++ {
            nx := x + dir[i][0]
            ny := y + dir[i][1]
            // 继续搜索联通的 1
            if isInBoard(nx, ny) && !visited[nx][ny] && grid[nx][ny] == '1' {
                bfs(nx, ny)
            }
        }
    }

    res := 0
    // 依次遍历查找, 如果为 1 且没有被访问过，那么就是新大陆
    // 在新大陆开启四个方向的 BFS 搜索
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                bfs(i, j)
                res++
            }
        }
    }

    return res
}
```

