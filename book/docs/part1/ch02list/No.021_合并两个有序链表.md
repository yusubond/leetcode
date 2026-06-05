## 021 合并两个有序链表-简单

题目：

将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 



分析：

算法1：利用哑结点

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    l1, l2 := list1, list2
    dumy := &ListNode{}
    pre := dumy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }
    return dumy.Next
}
```



算法2：递归版

每个节点只会访问一次，因此时间复杂度为O(n+m)；只有到达l1或l2的底部，递归才会返回，所以空间复杂度为O(n+m)

```go
// 递归版
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
  if l1 == nil {return l2}
  if l2 == nil {return l1}
  var newHead *ListNode
  if l1.Val < l2.Val {
    newHead = l1
    newHead.Next = mergeTwoLists(l1.Next, l2)
  } else {
    newHead = l2
    newHead.Next = mergeTwoLists(l1, l2.Next)
  }
  return newHead
}
```
