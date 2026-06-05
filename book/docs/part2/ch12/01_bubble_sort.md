## 冒泡排序

冒泡排序(Bubble Sort)的思想是：

每次比较两个相邻的两个元素，按照大小顺序进行交换，这样在一次的遍历中，可以将最大或最小的值交换值序列的最后一个。

时间复杂度：O(n x n)

空间复杂度：O(1)

```go
// 冒泡排序
// 比较相邻的两个元素，把较大的交换至后面
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
```

