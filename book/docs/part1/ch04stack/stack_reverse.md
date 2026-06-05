---
title: "第 4.2 节 栈的反转"
---



# 4.2 栈的反转

利用迭代思想，实现栈的反转，只能使用push(),pop(),empty(),size()操作。

```
栈：                  栈：
4  <-  top            1  <-  top
3            反转后    2
2                     3
1                     4
```



思路分析：

栈的标准`push`是向栈顶压入新元素，如果将push()改为向栈底压入新元素，那么就可以利用empty(),pop()来递归实现栈的反转。

```go
// pushBottom define
func (s *FixedSizeStack) pushBottom(x int) {
	if s.IsEmpty() {
		s.Push(x)
	} else {
		temp := s.Top()
		s.Pop()
		s.pushBottom(x)
		s.Push(temp)
	}
}

// Reverse define
func (s *FixedSizeStack) Reverse() {
	if !s.IsEmpty() {
		x := s.Top()
		s.Pop()
		s.Reverse()
		s.pushBottom(x)
	}
}
```

