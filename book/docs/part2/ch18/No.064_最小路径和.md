## 64 最小路径和-中等

题目：

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。


分析：

`dp[i][j]` 表示走到 `(i,j)` 的最小路径和。每个格子只能从「上」或「左」过来，取两者较小值再加上本格的值：

`dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])`

边界：第一行只能从左来（`dp[0][j] = dp[0][j-1] + grid[0][j]`），第一列只能从上来（`dp[i][0] = dp[i-1][0] + grid[i][0]`），起点 `dp[0][0] = grid[0][0]`。


```go
// date 2026/06/18
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    dp[0][0] = grid[0][0]
    for j := 1; j < n; j++ { // 第一行：只能从左来
        dp[0][j] = dp[0][j-1] + grid[0][j]
    }
    for i := 1; i < m; i++ { // 第一列：只能从上来
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[m-1][n-1]
}
```

> `min` 用 Go 1.21+ 内置函数即可；旧版本需自行定义 `func min(x, y int) int { if x < y { return x }; return y }`。


空间优化（滚动数组压缩到一维，空间 O(n)）：

与 No.063 同一个滚动技巧：`dp[j]` 更新前是「上方」`dp[i-1][j]`，`dp[j-1]` 更新后已是「左方」`dp[i][j-1]`，取较小者加 `grid[i][j]` 即可。

```go
// date 2026/06/18
func minPathSum1D(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, n)

    dp[0] = grid[0][0]
    for j := 1; j < n; j++ { // 第一行
        dp[j] = dp[j-1] + grid[0][j]
    }
    for i := 1; i < m; i++ {
        dp[0] += grid[i][0] // 第 1 列：只能从上方累加（dp[0] 更新前是上方值，更新后是本行值）
        for j := 1; j < n; j++ {
            dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
            //         ↑          ↑
            //        上方        左边
            //  dp[j]   本行尚未覆盖 → 仍是上一行的 dp[i-1][j]（来自上方）
            //  dp[j-1] 本行已覆盖   → 已是本行的   dp[i][j-1]（来自左边）
        }
    }
    return dp[n-1]
}
```

复杂度：时间 O(m·n)，空间 O(n)。
