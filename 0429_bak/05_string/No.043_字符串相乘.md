## 43 字符串相乘-中等

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。



分析：任务拆解，把大任务拆解成一个一个的小任务。

```go
// date 2020/03/28
/*
1. 先实现两个字符串相加；2.实现单个数字的相乘，3.实现零的追加。
*/
func multiply(num1 string, num2 string) string {
    if len(num1) == 0 { return num2 }
    if len(num2) == 0 { return num1 }
    if len(num1) == 1 && num1[0] == '0' { return "0" }
    if len(num2) == 1 && num2[0] == '0' { return "0" }
    var res string
    for i := len(num2)-1; i >= 0; i-- {
        t := product(num1, num2[i])
        t = appendZeros(t, len(num2)-1-i)
        res = addString(res, t)
    }
    return res
}

func appendZeros(s string, c int) string {
    if c <= 0 { return s }
    for c > 0 {
        s += fmt.Sprintf("%d", 0)
        c--
    }
    return s
}

func product(s1 string, s2 byte) string {
    var res string
    var carry, n1, t int
    i, n2 := len(s1)-1, int(s2 - '0')
    for i >= 0 || carry > 0 {
        n1 = 0
        if i >= 0 { n1 = int(s1[i]-'0') }
        t = n1 * n2 + carry
        carry = t / 10
        res = fmt.Sprintf("%d", t%10) + res
        i--
    }
    return res
}

func addString(s1, s2 string) string {
    var res string
    var carry, n1, n2, t int
    i, j := len(s1)-1, len(s2)-1
    for i >= 0 || j >= 0 || carry > 0 {
        n1, n2 = 0, 0
        if i >= 0 { n1 = int(s1[i] - '0') }
        if j >= 0 { n2 = int(s2[j] - '0') } 
        t = n1 + n2 + carry
        carry = t / 10
        res = fmt.Sprintf("%d", t%10) + res
        i--
        j--
    }
    return res
}
```
