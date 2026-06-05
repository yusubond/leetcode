## 918 环形子数组的最大和-中等

题目：

给定一个长度为 `n` 的**环形整数数组** `nums` ，返回 *`nums` 的非空 **子数组** 的最大可能和* 。

**环形数组** 意味着数组的末端将会与开头相连呈环状。形式上， `nums[i]` 的下一个元素是 `nums[(i + 1) % n]` ， `nums[i]` 的前一个元素是 `nums[(i - 1 + n) % n]` 。

**子数组** 最多只能包含固定缓冲区 `nums` 中的每个元素一次。形式上，对于子数组 `nums[i], nums[i + 1], ..., nums[j]` ，不存在 `i <= k1, k2 <= j` 其中 `k1 % n == k2 % n` 。



分析：

单调队列。

既然可以环形，那么我们将数组延长一倍，对于 i >= n 的元素，nums[i] = nums[i-n]。

问题就变成了：在长度为 2n 的数组上，求长度不超过 n 的最大子数组和。

思路：前缀和作差。

令 Si 表示前缀和，那么在不规定子数组的长度时，求最大的子数组和，就等于求 Si - Sj 的最大值，其中 j < i。

现在规定了，子数组的长度不能超过 n，那我们用单调队列维护该集合。

1. 遍历到 i 时，如果队头元素的坐标小于 i - n（即子数组长度超过 n），就需要出队，一直重复，知道队列为空或者队头元素的坐标大于或等于 i - n
2. 取 队头元素的值，计算 Si - Sj，并更新结果
3. 维护队列的单调性。Si 加入队列的时候，任何大于等于 Si 的元素都需要出队，然后 Si 在入队。

```go
// date 2023/11/30
func maxSubarraySumCircular(nums []int) int {
    type pair struct {
        idx int
        val int
    }
    n := len(nums)
    preSum := nums[0]
    res := nums[0]
    queue := make([]pair, 0, 32)
    queue = append(queue, pair{idx:0, val: preSum})

    for i := 1; i < 2*n; i++ {
        for len(queue) > 0 && queue[0].idx < i - n {
            queue = queue[1:]
        }

        preSum += nums[i%n]

        res = max(res, preSum - queue[0].val)

        for len(queue) > 0 && queue[len(queue)-1].val >= preSum {
            queue = queue[:len(queue)-1]
        }

        queue = append(queue, pair{idx:i, val: preSum})
    }

    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

