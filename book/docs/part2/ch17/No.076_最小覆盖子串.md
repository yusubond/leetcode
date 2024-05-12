## 76 最小覆盖子串-困难

题目：

给你一个字符串 `s` 、一个字符串 `t` 。返回 `s` 中涵盖 `t` 所有字符的最小子串。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""` 。



分析：

详见 readme

```go
// date 2023/11/28
func minWindow(s string, t string) string {
    tset := make(map[byte]int, 32)
    for i := 0; i < len(t); i++ {
        tset[t[i]]++
    }
    scnt := make(map[byte]int, 32)
    isContainT := func() bool {
        for k, v := range tset {
            if scnt[k] < v {
                return false
            }
        }
        return true
    }
    ans := -1
    al, ar := 0, 0
    left, right := 0, 0
    n := len(s)
    for right < n {
        scnt[s[right]]++
        right++

        for isContainT() {
            if ans == -1 {
                ans = right - left
                al, ar = left, right
            } else if right - left < ans {
                ans = right - left
                al, ar = left, right
            }
            scnt[s[left]]--
            if scnt[s[left]] == 0 {
                delete(scnt, s[left])
            }
            left++
        }
    }
    if ans == -1 {
        return ""
    }
    return string(s[al:ar])
}
```

