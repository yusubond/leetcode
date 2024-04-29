## Math

[TOC]

### 1、相关题目

#### 7 整数反转

给你一个 32 位的有符号整数 `x` ，返回将 `x` 中的数字部分反转后的结果。

```go
// date 2022-09-15
func reverse(x int) int {
    mMax, mMin := 1 << 31 - 1, -1 << 31
    res, flag := 0, 1
    if x < 0 {
        flag, x = -1, -x
    }
    for x > 0 {
        res = res * 10 + x % 10
        x /= 10
    }
    res *= flag
    if res > 0 && res > mMax || res < 0 && res <= mMin {
        res = 0
    }
    return res
}
```



#### 8 String to Integer

https://leetcode.com/problems/string-to-integer-atoi/

```go
/* 细节题: 需要各种异常场景
1. 忽略开头的空白字符
2. 数字的符号
3. 溢出
4. 无效输入
*/
func myAtoi(str string) int {
    myMAX := 1 << 31 - 1
    myMIN := -1 << 31
    sign := 1
    var res int
    i := 0
    n := len(str)
    for i < n && str[i] == ' ' {i++}
    if i == n {return res}
    if str[i] == '+' {
        i++
    } else if str[i] == '-' {
        i++
        sign = -1
    }
    for ; i < n; i++ {
        if str[i] - '0' < 0 || str[i] - '0' > 9 {break}
        t := int(str[i] - '0')
        if res > myMAX / 10 || res == myMAX / 10 && t > myMAX % 10 {
            if sign == 1 {
                return myMAX
            } else if sign == -1 {
                return myMIN
            }
        }
        res = res * 10 + t
    }
    return res * sign
}
```



#### 9 Palindrome Number

```go
// 算法1：直接比较数字的每一位，更通用的算法
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    d := 1
    for x / d >= 10 {
        d *= 10
    }
    for x > 0 {
        // 最高位
        q := x / d
        // 最低位
        p := x % 10
        if q != p { return false }
        // 移动至次高位，同时去掉最后一位
        x = x % d / 10
        d /= 100
    }
    return true
}
// 算法2：将数字从中间分开，判断左边和右边的数字是否相等
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    left := 0
    for x > left {
        left = left * 10  + x % 10
        x /= 10
    }
    return x == left || x == left/10
}
```

思路学习：算法2对于判断链表同样适用。参考题目234https://leetcode.com/problems/palindrome-linked-list/

#### 15 三数之和【M】

题目要求：给你一个包含 *n* 个整数的数组 `nums`，判断 `nums` 中是否存在三个元素 *a，b，c ，*使得 *a + b + c =* 0 ？请你找出所有满足条件且不重复的三元组。

**注意：**答案中不可以包含重复的三元组。

思路分析



```go
// date 2020/03/29
func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    res, n, sum := make([][]int, 0), len(nums), -1
    left, right := 0, n-1
    for i := 0; i < n; i++ {
        // 升序排列，大于零则退出
        if nums[i] > 0 { break }
        // 去重
        if i > 0 && nums[i] == nums[i-1] { continue }
        left, right = i+1, n-1
        for left < right {
            sum = nums[i] + nums[left] + nums[right]
            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else if sum == 0 {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                // 去重
                for left < right && nums[left] == nums[left+1] { left++ }
                for left < right && nums[right] == nums[right-1] { right-- }
                left++
                right--
            }
        }
    }
    return res
}
```

#### 89 格雷编码

题目要求：给定一个格雷编码的总位数n，输出器格雷编码序列（格雷编码序列：相邻的两个数字只有一位不同）。

算法：镜像反射法

```go
// date 2020/01/11
/* 算法：G(n) = G(n-1) + R(n-1)  // R(n-1)为G(n-1)反序+head
0 0 00
  1 01
    11
    10
*/
func grayCode(n int) []int {
  res, head := make([]int, 0), 1
  res = append(res, 0)
  for i := 0; i < n; i++ {
    n := len(res) - 1
    for j = n; j >= 0; j-- {
      res = append(res, head + res[j])
    }
    head <<= 1
  }
  return res
}
```

#### 四数求和

思路分析

算法：原理同三数求和

```go
func fourSum(nums []int, target int) [][]int {
  sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
  })
  res, n, sum := make([][]int, 0), len(nums), 0
  l, r := 0, n-1
  for i := 0; i < n; i++ {
    // nums[i]可以是负数 if nums[i] > target {break}
    if i > 0 && nums[i] == nums[i-1] {continue}
    for j := i+1; j < n; j++ {
      if j > i+1 && nums[j] == nums[j-1] {continue}
      l, r = j+1, n-1
      for l < r {
        sum = nums[i] + nums[j] + nums[l] + nums[r]
        if sum == target {
          t := make([]int, 0)
          t = append(t, nums[i], nums[j], nums[l], nums[r])
          res = append(res, t)
          for l < r && nums[l] == nums[l+1] {l++}
          for l < r && nums[r] == nums[r-1] {r--}
          l++
          r--
        } else if sum < target {
          l++
        } else {
          r--
        }
      }
    }
  }
  return res
}
```

#### 191 位1的个数【简单】

题目链接：https://leetcode-cn.com/problems/number-of-1-bits/

