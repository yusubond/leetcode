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



## 与 No.377 组合总和 Ⅳ 的关系

**No.070 本质上是 No.377 的特例**：`nums = [1, 2]`，`target = n`。

爬楼梯求的是**排列数**，而非组合数——先走 1 步再走 2 步，与先走 2 步再走 1 步，是两条不同的路径。

```go
// No.377 的写法，nums = [1, 2]
func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i <= target; i++ {       // 外层：目标
        for _, num := range nums {        // 内层：可选步长
            if i >= num {
                dp[i] += dp[i-num]
            }
        }
    }
    return dp[target]
}
```

对比 No.070 泛化写法——**结构完全一致**：

```go
// date 2023/11/07
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

**循环顺序决定排列还是组合**：

| 循环方式 | 结果 | 应用到爬楼梯 |
|---------|------|-------------|
| 外层目标 i，内层步长 → **排列数** | No.377 / No.070 泛化写法 | 1→2 和 2→1 是不同的路径 |
| 外层步长，内层目标 i → **组合数** | No.518 零钱兑换 2 | 对爬楼梯无意义（同组步长只计一次） |

**推广到任意步长**：若每次可爬 `[1, 2, 3]` 步，递推公式为 `dp[i] = dp[i-1] + dp[i-2] + dp[i-3]`；若步长集合为 `steps`，则为 `dp[i] = Σ dp[i-step]`。这就是 No.377 的泛化写法。

> **一句话**：爬楼梯 = 用步长集合 `steps` 凑出目标 `n` 的排列数，等价于 No.377(nums=steps, target=n)。斐波那契写法（`dp[i]=dp[i-1]+dp[i-2]`）只是 `steps=[1,2]` 时的特例优化。

