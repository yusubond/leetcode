## 40 组合总和2-中等

题目：

给定一个候选人编号的集合 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的每个数字在每个组合中只能使用 **一次** 。

**注意：**解集不能包含重复的组合。 



> **示例 1:**
>
> ```
> 输入: candidates = [10,1,2,7,6,1,5], target = 8,
> 输出:
> [
> [1,1,6],
> [1,2,5],
> [1,7],
> [2,6]
> ]
> ```
>
> **示例 2:**
>
> ```
> 输入: candidates = [2,5,2,1,2], target = 5,
> 输出:
> [
> [1,2,2],
> [5]
> ]
> ```



分析：

这道题是 39 题的强化版，其强化条件是 每个数字在组合中只能使用一次。

这意味，如果原数组中的数字不重复，那么回溯递归的时候要往下一个元素去找。

如果原数组的数字有重复，那么重复的数字只能使用有限次。

所以去重的逻辑很关键，在什么时候去重。

```go
// date 2023/12/26
func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0, 16)

    var backtrack func(nums []int, target, idx int, temp []int)
    backtrack = func(nums []int, target, idx int, temp []int) {
        if target <= 0 {
            if target == 0 {
                one := make([]int, len(temp))
                copy(one, temp)
                res = append(res, one)
            }
            return
        }
        for i := idx; i < len(nums); i++ {
          	// 这里还没有进入递归回溯的函数
          	// 所以，也就是说，nums[i] 已经搜索以后
          	// 如果存在重复，则跳过
          	// 即对于重复的3个元素，取第一个元素的回溯结果，
          	// 剩下的两个直接跳过，避免结果集中重复
            if i > idx && nums[i] == nums[i-1] {
                continue
            }
          	// 剪枝
            if nums[i] > target {
                break
            }
            temp = append(temp, nums[i])
          	// 元素只能使用一次，从 i+1 开始搜索
            backtrack(nums, target-nums[i], i+1, temp)
            temp = temp[:len(temp)-1]
        }
    }

    sort.Slice(candidates, func(i, j int) bool {
        return candidates[i] < candidates[j]
    })

    backtrack(candidates, target, 0, []int{})

    return res
}
```

