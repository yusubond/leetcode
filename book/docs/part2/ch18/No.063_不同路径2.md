## 63 不同路径2-中等

题目：

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。


分析：

经典二维动态规划，状态转移：

`dp[i][j] = 0`，若 `grid[i][j] == 1`（障碍物）；
`dp[i][j] = dp[i-1][j] + dp[i][j-1]`，否则。

关键在初始化：第一行、第一列一旦遇到障碍物，其后所有格子都不可达，必须 `break`，而不是逐个判断为 0（否则障碍物之后的格子会被错误地设成 1）。


```go
// date 2026/06/18
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // 初始化第一行：遇到障碍物就 break，后面全为 0
    for j := 0; j < n; j++ {
        if obstacleGrid[0][j] == 1 {
            break
        }
        dp[0][j] = 1
    }
    // 初始化第一列：同理
    for i := 0; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[m-1][n-1]
}
```


空间优化（滚动数组压缩到一维，空间 O(n)）：

`dp[i][j]` 只依赖「上一行同列」和「本行前一列」，故可用一维数组滚动。设 `dp[0] = 1` 作为哨兵，简化第一行的 `dp[j] += dp[j-1]`；遇到障碍物直接置 0。

```go
// date 2026/06/18
func uniquePathsWithObstacles1D(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([]int, n)

    dp[0] = 1 // 哨兵：方便第一行 dp[j] += dp[j-1]
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
            } else if j > 0 {
                // dp[j](更新后) = dp[j](更新前) + dp[j-1]  <- 滚动的来源
                // dp[j-1] <- 左方
                // dp[j](更新前) <- 上方
                dp[j] += dp[j-1]
            }
            // j == 0 时，dp[j] 继承上一行的值（遇障碍已被置 0）
        }
    }
    return dp[n-1]
}
```

关于一维 `dp[j]` 的两点关键：

- **小结**：`dp[j]` 更新**前**是「上方」`dp[i-1][j]`，更新**后**变成「当前」`dp[i][j]`；而 `dp[j-1]` 因从左到右更新已是「左方」`dp[i][j-1]`。故一句 `dp[j] += dp[j-1]` 即 `dp[i][j] = dp[i-1][j] + dp[i][j-1]`。哨兵 `dp[0]=1` 是起点 `dp[0][0]=1` 的种子，并顺着第一列向下流。（方向绝不能反过来——从右到左会破坏这个对齐。）

- **`j == 0` 的特殊性**：第一列只能「从上方」来，没有「左方」，递推退化为 `dp[i][0] = dp[i-1][0]`。代码里 `j==0` 时两个分支都不进——非障碍就原封不动继承上一行的值，正好实现该递推；障碍则置 0。

复杂度：时间 O(m·n)，空间 O(n)。
