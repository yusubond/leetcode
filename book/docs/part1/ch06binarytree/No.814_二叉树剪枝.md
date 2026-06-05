## 814 二叉树剪枝-中等
题目：

给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
返回移除了所有不包含 1 的子树的原二叉树。
节点 node 的子树为 node 本身加上所有 node 的后代。



分析：

算法1【推荐该算法】

先递归左右子树，然后判断子树是否为空，以及当前节点是否为零，如果是直接剪枝。

```go
// date 2023/10/27
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    if root.Left == nil && root.Right == nil && root.Val == 0 {
        return nil
    }
    return root
}
```





算法2

使用 `containsOne()` 递归判断节点以及节点的左右子树是否包含1，如果不包含，直接剪枝。

```go
// date 2023/10/27
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    if containsOne(root) {
        return root
    }
    return nil
}

func containsOne(root *TreeNode) bool {
    if root == nil {
        return false
    }
    l, r := containsOne(root.Left), containsOne(root.Right)
    if !l {
        root.Left = nil
    }
    if !r {
        root.Right = nil
    }
    return root.Val == 1 || l || r
}
```

