## 1137 第N个泰波那契数-简单

题目：

泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。


分析：

递推。

```go
// date 2023/11/11
func tribonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n < 3 {
        return 1
    }
    t0, t1, t2 := 0, 1, 1
    t := 0
    for i := 3; i <= n; i++ {
        t = t0+t1+t2
        t0, t1, t2 = t1, t2, t
    }
    return t
}
```
