package template

type (
	TwoStack struct {
		top1, top2 int
		data       []int
	}
)

func NewTwoStack(size int) *TwoStack {
	return &TwoStack{
		top1: -1,
		top2: size,
		data: make([]int, size, size),
	}
}

func (t *TwoStack) Push1(x int) {
	if !t.IsFull() {
		t.top1++
		t.data[t.top1] = x
	}
}

func (t *TwoStack) Push2(x int) {
	if !t.IsFull() {
		t.top2--
		t.data[t.top2] = x
	}
}

func (t *TwoStack) Pop1() int {
	if t.top1 >= 0 {
		x := t.data[t.top1]
		t.top1--
		return x
	}
	return -1
}

func (t *TwoStack) Pop2() int {
	if t.top2 < len(t.data) {
		x := t.data[t.top2]
		t.top2++
		return x
	}
	return -1
}

func (t *TwoStack) IsFull() bool {
	return t.top1+1 == t.top2
}
