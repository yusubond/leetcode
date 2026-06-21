## 275 H指数2-中等

题目：

给你一个整数数组 `citations` ，其中 `citations[i]` 表示研究者的第 `i` 篇论文被引用的次数，`citations` 已经按照 **升序排列** 。计算并返回该研究者的 h 指数。

[h 指数的定义](https://baike.baidu.com/item/h-index/3991452?fr=aladdin)：h 代表“高引用次数”（high citations），一名科研人员的 `h` 指数是指他（她）的 （`n` 篇论文中）**至少** 有 `h` 篇论文分别被引用了**至少** `h` 次。

请你设计并实现对数时间复杂度的算法解决此问题。



分析：

假设数组长度为 n，其实就是求 nums[i] >= n - i 的最右边界。

变相地等价于求数组中第一个小于 n - i 的元素。

```go
// date 2023/12/05
func hIndex(citations []int) int {
    n := len(citations)
    left, right := 0, n-1
    // find the first small key(key = n - mid)
    // the result is n - left
    for left <= right {
        mid := (left + right) / 2
        if citations[mid] >= n - mid {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return n - left
}
```

