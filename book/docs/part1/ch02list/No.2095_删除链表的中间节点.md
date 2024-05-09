## 2095 删除链表的中间节点-中等

题目：

给你一个链表的头节点 `head` 。**删除** 链表的 **中间节点** ，并返回修改后的链表的头节点 `head` 。

长度为 `n` 链表的中间节点是从头数起第 `⌊n / 2⌋` 个节点（下标从 **0** 开始），其中 `⌊x⌋` 表示小于或等于 `x` 的最大整数。

- 对于 `n` = `1`、`2`、`3`、`4` 和 `5` 的情况，中间节点的下标分别是 `0`、`1`、`1`、`2` 和 `2` 。



**解题思路**

这道题有2种解法。一种是先遍历求链表长度，计算出中间节点，再次遍历，到达中间节点的时候删除之。详见解法1。

另一个种，快慢指针。快指针走到表尾时，慢指针刚好走到中间节点，删除之。详见解法2。



```go
// date 2024/10/12
// 解法1
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    cur := head
    cnt, mid := 0, 0
    for cur != nil {
        mid++
        cur = cur.Next
    }
    mid = mid >> 1

    prev := head
    cur = head
    cnt = 0
    for cur != nil {
        if cnt == mid {
            prev.Next = cur.Next
            break
        }
        cnt++
        prev = cur
        cur = cur.Next
    }
    return head
}

// 解法2
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    prev, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = slow.Next
    return head
}
```

