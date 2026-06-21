## 239 滑动窗口最大值-困难

题目：

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

题目链接：https://leetcode.cn/problems/sliding-window-maximum/



**解题思路：单调队列**

固定窗口大小，形成窗口不难，难在如何快速取出窗口内的最大值。朴素做法每次扫描窗口求最大值为 O(nk)，借助「单调队列」可优化到 O(n)。

维护一个双端队列，里面存的是元素**下标**，且这些下标对应的值**从队首到队尾单调递减**，于是**队首永远是当前窗口的最大值**。

关键在「队尾弹出」：新元素 `nums[i]` 入队时，把队尾所有 `<= nums[i]` 的下标弹掉——它们既比新元素小，又比新元素更早出窗，只要新元素还在窗口内，它们就永远当不了最大值，可以直接淘汰。

每个元素 i 依次执行：

1. **队首出窗**：若队首下标 `< i-k+1`（超出窗口左边界 `[i-k+1, i]`），从队首弹出。
2. **维持单调递减**：从队尾弹出所有 `nums[队尾] <= nums[i]` 的下标。
3. **当前下标入队尾**。
4. **记录答案**：当 `i >= k-1`（第一个窗口形成后），队首即为本轮最大值。

每个元素最多入队、出队各一次，所以整体为 O(n)。

```go
// date 2026/06/21
// 单调队列：队列存下标，对应值单调递减，队首即窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    ans := make([]int, 0, n-k+1)
    deque := make([]int, 0, k) // 存下标，对应值从队首到队尾单调递减

    for i := 0; i < n; i++ {
        // 1. 队首出窗：下标超出 [i-k+1, i] 范围
        for len(deque) > 0 && deque[0] < i-k+1 {
            deque = deque[1:]
        }
        // 2. 维持单调递减：队尾 <= 当前值的下标全部弹出
        for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
            deque = deque[:len(deque)-1]
        }
        // 3. 当前下标入队尾
        deque = append(deque, i)
        // 4. 窗口形成（i == k-1 起），队首即本轮最大值
        if i >= k-1 {
            ans = append(ans, nums[deque[0]])
        }
    }

    return ans
}
```

> 用 `deque[1:]` 弹队首、`deque[:len-1]` 弹队尾，是把切片当双端队列用的地道写法；若担心频繁扩容，可换成定长数组 + `head/tail` 双指针，逻辑不变。

**手动验证**：`nums = [1,3,-1,-3,5,3,6,7]`，`k = 3`

| i | nums[i] | deque(下标) | deque(值)   | 输出 |
|---|---------|------------|-------------|------|
| 0 | 1       | [0]        | [1]         | —    |
| 1 | 3       | [1]        | [3]         | —    |
| 2 | -1      | [1,2]      | [3,-1]      | 3    |
| 3 | -3      | [1,2,3]    | [3,-1,-3]   | 3    |
| 4 | 5       | [4]        | [5]         | 5    |
| 5 | 3       | [4,5]      | [5,3]       | 5    |
| 6 | 6       | [6]        | [6]         | 6    |
| 7 | 7       | [7]        | [7]         | 7    |

结果为 `[3,3,5,5,6,7]` ✓

**复杂度**

- 时间复杂度 O(n)：每个下标最多入队、出队各一次。
- 空间复杂度 O(k)：队列最多存 k 个下标。



**解法2：值单调队列**

队列直接存值（不存下标），同样保持单调递减，队首即窗口最大值；窗口左端元素 `nums[left]` 出窗时，若恰好等于队首则一并弹出。由于相等的最大值在求最大值时彼此可替代，这种「按值出队」对本题等价正确，但下标版本更直观、不易出错，推荐解法1。

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
        if right-left >= k {
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
