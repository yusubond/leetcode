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
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
	// 第 0 行，只能一直向右走，所以均为1
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
	// 第 0 列，只能一直向下走，所以均为1
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
