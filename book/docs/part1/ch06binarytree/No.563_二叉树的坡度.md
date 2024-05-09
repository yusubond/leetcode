## 563 二叉树的坡度-中等

题目：

给你一个二叉树的根节点 root ，计算并返回 整个树 的坡度 。

一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。

整个树 的坡度就是其所有节点的坡度之和。



分析：深度优先搜索

递归的思路，求每个节点的左右子树和，并在求和过程中更新坡度值。

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
func findTilt(root *TreeNode) int {
    var ans int
    var sum func(root *TreeNode) int
    sum = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        ls, rs := sum(root.Left), sum(root.Right)
        
        ans += abs(ls, rs)

        return ls + rs + root.Val
    }

    sum(root)

    return ans
}

func abs(x, y int) int {
    if x > y {
        return x-y
    }
    return y-x
}
```

