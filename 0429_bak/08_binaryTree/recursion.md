### 递归Recursion

#### 相关题目

- 爬楼梯
- Pow(x, n)

#### 爬楼梯

思路分析

基本情况：n = 1, res = 1, n = 2, res = 2; f(n) = f(n-1) + f(n2)

```go
// 斐波那契数列 时间复杂度O(n) 空间复杂度O(1)
func climbStairs(n int) int {
  if n <= 2 {return n}
  c, p1, p2 := 0, 2, 1
  for i := 3; i <= n; i++ {
    c = p1 + p2
    p1, p2 = c, p1
  }
  return c
}
```

#### Pow(x,n)

```go
// 迭代快速幂等法
func myPow(x float64, n int) float64 {
  if n < 0 {
    x = float64(1) / x
    n = -n
  }
  res := float64(1)
  c := x
  for i := n; i >= 0; i >>= 1 {
    if i & 0x1 == 1 {
      res *= c
    }
    c = c * c
  }
  return res
}
```

