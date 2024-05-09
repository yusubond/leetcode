## 225 用队列实现栈-简单

题目：

请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（`push`、`top`、`pop` 和 `empty`）。

实现 `MyStack` 类：

- `void push(int x)` 将元素 x 压入栈顶。
- `int pop()` 移除并返回栈顶元素。
- `int top()` 返回栈顶元素。
- `boolean empty()` 如果栈是空的，返回 `true` ；否则，返回 `false` 。



分析：

每次压栈的时候，都把已有的元素先出队，再入队。这样，队头元素永远都是最后压栈的这个。

```go
// date 2023/11/30
type MyStack struct {
    queue []int
}


func Constructor() MyStack {
    return MyStack{
        queue: make([]int, 0, 16),
    }
}


func (this *MyStack) Push(x int)  {
    n := len(this.queue)
    this.queue = append(this.queue, x)
    i := 0
    for i < n {
        this.queue = append(this.queue, this.queue[i])
        i++
    }
    this.queue = this.queue[n:]
}


func (this *MyStack) Pop() int {
    res := this.queue[0]
    this.queue = this.queue[1:]
    return res
}


func (this *MyStack) Top() int {
    return this.queue[0]
}


func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
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

