## 16 最接近的三数之和-中等

题目：

给你一个长度为 `n` 的整数数组 `nums` 和 一个目标值 `target`。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



分析：

先排序，然后双指针。

```go
// date 2023/12/08
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
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
            } else {
                l++
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
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

