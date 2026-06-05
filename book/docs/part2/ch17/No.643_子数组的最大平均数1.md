## 643 子数组的最大平均数-简单

题目：

给你一个由 `n` 个元素组成的整数数组 `nums` 和一个整数 `k` 。

请你找出平均数最大且 **长度为 `k`** 的连续子数组，并输出该最大平均数。

任何误差小于 `10-5` 的答案都将被视为正确答案。



分析：

滑动窗口，直接计算。

注意元素只有一个，且 k 等于1的时候，平均数有可能小于零。

```go
// date 2023/11/21
func findMaxAverage(nums []int, k int) float64 {
    var ans float64
    ansSet := false
    left, right := 0, 0
    n := len(nums)
    ksum := 0
    for right < n {
        ksum += nums[right]
        right++
        if right - left >= k {
            res := float64(ksum) / float64(k)
            if !ansSet {
                ans = res
                ansSet = true
            } else if res > ans {
                ans = res
            }
            ksum -= nums[left]
            left++
        }
    }
    return ans
}
```

