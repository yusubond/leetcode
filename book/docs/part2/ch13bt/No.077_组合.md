## 77 组合-中等

题目：

给定两个整数 `n` 和 `k`，返回范围 `[1, n]` 中所有可能的 `k` 个数的组合。

你可以按 **任何顺序** 返回答案。



> **示例 1：**
>
> ```
> 输入：n = 4, k = 2
> 输出：
> [
>   [2,4],
>   [3,4],
>   [2,3],
>   [1,2],
>   [1,3],
>   [1,4],
> ]
> ```
>
> **示例 2：**
>
> ```
> 输入：n = 1, k = 1
> 输出：[[1]]
> ```



分析：

这道题也是经典的回溯模板。注意，元素不可重复。

```go
// date 2023/12/26
func combine(n int, k int) [][]int {
    res := make([][]int, 0, 16)

    var backtrack func(n, k int, temp []int)
    backtrack = func(idx, k int, temp []int) {
        if k == 0 {
            one := make([]int, len(temp))
            copy(one, temp)
            res = append(res, one)
            return
        }
        // [1, n] 都是候选
        for i := idx; i <= n; i++ {
            temp = append(temp, i)
            // 不可重复，所以从 i+1 开始
            backtrack(i+1, k-1, temp)
            temp = temp[:len(temp)-1]
        }
    }

    backtrack(1, k, []int{})

    return res
}
```

