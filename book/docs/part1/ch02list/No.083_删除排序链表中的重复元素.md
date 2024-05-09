## 83 删除排序链表中的重复元素-简单

题目：

给定一个已排序的链表的头 `head` ， *删除所有重复的元素，使每个元素只出现一次* 。返回 *已排序的链表* 。



**解题思路**

这道题有2种解法。一种是利用哑结点，遍历链表，只连接值不同的节点，详见解法1。

另一个解法是直接判断当前指针 cur 与当前指针的Next值是否一致，如果是，直接删掉后一个，详见解法2。

解法2更优雅一些。

```go
// date 2022/10/08
// 解法1
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dumy := &ListNode{}
    pre := dumy
    // add head
    pre.Next = head
    pre = head

    lastVal := head.Val

    cur := head.Next
    for cur != nil {
        if cur.Val != lastVal {  // find a new node
            pre.Next = cur
            pre = pre.Next
            lastVal = cur.Val // update last value
        }
        cur = cur.Next
    }
    pre.Next = nil
    return dumy.Next
}
// 解法2
func deleteDuplicates(head *ListNode) *ListNode {
    cur := head
    for cur != nil && cur.Next != nil {
        if cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
            continue
        }
        cur = cur.Next
    }
    return head
}
```

