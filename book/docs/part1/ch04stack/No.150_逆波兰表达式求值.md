## 150 逆波兰表达式求值-中等

题目：

给你一个字符串数组 `tokens` ，表示一个根据 [逆波兰表示法](https://baike.baidu.com/item/逆波兰式/128437) 表示的算术表达式。

请你计算该表达式。返回一个表示表达式值的整数。

**注意：**

- 有效的算符为 `'+'`、`'-'`、`'*'` 和 `'/'` 。
- 每个操作数（运算对象）都可以是一个整数或者另一个表达式。
- 两个整数之间的除法总是 **向零截断** 。
- 表达式中不含除零运算。
- 输入是一个根据逆波兰表示法表示的算术表达式。
- 答案及所有中间计算结果可以用 **32 位** 整数表示。



> **示例 1：**
>
> ```
> 输入：tokens = ["2","1","+","3","*"]
> 输出：9
> 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
> ```
>
> **示例 2：**
>
> ```
> 输入：tokens = ["4","13","5","/","+"]
> 输出：6
> 解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
> ```
>
> **示例 3：**
>
> ```
> 输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
> 输出：22
> 解释：该算式转化为常见的中缀算术表达式为：
>   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
> = ((10 * (6 / (12 * -11))) + 17) + 5
> = ((10 * (6 / -132)) + 17) + 5
> = ((10 * 0) + 17) + 5
> = (0 + 17) + 5
> = 17 + 5
> = 22
> ```



分析：

栈的基本操作。

遇到运算符，出栈计算，并把新结果入栈；

遇到数字，直接入栈。

```go
// date 2023/12/18
func evalRPN(tokens []string) int {
    stack := make([]string, 0, 16)
    n := len(tokens)
    for i := 0; i < n; i++ {
        char := tokens[i]
        if char == "+" || char == "-" || char == "*" || char == "/" {
            res := calc(stack[len(stack)-2], stack[len(stack)-1], char)
            stack = stack[:len(stack)-2]
            stack = append(stack, fmt.Sprintf("%d", res))
        } else {
            stack = append(stack, char)
        }
    }
    if len(stack) > 0 {
        v, _ := strconv.Atoi(stack[len(stack)-1])
        return v
    }
    return 0
}

func calc(x, y string, k string) int {
    v1, _ := strconv.Atoi(x)
    v2, _ := strconv.Atoi(y)
    switch k {
    case "+":
        return v1 + v2
    case "-":
        return v1 - v2
    case "*":
        return v1 * v2
    case "/":
        return v1 / v2   
    }
    return 0
}
```

