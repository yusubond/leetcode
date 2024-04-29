## 广度优先搜索

[TOC]



### 知识点回顾



### 相关题目

#### 695 岛屿的最大面积

题目要求：

思路分析：

```go
// date 2020/03/29
// 广度优先搜索，注意使用visited变量，可以快速搜索
func maxAreaOfIsland(grid [][]int) int {
    n := len(grid)
    if n < 1 { return 0 }
    m := len(grid[0])
    res := 0
    var visited [50][50]int
    // 广度优先搜索，找到岛屿的最大值
    var findLandArea func(x, y int) int
    findLandArea = func(x, y int) int {
        area := 1
        dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
        queue := make([][2]int, 0)
        queue = append(queue, [2]int{x, y})
        visited[x][y] = 1
        for len(queue) != 0 {
            l := len(queue)
            for i := 0; i < l; i++ {
                for j := 0; j < 4; j++ {
                    nx, ny := queue[i][0]+dx[j], queue[i][1]+dy[j]
                    if nx < 0 || nx >=n || ny < 0 || ny >= m { continue }
                    if grid[nx][ny] == 1 && visited[nx][ny] != 1 {
                        queue = append(queue, [2]int{nx, ny})
                        visited[nx][ny] = 1
                    }
                }
            }
            queue = queue[l:]
            area += len(queue)
        }
        return area
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 && visited[i][j] != 1 {
                area := findLandArea(i, j)
                if res < area { res = area }
            }
        }
    }
    return res
}
```

#### 1162 地图分析

题目要求：https://leetcode-cn.com/problems/as-far-from-land-as-possible/

思路分析：

```go
// date 2020/03/29
func maxDistance(grid [][]int) int {
  // 第一步，找到某个海洋距离最近的陆地
  dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
  var findNearestLand func(x, y int) int
  findNearestLand = func(x, y int) int {
    queue := make([][3]int, 0)   // 0, 1, 2分别是横坐标，纵坐标，和距离
    vis := [100][100]int{}       // 每一次搜索都需要初始化
    queue = append(queue, [3]int{x, y, 0})
    vis[x][y] = 1
    nx, ny := 0, 0
    for len(queue) != 0 {
      l := len(queue)
      for i := 0; i < l; i++ {
        for j := 0; j < 4; j++ {
          nx, ny = queue[i][0] + dx[j], queue[i][1] + dy[j]
          if nx < 0 || nx >= n || ny < 0 || ny >= n { continue }
          if vis[nx][ny] == 1 { continue }
          queue = append(queue, [3]int{nx, ny, queue[i][2]+1})
          vis[nx][ny] = 1
          if grid[nx][ny] == 1 { return queue[i][2]+1 }
        }
      }
      queue = queue[l:]
    }
    return -1
  }

  res, n := -1, len(grid)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      if grid[i][j] == 0 {
        res = max(res, findNearestLand(i, j))
      }
    }
  }
  return res
}

func max(x, y int) int {
    if x > y { return x }
    return y
}
```

