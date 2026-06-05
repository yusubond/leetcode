## 503 下一个更大的元素2-中等

题目：

给定一个循环数组 `nums` （ `nums[nums.length - 1]` 的下一个元素是 `nums[0]` ），返回 *`nums` 中每个元素的 **下一个更大元素*** 。

数字 `x` 的 **下一个更大的元素** 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 `-1` 。



> **示例 1:**
>
> ```
> 输入: nums = [1,2,1]
> 输出: [2,-1,2]
> 解释: 第一个 1 的下一个更大的数是 2；
> 数字 2 找不到下一个更大的数； 
> 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [1,2,3,4,3]
> 输出: [2,3,4,-1,4]
> ```



分析：

输入为循环数组，所以考虑把数组膨胀1倍，然后逆序遍历，求新数组的下一个更大元素。然后只返回结果集的前一半。

```go
// date 2023/12/18
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    nums2 := make([]int, 2*n, 2*n)
    for i, v := range nums {
        nums2[i] = v
        nums2[n+i] = v
    }
    stack := make([]int, 0, 16)
    ans := make([]int, 2*n)
    for i := len(nums2)-1; i >= 0; i-- {
        v := nums2[i]
        for len(stack) > 0 && v >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 {
            ans[i] = stack[len(stack)-1]
        } else {
            ans[i] = -1
        }
        stack = append(stack, v)
    }

    ans = ans[:n]
    return ans
}
```

