## 147 对链表进行插入排序-中等

题目：

给定单链表的头 head，使用 **插入排序** 对链表进行排序，并返回 排序后的链表的头。



分析：

为了代码简洁我们写一个已排序的链表中插入单个节点的辅助函数。

```go
// date 2023/10/16
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    cur := head.Next

    newHead := head
    newHead.Next = nil

    for cur != nil {
        after := cur.Next
        cur.Next = nil
        newHead = insertOneNode(newHead, cur)
        cur = after
    }
    return newHead
}

func insertOneNode(head, node *ListNode) *ListNode {
    if node.Val < head.Val {
        node.Next = head
        head = node
        return head
    }
    pre := head
    cur := head.Next
    for cur != nil && cur.Val <= node.Val {
        pre = cur
        cur = cur.Next
    }
    node.Next = cur
    pre.Next = node
    return head
}
```

