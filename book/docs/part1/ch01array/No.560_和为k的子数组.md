## 560 和为k的子数组-中等

题目：

给你一个整数数组 `nums` 和一个整数 `k` ，请你统计并返回 *该数组中和为 `k` 的子数组的个数* 。

子数组是数组中元素的连续非空序列。



> **示例 1：**
>
> ```
> 输入：nums = [1,1,1], k = 2
> 输出：2
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1,2,3], k = 3
> 输出：2
> ```



**解题思路**

子数组是指连续的非空序列。

- 暴力求解，详见解法1。既然是连续的子数组，那么就可以两层遍历，第一层遍历每个元素，第二层从该元素开始往前遍历求和，只要等于目标指，结果加一。目标值可能为零，元素也可能为零，所以第二层遍历不能够匹配到目标值后就跳出，所以遍历所有情况。
- 前缀和，详见解法2。子序和为k，也可以看成是前缀和之间的差为 k，所以我们可以这样做：声明变量 preSum，然后从前往后遍历，计算前缀和，并保存在 map 中，从 map 中查找 pre-k的个数。



```go
// date 2024/01/16
// 解法1
// 暴力求解
func subarraySum(nums []int, k int) int {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++{
        sum := 0
        for j := i; j >= 0; j-- {
            sum += nums[j]
            if sum == k {
                ans++
            }
        }
    }
    return ans
}
// 解法2
// 前缀和 + 哈希查找
func subarraySum(nums []int, k int) int {
    n := len(nums)
    ans := 0
    preSum := 0
    m := make(map[int]int)
    m[0] = 1
    for i := 0; i < n; i++{
        preSum += nums[i]
        
        if ct, ok := m[preSum-k]; ok {
            ans += ct
        }

        m[preSum]++
    }
    return ans
}
```

