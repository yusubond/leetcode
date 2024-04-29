package template

import "testing"

func TestNewDoubleLinkedList(t *testing.T) {
	list := NewDoubleLinkedList()
	list.Show()

	n1 := &DoubleListNode{val: 1}
	n2 := &DoubleListNode{val: 2}
	n3 := &DoubleListNode{val: 3}
	n4 := &DoubleListNode{val: 4}
	n5 := &DoubleListNode{val: 5}
	list.AddRear(n1)
	list.AddRear(n2)
	list.AddRear(n3)
	list.Show()
	list.AddFront(n4)
	list.AddFront(n5)
	list.Show()
	if 5 != list.Len() {
		t.Fatalf("fail")
	}

	list.RemoveRear()
	list.Show()

	list.RemoteFront()
	list.Show()
	if 3 != list.Len() {
		t.Fatalf("fail")
	}
}
