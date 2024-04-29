## 103 二叉树的锯齿形层序遍历-中等

题目：

给你二叉树的根节点 `root` ，返回其节点值的 **锯齿形层序遍历** 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。



分析：

正常层序遍历，每层存储结果的时候，先正向再反向。

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
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0, 16)
    if root == nil {
        return res
    }
    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)
    isLeftFirst := true
    
    for len(queue) != 0 {
        n := len(queue)
        lRes := make([]int, n)
        for i := 0; i < n; i++ {
            cur := queue[i]

            if isLeftFirst {
                lRes[i] = cur.Val
            } else {
                lRes[n-1-i] = cur.Val
            }

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        queue = queue[n:]
        res = append(res, lRes)
        isLeftFirst = !isLeftFirst
    }
    return res
}
```

