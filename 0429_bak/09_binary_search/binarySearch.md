## Binary Search \[二分查找\]

  * 1.[问题定义](#问题定义)
  * 2.[复杂度分析](#复杂度分析)
  * 3.[算法流程](#算法流程)
  * 4.[代码实现](#代码实现)
    * 4.1[递归实现](#二分查找的递归实现)
    * 4.2[迭代实现](#二分查找的迭代实现)

### 问题定义

给定一个具有n个元素的有序序列，实现一个能够查找元素x的函数。

### 复杂度分析

最简单的方法是**线性搜索**，即依次比较每个元素，其时间复杂度为O(n)。而二分查找利用序列有序的特性，将时间复杂度降为O(Logn)。

### 算法流程

二分查找是指在已排序的序列中查看数据，每次查找均将排序序列分为一半的有序序列。查找的初始为整个序列，如果搜索的关键字小于或等于序列的中间值，则对前半序列继续搜索；否则对后半序列继续搜索。

在每一次的比较中，几乎可以筛选出一半的元素是否符合预期值。

1. 比较x与序列中间值
2. 如果x等于中间值，则返回中间值的索引
3. 如果x大于中间值，则x落在中间值的右半个子序列，进而对右半子序列进行查找
4. 如果x小于中间值，则x落在中间值的左半个子序列，进而对左半子序列进行查找

### 代码实现

#### 二分查找的递归实现

* C++

```cpp
int BinarySearch(int arry[], int l, int r, int x) {
    if(l <= r) {
        int mid = l + (r - 1) / 2;
        if(x == arry[mid])
            return mid;
        if(x < arry[mid])
            return BinarySearch(arry, l, mid - 1, x);
        else
            return BinarySearch(arry, mid + 1, r, x);
    }
    return -1;
}
```

* Python

```python
def BinarySearch(arry, l, r, x)
    if l <= r:
        mid = l + (r - 1)/2
        if x == arry[mid]:
            return mid
        if x < arry[mid]:
            return BinarySearch(arry, l, mid - 1, x)
        else:
            return BinayrSearch(arry, mid + 1, r, x)
    else:
        return -1
```

### 二分查找的迭代实现

* C++

```cpp
int BinarySearch(int arry[], int l, int r, int x) {
    while(l <= r) {
        if(x == arry[mid])
            return mid;
        if(x < arry[mid])
            r = mid - 1;
        else
            l = mid + 1;
    }
    return -1;
}
```

* Python

```python
def BinarySearch(arry, l, r, x)
    while l <= r:
        mid = l + (r - 1)/2
        if x== arry[mid]:
            return mid
        elif x < arry[mid]:
            r = mid - 1
        else:
            l = mid + 1
    return -1
```
