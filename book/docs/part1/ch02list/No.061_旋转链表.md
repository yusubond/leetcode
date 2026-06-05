#### No.61 旋转链表【中等】

题目：

给你一个链表的头节点 head，旋转链表，将链表的每个节点向右移动 k 个位置。



分析：【推荐使用该算法】

把链表每个节点向右移动 k 个位置，相当于把倒数 k 个节点形成新链表的前半部分，原来前面的 l - k 个形成新链表的后半部分。那么，l - k - 1 位置（从0计算索引）的节点就是新链表的最后一个节点，其Next就是新头节点。

具体这样做：

1. 求链表求长度 l，并将其形成环
2. 对 k 取模，k = k % l
3. 重新从 head 开始走 l - k - 1 步，找到新链表的尾巴，断开即可。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// date 2023/10/11
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    pre := head
    l := 1
    for pre.Next != nil {
        l++
        pre = pre.Next
    }
    pre.Next = head

    k = k % l
    step := l - k - 1
    pre = head
    for step > 0 {
        pre = pre.Next
        step -= 1
    }
    head = pre.Next
    pre.Next = nil
    return head
}
```

