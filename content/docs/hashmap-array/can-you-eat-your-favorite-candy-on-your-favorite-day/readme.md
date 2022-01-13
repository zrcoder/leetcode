---
title: "5667. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [前缀和]
---

## [5667. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？](https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/)

难度中等

给你一个下标从 **0** 开始的正整数数组 `candiesCount` ，其中 `candiesCount[i]` 表示你拥有的第 `i` 类糖果的数目。同时给你一个二维数组 `queries` ，其中 `queries[i] = [favoriteTypei, favoriteDayi, dailyCapi]` 。

你按照如下规则进行一场游戏：

- 你从第 `0` 天开始吃糖果。
- 你在吃完 **所有** 第 `i - 1` 类糖果之前，**不能** 吃任何一颗第 `i` 类糖果。
- 在吃完所有糖果之前，你必须每天 **至少** 吃 **一颗** 糖果。

请你构建一个布尔型数组 `answer` ，满足 `answer.length == queries.length` 。`answer[i]` 为 `true` 的条件是：在每天吃 **不超过** `dailyCapi` 颗糖果的前提下，你可以在第 `favoriteDayi` 天吃到第 `favoriteTypei` 类糖果；否则 `answer[i]` 为 `false` 。注意，只要满足上面 3 条规则中的第二条规则，你就可以在同一天吃不同类型的糖果。

请你返回得到的数组 `answer` 。

 

**示例 1：**

```
输入：candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
输出：[true,false,true]
提示：
1- 在第 0 天吃 2 颗糖果(类型 0），第 1 天吃 2 颗糖果（类型 0），第 2 天你可以吃到类型 0 的糖果。
2- 每天你最多吃 4 颗糖果。即使第 0 天吃 4 颗糖果（类型 0），第 1 天吃 4 颗糖果（类型 0 和类型 1），你也没办法在第 2 天吃到类型 4 的糖果。换言之，你没法在每天吃 4 颗糖果的限制下在第 2 天吃到第 4 类糖果。
3- 如果你每天吃 1 颗糖果，你可以在第 13 天吃到类型 2 的糖果。
```

**示例 2：**

```
输入：candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
输出：[false,true,true,false,false]
```

 

**提示：**

- `1 <= candiesCount.length <= 105`
- `1 <= candiesCount[i] <= 105`
- `1 <= queries.length <= 105`
- `queries[i].length == 3`
- `0 <= favoriteTypei < candiesCount.length`
- `0 <= favoriteDayi <= 109`
- `1 <= dailyCapi <= 109`

## 分析

题意弄清楚后就比较简单了。

对于某个查询，第 favoriteDayi 天（从0开始）能否吃到 favoriteTypei 类型的糖果？

每天一颗，最少要吃 favoriteDayi+1 颗糖果；每天 dailyCapi 颗，最多能吃 dailyCapi * (favoriteDayi+1) 颗糖果。

而对于 favoriteTypei  类型的糖果，要吃到它必须先把它前面类型的都吃了。可以用一个前缀和数组求出每类糖果要吃到，必须先吃多少其他类型糖果。
> 天数从 0 开始，第 day 天意味着吃了 day+1 天糖果
> 
> 尽量避免用除法，防止向下取整丢失精度，改成乘法

```go
func canEat(candiesCount []int, queries [][]int) []bool {
	// prefixSum[i] 代表要吃到糖果 i， 先要吃多少颗其他类型的糖果
	prefixSum := make([]int, len(candiesCount)+1)
	for i, v := range candiesCount {
		prefixSum[i+1] = prefixSum[i] + v
	}
	answer := make([]bool, len(queries))
	for i, q := range queries {
		candy, day, limit := q[0], q[1], q[2]
		min, max := day+1, (day+1)*limit
		answer[i] = min <= prefixSum[candy+1] && max > prefixSum[candy]
	}
	return answer
}
```

时空复杂度都是`O(n)`。

> 每个查询互不干扰，前边的查询仅是查询，并不会真的吃掉糖果，这样也就不影响后边的查询。
> 
> 如果每次查询都真的吃掉一些糖果，问题改为在最好的日子吃最爱的糖，最多能吃多少，就会复杂很多，要考虑动态规划了。