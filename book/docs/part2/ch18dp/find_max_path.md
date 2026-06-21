## 二维空间中的最长递增子序列

  * 1.[问题定义](#问题定义)
  * 2.[解题思路](#解题思路)
  * 3.[算法实现](#算法实现)
  * 4.[问题延伸](#问题延伸)

### 问题定义

给定一个N * N的二维整型矩阵，每个元素可以访问前后左右四个元素，找出其中包含严格递增序列的最大节点个数。

![二维空间中严格递增序列的最大节点个数](http://on64c9tla.bkt.clouddn.com/Algorithm/find_max_path.png)

例如，如图所示，符合严格递增序列的最大节点个数分别为16，12，10，相应的路径由红色实线给出。

### 解题思路

1. 对比最短路径问题中保证边的权值最小的问题，利用二维矩阵的坐标保存对应元素的可获得的(满足严格递增序列的)最大节点个数。

2. 利用动态规划思想不断通过遍历节点，并更新每个节点的值，并进行边界条件的设定。

3. 边界条件的设定。当每次遍历后可获得全局的最大值不在发生变化后，即达到收敛条件。

### 算法实现

```cpp
int FindMaxPath(int data[][4], int n) {
  int num[n][n];
  fill(num[0], num[0] + n * n, 1);
  int maxnum = -1, temp = -1;
  int tempnum = 0;
  while(1) {
    // 遍历所有节点，并更新每个节点num[i][j]值
    for(int i = 0; i < n; i++) {
      // 只进行前后比较即可
      if(i == n-1) {
        for(int j = 0; j < n - 1; j++) {
          if(data[i][j] > data[i][j+1] && num[i][j] <= num[i][j+1])
            num[i][j] = num[i][j+1] + 1;
          if(data[i][j] < data[i][j+1] && num[i][j] >= num[i][j+1])
            num[i][j+1] = num[i][j] + 1;
          tempnum = num[i][j] > num[i][j+1] ? num[i][j] : num[i][j+1];
          temp = tempnum >= temp ? tempnum : temp;
        }
      } else {
      // 进行前后和上下比较
      for(int j = 0; j < n; j++) {
        if(j == n-1) {
          if(data[i][j] > data[i+1][j] && num[i][j] <= num[i+1][j])
            num[i][j] = num[i+1][j] + 1;
          if(data[i][j] < data[i+1][j] && num[i][j] >= num[i+1][j])
            num[i+1][j] = num[i][j] + 1;
          tempnum = num[i][j] > num[i+1][j] ? num[i][j] : num[i+1][j];
          temp = tempnum >= temp ? tempnum : temp;
        } else {
          if(data[i][j] > data[i+1][j] && num[i][j] <= num[i+1][j])
            num[i][j] = num[i+1][j] + 1;
          if(data[i][j] < data[i+1][j] && num[i][j] >= num[i+1][j])
            num[i+1][j] = num[i][j] + 1;
          if(data[i][j] > data[i][j+1] && num[i][j] <= num[i][j+1])
            num[i][j] = num[i][j+1] + 1;
          if(data[i][j] < data[i][j+1] && num[i][j] >= num[i][j+1])
            num[i][j+1] = num[i][j] + 1;
          tempnum = find(num[i][j], num[i+1][j], num[i][j+1]);
          temp = tempnum >= temp ? tempnum : temp;
        }
      }
    }
  }
  // 边界条件的判断
  if(temp == maxnum) {
    return maxnum;
    break;
  }
  // 更新全局最大值
  if(temp > maxnum)
    maxnum = temp;
}
}
```

### 问题延伸

其实，这个问题是对[最长递增子序列](longestincreasingsub.md)的延伸，要深刻理解其算法思想。同样，这个算法的思想可以拓展到三维数组，只不过每个节点需要比较的节点个数变为最少三个，最多六个，但是其思路都一样。
