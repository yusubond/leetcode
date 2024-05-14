## 424 替换后的最长重复字符

题目要求：

给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。

在执行上述操作后，返回包含相同字母的最长子字符串的长度。
题目链接：https://leetcode.cn/problems/longest-repeating-character-replacement



算法分析：

这道题也是滑动窗口的算法，关键是滑动窗口内我们关心什么样的数据？

在一个连续的区间内，如果我们能够统计出每个字符出现的次数ct，并找出现次数最多的字符次数maxCt，那么剩下的就是其他字符出现的次数`right - left - maxCt`，如果这个值小于等于k，那就是一个可行解。

当窗口不满足条件（即`right - left - maxCt > k`）时，将最左端元素移除，统计次数减一即可。



1. 假设闭区间 `[left, right]`表示一个窗口；
2. ct统计窗口内各个字符出现的次数，并不断更新

```go
// date 2022/09/30
func characterReplacement(s string, k int) int {
	left, right := 0, 0
	maxCt, ans, n := 0, 0, len(s)
	ct := [26]int{}
	var idx uint8
	// 构造窗口
	for right < n {
		idx = s[right] - 'A'
		ct[idx]++
		// 窗口长度只能增大或者维持不变, left指针增加0/1位
		// maxCt 一直维护最大值，因为我们求的是最长，如果找不到更长，不影响结果
		right++
		if maxCt < ct[idx] {
			maxCt = ct[idx]
		}
		// 窗口构造完成
		if right - left - maxCt <= k {
			ans = max(ans, right-left)
		} else {
			// 超过可替换次数k,开始增加left
			// 因为求最大值，left移动1位即可
			idx = s[left] - 'A'
			ct[idx]--
			left++
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

