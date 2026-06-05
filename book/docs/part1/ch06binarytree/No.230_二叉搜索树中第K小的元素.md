## 230 二叉搜索树中第K小的元素

题目：

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。



分析：

利用二叉搜索树的特点，其左根右（前序遍历）序列为非递减序列。

所以，直接递归遍历，数数即可。

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
func kthSmallest(root *TreeNode, k int) int {
    var res int
    var count int
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        count++
        if count == k {
            res = root.Val
            return
        }
        dfs(root.Right)
    }

    dfs(root)

    return res
}
```
