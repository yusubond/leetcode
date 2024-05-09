## 938 二叉搜索树的范围和-简单

题目：

给定二叉搜索树的根结点 `root`，返回值位于范围 *`[low, high]`* 之间的所有结点的值的和。



分析：

算法1：深度优先搜索

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
func rangeSumBST(root *TreeNode, low int, high int) int {
    var sum int

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        if low <= root.Val && root.Val <= high {
            sum += root.Val
        }
        dfs(root.Left)
        dfs(root.Right)
    }

    dfs(root)

    return sum
}
```

算法2：

递归

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
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    var sum int
    sum += rangeSumBST(root.Left, low, high)
    sum += rangeSumBST(root.Right, low, high)
    if root.Val <= high && root.Val >= low {
        sum += root.Val
    }
    return sum
}
```
