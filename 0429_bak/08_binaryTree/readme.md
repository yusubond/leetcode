## 二叉树

[TOC]

### 1.什么是二叉树

​		树是一种数据结构，树中的每个节点都包含一个键值和所有子节点的列表，对于二叉树来说，每个节点最多有两个子树结构，分别称为左子树和右子树。

#### 二叉树的深度

​	二叉树的深度是指二叉树的根结点到叶子结点的距离。

​	最大深度 = 根结点到叶子结点的最大距离

​	最小深度 = 根结点到叶子结点的最小距离。

​	常见的问题包括求取二叉树的最大/最大深度，一般可以使用递归算法或者DFS。

参见题目：

- [二叉树的最大深度](104.md)
- [二叉树的最小深度](111.md)



### 2.二叉树的遍历

按照 root 节点的访问顺序，可分为前序遍历，中序遍历和后序遍历。其顺序如下：

- 前序遍历：按「根-左-右」依次访问各节点
- 中序遍历：按「左-根-右」依次访问各节点
- 后序遍历：按「左-右-根」依次访问各节点

中序遍历的题目参见

- 098 验证二叉搜索树，迭代中序遍历
- 099 恢复二叉搜索树，迭代中序遍历



递归版本的遍历，只要理解其思想就很好写；而对于非递归版本的遍历，需要深入理解其结点的遍历顺序，并记录下来之前经过的结点，所以一定会用到栈。



层序遍历是一种特殊的遍历，从二叉树的深度来理解就是，将同一深度的结果按照从左到右的顺序形成一个数组，不同深度的结果再放到另一个数组里面，最终形成一个二维数组。



从广度搜索和深度搜索的角度来讲，层序遍历属于广度优先搜索，而前序，中序，后序遍历均为深度优先搜索。

参见题目

- No.094—二叉树的中序遍历
- No.102—二叉树的层序遍历
- No.144—二叉树的前序遍历
- No.145—二叉树的后序遍历



#### 2.1 前序遍历

前序遍历是指先访问root节点，然后左子树，最后是右子树。

```go
// date 2020/10/19
// 递归版
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 16)
    res = append(res, root.Val)
    if root.Left != nil {
        res = append(res, preorderTraversal(root.Left)...)
    }
    if root.Right != nil {
        res = append(res, preorderTraversal(root.Right)...)
    }
    return res
}

// 迭代版
// 前序遍历：root->left->right
func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil { return res }
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) != 0 {
        // find root and left
        for root != nil {
            res = append(res, root.Val)  // 先将根节点加入结果集
            stack = append(stack, root)  // 将遍历过的结点加入栈，后序出栈依次访问结点的右子树
            root = root.Left             // 遍历当前结点的左子树
        }
        // find the right
        if len(queue) != 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }
    return res
}
```

参见题目：

- 105—从前序和中序遍历序列构造二叉树



#### 2.2 中序遍历

中序遍历是指访问左子树，然后root节点，最后是右子树。

```go
// 递归版
func inOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  // 先访问左子树
  if root.Left != nil {
    res = append(res, inOrder(root.Left)...)
  }
  // 后访问根结点
  res = append(res, root.Val)
  // 最后访问右子树
  if root.Right != nil {
    res = append(res, inOrder(root.Right)...)
  }
  return res
}
```

这里跟前序遍历类似，只是取根节点值加入结果集的时机略有区别。

其主要思路为：

1. 持续检查其左子树，重复1，直到左子树为空【此时栈中保留的是一系列的左子树节点】
2. 出栈一次，并取根节点值加入结果集，然后检查最后一个左子树的右子树，重复1,2

```go
// 迭代版
// 需要stack结构
// left->root->right
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := make([]*TreeNode, 0, 16)
    res := make([]int, 0, 16)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        if len(stack) != 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = append(res, root.Val)
            root = root.Right
        }
    }
    return res
}
```

参见题目：

- 106—从中序和后序遍历序列构造二叉树
- 538—把二叉搜索树转换为累加树，递归中序遍历解法
- 783—二叉搜索树节点最小距离，递归中序遍历



#### 2.3 后序遍历

后序遍历是指先访问左子树，然后右子树，最后是root节点。

