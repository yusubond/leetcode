## 79 单词搜索-中等

题目：

给定一个 `m x n` 二维字符网格 `board` 和一个字符串单词 `word` 。如果 `word` 存在于网格中，返回 `true` ；否则，返回 `false` 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)
>
> ```
> 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
> 输出：true
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)
>
> ```
> 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
> 输出：true
> ```
>
> **示例 3：**
>
> ![img](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)
>
> ```
> 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
> 输出：false
> ```



分析：

这道题目可用DFS解决。

在地图上的任意一点，向四个方向 DFS 搜索，直到所有的单词字母都找到了就输出 ture，否则输出 false。

注意，辅助变量 visited，遍历过的元素不可再次遍历。

```go
// date 2023/12/26
var dir = [][]int{
    []int{-1, 0},  // x = x-1, y = y
    []int{0, -1},  // x = x, y = y-1
    []int{1, 0},   // x = x+1, y = y
    []int{0, 1},   // x = x, y = y+1
}

func exist(board [][]byte, word string) bool {
    n := len(board)
    m := len(board[0])
    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, m)
    }

    // 从每个点开始四个方向搜索
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if searchWord(board, visited, word, 0, i, j) {
                return true
            }
        }
    }

    return false
}

func isInBoard(board [][]byte, x, y int) bool {
    return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, idx, x, y int) bool {
    if idx == len(word)-1 {
        return board[x][y] == word[idx]
    }
    if board[x][y] == word[idx] {
        visited[x][y] = true
        // search the next
        for i := 0; i < 4; i++ {
            nx := x + dir[i][0]
            ny := y + dir[i][1]
            if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, idx+1, nx, ny) {
                return true
            }
        }

        visited[x][y] = false
    }
    return false
}
```

