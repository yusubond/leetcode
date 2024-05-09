# 二叉索引树

### 6 Binary Indexed Tree \[二叉索引树\]

#### 结构介绍

二叉索引树，Binary Indexed Tree，又名树状数组，或者Fenwick Tree，因为本算法由Fenwick创造。

#### 问题定义

给定一个数组array[0....n-1]，实现两个函数：1）求取前i项和；2）更新第i项值。

分析：

其中一个比较简单的解法就是通过循环累计(0-i-1)求和(需要O(n)时间)，根据下标更新值(需要O(1)时间)。另一种解法是创建一个前n项和数组(即，第i项存储前i项的和)，这样求和需要O(1)时间，更新需要O(n)时间。

而使用BIT数据结构，可以在O(Logn)时间内进行求和操作和更新操作。下面具体介绍一下。

#### 表现形式

BIT这种数据结构使用一般数组进行存储，每个节点存储 **部分** 输入数组的元素之和。BIT数组的元素个数与输入数组的元素个数一致，均为n。BIT数组的初始值均为0，通过更新操作获取输入数组部分元素之和，并存储在相应的位置上。

#### 更新操作

更新操作是指根据输入数组的元素来更新BIT数组的元素，进而根据BIT数组求前i项和。

```
update(index, val): Updates BIT for operation arr[index] += val
// Note that arr[] is not changed here.  It changes
// only BI Tree for the already made change in arr[].
1) Initialize index as index+1.
2) Do following while index is smaller than or equal to n.
...a) Add value to BITree[index]
...b) Go to parent of BITree[index].  Parent can be obtained by removing
     the last set bit from index, i.e., index = index + (index & (-index))
```

令BIT数组为：BIT[i] = Input[i-lowbit(i)+1] + Input[i-lowbit(i)+2] + ... + Input[i];

即从最左边的孩子，到自身的和，如BIT[12]= A9+A10+A11+A12

#### 求和操作

求和操作是指利用BIT数组求原来输入数组的前i项和。

```
getSum(index): Returns sum of arr[0..index]
// Returns sum of arr[0..index] using BITree[0..n].  It assumes that
// BITree[] is constructed for given array arr[0..n-1]
1) Initialize sum as 0 and index as index+1.
2) Do following while index is greater than 0.
...a) Add BITree[index] to sum
...b) Go to parent of BITree[index].  Parent can be obtained by removing
     the last set bit from index, i.e., index = index - (index & (-index))
3) Return sum.
```


说明：

1. 如果节点y是节点x的父节点，那么节点y可以由节点x通过移除Lowbit(x)获得。**Lowbit(natural)为自然数(即1,2,3…n)的二进制形式中最右边出现1的值**。即通过以下公式获得 $parent(i) = i - i & (-i)$。

2. 节点y的子节点x(对BIT数组而言)存储从y(不包括y)到x(包括x)(对输入数组而言)的的元素之和。即BIT数组的值由输入数组两者之间下标的节点对应关系获得。


计算前缀和Sum(i)的计算：顺着节点i往左走，边走边“往上爬”，把经过的BIT[i]累加起来即可。

![bit演示图](http://on64c9tla.bkt.clouddn.com/Algorithm/bit.jpg)

#### BIT工作原理

BIT工作原理得益于所有的正整数，均可以表示为多个2的幂次方之和。例如，19 = 16 + 2 + 1。BIT的每个节点都存储n个元素的总和，其中n是2的幂。例如，在上面的getSum（）的第一个图中，前12个元素的和可以通过最后4个元素（从9到12） ）加上8个元素的总和（从1到8）。 数字n的二进制表示中的设置位数为O（Logn）。 因此，我们遍历getSum（）和update（）操作中最多的O（Logn）节点。 构造的时间复杂度为O（nLogn），因为它为所有n个元素调用update（）。

#### 代码实现

```
void updateBIT(int BITree[], int n, int index, int val) {
  index = index + 1;
  while(index <= n) {
    BITree[index] += val;
    index += index & (-index);
  }
}

int getSum(int BITree[], int index) {
  int sum = 0;
  index = index + 1;
  while(index > 0) {
    sum += BITree[index];
    index -= index & (-index);
  }
  return sum;
}

int *conBIT(int arr[], int n) {
  int *BITree = new int[n + 1];
  for(int i = 0; i <= n; i++)
    BITree[i] = 0;
  for(int i = 0; i < n; i++)
    updateBIT(BITree, n, i, arr[i]);
  return BITree;
}
```
