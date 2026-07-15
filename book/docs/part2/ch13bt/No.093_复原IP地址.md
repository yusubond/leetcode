## 93 复原IP地址-中等

题目：

**有效 IP 地址** 正好由四个整数（每个整数位于 `0` 到 `255` 之间组成，且不能含有前导 `0`），整数之间用 `'.'` 分隔。

- 例如：`"0.1.2.201"` 和` "192.168.1.1"` 是 **有效** IP 地址，但是 `"0.011.255.245"`、`"192.168.1.312"` 和 `"192.168@1.1"` 是 **无效** IP 地址。

给定一个只包含数字的字符串 `s` ，用以表示一个 IP 地址，返回所有可能的**有效 IP 地址**，这些地址可以通过在 `s` 中插入 `'.'` 来形成。你 **不能** 重新排序或删除 `s` 中的任何数字。你可以按 **任何** 顺序返回答案。



> **示例 1：**
>
> ```
> 输入：s = "25525511135"
> 输出：["255.255.11.135","255.255.111.35"]
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "0000"
> 输出：["0.0.0.0"]
> ```
>
> **示例 3：**
>
> ```
> 输入：s = "101023"
> 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
> ```



**分析：**

一个合法 IPv4 地址 = 把 `s` **按原顺序**切成恰好 **4 段**，每段必须同时满足：

1. 数值在 `0 ~ 255`；
2. **不能有前导 `0`**（即 `"0"` 合法，但 `"01"`、`"001"` 非法）；
3. 长度只能是 1 ~ 3。

由此可得一个重要剪枝：4 段每段至少 1 位、最多 3 位，所以 `s` 长度必须在 **4 ~ 12** 之间，否则直接返回空。

