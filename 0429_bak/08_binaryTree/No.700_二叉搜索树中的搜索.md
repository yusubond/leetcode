## 700 二叉搜索树中的搜索-简单

题目：

给定二叉搜索树（BST）的根节点 `root` 和一个整数值 `val`。

你需要在 BST 中找到节点值等于 `val` 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 `null` 。



分析：

根据二叉搜索树的特性，直接递归搜索即可。

```go
// date 2023/10/30
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    if root.Val < val {
        return searchBST(root.Right, val)
    }
    return searchBST(root.Left, val)
}
```

