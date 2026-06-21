## 17 电话号码的字母组合-中等

题目：

给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。答案可以按 **任意顺序** 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png)



> **示例 1：**
>
> ```
> 输入：digits = "23"
> 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
> ```
>
> **示例 2：**
>
> ```
> 输入：digits = ""
> 输出：[]
> ```
>
> **示例 3：**
>
> ```
> 输入：digits = "2"
> 输出：["a","b","c"]
> ```



分析：

这个题目并不算严格意义上的回溯，一般的 DFS 即可。

把每个数字代表的字母形成不同的数组，存放在 map 里。

正序遍历数字组合，然后深度优先搜索；当数字组合为空时，即找到一个答案，追加到结果集。



```go
// date 2023/12/25
func letterCombinations(digits string) []string {
    res := make([]string, 0, 16)

    if len(digits) == 0 {
        return res
    }

    chars := map[string][]string{
        "2": []string{"a", "b", "c"},
        "3": []string{"d", "e", "f"},
        "4": []string{"g", "h", "i"},
        "5": []string{"j", "k", "l"},
        "6": []string{"m", "n", "o"},
        "7": []string{"p", "q", "r", "s"},
        "8": []string{"t", "u", "v"},
        "9": []string{"w", "x", "y", "z"},
    }

    var backtrack func(digs string, temp string)
    backtrack = func(digs string, temp string) {
        if len(digs) == 0 {
            res = append(res, temp)
            return
        }
        v := string(digs[0])
        all, ok := chars[v]
        if ok {
            for _, s := range all {
                backtrack(string(digs[1:]), temp + s)
            }
        }
    }

    backtrack(digits, "")

    return res
}
```

