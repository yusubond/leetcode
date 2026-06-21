## 139 单次拆分-中等

题目：

给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。


分析：

以下代码可AC。

我们定义 dp[i] 表示字符串 s[0...i] 可以拆分成若干个字典中存在的单词，那么对于s[0...i]来说，任意分割点j，满足0 <= j < i，那么有如下递推公式：

```
dp[i] = dp[j] && s[j:i] in wordset
```

```go
// date 2023/11/13
func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool, len(wordDict))
    for _, v := range wordDict {
        wordSet[v] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i := 0; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
```


以下代码不可AC。

这段代码的逻辑是依次匹配单词，如果成功继续匹配。

无法通过这类 case：

```
s = "cars"

words = ["car", "ca", "rs"]
```

```go
// date 2023/11/13
func wordBreak(s string, wordDict []string) bool {
    start := 0
    end := 0
    total := len(s)
    i := 0

    for i < total {
        old := i
        start = i
        end = 0
        for _, word := range wordDict {
            size := len(word)
            end = start + size

            if end > total {
                continue
            }
            tgt := s[start:end]
            //if isSame(word, tgt) && end == total {
            //    return true
            //}
            if isSame(word, tgt) {
                i = end
                break
            }
        }
        if i == old {
           return false
        }
    }
    if i == total {
        return true
    }
    return false
}

func isSame(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            return false
        }
    }
    return true
}
```
