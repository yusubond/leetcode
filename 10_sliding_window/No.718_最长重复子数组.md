## 718 最长重复子数组-中等

题目：

给两个整数数组 `nums1` 和 `nums2` ，返回 *两个数组中 **公共的** 、长度最长的子数组的长度* 。



> **示例 1：**
>
> ```
> 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
> 输出：3
> 解释：长度最长的公共子数组是 [3,2,1] 。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
> 输出：5
> ```



**解题思路**

分别以A对齐B，和 B 对齐 A 查找两个数组的最长重复子数组，返回结果最大值。

```go
// date 2024/03/12
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	m, n := len(nums1), len(nums2)

	for i := 0; i < m; i++ {
		x := min(n, m-i)
		r1 := maxLenght(nums1, nums2, i, 0, x)
		ans = max(ans, r1)
	}

	for j := 0; j < n; j++ {
		x := min(m, n-j)
		r1 := maxLenght(nums1, nums2, 0, j, x)
		ans = max(ans, r1)
	}

	return ans
}

func maxLenght(A, B []int, starta, startb int, size int) int {
	ans, k := 0, 0
	for i := 0; i < size; i++ {
		if A[starta+i] == B[startb+i] {
			k++
		} else {
			k = 0
		}
		ans = max(ans, k)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```



