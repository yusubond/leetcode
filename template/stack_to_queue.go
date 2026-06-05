package template

type (
	MyStack struct {
		val []int
	}
)

func NewMyStack() *MyStack {
	return &MyStack{
		val: make([]int, 0, 1024),
	}
}

func (s *MyStack) Push(x int) {
	s.val = append(s.val, x)
}

func (s *MyStack) Pop() int {
	x, n := -1, len(s.val)
	if n > 0 {
		x = s.val[n-1]
		s.val = s.val[:n-1]
	}
	return x
}

func (s *MyStack) Size() int {
	return len(s.val)
}

type (
	MyQueue struct {
		stackA *MyStack // 用于入队
		stackB *MyStack // 用于出队
	}
)

// NewMyQueue define
func NewMyQueue() *MyQueue {
	return &MyQueue{
		stackA: NewMyStack(),
		stackB: NewMyStack(),
	}
}

// EnQueue define
func (q *MyQueue) EnQueue(x int) {
	q.stackA.Push(x)
}

// DeQueue define
func (q *MyQueue) DeQueue() int {
	if q.stackB.Size() == 0 {
		for q.stackA.Size() > 0 {
			x := q.stackA.Pop()
			q.stackB.Push(x)
		}
	}
	return q.stackB.Pop()
}

// IsEmpty define
func (q *MyQueue) IsEmpty() bool {
	return q.stackA.Size() == 0 && q.stackB.Size() == 0
}
