---
title: "1365. 有多少小于当前数字的数字"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [排序, 计数排序]
---

## [1365. 有多少小于当前数字的数字](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/)

难度简单

给你一个数组 `nums`，对于其中每个元素 `nums[i]`，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 `nums[i]` 你必须计算出有效的 `j` 的数量，其中 `j` 满足 `j != i` **且** `nums[j] < nums[i]` 。

以数组形式返回答案。

 

**示例 1：**

```
输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]
解释： 
对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。 
对于 nums[1]=1 不存在比它小的数字。
对于 nums[2]=2 存在一个比它小的数字：（1）。 
对于 nums[3]=2 存在一个比它小的数字：（1）。 
对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
```

**示例 2：**

```
输入：nums = [6,5,4,8]
输出：[2,1,0,3]
```

**示例 3：**

```
输入：nums = [7,7,7,7]
输出：[0,0,0,0]
```

 

**提示：**

- `2 <= nums.length <= 500`
- `0 <= nums[i] <= 100`

函数签名：

```go
func smallerNumbersThanCurrent(nums []int) []int
```

## 分析

### 朴素实现

```go
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		for j, w := range nums {
			if i == j {
				continue
			}
			if w < v {
				res[i]++
			}
		}
	}
	return res
}
```

时间复杂度 O(n^2)，空间复杂度  O(1)。

### 标准库排序

可以借助标准库先排序，再针对每个元素用二分法找到比其小的元素个数。

```go
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	for i, v := range nums {
		res[i] = sort.Search(len(nums), func(i int) bool {
			return tmp[i] >= v
		})
	}
	return res
}
```

时间复杂度 O(nlogn)，时标准库排序函数的复杂度。

空间复杂度 O(n)，需要额外的数组 tmp 做排序。

### 计数排序

注意到数组元素的取值范围是 0-100，排序使用计数排序非常合适，可以将时间复杂度降低到线性。

```go
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}
	for i, v := range nums {
		if v == 0 {
			continue
		}
		res[i] = cnt[v-1]
	}
	return res
}
```

时间复杂 O(n+k), 空间复杂度 O(k)。其中 k 是值域的大小，在这个问题约束中最大为100。

## 拓展

如果把问题拓展到二维矩阵里，找到每个位置同行同列比自己小的元素个数呢？

> 显然朴素实现的复杂度是 O(n^3)，可以用排序的方式将为 O(n^2logn)，如果元素值比较集中，如像这个问题中的 0-100，也可以用技术排序将复杂度降到  O(n^2)。