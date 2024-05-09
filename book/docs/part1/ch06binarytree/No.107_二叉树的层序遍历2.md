## 107 二叉树的层序遍历2-中等

题目：

给你二叉树的根节点 `root` ，返回其节点值 **自底向上的层序遍历** 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）



分析：

这里先正常层序遍历，然后再将结果集反转一下。【推荐该算法】

想了下，暂时没有比这更好的算法，因为向数组的头部插入数据，涉及拷贝，所以并不能够将每一层的结果一次性放入指定位置。

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
func levelOrderBottom(root *TreeNode) [][]int {
    res := make([][]int, 0, 16)
    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)
    
    for len(queue) != 0 {
        n := len(queue)

        lRes := make([]int, 0, 16)
        for i := 0; i < n; i++ {
            cur := queue[i]
            lRes = append(lRes, cur.Val)

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        res = append(res, lRes)
        queue = queue[n:]
    }

    i, j := 0, len(res)-1
    for i < j {
        res[i], res[j] = res[j], res[i]
        i++
        j--
    }

    return res
}
```

