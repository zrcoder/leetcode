---
title: "1658. 将 x 减到 0 的最小操作数"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [前缀和, 哈希表, 滑动窗口]
---

## [1658. 将 x 减到 0 的最小操作数](https://leetcode-cn.com/problems/minimum-operations-to-reduce-x-to-zero/)

   难度中等

   给你一个整数数组 `nums` 和一个整数 `x` 。每一次操作时，你应当移除数组 `nums` 最左边或最右边的元素，然后从 `x` 中减去该元素的值。请注意，需要 **修改** 数组以供接下来的操作使用。

   如果可以将 `x` **恰好** 减到 `0` ，返回 **最小操作数** ；否则，返回 `-1` 。

​    

   **示例 1：**

   ```
   输入：nums = [1,1,4,2,3], x = 5
   输出：2
   解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
   ```

   **示例 2：**

   ```
   输入：nums = [5,6,7,8,9], x = 4
   输出：-1
   ```

   **示例 3：**

   ```
   输入：nums = [3,2,20,1,1,3], x = 10
   输出：5
   解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。
   ```

​    

   **提示：**

   - `1 <= nums.length <= 105`
   - `1 <= nums[i] <= 104`
   - `1 <= x <= 109`

函数签名：

```go
func minOperations(nums []int, x int) int
```

## 分析

问题稍微绕了一点，等价于找到和为定值 `sum-x` 的最长连续子数组 ，如果找到了，结果就是总长度减去这个子数组的长度。

程序第一部分逻辑如下：

```go
func minOperations(nums []int, x int) int {
	target := -x
	for _, v := range nums {
		target += v
	}
	if target == 0 {
		return len(nums)
	}
	if target < 0 {
		return -1
	}
	// calMaxSubSizeWithSum 找到和为 target 的最长子数组的长度，不存在返回 -1
	maxSize := calMaxSubSizeWithSum(target, nums)
	if maxSize == -1 {
		return -1
	}
	return len(nums) - maxSize
}
```

找到和为定值的最长连续子数组，这是一个经典的问题。

现在来看 `calMaxSubSizeWithSum` 函数，从朴素到最优的实现如下：

## 朴素解法

```go
func calMaxSubSizeWithSum(target int, nums []int) int {
	res := -1
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if target == sum && j-i+1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
```

时间复杂度 `O(n^2)`，空间复杂度 `O(1)`，超时了。

### 前缀和 + 哈希表

借助前缀和技巧，哈希表辅助，空间换时间：

```go
func calMaxSubSizeWithSum(target int, nums []int) int {
	res := -1
	// 键为前缀和，值为前缀末尾的索引
	dic := map[int]int{0: -1}
	sum := 0
	for i, v := range nums {
		sum += v
		if j, ok := dic[sum-target]; ok && i-j > res { // 子数组 [j+1:i] 的和为 target，长度大于 res
			res = i - j
		}
		// 这里不用判断 sum 是否已经存在于dic，
		// 因为数组里都是正整数，sum是不断递增的，dic 里肯定不存在
        // 如果数组元素有负数，则需要判断，为了保持子数组最长，存入的前缀越短越好，判断如果已经有当前 sum 则不更新前缀末尾索引
		dic[sum] = i
	}
	return res
}
```

时间复杂度 `O(n)`，空间复杂度 `O(n)`。

### 滑动窗口

基于双指针的滑动窗口解法。

用 left 和 right 两个指针代表窗口的左右边界，一开始两个指针都指向数组起始位置。

另用一个变量 sum 维护窗口里的元素和， 如果 sum < target，右边界右移，sum == target， 则更新结果且右边界右移；sum > target 则左指针右移。

> 如果数组中有负数，这个方法不能用，移动左右指针的贪心策略就错了。

总体来说 left 和 right 最坏情况下都遍历了一遍数组，复杂度是 `O(n)` 的。

```go
func calMaxSubSizeWithSum(target int, nums []int) int {
	res := -1
	left, right, sum := 0, 0, 0
	for right < len(nums) {
		sum += nums[right]
		right++
		for sum > target && left < right {
			sum -= nums[left]
			left++
		}
		if target == sum && right-left > res {
			res = right-left
		}
	}
	return res
}
```

时间复杂度 `O(n)`，空间复杂度 `O(1)`。