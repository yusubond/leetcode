## 215 数组中第K大的元素-中等

题目要求：返回数组中第K大元素

题目链接：https://leetcode-cn.com/problems/sort-an-array/

算法分析：利用小顶堆保存K的元素，则堆顶就是想要的结果。



首先将前k个元素保存，然后初始化小顶堆；继续遍历，如果元素大于堆顶元素，则与堆顶元素交换，重新调整小顶堆。

```go
// date 2022/09/24
func findKthLargest(nums []int, k int) int {
	res := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		if i < k {
			res = append(res, nums[i])
			if len(res) == k {
				makeMinHeap(res)
			}
		} else if nums[i] > res[0] {
			res[0] = nums[i]
			minHeapify(res, 0)
		}
	}
	return res[0]
}

// 采用自底向上的建堆方法
func makeMinHeap(nums []int) {
	for i := len(nums) >> 1 - 1; i >= 0; i-- {
		minHeapify(nums, i)
	}
}

//
func minHeapify(nums []int, k int) {
	if k > len(nums) {
		return
	}	
	temp, n := nums[k], len(nums)
	l, r := k<<1+1, k<<1+2
	for l < n {
		r = l+1
		if r < n && nums[r] < nums[l] {
			l++
		}
		if nums[k] < nums[l] {
			break
		}
		nums[k] = nums[l]
		nums[l] = temp
		k = l
		l = k<<1+1
	}
}
```

