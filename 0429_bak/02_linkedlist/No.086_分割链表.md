## 86 分割链表-中等

题目：

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。



分析：

两个哑结点分别保存两个分区的数据。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    dumy1 := &ListNode{}
    dumy2 := &ListNode{}
    d1, d2 := dumy1, dumy2
    pre := head
    for pre != nil {
        if pre.Val < x {
            d1.Next = pre
            d1 = d1.Next
        } else {
            d2.Next = pre
            d2 = d2.Next
        }
        pre = pre.Next
    }
    d2.Next = nil
    d1.Next = dumy2.Next

    return dumy1.Next
}
```