递归算法：

```go
// date 2020/10/19
// 递归
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 16)
    if root.Left != nil {
        res = append(res, postorderTraversal(root.Left)...)
    }
    if root.Right != nil {
        res = append(res, postorderTraversal(root.Right)...)
    }
    res = append(res, root.Val)
    return res
}

```

迭代算法：

迭代遍历的时候依然需要 stack 结构来保存已经遍历过的节点；同时借助 pre 指针保存上次出栈的节点，用于判断当前节点是否同时具有左右子树，还是只有单个子树。

1. 先将根节点入栈，循环遍历栈是否为空
2. 出栈，取出当前节点
   1. 如果当前节点没有左右子树，则为叶子节点，直接加入结果集
   2. 其次判断上次出栈的节点是否是当前节点的左右子树，如果是，表明当前节点的子树已经处理完毕，也需要加入结果集【这里依赖的是左右子树的入栈，因为在栈中左右子树不具备前继关系，至于根节点具备】
   3. 依次检查当前节点的右子树，左子树，重复1,2

```go
// 迭代版
// left->right-root
func postorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    var pre, cur *TreeNode   // 记录前驱节点和当前节点
    for len(stack) != 0 {
        // 出栈 当前结点
        cur = stack[len(stack)-1]
        // 如果当前结点为叶子结点，则直接加入结果集
        // 如果当前结点不是叶子结点，但是上次遍历结点为当前结点的左右子树时(说明当前结点只有单个子树，且子树已经处理完毕)，也加入结果集
        if cur.Left == nil && cur.Right ==  nil || pre != nil && (pre == cur.Left || pre == cur.Right) {
            res = append(res, cur.Val)
            // 出栈，继续检查
            stack = stack[:len(stack)-1]
            pre = cur
        } else {
            // 因为在出栈的时候检查结点，并追加到结果中
            // 所以，先入栈右子树，后入栈左子树
            if cur.Right != nil {
                stack = append(stack, cur.Right)
            }
            if cur.Left != nil {
                stack = append(stack, cur.Left)
            }
        }
    }
    
    return res
}
```

![image](images/image145.png)



参见题目：

- 106—从中序和后序遍历序列构造二叉树



#### 2.4 层序遍历

层序遍历是指逐层遍历树的结构，也称为广度优先搜索，算法从一个根节点开始，先访问根节点，然后遍历其相邻的节点，其次遍历它的二级、三级节点。

借助队列数据结构，先入先出的顺序，实现层序遍历。

```go
// date 2020/03/21
// 层序遍历
// bfs广度优先搜索
// 算法一：使用队列，逐层遍历
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := make([][]int, 0, 16)
    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)
    for len(queue) != 0 {
        n := len(queue)
        curRes := make([]int, 0, 16)
        for i := 0; i < n; i++ {
            cur := queue[i]
            curRes = append(curRes, cur.Val)
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        res = append(res, curRes)
        queue = queue[n:]
    }
    return res
}
```

算法2：dfs深度优先搜索

这里的思路类似求二叉树的最大深度，借助dfs搜索，在每一层追加结果。

```go
// date 2020/03/21
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0, 16)
    var dfs func(root *TreeNode, level int)
    dfs = func(root *TreeNode, level int) {
        if root == nil {
            return
        }
        if len(res) == level {
            res = append(res, make([]int, 0, 4))
        }
        res[level] = append(res[level], root.Val)
        level++
        dfs(root.Left, level)
        dfs(root.Right, level)
    }
    dfs(root, 0)
    return res
}
```

**注意**，从上面两种实现的方式来看，层序遍历既可以使用广度优先搜索，也可以使用深度优先搜索。

参见题目

- 102—二叉树的层序遍历
- 103—二叉树的锯齿形层序遍历
- 107—二叉树的层序遍历2，将底部结果放到开头，没有新意
- 199—二叉树的右视图，取层序遍历最后一个节点
- 404—左叶子之和
- 987—二叉树的垂序遍历，层序遍历，每个节点标注坐标，根据坐标排序





#### 2.5 垂序遍历

> 给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。
>
> 对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。
>
> 二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。
>
> 返回二叉树的 垂序遍历 序列。
>
> 来源：力扣（LeetCode）
> 链接：https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree
> 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

