## 关于单链表的其他问题 | set 1

[TOC]

单链表节点的定义

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
```



### 1、判断一个链表是否存在环

解法：利用 **快慢指针** 遍历链表，当快慢两个指针在某一时刻指向同一个节点时，则存在环，否则不存在环。

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

### 2、合并两个有序链表

注意保存新链表的头节点。

```go
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

### 3、在不给出链表的头指针条件下，删除链表中的某个节点

解法：将准备删除的节点的下一个节点的值拷贝至将要删除的节点，后删除下一个节点。需要考虑，这个节点是否为尾节点。

```go
func deleteNode(node *ListNode) {
    if node == nil {
        return
    }
    pre := node
    for node.Next != nil {
        pre = node
        node.Val = node.Next.Val
        node = node.Next
    }
    pre.Next = nil    
}
```

### 4、删除一个已排序链表中的重复元素

两种思路，详见 83 题。



### 5、删除一个未排序链表中重复元素

这类题的要求是只保留链表中元素值出现1次的节点。详见 1836 题。

需要借助 map 统计出现次数。

```go
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

### 6、将链表中的最后一个元素移动至链表的最前面

直接遍历，记录前驱节点和当前节点，当当前节点是最后一个节点时，直接指向 head，从其前驱节点断开即可。

```go
func MoveTheLastToFront(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	prev, cur := head, head
	for cur != nil {
		if cur.next == nil {
			cur.next = head
			prev.next = nil
			break
		}
		prev = cur
		cur = cur.next
	}
	return cur
}
```



### 7、删除偶数位置上的节点

```go
func DeleteEven(head *ListNode) *ListNode {
  cur := head
  for cur != nil && cur.Next != nil {
    cur.Next = cur.Next.Next
    cur = cur.Next
  }
  return head
}
```

### 8、反转固定个数的节点

例如：1->2->3->4->5->6->7->8->NULL;当k=3时，反转后为3->2->1->6->5->4->8->7->NULL。

```
Node* Reverse(Node *phead, int k) {
  Node *p = phead;
  Node *pnext = NULL;
  Node *pre = NULL;
  int count = 0;
  while(p != NULL && count < k) {
    pnext = p->next;
    p->next = pre;
    pre = p;
    p = pnext;
    count++;
  }
  if(pnext != NULL)
    phead->next = Reverse(next, k);
  return pre;
}
```

### 9、交替反转链表中固定个数的节点

输入：1->2->3->4->5->6->7->8 k=3
输出：3->2->1->4->5->6->8->7

```
Node *KReverse(Node *phead, int k) {
  Node *p = phead;
  Node *pnext = NULL;
  Node *pre = NULL;
  count = 0;
  while(p != NULL && count < k) {
    pnext = p->next;
    p->next = pre;
    pre = p;
    p = pnext;
    count++;
  }
  if(phead != NULL)
    phead->next = p;
  count = 0;
  while(count < (k - 1) && p != NULL) {
    p = p->next;
    count++;
  }
  if(p != NULL)
    p->next = KReverse(p->next, k);
  return pre;
}
```

### 10、单调递减链表

解释：Delete nodes which have a greater value on right side

保证原有链表节点相对位置不变的情况下，使之变为单调递减链表。

输入：12->15->10->11->5->6->2->3
输出：15->11->6->3

如果当前节点的值比其下一个节点的值小，则删除当前节点，保留其下一个节点。

解法：先将链表反转，这样程序的逻辑就变为如果后一个节点小于前一个节点，则删除后一个节点。这样做的目的是 **确定新链表的头部**。原来的程序逻辑不太容易确定链表的头部，但是尾部相对容易确定，

这就是第 2487 题。

```go
func (l *LinkedList) DeleteSmallerRight() {
	if l.size > 1 {
		l.Reverse()
		cur := l.head.next
		for cur != nil && cur.next != nil {
			if cur.val > cur.next.val {
				cur.next = cur.next.next
				l.size--
				continue
			}
			cur = cur.next
		}
		l.Reverse()
	}
}
```
