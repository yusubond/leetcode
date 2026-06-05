## 430 扁平化多级双向链表-中等

题目：

你会得到一个双链表，其中包含的节点有一个下一个指针、一个前一个指针和一个额外的 **子指针** 。这个子指针可能指向一个单独的双向链表，也包含这些特殊的节点。这些子列表可以有一个或多个自己的子列表，以此类推，以生成如下面的示例所示的 **多层数据结构** 。

给定链表的头节点 head ，将链表 **扁平化** ，以便所有节点都出现在单层双链表中。让 `curr` 是一个带有子列表的节点。子列表中的节点应该出现在**扁平化列表**中的 `curr` **之后** 和 `curr.next` **之前** 。

返回 *扁平列表的 `head` 。列表中的节点必须将其 **所有** 子指针设置为 `null` 。*



分析：

深度优先搜索

```go
// date 2023/10/17
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    dfs(root)
    return root
}

func dfs(root *Node) (last *Node) {
    cur := root
    for cur != nil {
        next := cur.Next
        if cur.Child != nil {
            chdRes := dfs(cur.Child)

            next = cur.Next  // save next

            // change cur Next to Child
            cur.Next = cur.Child
            cur.Child.Prev = cur

            if next != nil {
                chdRes.Next = next
                next.Prev = chdRes
            }

            cur.Child = nil
            last = chdRes
        } else {
            last = cur
        }

        // 遍历 next
        cur = next
    }
    return
}
```

