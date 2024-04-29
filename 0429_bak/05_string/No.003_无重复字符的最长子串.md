## 3 无重复字符的最长子串

题目：

给定一个字符串 `s` ，请你找出其中不含有重复字符的 **最长子串** 的长度。



**解题思路**

- 滑动窗口，详见解法1。具体是初始化一个数组 list，保存遍历过的元素；以及 left，right 两个指针，移动 right 并将元素加入 list；前置判断，如果 list 不为空，且包含当前元素，则移除最早加入的元素，Left 指针递增；左右指针的差就是不含重复字符的子串，求其最大长度即可。

```go
// date 2022/09/29
func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0
    ans, n := 0, len(s)
    list := make([]uint8, 0, 64)  // 需要判断是否重复，list存在窗口的元素
    // 构造窗口
    for right < n {
        // 当窗口不在满足条件，即当前元素s[right]已经出现重复
        // 增加left指针，使窗口缩写
        for len(list) != 0 && isContains(s[right], list) {
            list = list[1:]
            left++
        }
        // 增加right
        list = append(list, s[right])
        right++
        if right - left > ans {
            ans = right - left
        }
    }
    return ans
}

func isContains(c uint8, list []uint8) bool {
    for _, v := range list {
        if v == c {
            return true
        }
    }
    return false
}
```

