## 746 使用最小花费爬楼梯-简单

题目：

给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

分析

还是动态规划，对于到达 step 台阶，有如下递推公式：

```
total_cost[step] = min(total_cost[step-1], total_cost[step-2]) + cost[step]
```

```go
// date 2023/11/09
func minCostClimbingStairs(cost []int) int {
    // total cost for step
    // total_cost[step] = min(total_cost[step-1], total_cost[step-2]) + cost[step]

    n := len(cost)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return cost[0]
    }
    step1, step2 := cost[1], cost[0]
    var tempCost int
    for i := 2; i < len(cost); i++ {
        tempCost = min(step1, step2) + cost[i]
        step2 = step1
        step1 = tempCost
    }

    return min(step1, step2)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```