详见题目[987二叉树的垂序遍历](987.md)



#### 2.6 相关题目

- [094二叉树的中序遍历](094.md)

- [102二叉树的层序遍历](102.md)

- [102二叉树的锯齿形层序遍历](103.md)

- [105从前序与中序遍历序列构造二叉树](105.md)
- [106从后续与中序遍历序列构造二叉树](106.md)
- [107二叉树的层序遍历2](107.md)
- [144二叉树的前序遍历](144.md)
- [145二叉树的后序遍历](145.md)
- [331验证二叉树的前序序列化](331.md)
- [429N叉树的层序遍历](429.md)
- [637二叉树的层平均值](637.md)
- [1161最大层内元素和](1161.md)





### 3.递归解决树的问题

递归通常是解决树的相关问题最有效和最常用的方法之一，分为**自顶向下**和**自底向上**两种。

**自顶向下**的解决方案

自顶向下意味着在每个递归层级，需要先计算一些值，然后递归调用并将这些值传递给子节点，视为**前序遍历**。参见题目【二叉树的最大深度】

```go
// 通用的框架，自顶向下
1. return specific value for null node
2. update the answer if needed                      // answer <-- params
3. left_ans = top_down(root.left, left_params)      // left_params <-- root.val, params
4. right_ans = top_down(root.right, right_params)   // right_params <-- root.val, params 
5. return the answer if needed                      // answer <-- left_ans, right_ans
```

参见题目

- No.95—不同的二叉搜索数2，生成所有的二叉搜索数
- No.113—路径总和2，求满足条件的所有路径
- No.129—求根节点到叶子节点数字之和，携带已知内容递归到叶子节点
- No.257—二叉树的所有路径



**自底向上**的解决方案

自底向上意味着在每个递归层级，需要先对子节点递归调用函数，然后根据返回值和根节点本身得到答案，视为**后序遍历**。

```go
// 通用的框架，自底向上
1. return specific value for null node
2. left_ans = bottom_up(root.left)          // call function recursively for left child
3. right_ans = bottom_up(root.right)        // call function recursively for right child
4. return answers                           // answer <-- left_ans, right_ans, root.val
```

参见题目：

- 543—二叉树的直径，求左右深度之和。
- 865—具有所有最深节点的最小子树





**总结**：

当遇到树的问题时，先思考一下两个问题：

> 1.你能确定一些参数，从该节点自身解决出发寻找答案吗？
>
> 2.你可以使用这些参数和节点本身的值来决定什么应该是传递给它子节点的参数吗？

如果答案是肯定的，那么可以尝试使用自顶向下的递归来解决问题。

当然也可以这样思考：

> 对于树中的任意一个节点，如果你知道它子节点的结果，你能计算出该节点的答案吗？

如果答案是肯定的，可以尝试使用自底向上的递归来解决问题。



### 深度优先搜索

深度优先搜索的过程，过程中计算最终的结果集，和深搜本身的返回值不一致。这并不矛盾。反而，有些题目更好使。参见题目979-在二叉树中分配硬币。

参见题目：





### 广度优先搜索





### 4 平衡二叉树

一棵高度平衡二叉树的定义为:

> 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

参见题目：

- 108—将有序数组转换为二叉搜索树，递归实现



### 5 对称二叉树

如果一棵二叉树关于根结点左右对称，则称为对称二叉树，也称为镜像二叉树。

参见题目：

- 101—对称二叉树，递归，或者迭代两个树



---

