## 145 二叉树的后序遍历-简单

题目：

给你一棵二叉树的根节点 `root` ，返回其节点值的 **后序遍历** 。



分析：

算法1：递归

```go
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



算法2：迭代【推荐该算法】

迭代遍历的时候依然需要 stack 结构来保存已经遍历过的节点；同时借助 pre 指针保存上次出栈的节点，用于判断当前节点是否同时具有左右子树，还是只有单个子树。

1. 先将根节点入栈，循环遍历栈是否为空
2. 出栈，取出当前节点
   1. 如果当前节点没有左右子树，则为叶子节点，直接加入结果集
   2. 其次判断上次出栈的节点是否是当前节点的左右子树，如果是，表明当前节点的子树已经处理完毕，也需要加入结果集【这里依赖的是左右子树的入栈，因为在栈中左右子树不具备前继关系，至于根节点具备】
   3. 依次检查当前节点的右子树，左子树，重复1,2

```go
// 迭代版
// date 2023/10/18
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

