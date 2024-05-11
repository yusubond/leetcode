## 计数排序

计数排序的思想是：

将待排序数组的值转换为键存储在额外的数组空间。因为键是有序的，通过计数键出现的次数，将待排序的数组在复原出来。

这是一种线性时间复杂度的排序，而且要求输入的数组必须是有确定范围的整数。



时间复杂度为：O(n)

空间复杂度为：O(n)

如果输入数组的范围未知，需要先求一次最大值，找到数组的范围，这种情况下，其时间复杂度为：O(2N)



```go
// counting sort
func countSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	// find the maxValue
	maxV := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	cnt := make([]int, maxV+1, maxV+1)
	// counting the elem
	for _, v := range nums {
		cnt[v]++
	}

	idx := 0
	for i, ct := range cnt {
		for ct > 0 {
			nums[idx] = i
			idx++
			ct--
		}
	}
	return nums
}
```

