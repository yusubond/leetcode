## 112 路径总和-简单

题目：

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。





分析：

算法1：递归

这种递归是**自顶向下**的方式，通过结合当前节点值，向下一层级传递新的目标值，以判断到达叶子节点时是否满足条件。

```go
// date 2022/10/19
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    targetSum -= root.Val
    if root.Left == nil && root.Right == nil {
        return targetSum == 0
    }
    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
```



算法2：【推荐该算法】

迭代，类似层序遍历，当遍历到叶子节点时，判断值是否为零

```go
// date 2023/10/18
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    queue := make([]*TreeNode, 0, 16)
    csum := make([]int, 0, 16)
    queue = append(queue, root)
    csum = append(csum, targetSum-root.Val)
    n := len(queue)
    for len(queue) != 0 {
        n = len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            // 找到叶子节点
            if cur.Left == nil && cur.Right == nil && csum[i] == 0 {
                return true
            }
            if cur.Left != nil {
                queue = append(queue, cur.Left)
                csum = append(csum, csum[i] - cur.Left.Val)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
                csum = append(csum, csum[i] - cur.Right.Val)
            }
        }
        queue = queue[n:]
        csum = csum[n:]
    }
    return false
}
```

