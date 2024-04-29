## 369 给单链表加1-中等

题目：

给定一个用**链表**表示的非负整数， 然后将这个整数 *再加上 1* 。

这些数字的存储是这样的：最高位有效的数字位于链表的首位 `head` 。



> **示例 1:**
>
> ```
> 输入: head = [1,2,3]
> 输出: [1,2,4]
> ```
>
> 
>
> **示例** **2:**
>
> ```
> 输入: head = [0]
> 输出: [1]
> ```



**解题思路**

因为低位在单链表的尾部，对低位加1不是很方便，所以需要先将链表反转，加1计算后，在反转一次即可。

```go
// date 2024/01/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    l2 := reverseList(head)
    carry := 1
    cur := l2
    prev := l2
    for cur != nil {
        after := cur.Next
        temp := cur.Val + carry
        cur.Val = temp % 10
        carry = temp / 10

        prev = cur
        cur = after
    }
    if carry == 1 {
        prev.Next = &ListNode{Val: 1}
    }
    l3 := reverseList(l2)
    return l3
}

func reverseList(head *ListNode) *ListNode {
    var tail *ListNode
    cur := head
    for cur != nil {
        after := cur.Next
        cur.Next = tail
        tail = cur
        cur = after
    }
    return tail
}
```

