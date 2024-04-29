## 669 修剪二叉搜索树-中等

题目：

给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。


分析：


根据二叉搜索树特征，直接修剪。

```go
// date 2023/11/01
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val < low {
        // 说明左子树完全不需要，直接舍弃
        // 直接修剪右子树
        return trimBST(root.Right, low, high)
    }
    if root.Val > high {
        // 同理，舍弃右子树，直接修剪左子树
        return trimBST(root.Left, low, high)
    }
    root.Left = trimBST(root.Left, low, high)
    root.Right = trimBST(root.Right, low, high)
    return root
}
```
