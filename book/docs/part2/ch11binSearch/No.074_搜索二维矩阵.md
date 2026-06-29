## 74 搜索二维矩阵-中等

题目：

给你一个满足下述两条属性的 `m x n` 整数矩阵：

- 每行中的整数从左到右按非严格递增顺序排列。
- 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target` ，如果 `target` 在矩阵中，返回 `true` ；否则，返回 `false` 。

> **示例 1：**
>
> ```
> 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
> 输出：true
> ```
>
> **示例 2：**
>
> ```
> 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
> 输出：false
> ```

**关键观察**

两个条件合起来意味着：每行内部有序，且行与行首尾相接（行首 > 上行行尾），所以把所有行首尾拼接后仍是**全局有序**的。于是整张矩阵可“展平”为一个长度 `m·n` 的有序一维数组，问题退化为一维数组的二分查找。

> 题目为“非严格递增”（可能有重复）。对于“判断 target 是否存在”，含重复的有序数组二分依然成立，故两解法均适用。

**解法 1：两次二分（先定位行，再行内查找）**

先对每行首元素做 LastEqualOrSmaller 二分，找到唯一可能包含 target 的行；再在该行内做 FirstEqual 二分。

```go
// date 2023/12/04
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    up, down := 0, m-1
    // find the last equal or smaller
    for up <= down {
        mid := (up + down) / 2
        if matrix[mid][0] > target {
            down = mid - 1
        } else {
            up = mid + 1
        }
    }
    // find 
    if down >= 0 && down < m {
        // find col
        n := len(matrix[down])
        left, right := 0, n-1
        for left <= right {
            mid := (left + right) / 2
            if matrix[down][mid] == target {
                return true
            } else if matrix[down][mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return false
}
```

**解法 2：展平一维二分（推荐）**

把矩阵当作长度 `m·n` 的一维有序数组，一轮二分即可。一维下标 `mid` 与二维坐标的映射：`matrix[mid/n][mid%n]`。

```go
// date 2026/06/29
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    lo, hi := 0, m*n-1
    for lo <= hi {
        mid := lo + (hi-lo)/2 // 用 lo+(hi-lo)/2 防整数溢出
        v := matrix[mid/n][mid%n]
        if v == target {
            return true
        } else if v < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return false
}
```

**复杂度**

| 解法 | 时间 | 空间 | 备注 |
|---|---|---|---|
| 解法1 两次二分 | O(log m + log n) | O(1) | 先定位行再行内查找，逻辑直观 |
| 解法2 展平一维二分 | O(log(m·n)) | O(1) | 一轮二分，写法最简，常数更小 |

> 注：`log(m·n) = log m + log n`，两解法渐近相同；解法2 只有一次循环，常数因子更优。

**与 LC 240 的区别**

[No.240 搜索二维矩阵 II](./No.240_搜索二维矩阵2.md) 只保证“行升序、列升序”，但**行首不一定大于上行行尾**，不能展平成一维，需改用从右上角出发的 Z 字（阶梯）查找。两题条件不同，解法不同，注意区分。
