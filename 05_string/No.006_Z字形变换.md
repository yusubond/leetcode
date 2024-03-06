## Z字形变换-中等

题目：

将一个给定字符串 `s` 根据给定的行数 `numRows` ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 `"PAYPALISHIRING"` 行数为 `3` 时，排列如下：

```
P   A   H   N
A P L S I I G
Y   I   R
```

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：`"PAHNAPLSIIGYIR"`。

请你实现这个将字符串进行指定行数变换的函数：

```
string convert(string s, int numRows);
```



**解题思路**

按行存储。

Z字形变化的周期为 r + r - 2。

```go
// date 2024/03/06
func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	mat := make([][]byte, numRows)
	r := 0
	// 变化周期为 r + r - 2
	t := numRows*2 - 2
	for i, ch := range s {
		mat[r] = append(mat[r], byte(ch))
		if i%t < numRows-1 {
			r++
		} else {
			r--
		}
	}
	return string(bytes.Join(mat, nil))
}
```

