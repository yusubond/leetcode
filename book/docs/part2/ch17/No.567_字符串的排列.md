## 567 字符串的排列-中等

题目：

给你两个字符串 `s1` 和 `s2` ，写一个函数来判断 `s2` 是否包含 `s1` 的排列。如果是，返回 `true` ；否则，返回 `false` 。

换句话说，`s1` 的排列之一是 `s2` 的 **子串** 。



分析：

排列是说两个字符串中每个字符出现的次数一样。

初始化两个数组，分别记录s1和s2固定窗口中每个字符的出现次数，如果一致，则返回 true。

```go
// date 2023/11/20
func checkInclusion(s1 string, s2 string) bool {
    n, m := len(s1), len(s2)
    if n > m {
        return false
    }
    var ct1, ct2 [26]int
    for i := range s1 {
        ct1[s1[i]-'a']++
        ct2[s2[i]-'a']++
    }
    if ct1 == ct2 {
        return true
    }

    for i := n; i < m; i++ {
        // update ct2
        ct2[s2[i]-'a']++
        ct2[s2[i-n]-'a']--
        if ct1 == ct2 {
            return true
        }
    }

    return false
}
```

