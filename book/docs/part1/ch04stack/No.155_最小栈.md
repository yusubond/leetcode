## 155 最小栈-中等

题目：

设计一个支持 `push` ，`pop` ，`top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类:

- `MinStack()` 初始化堆栈对象。
- `void push(int val)` 将元素val推入堆栈。
- `void pop()` 删除堆栈顶部的元素。
- `int top()` 获取堆栈顶部的元素。
- `int getMin()` 获取堆栈中的最小元素。



> **示例 1:**
>
> ```
> 输入：
> ["MinStack","push","push","push","getMin","pop","top","getMin"]
> [[],[-2],[0],[-3],[],[],[],[]]
> 
> 输出：
> [null,null,null,null,-3,null,0,-2]
> 
> 解释：
> MinStack minStack = new MinStack();
> minStack.push(-2);
> minStack.push(0);
> minStack.push(-3);
> minStack.getMin();   --> 返回 -3.
> minStack.pop();
> minStack.top();      --> 返回 0.
> minStack.getMin();   --> 返回 -2.
> ```



分析：

因为要求在常数时间内检索到最小元素，所以考虑使用辅助栈。

数据栈 dataStack 保存所有的元素。小元素栈 minStack 保存当前栈中的最小元素。

入栈的时候，如果入栈元素 `x` 小于等于 minStack 的栈顶元素那么把入栈元素 val 也放入 minStack 中。这样做是为了保证两个连续相同的元素都可以入小元素栈。

出栈的时候，如果 minStack 的栈顶元素 等于 dataStack 的栈顶元素，那么 minStack 也要出栈。

```go
// date 2020/03/23
// 第一种方案：辅助栈stackMin只保留当前栈中的最小值，即辅助栈和数据栈stackData长度不相等
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0, 64),
		min:  make([]int, 0, 64),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	s1 := len(this.data)
	if s1 <= 0 {
		return
	}
	s2 := len(this.min)
	if s2 > 0 && this.min[s2-1] == this.data[s1-1] {
		this.min = this.min[:s2-1]
	}

	this.data = this.data[:s1-1]
}

func (this *MinStack) Top() int {
	s1 := len(this.data)
	if s1 <= 0 {
		return 0
	}
	return this.data[s1-1]
}

func (this *MinStack) GetMin() int {
	s2 := len(this.min)
	if s2 <= 0 {
		return 0
	}
	return this.min[s2-1]
}

// 第二种方案：辅助栈stackMin随时保留当前栈的最小值，保持和数据栈stackData长度相等
type MinStack struct {
    stackData []int
    stackMin []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stackData: make([]int, 0),
        stackMin: make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.stackMin) == 0 || this.stackMin[len(this.stackMin)-1] >= x {
        this.stackMin = append(this.stackMin, x)
    } else {
        this.stackMin = append(this.stackMin, this.stackMin[len(this.stackMin)-1])
    }
    this.stackData = append(this.stackData, x)
}


func (this *MinStack) Pop()  {
    if len(this.stackData) <= 0 {return}
    this.stackData = this.stackData[:len(this.stackData)-1]
    this.stackMin = this.stackMin[:len(this.stackMin)-1]
}


func (this *MinStack) Top() int {
    return this.stackData[len(this.stackData)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.stackMin) <= 0 { return 0 }
    return this.stackMin[len(this.stackMin)-1]
}
```

