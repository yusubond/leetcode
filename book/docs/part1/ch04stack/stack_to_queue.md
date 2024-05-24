# 队列与栈的有限操作间的转换

题目：

已知Stack类及其3个方法Push()、Pop()和 Count()，请用2个Stack实现Queue类的入队(Enqueue)、出队(Dequeue)方法以及判空(Empty)操作。

*注意，这里假定Pop()操作可以获得元素的值*。

这个题目主要考察对基本数据结构的理解，以及操作的转换。



思路分析：

两个栈A和B，栈A提供入队操作，栈B提供出队操作，算法如下：

1）如果栈B不为空，直接弹出栈B的元素；
2）如果栈B为空，则依次弹出栈A的元素，放入栈B中，再弹出栈B的元素.

```go
// Using Stack implement the Queue's Operations
// Stack's Operations: s.push(x) s.pop() s.size()
// Queue's Operations: q.enqueue(x) q.dequeue() q.empty()
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
```



栈的三个操作实现如下：

```go
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
```

