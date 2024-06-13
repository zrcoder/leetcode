---
title: 368. 最大整除子集
date: 2024-02-10T16:50:45+08:00
---

## [368. 最大整除子集](https://leetcode.cn/problems/largest-divisible-subset) (Medium)

给你一个由 **无重复** 正整数组成的集合 `nums` ，请你找出并返回其中最大的整除子集 `answer` ，子集中每一元素对 `(answer[i], answer[j])` 都应当满足：

- `answer[i] % answer[j] == 0` ，或
- `answer[j] % answer[i] == 0`

如果存在多个有效解子集，返回其中任何一个均可。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。

```

**示例 2：**

```
输入：nums = [1,2,4,8]
输出：[1,2,4,8]

```

**提示：**

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 2 * 10⁹`
- `nums` 中的所有整数 **互不相同**

## 分析

### 动态规划

如果仅需要返回最大长度，可以先排序，再通过双重循环的动态规划解决，现在需要能返回具体的子集，稍微复杂点，参考[最长上升子序列](/main/dp-and-greedy/longest-increasing-subsequence)中的说明来做。

```go
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	mi := 0
	for j, v := range nums {
		dp[j] = 1
		for i, u := range nums[:j] {
			if v%u == 0 && dp[i]+1 > dp[j] {
				dp[j] = dp[i] + 1
			}
		}
		if dp[j] > dp[mi] {
			mi = j
		}
	}
	res := make([]int, dp[mi])
	i := len(res) - 1
	res[i] = nums[mi]
	for j := mi - 1; j >= 0 && i > 0; j-- {
		if res[i]%nums[j] == 0 && i == dp[j] {
			i--
			res[i] = nums[j]
		}
	}
	return res
}

```
