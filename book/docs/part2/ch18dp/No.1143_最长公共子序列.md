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
	dp := make([][]int, s1+1)
	for i := 0; i <= s1; i++ {
		dp[i] = make([]int, s2+1)
	}
	for i := 1; i <= s1; i++ {
		for j := 1; j <= s2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[s1][s2]
}
```


空间优化（滚动数组压缩到一维，空间 O(n)）：

LCS 与 No.072 编辑距离**完全同构**：转移同样依赖三个邻居——上方 `dp[i-1][j]`、左方 `dp[i][j-1]`、**左上角** `dp[i-1][j-1]`（仅当字符相等时用）。因此滚动技巧一模一样：左上角会在 `dp[j-1]` 被覆盖后丢失，必须用 `prev` 变量提前存下来。

唯一比编辑距离更简单的地方：**第 0 行/列恒为 0**（空串的 LCS 长度是 0），所以 `dp` 初始化为全 0 即可，无需每行重设 `dp[0]`，`prev := dp[0]` 直接取 0。

```go
// date 2026/06/20
func longestCommonSubsequence1D(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([]int, n+1)
    for i := 1; i <= m; i++ {
        prev := dp[0] // dp[i-1][j-1]，初始为上一行第 0 列（恒为 0）
        for j := 1; j <= n; j++ {
            temp := dp[j] // 更新前的 dp[j] = 上一行 dp[i-1][j]（上方值），先存下来
            if text1[i-1] == text2[j-1] {
                dp[j] = prev + 1 // 左上角 + 1
            } else {
                dp[j] = max(dp[j], dp[j-1])
                //          ↑       ↑
                //        上方     左方
                //  dp[j]    本行未覆盖 → 上一行 dp[i-1][j]
                //  dp[j-1]  本行已覆盖 → 本行   dp[i][j-1]
            }
            prev = temp // 本轮旧 dp[j] 成为下一列 (j+1) 的左上角
        }
    }
    return dp[n]
}
```

> `max` 用 Go 1.21+ 内置函数即可；旧版本需自行定义。

复杂度：时间 O(m·n)，空间 O(n)。
