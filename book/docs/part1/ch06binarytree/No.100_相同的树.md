## 100 相同的树-简单

题目：

给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。



分析：

直接递归判断就可以。

从递归可以看到，这是一种**自底向上**的方式，需要结合左右子树的情况以及当前节点本身，确定最终结果。

```go
// date 2022/10/10
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

