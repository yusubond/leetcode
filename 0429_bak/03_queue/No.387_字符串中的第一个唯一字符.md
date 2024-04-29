## 387 字符串中的第一个唯一字符

题目：

给定一个字符串 `s` ，找到 *它的第一个不重复的字符，并返回它的索引* 。如果不存在，则返回 `-1` 。



分析：

利用 map 存储字符出现的次数，如果首次出现，map 值为下标索引，否则为 -1。

然后，对 map 进行遍历找最小值。

```go
// date 2023/11/30
func firstUniqChar(s string) int {
    chSet := make(map[rune]int, 32)
    for i, c := range s {
        _, ok := chSet[c]
        if ok {
            chSet[c] = -1
        } else {
            chSet[c] = i
        }
    }
    ans := len(s)
    for _, v:= range chSet {
        if v != -1 && v < ans {
            ans = v
        }
    }
    if ans == len(s) {
        return -1
    }
    return ans
}
```

