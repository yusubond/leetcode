## 234 回文链表-简单

题目：

给你一个单链表的头节点 head，请你判断该链表是否为回文链表。如果是，返回 true，否则，返回 false。



**解题思路**

- 数组，双指针，详见解法1。遍历链表，将元素存入数组，然后判断数组是否回文。时间复杂度`O(N)`，空间复杂度`O(N)`。
- 双指针，快慢指针，详见解法2。利用快慢指针找到链表的中点，同时反转前半段，然后比较快慢指针的元素。时间复杂度`O(N)`，空间复杂度`O(1)`。

快慢指针，找到中间节点，反转前半段。

```go
// date 2023/10/16
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1
// 数组
func isPalindrome(head *ListNode) bool {
    nums := make([]int, 0, 16)
    cur := head
    for cur != nil {
        nums = append(nums, cur.Val)
        cur = cur.Next
    }
    left, right := 0, len(nums)-1
    for left < right {
        if nums[left] != nums[right] {
            return false
        }
        left++
        right--
    }
    return true
}

// 解法2
// 快慢指针
func isPalindrome(head *ListNode) bool {
    var leftTail *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next

        slowNext := slow.Next
        slow.Next = leftTail
        leftTail = slow
        slow = slowNext
    }
    // 奇数个节点
    if fast != nil {
        slow = slow.Next
    }
    // now slow is the right part
    for slow != nil && leftTail != nil {
        if slow.Val != leftTail.Val {
            return false
        }
        slow = slow.Next
        leftTail = leftTail.Next
    }
    return true
}
```

