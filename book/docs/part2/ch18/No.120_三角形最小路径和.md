## 120 三角形最小路径和-中等

题目：

给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。


分析：

不能正向求，因为正向只能看到局部最优解，而不是全局最优解。

所以，从底往上递归，并更新节点值。更新到顶部就是最小路径。


```go
// date 2023/11/13
func minimumTotal(triangle [][]int) int {
    m := len(triangle)
    if m == 0 {
        return 0
    }

    for i := m-2; i >= 0; i-- {
        // update this line
        th := len(triangle[i])
        for j := th-1; j >= 0; j-- {
            if j + 1 < len(triangle[i+1]) {
                triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
            }
        }
    }

    return triangle[0][0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```
