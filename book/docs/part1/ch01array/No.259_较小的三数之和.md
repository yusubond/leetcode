## 259 较小的三数之和-中等

题目：

给定一个长度为 `n` 的整数数组和一个目标值 `target` ，寻找能够使条件 `nums[i] + nums[j] + nums[k] < target` 成立的三元组 `i, j, k` 个数（`0 <= i < j < k < n`）。



**解题思路**

排序和双指针。

```go
// date 2024/03/04
func threeSumSmaller(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i+2 < n && nums[i]+nums[i+1]+nums[i+2] >= target {
			break
		}
		left, right := i+1, n-1
		sum := target - nums[i]
		for left < right {
			if nums[left]+nums[right] < sum {
        // select nums[right]
        // i, left, right
				ans += right - left
				left++
			} else {
				right--
			}
		}
	}

	return ans
}
```

