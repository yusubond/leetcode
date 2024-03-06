## 5 最长回文字符串-中等

题目：

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



分析：

从中心到两边的双指针。

```go
// date 2023/12/09
func longestPalindrome(s string) string {
	res := ""
	n := len(s)
	var checkPalindrome func(left, right int)
	checkPalindrome = func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if len(s[left+1:right]) > len(res) {
			res = s[left+1 : right]
		}
	}

	for i := 0; i < n; i++ {
		checkPalindrome(i, i)
		checkPalindrome(i, i+1)
	}

	return res
}
```

