## 169 多数元素-中等

题目：

给定一个大小为 `n` 的数组 `nums` ，返回其中的多数元素。多数元素是指在数组中出现次数 **大于** `⌊ n/2 ⌋` 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。



> **示例 1：**
>
> ```
> 输入：nums = [3,2,3]
> 输出：3
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [2,2,1,1,1,2,2]
> 输出：2
> ```


**解题思路**

这道题是求数组中的众数，可有多种解法。

- 排序，详见解法1。因为众数的个数大于 n/2，所以排序以后，中间位置的元素肯定是众数。
- map 去重，详见解法2。具体是，遍历数组，存入 map，统计出现的次数；然后遍历 map 找到出现次数最多的 key。
- Boyer-Moore 投票算法，详见解法3。如果我们把众数记为 +1，其他数记为 -1，全部加起来，那么和必然大于0。从结果本身也能看出众数比其他数多，这就是 Boyer-Moore 算法。具体步骤就是：
  1. 维护一个候选众数 ans，和它出现的次数 cnt，初始值 ans = nums[0]， cnt = 1
  2. 遍历数组，如果元素 x 等于 ans，那么 cnt 递增；不相等，则 cnt 递减；如果 cnt 小于等于 0，则把当前元素 x 选为 ans，cnt = 1，重复1，2
  3. 最后得到的 ans 就是众数



```go
// date 2024/01/16
// 解法1
// 排序
func majorityElement(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    return nums[len(nums)>>1]
}

// 解法2
// map 统计出现的次数
func majorityElement(nums []int) int {
    set := make(map[int]int, 16)
    for _, v := range nums {
        set[v]++
    }
    ans, cnt := 0, 0
    for k, v := range set {
        if v > cnt {
            cnt = v
            ans = k
        }
    }

    return ans
}

// 解法3
// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    ans := nums[0]
    cnt := 1
    for i := 1; i < n; i++ {
        if ans == nums[i] {
            cnt++
        } else {
            cnt--
            if cnt <= 0 {
                ans = nums[i]
                cnt = 1
            }
        }
    }
    return ans
}
```

