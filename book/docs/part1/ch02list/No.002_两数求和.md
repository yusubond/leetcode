## 002 两数相加-中等

题目：

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



分析：

哑结点，开辟空间计算

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dumy := &ListNode{}
    pre := dumy
    tempRes, n1, n2, carrry := 0, 0, 0, 0
    for l1 != nil || l2 != nil || carrry != 0 {
        n1, n2 = 0, 0
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }
        tempRes = n1 + n2 + carrry
        carrry = tempRes / 10

        pre.Next = &ListNode{Val: tempRes % 10}
        pre = pre.Next
    }
    return dumy.Next
}
```
