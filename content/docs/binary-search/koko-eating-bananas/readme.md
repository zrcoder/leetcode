---
title: "875. 爱吃香蕉的珂珂"
date: 2021-04-28T19:47:33+08:00
weight: 49
tags: [二分搜索]
---

## [875. 爱吃香蕉的珂珂](https://leetcode-cn.com/problems/koko-eating-bananas/)

难度中等

珂珂喜欢吃香蕉。这里有 `N` 堆香蕉，第 `i` 堆中有 `piles[i]` 根香蕉。警卫已经离开了，将在 `H` 小时后回来。

珂珂可以决定她吃香蕉的速度 `K` （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 `K` 根。如果这堆香蕉少于 `K` 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 `H` 小时内吃掉所有香蕉的最小速度 `K`（`K` 为整数）。

**示例 1：**

```
输入: piles = [3,6,7,11], H = 8
输出: 4
```

**示例 2：**

```
输入: piles = [30,11,23,4,20], H = 5
输出: 30
```

**示例 3：**

```
输入: piles = [30,11,23,4,20], H = 6
输出: 23
```

**提示：**

- `1 <= piles.length <= 10^4`
- `piles.length <= H <= 10^9`
- `1 <= piles[i] <= 10^9`

函数签名：

```go
func minEatingSpeed(piles []int, h int) int
```

## 分析

首先可以确定，珂珂吃香蕉的速度是在这样的一个区间里：[1, max]，其中 max 是数组里最大元素。最小速度每小时 1 个香蕉，最大每小时 max 个，速度再大没有用。

对于一个给定的速度 k，怎么判断能不能在限定时间内吃完所有香蕉？只要计算实际花费的时间，和限定时间比较即可。

假设对于速度 k，花费的时间是 cost(k)，容易发现这是个单点递减函数。既然有单调性，那么可以用二分法来确定 k。

```go
func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	for _, v := range piles {
		if v > hi {
			hi = v
		}
	}
	hi++
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !check(mid, h, piles) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func check(k, h int, piles []int) bool {
	cost := 0
	for _, v := range piles {
		cost += (v+k-1)/k
		if cost > h {
			return false
		}
	}
	return true
}
```

时间复杂度 `O(NlogMax)`，其中 `N` 是数组长度，`Max` 是数组最大元素值。