## 206 反转链表-简单

题目：

给你单链表的头节点 `head` ，请你反转链表，并返回反转后的链表。



**解题思路**

直接迭代，时间复杂度O(n)，空间复杂度O(1)。

```go
// 算法1：迭代
func reverseList(head *ListNode) *ListNode {
    var pre, after *ListNode
    for head != nil {
        after = head.Next
        head.Next = pre
        pre, head = head, after
    }
    return pre
}
```

算法2：假设链表为n1->n2->...->nk-1->nk->nk+1->...nm；如果从节点nk+1到结点nm均已经反转，即

n1->n2->...->nk-1->nk->nk+1<-...<-nm那么nk的操作则需要nk.Next.Next = nk

时间复杂度O(n)，空间复杂度O(n)，因为递归会达到n层。

```go
// 算法2：递归
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {return head}
  pre := reverseList(head.Next)
  head.Next.Next = head
  head.Next = nil
  return pre
}
```

#### 