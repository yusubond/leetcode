## Min Cost Path \[最小成本路径\]

  * 1.[问题定义](#问题定义)
  * 2.[算法分析](#算法分析)
  * 3.[算法实现](#算法实现)

### 问题定义

给定一个包含m*n的整型矩阵，每个节点上的值代表路径成本，求从(0,0)节点出发到(x,y)节点（0<=x<m,0<=y<n)的最小路径，规定节点只能向下走，即节点(i,j)只能到达(i+1,j),(i,j+1),(i+1,j+1)。

如图所示，节点(2,2)的最小成本路径为8。

![最小成本路径](http://on64c9tla.bkt.clouddn.com/Algorithm/min_cost_path.png)

### 算法分析

从规定的节点走向可以得出：进入节点(m,n)，只能从(m-1,n)，(m-1,n-1)和(m,n-1)之一进入；那么求出到达这三个节点的最小成本路径，进行递归即可。

```
MinCost[m][n]表示从原点到节点(m,n)的最小成本路径，则
MinCost[m][n] = min(MinCost(m-1,n), MinCost(m-1,n-1), MinCost(m,n-1)) + cost[m][n];

初始条件，即边界条件：
// 第一行
for(int i = 1; i < n; i++)
  mincost[0][i] = mincost[0][i - 1] + cost[0][i];
// 第一列
for(int i = 1; i < m; i++)
  mincost[i][0] = mincost[i - 1][0] + cost[i][0];
// 起点
mincost[0][0] = cost[0][0]
```

### 算法实现

```cpp
int MinFuc(int x, int y, int z) {
  return return z < (x < y ? x : y) ? z : (x < y ? x : y);
}
int MinCost(int cost[M][N], int m, int n) {
  int res[M][N];
  res[0][0] = cost[0][0];
  if(m == 0 && n == 0)
    return res[0][0];
  int i, j;
  for(i = 1; i < M; i++)
    res[i][0] = res[i-1][0] + cost[i][0];
  for(j = 1; j < N; j++)
    res[0][j] = res[0][j-1] + cost[0][j];
  for(i = 1; i <= m; i++)
    for(j = 1; j <= n; j++)
      res[i][j] = MinFuc(res[i-1][j], res[i-1][j-1], res[i][j-1]) + cost[i][j];
  return res[m][n];
}
```
