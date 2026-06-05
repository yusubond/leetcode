## 540 有序数组中的单一元素-中等

题目：

给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 `O(log n)` 时间复杂度和 `O(1)` 空间复杂度。



分析：

```go
// date 2023/12/05
func singleNonDuplicate(nums []int) int {
    n := len(nums)

    // find the first index x that nums[x] != nums[x^1]
    left, right := 0, n-1
    for left < right {
        mid := (right - left) / 2 + left
        // 只要和相邻的元素相等，那么 x 一定在右边
        // left = mid + 1
        if nums[mid] == nums[mid^1] {
            left = mid+1
        } else {
            right = mid
        }
    }
    return nums[left]
}
```

