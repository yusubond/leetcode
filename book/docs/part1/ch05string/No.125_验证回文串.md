## 125 验证回文串-简单

题目：

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。



**解题思路**

去除不需要比较的字符，然后双指针首尾对撞。

```go
// date 2022/10/17
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    var res string
    for i := 0; i < len(s); i++ {
        if 'a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9' {
            res += string(s[i])
        }
    }
    i, j := 0, len(res)-1
    for i < j {
        if res[i] != res[j] {
            return false
        }
        i++
        j--
    }
    return true
}
// 写法2
// 双指针，直接在原字符串上操作，只判断字母和数字字符
func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left, right := 0, len(s)-1

    for left < right {
        if !isCharNum(s[left]) {
            left++
            continue
        }
        if !isCharNum(s[right]) {
            right--
            continue
        }
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }

    return true
}

func isCharNum(x uint8) bool {
    if x >= 'a' && x <= 'z' || x >= '0' && x <= '9' {
        return true
    }
    return false
}
```

