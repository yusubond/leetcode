## 622 设计循环队列-中等

题目：

设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：

- `MyCircularQueue(k)`: 构造器，设置队列长度为 k 。
- `Front`: 从队首获取元素。如果队列为空，返回 -1 。
- `Rear`: 获取队尾元素。如果队列为空，返回 -1 。
- `enQueue(value)`: 向循环队列插入一个元素。如果成功插入则返回真。
- `deQueue()`: 从循环队列中删除一个元素。如果成功删除则返回真。
- `isEmpty()`: 检查循环队列是否为空。
- `isFull()`: 检查循环队列是否已满。



分析：

循环队列可以用数组实现，用 front, rear 分别指向队列的队头元素和队尾元素。

初始化，front, rear 都指向 -1，这也是队列为空的判断标准。

为空判断：front, rear 是否指向 -1。

队满判断：队尾 rear 再走一步就到队头 front 了，则循环队列满了。

入队：如果为空，front++, rear++；满了则直接返回，否则 rear++

出队：如果为空，直接返回；如果 front == rear，那么只有一个元素，直接清空；否则 front++。



```go
// date 2023/11/30
type MyCircularQueue struct {
    queue []int
    front, rear int
    size int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        queue: make([]int, k, k),
        front: -1,
        rear: -1,
        size: k,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsEmpty() {
        this.front++
        this.rear++
        this.queue[this.rear] = value
        return true
    }
    if this.IsFull() {
        return false
    }
    this.rear = (this.rear+1) % this.size
    this.queue[this.rear] = value
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
        return true
    }
    this.front = (this.front+1) % this.size
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.front]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.front == -1
}


func (this *MyCircularQueue) IsFull() bool {
    return (this.rear + 1) % this.size == this.front
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
```

