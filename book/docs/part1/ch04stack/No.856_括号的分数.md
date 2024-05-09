## 856 括号的分数-中等

题目：

给定一个平衡括号字符串 `S`，按下述规则计算该字符串的分数：

- `()` 得 1 分。
- `AB` 得 `A + B` 分，其中 A 和 B 是平衡括号字符串。
- `(A)` 得 `2 * A` 分，其中 A 是平衡括号字符串。



> **示例 1：**
>
> ```
> 输入： "()"
> 输出： 1
> ```
>
> **示例 2：**
>
> ```
> 输入： "(())"
> 输出： 2
> ```
>
> **示例 3：**
>
> ```
> 输入： "()()"
> 输出： 2
> ```
>
> **示例 4：**
>
> ```
> 输入： "(()(()))"
> 输出： 6
> ```



分析：

还是入栈，出栈的思路。

遇到`)`开始出栈，通过栈中已有的元素来判断，题目中所说的三种情况：

- 如果前一个为-1，即temp == 0，表示得1分
- 如果前一个不为-1，即 stack 中的得分是 A+B 形式，需要累计；并且需要扩大1倍

```go
// date 2023/12/19
func scoreOfParentheses(s string) int {
    res := 0
    stack := make([]int, 0, 16) // score stack

    for _, str := range s {
        if str == '(' {
            stack = append(stack, -1)
        } else { // ')'
            // 取出之前的数据
            // 如果紧挨着 -1，得分为1
            // 否则就是 A+B 形式, 求和
            temp := 0
            for stack[len(stack)-1] != -1 {
                temp += stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }

            // 去掉 -1
            stack = stack[:len(stack)-1]

            if temp == 0 {
                stack = append(stack, 1)
            } else {
                stack = append(stack, 2*temp)
            }
        }
    }
    for _, v := range stack {
        res += v
    }

    return res
}
```

