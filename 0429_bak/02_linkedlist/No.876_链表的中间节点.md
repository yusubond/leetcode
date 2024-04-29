## 876 链表的中间节点-简单

题目：

给你单链表的头节点 head，请你找出并返回链表的中间节点。

如果有两个中间节点，则返回第二个节点。



分析：

快慢指针。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```

