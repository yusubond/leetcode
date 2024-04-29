## 219 存在重复元素2-简单

题目：

给定一个数组nums和整数k，判断是否存在两个不同的下标，使得`nums[i] == nums[j]`，同时`abs(i-j) <= k`，如果存在返回true，否则返回false。



分析：

推荐该算法，滑动窗口。

这里滑动窗口的技巧在与题目中所要求的`abs(i-j) <= k`，所以只判断连续的k个元素即可。

```go
// date 2022/09/27
func containsNearbyDuplicate(nums []int, k int) bool {
    size := len(nums)

    start, end := 0, 0

    for start < size {
        end = start+1
        for end - start <= k && end < size {
            if nums[start] == nums[end] {
                return true
            }
            end++
        }
        start++
    }
    
    return false
}
```



算法2：利用map，保存遍历过的元素

```go
// date 2022/09/27
func containsNearbyDuplicate(nums []int, k int) bool {
  m := make(map[int]int)
  for index, v := range nums {
    if j, ok := m[v]; ok && index - j <= k {
      return true
    }
    m[v] = index
  }
  return false
}
```

