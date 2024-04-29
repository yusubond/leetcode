## 101 对称二叉树-简单

题目：

给你一个二叉树的根节点 `root` ， 检查它是否轴对称。

轴对称的意思是节点关于整棵树的根节点 root 对称。



**解题思路**

解法1：子树互为镜像

如果一个树的左右子树关于根节点对称，那么这个树就是轴对称的。

很直观看就是左右子树互为镜像，那么两个树在什么情况下互为镜像呢？

> 如果两个树同时满足两个条件，则互为镜像。
>
> 1.它们的根节点具有相同的值；
>
> 2.一个树的右子树和另一个树的左子树对称；一个树的左子树与另一个树的右子树对称

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
func isSymmetric(root *TreeNode) bool {
    var isMirror func(r1, r2 *TreeNode) bool
    isMirror = func(r1, r2 *TreeNode) bool {
        if r1 == nil && r2 == nil {
            return true
        }
        if r1 == nil || r2 == nil {
            return false
        }
        return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
    }

    return isMirror(root, root)
}
```



解法2：【推荐该算法】

类似层序遍历，将子树的左右结点依次放入队列，通过判断队列中的连续的两个值是否相同来确认是否对称。

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
func isSymmetric(root *TreeNode) bool {
    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root, root)
    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i += 2 {
            r1, r2 := queue[i], queue[i+1]
            if r1 == nil && r2 == nil {
                continue
            }
            if r1 == nil || r2 == nil {
                return false
            }
            if r1.Val != r2.Val {
                return false
            }
            queue = append(queue, r1.Left, r2.Right, r1.Right, r2.Left)
        }
        queue = queue[n:]
    }
    return true
}
```
