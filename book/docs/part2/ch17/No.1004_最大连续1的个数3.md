## 1004 最大连续1的个数3-中等

题目：

给定一个二进制数组 `nums` 和一个整数 `k`，如果可以翻转最多 `k` 个 `0` ，则返回 *数组中连续 `1` 的最大个数* 。



分析：

对于数组区间 [left, right] 而言，只要区间内不超过 k 个 0，那么该区间就可以满足条件，其长度为 right-left+1。

由此，这个问题可以转化成：

> 对于任意的右端点 right，希望找到最小的左端点 left，使得 [left, right] 包含不超过 k 个 0

想要快速判断一个区间内 0 的个数，我们可以把数组中的 1 变成 0，0 变成 1，对数组求前缀和。

只要 left 的前缀和 lsum，与 右端点 right 的前缀和 rsum，满足 lsum < rsum - k，那么符合条件的区间就形成了。



```go
// date 2023/11/22
func longestOnes(nums []int, k int) int {
    var ans int
    left, lsum, rsum := 0, 0, 0

    // 问题变成求：索引 right 前 0 的个数不能超 k 个 的最长元素个数
    for right, v := range nums {
        // 求 0 的个数，即前缀和
        rsum += 1 - v

        // 当 lsum < rsum - k，说明已经超过了 k 个
        // left++
        for lsum < rsum - k {
            lsum += 1 - nums[left]
            left++
        }
        ans = max(ans, right-left+1)
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

