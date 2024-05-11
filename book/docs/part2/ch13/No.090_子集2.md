## 90 子集2-中等

题目：

给你一个整数数组 `nums` ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。返回的解集中，子集可以按 **任意顺序** 排列。



> **示例 1：**
>
> ```
> 输入：nums = [1,2,2]
> 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0]
> 输出：[[],[0]]
> ```



分析：

这道题跟 78 题类似，两者的区别是：

78 题，给定的数组中元素不重复，返回所有的子集

这道题，给定的数组中元素有重复，返回所有子集，要求子集去重。

核心思路还是回溯，增加两个辅助条件：

1. 回溯前，对数组进行排序，方便回溯的时候跳过重复元素
2. 回溯过程中，剪枝优化。剪枝的关键是初次遍历不剪枝，下一次的时候，如果元素跟前一个元素相同，直接跳过回溯。



```go
// date 2023/12/26
func subsetsWithDup(nums []int) [][]int {
    res := make([][]int, 0, 16)

    var backtrack func(start int, temp []int)
    backtrack = func(start int, temp []int) {
        one := make([]int, len(temp))
        copy(one, temp)
        res = append(res, one)

        for i := start; i < len(nums); i++ {
            // 去重的关键，本次不去重，下一次去重
            for i < len(nums) && i > start && nums[i] == nums[i-1] {
                i++
                continue
            }
            // 注意判断，去重后是否越界
            if i >= len(nums) {
                break
            }


            temp = append(temp, nums[i])
            backtrack(i+1, temp)
            temp = temp[:len(temp)-1]
        }
    }

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    backtrack(0, []int{})

    return res
}
```

