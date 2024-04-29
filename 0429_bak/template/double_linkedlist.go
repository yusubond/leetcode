package template

import (
	"fmt"
	"strings"
)

type (
	DoubleListNode struct {
		val       int
		pre, next *DoubleListNode
	}
	DoubleLinkedList struct {
		size       int
		head, tail *DoubleListNode
	}
)

func NewDoubleLinkedList() *DoubleLinkedList {
	list := &DoubleLinkedList{
		size: 0,
		head: &DoubleListNode{},
		tail: &DoubleListNode{},
	}
	list.head.next = list.tail
	list.tail.pre = list.head
	return list
}

func (l *DoubleLinkedList) AddFront(node *DoubleListNode) {
	node.pre = l.head
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
	l.size++
}

func (l *DoubleLinkedList) AddRear(node *DoubleListNode) {
	node.pre = l.tail.pre
	node.next = l.tail
	l.tail.pre.next = node
	l.tail.pre = node
	l.size++
}

func (l *DoubleLinkedList) RemoteFront() *DoubleListNode {
	if l.size > 0 {
		node := l.head.next
		l.removeNode(node)
		return node
	}
	return nil
}

func (l *DoubleLinkedList) RemoveRear() *DoubleListNode {
	if l.size > 0 {
		node := l.tail.pre
		l.removeNode(node)
		return node
	}
	return nil
}

func (l *DoubleLinkedList) removeNode(node *DoubleListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
}

func (l *DoubleLinkedList) Len() int {
	return l.size
}

func (l *DoubleLinkedList) Show() {
	k := l.size
	str := make([]string, 0, k)
	if k > 0 {
		fmt.Printf("total node: %d\n", k)
	}
	pre := l.head.next
	for k > 0 {
		str = append(str, fmt.Sprintf("%d", pre.val))
		pre = pre.next
		k--
	}
	if len(str) != 0 {
		fmt.Printf("list: %s\n", strings.Join(str, "<->"))
	}
}
