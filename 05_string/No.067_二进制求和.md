## 67 二进制求和-简单

题目：

给你两个二进制字符串 `a` 和 `b` ，以二进制字符串的形式返回它们的和。

>**示例 1：**
>
>```
>输入:a = "11", b = "1"
>输出："100"
>```
>
>**示例 2：**
>
>```
>输入：a = "1010", b = "1011"
>输出："10101"
>```



解题思路：

这道题跟 415 题没有区别，直接从尾部开始相加即可。

```go
// date 2022/10/17
func addBinary(a string, b string) string {
    ans := ""

    v1, v2, carry := 0, 0, 0
    i, j := len(a)-1, len(b)-1
    temp := 0

    for i >= 0 || j >= 0 || carry > 0 {
        v1, v2 = 0, 0
        if i >= 0 {
            v1 = int(a[i] - '0')
        }
        if j >= 0 {
            v2 = int(b[j] - '0')
        }
        temp = v1 + v2 + carry
        carry = temp  / 2
        ans = fmt.Sprintf("%d", temp % 2) + ans
        i--
        j--
    }

    return ans
}
```

