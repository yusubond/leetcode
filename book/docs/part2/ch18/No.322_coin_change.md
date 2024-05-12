## 322 Coin Change \[零钱兑换\]

  * 1.[问题定义](#问题定义)
  * 2.[最优子结构解法](#最优子结构解法)
  * 3.[重叠子问题解法](#重叠子问题解法)

### 问题定义

给定一个值M，和一组硬币面值的集合，问将硬币组合值为M的情况一共有几种，不考虑硬币的顺序问题。

例如，值为4，硬币面值为{1,2,3}。那么，一共有四种情况。{1,1,1,1}、{1,1,2}、{1,3}和{2,2}。

变相问题：

给定一个int数组A，数组中元素互不重复，给定一个数x，求所有求和能得到x的数字组合，组合中的元素来自A，可重复使用 

### 最优子结构解法

```
L(set, n, sum)表示（值为sum，硬币面值集合为set, n为硬币面值的个数）硬币的组合情况数量
L(set, n, sum) = L(set, n - 1, sum)   // 不包含面值为set[n]的硬币
               + L(set, n, sum - set[n - 1]); //至少包含一个面值为set[n-1]的硬币

边界条件，即递归结束条件
if(sum == 0)
  return 1
if(sum < 0)
  return 0;
if(n <= 0 && sum >= 1)
  return 0;
```

```cpp
// 简易版
int CoinChange(int coin[], int n, int sum) {
  if(sum == 0)
    return 1;
  if(sum < 0)
    return 0;
  if(n <= 0 && sum > 0)
    return 0;
  return CoinChange(coin, n - 1, sum) + CoinChange(coin, n, sum - coin[n - 1]);
}
```

![硬币兑换](http://on64c9tla.bkt.clouddn.com/Algorithm/coin_change.png)

上面的分析中，我们可以看到在递归地过程中有些值被重复计算了多次，因此，硬币兑换问题同时具有 **最优子结构** 和 **重叠子问题** 两种属性，据此，可对算法进行优化。

### 重叠子问题解法

对于set={3, 4, 7}和sum = 16的情况而言，sum=13/12/9的情况一样，无非是16-set[i]。因此，为了避免重复计算，将每一次计算的结果存储在table[sum+1][n]中，横坐标视为满足当前sum值的子问题结果，而纵坐标体现了集合中元素的包含关系。

```cpp
int CoinChange(int s[], int n, int sum) {
  int i, j, x, y;
  int table[sum][n];
  for(j = 0; j < n; j++)
    table[0][j] = 1;
  for(i = 1; i <= sum; i++) {
    for(j = 0; j < n; j++) {
      // 包含至少一个s[j]面值的硬币
      x = (i - s[j]) >= 0 ? table[i - s[j]][j] : 0;
      // 不包含s[j]面值的硬币
      y = (y >= 1) ? table[i][j - 1] : 0;
      table[i][j] = x + y;
    }
  }
  return table[sum][n - 1];
}
```
