## Queue

[TOC] 

### 基本定义

队列实现的是一种先进先出（FIFO）策略的线性表。

队列有队头(head)和队尾(tail)，当有一个元素入队时，放入队尾；出队时，即删除队头元素。

### 队列的基本操作

  + push(x)：将元素x加入到队列的尾部
  + pop()：弹出队列的第一个元素，但是，**不会返回被弹元素的值**
  + front()：返回队首元素
  + back()：返回队尾元素
  + empty()：判断队列是否为空，当队列为空时，返回true
  + size()：返回队列中的元素个数

注意：以上操作均基于c++模板类，位于头文件<queue>和<stack>中。以下是两个数据结构互相转换算法和程序。

### 队列的实现

从时间复杂度角度而言，push(), pop(), front()，back(), empty()操作只需要O(1)的时间。从实现角度而言，队列的实现方式有两种，**数组和线性表**。

```cpp
class MyQueue
{
private:
  int front, rear, size;
  unsigned capacity;
  int *array;
public:
  MyQueue(unsigned num);
  bool Push(int x);
  bool Pop();
  int Front();
  int Back();
  bool Empty();
  bool Full();
};
// 初始化
MyQueue::MyQueue(unsigned num) {
  capacity = num;
  array = new int[capacity];
  front = 0;
  rear = capacity - 1;
  size = 0;
}
bool MyQueue::Full() {
  if(size == capacity)
    return true;
  else
    return false;
}
bool MyQueue::Empty() {
  if(size == 0)
    return true;
  else
    return false;
}
// 入队
bool MyQueue::Push(int x) {
  if(size == capacity) {
    cout << "Queue is Overflow" << endl;
    return false;
  }
  rear = (rear + 1) % capacity;
  array[rear] = x;
  size += 1;
  return true;
}
// 出队
bool MyQueue::Pop() {
  if(size == 0) {
    cout << "Queue is Underflow" << endl;
    return false;
  }
  int x = array[front];
  front = (front + 1) % capacity;
  size -= 1;
  return x;
}
// 取队头元素
int MyQueue::Front() {
  if(size == 0) {
    cout << "Queue is Underflow" << endl;
    return -1;
  }
  return array[front];
}
// 取队尾元素
int MyQueue::Back() {
  if(size == 0) {
    cout << "Queue is Underflow" << endl;
    return -1;
  }
  return array[rear];
}
```

### 由栈实现队列

```cpp
#include <iostream>
#include <stack>
using namespace std;
class MyQueue {
private:
  stack<int> s;   // 用于入队
  stack<int> s2;  // 用于出队
public:
  void Push(int x);
  void Pop();
  int Front();
  int Back();
  bool Empty();
  int Size();
};
// 入队
void MyQueue::Push(int x) {
  s.push(x);
}
// 出队
void MyQueue::Pop() {
  if(s2.empty()) {
    if(s.empty()) {
      // s2为空，s为空，表示整个队列为空
      cout << "The Queue is Empty" << endl;
      return;
    } else {
      // s2为空，s不为空，则将s中的元素依次压入s2中
      while(!s.empty()) {
        s2.push(s.top());
        s.pop();
      }
    }
  }
  if(!s2.empty()) {
    // s2不为空，其栈顶元素就是最先进入的元素，即队首元素
    s2.pop();
  }
}
// 获取队首元素
int MyQueue::Front() {
  // 与出队原理一样，只是将pop()变为top()
  if(s2.empty()) {
    if(s.empty()) {
      cout << "The Queue is Empty" << endl;
      return -1;
    } else {
      while(!s.empty()) {
        s2.push(s.top());
        s.pop();
      }
    }
  }
  if(!s2.empty()) {
    return s2.top();
  }
}
// 获取队尾元素
int MyQueue::Back() {
  // 因为每次入队均将元素压入s中，因此s的栈顶元素即为队尾元素
  // 同样的道理，当s为空使，s2不为空时，s2中的栈顶->栈底 对应 队首->队尾
  // 因此，需要将s2中的元素依次压入s中
  if(s.empty()) {
    if(s2.empty()) {
      // s为空，s2为空，即整个队列为空
      cout << "The Queue is Empty" << endl;
      return -1;
    } else {
      // s为空，s2不为空，则将s2中的元素依次压入s中
      while(!s2.empty()) {
        s.push(s2.top());
        s2.pop();
      }
    }
  }
  if(!s.empty()) {
    // s不为空，其栈顶元素就是队列的队尾元素
    return s.top();
  }
}
// 判断队列是否为空
bool MyQueue::Empty() {
  return (s.empty() && s2.empty());
}
// 获取栈中元素的个数
int MyQueue::Size() {
  return (s.size() + s2.size());
}
```

### 相关题目

[TOC]

#### 621 Task Scheduler【任务调度器】

题目描述：

给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。

示例1：

```
输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
```

```go
// 算法1：优先安排出现次数最多的任务
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    var res int
    for m[25] > 0 {
        i := 0
        for i <= n {
            if m[25] == 0 {
                break
            }
            if i < 16 && m[25-i] > 0 {
                m[25-i]--
            }
            res++
            i++
        }
        sort.Slice(m, func(i, j int) bool {
            return m[i] < m[j]
        })
    }
    return res
}
// 算法二：设计，利用空闲时间[leetcode-cn]
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    max_val := m[25] - 1
    idle := n * max_val
    for i := 24; i >= 0 && m[i] > 0; i-- {
        if m[i] <= max_val {
            idle -= m[i]
        } else {
            idle -= max_val
        }
    }
    if idle > 0 {
        return idle + len(tasks)
    }
    return len(tasks)
}
```
