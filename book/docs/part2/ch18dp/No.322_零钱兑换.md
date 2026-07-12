## 322 零钱兑换-中等

题目：

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 **最少的硬币个数** 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



**关键观察**

1. **最优子结构**：凑出金额 `target` 的最优解，可以由"先凑出 `target - coin`，再加一枚 `coin`"得到。即 `opt[target] = min(opt[target], 1 + opt[target - coin])`。
2. **无限硬币**：每枚硬币可以重复使用，因此内层正序遍历金额（完全背包）即可，与 No.518 的循环顺序一致。
3. **无法凑出**：用 `amount + 1` 作为"不可达"的哨兵值（因为最坏情况也只需要 amount 枚 1 元硬币），最终 `dp[amount]` 仍为此值则返回 -1。



**解法一：DP 自底向上（完全背包）**

```go
// date 2023/11/08
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1 // 哨兵：不可达
    }
    // dp[0] = 0 已由 make 初始化

    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if i < coin {
                continue
            }
            dp[i] = min(dp[i], 1+dp[i-coin])
        }
    }

    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
```

交换内外层循环也正确（与 No.518 不同，这里求的是**最小值**而非组合数，顺序不影响最优解）：

```go
// date 2026/07/12
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1
    }
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            dp[i] = min(dp[i], 1+dp[i-coin])
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
```



**解法二：DFS + 记忆化（自顶向下）**

从 amount 出发，每次尝试减去一枚硬币，递归求解子问题。

```go
// date 2026/07/12
func coinChange(coins []int, amount int) int {
    memo := make([]int, amount+1)
    for i := range memo {
        memo[i] = -2 // -2 表示未计算；-1 表示不可达
    }
    memo[0] = 0

    var dfs func(rem int) int
    dfs = func(rem int) int {
        if rem < 0 {
            return -1
        }
        if memo[rem] != -2 {
            return memo[rem]
        }
        res := amount + 1
        for _, coin := range coins {
            sub := dfs(rem - coin)
            if sub >= 0 {
                res = min(res, 1+sub)
            }
        }
        if res > amount {
            memo[rem] = -1
        } else {
            memo[rem] = res
        }
        return memo[rem]
    }

    return dfs(amount)
}
```



**解法三：BFS（图的最短路径）**

将问题建模为图：节点是金额 0..amount，从金额 `cur` 到 `cur - coin` 有一条边（权为 1）。求从 amount 到 0 的最短路径。

BFS 按层遍历，第一次访问到某个金额时就是最少硬币数。**比 DP 的优势**：早停，不需要遍历所有状态。

```go
// date 2026/07/12
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    visited := make([]bool, amount+1)
    queue := []int{amount}
    visited[amount] = true
    steps := 0

    for len(queue) > 0 {
        steps++
        size := len(queue)
        for k := 0; k < size; k++ {
            cur := queue[0]
            queue = queue[1:]
            for _, coin := range coins {
                next := cur - coin
                if next == 0 {
                    return steps
                }
                if next > 0 && !visited[next] {
                    visited[next] = true
                    queue = append(queue, next)
                }
            }
        }
    }
    return -1
}
```

BFS 的图视角也可以反过来——从 0 出发，每次加一枚硬币，找最短到达 amount 的路径。效果等价，方向偏好取决于直觉。



**复杂度对比**

| 解法 | 时间 | 空间 | 特点 |
|------|------|------|------|
| DP 自底向上 | O(n × amount) | O(amount) | 最经典，实现简单 |
| DP 自顶向下 | O(n × amount) | O(amount) | 可能跳过一些不可达状态 |
| BFS | O(n × amount) 最坏 | O(amount) | 可早停，大量不可达时优势明显 |

> `n` 为硬币种类数。



**与相邻题的关系**

| 题目 | 目标 | 核心区别 |
|------|------|---------|
| **No.322（本题）** | 最少硬币数 | 求最小值，循环顺序无影响 |
| No.518 零钱兑换 2 | 组合数 | 求方案数，**外 coin 内 amount 是组合，反之是排列** |
| No.279 完全平方数 | 最少平方数 | 本题特例：coins = {1,4,9,16,…} |
| No.377 组合总和 Ⅳ | 排列数 | 求排列数，外层 amount 内层 coin |

> `min` 用 Go 1.21+ 内置函数即可；旧版本需自行定义。
