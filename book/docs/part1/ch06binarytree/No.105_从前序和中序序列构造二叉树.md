## 105 从前序与中序遍历序列构造二叉树-中等

题目：

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。



分析：

前序序列是按照「根左右」的顺序遍历，而中序序列是按照「左根右」的顺序遍历，所以根节点肯定是前序序列中的第一个元素，然后再中序序列中找到该根节点的位置，将其分为左右子树，递归构建。

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    index := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            index = i
            break
        }
    }
    root.Left = buildTree(preorder[1:index+1], inorder[0:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
    return root
}
```

