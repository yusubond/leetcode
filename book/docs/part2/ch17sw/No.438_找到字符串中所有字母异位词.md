## 438 找到字符串中所有字母异位词-中等

题目：

给定两个字符串 `s` 和 `p`，找到 `s` 中所有 `p` 的 **异位词** 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

**异位词** 指由相同字母重排列形成的字符串（包括相同的字符串）。



**解题思路**

- 滑动窗口，详见解法1。具体为，先处理 p 字符串，将其每个元素存入 map，记为 mp；滑动窗口遍历字符串s，并将元素存入 map，记为 ms，当滑动窗口区间等于字符串p的长度时，比较 ms, mp 是否相等，如果相等那么 left 就是答案之一。





```go
// date 2024/01/23
// 解法1 滑动窗口
func findAnagrams(s string, p string) []int {
	s1, s2 := len(s), len(p)
	mp := make(map[uint8]int, 16)
	for i := 0; i < s2; i++ {
		mp[p[i]]++
	}

	ans := make([]int, 0, 16)
	left, right := 0, 0
	ms := make(map[uint8]int, 16)

	for right < s1 {
		ms[s[right]]++
		right++
		if right-left >= s2 {
			if isSame(ms, mp) {
				ans = append(ans, left)
			}
			ms[s[left]]--
			if ms[s[left]] == 0 {
				delete(ms, s[left])
			}
			left++
		}
	}

	return ans
}

func isSame(ms, mp map[uint8]int) bool {
	if len(ms) != len(mp) {
		return false
	}
	for k, v := range ms {
		ov, ok := mp[k]
		if !ok || ov != v {
			return false
		}
	}
	return true
}
```

