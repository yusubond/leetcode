## 239 滑动窗口最大值-困难

题目：

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



**解题思路**

这道题可用滑动窗口解，且辅助优先队列。优先队列用于求窗口内的最大值。

这个问题是固定窗口大小，所以形成窗口不难，难在如何对窗口内的元素求最大值。

可以使用辅助数组，维护窗口内的递增子序列。每次新增元素的时候，用插入排序保证序列中比当前值小的元素都剔除。

一旦窗口形成，需要移除序列队头元素的时候，与窗口第一个元素进行判断。


```go
// date 2023/11/20
func maxSlidingWindow(nums []int, k int) []int {
    left, right := 0, 0
    n := len(nums)
    ans := make([]int, 0, 64)
    priQueue := make([]int, 0, 16)

    for right < n {
        for len(priQueue) != 0 && nums[right] > priQueue[len(priQueue)-1] {
            priQueue = priQueue[:len(priQueue)-1]
        }
        priQueue = append(priQueue, nums[right])
        right++
        if right - left >= k {
            ans = append(ans, priQueue[0])
            if priQueue[0] == nums[left] {
                priQueue = priQueue[1:]
            }
            left++
        }
    }

    return ans
}
```
