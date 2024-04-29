## 2487 从链表中移除节点-中等

题目：

给你一个链表的头节点 `head` 。

移除每个右侧有一个更大数值的节点。

返回修改后链表的头节点 `head` 。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2022/10/02/drawio.png)
>
> ```
> 输入：head = [5,2,13,3,8]
> 输出：[13,8]
> 解释：需要移除的节点是 5 ，2 和 3 。
> - 节点 13 在节点 5 右侧。
> - 节点 13 在节点 2 右侧。
> - 节点 8 在节点 3 右侧。
> ```
>
> **示例 2：**
>
> ```
> 输入：head = [1,1,1,1]
> 输出：[1,1,1,1]
> 解释：每个节点的值都是 1 ，所以没有需要移除的节点。
> ```



**解题思路**：

这道题的要求是，如果右侧有值更大的节点，那么当前节点就应该移除，使得最终的链表是单调递减的。

这道题可以这样解，新的表头不容易确定，所以可以先反转链表，然后移除每个节点后面小于当前节点的节点。最后在反转链表即可。

```go
// date 2024/01/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    lr := reverse(head)
    cur := lr
    for cur != nil {
        if cur.Next != nil && cur.Next.Val < cur.Val {
            cur.Next = cur.Next.Next
            continue
        }
        cur = cur.Next
    }
    res := reverse(lr)
    return res
}

func reverse(head *ListNode) *ListNode {
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

