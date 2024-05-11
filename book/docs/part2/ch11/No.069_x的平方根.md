## 69 x的平方根-简单

题目：

给你一个非负整数 `x` ，计算并返回 `x` 的 **算术平方根** 。

由于返回类型是整数，结果只保留 **整数部分** ，小数部分将被 **舍去 。**

**注意：**不允许使用任何内置指数函数和算符，例如 `pow(x, 0.5)` 或者 `x ** 0.5` 。



分析：

二分查找，找到最后一个key，key的平方等于或小于 x 。

```go
// data 2023/12/06
func mySqrt(x int) int {
    // 查找区间[0,x]
    left, right := 0, x
    for left <= right {
        mid := left + (right-left)/2
        v := mid * mid
        if v > x {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return right
}
```

