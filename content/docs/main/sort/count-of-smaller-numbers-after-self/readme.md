---
title: "数组中的逆序对"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [滑动窗口, 桶排序]
---

## [面试题51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。  
输入一个数组，求出这个数组中的逆序对的总数。
```
示例 1:
输入: [7,5,6,4]
输出: 5

限制：
0 <= 数组长度 <= 50000
```
## 分析
两层循环的朴素实现复杂度是 `O(n^2)`;  
可用归并排序，在归并过程中统计逆序对的个数或统计 counts 数组，时间复杂度降为 `O(nlgn)`
```go
func reversePairs(nums []int) int {
	var count int
	mergeSort(nums, &count)
	return count
}

func mergeSort(nums []int, count *int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	left := make([]int, mid)
	right := make([]int, len(nums)-mid)
	_ = copy(left, nums[:mid])
	_ = copy(right, nums[mid:])
	mergeSort(left, count)
	mergeSort(right, count)
	merge(left, right, nums, count)
}

func merge(left, right, nums []int, count *int) {
	var i, j, k int
	for ; i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] {
			*count += j // left[i] 要比right[0:j]共j个元素大
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
	for ; i < len(left); i, k = i+1, k+1 {
		*count += j // 左侧剩余的元素同样要比j个（等于len（right））right部分元素大
		nums[k] = left[i]
	}
	for ; j < len(right); j, k = j+1, k+1 {
		nums[k] = right[j]
	}
}
```
## [315. 计算右侧小于当前元素的个数](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/)
给定一个整数数组 nums，按要求返回一个新数组 counts。
数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例:

输入: [5,2,6,1]
输出: [2,1,1,0]
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.
## 分析
思路同上个问题
```go
type pair struct {
	val, index int
}

func countSmaller(nums []int) []int {
	n := len(nums)
	pairs := make([]pair, n) // 记录每个元素的值和索引,以免在排序过程中打乱顺序
	for i, v := range nums {
		pairs[i] = pair{val: v, index: i}
	}
	count := make([]int, n)
	mergeSort(pairs, count)
	return count
}

func mergeSort(pairs []pair, count []int) {
	if len(pairs) < 2 {
		return
	}
	mid := len(pairs) / 2
	left := make([]pair, mid)
	right := make([]pair, len(pairs)-mid)
	_ = copy(left, pairs[:mid])
	_ = copy(right, pairs[mid:])
	mergeSort(left, count)
	mergeSort(right, count)
	merge(left, right, pairs, count)
}

func merge(left, right, pairs []pair, count []int) {
	var i, j, k int
	for ; i < len(left) && j < len(right); k++ {
		if left[i].val <= right[j].val {
			count[left[i].index] += j // left[i]的值要比 right[0:j]共j个值大
			pairs[k] = left[i]
			i++
		} else {
			pairs[k] = right[j]
			j++
		}
	}
	for ; i < len(left); i, k = i+1, k+1 {
		count[left[i].index] += j // 左侧剩余的元素同样要比j个（等于len（right））right部分元素大
		pairs[k] = left[i]
	}
	for ; j < len(right); j, k = j+1, k+1 {
		pairs[k] = right[j]
	}
}
```