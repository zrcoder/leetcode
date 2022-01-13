---
title: "34. 在排序数组中查找元素的第一个和最后一个位置"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array)
在排序数组中查找元素的第一个和最后一个位置  
给定一个按照升序排列的整数数组 nums，和一个目标值 target。  
找出给定目标值在数组中的开始位置和结束位置。  
你的算法时间复杂度必须是 O(log n) 级别。  
如果数组中不存在目标值，返回 [-1, -1]。
```
示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
```
## 分析
### 手写二分法
```go
func searchRange(nums []int, target int) []int {
	leftIndex := search(nums, target) // 或用标准库api： leftIndex := sort.SearchInts(nums, target)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	return []int{leftIndex, searchFromRight(nums, target) - 1}
}
```
search 函数在nums里搜索target，返回新target应该插入的位置；如果nums里已经有target，则在第一个已有target元素之前插入  
nums已经排序，但可能有重复元素  
功能同标准库里的sort.SearchInts(nums, target)
```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] >= target:
			right = mid
		}
	}
	return left
}
```
searchFromRight 在nums里从右向左搜索target，返回新target应该插入的位置；  
如果nums里已经有target，则在最后一个target元素之后插入  
nums已经排序，但可能有重复元素
```go
func searchFromRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] <= target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}
	return left
}
```
### 纯用标准库
实际上 searchFromRight 是可以省去的，找右端点的时候，可以从左往右找第一个 `大于` 目标的位置，最后用这个位置减去 1 就是右端点。
```go
func searchRange(nums []int, target int) []int {
	leftIndex := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	return []int{leftIndex, sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})-1}
}
```