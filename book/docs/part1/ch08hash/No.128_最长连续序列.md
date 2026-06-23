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

要求 `O(n)`，所以不能先排序（排序是 `O(n log n)`）。需要借助**哈希表**实现 `O(1)` 查找。

- **单次遍历 + 动态边界更新**（推荐）。遍历时用哈希表记录包含当前数字的连续序列长度，通过查询相邻数字的端点值来合并序列。真正的一趟扫描，每个元素只做 O(1) 次操作，没有内层循环，常数最优，详见解法1。
- 哈希集合 + 只从起点计数。把所有数放进集合，只从序列起点向上数，也很直观。均摊 `O(n)` 但存在内层 `while` 循环，对极长的连续序列（如 `[1..10^5]`）内层迭代次数多，可能被判超时，详见解法2。
- 排序 + 线性扫描。排序后扫描相邻连续段，`O(n log n)`，不满足本题要求，仅作对比，详见解法3。

解法1 的关键在于**每个元素只遍历一次**：处理 `num` 时，查 `num-1` 和 `num+1` 在哈希表里已有的序列长度，三者合并即可得到包含 `num` 的序列长度。同时只需更新**序列两端**的长度值，因为后续新元素只会通过 `num-1`（查右邻居）和 `num+1`（查左邻居）来拼接，中间元素的值不会再被访问。

```go
// date 2026/06/23
// 解法1（推荐）
// 单次遍历 + 动态边界更新
// 时间复杂度 O(N)，空间复杂度 O(N)
func longestConsecutive(nums []int) int {
    // seqCt[num] 表示包含 num 的连续序列的长度
    seqCt := make(map[int]int)

    res := 0
    for _, num := range nums {
        // 跳过重复元素
        if seqCt[num] > 0 {
            continue
        }

        // 查询左右相邻数字所在序列的长度
        left := seqCt[num-1]  // 以 num-1 结尾的连续序列长度
        right := seqCt[num+1] // 以 num+1 开头的连续序列长度

        // 合并后的序列总长度
        sum := left + right + 1
        seqCt[num] = sum
        if sum > res {
            res = sum
        }

        // 只更新序列两端的边界值——后续新元素拼接时
        // 只会查 num-1 或 num+1，所以中间元素的值不需维护
        seqCt[num-left] = sum
        seqCt[num+right] = sum
    }
    return res
}
```

```go
// date 2026/06/21
// 解法2
// 哈希集合：只从序列起点向上计数
// 时间复杂度 O(N)，空间复杂度 O(N)
// 注意：内层 while 可能因极长连续序列而超时
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
            curNum, curLen := num, 1
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
// 解法3
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

**解法1（DP 边界更新）** 以 `nums = [100,4,200,1,3,2]` 为例：

```
num=100: left=seqCt[99]=0, right=seqCt[101]=0, sum=1
         seqCt[100]=1, 边界 seqCt[100]=1, seqCt[100]=1
         {100:1}

num=4:   left=seqCt[3]=0, right=seqCt[5]=0, sum=1
         {100:1, 4:1}

num=200: left=seqCt[199]=0, right=seqCt[201]=0, sum=1
         {100:1, 4:1, 200:1}

num=1:   left=seqCt[0]=0, right=seqCt[2]=0, sum=1
         {100:1, 4:1, 200:1, 1:1}

num=3:   left=seqCt[2]=0, right=seqCt[4]=1, sum=0+1+1=2
         seqCt[3]=2
         更新左边界 seqCt[3-0]=2 → seqCt[3]=2
         更新右边界 seqCt[3+1]=2 → seqCt[4]=2

num=2:   left=seqCt[1]=1, right=seqCt[3]=2, sum=1+2+1=4  ← 合并！
         seqCt[2]=4
         更新左边界 seqCt[2-1]=4 → seqCt[1]=4
         更新右边界 seqCt[2+2]=4 → seqCt[4]=4
         最终 {1:4, 100:1, 200:1, 4:4, 2:4, 3:2}
               ↑        ↑
            端点=4   端点=4  ← 只有端点值准确，供后续拼接

结果 4 ✓
```

**解法2（哈希集合起点计数）** 以 `nums = [100,4,200,1,3,2]` 为例：

- 集合 = `{1,2,3,4,100,200}`
- `100`：`99` 不在集合 → 是起点，向上找 `101` 不在，长度 1
- `4`：`3` 在集合 → **跳过**（不是起点，避免重复统计）
- `200`：是起点，长度 1
- `1`：是起点，向上 `2 → 3 → 4`，长度 4
- 最长 `4`
