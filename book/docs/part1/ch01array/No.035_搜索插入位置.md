## 35 搜索插入位置-简单

题目：

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。


分析：

```go
// date 2023/11/02
func searchInsert(nums []int, target int) int {
    i, s := 0, len(nums)
    for i < s {
        if nums[i] >= target {
            return i
        }
        i++
    }
    return i
}
```
