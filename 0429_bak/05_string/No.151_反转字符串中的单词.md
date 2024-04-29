## 151 反转字符串中的单词-中等

题目：

给你一个字符串 `s` ，请你反转字符串中 **单词** 的顺序。

**单词** 是由非空格字符组成的字符串。`s` 中使用至少一个空格将字符串中的 **单词** 分隔开。

返回 **单词** 顺序颠倒且 **单词** 之间用单个空格连接的结果字符串。

**注意：**输入字符串 `s`中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。



分析：

1. 先去除头尾的空格无效字符
2. 从尾部开始遍历，每找到一个单词就追加到结果中



```go
// date 2023/12/10
func reverseWords(s string) string {
    res := ""
    i, j := 0, len(s)-1
    for j >= 0 && s[j] == ' ' {
        j--
    }
    for i >= 0 && s[i] == ' ' {
        i++
    }
    start := i
    end := j

    for j >= start {
        for j >= start && s[j] != ' ' {
            j--
        }
        // now, the [j+1:end+1] is a word
        if res == "" {
            res = string(s[j+1:end+1])
        } else {
            res += " "
            res += string(s[j+1:end+1])
        }
        for j >= start && s[j] == ' ' {
            j--
        }
        end = j
    }
    return res
}
```

