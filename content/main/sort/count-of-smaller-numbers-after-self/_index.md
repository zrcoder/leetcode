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
    return mergeSort(nums)
}

func mergeSort(arr []int) int {
    n := len(arr)
    if n < 2 {
        return 0
    }

    left := append([]int{}, arr[:n/2]...)
    right := append([]int{}, arr[n/2:]...)

    return mergeSort(left) + mergeSort(right) + merge(arr, left, right)
}

func merge(arr, left, right []int) int {
    res := 0
    var i, j, k int
    for ; i < len(left) || j < len(right); k++ {
        if j == len(right) || i < len(left) && left[i] <= right[j] {
            res += j
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
    }
    return res
}
```
时间复杂度：`O(nlogn)`, 空间复杂度：`O(n)`。

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
	pairs := make([]pair, len(nums)) // 记录每个元素的值和索引,以免在排序过程中打乱顺序
	for i, v := range nums {
		pairs[i] = pair{val: v, index: i}
	}
	count := make([]int, len(nums))
	mergeSort(pairs, count)
	return count
}

func mergeSort(pairs []pair, count []int) {
    n := len(pairs)
	if n < 2 {
		return
	}
	left := append([]pair{}, pairs[:n/2]...)
    right := append([]pair{}, pairs[n/2:]...)
	mergeSort(left, count)
	mergeSort(right, count)
	merge(left, right, pairs, count)
}

func merge(left, right, pairs []pair, count []int) {
	var i, j, k int
	for ; i < len(left) || j < len(right); k++ {
		if j == len(right) || i < len(left) && left[i].val <= right[j].val {
			count[left[i].index] += j // left[i]的值要比 right[0:j]共j个值大
			pairs[k] = left[i]
			i++
		} else {
			pairs[k] = right[j]
			j++
		}
	}
}
```
时间复杂度：`O(nlogn)`, 空间复杂度：`O(n)`。