## 747 至少是其他数字两倍的最大数-简单

给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。

请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。



分析：

题目要求最大元素至少是数组中每个其他数字的两倍，因此只需要检查最大数和次大数即可，只要最大数是次大数的两倍，则返回结果，否则返回-1。

```go
// date 2022/10/07
func dominantIndex(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)-1
    }
    max1, max2 := 0, 0
    res := -1
    for i, v := range nums {
        if v > max1 {
            max2 = max1
            max1 = v
            res = i
        } else if v > max2 {
            max2 = v
        }
    }
    if max1 >= 2 * max2 {
        return res
    }
    return -1
}
```

