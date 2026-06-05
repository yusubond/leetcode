## 64 最小路径和-中等

题目：

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。


分析：

```go
// date 2020/03/08
// 时间复杂度O(M*N),空间复杂度O(1)
func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 { return 0 }
    n := len(grid[0])
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            // 如果存在右节点和下节点，选择其中最小的那个
            if i + 1 < m && j + 1 < n {
                grid[i][j] += min(grid[i+1][j], grid[i][j+1])
            // 如果只有右节点，则选择右节点
            } else if j + 1 < n {
                grid[i][j] += grid[i][j+1]
            // 如果只有下节点，则选择下节点
            } else if i + 1 < m {
                grid[i][j] += grid[i+1][j]
            }
        }
    }
    return grid[0][0]
}

func min(x, y int) int {
  if x < y { return x }
  return y
}
```
