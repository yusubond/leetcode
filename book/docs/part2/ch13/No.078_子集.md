## 78 子集-中等

题目：

给你一个整数数组 `nums` ，数组中的元素 **互不相同** 。返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。你可以按 **任意顺序** 返回解集。



> **示例 1：**
>
> ```
> 输入：nums = [1,2,3]
> 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0]
> 输出：[[],[0]]
> ```



分析：

迭代算法版本详见数组章节。以下是回溯算法。

这道题要求返回所有的子集，所以空集合也算一个。

因为要找出所以的子集，所以回溯的时候，没有明显的结束条件，查到数组的最后结束即可。

其次，注意拷贝结果的时候，注意深度拷贝。

```go
// date 2023/12/26
func subsets(nums []int) [][]int {
    res := make([][]int, 0, 16)

    var backtrack func(start int, temp []int)
    backtrack = func(start int, temp []int) {
      	// 任何大小的子集都算，所以上来就追加结果
        one := make([]int, len(temp))
        copy(one, temp)
        res = append(res, one)
				
      	// 开始遍历
        for i := start; i < len(nums); i++ {
            temp = append(temp, nums[i])
            backtrack(i+1, temp)  // 向下一个元素回溯遍历
            temp = temp[:len(temp)-1]  // 取消候选
        }
    }

    backtrack(0, []int{})

    return res
}
```

