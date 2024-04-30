## 15 三数之和-中等

题目：

给你一个整数数组 `nums` ，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k` ，同时还满足 `nums[i] + nums[j] + nums[k] == 0` 。请

你返回所有和为 `0` 且不重复的三元组。

**注意：**答案中不可以包含重复的三元组。



> **示例 1：**
>
> ```
> 输入：nums = [-1,0,1,2,-1,-4]
> 输出：[[-1,-1,2],[-1,0,1]]
> 解释：
> nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
> nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
> nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
> 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
> 注意，输出的顺序和三元组的顺序并不重要。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0,1,1]
> 输出：[]
> 解释：唯一可能的三元组和不为 0 。
> ```
>
> **示例 3：**
>
> ```
> 输入：nums = [0,0,0]
> 输出：[[0,0,0]]
> 解释：唯一可能的三元组和为 0 。
> ```



**解题思路**

算法：时间复杂度O(n2)

1-对数组进行升序排序，然后逐一遍历，固定当前元素nums[i]，使用左右指针遍历当前元素后面的所有元素，nums[l]和nums[r]，如果三数之和=0，则添加进结果集

2-当nums[i] > 0 结束循环 // 升序排序，三数之和必然大于零

3-当nums[i] == nums[i-1] 表明用重复元素，跳过

4-sum = 0零时，如果nums[l] == nums[l+1] 重复元素，跳过

5-sum = 0, nums[r] == nums[r-1] 重复元素，跳过

```go
// date 2022/09/14
func threeSum(nums []int) [][]int {
    ans := make([][]int, 0, 16)
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    n := len(nums)
    for i := 0; i < n; i++ {
        // case 0
        if nums[i] > 0 {
            break
        }
        // case 1
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left := i+1
        right := n-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else if sum == 0 {
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            }
        }
    }

    return ans
}
```

