---
title: "494. 目标和"
date: 2021-06-07T19:42:12+08:00
weight: 50
tags: [动态规划,背包,回溯,深度优先搜索]
---

## [494. 目标和](https://leetcode-cn.com/problems/target-sum/)

难度中等

给你一个整数数组 `nums` 和一个整数 `target` 。

向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：

- 例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

**示例 1：**

```
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
```

**示例 2：**

```
输入：nums = [1], target = 1
输出：1
```

**提示：**

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 100`

函数签名：

```go
func findTargetSumWays(nums []int, target int) int
```

## 分析

### DFS

因为数据范围较小，数组长度在 20 之内，可以直接用 DFS 回溯算法。

```go
func findTargetSumWays(nums []int, target int) int {
	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		return dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
	}
	return dfs(0, 0)
}
```

时间复杂度 `O(2^n)`，其中 `n` 指数组长度。

空间复杂度是 `O(n)`，递归栈的深度不超过 `n`。

### 记忆华搜索

```go
func findTargetSumWays(nums []int, target int) int {
	memo := make(map[string]int, len(nums))
	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		key := fmt.Sprintf("%d,%d", i, sum)
		if v, ok := memo[key]; ok {
			return v
		}
		memo[key] = dfs(i+1, sum+nums[i]) + dfs(i+1, sum-nums[i])
		return memo[key]
	}
	return dfs(0, 0)
}
```

时空复杂度都是 `O(n*target)`。

### 零一背包

假设数组所有元素和为 `sum`，保持正号的元素和为 `reg`，那么保持负号的元素和的绝对值就是 `sum-reg`，要使总和为 `target`，即：

`reg - (sum-reg) = target` 可得 `reg = (target+sum)/2`。

如果 `target+sum` 是奇数，可以直接返回 `0`；否则，问题相当于：在 `nums` 中选取若干元素，使其和为 `reg`。这是个典型的零一背包问题。

还可以优化时空复杂度：假设保持负号的元素和的绝对值为 `neg`， 那么保持正号的元素和为 `sum-neg`，由 `(sum-neg)-neg = target`得到 `neg = (sum-target)/2`，如果 `sum-target` 是奇数，可以直接返回 `0`；否则，问题相当于：在 `nums` 中选取若干元素，使其和为 `neg`，这比上边的 `reg` 规模更小。

```go
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < target || -sum > target {
		return 0
	}
	
	diff := sum - target
	if diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}
```

时间复杂度：`O(n*(sum-target))`，空间复杂度 `O(sum-target)`。