---
title: "零钱兑换"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [动态规划, 背包问题]
---

##  [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change)

给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

```
示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1

说明:
你可以认为每种硬币的数量是无限的。
```

函数签名：
```
func coinChange(coins []int, amount int) int
```
## 分析
不求 max 而求 min 的完全背包问题
```go
func coinChange(coins []int, amount int) int {
	infinityAmount := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = infinityAmount
	}
	dp[0] = 0
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-v]+1)
		}
	}
	if dp[amount] == infinityAmount {
		return -1
	}
	return dp[amount]
}
```

## [518. 零钱兑换 II](https://leetcode-cn.com/problems/coin-change-2)
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
```
示例 1:
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2:
输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。

示例 3:
输入: amount = 10, coins = [10]
输出: 1

注意:
你可以假设：

0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数
```

## 分析
完全背包问题。

```go
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] += dp[j-v]
		}
	}
	return dp[amount]
}
```

## 扩展
如果要求排列数而不是组合数，应该怎么做？比如这个问题：[377. 组合总和 Ⅳ](https://leetcode-cn.com/problems/combination-sum-iv/)

> 需要把两层循环的顺序颠倒，想想是为什么。

参考实现：
```go
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for sum := 1; sum <= target; sum++ {
		for _, v := range nums {
			if v <= sum {
				dp[sum] += dp[sum-v]
			}
		}
	}
	return dp[target]
}
```