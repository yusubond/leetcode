## 1008 前序遍历构造二叉搜索树-中等

题目：

给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。

保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。

二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大于 Node.val。

二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。



分析：

根据二叉搜索树特性，从前序序列中找到右子树节点，递归构造。

需要注意的是，要考虑左子树可能为空。

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
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    if len(preorder) == 1 {
        return &TreeNode{Val: preorder[0]}
    }

    root := &TreeNode{Val: preorder[0]}

    // find right index
    idx := 0
    for i, v := range preorder{
        if v > preorder[idx] {
            idx = i
            break
        }
    }
    // left
    if idx == 0 {
        // all val is left
        root.Left = bstFromPreorder(preorder[1:])
    } else {
        root.Left = bstFromPreorder(preorder[1:idx])
    }

    // find the right
    if idx != 0 {
        root.Right = bstFromPreorder(preorder[idx:])
    }

    return root
}
```
