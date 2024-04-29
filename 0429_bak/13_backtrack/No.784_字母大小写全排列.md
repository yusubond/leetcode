## 784 字母大小写全排列-中等

题目：

给定一个字符串 `s` ，通过将字符串 `s` 中的每个字母转变大小写，我们可以获得一个新的字符串。

返回 *所有可能得到的字符串集合* 。以 **任意顺序** 返回输出。



> **示例 1：**
>
> ```
> 输入：s = "a1b2"
> 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
> ```
>
> **示例 2:**
>
> ```
> 输入: s = "3z4"
> 输出: ["3z4","3Z4"]
> ```



分析：

这道题的解题思路是 DFS。

- 遇到数字不变，继续深搜
- 遇到小写字母，变成大写，对大写和小写分别继续搜索
- 遇到大写字母，变成小写，对大写和小写分别继续搜索



```go
// date 2023/12/27
func letterCasePermutation(s string) []string {
    res := make([]string, 0, 16)
    
    n := len(s)

    var backtrack func(idx int, temp string)
    backtrack = func(idx int, temp string) {
        if idx == n {
            res = append(res, temp)
            return
        }
        char := string(s[idx])
        if char >= "0" && char <= "9" {
            backtrack(idx+1, temp + char)
        } else if char >= "a" && char <= "z" {
            char2 := strings.ToUpper(char)
            backtrack(idx+1, temp+char)
            backtrack(idx+1, temp+char2)
        } else if char >= "A" && char <= "Z" {
            char2 := strings.ToLower(char)
            backtrack(idx+1, temp+char)
            backtrack(idx+1, temp+char2)
        }
    }

    backtrack(0, "")


    return res
}
```

