## 213 打家劫舍2-中等

题目：

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 **围成一圈** ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，**如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警** 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 **在不触动警报装置的情况下** ，今晚能够偷窃到的最高金额。



> **示例 1：**
>
> ```
> 输入：nums = [2,3,2]
> 输出：3
> 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1,2,3,1]
> 输出：4
> 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
>      偷窃到的最高金额 = 1 + 3 = 4 。
> ```
>
> **示例 3：**
>
> ```
> 输入：nums = [1,2,3]
> 输出：3
> ```



```go
// date 2024/03/03
func rob(nums []int) int {

    var robRange func(nums []int) int
    robRange = func(nums []int) int {
        y2, y1 := nums[0], max(nums[1], nums[0])
        for _, v := range nums[2:] {
            y0 := max(y1, y2+v)
            y2, y1 = y1, y0
        }
        return y1
    }

    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return max(nums[0], nums[1])
    }

    return max(robRange(nums[:n-1]), robRange(nums[1:]))
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

