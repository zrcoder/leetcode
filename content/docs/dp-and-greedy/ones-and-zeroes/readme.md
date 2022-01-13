---
title: "474. 一和零"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [474. 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/)

难度中等

给你一个二进制字符串数组 `strs` 和两个整数 `m` 和 `n` 。

请你找出并返回 `strs` 的最大子集的大小，该子集中 **最多** 有 `m` 个 `0` 和 `n` 个 `1` 。

如果 `x` 的所有元素也是 `y` 的元素，集合 `x` 是集合 `y` 的 **子集** 。

**示例 1：**

```
输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
```

**示例 2：**

```
输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
```

**提示：**

- `1 <= strs.length <= 600`
- `1 <= strs[i].length <= 100`
- `strs[i]` 仅由 `'0'` 和 `'1'` 组成
- `1 <= m, n <= 100`

函数签名：

```go
func findMaxForm(strs []string, m int, n int) int
```

## 分析

### 回溯

可以参考[子集系列问题]()，比较方便地写出回溯法代码如下：

```go
func findMaxForm(strs []string, m int, n int) int {
	var res, curLen, cur0s, cur1s int
	zeros := countZeros(strs)
	var backtrack func(i int)
	backtrack = func(i int) {
		if cur0s > m || cur1s > n {
			return
		}
		if curLen > res {
			res = curLen
		}
		if i == len(strs) {
			return
		}
		// 不使用 i 处字符串
		backtrack(i + 1)

		// 使用 i 处字符串
		c0 := zeros[i]
		c1 := len(strs[i]) - c0
		curLen, cur0s, cur1s = curLen+1, cur0s+c0, cur1s+c1
		backtrack(i + 1)
		curLen, cur0s, cur1s = curLen-1, cur0s-c0, cur1s-c1
	}
	backtrack(0)
	return res
}

func countZeros(strs []string) []int {
	res := make([]int, len(strs))
	for i, v := range strs {
		res[i] = strings.Count(v, "0")
	}
	return res
}
```

回溯函数里的参数 i 表示从 strs 的位置 i 开始尝试，之前的元素已经尝试过，不再考虑。

假设 strs 长度为 N，根据输入规模限制，可以忽略统计每个字符串中 0 的个数，时间主要花费在回溯上，每个字符串有选择和不选择两种情况，时间复杂度是 `O(2^N)`。

复杂度过高，实际在 LeetCode 也超时了，只过了一小半用例。

### 加上备忘录优化回溯

朴素回溯解法之所以复杂度高，是因为有非常多的重复计算，可以加上备忘录来优化。

在上边的代码基础上加备忘录比较艰难，必须先修改回溯函数的参数，并加上返回值。

部分定义函数 `var dfs func(m, n, i int) int` 表示限制最多 m 个0，n 个 1，在 strs 能找到的字符串最大个数；其中 i 的意义同一开始的回溯函数里的 i。

这样就可以加一个三维的备忘录来优化，得到记忆华搜索解法：

```go
func findMaxForm(strs []string, m int, n int) int {
	zeros := countZeros(strs)
	memo := make([][][]int, m+1)
	for i := range memo {
		memo[i] = make([][]int, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, len(strs))
		}
	}
	var dfs func(m, n, i int) int
	dfs = func(m, n, i int) int {
		if m < 0 || n < 0 {
			return -1
		}
		if i == len(strs) {
			return 0
		}
		if memo[m][n][i] > 0 {
			return memo[m][n][i]
		}
		c0 := zeros[i]
		c1 := len(strs[i]) - c0
		memo[m][n][i] = max(1+dfs(m-c0, n-c1, i+1), dfs(m, n, i+1))
		return memo[m][n][i]
	}
	return dfs(m, n, 0)
}
```

时间、空间复杂度都是 `O(mnN)`，其中 N 是 strs 的长度。

### 动态规划

根据上边自顶向下的记忆化搜索解法，就可以想到自底向上的动态规划解法。

实际上也可以立马联想到一类特殊的动态规划问题，即零一背包问题。且可以把空间复杂度降低为 `O(mn)`，在这个问题的输入约束下，这个优化非常大。

定义二维 dp 数组， `dp[i][j]` 代表限定 i 个 0、j 个 1 在 strs 中最多能找到多少个字符串。

```go
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for _, v := range strs {
		str0s := strings.Count(v, "0")
		str1s := len(v) - str0s
		for c0 := m; c0 >= str0s; c0-- {
			for c1 := n; c1 >= str1s; c1-- {
				dp[c0][c1] = max(dp[c0][c1], 1+dp[c0-str0s][c1-str1s])
				res = max(res, dp[c0][c1])
			}
		}
	}
	return res
}
```

时间复杂度 `O(mnN)`，其中 N 是 strs 的长度。

空间复杂度 `O(mn)`。