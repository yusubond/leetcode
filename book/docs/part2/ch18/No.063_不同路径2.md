## 63 不同路径2-中等

题目：

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。


分析：

直接递推。遇到障碍物，直接复用上一个值，或者变成零。

注意，初始化的时候，遇到障碍物也要变成零。


```go
// date 2023/11/11
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        if i == 0 {
            for j := 0; j < n; j++ {
                if obstacleGrid[i][j] == 1 {
                    dp[i][j] = 0
                } else {
                    dp[i][j] = 1
                }
            }
        }
    }
    for i := 0; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            dp[i][0] = 0
        } else {
            dp[i][0] = 1
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                continue
            }
            if i > 0 && j > 0 {
                if obstacleGrid[i-1][j] == 1 {
                    dp[i][j] = dp[i][j-1]
                } else if obstacleGrid[i][j-1] == 1 {
                    dp[i][j] = dp[i-1][j]
                } else if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
                    dp[i][j] = 0
                } else {
                    dp[i][j] = dp[i-1][j] + dp[i][j-1]
                }
            } else if i > 0 {
                if obstacleGrid[i-1][j] == 1 {
                    dp[i][j] = 0
                } else {
                    dp[i][j] = dp[i-1][j]
                }
            } else if j > 0 {
                if obstacleGrid[i][j-1] == 1 {
                    dp[i][j] = 0
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }
    return dp[m-1][n-1]
}
```
