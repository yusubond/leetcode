## 48 旋转图像-中等

题目：

给定一个 *n* × *n* 的二维矩阵 `matrix` 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在**[ 原地](https://baike.baidu.com/item/原地算法)** 旋转图像，这意味着你需要直接修改输入的二维矩阵。**请不要** 使用另一个矩阵来旋转图像。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg)
>
> ```
> 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
> 输出：[[7,4,1],[8,5,2],[9,6,3]]
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg)
>
> ```
> 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
> 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
> ```



**解题思路**

没啥好说的，细节题，模拟旋转。

```go
// date 2024/01/30
// 解法1 模拟
func rotate(matrix [][]int)  {
    n := len(matrix)
    left, right := 0, n-1
    up, down := 0, n-1
    for left < right && up < down {
        data := make([]int, 0, 16)
        // save data
        for j := left; j <= right; j++ {
            data = append(data, matrix[up][j])
        }
        // loop 1,4,7
        idx := right
        for i := up; i <= down; i++ {
            matrix[up][idx] = matrix[i][left]
            idx--
        }
        // loop 7,8,9
        idx = up
        for j := left; j <= right; j++ {
            matrix[idx][left] = matrix[down][j]
            idx++
        }
        // loop 9,6,3
        idx = left
        for i := down; i >= up; i-- {
            matrix[down][idx] = matrix[i][right]
            idx++
        }
        // restore data
        // 1,2,3
        idx = up
        for _, v := range data {
            matrix[idx][right] = v
            idx++
        }
        left++
        right--
        up++
        down--
    }
}
```

