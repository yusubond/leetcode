## 912 排序数组-中等

题目：

给你一个整数数组 `nums`，请你将该数组升序排列。



**解题思路**

```go
// date 2024/03/03
// 解法1 快速排序
// 需要随机选取基准，否则会超时
func sortArray(nums []int) []int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := division(nums, left, right)
	quickSort(nums, left, mid-1)
	quickSort(nums, mid+1, right)
}

func division(nums []int, left, right int) int {
	rdx := randRange(left, right)
	nums[left], nums[rdx] = nums[rdx], nums[left]
	base := nums[left]
	for left < right {
		for left < right && nums[right] >= base {
			right--
		}
		// now the nums[right] < base
		nums[left] = nums[right]
		for left < right && nums[left] <= base {
			left++
		}
		// now the nums[left] > base
		nums[right] = nums[left]
		nums[left] = base
	}
	return left
}

func randRange(l, r int) int {
	return rand.Intn(r-l) + l
}
```



`division`函数还有另外一种写法。

```go
// date 2024/03/04
// 快速排序
func sortArray(nums []int) []int {
	n := len(nums)
	quickSort2(nums, 0, n-1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := division2(nums, left, right)
	quickSort2(nums, left, mid-1)
	quickSort2(nums, mid+1, right)
}

func division2(nums []int, left, right int) int {
	rdx := randRange(left, right)
	swap(nums, rdx, right)
	base := nums[right]
	idx := left
	for left < right {
		if nums[left] < base {
			swap(nums, idx, left)
			idx++
		}
		left++
	}
	swap(nums, idx, right)
	return idx
}

func swap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
```



归并排序

```go
// date 2024/03/04
// 归并排序
func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	left := nums[0:mid]
	right := nums[mid:n]
	return mergeNums(mergeSort(left), mergeSort(right))
}

func mergeNums(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	nums := make([]int, n1+n2)
	idx, i, j := 0, 0, 0
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			nums[idx] = nums1[i]
			i++
		} else {
			nums[idx] = nums2[j]
			j++
		}
		idx++
	}
	for i < n1 {
		nums[idx] = nums1[i]
		idx++
		i++
	}
	for j < n2 {
		nums[idx] = nums2[j]
		idx++
		j++
	}
	return nums
}
```

