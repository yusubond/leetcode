## 109 有序链表转换为二叉搜索树-中等

题目：

给定一个单链表的头节点 head，其中的元素按升序排列，将其转换为高度平衡的二叉搜索树。

一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差不超过 1。



分析：

这个题目跟 108 题目类似，只不过从【有序数组】换成了【有序链表】。但思路是一样的，找到中间节点，递归构造。

```go
// date 2023/10/23
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return &TreeNode{Val: head.Val}
    }

    var pre, slow, fast *ListNode
    slow = head
    fast = head
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next

    }
    
    pre.Next = nil
    root := &TreeNode{Val: slow.Val}
    root.Left = sortedListToBST(head)

    slow = slow.Next
    if slow == nil {
        return root
    }
    root.Right = sortedListToBST(slow)

    return root
}
```

