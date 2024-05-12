## 300 最长上升子序列

题目：

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


分析：

题目中并没有求最长**连续**上升子序列，所以子序列只是保持相对位置的元素集合。

那么定义 dp[i] 表示包含元素nums[i]的最长上升序列的长度，那么两层遍历即可得到下面的递推公式：

```
当 0 <= j < i && nums[i] > nums[j] 时
dp[i] = max(dp[j]) + 1
```

初始状态为 dp[0] = 1，即一个元素组成的上升序列。


```go
// date 2023/11/09
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res := 1
    n := len(nums)

    dp := make([]int, n)
    dp[0] = 1

    for i := 1; i < n; i++ {
        dpj_max := 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dpj_max = max(dpj_max, dp[j])
            }
        }
        dp[i] = dpj_max + 1

        res = max(res, dp[i])
    }

    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

算法2：

```go
// date 2020/03/16
func lengthOfLIS(nums []int) int {
  if len(nums) == 0 { return 0}
  n := len(nums)
  lis := make([]int, 0, n)
  lis = append(lis, nums[0])
  for i := 1; i < n; i++ {
    // 在lis数组找到第一个比nums[i]大的元素，并替换它，否则进行追加
    m := len(lis)
    // nums[i] > lis[m-1]
    // nums[i] 直接参与最长上升序列
    if nums[i] > lis[m-1] {
      lis = append(lis, nums[i])
      continue
    }
    // 替换的必要性
    // 1. 保证 nums[i] 有机会参与最长上升序列,为后面更大的值剔除障碍
    // 2. 因为不要求连续性，lis 中的大值没有意义，替换相对于变相删除
    for j := 0; j < m; j++ {
      if nums[i] <= lis[j] {
        lis[j] = nums[i]
        break
      }
    }
  }
  return len(lis)
}
```
