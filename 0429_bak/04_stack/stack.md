## 栈

[TOC]

### 栈的基本操作

  + push(x)：将元素x加入到栈中
  + pop()：弹出栈顶元素，但是，**不会返回弹出元素的值**
  + top()：取出栈顶元素，但是，**不会弹出栈顶元素**
  + empty()：判断栈是否为空，当栈为空时，返回true
  + size()：返回栈中的元素个数

### 栈的实现

从时间复杂度角度而言，push(), pop(), top()，empty()操作只需要O(1)的时间。从实现角度而言，栈的的实现方式有两种，**数组和线性表**。

### C++

```cpp
#define MAX 1000
class Stack
{
private:
  int top;
public:
  int a[MAX];
  Stack() {top = -1;}
  bool push(int x);
  int pop();
  bool empty();
};
// 入栈
bool Stack::push(int x) {
  if(top >= MAX) {
    cout << "Stack Overflow";
    return false;
  } else {
    a[++top] = x;
    return true;
  }
}
// 出栈，一般情况下不带返回，直接删除元素；而自己写的函数可以允许删除的时候返回元素的值
int Stack::pop() {
  if(top < 0) {
    cout << "Stack Underflow";
    return 0;
  } else {
    int x = a[top--];
    return x;
  }
}
// 空判断
bool Stack::empty() {
  return (top < 0)
}
```

### Python

```python
class Stack:
    """模拟栈的基本操作"""
    # 初始化
    def __init__(self):
        self.item = []
    # 判空
    def isEmpty(self):
        return len(self.item) == 0
    # 压栈
    def push(self, data):
        self.item.append(data)
    # 出栈
    def pop(self):
        if not self.isEmpty():
            return self.item.pop()
    # 去栈顶元素
    def top(self):
        if not self.isEmpty():
            return self.item[len(self.item)-1]
    # 栈的大小
    def size(self):
        return len(self.item)
```

### 栈的应用

栈的主要应用包括：运算符的优先级，字符串反转。

### 由队列实现栈

```cpp
#include <iostream>
#include <queue>
using namespace std;
class MyStack {
private:
  queue<int> q;
public:
  void Push(int x);
  void Pop();
  int Top();
  bool Empty();
  int Size();
};
// 压栈
void MyStack::Push(int x) {
  q.push(x);
}
// 出栈
void MyStack::Pop() {
  if(q.size() == 0) {
    cout << "The Stack is Empty" << endl;
    return;
  }
  if(q.size() == 1)
    q.pop();
  else {
    int n = q.size();
    while(n > 1) {
      q.push(q.front());
      q.pop();
      n--;
    }
    q.pop();
  }
}
// 获取栈顶元素
int MyStack::Top() {
  if(q.size() == 0) {
    cout << "The Stack is Empty" << endl;
    return -1;
  }
  if(q.size() == 1)
    return q.front();
  else {
    int n = q.size();
    while(n > 1) {
      q.push(q.front());
      q.pop();
      n--;
    }
    return q.front();
  }
}
// 判断栈是否为空
bool MyStack::Empty() {
  return q.empty();
}
// 获取栈中元素的个数
int MyStack::Size() {
  return q.size();
}
```

### 栈的反转

利用迭代思想，实现栈的反转，只能使用push(),pop(),empty(),size()操作。

```
栈：                  栈：
4  ->  top            1  ->  top
3            反转后    2
2                     3
1                     4
```

分析：栈的标准push()是向栈顶压入新元素，如果将push()改为向栈底压入新元素，那么就可以利用empty(),pop()来递归实现栈的反转。

```cpp
stack<int> st;
void push_bottom(int x) {
  if(st.size() == 0)
    st.push(x);
  else {
    int temp = st.top();
    st.pop();
    push_bottom(x);
    st.push(temp);
  }
}
void Reverse() {
  if(st.size() > 0) {
    int x = st.top();
    st.pop();
    Reverse();
    push_bottom(x);
  }
}
```

