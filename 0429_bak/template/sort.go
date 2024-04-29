package template

// BubbleSort define
// 冒泡排序的思路
// 通过交换不断地把最大值移动到数组的后面
func BubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// SelectSort define
// 选择排序的思路是
// 从未排序的列表中选择最值，交换到数组的前面
func SelectSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		// swap the smallest to idx
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

// InsertSort define
// 插入排序的思路是
// 依次选择未排序的元素，插入到已排序（前面）的序列中
func InsertSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		v := nums[i] // the unsort elem

		// find the last equal or smaller than v
		// from end to start
		j := i - 1
		for j >= 0 && nums[j] > v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v
	}
	return nums
}

// QuickSort define
func QuickSort(nums []int) []int {
	left, right := 0, len(nums)-1
	quickSortArr(nums, left, right)
	return nums
}

// 选择 left 作为基准base，把数组分成小于等于和大于base的两部分
// 并返回两部分的分界线下标
func quickDivision(nums []int, left, right int) int {
	base := nums[left]
	for left < right {
		for left < right && nums[right] > base {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= base {
			left++
		}
		nums[right] = nums[left]
		nums[left] = base
	}
	return left
}

func quickSortArr(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := quickDivision(nums, left, right)
	quickSortArr(nums, left, p-1)
	quickSortArr(nums, p+1, right)
}

// MergeSort define
func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	left := nums[0:mid]
	right := nums[mid:n]
	return mergeArr(MergeSort(left), MergeSort(right))
}

// merge the sorted arrays
func mergeArr(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	ans := make([]int, 0, n1+n2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if nums1[i] <= nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	if i < n1 {
		ans = append(ans, nums1[i:]...)
	}
	if j < n2 {
		ans = append(ans, nums2[j:]...)
	}
	return ans
}
