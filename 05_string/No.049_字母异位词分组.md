## 49 字母异位词分组-中等

题目：

给你一个字符串数组，请你将 **字母异位词** 组合在一起。可以按任意顺序返回结果列表。

**字母异位词** 是由重新排列源单词的所有字母得到的一个新单词。

> **示例 1:**
>
> ```
> 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
> 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
> ```
>
> **示例 2:**
>
> ```
> 输入: strs = [""]
> 输出: [[""]]
> ```
>
> **示例 3:**
>
> ```
> 输入: strs = ["a"]
> 输出: [["a"]]
> ```



**解题思路**

所谓字母异位词，就是字母字符个数一样，且出现的次数也一样。

所以我们可以对每个字符串进行计数，统计每个字母出现的次数，并保存在数组中。

进而以数组为 key，进行查找合并。

```go
// date 2024/01/24
// 解法1
// 字符计数
func groupAnagrams(strs []string) [][]string {
    // key = [26]int
    // value = []string
    checkMap := map[[26]int][]string{}  // key = [26]int
    
    ans := make([][]string, 0, 16)

    for _, str := range strs {
        one := toCheck(str)
        checkMap[one] = append(checkMap[one], str)
    }

    for _, v := range checkMap {
        ans = append(ans, v)
    }

    return ans
}

func toCheck(str string) [26]int {
    res := [26]int{}
    for _, v := range str {
        res[v-'a']++
    }
    return res
}
```

