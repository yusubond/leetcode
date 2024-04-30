## 16 最接近的三数之和-中等

题目：

给你一个长度为 `n` 的整数数组 `nums` 和 一个目标值 `target`。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



> **示例 1：**
>
> ```
> 输入：nums = [-1,2,1,-4], target = 1
> 输出：2
> 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0,0,0], target = 1
> 输出：0
> ```



**解题思路**

这道题可用双指针解决。思路是先排序，然后对撞指针找最接近的答案。

```go
// date 2024/01/15
func threeSumClosest(nums []int, target int) int {
    // sort and left, right
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    ans := math.MaxInt32
    n := len(nums)
    for i := 0; i < n; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == target {
                return sum
            }
            // update the ans
            if abs(ans, target) > abs(sum, target) {
                ans = sum
            }
            if sum > target {
                r--
            } else {
                l++
            }
        }
    }


    return ans
}

func abs(x, y int) int {
    if x > y {
        return x-y
    }
    return y-x
}
```

