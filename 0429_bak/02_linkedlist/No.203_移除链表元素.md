## 203 移除链表元素-简单

题目：

给你一个链表的头节点 head 和一个整数 val，请你删除链表中所有满足 `node.val == val` 的节点，并返回新的头节点。



分析：

```go
// date 2023/10/16
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    dumy := &ListNode{}
    pre := dumy
    for head != nil {
        if head.Val != val {
            pre.Next = head
            pre = pre.Next
        }
        head = head.Next
    }
    pre.Next = nil
    return dumy.Next
}
```

