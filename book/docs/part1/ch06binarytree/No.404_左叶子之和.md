## 404 左叶子之和-简单

题目：

给定二叉树的根节点 `root` ，返回所有左叶子之和。



分析：先判断是叶子节点，再判断其是前驱节点的左子树

```go
// date 2023/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	var bfs func(root, pre *TreeNode)
	var res int

	bfs = func(root, pre *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && pre != nil && pre.Left == root {
			res += root.Val
		}

		bfs(root.Left, root)
		bfs(root.Right, root)
	}

	bfs(root, nil)

	return res
}
```
