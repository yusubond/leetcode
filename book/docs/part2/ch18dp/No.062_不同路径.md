## 62 不同路径-中等

题目：

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？



分析：

动归，直接计算 `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

注意边界条件。


```go
// date 2023/11/07
func uniquePaths(m int, n int) int {
  	// m row, n col
  	// dp[i][j] 走到坐标i,j的总解法
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
	// 0 col, just move down
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
	// 0 row, just move right
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}
```


空间优化（滚动数组压缩到一维，空间 O(n)）：

`dp[i][j] = dp[i-1][j] + dp[i][j-1]` 只依赖「上一行同列」与「本行前一列」，可用一维数组滚动：`dp[j]` 更新前是「上方」`dp[i-1][j]`，`dp[j-1]` 更新后已是「左方」`dp[i][j-1]`，于是 `dp[j] += dp[j-1]` 一句即可（与 No.063 同一套技巧，只是这里没有障碍物）。

```go
// date 2026/06/20
func uniquePaths1D(m int, n int) int {
    dp := make([]int, n)
    for j := 0; j < n; j++ { // 第一行：只有一种走法（一直向右）
        dp[j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] += dp[j-1]
            //   ↑      ↑
            // 上方 dp[i-1][j]（更新前）+ 左方 dp[i][j-1]（已更新）
            // j == 0 时 dp[0] 不变 = 1（第一列只有一种走法：一直向下）
        }
    }
    return dp[n-1]
}
```

复杂度：时间 O(m·n)，空间 O(n)。
