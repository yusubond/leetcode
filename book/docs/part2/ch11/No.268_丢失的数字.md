## 268 丢失的数字

题目：

给定一个包含 `[0, n]` 中 `n` 个数的数组 `nums` ，找出 `[0, n]` 这个范围内没有出现在数组中的那个数。

 

分析：

直接异或计算。

一个值和他本身异或，结果为零。

直接把 0-n 之间所有的数字和 nums 中的数字，异或一次。

没出现的数字，只会出现一次，其他数字出现两次。所以最终的结果就是丢失的数字。

```go
// date 2023/12/05
func missingNumber(nums []int) int {
    ans := int(0)
    n := len(nums)
    for i := 0; i < n; i++ {
        ans ^= nums[i]
        ans ^= i
    }

    ans ^= n

    return ans
}
```

