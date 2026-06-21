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



**解题思路**

这道题是路径探索类的回溯，通过剪枝去重，控制进入回溯的时机。

剪枝的关键是初次遍历不剪枝，下一次的时候，如果元素跟前一个元素相同，直接跳过回溯。

```go
// date 2023/12/26
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, 16)
	n := len(nums)

	var backtrack func(int, []int)
	backtrack = func(start int, path []int) {
		one := make([]int, len(path))
		copy(one, path)
		res = append(res, one)

		for i := start; i < n; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backtrack(0, []int{})

	return res
}
```

