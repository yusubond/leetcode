## 129 求从根节点到叶子节点数字之和-中等

给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。



分析：

自顶向下递归，并且携带已知的值；当遇到叶子节点时计算到结果中。

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
func sumNumbers(root *TreeNode) int {
    var ans int
    var dfs func(root *TreeNode, has int)
    dfs = func(root *TreeNode, has int) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            ans += has + root.Val
            return 
        }
        has = (has + root.Val) * 10
        if root.Left != nil {
            dfs(root.Left, has)
        }
        if root.Right != nil {
            dfs(root.Right, has)
        }
    }

    dfs(root, 0)

    return ans
}
```

