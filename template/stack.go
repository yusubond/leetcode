package template

import "fmt"

type (
	FixedSizeStack struct {
		cap int
		val []int
	}
)

func NewFixedSizeStack(n int) *FixedSizeStack {
	return &FixedSizeStack{
		cap: n,
		val: make([]int, 0, n),
	}
}

// Push define
func (s *FixedSizeStack) Push(x int) {
	if !s.IsFull() {
		s.val = append(s.val, x)
	}
}

// Pop define
func (s *FixedSizeStack) Pop() {
	if !s.IsEmpty() {
		s.val = s.val[:len(s.val)-1]
	}
}

// Top define
func (s *FixedSizeStack) Top() int {
	if !s.IsEmpty() {
		return s.val[len(s.val)-1]
	}
	return -1
}

// IsEmpty define
func (s *FixedSizeStack) IsEmpty() bool {
	return len(s.val) == 0
}

// IsFull define
func (s *FixedSizeStack) IsFull() bool {
	return len(s.val) == s.cap
}

// Size define
func (s *FixedSizeStack) Size() int {
	return len(s.val)
}

// pushBottom define
func (s *FixedSizeStack) pushBottom(x int) {
	if s.IsEmpty() {
		s.Push(x)
	} else {
		temp := s.Top()
		s.Pop()
		s.pushBottom(x)
		s.Push(temp)
	}
}

// Reverse define
func (s *FixedSizeStack) Reverse() {
	if !s.IsEmpty() {
		x := s.Top()
		s.Pop()
		s.Reverse()
		s.pushBottom(x)
	}
}

func (s *FixedSizeStack) Show() {
	if !s.IsEmpty() {
		fmt.Printf("stack = %d\n", s.val)
	}
}
