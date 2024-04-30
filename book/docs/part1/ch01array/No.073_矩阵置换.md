

## 73 矩阵置换-中等

题目：

给定一个 `*m* x *n*` 的矩阵，如果一个元素为 **0** ，则将其所在行和列的所有元素都设为 **0** 。请使用 **[原地](http://baike.baidu.com/item/原地算法)** 算法。

- 一个直观的解决方案是使用  `O(*m**n*)` 的额外空间，但这并不是一个好的解决方案。
- 一个简单的改进方案是使用 `O(*m* + *n*)` 的额外空间，但这仍然不是最好的解决方案。
- 你能想出一个仅使用常量空间的解决方案吗？



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg)
>
> ```
> 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
> 输出：[[1,0,1],[0,0,0],[1,0,1]]
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg)
>
> ```
> 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
> 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
> ```



**解题思路**

- 最朴素的算法，详见解法1。因为一行或一列中的零有可能重复，所以可以先进行一次遍历，把为零的行列信息保存在哈希表中，然后遍历哈希表，就地置换。该算法使用`O(M+N)`的额外空间。

- 第0行，第0列，特殊处理，详见解法2。具体为，先遍历第0行和第0列，用两个变量标记是否有零；然后整个遍历矩阵，把行列为零的信息保存到第0行和第0列；然后根据第0行第0列的信息置零；最后，在根据两个变量把第0行第0列置零。

  



```go
// date 2024/01/30
// 解法1 哈希表
func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    rset := make(map[int]int, 16)
    cset := make(map[int]int, 16)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rset[i] = i
                cset[j] = j
            }
        }
    }
    for row := range rset {
        for j := 0; j < n; j++ {
            matrix[row][j] = 0
        }
    }
    for col := range cset {
        for i := 0; i < m; i++ {
            matrix[i][col] = 0
        }
    }
}

// 解法2
// 两个变量，0行0列特殊处理
// 时间复杂度 O(MN)
// 空间复杂度 O(1)
func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    row0, col0 := false, false
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            row0 = true
        }
    }
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            col0 = true
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }


    if row0 {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if col0 {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
```

