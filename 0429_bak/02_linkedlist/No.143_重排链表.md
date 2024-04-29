## 143 重排链表-中等

题目：

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



分析：

快慢指针，找出链表的中间节点；slow 指针继续走，同时反转后半段链表。

然后将前半段和后半段链表合并。



要对每一段代码很熟悉，要精确的知道每段代码的结果是什么样子；然后大问题拆成小问题去解决。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    var tail *ListNode
    slow, fast := head, head
    spre := head
    for fast != nil && fast.Next != nil {
        spre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast != nil {
        spre = slow
        slow = slow.Next
    }
    // make part one, [head, spre]
    spre.Next = nil
    part1 := head

    for slow != nil {
        sn := slow.Next
        slow.Next = tail
        tail = slow
        slow = sn
    }

    head = merge(part1, tail)
}

func merge(l1, l2 *ListNode) *ListNode{
    dumy := &ListNode{}
    pre := dumy
    for l1 != nil && l2 != nil {
        l1f := l1.Next
        l2f := l2.Next

        l1.Next = l2
        pre.Next = l1
        pre = l2

        l1 = l1f
        l2 = l2f
    }
    if l1 != nil {
        pre.Next = l1
    }
    return dumy.Next
}
```

