package template

type (
	StackNode struct {
		val       int
		pre, next *StackNode
	}

	MiddleStack struct {
		top *StackNode
		mid *StackNode
		cnt int
	}
)

func NewMiddleStack() *MiddleStack {
	return &MiddleStack{
		top: nil,
		mid: nil,
		cnt: 0,
	}
}

func (m *MiddleStack) Push(x int) {
	node := &StackNode{val: x, next: m.top}
	m.cnt += 1
	if m.cnt == 1 {
		// the first node
		m.mid = node
	} else {
		m.top.pre = node
		if m.cnt&0x1 == 0x1 {
			// 当有奇数个元素时, mid 上移一次
			m.mid = m.mid.pre
		}
	}
	m.top = node
}

func (m *MiddleStack) Pop() int {
	if m.IsEmpty() {
		return -1
	}

	x := m.top
	m.top = x.next
	// remove x
	if m.top != nil {
		m.top.pre = nil
	}
	m.cnt -= 1

	// 当元素个数由奇数变成偶数时，mid 向 next(bottom) 移动
	if m.cnt&0x1 == 0x0 {
		m.mid = m.mid.next
	}

	return x.val
}

func (m *MiddleStack) GetMiddle() int {
	if m.IsEmpty() {
		return -1
	}
	return m.mid.val
}

func (m *MiddleStack) DeleteMiddle() {
	if m.IsEmpty() {
		return
	}
	m.cnt -= 1

	// save the mid first, and delete it
	node := m.mid
	// 当元素个数从偶数变成奇数时，mid 向 pre(top) 移动
	if m.cnt&0x1 == 0x1 {
		// update mid
		m.mid = node.pre
		if m.cnt == 1 {
			m.mid.next = nil
		} else {
			// delete node
			m.mid.next = node.next
			node.next.pre = m.mid
		}
	} else {
		// 当元素个数从奇数变成偶数，mid 向 next(bottom) 移动
		if m.cnt == 0 {
			m.top, m.mid = nil, nil
		} else {
			m.mid = node.next
			node.pre.next = m.mid
			m.mid.pre = node.pre
		}
	}
}

func (m *MiddleStack) IsEmpty() bool {
	return m.cnt == 0
}
