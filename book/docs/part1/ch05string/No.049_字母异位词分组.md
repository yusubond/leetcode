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

字母异位词的核心特征是**每个字母的出现次数相同**。所以思路是给每个字符串计算一个规范化的 key，相同 key 的归入同一组。区别在于 key 怎么构造：

- **解法1：计数数组作为 key**（Go 原生支持 `[26]int` 作为 map key）。对每个字符串遍历一次得到 26 个字母的计数数组，直接用作 key。时间复杂度 O(n·k)，空间 O(n)，写法最简洁。但因为 key 是定长 26 的数组，Go 运行时哈希时需遍历整个数组，字符串较短时数组哈希的常数开销相对较大。
- **解法2：排序作为 key**。将每个字符串排序，排序后的字符串作为 key。时间复杂度 O(n·k log k)，思路最直观，适合 k 很小（短字符串）的场景。
- **解法3：计数编码为字符串 key**。先计数得到 `[26]int`，再将非零字符编码成紧凑字符串（如 `"a3b1c2"`）作为 key。时间复杂度 O(n·k)，但 Go 对 string 类型的哈希比数组哈希快，只遍历非零字符，常数更优，尤其 k 较大时优势明显。

```go
// date 2024/01/24
// 解法1
// 计数数组作为 key（[26]int 直接当 map key）
// 时间复杂度 O(N*K)，空间复杂度 O(N)
// N 为字符串个数，K 为字符串平均长度
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

```go
// date 2026/06/23
// 解法2
// 排序作为 key
// 时间复杂度 O(N * K log K)，空间复杂度 O(N)
// 思路最直观，适合字符串很短的情况
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, s := range strs {
        b := []byte(s)
        sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
        key := string(b)
        groups[key] = append(groups[key], s)
    }
    res := make([][]string, 0, len(groups))
    for _, v := range groups {
        res = append(res, v)
    }
    return res
}
```

```go
// date 2026/06/23
// 解法3
// 计数编码为字符串 key（紧凑编码，只写非零字符）
// 时间复杂度 O(N*K)，空间复杂度 O(N)
// 字符串 key 的哈希开销小于 [26]int 数组 key，常数更优
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, s := range strs {
        key := encode(s)
        groups[key] = append(groups[key], s)
    }
    res := make([][]string, 0, len(groups))
    for _, v := range groups {
        res = append(res, v)
    }
    return res
}

func encode(s string) string {
    cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++
    }
    // 构建紧凑的字符串 key，只包含出现过的字符
    // 例如 "aab" → "a2b1"
    var sb strings.Builder
    for i, c := range cnt {
        if c > 0 {
            sb.WriteByte(byte('a' + i))
            sb.WriteString(strconv.Itoa(c))
        }
    }
    return sb.String()
}
```

