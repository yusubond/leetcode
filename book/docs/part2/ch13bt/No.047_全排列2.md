## 47 全排列2-中等

题目：

给定一个可包含重复数字的序列 `nums` ，***按任意顺序*** 返回所有不重复的全排列。

> **示例 1：**
>
> ```
> 输入：nums = [1,1,2]
> 输出：
> [[1,1,2],
>  [1,2,1],
>  [2,1,1]]
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1,2,3]
> 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
> ```



**解题思路**

本题也是回溯算法，通过剪枝去重即可，剪枝逻辑是原数组排序后，相同元素的全排列取第一个即可。

```go
// date 2023/11/07
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0, 16)

	var backtrack func([]int, []int)
	// uv unVisited
	backtrack = func(uv []int, path []int) {
		if len(uv) == 0 {
			one := make([]int, len(path))
			copy(one, path)
			res = append(res, one)
			return
		}

		for i, v := range uv {
			// 剪枝去重: 当 当前元素与其前驱元素相同时，取一次结果即可
			if i > 0 && uv[i] == uv[i-1] {
				continue
			}
			// pick up v
			path = append(path, v)
			unUsed := make([]int, 0, len(uv))
			unUsed = append(unUsed, uv[:i]...)
			unUsed = append(unUsed, uv[i+1:]...)
			backtrack(unUsed, path)

			path = path[:len(path)-1]
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backtrack(nums, []int{})

	return res
}
```

