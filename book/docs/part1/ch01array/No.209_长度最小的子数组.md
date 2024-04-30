---
title: "209 长度最小的子数组"
---

## 209 长度最小的子数组-中等

题目

给定一个含有 `n` 个正整数的数组和一个正整数 `target` **。**

找出该数组中满足其总和大于等于 `target` 的长度最小的 **连续子数组**

`[numsl, numsl+1, ..., numsr-1, numsr]` ，并返回其长度**。**如果不存在符合条件的子数组，返回 `0` 。



**思路分析**

虽然这里也是用双指针，但其实是滑动窗口的思想。

`sum`维护窗口内的元素和，`right`指针遍历数组，一旦`sum`满足条件，`left`指针缩小窗口。

```go
// date 2024/03/27
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := n + 1
	left, right := 0, 0
	sum := 0

	for right < n {
		sum += nums[right]
		right++
		for sum >= target {
			ans = min(ans, right-left)

			sum -= nums[left]
			left++
		}
	}

	if ans == n+1 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```

