## 238 除自身以外数组的乘积-中等

题目：

给你一个整数数组 `nums`，返回 *数组 `answer` ，其中 `answer[i]` 等于 `nums` 中除 `nums[i]` 之外其余各元素的乘积* 。

题目数据 **保证** 数组 `nums`之中任意元素的全部前缀元素和后缀的乘积都在 **32 位** 整数范围内。

请 **不要使用除法，**且在 `O(*n*)` 时间复杂度内完成此题。



> **示例 1:**
>
> ```
> 输入: nums = [1,2,3,4]
> 输出: [24,12,8,6]
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [-1,1,0,-3,3]
> 输出: [0,0,9,0,0]
> ```



**解题思路**

这道题可用前缀乘积和后缀乘积解决，详见解法1。具体为，初始化两个变量 left 和 right，分别表示前缀乘积和后缀乘积，初始值为1。

先从前向后遍历，left 负责记录前缀乘积，并更新 ans；在从后向前遍历，right 记录后缀乘积，并更新 ans。

```go
// date 2024/01/16
// 解法1
// 前缀乘积 * 后缀乘积
func productExceptSelf(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }

    ans := make([]int, n)
    left := 1
    // fill left
    for i := 0; i < n; i++ {
        ans[i] = left 
        left *= nums[i]
    }

    // fill right
    right := 1
    for i := n-1; i >= 0; i-- {
        ans[i] = ans[i] * right
        right *= nums[i]
    }
    return ans
}
```

