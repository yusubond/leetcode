## Longest Common Subsequence \[最长公共子序列\]

  * 1.[问题定义](#问题定义)
  * 2.[算法分析](#算法分析)
  * 3.[代码实现](#代码实现)

### 问题定义

Given two sequences, find the length of longest subsequence present in both of them.A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous.

即，给定两个序列，找出两个序列均包含的子序列的最大长度。**子序列是指以相同的相对顺序出现但不一定相邻的序列**。

例如，"ABCDGH"和"AEDFHR"两个序列的子序列"ADH"的长度为3；"AGGTAB"和"GXTXAYB"两个序列的子序列"GTAB"的长度是4。

计算机科学中，`diff`指令就是一种最长公共子序列的应用。

### 算法分析

1）**最优子结构**

假定用 $X[0...m-1]$ 和 $Y[0...n-1]$ 分别表示长度为m和n的两个序列。$L(X[0...m-1], Y[0...n-1])$ 表示X,Y两个序列的最长公共子序列长度值。

首先考虑两个序列最后一位元素，若相匹配( $X[m-1] == Y[n-1]$ )，则具有如下公式：

$L(X[0...m-1], Y[0...n-1]) = 1 + L(X[0...m-2], Y[0...n-2])$

如果不匹配( $X[m-1] != Y[n-1]$ )，则具有如下公式：

$L(X[0...m-1], Y[0...n-1]) = MAX(L(X[0...m-1], Y[0...n-2]), L(X[0...m-2], Y[0...n-1]))$

举个例子：

对于"AGGTAB"和"GXTXAYB"而言，$L("AGGTAB","GXTXAYB") = 1 + L("AGGTA", "GXTXAY")$

对于"ABCDGH"和"AEDFHR"而言，$L("ABCDGH", "AEDFHR") = MAX(L("ABCDGH", "AEDFH"), L("ABCDG", "AEDFHR"))$

![最长公共子序列演示](http://on64c9tla.bkt.clouddn.com/Algorithm/LCS.jpg)

2) **重叠子问题**

通过上面的分析可以得出这样的结论，通过不断的求解子问题可以获得最终的解。因此，可以通过递归实现该算法。

### 代码实现

* C++

```
int lcs(char* x, char* y, int m, int n) {
  if(m == 0 | n == 0)
    return 0;
  if(x[m-1] == y[n-1])
    return 1 + lcs(x, y, m - 1, n - 1);
  else
    return max(lcs(x, y, m, n - 1), LCS(x, y, m - 1, n));
}
```

* Python

```python
def lcs(x, y, m, n):
    if m == 0 or n == 0:
        return 0;
    elif x[m-1] == y[n-1]:
        return 1 + lcs(x, y, m - 1, n - 1);
    else
        return max(lcs(x, y, m, n - 1), lcs(x, y, m - 1, n));
```
