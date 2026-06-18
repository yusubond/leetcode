## 3437 全排列3-中等

给定一个整数 `n`，一个 **交替排列** 是没有 **两个** 相邻元素 **同时** 为奇数或偶数的前 `n` 个正整数的排列。

返回所有这样的 **交替排列** 以字典序排序。



**解题思路**

这道题也是回溯算法，通过剪枝控制进入回溯的时机，即：当 前继节点 和 当前节点 不同时为奇数或偶数，且没有遍历过时，才进入回溯选择。

```go
// date 2026.06.18
// 这里虽然写 backtrack 实际就是 dfs
func permute(n int) [][]int {
	res := make([][]int, 0, 16)

	visited := make(map[int]bool, n)
	path := []int{}

	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			one := make([]int, n)
			copy(one, path)
			res = append(res, one)
			return
		}
		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}
			if len(path) > 0 && path[len(path)-1]%2 == i%2 {
				continue
			}

			// pick up i
			visited[i] = true
			path = append(path, i)

			backtrack()

			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()

	return res
}
```

**复杂度分析**

设答案数为 `F(n)`（见下方公式），则时间复杂度为 `O(F(n) · n)`，空间复杂度 `O(n)`（递归栈 + `visited`）。由于 `F(n)` 随 `n` 增长极快，该解法只适用于较小的 `n`（约 `n ≤ 10`）。



**进阶：只计数（O(n)）**

很多平台的原题只问**有多少个**交替排列（如要求对 `1e9+7` 取模）。利用结构特性可以推出闭式解：

相邻元素奇偶性必须不同 ⇒ 奇数和偶数必须**交错**。前 `n` 个正整数中奇数有 `⌈n/2⌉` 个、偶数有 `⌊n/2⌋` 个，二者数量差恒 ≤ 1，所以总是可行的，只有两种交错模式：

- `n` 为偶数（奇 = 偶）：`奇偶奇偶...` 或 `偶奇偶奇...`，共 2 种；
- `n` 为奇数（奇 = 偶 + 1）：只能 `奇偶奇偶奇...`，1 种。

模式一旦确定，奇数内部、偶数内部各自任意排列互不影响，因此：

```katex
F(n) = \begin{cases}
2 \cdot m! \cdot m!, & n = 2m \\
(m+1)! \cdot m!, & n = 2m+1
\end{cases}
```

| n | 奇 | 偶 | 公式 | 结果 |
|---|---|---|------|------|
| 1 | 1 | 0 | `1!·0!` | 1 |
| 2 | 1 | 1 | `2·1!·1!` | 2 |
| 3 | 2 | 1 | `2!·1!` | 2 |
| 4 | 2 | 2 | `2·2!·2!` | 8 |
| 5 | 3 | 2 | `3!·2!` | 12 |

```go
// date 2026.06.18
// 仅计数，结果对 1e9+7 取模。时间 O(n)，空间 O(1)
const MOD = 1_000_000_007

func permuteCount(n int) int {
	odd, even := (n+1)/2, n/2
	fact := func(x int) int {
		r := 1
		for i := 2; i <= x; i++ {
			r = r * i % MOD
		}
		return r
	}
	if odd == even { // 两种交错模式：奇偶... / 偶奇...
		return 2 * fact(odd) % MOD * fact(even) % MOD
	}
	// odd == even+1：只能 奇偶奇...
	return fact(odd) * fact(even) % MOD
}
```

