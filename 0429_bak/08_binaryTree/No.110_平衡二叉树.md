## 110 平衡二叉树-简单

题目：

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。



分析：

因为每个节点都需要保证平衡，所以采用自底向上的深度优先搜索，

因为需要每个节点都保证是平衡，所以可以采用自底向上的递归。

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
func isBalanced(root *TreeNode) bool {
    res := true
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int {
        if root == nil {
            return depth
        }
        l := dfs(root.Left, depth)
        r := dfs(root.Right, depth)
        if abs(l, r) > 1 {
            res = false
        }
        if l > r {
            return l+1
        }
        return r+1
    }

    dfs(root, 0)

    return res
}

func abs(x, y int) int{
    if x > y {
        return x-y
    }
    return y-x
}
```

