## 641 设计循环双端队列-中等

题目：

设计实现双端队列。

实现 `MyCircularDeque` 类:

- `MyCircularDeque(int k)` ：构造函数,双端队列最大为 `k` 。
- `boolean insertFront()`：将一个元素添加到双端队列头部。 如果操作成功返回 `true` ，否则返回 `false` 。
- `boolean insertLast()` ：将一个元素添加到双端队列尾部。如果操作成功返回 `true` ，否则返回 `false` 。
- `boolean deleteFront()` ：从双端队列头部删除一个元素。 如果操作成功返回 `true` ，否则返回 `false` 。
- `boolean deleteLast()` ：从双端队列尾部删除一个元素。如果操作成功返回 `true` ，否则返回 `false` 。
- `int getFront()` )：从双端队列头部获得一个元素。如果双端队列为空，返回 `-1` 。
- `int getRear()` ：获得双端队列的最后一个元素。 如果双端队列为空，返回 `-1` 。
- `boolean isEmpty()` ：若双端队列为空，则返回 `true` ，否则返回 `false` 。
- `boolean isFull()` ：若双端队列满了，则返回 `true` ，否则返回 `false` 。



分析：

跟单端队列类似，处理好 front, rear 即可。

```go
// date 2023/11/30
type MyCircularDeque struct {
    queue []int
    front, rear int
    size int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        queue: make([]int, k, k),
        front: -1,
        rear: -1,
        size: k,
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsEmpty() {
        this.front++
        this.rear++
        this.queue[this.front] = value
        return true
    }
    if this.IsFull() {
        return false
    }
    if this.front == 0 {
        this.front = this.size
    }
    this.front--
    this.queue[this.front] = value
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
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


func (this *MyCircularDeque) DeleteFront() bool {
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


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
        return true
    }
    if this.rear == 0 {
        this.rear = this.size
    }
    this.rear--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.front]
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.rear]
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.front == -1
}


func (this *MyCircularDeque) IsFull() bool {
    return (this.rear + 1) % this.size == this.front
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
```

