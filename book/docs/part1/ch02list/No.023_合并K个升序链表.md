## 023 合并K个升序链表-困难

题目：

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



**解题思路**

复习递归算法，掌握自底向上和自上向下的递归方法。

- 最朴素的算法，顺次两两合并，详见解法1。递归算法，自底向上的两两合并。每个链表遍历一遍，所示时间复杂度O(n*k)，空间复杂度O(1)。
- 分治算法，类似归并排序，详见解法2。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1
// 最朴素的顺序两两合并
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    k := len(lists)
    l1, l2 := mergeKLists(lists[:k-1]), lists[k-1]
    dumy := &ListNode{}
    pre := dumy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }
    return dumy.Next
}

// 解法2
// 分治思想，归并排序思想的合并
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    
    var merge func(lists []*ListNode, left, right int) *ListNode
    merge = func(lists []*ListNode, left, right int) *ListNode {
        if left == right {
            return lists[left]
        }
        if left > right {
            return nil
        }
        mid := left + (right- left)/2
        return mergeTwoList(merge(lists, left, mid), merge(lists, mid+1, right))
    }

    return merge(lists, 0, len(lists)-1)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    }
    if l2 != nil {
        cur.Next = l2
    }
    return dummy.Next
}
```