编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为[汉明重量](https://baike.baidu.com/item/汉明重量)）。

思路分析：

```golang
func hammingWeight(num uint32) int {
    var c int
    for num > 0 {
        c += int(num & 0x1)
        num >>= 1
    }
    return c
}
```

#### 149 直线上最多的点数【H】

题目要求：给定二维平面上的n个点，求最多有多少个点在同一条直线上。

算法：哈希表，欧几里得最大公约数算法

```go
// data 2020/02/01
func maxPoints(points [][]int) int {
  if len(points) < 3 { return len(points) }
  n, count := len(points), make([]int, len(points))
  size, same := 1, 0
  for i := 0; i < n; i++ {
    m := make(map[[2]int]int, 0)
    count[i] = 1
    size, same = 1, 0
    for j := 0; j < n; j++ {
      if i == j { continue }
      // x坐标一样
      if points[i][0] == points[j][0] {
        if points[i][1] == points[j][1] {
          same++
        }
        size++
      } else {
     // x坐标不一样 求斜率
        dy := points[j][1] - points[i][1]
        dx := points[j][0] - points[i][0]
        g := gcd(dy, dx)
        if g != 0 {
          dy, dx = dy/g, dx/g
        }
        k := [2]int
        k[0], k[1] = dy, dx
        find := false
        for mk, mv := range m {
          if mk[0] == k[0] && mk[1] == k[1] {
            find = true
            m[mk]++
          }
        }
        if !find {
          m[k] = 1
        }
      }
    }
    for _, v := range m {
      if v + 1 > count[i] {
        count[i] = v + 1
      }
    }
    count[i] += same
    if count[i] < size {
      count[i] = size
    }
  }
  maxIndex := 0
  for i := 0; i < n; i++ {
    if count[i] > count[maxIndex] {
      maxIndex = i
    }
  }
  return count[maxIndex]
}

func gcd(x, y int) int {
  if y == 0 { return x }
  return gcd(y, x % y)
}
```

#### 231 2的幂【E】

题目要求：给定一个整数，判断其是否是2的幂次方。

思路分析：判断这个数是否不断的被2整数，直接结果等于1。如果某一次结果为奇数，直接返回false。

```golang
// date 2020/01/11
func isPowerOfTwo(n int) bool {
  for n > 1 {
    if n &0x1 == 1 {return false}
    n >>= 1
  }
  return n == 1
}
```

#### 292 Nim游戏

算法思路：只要是4的倍数，你一定输。

```go
// date 2020/01/09
func canWinNim(n int) bool {
  if n % 4 == 0 {
    return false
  }
  return true
}
```

#### 338 比特位计数

题目要求：https://leetcode-cn.com/problems/counting-bits/

思路分析：

```go
// date 2020/03/30
/*
递推，二进制的表示
如果当前为奇数，则在上一个偶数的基础上加一；如果当前为偶数，则和上一个偶数保持一致
*/
func countBits(num int) []int {
    res := make([]int, num+1)
    res[0] = 0
    for i := 1; i <= num; i++ {
        if i & 0x1 == 1 {
            res[i] = res[i-1]+1
        } else {
            res[i] = res[i>>1]
        }
    }
    return res
}
```

#### 461 汉明距离【简单】

题目链接：https://leetcode-cn.com/problems/hamming-distance/

两个整数之间的[汉明距离](https://baike.baidu.com/item/汉明距离)指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 `x` 和 `y`，计算它们之间的汉明距离。

思路分析：

```golang
func hammingDistance(x int, y int) int {
    return hammingWeight(x^y)
}

func hammingWeight(num int) int {
    var c int
    for num > 0 {
        c += num & 0x1
        num >>= 1
    }
    return c
}
```

#### 476 数字的补数

题目链接：https://leetcode-cn.com/problems/number-complement/

题目要求：给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。

思路分析：

```go
// date 2020/05/04
// 找到第一个比这个数大的二进制整数，进而求得全F,然后减去这个数即可
func findComplement(num int) int {
    x := 1
    for x <= num {
        x <<= 1
    }
    return x - 1 - num
}
```

#### 斐波那契数列

思路分析

记忆化递归

```go
func fib(N int) int {
  if N < 2 {return N}
  c, p1, p2 := 1, 1, 0
  for i := 2; i <= N; i++ {
  	c = p1+p2
    p1, p2 = c, p1
  }
  return c
}
```

#### 914 卡牌分组

题目要求：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/

思路分析：

第一种：暴力求解，先统计每个元素出现的次数C，然后对可能的X值进行两个判断，X必须是N的约数，也是C的约数。

```go
// date 2020/03/27
func hasGroupsSizeX(deck []int) bool {
    count := make([]int, 10000)
    // 统计每个数出现的次数
    for _, v := range deck {
        count[v]++
    }
    value := make([]int, 0)
    // 将出现的次数放入数组
    for _, c := range count {
        if c > 0 { value = append(value, c) }
    }
    // 暴露求解
    n := len(deck)
    for x := 2; x <= n; x++ {
        // x需要是n的约数
        if n % x == 0 {
            ok := true
            // 每个数字出现的次数也需要是x的倍数
            for _, v := range value {
                if v % x != 0 {
                    ok = false
                    break
                }
            }
            if ok { return true }
        }
    }
    return false
}
```

第二种：统计每个元素出现的次数C，能够分组的前提就是所有C的最大公约数。

```go
// data 2020/03/27
func hasGroupsSizeX(deck []int) bool {
    count := make(map[int]int)
    // 统计每个数出现的次数
    for _, v := range deck {
        if _, ok := count[v]; ok {
            count[v]++
        } else {
            count[v] = 1
        }
    }
    value := make([]int, 0)
    // 将出现的次数放入数组
    for _, c := range count {
        if c > 0 { value = append(value, c) }
    }
    g := -1
    for _, v := range value {
        if g == -1 {
            g = v
        } else {
            g = gcd(g, v)
        }
    }
    return g >= 2
}

func gcd(x, y int) int {
    if x == 0 { return y }
    return gcd(y%x, x)
}
```

