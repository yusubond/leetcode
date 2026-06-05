## 220 存在重复元素3-困难

题目：

给定一个数组nums和两个整数k,t。判断是否存在两个不同的下标，使得`nums[i]-nums[j] <= t`，同时`abs(i-j) <= k`，如果存在返回true，否则返回false。



分析：

滑动窗口，使用双指针控制窗口大小。该算法可行，但是某些测试用例会超时，所以需要寻找其他解法。

```go
// 该算法会超时
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    n := len(nums)
    left, right := 0, 0
    for left < n {
        right = left+1
        for right - left <= indexDiff && right < n {
            if abs(nums[left], nums[right]) <= valueDiff {
                return true
            }
            right++
        }
        left++
    }
    return false
}

func abs(x, y int) int {
    if x < y {
        return y-x
    }
    return x-y
}
```



