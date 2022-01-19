---
title: "33. 搜索旋转排序数组"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

难度中等

整数数组 `nums` 按升序排列，数组中的值 **互不相同** 。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转**，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,5,6,7]` 在下标 `3` 处经旋转后可能变为 `[4,5,6,7,0,1,2]` 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，如果 `nums` 中存在这个目标值 `target` ，则返回它的下标，否则返回 `-1` 。

**示例 1：**

```
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
```

**示例 2：**

```
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
```

**示例 3：**

```
输入：nums = [1], target = 0
输出：-1
```

**提示：**

- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 中的每个值都 **独一无二**
- 题目数据保证 `nums` 在预先未知的某个下标上进行了旋转
- `-10^4 <= target <= 10^4`

**进阶：**你可以设计一个时间复杂度为 `O(log n)` 的解决方案吗？

## 分析

虽然被旋转，但还是可以用二分法来查找。

旋转数组被旋转点分割成了左右两个有序的部分。

二分查找过程中，mid 的左右两侧定有一侧是有序的，这给二分创造了条件。
哪一侧有序，判断 target 是否在有序侧，是则抛弃无序的一侧，否则抛弃有序的一侧。

```go
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cur := nums[mid]
		if cur == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // nums[lo:mid] 有序
			if nums[lo] <= target && target < cur {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // nums[mid+1:hi+1] 有序
			if cur < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
```

当然也可以根据 mid 和 hi 处的元素大小比较来确定哪一侧有序：
```go
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[hi] {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}
```

## [81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

和上边问题唯一的区别是元素可能重复。

## 分析

类似上边问题的解法，但是需要考虑 `nums[lo]`、`nums[mid]` 和 `nums[hi]` 相等的情况，这种情况无法判断 lo 和 mid 是否在旋转点同侧，需要把左右指针都保守移动一步；这样最坏时间复杂的是 `O(n)`。

```go
func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cur := nums[mid]
		if cur == target {
			return true
		}
		if nums[lo] == cur && cur == nums[hi] {
			lo++
			hi--
		} else if nums[lo] <= nums[mid] { // nums[lo:mid+1] 有序，lo 和 mid 在旋转点同侧
			if nums[lo] <= target && target < cur {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // lo 在左半部分， mid 在右半部分
			if cur < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}
```

## [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

难度中等

假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 `[0,1,2,4,5,6,7]` 可能变为 `[4,5,6,7,0,1,2]` 。

请找出其中最小的元素。

**示例 1：**

```
输入：nums = [3,4,5,1,2]
输出：1
```

**示例 2：**

```
输入：nums = [4,5,6,7,0,1,2]
输出：0
```

**示例 3：**

```
输入：nums = [1]
输出：1
```

**提示：**

- `1 <= nums.length <= 5000`
- `-5000 <= nums[i] <= 5000`
- `nums` 中的所有整数都是 **唯一** 的
- `nums` 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转

##  分析

使用二分搜索模板二的变体，详见代码及注释。

> 因为要获取 nums[hi] 的值，`hi` 一开始取值 `len(nums)-1`。

```go
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] { // mid 和 hi 落在旋转点左右两侧
			lo = mid + 1
		} else { // mid 和 hi 在旋转点同侧，但因为一开始 hi 在整个数组最右，所以当前只可能同在旋转点右侧
			hi = mid
		}
	}
	return nums[lo]
}
```

## [154. 寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii)
类似上个问题，但是允许元素重复，如：

```
[3, 1, 2, 3]
[3, 1, 2, 2, 3]
```

## 分析

解法也类似上边的问题解法，只是需要考虑 `nums[mid] < nums[hi]` 的情况，这种情况下右指针保守缩进，只向左移动一步。

```go
func findMin1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] > nums[hi]: // mid落在旋转点左侧
			lo = mid + 1
		case nums[mid] < nums[hi]: // mid和hi在旋转点同侧，但因为一开始hi在整个数组最右，所以当前只可能同在旋转点右侧
			hi = mid
		default: // 相等时保守缩进，避免遗漏一些元素
			hi--
		}
	}
	return nums[lo]
}
```