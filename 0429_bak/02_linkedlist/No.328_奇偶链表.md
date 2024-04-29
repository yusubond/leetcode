## 328 奇偶链表-中等

题目：

给定单链表的头节点 head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排列的链表。

第一个节点的索引被认为是奇数，第二个节点的索引被认为是偶数。

请注意，偶数组和奇数组内部的相对位置应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(N) 的时间复杂度下解决这个问题。



分析：

这个问题过于简单了，就是按索引重新分组。

利用两个【哑结点】，分别保存索引为奇数和偶数的节点，然后重新拼接。

```go
// date 2023/10/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }
    d1, d2 := &ListNode{}, &ListNode{}
    p1, p2 := d1, d2
    idx := 0
    for head != nil {
        if idx % 2 == 0 {
            p1.Next = head
            p1 = p1.Next
        } else {
            p2.Next = head
            p2 = p2.Next
        }
        idx++
        head = head.Next
    }
    p2.Next = nil
    p1.Next = d2.Next
    return d1.Next
}
```



## 延展一下

不以索引的奇偶作为分组标准，而是以链表中结点值的奇偶性作为分组的标准。

请注意，以第一个节点值的奇偶性作为起始。即如果第一个节点值为奇数，那么所有的奇数在前。