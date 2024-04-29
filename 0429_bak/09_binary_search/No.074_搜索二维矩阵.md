## 74 搜索二维矩阵-中等

题目：

给你一个满足下述两条属性的 `m x n` 整数矩阵：

- 每行中的整数从左到右按非严格递增顺序排列。
- 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target` ，如果 `target` 在矩阵中，返回 `true` ；否则，返回 `false` 。



分析：

先对每一行的第一个元素进行 LastEqualOrSmaller 二分查找，找到相应的行，然后在进行 FirstEqual 查找。

```go
// date 2023/12/04
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    up, down := 0, m-1
    // find the last equal or smaller
    for up <= down {
        mid := (up + down) / 2
        if matrix[mid][0] > target {
            down = mid - 1
        } else {
            up = mid + 1
        }
    }
    // find 
    if down >= 0 && down < m {
        // find col
        n := len(matrix[down])
        left, right := 0, n-1
        for left <= right {
            mid := (left + right) / 2
            if matrix[down][mid] == target {
                return true
            } else if matrix[down][mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return false
}
```

