## 排序算法

### Quick Sort  \[快速排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[复杂度分析](#复杂度分析)
  * 3.[算法流程](#算法流程)

### 算法介绍

快速排序与归并排序均属于分治思想算法(Divide and Conquer algorithm)，即选取一个基准值(pivot)，然后将其他给定值置于基准值的左右，从而实现排序过程。基准值的选取可以有多种方式：

1. 选取第一个元素作为基准值
2. 选取最后一个元素作为基准值
3. 随机选取一个元素作为基准值
4. 选取中间值作为基准值

快速排序算法的关键是partition()函数，其实现的目标是在给定的序列和基准值x的情况下，将基准值x放在正确的位置上，所有小于x的元素放在x的左边，所有大于等于x的元素放在x的右边。

### 复杂度分析

通常情况下，快速排序的时间复杂度可用如下公式表示：

$T(n) = T(k) + T(n-k-1) + \theta(n)$

其中，前两项表示递归调用，第三项表示分区过程(即，partition)。k代表序列中小于等于基准的元素个数。因此，快速排序的时间复杂度与 **待排序列** 和 **基准值的选取** 有关。下面分三种情况分别讨论：

* 最差的情况

  当分区过程总是选取最大或最小的元素作为基准值时，则发生最差的情况。此时，待排序列实际上已经是有序序列(递增，或递减)。因此，上述时间复杂度公式变为：

  $T(n) = T(0) + T(n-1) + \theta(n)$，即时间复杂度为$\theta(n^2)$

* 最好的情况

  当分区过程每次都能选出中间值作为基准值时，则发生最好的情况，此时，时间复杂度公式为：

  $T(n) = 2T(n/2) + \theta(n)$，即时间复杂度为 $\theta(nLogn)$

* 一般情况

  其时间复杂度为 $\theta(nLogn)$

### 算法流程

递归形式的快速排序的伪代码如下：

```golang
// low-起始索引；high-结束索引
quicksort(arry[], low, high) {
  if(low < high) {
    // pi是partition()产生的基准值的索引
    pi = partition(arry, low, high);
    quicksort(arry, low, pi - 1);
    quicksort(arry, pi + 1, high);
  }
}
```

partition()函数实现序列的划分，其逻辑是总是选取最左边的元素作为基准值x(也可以选取其他元素最为基准值)，然后追踪小于基准值x的索引i；遍历整个序列，若当前元素小于基准值x，则交换当前值值与i所在的值，否则忽略当前值。其伪代码如下：

```golang
// 函数选取最后一个值作为基准值
// 将所有小于等于基准值的元素放在基准值的左边，将所有大于基准值的元素放在基准值的右边
// index表示元素应该被放在的位置,for循环之后区间[left, index]均小于基准值,区间[index+1, right]均大于基准值,再交换一次swap(index+1, right)
// 则index+1为基准值,区间[index+2, right]为大于等于基准值
// 最后返回基准值的位置index+1
partition(nums []int, left, right int) {
  p := nums[right]
  index := left - 1
  for left <= right-1 {
    if nums[left] <= p {
      index++
      swap(nums, index, left)
    }
    left++
  }
  swap(nums, index+1, rigth)
  return index+1
}
```

```golang
// 另一种实现方式
func quickSort(nums []int, left, right int) {
  p := division(nums, left, right)
  quickSort(nums, left, p-1)
  quickSort(nums, p+1, right)
}

func division(nums []int, left, right int) int {
  base := nums[left]
  for left < right {
    for left < right && nums[right] >= base { right-- }
    nums[left] = nums[right]
    for left < right && nums[left] <= base { left++ }
    nums[right] = nums[left]
    nums[left] = base
  }
  return left
}
```

### Merge Sort \[归并排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[代码实现](#代码实现)
  * 3.[复杂度分析](#复杂度分析)
  * 4.[归并算法的应用](#归并算法的应用)

### 算法介绍

与快速排序一样，归并排序也是采用分治思想。将待排序列分成两个子序列(直到不可再分)，然后对两个子序列进行有序合并。merge()函数是归并排序的关键。

![归并排序过程演示](http://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/Merge-Sort.png)

### 代码实现

根据归并算法的原理，很容易想到利用递归的方式进行实现，如下所示：

* C++

```cpp
// 将两个有序序列合并成一个序列
// 需要使用辅助空间
void merge(vector<int> &data, int first, int mid, int last) {
  int i = first, j = mid + 1;
  vector<int> temp;
  while(i <= mid && j <= last) {
    if(data[i] <= data[j])
      temp.push_back(data[i++]);
    else
      temp.push_back(data[j++]);
  }
  while(i <= mid)
    temp.push_back(data[i++]);
  while(j <= last)
    temp.push_back(data[j++]);
  for(int k = 0; k < temp.size(); k++)
    data[first + k] = temp[k];
  return;
}
void mergesort(vector<int> &data, int first, int last) {
  if(first < last) {
    int mid = first + (last - first) / 2;
    mergesort(data, first, mid);
    mergesort(data, mid + 1, last);
    mergedata(data, first, mid, last);
  }
  return;
}
```

* python

```python
def merge(left,right):
    result=[]
    i,j=0,0
    while i<len(left) and j<len(right):
        if left[i]<=right[j]:
            result.append(left[i])
            i+=1
        else:
            result.append(right[j])
            j+=1
    result+=left[i:]
    result+=right[j:]
    return result

def mergesort(arr):
    if len(arr)<=1:
        return arr
    mid=int(len(arr)/2)
    left=mergesort(arr[:mid])
    right=mergesort(arr[mid:])
    return merge(left,right)
```

### 复杂度分析

归并排序属于递归算法，其时间复杂度可用如下公式表示：

$T(n) = 2T(n/2) + \theta(n)$

因为归并排序总是将序列分成两个子序列，直到不可分解，然后在 **线性时间** 内进行合并。所以三种情况(最好，最坏，一般)的时间复杂度均是 $\theta(nLogn)$ 。

从稳定性而言，归并排序比快速排序稳定。

辅助空间需要O(n)。

### 归并算法的应用

* 线性表的归并排序

  对于线性表而言，其与数组的主要区别在于并不需要申请内存空间，可以在O(1)的时间内和O(1)的空间内实现元素的插入。因此，对线性表进行归并排序，合并数据的过程不需要额外的辅助空间。

### 相关题目

- 252 会议室
- 280 摆动排序【M】
- 56 合并区间【M】
- 912 排序数组【M】

- 1356 根据数字二进制下1的数目排序【E】

#### 252 会议室

思路分析

算法：直接对区间进行排序，判断是否有交叉

```go
func canAttendMeetings(intervals [][]int) bool {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0], intervals[j][0]
  })
  for i := 0; i < len(intervals) - 1; i++ {
    if intervals[i][1] > intervals[i+1][0] {
      return false
    }
  }
  return true
}
```

#### 280 摆动排序

题目要求：给定数组，要求使其奇数位上的元素都大于其两边的元素值。

思路分析：遍历一遍，逐个纠正。

```go
// date 2020/02/28
func wiggleSort(nums []int) {
  var t int
  for i := 0; i < len(nums)-1; i++ {
    if (i & 0x1 == 0) == (nums[i] > nums[i+1]) {
      t = nums[i+1]
      nums[i+1] = nums[i]
      nums[i] = t
    }
  }
}
```



#### 56 合并区间

思路分析

算法：先对区间进行排序；维护start, end值，注意处理最后一个

```go
func merge(intervals [][]int) [][]int {
  if len(intervals) <= 1 {return intervals}
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  res, i, n := make([][]int, 0), 0, len(intervals) // i保存当前start索引
  // init start, end
  start, end := intervals[0][0], intervals[0][1]
  for i < n {
    i++
    // 查找以start开始的最大区间，并更新end值
    for i < n {
      if intervals[i][1] <= end || intervals[i][0] <= end {
        if intervals[i][1] > end { end = intervals[i][1] }
        i++
      } else { break }
    }
    t := make([]int, 2)
    t[0], t[1] = start, end
    res = append(res, t)
    // i >= n 说明处理完毕，直接退出
    if i >= n {break}
    // update start, end
    start, end = intervals[i][0], intervals[i][1]
  }
}
```

#### 912 排序数组

##### 快速排序

快速排序的思想是分治算法，partition将数组分成两个子数组，如果两个子数组有序，则这个数组有序。

```go
// date 2020/03/01
func sortArray(nums []int) []int {
    if len(nums) == 1 { return nums }
    QuickSort(nums, 0, len(nums)-1)
    return nums
}

// 快速排序
// 返回值为index
func partition(nums []int, l, r int) int {
  p := nums[l]
  index := l
  for i := l; i <= r; i++ {
    if nums[i] <= p {
      swap(nums, index, i)
      index++
    }
  }
  swap(nums, l, index-1)
  return index-1
}

func QuickSort(nums []int, l, r int) {
  if l >= r { return }
  p := partition(nums, l, r)
  QuickSort(nums, l, p-1)
  QuickSort(nums, p+1, r)
}
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
// 第二种写法
// partition选最右边的值作为基准值
func partition(nums []int, left, right int) int {
  // index维护小于基准值的下标
  index, p := left-1, nums[right]
  for left < right {
    if nums[left] < p {
      index++
      swap(nums, index, left)
    }
    left++
  }
  // 将基准值放在index+1上
  swap(nums, index+1, right)
  return index+1
}
```



##### 选择排序

思路分析：选择未排序数组中最小的一个，然后交换值应该所在的位置。时间复杂度O(n^2) 空间复杂度O(1)

```go
func selectSort(nums []int) {
  n := len(nums)
  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      if nums[i] > nums[j] {
        swap(nums, i, j)
      }
    }
  }
}
func swap(nums []int, i, j int) {
  nums[i], nums[j] = nums[j], nums[i]
}
```

##### 插入排序

算法逻辑：一个元素肯定是有序的，然后依次遍历剩下的元素，将其插入到前面已排序的数组的中。

时间复杂度O(n^2) 空间复杂度O(1)

```go
func insertSort(nums []int) {
  temp, k := 0, 0
  for i := 1; i < len(nums); i++ {
    temp, k = nums[i], i-1
    for k > 0 && nums[k] > temp {
      nums[k+1] = nums[k]
      k--
    }
    nums[k+1] = temp
  }
}
```

##### 希尔排序

算法思想

插入排序的思想是使数组相邻的两个元素有序，而希尔排序是使数组中任意间隔为h的元素有序，然后h不断减小，又称为插入排序的快速排序。

```go
func sortArray(nums []int) []int {
    n, h := len(nums), 1
    for h < n/3 { h = 3*h + 1}
    for h >= 1 {
        for i := h; i < n; i++ {
            for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
                nums[j], nums[j-h] = nums[j-h], nums[j]
            }
        }
        h /= 3
    }
}
```

##### 归并排序

算法思想

如果两个数组有序，可以通过归并的方式将其合并成一个有序的数组。

```go
// date 2020/01/05
// 自顶向下的归并排序，递归
// 时间复杂度O(nlgn)
func sortArray(nums []int) []int {
    if len(nums) < 2 { return nums }
    mid := len(nums) >> 1
    l := sortArray(nums[:mid])
    r := sortArray(nums[mid:])
    return merge(l, r)
}

// 归并排序
// 归并排序的merge算法
func merge(a1, a2 []int) []int {
    if len(a1) == 0 { return a2 }
    if len(a2) == 0 { return a1 }
    n1, n2 := len(a1), len(a2)
    n := n1 + n2
    res := make([]int, n)
    i, j := 0, 0
    for k := 0; k < n; k++ {
        if i >= n1 {
            res[k] = a2[j]
            j++
        } else if j >= n2 {
            res[k] = a1[i]
            i++
        } else if a1[i] < a2[j] {
            res[k] = a1[i]
            i++
        } else {
            res[k] = a2[j]
            j++
        }
    }
    return res
}
```

##### 堆排序

算法思想：使用大顶堆

```golang
func sortArray(nums []int) []int {
    makeMaxHeap(nums)
    n := len(nums)
    for n > 0 {
        swap(nums, 0, n-1)
        n--
        maxHeapify(nums, 0, n)
    }
    return nums
}

// 堆排序
// 构建大顶堆
func makeMaxHeap(nums []int) {
    for i := len(nums) >> 1 - 1; i >= 0; i-- {
        maxHeapify(nums, i, len(nums))
    }
}
// 自上而下的堆有序化
func maxHeapify(nums []int, start, end int) {
    if start > end { return }
    temp, l, r := nums[start], start << 1 + 1, start << 1 + 2
    for l < end {
        // 找到左右结点的最大值，并记录在l中
        r = l + 1
        if r < end && nums[r] > nums[l] { l++ }
        // 如果父节点大于左右中的最大节点，则跳出
        if nums[start] > nums[l] { break }
        // 否则将最大值更新到父节点
        nums[start] = nums[l]
        nums[l] = temp
        // 接着往下调整
        start = l
        l = start << 1 + 1
    }
} 
// swap
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
```

#### 922 按奇偶排序数组II【E】

思路分析：记录每个不"在位"的索引，然后统一调整。注意不能两两依次调整，因为前后两个”不在位“的原因可以一样。

```go
// date 2020/02/28
func sortArrayByParityII(A []int) []int {
  p1, p2 := make([]int, 0), make([]int, 0)
  for i := 0; i < len(A); i++ {
    if i & 0x1 == 0 && A[i] & 0x1 != 0 { p1 = append(p1, i) }
    if i & 0x1 == 1 && A[i] & 0x1 != 1 { p2 = append(p2, i) }
  }
  i, n := 0, len(p1)
  var t int
  for i < n {
    t = A[p1[i]]
    A[p1[i]] = A[p2[i]]
    A[p2[i]] = t
    i++
  }
  return A
}
```



#### 1086 前五科的均分

思路分析：先对id排序，再对成绩排序。

```go
// date 2020/02/28
func highFive(items [][]int) [][]int {
  sort.Slice(items, func(i, j int) bool {
    if items[i][0] == items[j][0] {
      return items[i][1] > items[j][1] // 对成绩降序排序
    }
    return items[i][0] < items[j][0] // 对id升序排序
  })
  res, id, temp, c := make([][]int, 0), items[0][0], 0, 0
  for _, it := range items {
    if it[0] == id && c < 5 {
      temp += it[1]
      k++
      if k == 5 {
        r := make([]int, 2)
        r[0], r[1] = id, temp/5
        res = append(res, r)
      }
    } else if it[0] != id {
      id = it[0]
      temp = it[1]
      k = 1
    }
  }
  return res
}
```



#### 1356 根据数字二进制下1的数目排序

算法分析：求一个数字的二进制下2的个数。

```go
// date 2020/02/28
// 暴力解法
func sortByBits(arr []int) []int {
  sort.Slice(arr, func(i, j int) bool {
    if bits(arr[i]) == bits(arr[j]) {
      return arr[i] < arr[j]
    }
    return bits(arr[i]) < bits(arr[j])
  })
  return arr
}
func bits(x int) int {
  res := 0
  for x > 0 {
    res += x & 0x1
    x >>= 1
  }
  return res
}
```



