## Longest Increasing Subsequence \[最长递增子序列\]

  * 1.[问题定义](#问题定义)
  * 2.[算法分析与实现](#算法分析与实现)
  * 3.[问题延伸](#问题延伸)

### 问题定义

The Longest Increasing Subsequence (LIS) problem is to find the length of the longest subsequence of a given sequence such that all elements of the subsequence are sorted in increasing order.

即，给定一个序列，输出其所有子序列中满足递增关系的最长子序列的长度。**元素不一定要连续**。

### 算法分析与实现

1) **最优子结构**

令 $L(i)$ 表示以索引i所在位置的元素arr[i]为结尾元素的LIS的长度，其中arr[i]为输入序列。则具有如下公式：

$L(i) = 1 + max(L(j))$, 当 $0 < j < i$ , 并且 $arr[j] < arr[i]$ ;

$L(i) = 1$ , 当满足上述条件的 $j$ 不在时。

```cpp
// 复杂度为O(n^2)
int LIS(vector<int> data) {
  vector<int> num(data.size());
  int res = 1;;
  fill(num.begin(), num.end(), 1);
  for(int i = 1; i < data.size(); i++) {
    for(int j = 0; j < i; j++) {
      if(data[i] > data[j] && num[i] < num[j] + 1) {
        num[i] = num[j] + 1;
        if(num[i] > res)
          res = num[i];
      }
    }
  }
  return res;
}
// 复杂度为O(nLogn)
#include <iostream>
#include <vector>

// Binary search (note boundaries in the caller)
int CeilIndex(std::vector<int> &v, int l, int r, int key) {
  while (r-l > 1) {
    int m = l + (r-l)/2;
    if (v[m] >= key)
      r = m;
    else
      l = m;
    }
  return r;
}

int LongestIncreasingSubsequenceLength(std::vector<int> &v) {
  if (v.size() == 0)
    return 0;
  std::vector<int> tail(v.size(), 0);
  int length = 1; // always points empty slot in tail
  tail[0] = v[0];
  for (size_t i = 1; i < v.size(); i++) {
    if (v[i] < tail[0])
      tail[0] = v[i];
    else if (v[i] > tail[length-1])
      tail[length++] = v[i];
    else
      tail[CeilIndex(tail, -1, length-1, v[i])] = v[i];
  }
  return length;
}
```

其实现过程，可以通过一个例子说明。

![最长递增子序列](http://on64c9tla.bkt.clouddn.com/Algorithm/lis_1.png)
![最长递增子序列](http://on64c9tla.bkt.clouddn.com/Algorithm/lis_2.png)
![最长递增子序列](http://on64c9tla.bkt.clouddn.com/Algorithm/lis_3.png)
![最长递增子序列](http://on64c9tla.bkt.clouddn.com/Algorithm/lis_4.png)

值得注意的是，根据LIS衍生出来的最长非递减序列，最长严格递减序列，其道理是一样的。

### 问题延伸

**最长连续递增子序列**

* C++

```cpp
int lis(vector<int> data) {
  int result = -1, temp = 1;
  for(int i = 1; i < data.size(); i++) {
    if(data[i-1] < data[i]) {
      temp++;
      if(temp > result)
        result = temp;
    } else {
      temp = 1;
    }
  }
  return result;
}
```
