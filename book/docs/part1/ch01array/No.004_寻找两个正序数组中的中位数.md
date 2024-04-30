## 004 寻找两个正序数组中的中位数-困难

题目：

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。


分析：


```go
// date 2023/11/01
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s1, s2 := len(nums1), len(nums2)
	total := s1+s2
	if total == 0 {
		return 0
	}
	targetIdx := total/2
	onlyOne := total % 2 == 1
	if s1 == 0 {
		if onlyOne {
			return float64(nums2[targetIdx])
		} else {
			return (float64(nums2[targetIdx-1]) + float64(nums2[targetIdx]))/2
		}
	}
	if s2 == 0 {
		if onlyOne {
			return float64(nums1[targetIdx])
		} else {
			return (float64(nums1[targetIdx-1]) + float64(nums1[targetIdx]))/2
		}
	}

	pre, cur := 0, 0
	i, j, nIdx := 0, 0, 0
	for i < s1 && j < s2 {
		if nums1[i] < nums2[j] {
			pre, cur = cur, nums1[i]
			i++
		} else {
			pre, cur = cur, nums2[j]
			j++
		}
		if nIdx == targetIdx {
			if onlyOne {
				return float64(cur)
			}
			return (float64(pre)+float64(cur))/2
		}
		nIdx++
	}
	for i < s1 {
		pre, cur = cur, nums1[i]
		i++
		if nIdx == targetIdx {
			if onlyOne {
				return float64(cur)
			}
			return (float64(pre)+float64(cur))/2
		}
		nIdx++
	}
	for j < s2 {
		pre, cur = cur, nums2[j]
		j++
		if nIdx == targetIdx {
			if onlyOne {
				return float64(cur)
			}
			return (float64(pre)+float64(cur))/2
		}
		nIdx++
	}

	if onlyOne {
		return float64(cur)
	}
	return (float64(pre)+float64(cur))/2
}
```
