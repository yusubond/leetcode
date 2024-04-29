## 543 二叉树的直径

题目：

给你一棵二叉树的根节点，返回该树的 **直径** 。

二叉树的 **直径** 是指树中任意两个节点之间最长路径的 **长度** 。这条路径可能经过也可能不经过根节点 `root` 。

两节点之间路径的 **长度** 由它们之间边数表示。



分析：

所谓直径就是每个节点的左右子树深度之和的全局最大值。

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
func diameterOfBinaryTree(root *TreeNode) int {
    var ans int

    var depth func(root *TreeNode) int
    depth = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l, r := depth(root.Left), depth(root.Right)
        // update ans
        if l + r > ans {
            ans = l + r
        }

        if l > r {
            return l+1
        }
        return r+1
    }

    depth(root)

    return ans
}
```

