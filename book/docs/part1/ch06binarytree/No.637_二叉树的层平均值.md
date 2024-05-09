## 637 二叉树的层平均值-简单

题目：

给定一个非空二叉树的根节点 `root` , 以数组的形式返回每一层节点的平均值。与实际答案相差 `10-5` 以内的答案可以被接受。



分析：

没啥好说的，就是算它。

```go
// date 2023/10/25
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    res := make([]float64, 0, 16)
    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        levelSum := 0
        for i := 0; i < n; i++ {
            cur := queue[i]

            levelSum += cur.Val

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        res = append(res, float64(levelSum) / float64(n))

        queue = queue[n:]
    }

    return res
}
```

