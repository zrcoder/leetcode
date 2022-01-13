---
title: "4. 寻找两个有序数组的中位数"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [4. 寻找两个有序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays)
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。  
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。  
你可以假设 nums1 和 nums2 不会同时为空。
```
示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
```
## 分析
对于一个有序数组，如果元素个数为奇数，中位数即中间元素的值；  
若元素个数为偶数，中位数为中间两个元素的平均值。  
对于两个或多个有序数组，其合并后的中位数并非每个数组中位数的平均值，如：

    [1, 3, 5] // 中位数3
    [8, 10] // 中位数9
    // 合并后的数组
    [1, 3, 5, 8, 10] // 中位数5, 并非3和9的平均数
    
所以，必须对两个数组合并，合并后依然有序

这样朴素实现的时间与空间复杂度均为O(m+n)。
### 1. 改进朴素实现，不用真的 merge
```go
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	lastR, currentR := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		lastR = currentR
		if start1 < m && (start2 >= n || nums1[start1] < nums2[start2]) {
			currentR = nums1[start1]
			start1++
		} else {
			currentR = nums2[start2]
			start2++
		}
	}
	if (m+n)%2 == 1 {
		return float64(currentR)
	}
	return float64(lastR+currentR) * 0.5
}
```
时间复杂度无太大改进，空间复杂度改进为 O(1)。
### 2. 问题转化为求数组第 k个元素
对于两个数组，假设长度分别是m、n，求合并后的中位数即求：

    i. 合并后第(m+n)/2 + 1 个元素（m+n为奇数）
    ii. 合并后第(m+n)/2 个元素与第(m+n)/2 + 1个元素的平均值（m+n为偶数）

令 k = (m+n)/2，可以分别取两个数组第k/2个元素，通过比较这两个元素的大小，可以批量地减少搜索范围

    1.如果 nums1[k/2] < nums2[k/2], 说明合并后的第 k 个元素肯定不在 nums1[0:k/2+1] 区间里
    可以继续在 nums1[k/2+1:] 和 nums2 中搜索第 k-(k/2+1) 个元素
    2.反之，可以排除 nums2 的前 k/2 个元素继续搜索

> 对于总数为奇数和偶数的两种情况需要稍作区分。

时间 `O(log(m+n))`，空间`O(1)`
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	if size == 0 {
		return 0.0
	}
	if size%2 == 1 {
		return getKth(nums1, nums2, size/2+1)
	}
	return (getKth(nums1, nums2, size/2) + getKth(nums1, nums2, size/2+1)) * 0.5
}
func getKth(nums1, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return getKth(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	i, j := min(m-1, k/2-1), min(n-1, k/2-1)
	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2[j+1:], k-(j+1))
	}
	return getKth(nums1[i+1:], nums2, k-(i+1))
}
```
### 3. 二分法
用i，j两个指针将两个数组分别划分为两部分，  
将 nums1 的左半部分和 nums2 的左半部分合起来看作合并后的左半部分，  
同样可以得到合并后右半部分  
    
                            |
    nums1       0, ..., i-1,| i, ..., m-1
                            |
    nums2 0, 1, ...,    j-1,| j, ..., n-1
                            |
                  左半部分   |  右半部分

如果能保证左右部分的大小相当（m+n为偶数则左右部分大小相等；为奇数则左半部分比右半部分多一个），  
也就找到了合并后的中位数

    m+n为偶数时：
    i+j = m-i + n-j 即i+j = （m+n）/2
    m+n为奇数时：
    i+j = m-i + n-j + 1也就是 i+j = (m+n+1)/2
    因整数除法特性，可以统一为i+j = (m+n+1)/2
    
    注意到确定了i，就确定了j， j = (m+n+1)/2 - i；

数组已排序，用二分搜索法来确定i:

    因为两个数组都是有序的，所以 nums1[i-1] <= nums1[i]，nums2[i-1] <= nums2[i] 是天然具备的，
    所以只需要保证 nums2[j-1] < = nums1[i] 和 nums1[i-1] <= nums2[j];对不满足的情况分两种情况讨论：
    nums2[j-1] > nums1[i]
    此时需要增加i
    nums1[i-1] > nums2[j]
    此时要减少i

时间`O(log(min(m,n)))`，空间`O(1)`

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		// 降低时间复杂度，同时方便后边边界处理
		return findMedianSortedArrays(nums2, nums1)
	}
	lo, hi := 0, m
	for lo <= hi {
		i := lo + (hi-lo)/2
		j := (m+n+1)/2 - i
		if i < m && j > 0 && nums1[i] < nums2[j-1] {
			lo = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] {
			hi = i - 1
		} else {
			return cal(nums1, nums2, i)
		}
	}
	return 0
}

func cal(nums1, nums2 []int, i int) float64 {
	m, n := len(nums1), len(nums2)
	j := (m+n+1)/2 - i
	leftMax := 0
	if i == 0 {
		leftMax = nums2[j-1]
	} else if j == 0 {
		leftMax = nums1[i-1]
	} else {
		leftMax = max(nums1[i-1], nums2[j-1])
	}

	if (m+n)%2 == 1 {
		return float64(leftMax)
	}
	rightMin := 0
	if i == m {
		rightMin = nums2[j]
	} else if j == n {
		rightMin = nums1[i]
	} else {
		rightMin = min(nums1[i], nums2[j])
	}
	return float64(leftMax+rightMin) * 0.5
}
```
