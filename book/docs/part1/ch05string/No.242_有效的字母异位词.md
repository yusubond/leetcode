## 242 有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



分析

算法1：题目中说明了只包含小写字母，因此可以用数组判断。

```go
func isAnagram(s string, t string) bool {
  if len(s) != len(t) {return false}
  ms, mt := [26]int{}, [26]int{}
  for i := 0; i < len(s); i++ {
    ms[s[i] - 'a']++
    mt[t[i] - 'a']++
  }
  return ms == mt
}
```

#### 