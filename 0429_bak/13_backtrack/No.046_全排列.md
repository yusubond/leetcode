## 46 全排列-中等

题目：

给定一个不含重复数字的数组 `nums` ，返回其 *所有可能的全排列* 。你可以 **按任意顺序** 返回答案。

 

> **示例 1：**
>
> ```
> 输入：nums = [1,2,3]
> 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0,1]
> 输出：[[0,1],[1,0]]
> ```
>
> **示例 3：**
>
> ```
> 输入：nums = [1]
> 输出：[[1]]
> ```



分析：

经典的回溯算法。

这一版的回溯算法条件是：

1. 结束条件：选择列表为空
2. 遍历的时候，把已经遍历过的元素，不在添加到选择列表中



```go
// date 2023/12/25
func permute(nums []int) [][]int {
    res := make([][]int, 0, 16)

    var backtrack func(nums []int, temp []int)

    backtrack = func(nums []int, temp []int) {
        if len(nums) == 0 {
            res = append(res, temp)
            return
        }
        for i, v := range nums {
            newTemp := append(temp, v)
            unUsed := make([]int, 0, 16)
            unUsed = append(unUsed, nums[:i]...)
            unUsed = append(unUsed, nums[i+1:]...)
            backtrack(unUsed, newTemp)
        }
    }

    backtrack(nums, []int{})

    return res
}

// 解法2
// 也是回溯算法，通过标记是否已经遍历过实现去重
func permute(nums []int) [][]int {
    ans := make([][]int, 0, 16)
    n := len(nums)
    visited := make([]bool, n)

    var backtrack func(path []int)
    backtrack = func(path []int) {
        if len(path) == n {
            one := make([]int, n)
            copy(one, path)
            ans = append(ans, one)
            return
        }
        for i, visit := range visited {
            if !visit {
                path = append(path, nums[i])
                visited[i] = true

                backtrack(path)
                
                visited[i] = false
                path = path[:len(path)-1]
            }
        }
    }

    backtrack([]int{})

    return ans
}
```



