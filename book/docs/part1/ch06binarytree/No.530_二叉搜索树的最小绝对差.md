## 530 二叉搜索树的最小绝对差-中等

题目：

给你一个二叉搜索树的根节点 `root` ，返回 **树中任意两不同节点值之间的最小差值** 。

差值是一个正数，其数值等于两值之差的绝对值。



分析：

递归中序遍历，求相邻元素的最小值。

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
func getMinimumDifference(root *TreeNode) int {
    res := 1 << 32 - 1
    preVal := -1

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)

        if preVal != -1 && root.Val - preVal < res {
            res = root.Val - preVal
        }
        preVal = root.Val
        dfs(root.Right)
    }

    dfs(root)

    return res
}
```
