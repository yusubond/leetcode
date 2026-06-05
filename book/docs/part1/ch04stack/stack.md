---
title: "第 4 章 栈"
---

# 第 4 章 栈

## 1.栈

栈是一种 LIFO (Last In First Out) 的数据结构。

根据栈中元素总个数是否固定，可分为固定大小栈和动态大小栈。

固定大小栈：顾名思义，栈有固定大小，不能动态增长和收缩。如果栈已满并尝试向其添加元素，会发生上溢错误；如果栈已空并尝试从中删除元素，会发生下溢错处。

动态大小栈：栈的大小可动态增长和收缩。



## 2.基本操作

  + push(x)：将元素x加入到栈中
  + pop()：弹出栈顶元素，但**不会返回弹出元素的值**
  + top()：取出栈顶元素，但**不会弹出栈顶元素**
  + isEmpty()：判断栈是否为空，当栈为空时，返回true
  + isFull()：判断栈是否已满，当栈已满时，返回 true
  + size()：返回栈中的元素个数



## 3.栈的实现

从时间复杂度角度而言，push(), pop(), top()，empty()操作只需要O(1)的时间。从实现角度而言，栈的的实现方式有两种，**数组和线性表**。

下面实现固定大小的栈。

```go
type (
	FixedSizeStack struct {
		cap int
		val []int
	}
)

func NewFixedSizeStack(n int) *FixedSizeStack {
	return &FixedSizeStack{
		cap: n,
		val: make([]int, 0, n),
	}
}

// Push define
func (s *FixedSizeStack) Push(x int) {
	if !s.IsFull() {
		s.val = append(s.val, x)
	}
}

// Pop define
func (s *FixedSizeStack) Pop() {
	if !s.IsEmpty() {
		s.val = s.val[:len(s.val)-1]
	}
}

// Top define
func (s *FixedSizeStack) Top() int {
	if !s.IsEmpty() {
		return s.val[len(s.val)-1]
	}
	return -1
}

// IsEmpty define
func (s *FixedSizeStack) IsEmpty() bool {
	return len(s.val) == 0
}

// IsFull define
func (s *FixedSizeStack) IsFull() bool {
	return len(s.val) == s.cap
}

// Size define
func (s *FixedSizeStack) Size() int {
	return len(s.val)
}

```



## 4.栈的优点和缺点

栈的优点：

- 简单性。栈是一种简单且易于理解的数据结构，使其适用于广泛的应用程序。
- 效率。压栈和出栈可在`O(1)`内执行，提供了对数据的高效访问。
- 先入后出。栈的先入后出原则在函数调用和表达式求值等场景有很多经典应用。
- 有限的内存使用。栈只需要存储已入栈的元素，内存效率更高。



栈的缺陷：

- 访问受限。只能从顶部访问栈中的元素，因此很难检索或者修改栈中间的元素。
- 溢出的可能性。如果压入栈中的元素多于栈所能容纳的元素，就会发生溢出错误，从而导致数据丢失。
- 不适合随机访问。栈不允许随机访问元素，这使得它不适合需要特定顺序访问元素的应用。
- 容量有限。栈有固定的容量，如果需要存储的元素数量未知，这可能会成为限制。
