## 1052 爱生气的书店老板-中等

题目：

有一个书店老板，他的书店开了 `n` 分钟。每分钟都有一些顾客进入这家商店。给定一个长度为 `n` 的整数数组 `customers` ，其中 `customers[i]` 是在第 `i` 分钟开始时进入商店的顾客数量，所有这些顾客在第 `i` 分钟结束后离开。

在某些时候，书店老板会生气。 如果书店老板在第 `i` 分钟生气，那么 `grumpy[i] = 1`，否则 `grumpy[i] = 0`。

当书店老板生气时，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的。

书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 `minutes` 分钟不生气，但却只能使用一次。

请你返回 *这一天营业下来，最多有多少客户能够感到满意* 。



分析：

有些顾客是确定满意的，有些是确定不满意的，那么想求满意顾客的最大数量，其实就是求 minutes 窗口内不满意的顾客最大数量 extra，把这些变成满意的，那么总数量就是：

确定满意的 total + 【minutes窗口内的 extra】

```go
// date 2023/11/22
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    var total, extra, temp int
    n := len(customers)
    left, right := 0, 0
    for right < n {
        // 确定满意的顾客
        total += (1 - grumpy[right]) * customers[right]

        // 
        temp += grumpy[right] * customers[right]  // 不满意的顾客
        right++
        extra = max(extra, temp)
        for right - left >= minutes {
            temp -= grumpy[left] * customers[left]
            left++
        }
    }
    return total + extra
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

