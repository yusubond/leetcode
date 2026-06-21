## 最小编辑距离

  * 1.[问题定义](#问题定义)
  * 2.[算法分析](#算法分析)
  * 3.[代码实现](#代码实现)

### 问题定义

给定两个字符串str1和str2，以及三种操作 *插入、删除和替换*, 问str1变为str2所需的最小操作次数。

例如，str1 = sunday, str2 = saturday, 最小操作次数为3，即插入a和t，并将n替换为r。

### 算法分析

对于这种问题，首先应该想到如何将大问题化解为小问题，即如果一两个字母该如何处理。然后运用递归思想不断地进行求解。

```
令L(s1, s2, n, m)表示字符串s1变成s2所需要的最少操作数，n和m为两个字符串的长度，那么则有以下公式：

当s1[n-1] = s2[m-1]，L(s1, s2, n, m) = L(s1, s2, n-1, m-1);
当s1[n-1] != s2[m-1]，三种操作均可以将其变成相等的情况，取三个操作的最小值即可
    L(s1, s2, n, m) = 1 + min(L(s1, s2, n, m - 1),        //插入
                              L(s1, s2, n - 1, m),        //删除
                              L(s1, s2, n - 1, m - 1));   //替换
递归结束条件，即基本条件
当m=0时，需要全部删除，return n;
当n=0时，需要全部插入，return m;
```

### 代码实现

```cpp
int MinFuc(int x, int y, int z) {
  return z < (x < y ? x : y) ? z : (x < y ? x : y);
}
int EditDistance(string s1, string s2, int n, int m) {
  if(n == 0)
    return m;
  if(m == 0)
    return n;
  if(s1[n - 1] == s2[m - 1])
    return EditDistance(s1, s2, n - 1, m - 1);
  else
    // 插入、删除，替换
    return 1 + MinFuc(EditDistance(s1, s2, n, m - 1), EditDistance(s1, s2, n - 1, m), EditDistance(s1, s2, n - 1, m - 1));
}
```
