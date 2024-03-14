## No.092 反转链表2

题目：

给你单链表的头指针 head 和两个整数 left, right，其中 left <= right。请你反转从位置 left 到 位置 right 的链表节点，返回反转后的链表。



分析：

使用一趟扫描，在扫描过程中：

1. 先找到 left 及其前驱节点 leftpre，并保存
2. 同时将 left 节点保存为 lefttail，继续遍历，并反转 [left, right] 中的节点
3. 找到 right 节点之后，更新 left, right 节点的Next值
4. 确定新的 head

```go
// date 2023/10/11
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    // leftpre, left, right, rightpost
    // leftpre.Next = right
    // left.Next = rightpost
    if head == nil || head.Next == nil {
        return head
    }
    // 区间内没有节点，不需要反转
    if left == right {
        return head
    }

    var leftPre, tail, cur *ListNode
    cur = head
    for left > 1 {
        leftPre = cur
        cur = cur.Next
        left--
        right--
    }
    // now cur is the left node
    // save it
    l := cur
    for right > 1 {
        after := cur.Next
        cur.Next = tail
        tail = cur
        cur = after
        right--
    }
    // now cur is the right node
    l.Next = cur.Next
    cur.Next = tail
    // when left > 1, the head is newHead
    if leftPre != nil {
        leftPre.Next = cur
        return head
    }
    // when left = 1, the right is new head
    return cur
}
```

![image](images/image092.png)



解法2：实现区间节点反转函数。

```go
// date 2024/03/14
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right {
        return head
    }
    cur := head
    leftPre := head
    for left > 1 {
        leftPre = cur
        cur = cur.Next
        left--
        right--
    }
    // now cur is the left
    leftNode := cur
    for right > 1 {
        cur = cur.Next
        right--
    }
    rightNode := cur.Next

    if leftNode == head {
        return reverse(leftNode, rightNode)
    }
    leftPre.Next = reverse(leftNode, rightNode)
    return head
}

// [first, last)
func reverse(first, last *ListNode) *ListNode {
	prev := last
	for first != last {
		temp := first.Next

		first.Next = prev
		prev = first

		first = temp
	}
	return prev
}
```

