## 20 有效的括号

题目：

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。



**解题思路**


stack数据结构，压栈及出栈

算法如下：时间复杂度O(n) 空间复杂度O(1)

```go
// date 2020/01/06
/* 算法：
 0. 如果字符串长度为单数，直接返回false.
 1. 遇到左括号，入栈
 2. 遇到右括号，判断栈是否为空，如果为空，返回false
 3. 如果不为空，出栈
 4. 遍历结束后，判断栈是否为空。
*/
func isValid(s string) bool {
  if len(s) & 0x1 == 1 {return false}
  stack := make([]rune, 0)
  var c rune
  for _, v := range s {
    if v == '(' || v == '{' || v == '[' {
      stack = append(stack, v)
    } else {
      if len(stack) == 0 {return false}
      c = stack[len(stack)-1]
      if v == ')' && c == '(' || v == '}' && c == '{' || v == ']' && c == '[' {
        stack = stack[:len(stack)-1]
      } else {
        return false
      }
    }
  }
  return len(stack) == 0
}
```

