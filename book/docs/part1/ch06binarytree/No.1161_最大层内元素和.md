## 1161 最大层内元素和-中等

题目：

给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。

请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。



分析：

这里就是层序遍历，根据每层的结果更新最终的结果即可。

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
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)
    n := len(queue)
    ls := 0 // level sum
    lsMax := math.Inf(-1)
    level, res := 1, 1
    for len(queue) != 0 {
        n = len(queue)
        ls = 0
        for i := 0; i < n; i++ {
            cur := queue[i]
            ls += cur.Val
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        if float64(ls) > lsMax {
            lsMax = float64(ls)
            res = level
        }
        queue = queue[n:]
        level++
    }
    return res
}
```

