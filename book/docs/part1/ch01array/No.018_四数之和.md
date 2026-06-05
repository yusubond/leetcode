## 18 四数之和-中等

题目：

给你一个由 `n` 个整数组成的数组 `nums` ，和一个目标值 `target` 。请你找出并返回满足下述全部条件且**不重复**的四元组 `[nums[a], nums[b], nums[c], nums[d]]` （若两个四元组元素一一对应，则认为两个四元组重复）：

- `0 <= a, b, c, d < n`
- `a`、`b`、`c` 和 `d` **互不相同**
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

你可以按 **任意顺序** 返回答案 。



分析：

这道题属于简单的复杂度。思路比较简单，但要注意各种条件进行剪枝，检查不必要的计算。

去重是指如果当前值和前一个值一样，跳过检查。

1. 整体思路是，先排序，然后两层循环，两个指针

2. 第一层循环：

   剪枝1：连续的四个值已经超过目标值，直接退出

   剪枝2：去重

   剪枝3：如果第一层循环的当前值加上尾部的三个值，小于目标值，那么直接 continue 掉

3. 第二层循环：

   剪枝1：第一层的当前值 加上 第二层循环连续的三个值，如果超过目标值，直接 break 掉

   剪枝2：第二层去重

   剪枝3：第一层的当前值 加上 第二层的当前值，加上最后的两个值，如果小于目标值，continue 掉

4. 双指针

   left = j + 1, right := n-1

   剪枝1：去重

```go
// date 2023/12/11
func fourSum(nums []int, target int) [][]int {
    // sort
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    ans := make([][]int, 0, 4)
    n := len(nums)

    for i := 0; i < n-3; i++ {
        // 连续四个已经超过目标值，不会再有解，break
        if nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target {
            break
        }
        // 加上最大的三个 小于 目标值,继续检查下一个
        if nums[i] + nums[n-3] + nums[n-2] + nums[n-1] < target {
            continue
        }
        // 去重
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // 注意这里的 for 循环条件也可以这样
        // for j := i+1; j < n-2 && nums[i] + nums[j] + nums[j+1] + nums[j+2] <= target; j++ {
        for j := i+1; j < n-2; j++ {
            if nums[i] + nums[j] + nums[j+1] + nums[j+2] > target {
                break
            }
            if nums[i] + nums[j] + nums[n-2] + nums[n-1] < target {
                continue
            }
            // 去重
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            left, right := j+1, n-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target {
                    // add to res
                    ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
                    left++
                    for left < right && nums[left] == nums[left-1] {
                        left++
                    }
                    right--
                    for left < right && nums[right] == nums[right+1] {
                        right--
                    }
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }

    return ans
}
```

