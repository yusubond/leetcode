## 70 爬楼梯-简单

题目：

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



**解题思路**

直接递推。c[i] = c[i-1] + c[i-2]

```go
// date 2023/11/07
func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    p1, p2 := 1, 2
    var res int
    for i := 3; i <= n; i++ {
        res = p1 + p2
        p1, p2 = p2, res
    }

    return res
}
```



## 动归的泛化写法

```go
// 递推公式 dp[i] 表示第 i 个台阶有多少种走法
// dp[i] = dp[i-1] + dp[i-2]，即两种情况
// 第一种：到达 i-1 的总解数，再走 1 步就到达 i
// 第二种：到达 i-2 的总解数，再走 2 步就到达 i
func climbStairs(n int) int {
    dp := make([]int, n+1)
    steps := []int{1, 2}
    dp[0] = 1

    for i := 1; i <= n; i++ {
        for _, step := range steps {
            if i < step {
                continue
            }
            dp[i] += dp[i-step]
        }
    }

    return dp[n]
}
```

