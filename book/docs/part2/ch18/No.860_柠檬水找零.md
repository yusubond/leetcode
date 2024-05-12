## 860 柠檬水找零-简单

题目：

在柠檬水摊上，每一杯柠檬水的售价为 `5` 美元。顾客排队购买你的产品，（按账单 `bills` 支付的顺序）一次购买一杯。

每位顾客只买一杯柠檬水，然后向你付 `5` 美元、`10` 美元或 `20` 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 `5` 美元。

注意，一开始你手头没有任何零钱。

给你一个整数数组 `bills` ，其中 `bills[i]` 是第 `i` 位顾客付的账。如果你能给每位顾客正确找零，返回 `true` ，否则返回 `false` 。



分析：

直接遍历，记录手里 5 块和 10 块的张数；找零优先支出 10 块，不足再找 5 块；如果找不开直接返回 false。

```go
func lemonadeChange(bills []int) bool {
    a, b := 0, 0

    for _, v := range bills {
        if v == 5 {
            a++
        } else if v == 10 {
            if a == 0 {
                return false
            }
            a--
            b++
        } else if v == 20 {
            if a > 0 && b > 0 {
                a--
                b--
            } else if a >= 3 {
                a -= 3
            } else {
                return false
            }
        }
    }

    return true
}
```

