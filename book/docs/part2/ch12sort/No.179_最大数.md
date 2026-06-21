## 179 最大数-中等

给定一组非负整数 `nums`，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

**注意：**输出结果可能非常大，所以你需要返回一个字符串而不是整数。



分析：

对数组进行字母表排序。

```go
// date 2026.06.21
func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	for i, v := range nums {
		str[i] = strconv.Itoa(v)
	}

	// sort num 符合字母表排序
	sort.Slice(str, func(i, j int) bool {
		if str[i][0] == str[j][0] {
			return str[i]+str[j] > str[j]+str[i]
		}
		return str[i] > str[j]
	})

	ans := strings.Join(str, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
```

