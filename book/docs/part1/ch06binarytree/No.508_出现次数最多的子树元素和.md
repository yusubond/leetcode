## 508 出现次数最多的子树元素和-中等

题目：

给你一个二叉树的根结点 `root` ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的 **「子树元素和」** 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。



分析：

深度优先搜索，维护子树元素和及其出现的次数 map；然后对 map 过滤。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	tSum := make(map[int]int, 4)
	maxTimes := 0

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := root.Val + dfs(root.Right) + dfs(root.Left)
		tSum[sum]++
		if tSum[sum] > maxTimes {
			maxTimes = tSum[sum]
		}
		return sum
	}

	dfs(root)

	res := make([]int, 0, 4)
	for k, v := range tSum {
		if v == maxTimes {
			res = append(res, k)
		}
	}

	return res
}
```

