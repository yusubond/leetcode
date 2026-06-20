## 1 两数之和-简单

题目：

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** *`target`* 的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



**解题思路**

题目中已经说明每种输入只会对应一种输出，所以可以有多种解法。

- 两层循环。第一层遍历当前 x，然后从元素的 x 的下一个元素 y 开始遍历，如果 x + y 等于目标值，则返回两者的坐标，详见解法1。

- 哈希表。利用map存储已经遍历过的元素和下标，达到O(1)查找，详见解法2。时间和空间复杂度都是`O(N)`。

- 排序 + 双指针。带着原下标排序后，左右指针从两端向中间逼近。时间`O(N log N)`、空间`O(N)`；本题不是最优，但是通往 167（有序两数之和）、15（三数之和）等题的通用套路，详见解法3。

```go
// date 2020/09/14
// 解法1
// 两层循环
// 时间复杂度O(N^2)
func twoSum(nums []int, target int) []int {
    res := make([]int, 0, 2)
    n := len(nums)

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if nums[i] + nums[j] == target {
                res = append(res, i, j)
                return res
            }
        }
    }
    
    return res
}
// 解法2
// map 查找
func twoSum(nums []int, target int) []int {
    set := make(map[int]int, len(nums))
    res := []int{}
    for j, v := range nums {
        d := target - v
        if i, ok := set[d]; ok && i != j {
            res = append(res, i, j)
            break
        } else {
            set[v] = j
        }
    }
    return res
}
// 解法3
// 排序 + 双指针（需导入 sort 包）
// 时间复杂度 O(N log N)，空间 O(N)
func twoSum(nums []int, target int) []int {
    type pair struct{ val, idx int }
    arr := make([]pair, len(nums))
    for i, v := range nums {
        arr[i] = pair{v, i}
    }
    sort.Slice(arr, func(a, b int) bool { return arr[a].val < arr[b].val })

    l, r := 0, len(arr)-1
    for l < r {
        s := arr[l].val + arr[r].val
        if s == target {
            return []int{arr[l].idx, arr[r].idx}
        } else if s < target {
            l++ // 和太小，左指针右移
        } else {
            r-- // 和太大，右指针左移
        }
    }
    return nil
}
```

面试要点

这道题是最高频的面试题之一，除了写对代码，下面几点更体现工程素养。

**动手前先澄清**（别立刻编码）：确认"恰好一个答案、找不到返回什么、能否有负数 / 重复元素"。这 30 秒本身就是加分项。

**复杂度对比**

| 解法 | 时间 | 空间 | 备注 |
|------|------|------|------|
| 暴力两层循环 | O(N²) | O(1) | 面试起点，N 小可接受 |
| 哈希表 | O(N) | O(N) | ✅ 本题最优 |
| 排序 + 双指针 | O(N log N) | O(N) | 输入已排序则 O(N) / O(1) |

**加分动作**

- 先讲暴力再优化："先用 O(N²) 保证正确，再用哈希表降到 O(N)。"——展示思维推进。
- 写完用一个例子手走一遍（如 `[2,7,11,15], target=9`），顺带覆盖重复值 / 自配边界。
- 主动报复杂度并说清权衡：用 O(N) 空间换 O(N) 时间，值得。

**高频追问**

- 输入已排序？→ 双指针，O(N) 时间 O(1) 空间（即 167 题）。
- 返回数值而非下标？→ 排序 + 双指针更自然，无需保留下标。
- 求所有满足的对 / 答案不唯一？→ 排序 + 双指针收集，或哈希表存"值 → 下标列表"（见下方扩展题目）。
- 3Sum / 4Sum？→ 固定一个 / 两个数 + 双指针。
- 数据量放不进内存？→ 外部排序 / 分块（MapReduce 思路）。
- 能否优于 O(N)？→ 不能，至少要把每个元素看一遍（Ω(N) 下界）。



扩展题目：两数之和

**问题延伸：给定一个无序且元素可重复的数组a[],以及一个数sum，求a[]中是否存在两个元素的和等于sum，并输出这两个元素的下标。答案可能不止一种，请输出所有可能的答案结果集。**

这是VMware面试中的一道题，以下是个人的解法。这个算法的时间复杂度为O(n)，但是需要辅助空间。

```go
// date 2021/02/21
// 解法一：两层循环
func twoSum(nums []int, target int) [][]int {
  res := make([][]int, 0, 16)
  idx := make(map[int][]int, 16)
  for i, v := range nums {
    d := target - v
    if oldIdx, ok := idx[d]; ok && len(oldIdx) != 0 {
      for _, j := range oldIdx {
        res = append(res, []int{j, i})
      }
    }
    _, ok := idx[v]
    if !ok {
      idx[v] = []int{}
    }
    idx[v] = append(idx[v], i)
  }
  return res
}
```

