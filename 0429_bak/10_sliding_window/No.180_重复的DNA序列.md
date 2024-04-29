## 180 重复的DNA序列

题目：

**DNA序列** 由一系列核苷酸组成，缩写为 `'A'`, `'C'`, `'G'` 和 `'T'`.。

- 例如，`"ACGAATTCCG"` 是一个 **DNA序列** 。

在研究 **DNA** 时，识别 DNA 中的重复序列非常有用。

给定一个表示 **DNA序列** 的字符串 `s` ，返回所有在 DNA 分子中出现不止一次的 **长度为 `10`** 的序列(子字符串)。你可以按 **任意顺序** 返回答案。



分析：

滑动窗口。

通过 map 记录出现过的子字符串的次数，最终在梳理结果。

```go
// date 2023/11/20
func findRepeatedDnaSequences(s string) []string {
    left, right := 0, 0
    ans := make([]string, 0, 64)
    set := make(map[string]int, 64)

    for right < len(s) {

        right++
        if right - left >= 10 {
            ss := string(s[left:right])
            set[ss]++
            left++
        }
    }

    for k, v := range set {
        if v > 1 {
            ans = append(ans, k)
        }
    }

    return ans
}
```