* Python

```Python
def insertAtBottom(stack, item):
    if isEmpty(stack):
        push(stack, item)
    else:
        temp = pop(stack)
        insertAtBottom(stack, item)
        push(stack, temp)

# Below is the function that reverses the given stack
# using insertAtBottom()
def reverse(stack):
    if not isEmpty(stack):
        temp = pop(stack)
        reverse(stack)
        insertAtBottom(stack, temp)
```

### 栈的排序

注意：利用递归思想，只能使用栈的基本操作。

### 算法分析

```
// 排序
sortStack(stack S)
	if stack is not empty:
		temp = pop(S);  
		sortStack(S);
		sortedInsert(S, temp);
// 向已排序的栈中插入元素
sortedInsert(Stack S, element)
	if stack is empty OR element > top element
		push(S, elem)
	else
		temp = pop(S)
		sortedInsert(S, element)
		push(S, temp)
```

### 算法实现

* C++

```cpp
// 递归实现向已排序栈中插入元素
void SortedInsert(stack<int> &s, int x)
{
  if(s.empty() || x > s.top()) {
    s.push(x);
    return;
  }
  int temp = s.top();
  s.pop();
  SortedInsert(s, x);
  s.push(temp);
}
// 排序
void StackSort(stack<int> &s) {
  if(!s.empty()) {
    int temp = s.top();
    s.pop();
    StackSort(s);
    SortedInsert(s, temp);
  }
}
```

* Python

```Python
def SortedInsert(s, x):
    if s.isEmpty() or x > s.top():
        s.push(x)
        return;
    else:
        temp = s.top()
        s.pop()
        SortedInsert(s, x)
        s.push(temp)

def StackSort(s):
    if not s.isEmpty():
        temp = s.top()
        s.pop()
        StackSort(s)
        SortedInsert(s, temp)
```

### SpecialStack

实现一种特殊的栈结构，让其push(),pop(),empty(),getmin()的四种操作的时间复杂度为O(1)。

```
// 继承上面栈的实现部分，Stack类
class SpecialStack: public Stack
{
private:
  Stack min;
public:
  int pop();
  void push(int x);
  int getmin();
};
// 入栈
void SpecialStack::push(int x) {
  if(empty()) {
    Stack::push(x);
    min.push(x);
  } else {
    Stack::push(x);
    int y = min.pop();
    min.push(y);
    if(x < y)
      min.push(x);
    else
      min.push(y);
  }
}
// 出栈
int SpecialStack::pop() {
  int x = Stack::pop();
  min.pop();
  return x;
}
// getmin
int SpecialStack::getmin() {
  int x = min.pop();
  min.push(x);
  return x;
}
```

### 相关题目

#### 二叉搜索树的前序序列判断

问题描述

给定一个数组，判断其是否为一个二叉搜索树的前序遍历，并且要求时间复杂度为O(n)。

提示：左节点的值 小于 根节点的值 小于或等于 右节点的值，满足二叉搜索树的特征。

算法分析

这个问题与NGE问题类似，可以采用栈进行求解。算法如下：

```
1. 创建一个空栈
2. 初始化root = INT_MIN
3. 遍历每一个节点pre[i]
  a) If pre[i] is smaller than current root, return false.
  b) Keep removing elements from stack while pre[i] is greater than stack top.
     Make the last removed item as new root.
     At this point, pre[i] is greater than the removed root.
  c) push pre[i] to stack (All elements in stack are in decreasing order.
```

算法实现

```cpp
bool BSTCheck(int pre[], int n) {
  stack<int> s;
  int root = INT_MIN;
  for(int i = 0; i < n; i++) {
    if(pre[i] < root)
      return false;
    while(!s.empty() && s.top() < pre[i]) {
      // 进入到while(),则说明存在右子树
      root = s.top();
      s.pop();
    }
    // 压入左子树的根节点
    s.push(pre[i]);
  }
  return true;
}
```

