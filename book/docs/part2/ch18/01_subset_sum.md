## 子集和问题

  * 1.[问题定义](#问题定义)
  * 2.[算法分析](#算法分析)
  * 3.[算法实现](#算法实现)
  * 4.[问题扩展](#问题扩展)
    * 4.1 [分区和相等](#分区和相等)
    * 4.2 [分区和差值最小](#分区和差值最小)

### 问题定义

给定一个非负整型数值集合set，和一个值sum。如果集合的某个子集的和等于这个值，则返回true，否则返回false。

例如：集合为set = {12, 13, 4, 8, 5, 10}和sum = 9，那么SUM{4, 5} = 9，因此返回true。

### 算法分析

```
SubSetSum(set, n, sum) = SubSetSum(set, n - 1, sum) || SubSetSum(set, n - 1, sum - set[n - 1])
// 初始条件
SubSetSum(set, n, sum) = true, if sum = 0;
SubSetSum(set, n, sum) = false, if sum > 0 && n = 0;
```

### 算法实现

* C++

```cpp
// 递归版
bool SubSetSum(vector<int> data, int n, int sum) {
  if(sum == 0)
    return true;
  if(sum != 0 && n == 0)
    return false;
  if(data[n-1] > sum)
    return SubSetSum(data, n - 1, sum);
  return SubSetSum(data, n - 1, sum) || SubSetSum(data, n - 1, sum - data[n - 1]);
}
```

### 问题拓展

### 分区和相等

问题1：将一个集合分成两个和相等的子集

给定一个非负整型集合，若能将其分为两个和相等的子集，则返回true，否则返回false。

### 解法分析

**问题1**

第一步，首先计算集合中元素的总和，如果和为奇数则返回false；如果和为偶数，进行第二步判断。

第二步，将上述算法中的sum，替换成sum/2即可。

```
bool SubSetSum(vector<int> data, int n, int sum) {
  if(sum == 0)
    return true;
  if(sum != 0 && n == 0)
    return false;
  if(data[n - 1] > sum)
    return SubSetSum(data, n - 1, sum);
  return SubSetEqual(data, n - 1, sum) || SubSetEqual(data, n - 1, (sum - data[n - 1]));
}
bool SubSetEqual(vector<int> data) {
  int sum = 0;
  for(int i = 0; i < data.size(); i++) {
    sum += data[i];
  }
  if((sum % 2) == 1)
    return false;
  else
    return SubSetSum(data, data.size(), sum / 2);
}
```

### 分区和差值最小

问题2：将一个集合分成两个和的差值最小的子集

给定一个非负整型集合，将其分成两个和的差值最小的子集，并返回差值。

**问题2**

设集合总和为 `sum`，把集合分成两部分 S1、S2，其中 S1 + S2 = sum。两部分的差值为：

```
|S1 - S2| = |S1 - (sum - S1)| = |2·S1 - sum|
```

要让差值最小，就要让 S1 尽可能接近 `sum/2`。因此问题转化为：**在所有可凑出的子集和中，找出不超过 `sum/2` 的最大值 S1**，最小差值即为 `sum - 2·S1`。

```cpp
int MinDiff(vector<int> data) {
  int sum = 0;
  for (int x : data) sum += x;
  int half = sum / 2;

  // dp[j] = 能否凑出和 j
  vector<bool> dp(half + 1, false);
  dp[0] = true;
  for (int v : data) {
    for (int j = half; j >= v; j--) { // 逆序：每个数只用一次
      dp[j] = dp[j] || dp[j - v];
    }
  }
  // 找不超过 half 的最大可凑出和
  int s1 = 0;
  for (int j = half; j >= 0; j--) {
    if (dp[j]) { s1 = j; break; }
  }
  return sum - 2 * s1;
}
```

例如 `data = {1, 6, 11, 5}`，sum = 23，half = 11。可凑出的、不超过 11 的最大和是 11（子集 `{11}`），最小差值 = 23 − 2×11 = **1**（`{11}` 与 `{1, 6, 5}` 的和分别为 11 与 12）。

> 与问题1「分区和相等」对照：问题1 问"能否恰好平分"（判断 `dp[half]` 是否为 true），问题2 问"最接近平分"（在 `dp[]` 中取 ≤ half 的最大 true）。两者共享同一张子集和表，只是答案的取法不同。
