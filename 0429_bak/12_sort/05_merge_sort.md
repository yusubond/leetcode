## Merge Sort \[归并排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[代码实现](#代码实现)
  * 3.[复杂度分析](#复杂度分析)
  * 4.[归并算法的应用](#归并算法的应用)
    

### 算法介绍

与快速排序一样，归并排序也是采用分治思想。将待排序列分成两个子序列(直到不可再分)，然后对两个子序列进行有序合并。merge()函数是归并排序的关键。

![归并排序过程演示](http://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/Merge-Sort.png)

### 代码实现

根据归并算法的原理，很容易想到利用递归的方式进行实现，如下所示：

```go
// merge sort
// 递归
// 不断地对半拆分，直到子序列长度小于2（即有序）
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

// 将两个有序的数组合并成一个有序数组
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

### 复杂度分析

归并排序属于递归算法，其时间复杂度可用如下公式表示：

$T(n) = 2T(n/2) + \theta(n)$

因为归并排序总是将序列分成两个子序列，直到不可分解，然后在 **线性时间** 内进行合并。所以三种情况(最好，最坏，一般)的时间复杂度均是 $\theta(nLogn)$ 。

从稳定性而言，归并排序比快速排序稳定。

辅助空间需要O(n)。

### 归并算法的应用

* 线性表的归并排序

  对于线性表而言，其与数组的主要区别在于并不需要申请内存空间，可以在O(1)的时间内和O(1)的空间内实现元素的插入。因此，对线性表进行归并排序，合并数据的过程不需要额外的辅助空间。
