---
title: "第 7 章 二分查找"
bookCollapseSection: true
---


# 第 7 章 二分查找

- [7.1 二分查找](./binarysearch.md)
- [7.2 二分查找的变种](./readme.md)



二分查找：对有序数据的查找，最简单的就是线性搜索，时间复杂度 O(N)；二分查找，利用数据的有序性，可将时间复杂度降低至 O(LogN)。



代码结构如下：
```go
func BinarySearch(nums []int, target int) bool {
  left, right := 0, len(nums)-1
  for left <= right {
    mid := left + (right-left)/2
    v := nums[mid]
    if v == target {
      return true
    } else if v < target {
      left = mid+1
    } else {
      right = mid-1
    }
  }
  return false
}
```

