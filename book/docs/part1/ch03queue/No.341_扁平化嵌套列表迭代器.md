## 341 扁平化嵌套列表迭代器-中等

题目：

给你一个嵌套的整数列表 `nestedList` 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。

实现扁平迭代器类 `NestedIterator` ：

- `NestedIterator(List<NestedInteger> nestedList)` 用嵌套列表 `nestedList` 初始化迭代器。
- `int next()` 返回嵌套列表的下一个整数。
- `boolean hasNext()` 如果仍然存在待迭代的整数，返回 `true` ；否则，返回 `false` 。

你的代码将会用下述伪代码检测：

```
initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res
```

如果 `res` 与预期的扁平化列表匹配，那么你的代码将会被判为正确。



分析：

深度优先搜索。

嵌套的整数列表实际上是一个树形结构，树上的叶子节点对应一个整数，非叶子节点对应一个列表。

所以直接对嵌套结构进行深度优先搜索，得到的就是迭代器的遍历顺序。

把遍历到的元素存入数组即可。



```go
// date 2023/11/30
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    nums := make([]int, 0, 32)
    var dfs func([]*NestedInteger)
    dfs = func(nestedList []*NestedInteger) {
        for _, nest := range nestedList {
            if nest.IsInteger() {
                nums = append(nums, nest.GetInteger())
            } else {
                dfs(nest.GetList())
            }
        }
    }

    dfs(nestedList)
    return &NestedIterator{vals: nums}
}

func (this *NestedIterator) Next() int {
    res := this.vals[0]
    this.vals = this.vals[1:]
    return res   
}

func (this *NestedIterator) HasNext() bool {
    return len(this.vals) != 0
}
```

