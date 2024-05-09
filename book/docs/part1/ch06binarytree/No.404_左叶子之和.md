## 404 左叶子之和-简单

题目：

给定二叉树的根节点 `root` ，返回所有左叶子之和。



分析：

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
    var res int

    var dfs func(root *TreeNode, isLeft bool)
    dfs = func(root *TreeNode, isLeft bool) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil && isLeft {
            res += root.Val
            return
        }
        if root.Left != nil {
            dfs(root.Left, true)
        }
        if root.Right != nil {
            dfs(root.Right, false)
        }
    }

    dfs(root, false)

    return res
}
```

