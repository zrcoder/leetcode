---
title: 2544. 交替数字和
date: 2023-07-12T17:08:33+08:00
---

## [2544. 交替数字和](https://leetcode.cn/problems/alternating-digit-sum) (Easy)

给你一个正整数 `n` 。 `n` 中的每一位数字都会按下述规则分配一个符号：

- **最高有效位** 上的数字分配到 **正** 号。
- 剩余每位上数字的符号都与其相邻数字相反。

返回所有数字及其对应符号的和。

**示例 1：**

```
输入：n = 521
输出：4
解释：(+5) + (-2) + (+1) = 4
```

**示例 2：**

```
输入：n = 111
输出：1
解释：(+1) + (-1) + (+1) = 1

```

**示例 3：**

```
输入：n = 886996
输出：0
解释：(+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0

```

**提示：**

- `1 <= n <= 10⁹`

## 分析

常规解法：
```go
func alternateDigitSum(n int) int {
	res := 0
	cnt := 0
	sign := 1
	for n > 0 {
		r := n % 10
		res += sign * r
		sign = -sign
		n /= 10
		cnt++
	}
	if cnt%2 == 0 {
		return -res
	}
	return res
}
```
一个更简洁的解法：

```go
func alternateDigitSum(n int) int {
	res := 0
	for n > 0 {
		res = n%10 - res
		n /= 10
	}
	return res
}

```
