## 701 二叉搜索树的插入操作-中等

题目：

给定二叉搜索树（BST）的根节点 `root` 和要插入树中的值 `value` ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 **保证** ，新值和原始二叉搜索树中的任意节点值都不同。

**注意**，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 **任意有效的结果** 。



分析：


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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if root.Val > val {
        if root.Left == nil {
            root.Left = &TreeNode{Val: val}
        } else {
            insertIntoBST(root.Left, val)
        }
    }
    if root.Val < val {
        if root.Right == nil {
            root.Right = &TreeNode{Val: val}
        } else {
            insertIntoBST(root.Right, val)
        }
    }

    return root
}
```
