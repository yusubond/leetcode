## 783 二叉搜索树节点最小距离-简单

题目：

给你一个二叉搜索树的根节点 `root` ，返回 **树中任意两不同节点值之间的最小差值** 。

差值是一个正数，其数值等于两值之差的绝对值。



分析：

中序遍历，相邻的两个差值最小。

```go
// date 2023/10/27
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    var ans, pre int
    ans = math.MaxInt64 - 1
    pre = -1
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if pre != -1 && root.Val - pre < ans {
            ans = root.Val - pre
        }
        pre = root.Val
        dfs(root.Right)
    }

    dfs(root)

    return ans
}
```

迭代版

```go
// date 2023/10/27
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    var ans, pre int
    ans = math.MaxInt64 - 1
    pre = -1

    stack := make([]*TreeNode, 0, 16)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) != 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if pre != -1 && root.Val - pre < ans {
                ans = root.Val - pre
            }
            // update pre
            pre = root.Val
            root = root.Right
        }
    }

    return ans
}
```

