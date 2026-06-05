## 141 环形链表-简单

题目：

给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。



**解题思路**

通过快慢指针是否相遇来判断链表中是否存在环。

```go
// date 2023/10/11
// 算法1：快慢指针
// 时间复杂度分析
// 如果链表中不存在环，则fast指针最多移动N/2次；如果链表中存在环，则fast指针需要移动M次才能与slow指针相遇，M为环的元素个数。
// 总体来讲，时间复杂度为O(N)
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

