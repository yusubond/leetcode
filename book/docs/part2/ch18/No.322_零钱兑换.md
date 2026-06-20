## 322 零钱兑换-中等

题目：

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



**解题思路**

该问题具有最优子结构，设目标值为 target，目标值对应的最优解为 opt[target]，那么在所有的硬币里面可以得到下面的递推公式：

```
opt[target] = min(opt[target], 1+opt[target-coin])
```


```go
// date 2023/11/08
func coinChange(coins []int, amount int) int {
	// dp[i] 凑出 i 金额的最少硬币数量
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			// pick up this coin
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
```
