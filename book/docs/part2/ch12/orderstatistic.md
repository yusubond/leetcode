## Order Statistc [顺序统计]

  * 1.[问题定义](#问题定义)
  * 2.[算法分析](#算法分析)
    * 2.1[一般解法](#一般解法)
    * 2.2[随机快速排序](#随机快速排序)
    * 2.3[最优解法](#最优解法)

### 问题定义

```
K’th Smallest/Largest Element in Unsorted Array.
即，给定一组未排序序列，输出其中第K个最大(最小)元素。
```
该问题是对排序算法的延伸和扩展。

### 算法分析

对于这个问题，其一般思路就是先对数据进行排序，然后输出对应的Kth最大(最小)值。因此，排序算法是这个问题的关键。

### 一般解法

一个比较简单的解法，就是选取一个复杂度为O(nLogn)的排序算法(比如快速排序，归并排序，堆排序等)对数据进行排序，然后输出第k个元素。

时间复杂度为：O(nLogn)

### 随机快速排序

快速排序中基准值的选取有三种方式，即首元素，尾元素，和随机元素。因此，我们选择随机选取基准的方式，实现快速排序中的分区部分，具体实现如下：

* C++

```
int randpartitioin(vector<int> &data, int low, int high) {
  int n = high - low + 1;
  int pivot = rand() % n;
  int temp = data[pivot];
  while(low < high) {
    while(low < high && data[high] >= temp) high--;
    data[pivot] = data[high];
    while(low < high && data[low] <= temp) low++;
    data[high] = data[low];
    data[low] = temp;
  }
  return low;
}
int KthSmallest(vector<int> &data, int low, int high, int k) {
  if(k > 0 && k <= high - low + 1) {
    int pos = randpartitioin(data, low, high);
    if(pos - low == k - 1)
      return data[pos];
    if(pos - low > k - 1)
      return KthSmallest(data, low, pos - 1, k);
    else
      return KthSmallest(data, pos, high, k - pos + low -1);
  } else {
    return -1;
  }
}
```

最坏的情况下，时间复杂度为O(nLogn)，即每次随机选取的基准值刚好是序列中的最值。一般情况下，其 **预期时间复杂度为O(n)**。

### 最优解法

在解法二的基础上，考虑基准值选取的平衡性，即每次选出的基准值尽量保证左右两个子序列元素个数的相差不多。

这样的解法，即是最坏的情况下，时间复杂度也是O(n)。
