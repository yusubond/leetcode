# Two Stack

利用数组Array实现一种双栈数据结构twoStack，towStack满足以下功能：

+ push1(int x)：向栈1压入元素x
+ push2(int x)：向栈2压入元素x
+ pop1()：删除栈1中的栈顶元素，并返回删除的元素值
+ pop2()：删除栈2中的栈顶元素，并返回删除的元素值

注意：实现twoStack的过程中尽量保证空间效率。



思路分析：

方法1：共有1个数组，对半分。

将数组分为两部分，前`[0,n/2]`用于栈1，后`[n/2+1, n]`用于栈2，但这样分配存在一个问题，当其中一个栈已满时，另一个栈可能还有可用的空间，造成空间上的浪费。

方法2：还是共用1个数组，动态维护两个栈的栈顶。

具体为，分别将`0`和`n-1`作为两个栈的栈底，从数组两端向中间存储数据，当两个栈的栈顶`top1`和`top2`相等时，栈已满。



下面是方法2的实现。

```go
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
```

