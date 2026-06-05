## 162 寻找峰值-中等

题目：

峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 `nums`，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 **任何一个峰值** 所在位置即可。

你可以假设 `nums[-1] = nums[n] = -∞` 。

你必须实现时间复杂度为 `O(log n)` 的算法来解决此问题。



**解题思路**

```go
// date 2024/03/04
func findPeakElement(nums []int) int {
	n := len(nums)
	var getNum func(i int) int
	getNum = func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if getNum(mid-1) < getNum(mid) && getNum(mid) > getNum(mid+1) {
			return mid
		}
		if getNum(mid) < getNum(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
```

