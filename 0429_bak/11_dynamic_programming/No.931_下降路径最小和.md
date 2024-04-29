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
