## 367 有效的完全平方数-简单

题目：

给你一个正整数 `num` 。如果 `num` 是一个完全平方数，则返回 `true` ，否则返回 `false` 。

**完全平方数** 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如 `sqrt` 。



分析：

如果存在一个数 x 的平方等于 num，那么这个数一定是`1<=x <= num`。

所以我们直接对[1, num]区间中的数进行二分查找。

```go
// date 2023/12/05
func isPerfectSquare(num int) bool {
    left, right := 0, num

    for left <= right {
        mid := (left + right) / 2
        v := mid * mid
        if v == num {
            return true
        } else if v < num {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

