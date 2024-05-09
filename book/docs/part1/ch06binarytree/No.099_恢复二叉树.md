##  99 恢复二叉树-中等

题目：

给你二叉搜索树的根节点 `root` ，该树中的 **恰好** 两个节点的值被错误地交换。*请在不改变其结构的情况下，恢复这棵树* 。



分析：

二叉树的中序遍历肯定为递增序列，所以迭代的方式遍历，根据当前节点与前继节点的大小关系来判断是否被交换过，并记录。

在递增序列中交换两个值，那么第一次打破递增关系时，一定是前驱节点大于当前节点，前驱节点为被换过的。

第二次打破递增关系时，同样前驱节点大于当前节点，但是当前节点是被换过的。

```sh
原始序列：1, 2, 3, 4, 5, 6, 7
交换后的序列：1, 2, 6, 4, 5, 3, 7
第一次为[6,4], 6为置换过的
第二次为[5,3], 3为置换过的
```

下面看具体算法：

```go
// date 2023/10/22
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    stack := make([]*TreeNode, 0, 16)
    // x, y 记录被交换的两个节点
    // pre 记录前驱节点，通过判断当前节点与前驱节点的关系
    // 找到两个被错误交换的节点，然后进行交换
    var x, y, pre *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pre != nil && cur.Val < pre.Val {
            y = cur
            if x == nil {
                x = pre
            } else {
                break
            }
        }
        pre = cur
        root = cur.Right
    }
    x.Val, y.Val = y.Val, x.Val
}
```

