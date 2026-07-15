## 131 分割回文串-中等

题目：

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是**回文串**。返回 `s` 所有可能的分割方案。



> **示例 1：**
>
> ```
> 输入：s = "aab"
> 输出：[["a","a","b"],["aa","b"]]
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "a"
> 输出：[["a"]]
> ```



**分析：**

本题与 [No.093 复原IP地址](./No.093_复原IP地址.md) 是同一类「**字符串分割**」回溯，区别只有两点：

| | No.093 复原IP地址 | No.131 分割回文串 |
|---|---|---|
| 段的合法性 | 数值 0~255、无前导 0 | 是回文串 |
| 段数 | 固定 4 段 | 任意段 |
| 终止条件 | 凑满 4 段 **且** 用完整个串 | 用完整个串即可 |

正因为**段数不固定**，本题的递归终止就只有一条——「切到字符串末尾 `start == n`」。这正好对应我们讨论 No.093 终止条件时的另一种自然写法：以「用完字符串」作为 gate，无需 `len == k` 的深度上界。框架依然是回溯三要素：

- **路径** `path`：已经切好的回文子串；
- **选择列表**：在 `start` 之后枚举切点 `end`，只要 `s[start..end]` 是回文就切下来；
- **结束条件**：`start == n`，把当前 `path` 收集为一种方案。

回文判断有两种做法：(1) 现场用双指针判断；(2) 先用区间 DP 预处理一张 `g[i][j]` 回文表，回溯时 O(1) 查表。下面分别给出。

### 方法一：回溯 + 双指针判回文 ✅ 推荐

最直白的写法。每个切点用双指针 O(n) 判断 `s[start..end]` 是否回文。`n ≤ 16`，规模极小，这种写法完全够用，且代码最短。

```go
// date 2026/07/15
func partition(s string) [][]string {
	res := make([][]string, 0, 16)
	n := len(s)

	// start: 当前切到的下标；path: 已经切好的回文子串
	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		if start == n { // 切到末尾 → 收集一种方案
			cp := make([]string, len(path))
			copy(cp, path) // 必须 copy，见下方说明
			res = append(res, cp)
			return
		}
		// 枚举切点 end：若 s[start..end] 是回文，就切下来递归
		for end := start; end < n; end++ {
			if isPal(s, start, end) {
				path = append(path, s[start:end+1]) // 做选择
				dfs(end+1, path)
				path = path[:len(path)-1] // 撤销选择（回溯）
			}
		}
	}

	dfs(0, make([]string, 0, n))
	return res
}

// 双指针判断 s[i..j] 是否回文
func isPal(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

> **为什么收集时要 `copy`，而 No.093 不用？** No.093 用 `strings.Join` 把段拼成了一个**不可变的字符串**再存，所以安全；本题存的是 `[]string` 切片，而 `path` 的底层数组在回溯过程中会被反复 `append`/截短，若不 `copy`，后续分支会覆盖掉已收集方案里的内容，最终结果会变成一堆相同的错误切片。这是「收集切片」型回溯的必踩坑点。

### 方法二：回溯 + DP 预处理回文表

回溯过程中同一个子串 `s[i..j]` 会被反复判断多次。可以先用**区间 DP** 预处理一张布尔表 `g[i][j]`（`true` 表示 `s[i..j]` 是回文），回溯时 O(1) 查表。

转移方程（从短的区间推长的）：

$$g[i][j] = \begin{cases} \text{true}, & s[i]=s[j] \text{ 且 } (j-i<2 \text{ 或 } g[i+1][j-1]) \end{cases}$$

- `j-i<2`（长度 1 或 2）只需两端字符相等；
- 更长的区间依赖更短的 `g[i+1][j-1]`，所以外层 `i` 从大到小、内层 `j` 从小到大。

```go
// date 2026/07/15
func partitionDP(s string) [][]string {
	n := len(s)

	// —— 1) 区间 DP 预处理回文表 g[i][j] ——
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- { // 从右下往左上填
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || g[i+1][j-1]) {
				g[i][j] = true
			}
		}
	}

	// —— 2) 回溯，O(1) 查表判回文 ——
	res := make([][]string, 0, 16)
	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		if start == n {
			cp := make([]string, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}
		for end := start; end < n; end++ {
			if g[start][end] { // O(1) 查表
				path = append(path, s[start:end+1])
				dfs(end+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0, make([]string, 0, n))
	return res
}
```

**运行示例：** `s = "aab"`（n=3，方法一的回溯过程）

```
start=0:
  end=0  "a"    回文 ✓ → path=["a"]        → dfs(1)
  end=1  "aa"   回文 ✓ → path=["aa"]       → dfs(2)
  end=2  "aab"  非回文 ✗

由 path=["a"] 进入 start=1:
  end=1  "a"   回文 ✓ → path=["a","a"]     → dfs(2)
  end=2  "ab"  非回文 ✗

由 path=["a","a"] 进入 start=2:
  end=2  "b"   回文 ✓ → path=["a","a","b"] → dfs(3)
    start=3==n → 收集 ["a","a","b"] ✓

由 path=["aa"] 进入 start=2:
  end=2  "b"   回文 ✓ → path=["aa","b"]    → dfs(3)
    start=3==n → 收集 ["aa","b"] ✓

结果：[["a","a","b"],["aa","b"]]
```

对应的回文表 `g`（行 `i`、列 `j`）：

| | j=0 | j=1 | j=2 |
|---|---|---|---|
| i=0 `a` | ✓ `a` | ✓ `aa` | ✗ `aab` |
| i=1 `a` | | ✓ `a` | ✗ `ab` |
| i=2 `b` | | | ✓ `b` |

### 复杂度分析

最坏情况是 `s` 全为相同字符（如 `"aaaa"`），此时每个前缀都是回文，`n-1` 个间隙每个都可切可不切，方案数为 $2^{n-1}$。

| 方法 | 时间复杂度 | 空间复杂度 | 备注 |
|------|-----------|-----------|------|
| 方法一 双指针判回文 | O(n·2ⁿ) | O(n) 递归栈 | 每个切点 O(n) 判回文；代码最短 |
| 方法二 DP 预处理回文表 | O(n² + 2ⁿ) | O(n²) 回文表 | 判回文 O(1)；预处理一次性 O(n²) |

两种方法都受「输出规模 2ⁿ」主导：构造每一份方案本身就要 O(n)。`n ≤ 16` 时 2¹⁵ ≈ 3.3 万，两种写法都能轻松通过。

**如何选择：** 本题 `n ≤ 16`，写**方法一**最简洁、面试最稳；**方法二**的真正价值在 [No.132 分割回文串 II] 这类需要大量回文查询、或 `n` 更大的题目里——把回文表预处理出来后可以反复 O(1) 复用。可以把方法二当作「为 132 铺路」的写法。

### 相邻题

- [No.093 复原IP地址](./No.093_复原IP地址.md)：同为字符串分割回溯；93 段数固定为 4、终止条件是「4 段且用完」，131 段数任意、终止条件只是「用完」。最值得直接对照，体会终止条件的差异。
- [132 分割回文串 II](https://leetcode.cn/problems/palindrome-partitioning-ii/)：求**最少**分割数，DP；本题方法二的回文表正是它的基础。
- [1278 分割回文串 III](https://leetcode.cn/problems/palindrome-partitioning-iii/)：分成恰好 `k` 段、允许改字符使每段回文，求最少修改次数，区间 DP。
- [1745 分割回文串 IV](https://leetcode.cn/problems/palindrome-partitioning-iv/)：能否切成恰好 **3 段**回文——又回到「固定段数」主题，与 No.093 呼应。
