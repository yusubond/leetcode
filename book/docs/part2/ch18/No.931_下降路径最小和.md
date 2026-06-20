## 931 下降路径最小和-中等

题目：

给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。


分析：

自底向上递推，注意求边界条件。

只有一列时，求所有行的和；只有一行时，求行内最小值。


```go
// date 2023/11/13
func minFallingPathSum(matrix [][]int) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    res := 0
    if m == 1 {
        for j := 0; j < len(matrix[0]); j++ {
            if j == 0 {
                res = matrix[0][0]
            } else {
                res = xy(res, matrix[0][j])
            }
        }
        return res
    }

    n := len(matrix[0])
    if n == 1 {
        res = 0
        for i := 0; i < m; i++ {
            res += matrix[i][0]
        }
        return res
    }

    for i := m-2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            if j > 0 && j+1 < n {
                matrix[i][j] += min(matrix[i+1][j-1], matrix[i+1][j], matrix[i+1][j+1])
            } else if j > 0 {
                matrix[i][j] += xy(matrix[i+1][j-1], matrix[i+1][j])
            } else if j + 1 < n {
                matrix[i][j] += xy(matrix[i+1][j], matrix[i+1][j+1])
            }
            if i == 0 {
                if j == 0 {
                    res = matrix[0][0]
                    continue
                } else if j > 0 {
                    res = xy(res, matrix[i][j])
                }
            }
        }
    }

    return res
}

func min(x, y, z int) int {
    return xy(xy(x, y), z)
}

func xy(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```


**重构版**：上面原实现用 `min`/`xy` 两套取小函数、并对单行/单列做特判，较为冗长。下面给出更干净的写法——**自顶向下**填表（单行/单列无需特判，DP 自然处理）。

干净二维（自顶向下）：`dp[i][j]` 表示从第一行走到 `(i,j)` 的下降路径最小和，从上一行 `j-1/j/j+1` 三列取最小再加本格：

```go
// date 2026/06/20
func minFallingPathSum2D(matrix [][]int) int {
    n := len(matrix)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    copy(dp[0], matrix[0]) // 第一行：起点
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            best := dp[i-1][j]                                  // 正上方
            if j > 0 && dp[i-1][j-1] < best { best = dp[i-1][j-1] }   // 左上
            if j+1 < n && dp[i-1][j+1] < best { best = dp[i-1][j+1] } // 右上
            dp[i][j] = matrix[i][j] + best
        }
    }
    res := dp[n-1][0]
    for j := 1; j < n; j++ {
        if dp[n-1][j] < res { res = dp[n-1][j] }
    }
    return res
}
```

空间优化（滚动数组压缩到一维，空间 O(n)）：

和编辑距离、LCS 一样，`dp[i][j]` 依赖上一行的 `j-1/j/j+1` 三个格子。从左到右滚动时，**左上 `dp[i-1][j-1]` 会被 `dp[j-1]` 覆盖丢失**，需用 `prev` 保存；而右上 `dp[j+1]` 尚未更新可直接读。

```go
// date 2026/06/20
func minFallingPathSum1D(matrix [][]int) int {
    n := len(matrix)
    dp := make([]int, n)
    copy(dp, matrix[0])
    for i := 1; i < n; i++ {
        prev := dp[0] // 上一行 j-1（左上）
        for j := 0; j < n; j++ {
            temp := dp[j]                                   // 上一行 j（上方），先存下来
            best := dp[j]                                   // 正上方
            if j > 0 && prev < best { best = prev }         // 左上（上一行 j-1）
            if j+1 < n && dp[j+1] < best { best = dp[j+1] } // 右上（尚未覆盖）
            dp[j] = matrix[i][j] + best
            prev = temp                                     // 下一列的左上 = 本列的上方
        }
    }
    res := dp[0]
    for j := 1; j < n; j++ {
        if dp[j] < res { res = dp[j] }
    }
    return res
}
```

复杂度：时间 O(n²)，空间 O(n)。
