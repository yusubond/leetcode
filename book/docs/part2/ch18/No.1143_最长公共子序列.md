## 1143 最长公共子序列-中等

题目：


给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。



**解题思路**

动规，这题跟 No.72 编辑距离很像，思路是一样的。

定义 `lcs[i][j]` 表示两个字符串s1,s2中前 `i`,`j`个字符的最长公共子序列长度，那么有以下转移方程。

```
1. s1[i] == s2[j]
	lcs[i][j] = lcs[i-1][j-1] + 1

2. s1[i] != s2[j]
	lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1]
```


```go
// date 2023/11/13
func longestCommonSubsequence(text1 string, text2 string) int {
    s1, s2 := len(text1), len(text2)
    lcs := make([][]int, s1+1)
    for i := 0; i <= s1; i++ {
        lcs[i] = make([]int, s2+1)
    }
    for i := 0; i < s1; i++ {
        for j := 0; j < s2; j++ {
            if text1[i] == text2[j] {
                lcs[i+1][j+1] = lcs[i][j] + 1
            } else {
                lcs[i+1][j+1] = max(lcs[i][j+1], lcs[i+1][j])
            }
        }
    }
    return lcs[s1][s2]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```
