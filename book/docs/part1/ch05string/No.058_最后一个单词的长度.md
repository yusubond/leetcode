## 58 最后一个单词的长度-简单

给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。



分析：

简单的指针技巧。

给定一个字符串，返回其最后一个单词的长度。注意字符串末尾的空字符串。

```go
// date 2020/01/19
func lengthOfLastWord(s string) int {
  var res int
  for i := len(s) - 1; i >= 0; i-- {
    if res == 0 && s[i] == ' ' {continue}
    if s[i] == ' ' {return res}
    res++
  }
  return res
}
```

#### 
