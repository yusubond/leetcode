## 739 每日温度-中等

题目：

给定一个整数数组 `temperatures` ，表示每天的温度，返回一个数组 `answer` ，其中 `answer[i]` 是指对于第 `i` 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 `0` 来代替。



> **示例 1:**
>
> ```
> 输入: temperatures = [73,74,75,71,69,72,76,73]
> 输出: [1,1,4,2,1,1,0,0]
> ```
>
> **示例 2:**
>
> ```
> 输入: temperatures = [30,40,50,60]
> 输出: [1,1,1,0]
> ```
>
> **示例 3:**
>
> ```
> 输入: temperatures = [30,60,90]
> 输出: [1,1,0]
> ```



**解题思路**

answer 数组表示下一个更高的温度出现在几天后，一种朴素的想法是从当前开始往后遍历，找第一个大于该元素的值。那么，也可以这样想，从当前元素开始，往前找，找到比它小的就更新。

用栈来维护元素的下标，出栈的条件是当前温度 cur 大于栈中下标所对应的温度，那么 cur 就可以用来更新 answer数组。

入栈，依次入栈就可以。

```go
// date 2023/12/21
func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, 0, 16)
    n := len(temperatures)
    ans := make([]int, n, n)

    for i := 0; i < n; i++ {
        v := temperatures[i]

        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            preIdx := stack[len(stack)-1]
            ans[preIdx] = i - preIdx
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
        // 没有出栈之前，一直入栈，那么则说明 v 一直小于栈顶元素
        // 栈是单调递减的，对应的数组也是单调递减的
        // 栈维护的是温度单调递减
        // 只要出现一个v比栈顶大的，那么栈里面比 v 小的都可以更新掉
    }

    return ans
}
```

