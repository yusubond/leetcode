## 1295 统计位数为偶数的数字-简单

题目：

给你一个整数数组 `nums`，请你返回其中位数为 **偶数** 的数字的个数。



分析：

注意，题目中的位数是指十进制的位数。

```go
// date 2020/04/19
func findNumbers(nums []int) int {
    res := 0

    for _, v := range nums {
        if isEvenDigits(v) {
            res++
        }
    }

    return res
}

func isEvenDigits(num int) bool {
    dig := 0
    for num > 0 {
        dig++
        num /= 10
    }
    return dig & 0x1 == 0
}
```

