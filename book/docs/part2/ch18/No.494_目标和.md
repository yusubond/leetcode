## 494 目标和-中等

题目：

给你一个非负整数数组 `nums` 和一个整数 `target` 。

向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：

- 例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。



> **示例 1：**
>
> ```
> 输入：nums = [1,1,1,1,1], target = 3
> 输出：5
> 解释：一共有 5 种方法让最终目标和为 3 。
> -1 + 1 + 1 + 1 + 1 = 3
> +1 - 1 + 1 + 1 + 1 = 3
> +1 + 1 - 1 + 1 + 1 = 3
> +1 + 1 + 1 - 1 + 1 = 3
> +1 + 1 + 1 + 1 - 1 = 3
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1], target = 1
> 输出：1
> ```



分析：

**关键转化：化归为「子集和 / 0-1 背包」**

把 nums 分成两组：加号组 P（前面填 `+`）、减号组 N（前面填 `-`），要求 `sum(P) - sum(N) = target`。设所有数之和为 `totalSum`，则有：

- `sum(P) - sum(N) = target`
- `sum(P) + sum(N) = totalSum`

两式相加得 `2·sum(P) = target + totalSum`，即 `sum(P) = (target + totalSum) / 2`。

于是问题转化为：**从 nums 中选若干个数，使其和恰好等于 `(target + totalSum) / 2`，求选法数** —— 即「装满容量为 bag 的背包有多少种装法」，经典的 0-1 背包计数。

### 方法一：动态规划（0-1 背包，最优）

```go
// date 2023/12/27
func findTargetSumWays(nums []int, target int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    diff := target + totalSum
    if diff < 0 || diff%2 != 0 { // sum(P) 必须是非负整数
        return 0
    }
    bag := diff / 2

    // dp[j] 表示和为 j 的子集选法数
    dp := make([]int, bag+1)
    dp[0] = 1 // 空集，和为 0，有 1 种
    for _, num := range nums {
        for j := bag; j >= num; j-- { // 逆序：保证每个数只选一次
            dp[j] += dp[j-num]
        }
    }
    return dp[bag]
}
```

- 时间复杂度 `O(n·bag)`，空间复杂度 `O(bag)`。
- 内层循环必须**逆序**：正序会让同一个数在同一轮被重复计入（那是完全背包的写法）；逆序保证 `dp[j-num]` 读到的是「上一层」的值。
- 边界判断：`diff < 0` 或 `diff` 为奇数都说明无解；若推出的 `bag > totalSum`，DP 自然会得到 0。

### 方法二：记忆化 DFS

在暴力 DFS 的基础上加备忘录，状态 `(idx, target)` 去重。这里把 DFS 改成「返回结果」的形式，天然适合记忆化：

```go
// date 2023/12/27
func findTargetSumWays(nums []int, target int) int {
    memo := make(map[[2]int]int)
    var dfs func(idx, target int) int
    dfs = func(idx, target int) int {
        if idx == len(nums) {
            if target == 0 {
                return 1
            }
            return 0
        }
        key := [2]int{idx, target}
        if v, ok := memo[key]; ok {
            return v
        }
        res := dfs(idx+1, target-nums[idx]) + dfs(idx+1, target+nums[idx])
        memo[key] = res
        return res
    }
    return dfs(0, target)
}
```

- 状态数为 `O(n·sum)`，每个状态计算 `O(1)`，相比暴力 DFS 的 `O(2ⁿ)` 大幅优化。
- `map` 的 key 用 `[2]int{idx, target}`，`target` 取负值也能正常作为 key。

### 方法三：暴力 DFS（回溯）

每个位置选 `+` 或 `-`，复杂度 `O(2ⁿ)`，n 较大时会超时。其递推关系为：

```sh
// dfs(idx, target) 表示从下标 idx 开始，剩余目标和为 target 的解的个数
dfs(idx, target) = dfs(idx+1, target-nums[idx]) + dfs(idx+1, target+nums[idx])
```

```go
// date 2023/12/27
func findTargetSumWays(nums []int, target int) int {
    res := 0
    n := len(nums)

    var backtrack func(idx, target int)
    backtrack = func(idx, target int) {
        if idx == n {
            if target == 0 {
                res += 1
            }
            return
        }
        backtrack(idx+1, target-nums[idx])  // 选 + nums[idx]
        backtrack(idx+1, target+nums[idx])  // 选 - nums[idx]
    }

    backtrack(0, target)

    return res
}
```
