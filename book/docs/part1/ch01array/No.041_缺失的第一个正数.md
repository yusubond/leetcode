## 41 缺失的第一个正数-困难

题目：

给你一个未排序的整数数组 `nums` ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 `O(n)` 并且只使用常数级别额外空间的解决方案。



**解题思路**

- 排序，详见解法1。要求最小的正数，可以先讲数组排序，然后从头开始遍历；遍历的时候注意1）跳过小于等于零的元素；2）跳过重复的元素。假定结果 ans 为1，遍历的时候如果元素等于 ans，那么 ans+1，如果不等于那么 ans 就是答案。

- 哈希表，详见解法2。

  最朴素的做法就是遍历原数组把大于等于零的元素都存入哈希表；然后从 1 开始枚举正整数从哈希表里查找，第一个不在哈希表里的正整数就是答案。

  这样的设计需要额外的空间，题目要求常数级的空间，所以需要考虑把数组设计成哈希表的替代品。

  一个事实。

  对于一个长度为 N 的数组，其中没有出现的最小正整数只能在区间`[1, N+1]`中。如果区间`[1, N]`中所有元素都在数组中出现了，那么答案就是 N+1，否则答案就在`[1,N]`中。

  基于这样的事实，我们可以这样设计：

  1）第一次遍历，把小于等于零的元素都修改为 N+1

  2）第二次遍历，对元素求绝对值x，如果绝对值x在区间`[1,N]`中，那么把数组`x-1`位置的元素变成其值的负数。

  3）第三次遍历，**经过第二次遍历数组中出现过的元素，已经把元素映射成索引所对应位置上的元素变成负数**，所以此时再出现的正数，其索引+1就是答案。

```go
// date 2024/01/24
// 解法1 排序
func firstMissingPositive(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    ans := 1
    for i, v := range nums {
        if v <= 0 {
            continue
        }
        if i > 0 && v == nums[i-1] {
            continue
        } 
        if v == ans {
            ans++
        } else {
            break
        }
    }
    return ans
}

// 解法2
// 设计哈希表
func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] <= 0 {
            nums[i] = n+1
        }
    }
    for i := 0; i < n; i++ {
        v := abs(nums[i])
        if v <= n {
            nums[v-1] = -abs(nums[v-1])
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i+1
        }
    }
    return n+1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

