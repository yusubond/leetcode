## 394 字符串编码-中等

题目：

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: `k[encoded_string]`，表示其中方括号内部的 `encoded_string` 正好重复 `k` 次。注意 `k` 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 `k` ，例如不会出现像 `3a` 或 `2[4]` 的输入。



> **示例 1：**
>
> ```
> 输入：s = "3[a]2[bc]"
> 输出："aaabcbc"
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "3[a2[c]]"
> 输出："accaccacc"
> ```
>
> **示例 3：**
>
> ```
> 输入：s = "2[abc]3[cd]ef"
> 输出："abcabccdcdcdef"
> ```
>
> **示例 4：**
>
> ```
> 输入：s = "abc3[cd]xyz"
> 输出："abccdcdcdxyz"
> ```



分析：

这道题跟 224 类似，关键是遇到什么字符的时候，需要入栈；以及遇到什么字符开始出栈。

出栈的时候，要把内容一次性处理。

```go
// date 2023/12/19
func decodeString(s string) string {
    strStack := make([]string, 0, 16)

    for _, str := range s {
        if len(strStack) == 0 || len(strStack) > 0 && str != ']' {
            strStack = append(strStack, string(str))
        } else {
            // 遇到 ']'
            // 说明前面有可以处理的字符和数字
            // 开始出栈，先出栈 字符串
            tmp := ""
            for strStack[len(strStack)-1] != "[" {
                tmp = strStack[len(strStack)-1] + tmp  // 连接字符
                strStack = strStack[:len(strStack)-1]
            }
            // 出栈 '['
            strStack = strStack[:len(strStack)-1]

            // 出栈 数字
            idx, repeat := 0, ""
            for idx = len(strStack)-1; idx >= 0; idx-- {
                if strStack[idx] >= "0" && strStack[idx] <= "9" {
                    repeat = strStack[idx] + repeat
                } else {
                    break
                }
            }
            // 去掉数字
            strStack = strStack[:len(strStack)-len(repeat)]

            repeatVal, _ := strconv.Atoi(repeat)
            copyTmp := tmp
            for repeatVal > 1 {
                tmp += copyTmp
                repeatVal--
            }

            // 将重复完成的字符串在入栈
            strStack = append(strStack, tmp)
        }
    }

    res := ""
    for _, v := range strStack {
        res += v
    }

    return res
}
```

