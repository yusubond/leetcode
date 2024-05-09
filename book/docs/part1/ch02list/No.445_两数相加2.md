## 445 两数相加-中等

题目：

给你两个 **非空** 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



> **示例1：**
>
> ![img](https://pic.leetcode-cn.com/1626420025-fZfzMX-image.png)
>
> ```
> 输入：l1 = [7,2,4,3], l2 = [5,6,4]
> 输出：[7,8,0,7]
> ```
>
> **示例2：**
>
> ```
> 输入：l1 = [2,4,3], l2 = [5,6,4]
> 输出：[8,0,7]
> ```
>
> **示例3：**
>
> ```
> 输入：l1 = [0], l2 = [0]
> 输出：[0]
> ```



分析：

先讲原始链表反转，然后顺次相加。

```go
// date 2023/12/18
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }

    l1 = reverseList(l1)
    l2 = reverseList(l2)

    v1, v2, carry := 0, 0, 0
    var res *ListNode

    for l1 != nil || l2 != nil || carry > 0 {
        v1, v2 = 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        temp := v1 + v2 + carry
        carry = temp / 10

        head := &ListNode{Val: temp%10}
        head.Next = res
        res = head
    }

    return res
}

func reverseList(l *ListNode) *ListNode {
    if l == nil || l.Next == nil {
        return l
    }
    var tail *ListNode
    pre := l
    for pre != nil {
        pn := pre.Next
        pre.Next = tail
        tail = pre
        pre = pn
    }
    return tail
}
```

