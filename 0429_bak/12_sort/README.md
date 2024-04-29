## 排序算法

### 1 冒泡排序

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



### 2 插入排序

插入排序的思想是：

依次选择未排序的元素，将其插入到已排序的合适位置。

```go
// 插入排序
// 每次选择未排序的元素，将其插入到已排序的合适位置。
func insertSort(nums []int) []int {
	n := len(nums)

	for i := 1; i < n; i++ {
		v := nums[i]  // the unsort elem
		j := i - 1
    // find the last euqal or small than v
		for j >= 0 && nums[j] > v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v
	}

	return nums
}
```



### 3 选择排序

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



### 4 归并排序

归并排序（merge sort）是建立在归并操作上的一种有效的排序算法。该算法是分治思想的一个典型应用。

其思想是：

1. 将待排序数组不断地进行折半分开，直到不能再分（即每一小段是有序的）

2. 将 1 中每个有序的小段合并起来

因为折半拆分，其时间复杂度是O(nlogn)。代价是需要额外的空间保存结果。

```go
// merge sort
func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n / 2
	left := nums[0:mid]
	right := nums[mid:n]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(nums1, nums2 []int) []int {
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
```



### 快速排序



```go
func quickSort(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}
	p := division(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
	return nums
}

// 选取 left 作为基准base, 将数组分成小于和大于基准的两部分
// 并返回 base 的新下标
func division(nums []int, left, right int) int {
	base := nums[left]
	for left < right {
		// find the last small than base, and put is to left
		for left < right && nums[right] >= base {
			right--
		}
		nums[left] = nums[right]
		// find the first larger than base, and put it to right
		for left < right && nums[left] <= base {
			left++
		}
		nums[right] = nums[left]

		// save base to left
		nums[left] = base
	}
	return left
}
```

