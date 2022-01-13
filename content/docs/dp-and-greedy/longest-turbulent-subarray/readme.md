---
title: "978. 最长湍流子数组"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [978. 最长湍流子数组](https://leetcode-cn.com/problems/longest-turbulent-subarray/)

难度中等

当 `A` 的子数组 `A[i], A[i+1], ..., A[j]` 满足下列条件时，我们称其为*湍流子数组*：

- 若 `i <= k < j`，当 `k` 为奇数时， `A[k] > A[k+1]`，且当 `k` 为偶数时，`A[k] < A[k+1]`；
- **或** 若 `i <= k < j`，当 `k` 为偶数时，`A[k] > A[k+1]` ，且当 `k` 为奇数时， `A[k] < A[k+1]`。

也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 `A` 的最大湍流子数组的**长度**。



**示例 1：**

```
输入：[9,4,2,10,7,8,8,1,9]
输出：5
解释：(A[1] > A[2] < A[3] > A[4] < A[5])
```

**示例 2：**

```
输入：[4,8,12,16]
输出：2
```

**示例 3：**

```
输入：[100]
输出：1
```



**提示：**

1. `1 <= A.length <= 40000`
2. `0 <= A[i] <= 10^9`

函数签名：

```go
func maxTurbulenceSize(arr []int) int
```

## 分析

可以用双指针滑动窗口或动态规划的解法。时间复杂度都是 O(n)，空间复杂度都是常数级。

### 滑动窗口

```go
const (
	equal = iota
	lager
	smaller
	unknown
)

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	flag, res, cur := unknown, 1, 1
	for i, j := 0, 1; j < len(arr); j++ {
		curFlag := equal
		if arr[j] > arr[j-1] {
			curFlag = lager
		} else if arr[j] < arr[j-1] {
			curFlag = smaller
		}
		if curFlag == equal {
			cur = 1
			i = j
		} else if flag == curFlag {
			cur = 2
			i = j - 1
		} else {
			cur = j - i + 1
		}
		flag = curFlag

		if cur > res {
			res = cur
		}
	}
	return res
}
```

### 动态规划

```go
func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	// 分别记录以当前元素结尾，且当前元素比前一个元素大/小的子序列的长度
	larger, smaller := 1, 1
	res := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			larger, smaller = smaller+1, 1
		} else if arr[i] < arr[i-1] {
			larger, smaller = 1, larger+1
		} else {
			larger, smaller = 1, 1
		}
		res = max(res, max(larger, smaller))
	}
	return res
}
```
