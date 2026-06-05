# 二叉搜索树

### 什么是二叉搜索树

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

### 二叉搜索树的基本操作

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

### 相关题目

- 108 将有序数组转换为二叉搜索树【M】
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

#### 