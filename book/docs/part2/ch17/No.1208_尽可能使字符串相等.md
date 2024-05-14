## 1208 尽可能使字符串相等

题目：

给你两个长度相同的字符串，s 和 t。

将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。

如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

> **示例 1：**
>
> ```
> 输入：s = "abcd", t = "bcdf", maxCost = 3
> 输出：3
> 解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "abcd", t = "cdef", maxCost = 3
> 输出：1
> 解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。
> ```
>
> **示例 3：**
>
> ```
> 输入：s = "abcd", t = "acde", maxCost = 0
> 输出：1
> 解释：a -> a, cost = 0，字符串未发生变化，所以最大长度为 1。
> ```



算法：

因为题目中并没有要求一定要从 s 字符串的第一个字符开始替换，所以只要能够有可以替换的就算。

所以选择滑动窗口，窗口大小不固定。

```go
func equalSubstring(s, t string, maxCost int) int {
  left, right := 0, 0
  n := len(s)
  sumCost, ans := 0, 0
  // 构造窗口
  for right < n {
    sumCost += abs(s[right], t[right])   // 计算每一点的费用
    right++   // 增加right指针，扩大窗口
    for sumCost > maxCost { // 因为求最大值，所以不满足条件时需要增加left指针，缩小窗口
      sumCost -= abs(s[left], t[left])
      left++
    }
    // 此时窗口已经形成，检查是否需要更新结果
    if right - left > ans {
      ans = right - left
    }
  }
  return ans
}

func abs(x, y uint8) int {
  if x > y {
    return int(x-y)
  }
  return int(y-x)
}
```

