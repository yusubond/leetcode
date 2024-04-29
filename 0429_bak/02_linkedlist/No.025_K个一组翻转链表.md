## 25 K个一组翻转链表-困难

题目：

给你链表的头节点 `head` ，每 `k` 个节点一组进行翻转，请你返回修改后的链表。

`k` 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 `k` 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)
>
> ```
> 输入：head = [1,2,3,4,5], k = 2
> 输出：[2,1,4,3,5]
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)
>
> ```
> 输入：head = [1,2,3,4,5], k = 3
> 输出：[3,2,1,4,5]
> ```



**解题思路**

这道题可以递归翻转，具体为：

1. 先实现一个区间翻转的函数，满足K个直接进入区间翻转
2. 不足 k 个直接返回

```go
// date 2024/01/11
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    node := head
    for i := 0; i < k; i++ {
        if node == nil {
            return head
        }
        node = node.Next
    }
    newHead := reverse(head, node)  // reverse 之后 head.Next = node
    head.Next = reverseKGroup(node, k)
    return newHead
}

// reverse [first, last) then return new head
// not contains the last node
func reverseBetween(first, last *ListNode) *ListNode {
    prev := last
    for first != last {
        tmp := first.Next
        first.Next = prev
        prev = first
        first = tmp
    }
    return prev
}
```

![image](images/img025.svg)
