## 977 有序数组的平方-简单

题目：

给你一个按 **非递减顺序** 排序的整数数组 `nums`，返回 **每个数字的平方** 组成的新数组，要求也按 **非递减顺序** 排序。



分析：

对撞指针，因为可能存在负数，所以两头的平方比较大

```go
// date 2021/03/13
func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n, n)
    left, right := 0, n-1
    idx := n-1

    for left <= right {
        if left == right {
            res[idx] = sqar(nums[left])
            break
        }
        sl, sr := sqar(nums[left]), sqar(nums[right])
        if sr > sl {
            res[idx] = sr
            right--
        } else {
            res[idx] = sl
            left++
        }
        idx--
    }

    return res
}

func sqar(x int) int {
    return x*x
}
```

