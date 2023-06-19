---
title: 1262. 可被三整除的最大和
---

## [1262. 可被三整除的最大和](https://leetcode.cn/problems/greatest-sum-divisible-by-three) (Medium)

给你一个整数数组 `nums`，请你找出并返回能被三整除的元素最大和。

**示例 1：**

```
输入：nums = [3,6,5,1,8]
输出：18
解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
```

**示例 2：**

```
输入：nums = [4]
输出：0
解释：4 不能被 3 整除，所以无法选出数字，返回 0。

```

**示例 3：**

```
输入：nums = [1,2,3,4,4]
输出：12
解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。

```

**提示：**

- `1 <= nums.length <= 4 * 10^4`
- `1 <= nums[i] <= 10^4`

## 分析

动态规划，仅需遍历一次数组，记录累加和，将累加和分为三类：模3得0、模3得1、模3得2. 这三类之间可以转移，最终返回第一类的和。

时间复杂度 O(n), 空间复杂度 O(1), 其中 n 是数组长度。

```go
func maxSumDivThree(nums []int) int {
	r := [3]int{} // r[i] 记录模3得i的累加和
	for _, v := range nums {
		a := r[0] + v
		b := r[1] + v
		c := r[2] + v
		r[a%3] = max(r[a%3], a)
		r[b%3] = max(r[b%3], b)
		r[c%3] = max(r[c%3], c)
	}
	return r[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```
