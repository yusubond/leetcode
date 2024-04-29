## 236 二叉树的最近公共祖先-中等

题目：

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”



分析：

这里主要是把思路转变一下，通过寻找 p, q 在二叉树的位置来判断。

如果 p, q 分别位于二叉树的左右子树中，那么 root 就是最近的公共祖先。

否则，在左右子树中分别递归查找。

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
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil || root == p || root == q {
         return root
     }
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)
     if left != nil && right != nil {
         return root
     }
     if left == nil {
         return right
     }
     return left
}
```

