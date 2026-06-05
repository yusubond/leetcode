## 538 把二叉搜索树转换为累加树-中等

题目：

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。



分析：

对于二叉搜索树其中序遍历序列肯定是递增序列，而要转换为 大于或等于 节点值之和，那么可以采用 中序 的逆序遍历，然后再遍历过程中直接求和。

```go
// date 2023/10/25
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    var sum int
    var deorder func(root *TreeNode)
    deorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        deorder(root.Right)
        sum += root.Val
        root.Val = sum
        deorder(root.Left)
    }

    deorder(root)

    return root
}
```

