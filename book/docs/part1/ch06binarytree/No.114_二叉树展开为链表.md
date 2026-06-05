## 114 二叉树展开为链表-中等

题目：

给你二叉树的根结点 `root` ，请你将它展开为一个单链表：

- 展开后的单链表应该同样使用 `TreeNode` ，其中 `right` 子指针指向链表中下一个结点，而左子指针始终为 `null` 。
- 展开后的单链表应该与二叉树 [**先序遍历**](https://baike.baidu.com/item/先序遍历/6442839?fr=aladdin) 顺序相同。



分析：

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
func flatten(root *TreeNode)  {
    cur := root

    for cur != nil {
        if cur.Left != nil {
            next := cur.Left
            pre := next
            for pre.Right != nil {
                pre = pre.Right
            }
            pre.Right = cur.Right

            cur.Left = nil
            cur.Right = next
        }

        cur = cur.Right
    }
}
```

