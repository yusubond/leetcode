## 225 用队列实现栈-简单

题目：

请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（`push`、`top`、`pop` 和 `empty`）。

实现 `MyStack` 类：

- `void push(int x)` 将元素 x 压入栈顶。
- `int pop()` 移除并返回栈顶元素。
- `int top()` 返回栈顶元素。
- `boolean empty()` 如果栈是空的，返回 `true` ；否则，返回 `false` 。



分析：

两个队列，入栈的时候倒一手。具体为：

`recv`队列：维护 push 的数据，每次 push 时，先将新入栈的元素入队`recv`。如果`send`队列不为空，表示之前没有pop出所有的元素，所以需要将`send`队列的元素也依次入队`recv`。如此操作后，`send`队列为空，`recv`队列中最新元素在队头，次新的在队尾。

然后，交换`recv`和`send`队列。

pop 的时候，依次从`send`队列出队即可。

```go
// date 2023/11/30
type MyStack struct {
	recv []int
	send []int
}

func Constructor() MyStack {
	return MyStack{
		recv: make([]int, 0),
		send: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.recv = append(this.recv, x)
	for len(this.send) > 0 {
		this.recv = append(this.recv, this.send[0])
		this.send = this.send[1:]
	}
	// exchange recv and send
	this.recv, this.send = this.send, this.recv
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	x := this.send[0]
	this.send = this.send[1:]
	return x
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.send[0]
}

func (this *MyStack) Empty() bool {
	return len(this.send) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
```

