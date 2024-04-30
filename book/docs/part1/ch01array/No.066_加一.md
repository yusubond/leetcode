## 66 加一-简单

题目：

给定一个由 **整数** 组成的 **非空** 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储**单个**数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



分析：

细节题，注意进位

```go
//date 2022/09/20
func plusOne(digits []int) []int {
    carry, temp := 1, 0
    for i := len(digits)-1; i >= 0; i-- {
        temp = digits[i]+carry
        carry = temp / 10
        digits[i] = temp % 10
    }
    if carry == 1 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
```

