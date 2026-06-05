## 1302 层数最深叶子节点的和-中等

题目：

给你一棵二叉树的根节点 `root` ，请你返回 **层数最深的叶子节点的和** 。



分析：

就是层序遍历，只存储当前层结果，那么最后一次结果就是答案。

```go
// date 2023/20/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    var res int
    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        res = 0
        for i := 0; i < n; i++ {
            cur := queue[i]
            res += cur.Val
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        queue = queue[n:]
    }

    return res
}
```