### 6 Binary Indexed Tree \[二叉索引树\]

  * 1.[结构介绍](#结构介绍)
  * 2.[问题定义](#问题定义)
  * 3.[表面形式](#表现形式)
    * 3.1[更新操作](#更新操作)
    * 3.2[求和操作](#求和操作)
  * 4.[工作原理](#BIT工作原理)
  * 5.[代码实现](#代码实现)

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





---

### 二叉搜索树

顾名思义，二叉搜索树是以一棵二叉树来组织的，其满足一下条件：x为二叉树的一个节点，如果y是x左子树的一个节点，那么`y.Val <= x.Val`；如果y是x右子树的一个节点，则`y.Val >= x.Val`

根据二叉树的性质可得知，二叉搜索树的中序遍历为升序序列，相反如果交换左右子树的遍历顺序亦可得到降序序列。

```go
// inorder伪代码
func inOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, inOrder(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inOrder(root.Right)...)
  }
  return res
}

func decOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, decOrder(root.Right)...)
    res = append(res, root.Val)
    res = append(res, decOrder(root.Left)...)
  }
  return res
}
```

#### 二叉搜索树的基本操作

二叉搜索树的基本操作包括查询，最小关键字元素，最大关键字元素，前继节点，后继节点，插入和删除，这些基本操作所花费的时间与二叉搜索树的高度成正比。

**查询**

如果二叉搜索树的高度为h，则查询的时间复杂度为O(h)。

```go
// 查询伪代码
// 递归版
func searchBST(root *TreeNode, key int) *TreeNode {
  if root == nil || root.Val == key { return root }
  if root.Val < key {
    return searchBST(root.Right, key)
  }
  return searchBST(root.Left, key)
}
// 迭代版
func searchBST(root *TreeNode, key int) *TreeNode {
  for root != nil && root.Val != key {
    if root.Val < key {
      root = root.Right
    } else {
      root = root.Left
    }
  }
  return root
}
```

**最大关键字元素**

如果二叉搜素树的高度为h，则最大关键字元素的操作时间复杂度为O(h)。

```go
// 最大关键字元素的伪代码
func maximumBST(root *TreeNode) *TreeNode {
  for root.Right != nil { root = root.Right }
  return root
}
// 递归版
func maximumBST(root *TeeNode) *TreeNode {
  if root.Right != nil {
    return maximumBST(root.Right)
  }
  return root
}
```

**最小关键字元素**

如果二叉搜素树的高度为h，则最大关键字元素的操作时间复杂度为O(h)。

```go
// 最小关键字元素的伪代码
func minimumBST(root *TreeNode) *TreeNode {
  for root.Left != nil { root = root.Left }
  return root
}
// 递归版
func minimumBST(root *TreeNode) *TreeNode {
  if root.Left != nil {
    return minimum(root.Left)
  }
  return root
}
```

**前继节点**

前继节点是指给定一个节点x，如果存在x的前继节点则返回前继节点，否则返回nil。

如果二叉搜素树的高度为h，则查找前继节点的操作时间复杂度为O(h)。

```go
func predecessorBST(x *TreeNode) *TreeNode {
  // x节点存在左子树，则返回左子树中最大的节点
  if x.Left != nil {
    return maximum(root.Left)
  }
  // 如果x节点没有左子树，则找到其父节点，且父节点的右子树包含x
  y := x.Present
  for y != nil && x == y.Left {
    x = y
    y = y.Present
  }
  return y
}
```

**后继节点**

后继节点是指给定一个节点x，如果存在x的后继节点则返回后继节点，否则返回nil。

如果二叉搜素树的高度为h，则查找后继节点的操作时间复杂度为O(h)。

```go
// 后继节点的伪代码
func successorBST(x *TreeNode) *TreeNode {
  // x节点存在右子树，则返回右子树中最小的节点
  if x.Right != nil {
    return minimum(root.Right)
  }
  // 如果x节点没有右子树，则找到其父节点，且父节点的左子树包含x
  y := x.Present
  for y != nil && x == y.Right {
    x = y
    y = y.Present
  }
  return y
}
```

**插入**

插入和删除会引起二叉搜索树所表示的动态集合的变化。一定要修改数据结构来反映这种变化，单修改要保持二叉搜索树的性质不变。

```go
// x.Val = val, x.Right = x.Left = x.Present = nil
// 插入的伪代码
func insertBST(root, x *TreeNode) {
  var y *TreeNode
  for root != nil {
    y = root
    if x.Val < root.Val {
      root = root.Left
    } else {
      root = root.Right
    }
  }
  x.Present = y
  if y == nil { // empty tree
    root = x
  } else if x.Val < y.Val {
    y.Left = x
  } else {
    y.Right = x
  }
}
// 没有父指针
// 递归版
func insertBST(root, x *TreeNode) *TreeNode {
  if root == nil {
    root = x
    return root
  }
  if root.Val < x.Val {
    return insertBST(root.Right)
  }
  return insertBST(root.Left)
}
// 迭代版
func insertBST(root, x *TreeNode) *TreeNode {
  if root == nil {
    root == x
    return root
  }
  pre, head := root, root
  for head != nil {
    pre = head
    if root.Val < x.Val {
      head = head.Right
    } else {
      head = head.Left
    }
  }
  if y.Val < x.Val {
    y.Right = x
  } else {
    y.Left = x
  }
  return root
}
```

**删除**

从一棵二叉搜索树中删除一个节点x的整个策略分为三个情况，分别是：

1）如果x没有孩子节点，那么直接删除即可；

2）如果x只有一个孩子节点，那么将该孩子节点提升到x即可；

3）如果x有两个孩子，那么需要找到x的后继节点y（一定存在x的右子树中），让y占据x的位置，那么x原来的右子树部分称为y的右子树，x的左子树称为y的新的左子树。

```go
func deleteBST(root *TreeNode, key int) *TreeNode{
  if root == nil { return root }
  if root.Val < key {
    return deleteBST(root.Right, key)
  } else if root.Val > key {
    return deleteBST(root.Left, key)
  } else {
    // no child node
    if root.Left == nil && root.Right == nil {
      root = nil
    } else if root.Right != nil {
      root.Val = postNodeVal(root.Right)
      root.Right = deleteBST(root.Right, root.Val)
    } else {
      root.Val = preNodeVal(root.Left)
      root.Left = deleteBST(root.Left, root.Val)
    }
  }
  return root
}

// find post node val in right
func postNodeVal(root *TreeNode) int {
  for root.Left != nil { root = root.Left }
  return root.Val
}
// find pre node val in left
func preNodeVal(root *TreeNode) int {
  for root.Right != nil { root = root.Right }
  return root.Val
}
```

### 二叉搜索树相关题目

- 108 将有序数组转换为二叉搜索树【M】
- 109—将有序链表转换为二叉搜索树
- 235 二叉搜素树的最近公共祖先
- 230 二叉搜索树中第K小的元素
- 450 删除二叉搜索树中的节点
- 1038 从二叉搜索树到更大和树【M】
- 1214 查找两棵二叉搜索树之和【M】
- 面试题 17.12 BiNode【E】
- 面试题54 二叉搜索树的第K大节点

#### 108 将有序数组转换为二叉搜索树

题目要求：给定一个升序数据，返回一棵高度平衡二叉树。

思路分析：

二叉搜索树的定义：

1）若任意节点的左子树不为空，则左子树上所有的节点值均小于它的根节点值；

2）若任意节点的右子树不为空，则右子树上所有的节点值均大于它的根节点值；

3）任意节点的左，右子树均为二叉搜索树；

4）没有键值相等的点。

如何构造一棵树，可拆分成无数个这样的子问题：构造树的每个节点，以及节点之间的关系。对于每个节点来说，都需要：

1）选取节点；

2）构造该节点的左子树；

3）构造该节点的右子树。

算法一：对于升序序列，可以选择中间的节点作为根节点，然后递归调用。因为每个节点只遍历一遍，所以时间复杂度为O(n)；因为递归调用，空间复杂度为O(log(n))。

```go
//date 2020/02/20
// 算法一：递归版
func sortedArrayToBST(nums []int) *TreeNode {
  if len(nums) == 0 { return nil }
  // 选择根节点
  mid := len(nums) >> 1
  // 构造左子树和右子树
  left := nums[:mid]
  right := nums[mid+1:]
  
  return &TreeNode{
    Val: nums[mid],
    Left: sortedArrayToBST(left),
    Right: sortedArrayToBST(right),
  }
}
```

#### 230 二叉搜索树中第K小的元素

题目要求：给定一个二叉搜索树，返回其第K小的元素，你可以假设K总是有效的，即1<=k<=n。

算法：因为是二叉搜索树，先得到其中序遍历结果，然后直接返回k-1所在的元素。

```go
// date 2020/01/11
func kthSmallest(root *TreeNode) int {
  data := inOrder(root)
  if len(data) < k {return 0}
  return data[k-1]
}

func inOrder(root *TreeNode) []int {
  if root == nil {return nil}
  res := make([]int, 0)
  if root.Left != nil { res = append(res, inOrder(root.Left)...) }
  res = append(res, root.Val)
  if root.Right != nil { res = append(res, inOrder(root.Right)...) }
}
```


#### 235 二叉搜索树的最近公共祖先

题目要求：给定一个二叉搜素树和两个节点，返回其最近公共祖先。

算法：递归，并充分利用二叉搜索树的性质`left.Val < root.Val < right.Val`。

```go
// date 2020/02/18
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil || root == p || root == q {return root }
  if p == nil { return q }
  if q == nil { return p }
  if root.Val > q.Val && root.Val > p.Val { return lowestCommonAncestor(root.Left) }
  if root.Val < q.Val && root.Val < p.Val { return lowestCommonAncestor(root.Right) }
  return root
}
```



#### 450 删除二叉搜索树中的节点

题目要求：给定一个二叉搜索树和key，删除key节点，并返回树的根节点

算法一：二叉搜索树的中序遍历序列是个升序序列，则欲要删除的节点需要用其前驱节点或后驱节点来代替。

1）如果为叶子节点，则直接删除；

2）如果不是叶子节点且有右子树，则用后驱节点代替。后驱节点为右子树中的较低位置；

3）如果不是叶子节点且有左子树，则用前驱节点代替。前驱节点为左子树中的较高位置；  

```go
func deleteBST(root *TreeNode, key int) *TreeNode{
  if root == nil { return root }
  if root.Val < key {
    return deleteBST(root.Right, key)
  } else if root.Val > key {
    return deleteBST(root.Left, key)
  } else {
    // no child node
    if root.Left == nil && root.Right == nil {
      root = nil
    } else if root.Right != nil {
      root.Val = postNodeVal(root.Right)
      root.Right = deleteBST(root.Right, root.Val)
    } else {
      root.Val = preNodeVal(root.Left)
      root.Left = deleteBST(root.Left, root.Val)
    }
  }
  return root
}

// find post node val in right
func postNodeVal(root *TreeNode) int {
  for root.Left != nil { root = root.Left }
  return root.Val
}
// find pre node val in left
func preNodeVal(root *TreeNode) int {
  for root.Right != nil { root = root.Right }
  return root.Val
}
```

#### 1038 从二叉搜索树到更大和树

题目要求：给出二叉**搜索**树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 `node` 的新值等于原树中大于或等于 `node.val` 的值之和

思路分析：

`root.Val = root.Val + root.Right.Val`

`root.Left.Val = root.Left.Val + root.Val`

首先得到二叉搜索树的中序遍历序列的逆序列，然后从序列的第一个开始，将累加值记录到当前树的节点上。

```go
// date 2020/02/23
func bsttoGST(root *TreeNode) *TreeNode{
  array := decOrder(root)
  val := 0
  for i := 0; i < len(array); i++ {
    val += array[i]
    setNewVal(root, array[i], val)
  }
}
func decOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, decOrder(root.Right)...)
    res = append(res, root.Val)
    res = append(res, decOrder(root.Left)...)
  }
  return res
}

func setNewVal(root *TreeNode, key, val int) {
  if root != nil {
    if root.Val == key {
      root.Val = val
    } else if root.Val > key {
      setNewVal(root.Left, key, val)
    } else {
      setNewVal(root.Right, key, val)
    }
  }
}
```

#### 面试题 BiNode

题目要求：将一颗二叉搜索树转换成单链表，Right指针指向下一个节点。

思路分析：

要求返回头节点，所以先转换left。

```go
// date 2020/02/19
func convertBiNode(root *TreeNode) *TreeNode {
  if root == nil { return nil }
  // 左子树不为空，需要转换
  if root.Left != nil {
    // 得到左子树序列，并找到最后一个节点
    left :=  convertBiNode(root.Left)
    pre := left
    for pre.Right != nil { pre = pre.Right }
    // 将最后一个节点right指向root
    pre.Right = root
    root.Left = nil
    // 得到右子树序列
    root.Right = convertBiNode(root.Right)
    return left
  }
  root.Right = convertBiNode(root.Right)
  return root
}
```

#### 面试题54 二叉搜索树的第K大节点

题目要求：给定一个二叉搜索树，找出其中第k大节点。

算法一：二叉搜索树的中序遍历为升序序列，按照Right->root->Left的顺序遍历可得到降序序列，然后返回第k节点。

```go
// date 2020/02/18
func kthLargest(root *TreeNode, k int) int {
  if root == nil || k < 0 { return -1 }
  nums := decOrder(root)
  if len(nums) < k {return -1}
  return nums[k-1]
}

// right->root->left
func decOrder(root *TreeNode) []int {
  if root == nil { return []int{} }
  res := make([]int, 0)
  if root.Right != nil {
    res = append(res, decOrder(root.Right)...)
  }
  res = append(res, root.Val)
  if root.Left != nil {
    res = append(res, decOrder(root.Left)...)
  }
  return res
}
```



### [相关题目](leetcode_bt.md)

#### 156 上下翻转二叉树【中等】

题目要求：给定一个二叉树，其中所有的右节点要么是具有兄弟节点（拥有相同父节点的左节点）的叶节点，要么为空，将此二叉树上下翻转并将它变成一棵树， 原来的右节点将转换成左叶节点。返回新的根。

算法分析：

```go
// date 2020/02/25
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var parent, parent_right *TreeNode
  for root != nil {
    // 0.保存当前left, right
    root_left := root.Left
    root.Left = parent_right  // 重新构建root,其left为上一论的right
    parent_right := root.Right // 保存right
    root.Right = parent // 重新构建root,其right为上一论的root
    parent = root
    root = root_left
  }
  return parent
}
```

#### 270 最接近的二叉搜索树值【简单】

题目要求：给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

算法分析：使用层序遍历，广度优先搜索。

```go
// date 2020/02/23
func closestValue(root *TreeNode, target float64) int {
    if root == nil { return 0 }
    var res int
    cur := math.Inf(1)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        n := len(stack)
        for i := 0; i < n; i++ {
            if math.Abs(float64(stack[i].Val) - target) < cur {
                cur = math.Abs(float64(stack[i].Val) - target)
                res = stack[i].Val
            }
            if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
            if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
        }
        stack = stack[n:]
    }
    return res
}
```

#### 543 二叉树的直径【简单】

题目要求：给定一棵二叉树，返回其直径。

算法分析：左右子树深度和的最大值。

```go
// date 2020/02/23
func diameterOfBinaryTree(root *TreeNode) int {
    v1, _ := findDepth(root)
    return v1
}

func findDepth(root *TreeNode) (int, int) {
  if root == nil { return 0, 0 }
  var v int
  v1, l := findDepth(root.Left)
  v2, r := findDepth(root.Right)
  if v1 > v2 {
    v = v1
  } else {
    v = v2
  }
  if l+r > v { v = l+r }
  if l > r { return v, l+1}
  return v, r+1
}
```

#### 662 二叉树的最大宽度【中等】

题目要求：给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
		链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree

算法分析：使用栈stack逐层遍历，计算每一层的宽度，注意全空节点的时候要返回。时间复杂度O(n)，空间复杂度O(logn)。

```go
// date 2020/02/26
func widthOfBinaryTree(root *TreeNode) int {
  res := 0
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  var n, i, j int
  var isAllNil bool
  for 0 != len(stack) {
    i, j, n = 0, len(stack)-1, len(stack)
    // 去掉每一层两头的空节点
    for i < j && stack[i] == nil { i++ }
    for i < j && stack[j] == nil { j-- }
    // 计算结果
    if j-i+1 > res { res = j-i+1 }
    // 添加下一层
    isAllNil = true
    for ; i <= j; i++ {
      if stack[i] != nil {
        stack = append(stack, stack[i].Left, stack[i].Right)
        isAllNil = false
      } else {
        stack = append(stack, nil, nil)
      }
    }
    if isAllNil { break }
    stack = stack[n:]
  }
  return res
}
```

#### 669 修剪二叉搜索树【简单】

题目要求：给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

算法分析：

```go
// date 2020/02/25
func trimBST(root *TreeNode, L, R int) *TreeNode {
  if root == nil { return nil }
  if root.Val < L { return trimBST(root.Right, L, R) }
  if root.Val > R { return trimBST(root.Left, L, R) }
  root.Left = trimBST(root.Left, L, R)
  root.Right = trimBST(root.Right, L, R)
  return root
}
```

#### 703 数据流中第K大元素【简单】

解题思路：

0。使用数据构建一个容量为k的小顶堆。

1。如果小顶堆堆满后，向其增加元素，若大于堆顶元素，则重新建堆，否则掉弃。

2。堆顶元素即为第k大元素。

```go
// date 2020/02/19
type KthLargest struct {
  minHeap []int
  size
}

func Constructor(k int, nums []int) KthLargest {
  res := KthLargest{
    minHeap: make([]int, k),
  }
  for _, v := range nums {
    res.Add(v)
  }
}

func (this *KthLargest) Add(val int) int {
  if this.size < len(this.minHeap) {
    this.size++
    this.minHeap[this.size-1] = val
    if this.size == len(this.minHeap) {
      this.makeMinHeap()
    }
  } else if this.minHeap[0] < val {
    this.minHeap[0] = val
    this.minHeapify(0)
  }
  return this.minHeap[0]
}

func (this *KthLargest) makeMinHeap() {
  // 为什么最后一个根节点是n >> 1 - 1
  // 长度为n，最后一个叶子节点的索引是n-1; left = i >> 1 + 1; right = i >> 1 + 2
  // 父节点是 n >> 1 - 1
  for i := len(this.minHeap) >> 1 - 1; i >= 0; i-- {
    this.minHeapify(i)
  }
}

func (this *KthLargest) minHeapify(i int) {
  if i > len(this.minHeap) {return}
  temp, n := this.minHeap[i], len(this.minHeap)
  left := i << 1 + 1
  right := i << 1 + 2
  for left < n {
    right = left + 1
    // 找到left, right中较小的那个
    if right < n && this.minHeap[left] > this.minHeap[right] { left++ }
    // left, right 均小于root否和小顶堆，跳出
    if this.minHeap[left] >= this.minHeap[i] { break }
    this.minHeap[i] = this.minHeap[left]
    this.minHeap[left] = temp
		// left发生变化，继续调整
    i = left
    left = i << 1 + 1
  }
}
```

#### 814 二叉树剪枝【中等】

题目要求：给定一个二叉树，每个节点的值要么是0，要么是1。返回移除了所有不包括1的子树的原二叉树。

算法分析：

使用containOne()函数判断节点以及节点的左右子树是否包含1。

```go
// date 2020/02/25
func pruneTree(root *TreeNode) *TreeNode {
  if containsOne(root) { return root }
  return nil
}

func containsOne(root *TreeNode) bool {
  if root == nil { return false }
  l, r := containsOne(root.Left), containsOne(root.Right)
  if !l { root.Left = nil }
  if !r { root.Right = nil }
  return root.Val == 1 || l || r
}
```

#### 1104 二叉树寻路【中等】

题目要求：https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/

算法分析：

```go
// date 2020/02/25
func pathInZigZagTree(label int) []int {
  res := make([]int, 0)
  for label != 1 {
    res = append(res, label)
    label >>= 1
    y := highBit(label)
    lable = ^(label-y)+y
  }
  res = append(res, 1)
  i, j := 0, len(res)-1
  for i < j {
    t := res[i]
    res[i] = res[j]
    res[j] = t
    i++
    j--
  }
  return res
}label = label ^(1 << (label.bit_length() - 1)) - 1

func bitLen(x int) int {
  res := 0
  for x > 0 {
    x >>= 1
    res++
  }
  return res
}
```

#### 面试题27 二叉树的镜像【简单】

题目链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

题目要求：给定一棵二叉树，返回其镜像二叉树。

```go
// date 2020/02/23
// 递归算法
func mirrorTree(root *TreeNode) *TreeNode {
  if root == nil { return root }
  left := root.Left
  root.Left = root.Right
  root.Right = left
  root.Left = mirrorTree(root.Left)
  root.Right = mirrorTree(root.Right)
  return root
}
```