这本质上是一个**字符串分割**问题，和 [131 分割回文串](https://leetcode.cn/problems/palindrome-partitioning/) 同构——只是把「段必须是回文」换成「段必须是合法 IP 段」，并且段数固定为 4。下面给出三种解法：

- **方法一（回溯，按段分割）**：用 `start` 指针标记当前切到的位置，每层尝试截取 1~3 位作为下一段，合法就递归；凑满 4 段且恰好用完整个串即为一个答案。✅ 推荐
- **方法二（三重循环枚举切点）**：4 段只需 3 个切点，直接枚举 3 个切点位置并逐段校验，无需递归。
- **方法三（逐位归类回溯）**：逐个数字决定「并入上一段」还是「另起一段」，是原解法的思路。

### 方法一：回溯（按段分割）✅ 推荐

标准回溯三要素：**路径**（已切好的 `segments`）、**选择列表**（下一段截 1/2/3 位）、**结束条件**（凑满 4 段且 `start == n`）。配合两条剪枝（长度范围、剩余字符能否恰好分给剩余段数），代码非常干净，是面试首选写法。

```go
// date 2026/07/15
import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0, 8)
	n := len(s)
	if n < 4 || n > 12 { // 剪枝：合法 IP 长度只能是 4~12
		return res
	}

	// start: 当前切到的下标；segments: 已经切好的若干段
	var dfs func(start int, segments []string)
	dfs = func(start int, segments []string) {
		if len(segments) == 4 { // 凑满 4 段
			if start == n { // 且恰好用完整个串
				res = append(res, strings.Join(segments, "."))
			}
			return
		}
		// 剪枝：剩余字符必须能恰好分给 (4-len) 段，每段 1~3 位
		need := 4 - len(segments)
		if start+need > n || start+3*need < n {
			return
		}
		// 选择列表：尝试截取长度 1、2、3 的段
		for l := 1; l <= 3 && start+l <= n; l++ {
			seg := s[start : start+l]
			if valid(seg) {
				segments = append(segments, seg)       // 做选择
				dfs(start+l, segments)                 // start+l 作为下一次起点
				segments = segments[:len(segments)-1]  // 撤销选择（回溯）
			}
		}
	}

	dfs(0, make([]string, 0, 4))
	return res
}

// valid: 无前导 0，且数值在 0~255
func valid(seg string) bool {
	if len(seg) > 1 && seg[0] == '0' {
		return false // 前导 0，如 "01"、"001"
	}
	v, _ := strconv.Atoi(seg)
	return v <= 255
}
```

> **为什么「`append` 后用 `segments[:len-1]` 撤销」是安全的？** 初始用 `make([]string, 0, 4)` 预留容量 4，递归深度最多 4，`append` 永远不会触发扩容（底层数组不变），因此撤销只需把长度缩回一位，不会破坏父调用的切片。这是 Go 写回溯的常用技巧——若不预留容量，则必须每次 `copy` 出新切片，否则兄弟分支会互相覆盖。

### 方法二：三重循环枚举切点

4 段由 3 个切点 `i < j < k` 唯一决定：四段分别是 `s[:i]`、`s[i:j]`、`s[j:k]`、`s[k:]`。每段长度 1~3，切点取值范围极小，直接三重循环枚举并逐段校验，**完全不需要递归**。思路最直白，适合在面试中快速写对。

```go
// date 2026/07/15
import "strconv"

func restoreIpAddressesLoop(s string) []string {
	n := len(s)
	res := make([]string, 0, 8)
	if n < 4 || n > 12 {
		return res
	}

	ok := func(seg string) bool {
		if len(seg) > 1 && seg[0] == '0' {
			return false
		}
		v, _ := strconv.Atoi(seg)
		return v <= 255
	}

	// i,j,k 为前三段的右端点(不含)；第四段为 s[k:]
	for i := 1; i <= 3 && i < n; i++ {
		for j := i + 1; j <= i+3 && j < n; j++ {
			for k := j + 1; k <= j+3 && k < n; k++ {
				if n-k > 3 { // 第四段长度只能 1~3
					continue
				}
				a, b, c, d := s[:i], s[i:j], s[j:k], s[k:]
				if ok(a) && ok(b) && ok(c) && ok(d) {
					res = append(res, a+"."+b+"."+c+"."+d)
				}
			}
		}
	}
	return res
}
```

### 方法三：逐位归类回溯（原解法）

前两种方法都以「段」为单位做选择。本解法换一个视角——**逐个数字**处理，每读到一个数字 `num`，只有两种选择：

1. **并入上一段**：把上一段的值更新为 `last*10 + num`。仅当上一段当前值非 0（避免前导 0）且合并后 ≤ 255 时允许；
2. **另起一段**：把 `num` 作为新的一段追加。仅当当前段数 < 4 时允许。

`idx == 0` 是首数字的特例——必须另起第 1 段（首段允许为 `"0"`），所以单独处理。这样每位至多 2 个分支，`n ≤ 12` 时实际很快。

```go
// date 2023/12/26（原解法，改名以区分）
import "strconv"

func restoreIpAddressesByDigit(s string) []string {
	res := make([]string, 0, 16)

	var dfs func(s string, idx int, ip []int)
	dfs = func(s string, idx int, ip []int) {
		if idx == len(s) {
			if len(ip) == 4 {
				res = append(res, toString(ip))
			}
			return
		}
		if idx == 0 {
			// IP 地址的第一个值可以为零
			// 所以这里不需要判断 num == 0
			num, _ := strconv.Atoi(string(s[0]))
			ip = append(ip, num)
			dfs(s, idx+1, ip)
		} else {
			// 非 IP 地址的第一个值,都需要判断不为零
			num, _ := strconv.Atoi(string(s[idx]))
			next := ip[len(ip)-1]*10 + num
			// 如果多个可以当做一个
			if next <= 255 && ip[len(ip)-1] != 0 {
				ip[len(ip)-1] = next
				dfs(s, idx+1, ip)
				ip[len(ip)-1] /= 10  // 撤销 num 选择
			}
			if len(ip) < 4 {
				ip = append(ip, num)
				dfs(s, idx+1, ip)
				ip = ip[:len(ip)-1]
			}
		}
	}

	dfs(s, 0, []int{})

	return res
}

func toString(nums []int) string {
	res := strconv.Itoa(nums[0])

	for i := 1; i < len(nums); i++ {
		res += "." + strconv.Itoa(nums[i])
	}

	return res
}
```

**运行示例 1：** `s = "25525511135"`（n=11，方法一的回溯过程）

```
成功分支 ①  → "255.255.11.135"
  start=0  选 "255"   segments=["255"]
  start=3  选 "255"   segments=["255","255"]
  start=6  选 "11"    segments=["255","255","11"]      剩 "135"
  start=8  选 "135"   segments=["255","255","11","135"]  4 段且 start==11 ✓

成功分支 ②  → "255.255.111.35"
  start=6  改选 "111"  segments=["255","255","111"]     剩 "35"
  start=9  选 "35"     4 段且 start==11 ✓

剪枝示例：
  · start=0 想截 4 位 "2552" → 循环 l≤3，根本不会尝试；
  · ["255","255","1"] 后 start=7 剩 "1135"，need=1 段最多吃 3 位，
    而 7+3*1=10 < 11 → 剩余字符太多，剪枝 return，不再往下搜。
```

**运行示例 2：** `s = "0000"`（n=4）

每段都只能截 `"0"`（截 `"00"`/`"000"` 因前导 0 被否决），唯一路径 `["0","0","0","0"]` → 输出 `["0.0.0.0"]`。

### 复杂度分析

| 方法 | 时间复杂度 | 空间复杂度 | 备注 |
|------|-----------|-----------|------|
| 方法一 回溯(按段分割) | O(3⁴·n) ≈ O(n) | O(n) 递归栈 + O(答案) | 最经典、可剪枝，面试首选 |
| 方法二 三重循环 | O(3³·n) ≈ O(n) | O(1) + O(答案) | 无递归，最直白 |
| 方法三 逐位归类 | O(2ⁿ) 最坏 | O(n) 递归栈 | 视角独特，n≤12 实际极快 |

三种方法的搜索空间都被「恰好 4 段、每段 ≤3 位」严格限制：合法划分总数 ≤ 3⁴ = 81 种，因此时间随 `n` 几乎不变，主要开销在拼接输出字符串。`n ≤ 12` 时三种写法都能轻松通过。

**如何选择：** 面试写**方法一**最稳妥（框架通用、剪枝清晰）；若追求代码最短、不想写递归，用**方法二**；**方法三**适合作为「换个决策视角」的思维训练。

### 相邻题

- [131 分割回文串](https://leetcode.cn/problems/palindrome-partitioning/)：同为「字符串分割」回溯；区别在于 131 段数任意且段需是回文，本题段数固定为 4 且段需是合法 IP 段。最值得对照练习。
- [468 验证 IP 地址](https://leetcode.cn/problems/validate-ip-address/)：判断给定串是否为合法 IPv4/IPv6；本题的「段合法性判断」正是它的子问题。
- [132 分割回文串 II](https://leetcode.cn/problems/palindrome-partitioning-ii/)：分割类问题从「回溯枚举」到「DP 计数」的延伸（求最少分割数）。
- [No.022 括号生成](./No.022_括号生成.md)：同在本章，同为「构造所有合法结构」的 DFS，可对照「结束条件 + 选择」的写法。
