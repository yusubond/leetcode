## 198 打家劫舍-中等

题目：

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。


分析：

动归，递推。每天晚上你都可以选择偷或者不偷，定义dp[i] 表示第i天的偷窃的总金额，那么有如下递推公式：

```
dp[i] = max(dp[i-1], dp[i-2]+nums[i])
```

```go
// date 2023/11/11
func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    n := len(nums)
    dp := make([]int, n)
    // dp[i] = 当晚偷窃的最大值
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2] + nums[i])
    }

    return dp[n-1]
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
// 简洁版
// 因为当晚可获得最大金额至于前两天的状态有关，所以可用两个变量维护。
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    // y1 前一天可偷的最大值
    // y2 前两天可偷的最大值
    y2 := nums[0]
    y1 := max(nums[0], nums[1])
    y0 := max(y1, y2)
    for i := 2; i < n; i++ {
        y0 = max(y1, y2+nums[i])
        y2, y1 = y1, y0
    }

    return y0
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```
