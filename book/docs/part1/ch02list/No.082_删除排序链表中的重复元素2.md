## 82 删除排序链表中的重复元素2-中等

题目：

给定一个已排序的链表的头 `head` ， *删除原始链表中所有重复数字的节点，只留下不同的数字* 。返回 *已排序的链表* 。

这一题跟83题的区别在于，83题是要求删除重复节点，保证所有节点只出现一次。

而82题则要求删除所有重复出现的元素，只保留出现过一次的节点。



分析：

哑结点 和 快慢指针

其中，慢指针slow 从哑结点开始，并且 slow.next 指向 head，即未检查的数据（潜在的结果）。slow 指针保存最终的结果集。

快指针 fast 从 head 开始。

当 fast 与 fast.next 值不相等时，fast 是一个可能的结果，然后在通过 fast 与 slow 的距离来判断，fast 是否满足条件。

slow.next 与 fast 之间的数据就是重复元素。

```go
// date 2022/10/08
// 83升级版，所有重复元素都不要，只保留出现过一次的元素
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Val: -1}
    dummy.Next = head
    slow := dummy
    fast := head
    for fast.Next != nil {
        // 可能需要更新
        if fast.Val != fast.Next.Val {
            if slow.Next == fast {
                // 说明slow和fast之间没有重复元素，更新slow节点
                slow = fast
            } else {
                // 说明slow和fast之间有重复元素，更新slow的next节点
                slow.Next = fast.Next
            }
        }
        fast = fast.Next
    }
    // 查看尾部是否有重复
    // slow.Next != nil && slow.Next == fast 表示尾部没有重复，已经在17行进行更新了
    if slow.Next != nil && slow.Next != fast && slow.Next.Val == fast.Val {
        // 表示fast为重复元素，需要舍弃
        slow.Next = nil
    }
    return dummy.Next
}

// 解法2
//
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dumy := &ListNode{0, head}
	cur := dumy
	for cur.Next != nil && cur.Next.Next != nil {
		// 如果连续的两个值相等
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dumy.Next
}
```
