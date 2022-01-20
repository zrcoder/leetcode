---
title: "归并排序"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

归并排序是分治思想的典型应用。有一类巧妙的问题，在归并排序过程中做一点额外的工作，往往可以较低的代价解决它们。

## [327. 区间和的个数](https://leetcode-cn.com/problems/count-of-range-sum/)

难度困难

给定一个整数数组 `nums`，返回区间和在 `[lower, upper]` 之间的个数，包含 `lower` 和 `upper`。
区间和 `S(i, j)` 表示在 `nums` 中，位置从 `i` 到 `j` 的元素之和，包含 `i` 和 `j` (`i` ≤ `j`)。

**说明:**
最直观的算法复杂度是 *O*(*n*^2) ，请在此基础上优化你的算法。

**示例:**

```
输入: nums = [-2,5,-1], lower = -2, upper = 2,
输出: 3 
解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。
```

函数签名：

```go
func countRangeSum(nums []int, lower int, upper int) int
```

## 分析

### 方法零：朴素解法

```go
func countRangeSum(nums []int, lower int, upper int) int {
    var res = 0
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum >= lower && sum <= upper {
                res++
            }
        }
    }
    return res
}
```

时间复杂度 `O(n^2)`， 空间复杂度 `O(1)`。

### 方法一：分治思想

设前缀和数组为 *preSum*，则问题等价于求所有的下标对 (*i*,*j*)，满足

```
preSum[j] − preSum[i-1] ∈ [lower,upper]
```

> 在前缀和头部插入值为 0 的元素，可简化边界处理

在朴素解法基础上加上前缀和技巧做一个改变：

```go
func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	var res = 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			s := preSum[j+1] - preSum[i]
			if s >= lower && s <= upper {
				res++
			}
		}
	}
	return res
}
```
这个改动不但没有优化时间复杂度，还额外增加了空间复杂度。当前于事无补，但是借助前缀和技巧，可以用分治的思想优化时间复杂度。

将 `preSum` 划分为左右两个子数组 n1、 n2，可分别求出n1、 n2 中满足要求的下标对个数，相加后再加上 一个坐标在 n1 而另一个坐标在 n2 且满足题意的坐标对个数，就得到了结果：

```
0, 1, ..., n/2-1 | n/2, n/2 + 1, ..., n-1
<----- n1 ------>|<--------  n2 -------->

result(all) = result(n1) + result(n2) + count(i, j) (i ∈ n1， j ∈ n2)
```

对于 count 函数，如果左右子数组都有序，问题将变得简单，时间复杂度也会降为线性：

对于两个**升序排列**的数组 *n*1、*n*2，找出所有的下标对 (i,j)，满足

```
n2[j] − n1[i] ∈ [lower,upper]
```

两个数组有序。可以在 n2 中维持两个指针 l 、r，起初 l 指向 n2 的初始位置， r 将在随后被赋值。

先考察 n1 的第一个元素。先不断右移 l 直到 `n2[l] >= n1[0] + lower`，此时 l 及其右侧元素都不小于 `n1[0] + lower`；随后使 r = l 并不断右移 r 指针， 直到 `n2[r] > n1[0] + upper`， 此时 r 左侧元素都不大于 `n1[0] + upper`。至此区间 `[l, r)` 中所有索引 j 都满足

```
n2[j] - n1[0] ∈ [lower, upper]
```

接下来考察 n1 的第二个元素。因 n1 递增， 不难发现 l、r 只能向右移动。

因此，不断执行上述过程，对于 n1 中的每一个下标，都记录对应的区间 [l, r) 的大小。最终就能统计得到满足条件的下表对 (i, j) 的数量。

**整个过程即是对 `preSum` 数组归并排序，期间做一点额外的工作，即统计满足题意的坐标对的个数。**

```go
var lo, hi int

func countRangeSum(nums []int, lower, upper int) int {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}
	lo, hi = lower, upper
	return mergeCount(preSum)
}

func mergeCount(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}
	n1 := append([]int{}, arr[:n/2]...)
	n2 := append([]int{}, arr[n/2:]...)
	cnt := mergeCount(n1) + mergeCount(n2) // 递归完成后， n1、n2 均为有序
	// 统计分别在 n1、 n2 中满足要求的下标对的数量
	cnt += calPairs(n1, n2)
	// n1、n2 归并填入 arr，使 arr 有序
	merge(arr, n1, n2)
	return cnt
}

func calPairs(n1, n2 []int) int {
	res := 0
	var l, r int
	for _, v := range n1 {
		for ; l < len(n2) && n2[l] < v+lo; l++ {
		}
		for r = l; r < len(n2) && n2[r] <= v+hi; r++ {
		}
		res += r - l
	}
	return res
}

func merge(arr, n1, n2 []int) {
	p1, p2 := 0, 0
	for i := range arr {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			arr[i] = n1[p1]
			p1++
		} else {
			arr[i] = n2[p2]
			p2++
		}
	}
}
```

#### 复杂度分析

时间复杂度 `O(nlogn)`， 其中 `n` 是数组长度。设执行时间为 `T(n)`, 则两次递归调用的时间分别为 `T(n/2)`， 还需要 `O(n)` 的时间求下标对梳理并合并数组，所以 `T(n) = 2*T(n) + O(n)`，根据主定理有 `T(n) = O(nlogn)`。

空间复杂度 `O(n)`。设空间占用 `M(n)`， 递归栈空间为 `M(n/2)`， 合并数组需要空间 `O(n)`，所以 `M(n) = M(n/2) + O(n)`，根据主定理，有 `M(n) = O(n)`。

## [493. 翻转对](https://leetcode-cn.com/problems/reverse-pairs/)

难度困难

给定一个数组 `nums` ，如果 `i < j` 且 `nums[i] > 2*nums[j]` 我们就将 `(i, j)` 称作一个***重要翻转对***。

你需要返回给定数组中的重要翻转对的数量。

**示例 1:**

```
输入: [1,3,2,3,1]
输出: 2
```

**示例 2:**

```
输入: [2,4,3,5,1]
输出: 3
```

**注意:**

1. 给定数组的长度不会超过`50000`。
2. 输入数组中的所有数字都在32位整数的表示范围内。

函数签名：

```go
func reversePairs(nums []int) int
```

## 分析

朴素实现显然是 O(n^2) 的时间复杂度。

如果只需要找到 i 后边的第一个满足 nums[i] > nums[j]*2 的 j，可以使用单调栈在线性时间复杂度内解决。但是要找出所有的 j，则没这么简单。

可以复用上面区间和个数问题的思路，用归并排序的方式解决。代码基本相同，只有 calPairs 函数实现逻辑不一样：

```go
func reversePairs(nums []int) int {
    return mergeCount(nums)
}

// mergeCount、merge 同上，略

func calPairs(n1, n2 []int) int {
    res := 0
    i, j := 0, 0
    for i < len(n1) && j < len(n2) {
        if n1[i] > 2*n2[j] {
            res += len(n1)-i
            j++
        } else {
            i++
        }
    }
    return res
}
```

 复杂度也同上个问题。
