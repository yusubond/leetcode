## 121 买卖股票的最佳时机

题目：

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


分析：

因为只能交易一次，所以遍历的时候，记录前面出现的最小值，并用当前值计算出利润，并更新结果。

```go
// date 2023/11/08
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	profit := 0
	minPrice := prices[0]

	for _, curPrice := range prices {
		if curPrice > minPrice && profit < curPrice-minPrice {
			profit = curPrice - minPrice
		}
		// update the min price
		if curPrice < minPrice {
			minPrice = curPrice
		}
	}

	return profit
}
```
