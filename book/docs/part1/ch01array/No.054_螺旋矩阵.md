## 54 螺旋矩阵-中等

题目：

给你一个 `m` 行 `n` 列的矩阵 `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)
>
> ```
> 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
> 输出：[1,2,3,6,9,8,7,4,5]
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg)
>
> ```
> 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
> 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
> ```



**解题思路**



```go
// date 2024/01/30
// 解法1 模拟输出
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0, 16)
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, n-1
	up, down := 0, m-1
	for left <= right && up <= down {
		for j := left; j <= right; j++ {
			ans = append(ans, matrix[up][j])
		}
		up++
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// up, right 已经发生变化，需要重新判断
		if left <= right && up <= down {
			for j := right; j >= left; j-- {
				ans = append(ans, matrix[down][j])
			}
			down--
			for i := down; i >= up; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}
```

