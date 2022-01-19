---
title: "970. 强整数"
date: 2021-04-30T17:52:25+08:00
weight: 50
math: true
tags: [模拟]
---

## [970. 强整数](https://leetcode-cn.com/problems/powerful-integers/)

难度中等

给定两个正整数 `x` 和 `y`，如果某一整数等于 `x^i + y^j`，其中整数 `i >= 0` 且 `j >= 0`，那么我们认为该整数是一个*强整数*。

返回值小于或等于 `bound` 的所有*强整数*组成的列表。

你可以按任何顺序返回答案。在你的回答中，每个值最多出现一次。

**示例 1：**

```
输入：x = 2, y = 3, bound = 10
输出：[2,3,4,5,7,9,10]
解释： 
2 = 2^0 + 3^0
3 = 2^1 + 3^0
4 = 2^0 + 3^1
5 = 2^1 + 3^1
7 = 2^2 + 3^1
9 = 2^3 + 3^0
10 = 2^0 + 3^2
```

**示例 2：**

```
输入：x = 3, y = 5, bound = 15
输出：[2,4,6,8,10,14]
```

**提示：**

- `1 <= x <= 100`
- `1 <= y <= 100`
- `0 <= bound <= 10^6`

函数签名：

```go
func powerfulIntegers(x, y, bound int) []int
```

## 分析

这个问题很独特，独特在没有比朴素解法更好的解法~

朴素解法很容易想到，先分析下复杂度怎么样。

假设 `x`、`y` 都比 1 大，要满足 $x^i+y^j \le bound$ ，显然 `i` 和 `j` 有上限，分别是 $log_{x}bound$ 和 $log_{y}bound$ 。

根据题目约束，这两个值不会大于 20，这样就可以用朴素解法。

```go
func powerfulIntegers(x, y, bound int) []int {
	px, py := 0, 0 // x 和 y 的指数上限。如果 x 是 1，px 就是 0，y 和 py 同理
	if x > 1 {
		px = int(math.Log2(float64(bound)) / math.Log2(float64(x))) // log(x, bound)
	}
	if y > 1 {
		py = int(math.Log2(float64(bound)) / math.Log2(float64(y))) // log(y, bound)
	}
	set := map[int]bool{} // 用集合去重
	powX := 1             // x^i
	for i := 0; i <= px; i++ {
		powY := 1 // y^j
		for j := 0; j <= py; j++ {
			val := powX + powY
			if val > bound {
				break
			}
			set[val] = true
			powY *= y
		}
		powX *= x
	}
	res := make([]int, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}
```

时间复杂度是 `O(XY)`， `X` 和 `Y` 就是两个指数上限值，各自不会超过 20 。