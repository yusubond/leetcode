## 713 乘积小于K的子数组

题目：

给你一个整数数组 `nums` 和一个整数 `k` ，请你返回子数组内所有元素的乘积严格小于 `k` 的连续子数组的数目。



分析：

两层循环直接计算。

注意，单个元素小于 K 也要计算在内。

```go
// date 2023/11/21
func numSubarrayProductLessThanK(nums []int, k int) int {
    var ans int

    sum := 1
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] < k {
            ans++
        }
        sum = nums[i]
        for j := i+1; j < n; j++ {
            sum *= nums[j]
            if sum >= k {
                break
            } else {
                ans++
            }
        }
    }

    return ans
}
```

