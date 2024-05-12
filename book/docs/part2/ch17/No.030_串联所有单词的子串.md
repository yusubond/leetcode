## 30 串联所有单词的子串

题目：

给定一个字符串 `s` 和一个字符串数组 `words`**。** `words` 中所有字符串 **长度相同**。

 `s` 中的 **串联子串** 是指一个包含 `words` 中所有字符串以任意顺序排列连接起来的子串。

- 例如，如果 `words = ["ab","cd","ef"]`， 那么 `"abcdef"`， `"abefcd"`，`"cdabef"`， `"cdefab"`，`"efabcd"`， 和 `"efcdab"` 都是串联子串。 `"acdbef"` 不是串联子串，因为他不是任何 `words` 排列的连接。

返回所有串联子串在 `s` 中的开始索引。你可以以 **任意顺序** 返回答案。



分析：

滑动窗口。

注意判断子串是否在 words 中，要求单词的数量。

```go
// date 2023/11/20
func findSubstring(s string, words []string) []int {
    ans := make([]int, 0, 64)
    s1 := len(words)
    s2 := len(words[0])
    total := s1 * s2

    idx := 0
    for idx + total <= len(s) {
        s3 := string(s[idx:idx+total])
        if isInWords(s3, words) {
            ans = append(ans, idx)
        }
        idx++
    }

    return ans
}

func isInWords(s string, words []string) bool {
    set := make(map[string]int, 64)
    for i := 0; i < len(words); i++ {
        set[words[i]]++
    }

    size := len(words[0])
    base := 0
    for base + size <= len(s) {
        s1 := string(s[base:base+size])
        _, ok := set[s1]
        if !ok {
            return false
        }
        set[s1]--
        if set[s1] == 0 {
            delete(set, s1)
        }
        base += size
    }
    return len(set) == 0
}
```

