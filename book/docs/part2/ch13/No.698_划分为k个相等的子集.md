## 698 划分为k个相等的子集-中等

题目：

给定一个整数数组 `nums` 和一个正整数 `k`，找出是否有可能把这个数组分成 `k` 个非空子集，其总和都相等。



> **示例 1：**
>
> ```
> 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
> 输出： True
> 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [1,2,3,4], k = 3
> 输出: false
> ```



分析：

这道题的解题思路比较复杂。

既有回溯，又要考虑剪枝，否则会超时。

```go
// date 2023/12/27
func canPartitionKSubsets(nums []int, k int) bool {
    sum, theMax := calcSumAndMax(nums)
    if sum % k  != 0 {
        return false
    }
    target := sum / k
    if theMax > target {
        return false
    }
    // 剩下的问题就是 求子数组和为 target
    // 也就是把所有的元素放入 K 个桶
    backet := make([]int, k)
    n := len(nums)
    var backtrack func(idx int, buck []int) bool
    backtrack = func(idx int, buck []int) bool {
        if idx == n {
            return true
        }
        for i := 0; i < k; i++ {
            // 剪枝1
            // 如果当前桶和上一个桶一样，当前值已经在上一个桶计算过了，可以跳过
            if i > 0 && buck[i] == buck[i-1] {
                continue
            }
            // 剪枝2
            // 如果桶超过容量，直接跳过
            if buck[i] + nums[idx] > target {
                continue
            }
            buck[i] += nums[idx]
            if backtrack(idx+1, buck) { // 如果有解，直接返回
                return true
            }
            buck[i] -= nums[idx]
        }
        return false
    }

    return backtrack(0, backet)
}

func calcSumAndMax(nums []int) (sum int, max int) {
    sum = 0
    max = nums[0]
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    return
}
```