#### Next Greater Element

问题定义

给定一个整型数组array[n]，输出数组中每个元素对应的下一个比其大元素(NGE)。其规则如下：

+ 对于最后一个元素，其NGE = -1;
+ 对于一个非递增数组，所有元素的NGE均为-1

例如，数组{4, 5, 2, 25}对应的NGE数组为{5, 25, 25, -1}。

算法分析

方法一：两层循环，外层循环负责检查每个元素，内层循环负责选取后面元素中第一个比其大的元素。

方法二：利用stack。

1） 创建两个栈s和s2，s，s2均存储存储当前元素至末尾元素的的元素序列，s用于判决比较，s2用于覆盖s。
2）当s不为空时，若s栈顶元素大于当前元素，则NGE = s.top()，否则，s.pop()后继续判断。
3）若s为空时，仍没有找到大于当前元素的值，则NGE = -1。

算法实现

```cpp
void NextGreaterElement(int arr[], int n) {
  if(n < 1)
    return;
  int nge[n];
  stack<int> s, s2;
  nge[n - 1] = -1;
  s.push(arr[n - 1]);
  s2.push(arr[n - 1]);
  int i, j;
  for(i = n - 2; i >= 0; i--) {
    s2.push(arr[i]);
    while(!s.empty()) {
      int x = s.top();
      if(arr[i] < x) {
        nge[i] = x;
        break;
      } else {
        s.pop();
      }
    }
    if(s.empty()) {
      nge[i] = -1;
    }
    s = s2;
  }
  for(i = 0; i < n; i++)
    cout << arr[i] << "--" << nge[i] << endl;;
}
```

#### Middle Stack

问题定义

设计并实现一种特殊栈结构，具有以下四种操作，且四种操作的时间复杂度均为O(1)。

1) push() which adds an element to the top of stack.
2) pop() which removes an element from top of stack.
3) findMiddle() which will return middle element of the stack.
4) deleteMiddle() which will delete the middle element.
其中push()和pop()操作属于栈的标准操作。

算法分析

栈的底层实现是利用数组或线性表来进行构建的。对于push(),pop()和findMiddle()操作，数组可满足O(1)时间复杂度，但是deleteMiddle()设计元素的移动，数组无法满足。同样，单链表也无法满足两个方向(即，push()和pop())上的中间元素的移动。

因此，我们利用双向链表来实现这一特殊的栈结构。

算法实现

```cpp
// 双链表
struct Node {
  struct Node *pre, *next;
  int val;
};
// 栈结构
struct MyStack {
  struct Node *head;
  struct Node *mid;
  int count;
};
// 初始化
struct MyStack *CreateStack() {
  MyStack *ms = (MyStack*)malloc(sizeof(struct MyStack));
  ms->count = 0;
  return ms;
}
// 入栈
void Push(struct MyStack *ms, int x) {
  struct Node *data = (struct Node*)malloc(sizeof(struct Node));
  data->val = x;
  data->pre = NULL;
  data->next = ms->head;
  ms->count += 1;
  // 第一个元素
  if(ms->count == 1) {
    ms->mid = data;
  } else {
    // 让栈顶的元素指向新元素
    ms->head->pre = data;
    // 当有奇数个元素时，mid指针才发生一次移动
    if(ms->count & 1)
      ms->mid = ms->mid->pre;
  }
  ms->head = data;
}
// 出栈
int Pop(struct MyStack *ms) {
  if(ms->count == 0) {
    cout << "Stack is Empty" << endl;
    return -1;
  }
  struct Node *temp = ms->head;
  int x = temp->val;
  ms->head = temp->next;
  if(ms->head != NULL)
    ms->head->pre = NULL;
  ms->count -= 1;
  // 元素个数由奇数变为偶数时，mid指针需要移动
  if(!(ms->count & 1))
    ms->mid = ms->mid->next;
  free(temp);
  return x;
}
// 取”栈中“元素
int FindMiddle(struct MyStack *ms) {
  if(ms->count == 0) {
    cout << "Stack is Empty" << endl;
    return -1;
  } else {
    return ms->mid->val;
  }
}
// 删除”栈中“元素
int DeleteMiddle(struct MyStack *ms) {
  if(ms->count == 0) {
    cout << "Stack is Empty" << endl;
    return -1;
  }
  struct Node *temp = ms->mid;
  int x = temp->val;
  ms->count -= 1;
  // 元素个数由偶数变为奇数，mid指针向pre方向移动
  if((ms->count) & 1) {
    ms->mid = temp->pre;
    if(ms->count == 1) {
      ms->mid->next = NULL;
    } else {
    temp->next->pre = ms->mid;
    ms->mid->next = temp->next;
    }
  } else {
  // 元素个数由奇数变为偶数，mid指针向next方向移动
  if(ms->count == 0) {
    ms->head = NULL;
    ms->mid = NULL;
  } else {
    ms->mid = temp->next;
    temp->pre->next = ms->mid;
    ms->mid->pre = temp->pre;
    }
  }
  free(temp);
  return x;
}
```

