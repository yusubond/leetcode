## 53 最大子数组和-中等

题目：

给你一个整数数组 `nums` ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**子数组** 是数组中的一个连续部分。



**解题思路**

解法1：动态规划。

如果我们用`f(i)`表示以第`i`个数结尾的「连续子数组和」，那么题目就是求：

```sh
max(f(i)), 0<=i<=n-1
```

对于数组每个位置上的元素，我们都只有两个选择，要么加入前一个元素对应的子数组，要么单独成为一段子数组，这取决于`nums[i]`和`f(i-1)+nums[i]`的大小。于是就有下面的动态规划转移方程：

```sh
f(i) = max{nums[i], f(i-1)+nums[i]}
```

很明显，当`f(i-1)`小于零时，`f(i) = nums[i]`

考虑到上面的动态转移方程`f(i)`之和`f(i-1)`相关，我们可用一个变量 pre 来维护对于当前`f(i)`的`f(i-1)`值是多少。具体代码详见解法1。

```go
// date 2023/11/04
// 解法1
func maxSubArray(nums []int) int {
    ans := nums[0]
    // f(i) = max{f(i-1)+nums[i], nums[i]}
    // pre = f(i-1)
    pre := 0
    for _, v := range nums {
        pre = max(pre+v, v)
        ans = max(pre, ans)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```
