## 673 最长递增子序列的个数-中等

题目：

给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。


分析：

dp 表示递增子序列的长度。

cnt 表示以 nums[i] 结尾的长度为dp[i]的个数。

```go
// date 2023/11/13
func findNumberOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    maxLen := 0

    ans := 0

    for i := 0; i < n; i++ {
        dp[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if dp[j] + 1 > dp[i] {
                    dp[i] = dp[j] + 1
                    cnt[i] = cnt[j]   // dp[i] 在更新，所以重置cnt[i]个数
                } else if dp[j] + 1 == dp[i] {
                    cnt[i] += cnt[j]  // 说明存在多个，累加
                }

            }
        }
		// 最终结果集，更新或累加
        if dp[i] > maxLen {
            maxLen = dp[i]
            ans = cnt[i]
        } else if dp[i] == maxLen {
            ans += cnt[i]
        }
    }
    return ans
}
```


**关于方法一转移的细节**：`dp[i]` 是以 `nums[i]` 结尾的 LIS 长度，`cnt[i]` 是对应方案数。扫 `j < i` 时：

- `dp[j]+1 > dp[i]`：发现更长的，**重置** `cnt[i] = cnt[j]`（更短的方案作废）；
- `dp[j]+1 == dp[i]`：长度相同，**累加** `cnt[i] += cnt[j]`。

最后再对全局 `maxLen` 做同样的「重置/累加」收尾。复杂度：时间 O(n²)，空间 O(n)。

**方法二：O(n log n) 树状数组（Fenwick）**

当 `n` 很大（如 2×10⁵）时 O(n²) 会超时。优化思路：把「以值 `x` 结尾」的信息挂到**值域**上，用树状数组维护每个值处的 `(最长长度, 方案数)`，查询「结尾严格小于 `x` 的前缀」即可 O(log n) 得到可接上的最长长度及其方案数。

节点存 `(l, c)`：`query(i)` 取前缀 `[1..i]` 中**长度最大**者，长度相同则 **cnt 相加**。对每个 `x`：先 `query(rank(x)-1)` 得到结尾 `<x` 的 `(l, c)`，则 `x` 形成 `(l+1, c)`，再 `update(rank(x), (l+1, c))`。`rank` 为离散化后的 1-based 排名。

> `query` 初值取 `(0, 1)`：查到长度 0（无更小元素）时方案数按 1 算，保证单元素 `x` 得到 `(1, 1)`。

```go
// date 2026/06/20   需 import "sort"
func findNumberOfLIS(nums []int) int {
    // 离散化：值映射到 1-based 排名
    sorted := append([]int(nil), nums...)
    sort.Ints(sorted)
    uniq := sorted[:0]
    for i, v := range sorted {
        if i == 0 || v != sorted[i-1] {
            uniq = append(uniq, v)
        }
    }
    rank := func(v int) int { return sort.SearchInts(uniq, v) + 1 }

    t := make([]node, len(uniq)+1)
    for _, x := range nums {
        r := rank(x)
        q := query(t, r-1)              // 结尾严格小于 x 的 (maxLen, cnt)
        update(t, r, node{q.l + 1, q.c}) // 接上 x
    }
    return query(t, len(uniq)).c
}

type node struct{ l, c int }

// 前缀 [1..i] 中长度最大者；长度相同则 cnt 相加
func query(t []node, i int) node {
    res := node{0, 1} // 长度 0 视为 1 种方案，保证单元素得 (1,1)
    for ; i > 0; i -= i & (-i) {
        res = merge(res, t[i])
    }
    return res
}
func update(t []node, i int, v node) {
    for ; i < len(t); i += i & (-i) {
        t[i] = merge(t[i], v)
    }
}
func merge(a, b node) node {
    if a.l > b.l { return a }
    if b.l > a.l { return b }
    return node{a.l, a.c + b.c}
}
```

复杂度：时间 O(n log n)，空间 O(n)。
