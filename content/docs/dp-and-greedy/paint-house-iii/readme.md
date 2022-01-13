---
title: "1473. 粉刷房子 III"
date: 2021-05-11T19:55:45+08:00
weight: 50
tags: [记忆化搜索, 动态规划]
---

#### [1473. 粉刷房子 III](https://leetcode-cn.com/problems/paint-house-iii/)

难度困难

在一个小城市里，有 `m` 个房子排成一排，你需要给每个房子涂上 `n` 种颜色之一（颜色编号为 `1` 到 `n` ）。有的房子去年夏天已经涂过颜色了，所以这些房子不可以被重新涂色。

我们将连续相同颜色尽可能多的房子称为一个街区。（比方说 `houses = [1,2,2,3,3,2,1,1]` ，它包含 5 个街区 ` [{1}, {2,2}, {3,3}, {2}, {1,1}]` 。）

给你一个数组 `houses` ，一个 `m * n` 的矩阵 `cost` 和一个整数 `target` ，其中：

- `houses[i]`：是第 `i` 个房子的颜色，**0** 表示这个房子还没有被涂色。
- `cost[i][j]`：是将第 `i` 个房子涂成颜色 `j+1` 的花费。

请你返回房子涂色方案的最小总花费，使得每个房子都被涂色后，恰好组成 `target` 个街区。如果没有可用的涂色方案，请返回 **-1** 。 

**示例 1：**

```
输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
输出：9
解释：房子涂色方案为 [1,2,2,1,1]
此方案包含 target = 3 个街区，分别是 [{1}, {2,2}, {1,1}]。
涂色的总花费为 (1 + 1 + 1 + 1 + 5) = 9。
```

**示例 2：**

```
输入：houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
输出：11
解释：有的房子已经被涂色了，在此基础上涂色方案为 [2,2,1,2,2]
此方案包含 target = 3 个街区，分别是 [{2,2}, {1}, {2,2}]。
给第一个和最后一个房子涂色的花费为 (10 + 1) = 11。
```

**示例 3：**

```
输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[1,10],[10,1],[1,10]], m = 5, n = 2, target = 5
输出：5
```

**示例 4：**

```
输入：houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
输出：-1
解释：房子已经被涂色并组成了 4 个街区，分别是 [{3},{1},{2},{3}] ，无法形成 target = 3 个街区。
```

 **提示：**

- `m == houses.length == cost.length`
- `n == cost[i].length`
- `1 <= m <= 100`
- `1 <= n <= 20`
- `1 <= target <= m`
- `0 <= houses[i] <= n`
- `1 <= cost[i][j] <= 10^4`

函数签名：

```go
func minCost(houses []int, cost [][]int, m int, n int, target int) int
```

## 分析

### 记忆化搜索

回溯穷举是最容易写的，但肯定会超时，也可以尝试加上备忘录来优化，将朴素回溯修改成记忆化搜索。

朴素回溯的实现要改成记忆化搜索，需要尽可能能抽象出函数入参，在这个问题里，入参对应的维度有：当前处理的房间数、当前已经形成的街区数、以及当前最后一个房间的颜色，当然函数应该返回这三个状态对应的最小花费。

```go
const inf = math.MaxInt32

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, target+1)
			fillInts(memo[i][j], -1)
		}
	}
	var dfs func(i, preColor, areas int) int
	dfs = func(i, preColor, areas int) int {
		if areas > target {
			return inf
		}
		if i == m {
			return choose(areas == target, 0, inf)
		}
		if memo[i][preColor][areas] != -1 {
			return memo[i][preColor][areas]
		}
		res := inf
		targetColor := houses[i]
		if targetColor == 0 {
			for c := 1; c <= n; c++ {
				tmpAreas := choose(i == 0, 1, choose(preColor == c, areas, areas+1))
				res = min(res, dfs(i+1, c, tmpAreas)+cost[i][c-1])
			}
		} else {
			tmpAreas := choose(i == 0, 1, choose(preColor == targetColor, areas, areas+1))
			res = min(res, dfs(i+1, targetColor, tmpAreas))
		}
		memo[i][preColor][areas] = res
		return res
	}

	res := dfs(0, 0, 1)
	if res == inf {
		return -1
	}
	return res
}
```

辅助函数：

```go
func fillInts(s []int, val int) {
	for i := range s {
		s[i] = val
	}
}

// 一个类似三目运算符的函数~
func choose(chooseFirst bool, first, second int) int {
	if chooseFirst {
		return first
	}
	return second
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

时间复杂度 O(m⋅n^2 ⋅target)， 空间复杂度 O(m⋅n⋅target)。

### 动态规划

再进一步，可以把记忆化搜索改成动态规划。定义 dp 三维数组，`dp[i][j][k]` 表示考虑前 i 间房子、最后一件房子颜色为 j、形成 k 个街区的所有方案中最少花费。

怎么做转移？分情况讨论：

1. 第 i 间房子已经上色，即 `houses[i] != 0` 这时候 j 只考虑值为 houses[i] 的区块才有意义，其余情况取正无穷表示。同时根据第 i 间房和第 i-1 间房颜色是否相同，能确定是否新增了一个街区。
2. 第 i 间房子还没有上色，这样要考虑所有能被粉刷的颜色。与上面同样的逻辑确定是否新增了一个街区。

- dp 为什么会是这样三个维度？

> 首先，很容易确定房间编号维度 i 和分区数量维度 k，其次，为了在转移过程中能清楚知道哪些状态转移过来需要增加街区数，哪些不需要，所以需要引入最后一个房间颜色 j 的维度。

- 这样三个维度的规模怎么样？

> i 最大值 m，j 最大值 n， k 最大值 target，根据题目约束，是能接受的，实际上这个问题想不到更好的解法了。

```go
const inf = math.MaxInt32

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			dp[i][j][0] = inf
		}
	}
	for i := 1; i <= m; i++ {
		color := houses[i-1]
		for j := 1; j <= n; j++ {
			for k := 1; k <= target; k++ {
				if k > i {
					dp[i][j][k] = inf
					continue
				}
				if color != 0 { // 第 i 间房子已经上色
					if j != color {
						dp[i][j][k] = inf
						continue
					}
					tmp := inf
					for p := 1; p <= n; p++ {
						if p != j {
							tmp = min(tmp, dp[i-1][p][k-1])
						}
					}
					dp[i][j][k] = min(dp[i-1][j][k], tmp)
				} else { // 第 i 间房子还未上色
					tmp := inf
					for p := 1; p <= n; p++ {
						if p != j {
							tmp = min(tmp, dp[i-1][p][k-1])
						}
					}
					dp[i][j][k] = min(dp[i-1][j][k], tmp) + cost[i-1][j-1]
				}
			}

		}
	}
	res := inf
	for j := 1; j <= n; j++ {
		res = min(res, dp[m][j][target])
	}
	if res == inf {
		return -1
	}
	return res
}
```

复杂度同记忆化搜索方法，值得注意的是空间复杂度可以优化。
