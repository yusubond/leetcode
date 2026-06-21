## 740 删除并获得点数

题目：

给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。


分析：

进阶版的打家劫舍。

先统计数组中每个元素出现的总点数，然后再做打家劫舍。

```go
// date 2023/11/11
func deleteAndEarn(nums []int) int {
    maxVal := 0
    for _, v := range nums {
        maxVal = max(maxVal, v)
    }
    sum := make([]int, maxVal+1)
    for _, v := range nums {
        sum[v] += v
    }

    // 对 sum 打家劫舍
    if len(sum) == 1 {
        return sum[0]
    }
    if len(sum) == 2 {
        return max(sum[0], sum[1])
    }
    f1, f2 := sum[0], sum[1]
    f := 0
    for i := 2; i < len(sum); i++ {
        f = max(f2, f1 + sum[i])
        f1, f2 = f2, f
    }
    return f
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```
