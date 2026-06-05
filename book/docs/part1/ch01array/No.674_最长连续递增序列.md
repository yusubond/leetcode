## 674 最长连续递增序列-简单

题目：

给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。




分析：

直接遍历，统计即可。

```go
// date 2022/10/01
func findLengthOfLCIS(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    ans, curMax := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            curMax++
        } else {
            curMax = 1
        }
        if ans < curMax {
            ans = curMax
        }
    }
    return ans
}
```

