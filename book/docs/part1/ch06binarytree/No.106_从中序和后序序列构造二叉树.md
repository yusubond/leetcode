## 106从中序与后序遍历序列构造二叉树-中等

题目：

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。



分析：

后序序列中的最后一个元素为根节点，然后再中序序列中找到根节点的位置，依次递归构造左右子树。

```go
// date 2023/10/23
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    n := len(postorder)
    index := 0
    root := &TreeNode{Val: postorder[n-1]}

    for i := 0; i < n; i++ {
        if inorder[i] == postorder[n-1] {
            index = i
            break
        }
    }
    // inorder   左右根
    // postorder 左右根
    root.Left = buildTree(inorder[0:index], postorder[0:index])   // 一样长
    root.Right = buildTree(inorder[index+1:], postorder[index:n-1])  // 一样的起点

    return root
}
```

