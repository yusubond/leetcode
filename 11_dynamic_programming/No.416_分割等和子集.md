## 416 分割等和子集

题目：

给你一个 **只包含正整数** 的 **非空** 数组 `nums` 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



>**示例 1：**
>
>```
>输入：nums = [1,5,11,5]
>输出：true
>解释：数组可以分割成 [1, 5, 5] 和 [11] 。
>```
>
>**示例 2：**
>
>```
>输入：nums = [1,2,3,5]
>输出：false
>解释：数组不能分割成两个元素和相等的子集。
>```





**解题思路**

```sh
// true 表示能够从数组中选取若干个元素，使其和为 target
dp[target] = true or false
```

```go
// date 2024/02/29
func canPartition(nums []int) bool {
    total, maxV := getSumAndMax(nums)
    if total%2 == 1 {
        return false
    }
    target := total >> 1
    if maxV > target {
        return false
    }
    if maxV == target {
        return true
    }
    n := len(nums)
    // dp[i][j]
    // 0 <= i <= n
    dp := make([][]bool, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]bool, target+1)
    }
    dp[0][0] = true

    for i := 1; i < n+1; i++ {
        // cur elem v
        v := nums[i-1]
        for j := 1; j < target+1; j++ {
            if j < v {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
            }
        }
    }

    return dp[n][target]
}

func getSumAndMax(nums []int) (int, int) {
    if len(nums) == 1 {
        return nums[0], nums[0]
    }
    res := 0
    max := nums[0]
    for _, v := range nums {
        res += v
        if v > max {
            max = v
        }
    }
    return res, max
}
```

