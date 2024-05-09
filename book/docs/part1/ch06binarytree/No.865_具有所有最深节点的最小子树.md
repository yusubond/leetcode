## 865 具有所有最深节点的最小子树

题目：

给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。

返回包含原始树中所有 最深节点 的 最小子树 。

如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。

一个节点的 子树 是该节点加上它的所有后代的集合。



分析：

该问题就是求所有叶子节点的最近公共祖先。

定义函数 find 寻找每个节点的深度和最近公共祖先。如果深度相等，说明左右子树中都有最深的叶子节点，那么当前节点就是所有叶子节点的最近公共祖先。

如果左子树深度大，那么返回左子树结果。

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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    // 问题就是 求所有叶子节点的最近公共祖先
    _, res := find(root)
    return res
}

// 求 子树的深度 和 子树叶子节点的最近公共祖先
func find(root *TreeNode) (int, *TreeNode) {
    if root == nil {
        return 0, nil
    }
    ld, ln := find(root.Left)
    rd, rn := find(root.Right)
    if ld > rd {
        return ld+1, ln
    }
    if ld < rd {
        return rd+1, rn
    }
    return ld+1, root
}
```
