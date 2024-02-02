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

数据栈 dataStack 保存所有的元素。

小元素栈 minStack 保存当前的最小元素。

入栈的时候，如果 minStack 的栈顶元素 大于等于入栈元素，那么把入栈元素 val 也放入 minStack 中，这样做是为了保证两个连续相同的元素都可以入小元素栈。

出栈的时候，如果 minStack 的栈顶元素 等于 dataStack 的栈顶元素，那么 minStack 也要出栈。

```go
// date 2020/03/23
// 第一种方案：辅助栈stackMin只保留当前栈中的最小值，即辅助栈和数据栈stackData长度不相等
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
    }
    this.stackData = append(this.stackData, x)
}


func (this *MinStack) Pop()  {
    if len(this.stackData) <= 0 {return}
    if len(this.stackMin) > 0 && this.stackMin[len(this.stackMin)-1] == this.stackData[len(this.stackData)-1] {
        this.stackMin = this.stackMin[:len(this.stackMin)-1]
    }
    this.stackData = this.stackData[:len(this.stackData)-1]
}


func (this *MinStack) Top() int {
    return this.stackData[len(this.stackData)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.stackMin) <= 0 { return 0 }
    return this.stackMin[len(this.stackMin)-1]
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

