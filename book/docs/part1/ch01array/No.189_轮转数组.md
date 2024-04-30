## 189 轮转数组-中等

题目：

给定一个数组，将数组中的元素向右移动K个位置，K是非负整数。



> **示例 1:**
>
> ```
> 输入: nums = [1,2,3,4,5,6,7], k = 3
> 输出: [5,6,7,1,2,3,4]
> 解释:
> 向右轮转 1 步: [7,1,2,3,4,5,6]
> 向右轮转 2 步: [6,7,1,2,3,4,5]
> 向右轮转 3 步: [5,6,7,1,2,3,4]
> ```
>
> **示例 2:**
>
> ```
> 输入：nums = [-1,-100,3,99], k = 2
> 输出：[3,99,-1,-100]
> 解释: 
> 向右轮转 1 步: [99,-1,-100,3]
> 向右轮转 2 步: [3,99,-1,-100]
> ```
>
>  
>
> **提示：**
>
> - `1 <= nums.length <= 105`
> - `-231 <= nums[i] <= 231 - 1`
> - `0 <= k <= 105`
>
>  
>
> **进阶：**
>
> - 尽可能想出更多的解决方案，至少有 **三种** 不同的方法可以解决这个问题。
> - 你可以使用空间复杂度为 `O(1)` 的 **原地** 算法解决这个问题吗？



思路分析：【推荐该算法】

向右轮转 K 个位置，就是将数组的后 K 个元素放到数组的头部。那么，可以通过下面的方式得到：

1. 先整体反转数组
2. 再反转前 K-1 个元素，最后反转 后面的 n-k 个元素。

```go
// date 2022/09/22
func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}

func reverse(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```

算法图解：

![image-20220924164237944](./images/img189.svg)



其他算法：

暴力解法。时间复杂度O(n*k)，空间复杂度O(1)

先移动1步，再移动k步

```go
// date 2022/09/24
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
    for k > 0 {
        moveOneStep(nums)
        k--
    }
}

func moveOneStep(nums []int) {
    if len(nums) < 2 {
        return
    }
    n := len(nums)
    tail := nums[n-1]
    copy(nums[1:], nums[0:n-1])
    nums[0] = tail
}
```

算法图解：

![image-20220924164237944](../images/image_189.png)
