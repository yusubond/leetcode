## 946 验证栈序列-中等

题目：

给定 `pushed` 和 `popped` 两个序列，每个序列中的 **值都不重复**，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 `true`；否则，返回 `false` 。



> **示例 1：**
>
> ```
> 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
> 输出：true
> 解释：我们可以按以下顺序执行：
> push(1), push(2), push(3), push(4), pop() -> 4,
> push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
> ```
>
> **示例 2：**
>
> ```
> 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
> 输出：false
> 解释：1 不能在 2 之前弹出。
> ```



分析：

思路就是按照入栈队列依次入栈，取栈顶元素看看是不是在出栈序列中，如果在出栈。

如果入栈序列和出栈序列能够契合，那么出栈的指针就会走到最后，否则不会。

```go
// date 2023/12/18
func validateStackSequences(pushed []int, popped []int) bool {
    stack := make([]int, 0, 16)
    j := 0
    n := len(pushed)

    for _, v := range pushed {
        stack = append(stack, v)

        for len(stack) > 0 && j < n && stack[len(stack)-1] == popped[j] {
            stack = stack[:len(stack)-1]
            j++
        }
    }
    return j == n
}
```

