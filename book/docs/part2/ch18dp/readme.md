## Dynamic Programming [动态编程]

[TOC]

动态编程是一种算法范例，其思想是通过将其分解为子问题来解决给定的复杂问题，并存储子问题的结果，以避免再次计算相同的结果。 那么，如何使用动态规划来解决问题？答案是有的。如果一个问题满足以下两个属性之一，那么表明该问题可以使用动态编程来解决。

* 1）重叠子问题
* 2）最优子结构

### 重叠子问题

类似分治思想，动态规划可以将解决方案与子问题结合。动态规划主要适用于一次又一次地需要子问题的答案来进行求解的问题。因此，在动态规划中将子问题的答案存在表中，以避免重复计算。

例如，二分查找就没有共同的子问题，无法使用动态规划求解。但是，求解斐波那契数列就可以使用动态规划。

一个简单的计算斐波那契数列的递归函数如下：

```cpp
int fib(int n) {
  if(n <= 1)
    return n;
  return fib(n - 1) + fin(n - 2);
}
```

假设求取fib(5)。

![斐波那契数列](http://on64c9tla.bkt.clouddn.com/Algorithm/fib_5.png)

我们可以看到，fib(3)被调用了两次，fib(2)、fib(1)和fib(0)调用的次数就更多了。为了避免重复计算，可以将已经计算出来的值存储到数组中，这样下次计算直接调用数组中存储的值即可。这就是重叠子问题。有两种方式可以实现：

1）记忆（自上而下）

问题的记忆式程序与递归版本相似，在计算解决方案之前，它会对查询表进行小修改。 我们初始化一个所有初始值为NIL的查找数组。 每当我们需要解决一个子问题时，我们首先查看查找表。 如果预先计算的值在那里，那么我们返回该值，否则我们计算该值并将结果放在查找表中，以便稍后重新使用。

```cpp
// lookup[n]初始化为-1
int fib(int n) {
  if(lookup[n] == -1) {
    if(n <= 1)
      lookup[n] = 1;
    else
      lookup[n] = fib(n - 1) + fib(n - 2);
  }
  return lookup[n];
}
```
2) 制表（自下而上）

制表（从上到下）：给定问题的列表程序以自下而上的方式构建表，并从表返回最后一个条目。 例如，对于相同的斐波那契数，我们首先计算fib（0），然后计算fib（1）然后fib（2）然后fib（3）等等。 从字面上看，我们正在建立自下而上的子问题的解决方案。

```cpp
int fib(int n) {
  int f[n+1];
  f[0] = 1, f[1] = 1;
  for(int i = 2; i <= n; i++)
    f[n] = f[n - 1] + f[n - 2];
  return f[n];
}
```

### 最优子结构

**如果一个问题的最优解可以通过其子问题的最优解获得，那么称这个问题具有最优子结构属性**。例如，最短路径问题就具有最优子结构属性：

假设节点 x 位于起始节点 u 与目的节点 v 之间的最短路径上，那么u-v之间的最短路径就包含了u-x间的最短路径和x-v间的最短路径。

所有的两点之间的最短路径算法，像 Floyd–Warshall 和 Bellman–Ford (贝尔曼-福特算法) 都是典型的动态规划。

相反，最长路径问题就没有最优子结构。这里的最长路径是指两点之间最长的简单路径（不包含环）。

