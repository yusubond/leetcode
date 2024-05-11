## 216 组合总和3-中等

题目：

找出所有相加之和为 `n` 的 `k` 个数的组合，且满足下列条件：

- 只使用数字1到9
- 每个数字 **最多使用一次** 

返回 *所有可能的有效组合的列表* 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。



> **示例 1:**
>
> ```
> 输入: k = 3, n = 7
> 输出: [[1,2,4]]
> 解释:
> 1 + 2 + 4 = 7
> 没有其他符合的组合了。
> ```
>
> **示例 2:**
>
> ```
> 输入: k = 3, n = 9
> 输出: [[1,2,6], [1,3,5], [2,3,4]]
> 解释:
> 1 + 2 + 6 = 9
> 1 + 3 + 5 = 9
> 2 + 3 + 4 = 9
> 没有其他符合的组合了。
> ```
>
> **示例 3:**
>
> ```
> 输入: k = 4, n = 1
> 输出: []
> 解释: 不存在有效的组合。
> 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
> ```



分析：

这道题也是经典的回溯算法。需要注意的是：

1. 结束条件：既要判断剩余空位为零，也要判断等于目标和
2. 如果空位不足，直接返回
3. 剪枝：如果候选列表中的元素大于目标和，直接退出

```go
// date 2023/12/26
func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0, 16)

    // k 表示还剩余的空位
    var backtrack func(start, k int, target int, temp []int)
    backtrack = func(start, k int, target int, temp []int) {
        if k == 0 && target == 0 {
            one := make([]int, len(temp))
            copy(one, temp)
            res = append(res, one)
        }
        if k <= 0 {
            return
        }
        for i := start; i <= 9; i++ {
            if i > target {
                break
            }
            temp = append(temp, i)
            backtrack(i+1, k-1, target-i, temp)
            temp = temp[:len(temp)-1]
        }
    }

    backtrack(1, k, n, []int{})

    return res
}
```

