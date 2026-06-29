## 162 寻找峰值-中等

题目：

峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 `nums`，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 **任何一个峰值** 所在位置即可。

你可以假设 `nums[-1] = nums[n] = -∞` 。

你必须实现时间复杂度为 `O(log n)` 的算法来解决此问题。

> **示例 1：**
>
> ```
> 输入：nums = [1,2,3,1]
> 输出：2
> 解释：3 是峰值元素，你的函数应该返回其索引 2。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [1,2,1,3,5,6,4]
> 输出：1 或 5
> 解释：你的函数可以返回索引 1（峰值元素为 2），或者返回索引 5（峰值元素为 6）。
> ```

**关键观察：无序数组为何能二分**

这并非在“有序”数组上二分，而是利用“爬山”的**方向单调性**——每比较一次都能安全地砍掉一半。

- `nums[-1] = nums[n] = -∞` 保证两端必然“下降”，所以峰值**一定存在**（全局最大值本身就是峰值之一）。
- 若 `nums[mid] < nums[mid+1]`，则 `[mid+1, n-1]` 内必有峰值：从 `mid+1` 出发向右沿“严格上升”方向走，由于右端是 `-∞`，迟早会遇到“先升后降”的拐点，那就是峰值。
- 反之若 `nums[mid] > nums[mid+1]`，则峰值在左侧 `[left, mid]`。
- 依据这一比较，每次都能把搜索区间缩小一半，故为 `O(log n)`。

> LeetCode 原题附带约束 `nums[i] != nums[i+1]`（相邻元素互不相等），这正是“峰值必然存在、二分只需做严格比较”的前提；否则可能出现 `[2,2,2]` 这种无峰值的平台。

**解法 1：判断 mid 是否为峰值，否则向高处走**

用 `getNum` 把越界访问（`-1` 与 `n`）统一映射为 `-∞`（即 `math.MinInt64`），从而无需单独处理端点；找到峰值即返回，否则按 `mid` 与 `mid+1` 的大小关系决定向哪侧走。

```go
// date 2024/03/04
func findPeakElement(nums []int) int {
	n := len(nums)
	var getNum func(i int) int
	getNum = func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if getNum(mid-1) < getNum(mid) && getNum(mid) > getNum(mid+1) {
			return mid
		}
		if getNum(mid) < getNum(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
```

**解法 2：只比 mid 与 mid+1（推荐）**

更进一步：**不必单独判断 mid 是否为峰值**。只需比较 `nums[mid]` 与 `nums[mid+1]`，往“更高”的一侧收缩，区间最终收敛到唯一一个峰值点。

循环用 `left < right`，于是恒有 `mid < right`，故 `mid+1` 永不越界，省去 `getNum` 边界处理。

```go
// date 2026/06/29
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// mid 比右邻居高 → 右侧在下降，峰值落在 [left, mid]
			right = mid
		} else {
			// mid 比右邻居低（含相等，但相等不出现）→ 右侧在上升，峰值落在 [mid+1, right]
			left = mid + 1
		}
	}
	return left // left == right，即为某个峰值
}
```

**复杂度**

| 解法 | 时间 | 空间 | 备注 |
|---|---|---|---|
| 解法1 判断峰值 + 爬坡 | O(log n) | O(1) | 用 `getNum` 把越界统一为 -∞，逻辑直观 |
| 解法2 只比 mid 与 mid+1 | O(log n) | O(1) | 无需边界处理，写法最简，常数更小 |

> 两解法渐近相同；解法2 省去 `getNum` 闭包与单独的峰值判断，更不易写错，推荐使用。

**相邻题**

- [No.540 有序数组中的单一元素](./No.540_有序数组中的单一元素.md)：同为“在非单调数组上二分”——利用相邻对的奇偶性做单调判断，每次砍一半，与本题思想可互相对照。
- [No.074](./No.074_搜索二维矩阵.md) / [No.240](./No.240_搜索二维矩阵2.md)：经典有序 / 双序结构上的二分，结构不同但同属二分章节。
