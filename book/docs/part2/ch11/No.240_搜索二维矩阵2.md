## 240 搜索二维矩阵-中等

题目：

编写一个高效的算法来搜索 `*m* x *n*` 矩阵 `matrix` 中的一个目标值 `target` 。该矩阵具有以下特性：

- 每行的元素从左到右升序排列。
- 每列的元素从上到下升序排列。



分析：

因为矩阵是每行升序，每列升序，所以可以以右上角的元素为基准。

如果右上角元素大于目标值，那么这一列不会有答案；

如果右上角元素小于目标值，那么这一行不会有答案；

```go
// date 2023/12/05
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    // Z 字查找
    x, y := 0, n-1
    for x < m && y >= 0 {
        if matrix[x][y] == target {
            return true
        } else if matrix[x][y] > target {  // 说明这一列不可能有
            y--
        } else { // 说明这一行不可能存在答案
            x++
        }
    }
    return false
}
```

