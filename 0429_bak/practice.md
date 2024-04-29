## 题目练习

题目1：

把一张 100 元面值的纸币，换成1，2，5，10，20，50元面值的纸币，共有多少种换法？

如果限制换来的纸币数量不超过 N 张，有多少中换法？



分析：

求组合数，详见第 518 题。

```go
func change(coins []int, total int) int {
	//coins := []int{1, 2, 5, 10, 20, 50}
	//total := 100
	dp := make([]int, total+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= total; i++ {
			if i < coin {
				continue
			}
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[total]
}
```








题目2：

要求构造一个函数，实现以下功能：

输入一个列表 list，返回一个计数列表 count，count[i] 表示 list 中第 i 个元素右边有多少个数小于list[i]。

要求算法复杂度度尽可能低。
