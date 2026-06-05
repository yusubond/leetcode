## 623 在二叉树中增加一行-中等

题目：

给定一个二叉树的根 `root` 和两个整数 `val` 和 `depth` ，在给定的深度 `depth` 处添加一个值为 `val` 的节点行。

注意，根节点 `root` 位于深度 `1` 。

加法规则如下:

- 给定整数 `depth`，对于深度为 `depth - 1` 的每个非空树节点 `cur` ，创建两个值为 `val` 的树节点作为 `cur` 的左子树根和右子树根。
- `cur` 原来的左子树应该是新的左子树根的左子树。
- `cur` 原来的右子树应该是新的右子树根的右子树。
- 如果 `depth == 1 `意味着 `depth - 1` 根本没有深度，那么创建一个树节点，值 `val `作为整个原始树的新根，而原始树就是新根的左子树。

 

分析：

层序遍历。找到目标层的上一次，然后插入。

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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        r2 := &TreeNode{Val: val}
        r2.Left = root
        return r2
    }
    curs := make([]*TreeNode, 0, 16)
    curs = append(curs, root)
    depth--
    for depth > 1 {
        n := len(curs)
        for i := 0; i < n; i++ {
            cur := curs[i]
            if cur.Left != nil {
                curs = append(curs, cur.Left)
            }
            if cur.Right != nil {
                curs = append(curs, cur.Right)
            }
        }
        curs = curs[n:]
        depth--
    }
    if len(curs) != 0 {
        n := len(curs)
        for i := 0; i < n; i++ {
            cur := curs[i]
            r1, r2 := &TreeNode{Val: val}, &TreeNode{Val: val}
            // save origin left and right
            rl, rr := cur.Left, cur.Right
            cur.Left = r1
            cur.Right = r2
            r1.Left = rl
            r2.Right = rr
        }
    }
    return root
}
```

