## 46 全排列-中等

题目：

给定一个不含重复数字的数组 `nums` ，返回其 *所有可能的全排列* 。你可以 **按任意顺序** 返回答案。

 

> **示例 1：**
>
> ```
> 输入：nums = [1,2,3]
> 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0,1]
> 输出：[[0,1],[1,0]]
> ```
>
> **示例 3：**
>
> ```
> 输入：nums = [1]
> 输出：[[1]]
> ```



分析：

经典的回溯算法，算法逻辑是：

1. 结束条件：可选列表为空
2. 遍历的时候，把已经遍历过的元素，不在添加到选择列表中



```go
// date 2023/12/25
// 可对比下两种写法，写法 1 更直观，推荐 1
// 写法1
func permute(nums []int) [][]int {
	res := make([][]int, 0, 16)

	var backtrack func([]int, []int)
	backtrack = func(unVisited []int, path []int) {
		if len(unVisited) == 0 {
			one := make([]int, len(path))
			copy(one, path)
			res = append(res, one)
			return
		}
		for i, v := range unVisited {
      // pick up v
			path = append(path, v)       // 选择，在原有路径上追加
			unUsed := make([]int, 0, len(unVisited))
			unUsed = append(unUsed, unVisited[:i]...)
			unUsed = append(unUsed, unVisited[i+1:]...)
			backtrack(unUsed, path)
			path = path[:len(path)-1]   // 撤销选择
		}
	}

	backtrack(nums, []int{})

	return res
}

// 写法2
func permute(nums []int) [][]int {
	res := make([][]int, 0, 16)

	var backtrack func([]int, []int)
	backtrack = func(unVisited []int, path []int) {
		if len(unVisited) == 0 {
			res = append(res, path)     // 因为 41 行总是新变量，这里可以直接用
			return
		}
		for i, v := range unVisited {
      vPath := append(path, v)     // 注意这里是 := 意味着选择 v 之后，使用新切片
			unUsed := make([]int, 0, len(unVisited))
			unUsed = append(unUsed, unVisited[:i]...)
			unUsed = append(unUsed, unVisited[i+1:]...)
			backtrack(unUsed, vPath)
			// path = path[:len(path)-1]
		}
	}

	backtrack(nums, []int{})

	return res
}
```



