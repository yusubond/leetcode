## 572 另一棵树的子树-中等

题目：

给你两棵二叉树 `root` 和 `subRoot` 。检验 `root` 中是否包含和 `subRoot` 具有相同结构和节点值的子树。如果存在，返回 `true` ；否则，返回 `false` 。

二叉树 `tree` 的一棵子树包括 `tree` 的某个节点和这个节点的所有后代节点。`tree` 也可以看做它自身的一棵子树。



分析：【推荐该算法】

这道题跟相同的树很类似，判断是不是子树，就等于判断树的子树是否与另一个树相同。

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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    return checkIsSame(root, subRoot) ||  isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkIsSame(r1, r2 *TreeNode) bool {
    if r1 == nil && r2 == nil {
        return true
    }
    if r1 == nil || r2 == nil {
        return false
    }
    if r1.Val == r2.Val {
        return checkIsSame(r1.Left, r2.Left) && checkIsSame(r1.Right, r2.Right)
    }
    return false
}
```



算法2：

这个算法是通过将二叉树的延展成先序序列，并且延展的时候把空节点也当成字符充填进去。

那么，是不是子树，就等价于判断一个字符串中是否存在子串。