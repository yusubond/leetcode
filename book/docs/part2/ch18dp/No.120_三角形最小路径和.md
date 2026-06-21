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


变体：自顶向下一维 DP（不修改输入）

上面是「自底向上 + 原地修改 triangle」，额外空间已是 O(1)。若不想破坏输入，可用一维 `dp` 自顶向下滚动：`dp[j]` 表示到达当前行位置 `j` 的最小路径和。每行长度递增，需**从右往左**更新——更新 `dp[j]` 时，`dp[j]`（更新前）是上一行同列（上方），`dp[j-1]`（更新前）是上一行左列（左上），从右往左保证两者都还是上一行的旧值。

```go
// date 2026/06/20
func minimumTotal1D(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, n)
    dp[0] = triangle[0][0]
    for i := 1; i < n; i++ {
        dp[i] = dp[i-1] + triangle[i][i]                 // 最右：只能从左上来
        for j := i - 1; j >= 1; j-- {
            dp[j] = triangle[i][j] + min(dp[j-1], dp[j]) // min(左上 dp[j-1], 上方 dp[j])
        }
        dp[0] += triangle[i][0]                          // 最左：只能从上方来
    }
    res := dp[0]
    for j := 1; j < n; j++ {
        if dp[j] < res { res = dp[j] }
    }
    return res
}
```

> 原地自底向上版本空间更省（O(1) 额外）；这里的一维版优势在于**不破坏输入**，且写法与网格 DP 的滚动一致。复杂度：时间 O(n²)，空间 O(n)。
