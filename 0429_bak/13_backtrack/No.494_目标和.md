## 494 目标和-中等

题目：

给你一个非负整数数组 `nums` 和一个整数 `target` 。

向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：

- 例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

 

> **示例 1：**
>
> ```
> 输入：nums = [1,1,1,1,1], target = 3
> 输出：5
> 解释：一共有 5 种方法让最终目标和为 3 。
> -1 + 1 + 1 + 1 + 1 = 3
> +1 - 1 + 1 + 1 + 1 = 3
> +1 + 1 - 1 + 1 + 1 = 3
> +1 + 1 + 1 - 1 + 1 = 3
> +1 + 1 + 1 + 1 - 1 = 3
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1], target = 1
> 输出：1
> ```



分析：

这道题目可以用 DFS 解决。既然可以在每个数字的前面添加`+`或者`-`，那么就有如下的搜索公式：

```sh
// 定义 dfs(0, target) 表示从下标0开始搜索，目标和为 target 的解的个数，那么
dfs(0, target) = dfs(1, target-nums[0]) + dfs(1, target+nums[0])
```



```go
// date 2023/12/27
func findTargetSumWays(nums []int, target int) int {
    res := 0
    n := len(nums)

    var backtrack func(idx, target int)
    backtrack = func(idx, target int) {
        if idx == n {
            if target == 0 {
                res += 1
            }
            return
        }
        backtrack(idx+1, target-nums[idx])  // +nums[i]
        backtrack(idx+1, target+nums[idx])  // -nums[i]
    }

    backtrack(0, target)

    return res
}
```

