## 673 最长递增子序列的个数-中等

题目：

给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。


分析：

dp 表示递增子序列的长度。

cnt 表示以 nums[i] 结尾的长度为dp[i]的个数。

```go
// date 2023/11/13
func findNumberOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    maxLen := 0

    ans := 0

    for i := 0; i < n; i++ {
        dp[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if dp[j] + 1 > dp[i] {
                    dp[i] = dp[j] + 1
                    cnt[i] = cnt[j]   // dp[i] 在更新，所以重置cnt[i]个数
                } else if dp[j] + 1 == dp[i] {
                    cnt[i] += cnt[j]  // 说明存在多个，累加
                }

            }
        }
		// 最终结果集，更新或累加
        if dp[i] > maxLen {
            maxLen = dp[i]
            ans = cnt[i]
        } else if dp[i] == maxLen {
            ans += cnt[i]
        }
    }
    return ans
}
```
