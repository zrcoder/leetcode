---
title: "组合总和"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [回溯, 记忆化搜索, 动态规划]
---

## [39. 组合总和](https://leetcode-cn.com/problems/combination-sum)

难度中等

给定一个**无重复元素**的数组 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的数字可以无限制重复被选取。

**说明：**

- 所有数字（包括 `target`）都是正整数。
- 解集不能包含重复的组合。

**示例 1：**

```
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
```

**示例 2：**

```
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

**提示：**

- `1 <= candidates.length <= 30`
- `1 <= candidates[i] <= 200`
- `candidate` 中的每个元素都是独一无二的。
- `1 <= target <= 500`

## 分析

### 回溯

{{< tabs >}}
{{% tab name="通用写法" %}}
```go
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, i int)
	dfs = func(t, i int) {
		if t < 0 || i == len(candidates) {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}
		// 使用 i 处元素
		cur = append(cur, candidates[i])
		// 元素可无限制重复使用，这里 i 不加1
		dfs(t-candidates[i], i)
		cur = cur[:len(cur)-1]
		// 不使用 i 处元素，这里 i + 1
		dfs(t, i+1)
	}
	dfs(target, 0)
	return res
}
```
{{% /tab %}}
{{% tab name="循环写法" %}}
```go
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, start int)
	dfs = func(t, start int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}
		// 从 start 开始，而不是从 0 开始，防止重复的组合出现
		for j := start; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			dfs(t-candidates[j], j)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return res
}
```
{{% /tab %}}
{{% /tabs %}}

辅助函数：
```go
func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}
```

### 拓展

如果这个问题是问这样的组合数有多少种呢？实际就是完全背包问题。参考代码如下：

```go
func combinations(candidates []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for _, v := range candidates {
        for c := v; c <= target; c++ {
            dp[c] += dp[c-v]
        }
    }
    return dp[target]
}
```

> 1. 注意到 target 的值不是很大。
> 2. 和《组合总和 IV》对比，看看有什么不同。

## [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii)

难度中等

给定一个数组 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的每个数字在每个组合中只能使用一次。

**说明：**

- 所有数字（包括目标数）都是正整数。
- 解集不能包含重复的组合。

**示例 1:**

```
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

**示例 2:**

```
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
```

## 分析

主要难点在去重，可以事先将数组排序，回溯过程中调整策略。

{{< tabs >}}
{{% tab name="通用写法" %}}
```go
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(t, start int)
	backtrack = func(t, start int) {
		if t == 0 {
            res = append(res, copySlice(cur))
			return
		}
		if t < 0 || start == len(candidates) {
			return
		}
		// 选择 start 处的元素
		cur = append(cur, candidates[start])
		backtrack(t-candidates[start], start+1)
		cur = cur[:len(cur)-1]
		// 不选择 start 处的元素
		// 也不能选择紧跟 start 后与 start 处元素相同的元素
		i := start+1
		for i < len(candidates) && candidates[i] == candidates[start] {
			i++
		}
		backtrack(t, i)
	}
	sort.Ints(candidates)
	backtrack(target, 0)
	return res
}
```
{{% /tab %}}
{{% tab name="循环写法" %}}
```go
func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
	var cur []int
	var backtrack func(t, start int)
	backtrack = func(t, start int) {
		if t == 0 {
            res = append(res, copySlice(cur))
			return
		}
		for i := start; i < len(candidates); i++ {
			if t - candidates[i] < 0 {
				return
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(t-candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}
	sort.Ints(candidates)
	backtrack(target, 0)
	return res
}
```
{{% /tab %}}
{{% /tabs %}}

## [216. 组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii)

难度中等

找出所有相加之和为 ***n*** 的 ***k\*** 个数的组合***。\***组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

**说明：**

- 所有数字都是正整数。
- 解集不能包含重复的组合。

**示例 1:**

```
输入: k = 3, n = 7
输出: [[1,2,4]]
```

**示例 2:**

```
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
```

## 分析

通用回溯解法：

```go
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var cur []int
	max := 9
	if n < max {
		max = n
	}
	var backtrack func(target, num int)
	backtrack = func(target, num int) {
		if target == 0 && len(cur) == k { 
			res = append(res, copySlice(cur))
			return
		}
		if target <= 0 || num > max {
			return
		}

		cur = append(cur, num)
		backtrack(target-num, num+1)
		cur = cur[:len(cur)-1]

		backtrack(target, num+1)
	}
	backtrack(n, 1)
	return res
}
```

## [377. 组合总和 Ⅳ](https://leetcode-cn.com/problems/combination-sum-iv/)

难度中等

给你一个由 **不同** 整数组成的数组 `nums` ，和一个目标整数 `target` 。请你从 `nums` 中找出并返回总和为 `target` 的元素组合的个数。

题目数据保证答案符合 32 位整数范围。

**示例 1：**

```
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
```

**示例 2：**

```
输入：nums = [9], target = 3
输出：0
```

**提示：**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 1000`
- `nums` 中的所有元素 **互不相同**
- `1 <= target <= 1000`

**进阶：**如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

## 分析
看题目示例，这个问题实际求的是排列数，而不是组合数。

### 朴素回溯
```go
func combinationSum4Timeout(nums []int, target int) int {
	var res int
	var backtrack func(t int)
	backtrack = func(t int) {
		if t == 0 {
			res++
		}
		if t < 0 {
			return
		}
		for _, v := range nums {
			backtrack(t-v)
		}
	}
	backtrack(target)
	return res
}
```

超时了。

### 记忆化搜索

即给回溯加上备忘录来优化。

```go
func combinationSum4Memo(nums []int, target int) int {
   memo := make(map[int]int, 0)
   var backtrack func(t int) int
   backtrack = func(t int) int {
   	if t == 0 {
   		return 1
   	}
   	if t < 0 {
   		return -1
   	}
   	if v, ok := memo[t]; ok {
   		return v
   	}
   	res := 0
   	for _, v := range nums {
   		if backtrack(t-v) != -1 {
   			res += backtrack(t - v)
   		}
   	}
   	memo[t] = res
   	return res
   }
   res := backtrack(target)
   if res == -1 {
   	return 0
   }
   return res
}
```

AC了。
时间复杂度是 `O(n*target)`, `n` 为 nums 大小
空间复杂度 `O(target)`

### 动态规划

根据以上记忆化解法，可以得到动态规划解法。实际上正是完全背包问题的一种，不过要注意两层循环的顺序，对`target`的枚举放在最外层，想想是为什么。

```go
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, v := range nums {
			if v <= t {
				dp[t] += dp[t-v]
			}
		}
	}
	return dp[target]
}
```

时空复杂度同记忆化搜索方法，不过对于某些输入，空间上有点浪费，不如记忆化里边用 map。

