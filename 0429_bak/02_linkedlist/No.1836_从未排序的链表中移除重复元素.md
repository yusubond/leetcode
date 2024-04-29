## 1836 从未排序的链表中移除重复元素-中等

题目：

给定一个链表的第一个节点 `head` ，找到链表中所有出现**多于一次**的元素，并删除这些元素所在的节点。

返回删除后的链表。



> ```
> 输入: head = [1,2,3,2]
> 输出: [1,3]
> 解释: 2 在链表中出现了两次，所以所有的 2 都需要被删除。删除了所有的 2 之后，我们还剩下 [1,3] 。
> ```



**解题思路**

这道题的要求是把出现次数大于1次的节点都去掉，并保持原节点的相对位置不变。

一种思路是借助 map 记录节点出现的次数，详见解法1。

具体是第一次遍历，把节点的值存入 map，记录次数；第二次遍历，借助哑结点把出现次数为1的节点加入哑结点所在的链表，最后返回哑结点的链表。这个解法有个优化点，就是第一次遍历的时候，后面重复的节点可以先删除，以便第二次连接的时候少操作一部分。



```go
// date 2024/01/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    cnt := make(map[int]int, 4)
    cur := head
    prev := head
    for cur != nil {
        cnt[cur.Val]++
        // 优化1: 遍历的同时直接去掉后面出现的 重复的
        if cnt[cur.Val] > 1 && cur.Next != nil {
            prev.Next = cur.Next
        }
        prev = cur
        cur = cur.Next
    }

    dummy := &ListNode{}
    pre := dummy
    cur = head
    for cur != nil {
        if cnt[cur.Val] == 1 {
            pre.Next = cur
            pre = cur
        }
        cur = cur.Next
    }
    pre.Next = nil
    return dummy.Next
}
```

