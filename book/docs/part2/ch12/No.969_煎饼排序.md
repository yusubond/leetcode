## 969 煎饼排序-中等

题目：

给你一个整数数组 `arr` ，请使用 **煎饼翻转** 完成对数组的排序。

一次煎饼翻转的执行过程如下：

- 选择一个整数 `k` ，`1 <= k <= arr.length`
- 反转子数组 `arr[0...k-1]`（**下标从 0 开始**）

例如，`arr = [3,2,1,4]` ，选择 `k = 3` 进行一次煎饼翻转，反转子数组 `[3,2,1]` ，得到 `arr = [**1**,**2**,**3**,4]` 。

以数组形式返回能使 `arr` 有序的煎饼翻转操作所对应的 `k` 值序列。任何将数组排序且翻转次数在 `10 * arr.length` 范围内的有效答案都将被判断为正确。



分析：

其核心就是寻找前 N 个中的最大值，通过两次反转，把最大值放到最后面。

```go
// date 2023/12/08
func pancakeSort(arr []int) []int {
    n := len(arr)
    if n == 1 {
        return arr
    }

    ans := make([]int, 0, 16)
    
    var sort func(nums []int, total int)
    
    sort = func(nums []int, total int) {
        // find the maxV idx
        if total < 2 {
            return
        }
        maxIdx := 0
        maxV := nums[0]
        for i := 0; i < total; i++ {
            if nums[i] > maxV {
                maxIdx = i
                maxV = nums[i]
            }
        }
      	// 如果 max 是最后一个，说明不需要反转，直接往前递归
        if maxIdx == total-1 {
            sort(nums, total-1)
            return
        }
      	// 两次反转，然后递归
        ans = append(ans, maxIdx+1)
        ans = append(ans, total)
        reverse(nums, 0, maxIdx)
        reverse(nums, 0, total-1)
        sort(nums, total-1)
    }

    sort(arr, n)

    return ans
}

func reverse(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```