#### 队列与栈的有限操作间的转换

题目一：已知Stack类及其3个方法Push()、Pop()和 Count()，请用2个Stack实现Queue类的入队(Enqueue)、出队(Dequeue)方法以及判空(Empty)操作。*注意，这里假定Pop()操作可以获得元素的值*。

题目主要考察对基本数据结构的理解，以及操作的转换。

面试公司：VMware

**提示**：

两个栈A和B，栈A提供入队操作，栈B提供出队操作，算法如下：

1）如果栈B不为空，直接弹出栈B的元素；
2）如果栈B为空，则依次弹出栈A的元素，放入栈B中，再弹出栈B的元素.

```cpp
// Using Stack implement the Queue's Operations
// Stack's Operations: s.push(x) s.pop() s.size()
// Queue's Operations: q.enqueue(x) q.dequeue() q.empty()
class MyQueue {
private:
  stack<int> s;   // 用于入队
  stack<int> s2;  // 用于出队
public:
  void enqueue(int x);
  void dequeue(int &x);
  bool empty();
};
// 入队
void MyQueue::enqueue(int x) {
  s.push(x);
}
// 出队
void MyQueue::dequeue(int &x) {
  int temp;
  if(s2.size() == 0) {
    if(s.size() == 0) {
      cout << "The Queue is Empty" << endl;
    } else {
      while(s.size() > 0) {
        s.pop(&temp);
        s2.push(temp);
      }
    }
  }
  if(s2.size() != 0)
    s2.pop(&x);
}
// 判断队列是否为空
bool MyQueue::empty() {
  return ((s.size() == 0) && (s2.size() == 0));
}
```

#### 由队列实现栈

Queue的操作包括enqueue(),dequeue()，Stack的操作包括push(),pop()。

方法一

使用两个queue，即q1,q2。每个新加入的元素，均位于q1的队头，因此，pop()操作，直接dequeue即可。q2用于保存新加入的元素。这个情况下，push()操作具有较高的复杂度。

算法分析

```
push(s, x) // x is the element to be pushed and s is stack
  1) Enqueue x to q2
  2) One by one dequeue everything from q1 and enqueue to q2.
  3) Swap the names of q1 and q2

pop(s)
  1) Dequeue an item from q1 and return it.
```

方法二

这个方法与方法一类似，只是pop()操作的复杂度相对高。

算法分析

```
push(s,  x)
  1) Enqueue x to q1 (assuming size of q1 is unlimited).

pop(s)  
  1) One by one dequeue everything except the last element from q1 and enqueue to q2.
  2) Dequeue the last item of q1, the dequeued item is result, store it.
  3) Swap the names of q1 and q2
  4) Return the item stored in step 2.
```

方法三

只使用一个Queue，即每次push()的时候将元素放在队列的队头。

算法实现

