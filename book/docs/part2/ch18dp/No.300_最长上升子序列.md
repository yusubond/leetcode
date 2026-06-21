## 300 最长上升子序列

题目：

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。




分析：

题目中并没有求最长**连续**上升子序列，所以子序列只是保持相对位置的元素集合。

那么定义 dp[i] 表示包含元素nums[i]的最长上升序列的长度，那么两层遍历即可得到下面的递推公式：

```
// dp[i] 全量初始化为1, 表示 nums[i] 自身就是长度为 1 的上升子序列
当 0 <= j < i && nums[i] > nums[j] 时
dp[i] = max(dp[i], dp[j]+1)
```

算法 1：标准的递推，时间复杂段 O(n2)


```go
// date 2023/11/09
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	res := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1  // nums[1] 自身就是长度为 1 的上升序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}

	return res
}
```

算法2：贪心+二分查找

维护的 `lis` 始终保持**严格递增**——它的语义是「长度为 `k+1` 的递增子序列的**最小结尾值**」。

既然单调递增，找「第一个 `>= x` 的位置」就可以用**二分**，把内层循环从 `O(n)` 降到 `O(log n)`。Go 用标准库 `sort.Search`（需 `import "sort"`）：

```go
// date 2026/06/20
func lengthOfLIS(nums []int) int {
	lis := make([]int, 0, len(nums))

	for _, v := range nums {
    // lis 严格递增，二分找第一个 >= v 的位置
		pos := sort.Search(len(lis), func(i int) bool {
			return lis[i] >= v
		})

		if pos == len(lis) {
			lis = append(lis, v)  // v 比所有结尾都大，扩展更长序列
		} else {
			lis[pos] = v   // 替换，降低同长度序列末尾值
		}
	}

	return len(lis)
}
```

> ⚠️ `len(lis)` 就是答案，但 `lis` 这个数组**本身不一定是真实的 LIS**（替换会打乱元素顺序）。它只是「各长度最小结尾」的拼接；要还原真实子序列还需另存前驱指针。

复杂度：时间 `O(n log n)`，空间 `O(n)`。
