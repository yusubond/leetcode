## 987 二叉树的垂序遍历-困难

题目：

给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。

对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。

二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。返回二叉树的 垂序遍历 序列。



算法分析：

从题目垂直遍历的要求中获取到三个信息，也称为三个目标：

1. 要按照从左到右的顺利存放最终结果
2. 每一列按照从上到下的顺序遍历，其实隐含的就是层序遍历的思想，要一层一层处理
3. 同一层的时候，如果同一列上存在多个值，要升序排列一下



因此大框架上我们还是选择层序遍历。

1. 用一个队列 `queue` 保存每一层的节点，同时在辅助一个队列 `colIdx` 保存 `queue`中每个节点的列信息。其次初始化两个变量 `lv` 和 `rv` 分别表示整棵树最左和最右边界，用于最后整理总结果集。

2. 每次出队的时候，更新 `lv` 和 `rv`，完成1目标。

3. 在层序遍历中，当遍历完当前层，取得当前层结果，因为当前层的结果是个中间状态，需要进行升序排序后，才能放入总结果集中。完成3目标
4. 当所有层都遍历后，再根据左右边界，重新整理一次结果。



```go
// date 2022/10/22
func verticalTraversal(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    colRes := make(map[int][]int, 16) // 存储每一列的所有结果，key为列值
    lv, rv := int(0), int(0)          // 整棵树的左右边界
    queue := make([]*TreeNode, 0, 16)
    colIdx := make([]int, 0, 16)
    queue = append(queue, root)
    colIdx = append(colIdx, 0)

    for len(queue) != 0 {
        n := len(queue)
        cr := make(map[int][]int, 16)  // 记录当前层结果
        // 遍历当前层
        for i := 0; i < n; i++ {
            // 出队
            cur := queue[i]
            idx := colIdx[i]
            // 更新左右边界
            if idx > rv {
                rv = idx
            }
            if idx < lv {
                lv = idx
            }
						// 更新当前层结果
            ov, ok := cr[idx]
            if !ok {
                ov = make([]int, 0, 16)
            }
            ov = append(ov, cur.Val)
            cr[idx] = ov
            // 继续检查左右结点
            if cur.Left != nil {
                queue = append(queue, cur.Left)
                colIdx = append(colIdx, idx-1)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
                colIdx = append(colIdx, idx+1)
            }
        }
        // 排序当前层结果，并加入总结果集
        for k, v := range cr {
            if len(v) > 1 {
                sort.Slice(v, func(i, j int) bool {
                    return v[i] < v[j]
                })
            }
            if len(v) != 0 {
                if old, ok := colRes[k]; ok {
                    old = append(old, v...)
                    colRes[k] = old
                } else {
                    colRes[k] = v
                }
            }
        }
        // 检查下一层
        queue = queue[n:]
        colIdx = colIdx[n:]
    }

    // 按左右边界整理总结果
    ans := make([][]int, 0, 16)
    for i := lv; i <= rv; i++ {
        if v, ok := colRes[i]; ok && len(v) != 0 {
            ans = append(ans, v)
        }
    }
    return ans
}
```