## 224 基本计算器-困难

题目：

给你一个字符串表达式 `s` ，请你实现一个基本计算器来计算并返回它的值。

注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 `eval()` 。



> **示例 1：**
>
> ```
> 输入：s = "1 + 1"
> 输出：2
> ```
>
> **示例 2：**
>
> ```
> 输入：s = " 2-1 + 2 "
> 输出：3
> ```
>
> **示例 3：**
>
> ```
> 输入：s = "(1+(4+5+2)-3)+(6+8)"
> 输出：23
> ```



分析：

还是入栈、出栈的思路，有几个点需要注意。

- 可能存在空格，需要跳过
- `-`是有效的，所以可能出现负负得正的情况，所以记录每次计算出来的结果
- 遇到左括号，需要重新计算，主要保存之前的结果和符号
- 遇到右括号，取出保存的结果和符号，更新当前结果

```go
// date 2023/12/18
func calculate(s string) int {
    i := 0
    res := 0
    sign := 1 // 保存加减状态
    stack := make([]int, 0, 16)  // 保存括号外的中间结果和加减符号

    n := len(s)
    for i < n {
        if s[i] == ' ' {
            i++
        } else if s[i] <= '9' && s[i] >= '0' {
            // find a num
            base, v := 10, int(s[i] - '0')
            for i+1 < n && s[i+1] <= '9' && s[i+1] >= '0' {
                v = v * base + int(s[i+1] - '0')
                i++
            }
            res += v * sign
            i++
        } else if s[i] == '+' {
            sign = 1
            i++
        } else if s[i] == '-' {
            sign = -1
            i++
        } else if s[i] == '(' { // 遇到左括号，保存之前的结果，重新计算
            stack = append(stack, res, sign)
            res = 0
            sign = 1
            i++
        } else if s[i] == ')' { // 取出之前的结果和符号，计算
            if len(stack) > 1 {
                res = res * stack[len(stack)-1] + stack[len(stack)-2]
                stack = stack[:len(stack)-2]
            }
            //res = res
            i++
        }
    }
    return res
}
```

