## 237 删除链表中的点-中等

题目：

有一个单链表 head，我们想删除链表中的一个节点 node。

已知给你一个需要删除的节点 node，你无法访问第一个节点 head。

链表的所有值都是唯一的，并且保证给定的节点 node 不是链表中的最后一个节点。

注意，删除节点并不是指从内存中删除它。这里的意思是：

- 给定节点的值不应该存在于链表中。
- 链表节点的数量应该减少1。
- `node`前面的所有值顺序相同。
- `node`后面的所有值顺序相同。



分析：

采用 后者赋值 的方法，让后面的值依次覆盖 node 开始的值就可以；然后把最后一个节点删除。

```go
// date 2023/10/13
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    if node == nil {
        return
    }
    pre := node
    for node.Next != nil {
        pre = node
        node.Val = node.Next.Val
        node = node.Next
    }
    pre.Next = nil    
}
```

