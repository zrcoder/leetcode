---
title: 1493. 删掉一个元素以后全为 1 的最长子数组
date: 2023-07-05T11:42:57+08:00
---

## [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element) (Medium)

给你一个二进制数组 `nums` ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。

**提示 1：**

```
输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
```

**示例 2：**

```
输入：nums = [0,1,1,1,0,1,1,0,1]
输出：5
解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
```

**示例 3：**

```
输入：nums = [1,1,1]
输出：2
解释：你必须要删除一个元素。
```

**提示：**

- `1 <= nums.length <= 10⁵`
- `nums[i]` 要么是 `0` 要么是 `1` 。

## 分析

维护一个滑动窗口, 其中最多1个0. 仅需要遍历一次就可以找到满足题意的最长子串.

时间复杂度: O(n),  空间复杂度: O(1).

```go
func longestSubarray(nums []int) int {
	lo := 0 // 窗口左端点
	zeros := 0
	res := 0
	for hi := range nums { // 窗口右端点
		if nums[hi] == 0 {
			zeros++
		}
		for zeros == 2 {
			if nums[lo] == 0 {
				zeros--
			}
			lo++
		}
		res = max(res, hi-lo)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```
