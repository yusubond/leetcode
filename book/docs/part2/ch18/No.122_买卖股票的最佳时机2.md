## 122 买卖股票的最佳时机2

题目：

给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。


分析：

因为最多只能持有一只股票，但是可以在同一天买卖。所以总利润就是前后两个价格的差求和。


```go
// date 2023/11/08
func maxProfit(prices []int) int {
    var res int

    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            res += prices[i] - prices[i-1]
        }
    }

    return res
}
```
