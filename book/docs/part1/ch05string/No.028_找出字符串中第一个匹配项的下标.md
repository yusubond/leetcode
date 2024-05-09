## 28 找出字符串中第一个匹配项的下标-简单

题目：

给你两个字符串 `haystack` 和 `needle` ，请你在 `haystack` 字符串中找出 `needle` 字符串的第一个匹配项的下标（下标从 0 开始）。如果 `needle` 不是 `haystack` 的一部分，则返回 `-1` 。



分析：

单指针，逐位匹配。

超过匹配长度时，注意剪枝。

```go
// date 2023/12/10
func strStr(haystack string, needle string) int {
    s1, s2 := len(haystack), len(needle)
    if s1 < s2 {
        return -1
    }
    for i := 0; i < s1; i++ {
        if i + s2 > s1 {
            return -1
        }
        left := i
        j := 0
        for j < s2 && haystack[left] == needle[j] {
            left++
            j++
        }
        if j == s2 {
            return i
        }
    }
    return -1
}
```