```cpp
class MyStack
{
  queue<int> q;
public:
  void push(int x);
  void pop();
  int top();
  bool empty();
};
// 压栈
void MyStack::push(int x) {
  int s = q.size();
  q.push(x);
  // 将之前的元素依次加入到队列的尾部
  for(int i = 0; i < s; i++) {
    q.push(q.front());
    q.pop();
  }
}
// 出栈
void MyStack::pop() {
  if(q.empty())
    cout << "Stack is Underflow";
  else
    q.pop();
}
// 栈顶元素
int MyStack::top() {
  if(q.empty()) {
    cout << "Stack is Underflow";
    return -1;
  } else {
    return q.front();
  }
}
// 判空
bool MyStack::empty() {
  return q.empty();
}
```

#### twoStack

利用数组Array实现一种双栈数据结构twoStack，towStack满足以下功能：

+ push1(int x)：向栈1压入元素x
+ push2(int x)：向栈2压入元素x
+ pop1()：删除栈1中的栈顶元素，并返回删除的元素值
+ pop2()：删除栈2中的栈顶元素，并返回删除的元素值

注意：实现twoStack的过程中尽量保证空间效率。

分析：

第一种方法，将数组分为两部分，array[0]~array[n/2]用于栈1，array[n/2+1]~array[n-1]用于栈2。但是这个情况下存在的问题是，当其中的一个栈满时，而另一个栈可能还有可用空间，造成空间上的浪费。

第二种方法，分别将数组array[0]和array[n-1]作为两个栈的栈底，从数组两端向中间开始存储数据，当index1(栈1)和index2(栈2)相等时，栈满。

下面基本第二种方法进行实现：

* C++

```cpp
class twoStack()
{
private:
  int *arr;
  int size;
  int top1, top2;
public:
  twoStack(int n);
  bool push1(int x);
  bool push2(int x);
  int pop1();
  int pop2();
}
// 初始化
void twoStack::twoStack(int n) {
  size = n;
  arr = new int[n];
  top1 = -1;
  top2 = size;
}
// 栈1-入栈
bool twoStack::push1(int x) {
  if(top1 < top2 - 1) {
    arr[++top1] = x;
    return true;
  } else {
    cout << "Stack is Overflow";
    exit(1);
  }
}
// 栈2-入栈
bool twoStack::push2(int x) {
  if(top1 < top2 - 1) {
    arr[--top2] = x;
    return true;
  } else {
    cout << "Stack is Overflow";
    exit(1);
  }
}
// 栈1-出栈
int twoStack::pop1() {
  if(top1 >= 0) {
    int x = arr[top1];
    top1--;
    return x;
  } else {
    cout << "Stack1 is Empty";
    exit(1);
  }
}
// 栈2-出栈
int twoStack::pop2() {
  if(top2 <= size) {
    int x = arr[top2];
    top2++;
    return x;
  } else {
    cout << "Stack2 is Empty";
    exit(1);
  }
}
```

* Python

```Python
class twoStack:

    def __init__(self, n):
        self.size = n;
        self.arr = [None] * n;
        self.top1 = -1;
        self.top2 = self.size

    def push1(self, x):
        if self.top1 < self.top2 - 1:
            self.top1 = self.top1 + 1
            self.arr[self.top1] = x
        else:
            print("Stack is Overflow")
            exit(1)

    def push2(self, x):
        if self.top1 < self.top2 - 1:
            self.top2 = self.top2 - 1
            self.arr[self.top2] = x
        else:
            print("Stack is Overflow")
            exit(1)

    def pop1(self):
        if self.top1 >= 0:
            x = self.arr[self.top1]
            self.top1 = self.top1 - 1
            return x
        else:
            print("Stack is Underflow")
            exit(1)

    def pop2(self):
        if self.top2 <= self.size:
            x = self.arr[self.top2]
            self.top2 = self.top2 - 1
            return x
        else:
            print("Stack is Underflow")
            exit(1)
```
