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
    return fasle;
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
