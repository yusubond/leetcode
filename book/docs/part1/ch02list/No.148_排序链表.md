## 148 排序链表-中等

题目：

给你链表的头节点 head，请你将其按升序排列并返回排序后的链表。

你可以在`o(NlogN)`时间复杂度和常数级空间复杂度下，对链表进行排序吗？



**解题思路**

要想实现 `O(nlogn)`的时间复杂度，可以考虑归并排序。具体为，自顶向下的归并排序：

1. 找到链表的中点，以中点为界分成两个子链。
2. 对子链分别进行排序。
3. 对排序的子链，进行合并。

这里依赖基本 case，即子链中只有一个元素时，就是递归的返回原点。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    spre, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        spre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast != nil {
        spre = slow
        slow = slow.Next
    }
    spre.Next = nil
    left, right := sortList(head), sortList(slow)
    return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    prev := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            prev.Next = l1
            l1 = l1.Next
        } else {
            prev.Next = l2
            l2 = l2.Next
        }
        prev = prev.Next
    }
    if l1 != nil {
        prev.Next = l1
    }
    if l2 != nil {
        prev.Next = l2
    }
    return dummy.Next
}
```



