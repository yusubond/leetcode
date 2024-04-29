package template

import (
	"fmt"
	"strings"
)

type (
	Node struct {
		val  int
		next *Node
	}
	LinkedList struct {
		size int
		head *Node
	}
)

func NewLinkedList() *LinkedList {
	return &LinkedList{
		size: 0,
		head: &Node{},
	}
}

func (l *LinkedList) AddFront(node *Node) {
	node.next = l.head.next
	l.head.next = node
	l.size++
}

func (l *LinkedList) AddRear(node *Node) {
	pre := l.head
	for pre.next != nil {
		pre = pre.next
	}
	pre.next = node
	l.size++
}

func (l *LinkedList) Size() int {
	return l.size
}

// Reverse define
// 翻转链表
func (l *LinkedList) Reverse() {
	if l.size > 0 {
		var tail *Node
		pre := l.head.next
		for pre != nil {
			after := pre.next
			pre.next = tail
			tail = pre
			pre = after
		}
		l.head.next = tail
	}
}

// InsertSort define
// 对链表进行插入排序
func (l *LinkedList) InsertSort() {
	if l.size < 2 {
		return
	}
	dummy := &Node{}
	first := true
	pre := l.head.next
	for pre != nil {
		after := pre.next
		// insert pre to dummy
		pre.next = nil
		if first {
			first = false
			dummy.next = pre
		} else {
			cur := dummy.next
			curP := dummy
			for cur != nil && cur.val < pre.val {
				curP = cur
				cur = cur.next
			}
			curP.next = pre
			pre.next = cur
		}

		pre = after
	}

	l.head.next = dummy.next
}

// MoveTheLastToFront define
// 把链表中最后一个元素移到表头
func (l *LinkedList) MoveTheLastToFront() {
	if l.size > 0 {
		prev, cur := l.head.next, l.head.next
		for cur != nil {
			if cur.next == nil {
				cur.next = l.head.next
				prev.next = nil
				l.head.next = cur
				break
			}
			prev = cur
			cur = cur.next
		}
	}
}

// DeleteEvenNode define
// 删除偶数位置上的节点
func (l *LinkedList) DeleteEvenNode() {
	if l.size > 1 {
		cur := l.head.next
		for cur != nil && cur.next != nil {
			cur.next = cur.next.next
			cur = cur.next
			l.size--
		}
	}
}

// DeleteSmallerRight define
/*
保证原有链表节点相对位置不变的情况下，使之变为单调递减链表。
输入：12->15->10->11->5->6->2->3
输出：15->11->6->3
*/
/*
思路：先反转链表，然后遍历，如果后一个节点小于当前节点，则删除后一个节点。这与题意一致。
先反转的目的是更容易确定新的表头
*/
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

// ReverseBetween define
// 92. 区间反转
// 单链表一个共有 n 的节点，反转闭区间[left, right]中的节点，其中
// 1 <= left <= right <= n
func (l *LinkedList) ReverseBetween(left, right int) {
	// 只有一个节点的时候，不存在反转
	for l.size > 2 {
		// 必须区间只有一个节点，反转之后，原链表不变，直接返回即可
		if left == right {
			return
		}
		// 参数校验
		if left < 1 || left > l.size {
			return
		}
		if right < 1 || right > l.size {
			return
		}
		if left > right {
			return
		}
		// leftpre, left, right, post
		// leftpre.Next = right, if leftpre is nil, right is the new head
		// left.Next = post
		var leftPre *Node
		cur := l.head.next
		for left > 1 {
			leftPre = cur
			cur = cur.next
			left--
			right--
		}
		// now the cur is left node
		leftNode := cur
		tail := leftNode
		for right > 1 {
			after := cur.next
			cur.next = tail
			tail = cur
			cur = after
			right--
		}
		// now cur is the right node
		rightNode := cur
		leftNode.next = rightNode.next
		rightNode.next = tail
		if leftPre != nil {
			leftPre.next = rightNode
			return
		}
		l.head.next = rightNode
	}
}

func (l *LinkedList) Show() {
	if l.size > 0 {
		str := make([]string, 0, l.size)
		pre := l.head.next
		for pre != nil {
			str = append(str, fmt.Sprintf("%d", pre.val))
			pre = pre.next
		}

		fmt.Printf("total node: %d\n", l.size)
		fmt.Printf("linkedlist = %s\n", strings.Join(str, "<-"))
		fmt.Println()
	}
}

func MoveTheLastToFront(head *Node) *Node {
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
