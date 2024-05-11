## 选择排序

选择排序的思想是：

每次选择未排序中的最值，将其放到已排序的尾部。

优势：可就地替换，不占用额外空间

时间复杂度：O(n x n)

空间复杂度：O(1)

```go
// 选择排序
func selectSort(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		// swap
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}
```

