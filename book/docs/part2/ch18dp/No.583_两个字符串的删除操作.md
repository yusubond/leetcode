## 583 两个字符串的删除操作-中等

题目：

给定两个单词 word1 和 word2，返回使得 word1 和 word2 **相同**所需的最小步数。每步可以删除任意一个字符串中的一个字符。



**关键观察**

这题与 [No.1143 最长公共子序列](../No.1143_最长公共子序列/) 和 [No.72 编辑距离](../No.072_编辑距离/) 高度相关。

要使两个字符串相同，等价于：保留它们的最长公共子序列（LCS），删除其余字符。因此：

```
最小删除步数 = len(word1) + len(word2) - 2 × LCS(word1, word2)
```

例如 `word1 = "sea"`, `word2 = "eat"`：
- LCS = "ea"，长度为 2
- 最小步数 = 3 + 3 - 2×2 = 2
- 操作：删除 word1 的 's'，删除 word2 的 't'，得到 "ea"



**解法一：直接 DP（定义删除次数）**

定义 `dp[i][j]` 表示使 `word1[0:i]` 和 `word2[0:j]` 相同的最小删除步数。

转移方程：

```
1. word1[i-1] == word2[j-1]
    dp[i][j] = dp[i-1][j-1]  // 字符相等，无需删除

2. word1[i-1] != word2[j-1]
    dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
    // dp[i-1][j] + 1：删除 word1 的当前字符
    // dp[i][j-1] + 1：删除 word2 的当前字符
```

初始条件：
- `dp[i][0] = i`（word2 为空，需删光 word1 前 i 个字符）
- `dp[0][j] = j`（word1 为空，需删光 word2 前 j 个字符）

```go
// date 2026/07/12
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
        dp[i][0] = i // word2 为空，删光 word1
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j // word1 为空，删光 word2
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
```



**解法二：LCS 转化（更简洁）**

先求 LCS 长度，再用公式计算：

```go
// date 2026/07/12
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    lcs := longestCommonSubsequence(word1, word2)
    return m + n - 2*lcs
}

func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}
```



**空间优化（滚动数组，解法一）**

`dp[i][j]` 依赖上方 `dp[i-1][j]`、左方 `dp[i][j-1]`、左上角 `dp[i-1][j-1]`。压缩到一维时，左上角需要 `prev` 变量保存。

```go
// date 2026/07/12
func minDistance1D(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([]int, n+1)
    for j := 0; j <= n; j++ { // 第 0 行：word1 为空，需删除 word2 前 j 个
        dp[j] = j
    }

    for i := 1; i <= m; i++ {
        prev := dp[0] // dp[i-1][j-1]，初始为上一行第 0 列
        dp[0] = i     // 第 0 列：word2 为空，需删光 word1 前 i 个
        for j := 1; j <= n; j++ {
            temp := dp[j] // 更新前的 dp[j] = 上一行 dp[i-1][j]（上方值）
            if word1[i-1] == word2[j-1] {
                dp[j] = prev // 相等，无需删除，直接继承左上角
            } else {
                dp[j] = min(dp[j], dp[j-1]) + 1
                //         ↑        ↑
                //       上方      左方
                //  dp[j]    本行未覆盖 → 上一行 dp[i-1][j]
                //  dp[j-1]  本行已覆盖 → 本行   dp[i][j-1]
            }
            prev = temp // 本轮旧 dp[j] 成为下一列 (j+1) 的左上角
        }
    }
    return dp[n]
}
```

> `min` 和 `max` 用 Go 1.21+ 内置函数即可；旧版本需自行定义。



**复杂度对比**

| 解法 | 时间 | 空间 |
|------|------|------|
| 二维 DP（直接） | O(m·n) | O(m·n) |
| 二维 DP（LCS 转化） | O(m·n) | O(m·n) |
| 一维滚动数组 | O(m·n) | O(n) |



**与相邻题的关系**

| 题目 | 操作 | 核心关系 |
|------|------|---------|
| No.72 编辑距离 | 插入、删除、替换 | 本题只保留删除，是编辑距离的特例 |
| No.1143 最长公共子序列 | 保留字符 | LCS 是本题的等价转化：`步数 = m+n-2×LCS` |
| **No.583（本题）** | 仅删除 | `dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1` |

本质上，No.583 的 DP 表与编辑距离的 DP 表结构相同，只是缺少了替换操作 `dp[i-1][j-1] + 1` 这一分支。
