## 128 最长连续序列-中等

题目：

给定一个未排序的整数数组 `nums`，找出数字连续的**最长序列**（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 `O(n)` 的算法解决此问题。



**示例 1：**

```
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
```

**示例 2：**

```
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
```



**解题思路**

要求 `O(n)`，所以不能先排序（排序是 `O(n log n)`）。需要借助**哈希集合**实现 `O(1)` 查找。

- 哈希集合。把所有数放进集合（顺便去重），再通过「只从序列起点向上计数」的方式统计，详见解法1。时间 `O(n)`、空间 `O(n)`。
- 排序 + 线性扫描。排序后相邻相同的连续段即可求出长度。时间 `O(n log n)`、空间 `O(log n)`，不满足本题的 `O(n)` 要求，但实现简单，可作为对比，详见解法2。

解法1 的关键在于**只从「序列起点」开始计数**：当 `num - 1` 也在集合中时，说明 `num` 不是起点，直接跳过，避免重复统计。这样每段连续序列只会被从起点完整扫描一次。虽然外层 `for` 嵌套内层 `while`，但每个元素只会被计入唯一一条序列，内层 `while` 的总执行次数等于元素总数 `n`，因此整体均摊为 `O(n)`。

```go
// date 2026/06/21
// 解法1
// 哈希集合：只从序列起点向上计数
// 时间复杂度 O(N)，空间复杂度 O(N)
func longestConsecutive(nums []int) int {
    // 把所有数放进集合：去重 + O(1) 查找
    set := make(map[int]bool, len(nums))
    for _, num := range nums {
        set[num] = true
    }

    longest := 0
    for num := range set {
        // 只从「序列起点」开始计数：num-1 不在集合中，
        // 说明 num 是某段连续序列的开头，避免重复计数
        if !set[num-1] {
            curNum := num
            curLen := 1
            // 持续向上找连续的下一个数
            for set[curNum+1] {
                curNum++
                curLen++
            }
            if curLen > longest {
                longest = curLen
            }
        }
    }
    return longest
}
```

```go
// date 2026/06/21
// 解法2
// 排序 + 线性扫描（不满足 O(n)，仅作对比）
// 时间复杂度 O(N log N)，空间复杂度 O(log N)
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Ints(nums)

    longest, curLen := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            continue // 跳过重复元素
        }
        if nums[i] == nums[i-1]+1 {
            curLen++ // 仍然连续
        } else {
            curLen = 1 // 断开，重新开始
        }
        if curLen > longest {
            longest = curLen
        }
    }
    return longest
}
```

以 `nums = [100,4,200,1,3,2]` 为例：

- 集合 = `{1,2,3,4,100,200}`
- `100`：`99` 不在集合 → 是起点，向上找 `101` 不在，长度 1
- `4`：`3` 在集合 → **跳过**（不是起点，避免重复统计）
- `200`：是起点，长度 1
- `1`：是起点，向上 `2 → 3 → 4`，长度 4
- 最长 `4`
