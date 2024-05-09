## 965 单值二叉树-中等

题目：

如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true，否则返回 false。


分析：

递归比较即可。

```go
// date 2023/10/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return false
    }


    var isSame func(root *TreeNode, target int) bool
    isSame = func(root *TreeNode, target int) bool {
        if root == nil {
            return true
        }
        return root.Val == target && isSame(root.Left, target) && isSame(root.Right, target)
    }

    return isSame(root, root.Val)
}
```
