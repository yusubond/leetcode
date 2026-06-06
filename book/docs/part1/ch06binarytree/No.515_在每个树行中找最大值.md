## 515 在每个树行中找最大值-中等

题目：

给定一棵二叉树的根节点 `root` ，请找出该二叉树中每一层的最大值。



分析：

层序遍历，找每一层最大值

```go
// date 2026/06/06
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    res := make([]int, 0, 8)
    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0, 4)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        maxVal := queue[0].Val

        for i := 0; i < n; i++ {
            cur := queue[i]
            if cur.Val > maxVal {
                maxVal = cur.Val
            }

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        res = append(res, maxVal)

        queue = queue[n:]
    }

    return res
}
```

