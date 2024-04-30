## 905 按奇偶排序数组-简单

题目：

给定一个数组，通过就地修改使所有偶数位于所有基数的前面。你可以返回满足此条件的任何数组作为答案。



分析：

对撞指针，然后交换。

```go
// date 2022/10/7
func sortArrayByParity(nums []int) []int {
    left, right := 0, len(nums)-1
    for left < right {
        for left < right && nums[left] & 0x1 == 0x0 {
            left++
        }
        for left < right && nums[right] & 0x1 == 0x1 {
            right--
        }
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
    return nums
}
```