![最长路径问题](http://on64c9tla.bkt.clouddn.com/Algorithm/longestpath.png)

如图所示，q-t之间的最长路径为q->r->t或q->s->t。但是q-s之间的最长路径确实q->r->t->s。

### 基本思想和解题步骤

动态规划解决的关键是状态定义，状态的转移，初始化和边界条件。

**状态定义**：

**状态的转移**：

**初始化**：

**边界条件：**

**边界条件**：

**边界条件**：

### 相关题目

#### 62 [不同路径](https://leetcode-cn.com/problems/unique-paths/)【中等】

题目要求：给定一个二维数组m x n，返回机器人从二维数组的左上角->右下角的所有不同路径。

算法1：排列组合 res = f(m+n-2)/f(m-1)/f(n-1)  f为阶乘

算法2：该问题具有**重复子问题**和**最优子结构**，问题的解可以由子问题的解构成，所以可以用动态规划解决

```go
// date 2020/01/11
/* 算法2：递归方程为dp[i][j] = dp[i-1][j] + dp[i][j-1]
   边界条件：当i == 0 || j == 0; dp[i][j] = 1
*/
func uniquePaths(m, n int) int {
  dp := make([][]int, m)
  for i := 0; i < m; i++ {
    dp[i] = make([]int, n)
  }
  for i := 0; i < n; i++ {dp[0][i] = 1}  // 第0行，只能一直向右走，所以均为1
  for i := 0; i < m; i++ {dp[i][0] = 1}  // 第0列，只能一直向下走，所以均为1
  for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
      dp[i][j] = dp[i-1][j] + dp[i][j-1]
    }
  }
  return dp[m-1][n-1]
}
```

#### 64 [最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)【中等】

题目要求：https://leetcode-cn.com/problems/minimum-path-sum/

给定一个包含非负整数的 `m x n` 网格 `grid` ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

**说明：**每次只能向下或者向右移动一步。

思路分析：自底向上递推

递推公式：`opt[i, j] = grid[i][j] + min(opt[i+1][j], opt[i][j+1])`，注意边界情况

```go
// date 2020/03/08
// 时间复杂度O(M*N),空间复杂度O(1)
func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 { return 0 }
    n := len(grid[0])
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            // 如果存在右节点和下节点，选择其中最小的那个
            if i + 1 < m && j + 1 < n {
                grid[i][j] += min(grid[i+1][j], grid[i][j+1])
            // 如果只有右节点，则选择右节点
            } else if j + 1 < n {
                grid[i][j] += grid[i][j+1]
            // 如果只有下节点，则选择下节点
            } else if i + 1 < m {
                grid[i][j] += grid[i+1][j]
            }
        }
    }
    return grid[0][0]
}

func min(x, y int) int {
  if x < y { return x }
  return y
}
```

#### 70 [ 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)【简单】

题目要求：https://leetcode-cn.com/problems/climbing-stairs/

*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。*

*每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？*

*注意：给定 n 是一个正整数。*

思路分析：这个问题和斐波那契数列很像，可以有三种解法，一个比一个更优。

递推公式等于c[n] = c[n-1] + c[n-2]，基本情况n=1, c[n]=1;n=2,c[n]=2

算法一：直接递推

```go
// date 2020/03/08
// 算法一：自顶向下直接递推，超时。
// 时间复杂度O(2^N),空间复杂度O(1)
func climbStairs(n int) int {
  if n < 3 { return n}
  return climbStairs(n-1) + climbStairs(n-2)
}
// 算法二：自底向上递推，并且记录中间结果。
// 时间复杂度O(N),空间复杂度O(N)
func climbStairs(n int) int {
  if n < 3 { return n }
  opt := make([]int, n+1)
  opt[1], opt[2] = 1, 2
  for i := 3; i <= n; i++ {
    opt[i] = opt[i-1] + opt[i-2]
  }
  return opt[n]
}
// 算法三：自底向上的递推，只保留有用的结果
// 时间复杂度O(N),空间复杂度O(1)
func climbStairs(n int) int {
  if n < 3 { return n }
  first, second := 1, 2
  res := 0
  for i := 3; i <= n; i++ {
    res = first + second
    first, second = second, res
  }
  return res
}
```

#### 72 编辑距离

题目要求：https://leetcode-cn.com/problems/edit-distance/

思路分析：

```go
// date 2020/03/16
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    // dp[i][j] 表示将 word[0...i] 变成 word2[0...j] 所需要的最小编辑数
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
        dp[i][0] = i  // j = 0, word2 is empty, delete all word1 elem
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j  // i = 0, word1 is empty, insert all elem to word1
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
                // dp[i-1][j] delete 1 elem in word1
                // dp[i][j-1] insert 1 elem to word2 等价于从 word1 删除 1 个
                // dp[i-1][j-1] replace 1 elem of word1
            }
        }
    }

    return dp[m][n]
}

// word1[0...i-1] -> word2[0....j] = dp[i-1][j]
// word1[0.....i] -> word2[0....j] = dp[i][j]
// 在 dp[i-1][j] 已知的情况下，word1 中 删除 word1[i] 即可
// dp[i][j] = dp[i-1][j], delete 1 elem of word1

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```



#### 121 [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)【简单】

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

思路分析：只能交易一次。

遍历一遍，记录之前的最小值。如果当前值大于当前最小值，则计算利润；如果当前值小于当前最小值，则更新当前最小值，具体算法如下

```go
// date 2020/03/14
// 时间复杂度O(N),空间复杂度O(1)
// 非动态规划题解
func maxProfit(prices []int) int {
  if 0 == len(prices) { return 0 }
  res, cur_min := 0, prices[0]
  for _, v := range prices {
    if v > cur_min && res < v - cur_min { res = v - cur_min }
    if v < cur_min { cur_min = v }
  }
  return res
}
```

#### 122 [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)【简单】

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

思路分析：可以买卖无数次

遍历一遍，只要后面的值大于前面的值，即算作利润。

```go
// date 2020/03/14
// 时间复杂度O(N),空间复杂度O(1)
func maxProfit(prices []int) int {
  profit := 0
  for i := 0; i < len(prices)-1; i++ {
    if prices[i] < prices[i+1] {
      profit =+ prices[i+1] - prices[i]
    }
  }
  return profit
}
```

#### 123 [买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)【困难】

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

思路分析：最多可以交易两笔

遍历一遍，分别从数组[0...i]和[i+1...N]中获取最大利润

```go
// date 2020/03/14
// 时间复杂度O(N^2),空间复杂度O(1)
func maxProfit(prices []int) int {
  res, v1, v2 := 0, 0, 0
  for i := 0; i < len(prices); i++ {
    v1 = maxProfitOnce(prices[:i])
    v2 = maxProfitOnce(prices[i+1:])
    if v1 + v2 > res { res = v1 + v2 }
  }
  return res
}

func maxProfitOnce(prices []int) int {
  if 0 == len(prices) { return 0 }
  res, cur_min := 0, prices[0]
  for _, v := range prices {
    if v > cur_min && res < v - cur_min { res = v - cur_min }
    if v < cur_min { cur_min = v }
  }
  return res
}
// 官方题解
// 思路2
func maxProfit(prices []int) int {
    if 0 == len(prices) { return 0 }
    // 代表当前所拥有的钱数
    // buy1 第一次买入, sell1 第一次卖出, buy2 第二次买入, sell2 第二次卖出
    buy1, sell1 := -prices[0], 0
    buy2, sell2 := -prices[0], 0
    for i := 1; i < len(prices); i++ {
        buy1 = max(buy1, -prices[i]) // 保持不变，或者买入
        sell1 = max(sell1, buy1+prices[i]) // 保持不变，或者卖出
        buy2 = max(buy2, sell1-prices[i]) // 保持不变，或者在第一次卖出的情况下买入
        sell2 = max(sell2, buy2+prices[i]) // 保持不变，或者在第二次买入的情况卖出
    }
    return sell2
}

func max(x, y int) int {
    if x > y { return x }
    return y
}
```

#### 188 买卖股票的最佳时机IV

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

思路分析：

#### 322 [零钱兑换](https://leetcode-cn.com/problems/coin-change/)【中等】

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。

题目分析：

问题关键是如果用最少个数的硬币凑成目标值，所以该问题具有最优子结构，设目标值值target，目标值对应的最优解为opt[target]，则其递推公式为`opt[target] = min(opt[target], 1+opt[target-coin])`

同时，需要边界条件，当无法找到解时，需要返回-1。

```golang
// date 2021/03/16
func coinChange(coins []int, amount int) int {
    // opt[i]表示金额i的最优解，也就是凑成i金额的最少硬币个数
    opt := make([]int, amount+1)
    for i := 0; i <= amount; i++ { opt[i] = amount+1 }
    // 金额0的最优解是0
    opt[0] = 0
    for target := 1; target <= amount; target++ {
        for _, coin := range coins {
            if target < coin { continue }
            opt[target] = min(opt[target], 1+opt[target-coin])
        }
    }
    // 判断是否有解
    if opt[amount] == amount+1 { return -1 }
    return opt[amount]
}

func min(x, y int) int {
    if x < y { return x }
    return y
}
```



#### 264 丑树



题目要求：https://leetcode-cn.com/problems/ugly-number-ii/

算法分析：

```go
// date 2020/03/01
// 动态规划
func nthUglyNumber(n int) int {
    if n < 1 { return 1 }
    dp := make([]int, n)
    dp[0] = 1
    p1, p2, p3 := 0, 0, 0
    for i := 1; i < n; i++ {
        dp[i] = min(2 * dp[p1], min(3 * dp[p2], 5 * dp[p3]))
        if dp[i] == 2 * dp[p1] { p1++ }
        if dp[i] == 3 * dp[p2] { p2++ }
        if dp[i] == 5 * dp[p3] { p3++ }
    }
    return dp[n-1]
}

func min(x, y int) int {
    if x < y { return x }
    return y
}
```

#### 300 最长上升子序列

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence

```
算法一：动态规划，
定义dp方程，dp[i]表示包含nums[i]元素的最长上升子序列
dp[i] = max(dp[j]) + 1; 0 <= j < i && nums[j] < nums[i]
具体算法如下：
```

```go
// date 2020/03/14
// 时间复杂度O(N^2),空间复杂度O(N)
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 { return 0 }
    res, cur_max, n := 1, 0, len(nums)
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < len(nums); i++ {
        // 找到[0,i]之间最大的最长上升子序列
        cur_max = 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                cur_max = max(cur_max, dp[j])
            }
        }
        dp[i] = cur_max + 1  // 如果cur_max为0，+1表示只包含nums[i]自己
        // 否则即为[0,i]的结果集
        res = max(res, dp[i])
    }
    return res
}

func max(x, y int) int {
    if x > y { return x }
    return y
}
```

```
算法二：
思路分析：维护一个最长上升子序列数组lis，然后每次更新它，最后数组lis的长度就是结果。
更新有序数组可以用二分查找，时间复杂度为O(logN)。
需要知道的是最后的lis结果并不一定是真正的结果，但其长度是所要的结果。
```

```go
// date 2020/03/16
func lengthOfLIS(nums []int) int {
  if len(nums) == 0 { return 0}
  n := len(nums)
  lis := make([]int, 0, n)
  lis = append(lis, nums[0])
  for i := 1; i < n; i++ {
    // 在lis数组找到第一个比nums[i]大的元素，并替换它，否则进行追加
    m := len(lis)
    // nums[i] > lis[m-1]
    // nums[i] 直接参与最长上升序列
    if nums[i] > lis[m-1] {
      lis = append(lis, nums[i])
      continue
    }
    // 替换的必要性
    // 1. 保证 nums[i] 有机会参与最长上升序列,为后面更大的值剔除障碍
    // 2. 因为不要求连续性，lis 中的大值没有意义，替换相对于变相删除
    for j := 0; j < m; j++ {
      if nums[i] <= lis[j] {
        lis[j] = nums[i]
        break
      }
    }
  }
  return len(lis)
}
```

#### 面试题17.16 按摩师

题目要求：https://leetcode-cn.com/problems/the-masseuse-lcci/

思路分析：动态规划，动态递推方程`opt[i] = max(opt[i-1], num[i]+max(opt[0...i-2]))`

```go
// date 2020/03/24
func massage(nums []int) int {
    if len(nums) == 0 { return 0 }
    if len(nums) == 1 { return nums[0] }
    res, cur_res, n := 0, 0, len(nums)
    opt := make([]int, n)
    opt[0], opt[1] = nums[0], nums[1]
    res = max(opt[0], opt[1])
    for i := 2; i < n; i++ {
        // 找到0...i-2中最大值
        cur_res = 0
        for j := i-2; j >= 0; j-- {
            if nums[i] + opt[j] > cur_res {
                cur_res = nums[i] + opt[j]
            }
        }
        // 更新opt[i]
        opt[i] = max(opt[i-1], cur_res)
        if opt[i] > res { res = opt[i] }
    }
    return res
}

func max(x, y int) int {
    if x > y { return x }
    return y
}
```

