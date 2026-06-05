## 232 用栈实现队列-简单

题目：

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（`push`、`pop`、`peek`、`empty`）：

实现 `MyQueue` 类：

- `void push(int x)` 将元素 x 推到队列的末尾
- `int pop()` 从队列的开头移除并返回元素
- `int peek()` 返回队列开头的元素
- `boolean empty()` 如果队列为空，返回 `true` ；否则，返回 `false`

**说明：**

- 你 **只能** 使用标准的栈操作 —— 也就是只有 `push to top`, `peek/pop from top`, `size`, 和 `is empty` 操作是合法的。
- 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。



分析：

栈是先入后出的结构，队列是先入先出的结构。

如果要用栈实现队列，需要2个栈：

`stack_push`负责 push 存储入队的元素；另一个`stack_pop`负责 pop 出队，当`stack_pop` 为空时，从 `stack_push` 中倒一次手。

```go
// date 2023/11/29
type MyQueue struct {
    stackPush []int
    stackPop  []int
}


func Constructor() MyQueue {
    return MyQueue{
        stackPush: make([]int, 0, 16),
        stackPop: make([]int, 0, 16),
    }
}


func (this *MyQueue) Push(x int)  {
    this.stackPush = append(this.stackPush, x)
}


func (this *MyQueue) Pop() int {
    if len(this.stackPop) == 0 {
        n := len(this.stackPush)-1
        for n >= 0 {
            this.stackPop = append(this.stackPop, this.stackPush[n])
            n--
        }
        this.stackPush = this.stackPush[:0]
    }
    n := len(this.stackPop)
    v := this.stackPop[n-1]
    this.stackPop = this.stackPop[:n-1]
    return v
}


func (this *MyQueue) Peek() int {
    n := len(this.stackPop)
    if n > 0 {
        return this.stackPop[n-1]
    }
    m := len(this.stackPush)-1
    for m >= 0 {
        this.stackPop = append(this.stackPop, this.stackPush[m])
        m--
    }
    this.stackPush = this.stackPush[:0]
    return this.stackPop[len(this.stackPop)-1]
}


func (this *MyQueue) Empty() bool {
    return len(this.stackPush) == 0 && len(this.stackPop) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
```

