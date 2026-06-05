## 415 字符串相加-简单

题目：

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。



算法分析：

从尾部到头部依次计算。

```go 
// date 2020/03/28
func addStrings(num1 string, num2 string) string {
    ans := ""
    carry, v1, v2 := 0, 0, 0
    i, j := len(num1)-1, len(num2)-1
    for i >= 0 || j >= 0 || carry != 0 {
        v1, v2 = 0, 0
        if i >= 0 {
            v1 = int(num1[i] - '0')
        }
        if j >= 0 {
            v2 = int(num2[j] - '0')
        }
        t := carry + v1 + v2
        carry = t / 10
        ans = fmt.Sprintf("%d", t % 10) + ans
        i--
        j--
    }
    return ans
}
```

