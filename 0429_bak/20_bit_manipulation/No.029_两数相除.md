## 29 两数相除-中等

题目：

给你两个整数，被除数 `dividend` 和除数 `divisor`。将两数相除，要求 **不使用** 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（`truncate`）其小数部分。例如，`8.345` 将被截断为 `8` ，`-2.7335` 将被截断至 `-2` 。

返回被除数 `dividend` 除以除数 `divisor` 得到的 **商** 。

**注意：**假设我们的环境只能存储 **32 位** 有符号整数，其数值范围是 `[−231, 231 − 1]` 。本题中，如果商 **严格大于** `231 − 1` ，则返回 `231 − 1` ；如果商 **严格小于** `-231` ，则返回 `-231` 。



**解题思路**

这道题可用二分查找。把商作为搜索的目标，商的取值范围是[0, 被除数]，所以从区间 0 到 被除数 搜索。

```go
// 当 （商+1）* 除数 > 被除数 并且 商 * 除数 <= 被除数
// 或者 商+1）* 除数 >= 被除数 并且 商 * 除数 < 被除数
// 就算找到商了
```

注意，二分查找容易写错的点：

1. `low <= high`，这是循环的条件
2. `mid = low + (right-low)/2`，防止溢出
3. `low = mid-1, high = mid+1`，边界的更新要减1



```go
// date 2024/01/11
func divide(dividend int, divisor int) int {
    res, sign := 1, -1
    if dividend == 0 {
        return 0
    }
    if divisor == 1 {
        return dividend
    }

    if dividend < 0 &&  divisor < 0 || dividend > 0 && divisor > 0 {
        sign = 1
    }

    res = binarySearch(0, abs(dividend), abs(dividend), abs(divisor))
    if res > math.MaxInt32 {
        res = math.MaxInt32
    }
    if res < math.MinInt32 {
        res = math.MinInt32
    }
    return res * sign
}

// x/y
func binarySearch(left, right int, x, y int) int {
    mid := left + (right-left)>>1
    if (mid+1) * y > x && mid * y <= x || (mid+1) * y >= x && mid * y < x {
        if (mid+1) * y == x {
            return mid+1
        }
        return mid
    }
    if (mid+1) * y > x && mid * y > x {
        return binarySearch(left, mid-1, x, y)
    }
    if (mid+1) * y < x && mid * y < x {
        return binarySearch(mid+1, right, x, y)
    }

    return 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

