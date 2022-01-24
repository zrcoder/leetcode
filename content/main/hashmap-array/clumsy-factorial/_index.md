---
title: "1006. 笨阶乘"
date: 2021-04-19T22:04:56+08:00
weight: 1
math: true

tags: [模拟, 数学]
---

## [1006. 笨阶乘](https://leetcode-cn.com/problems/clumsy-factorial/)

难度中等

通常，正整数 `n` 的阶乘是所有小于或等于 `n` 的正整数的乘积。例如，`factorial(10) = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1`。

相反，我们设计了一个笨阶乘 `clumsy`：在整数的递减序列中，我们以一个固定顺序的操作符序列来依次替换原有的乘法操作符：乘法(*)，除法(/)，加法(+)和减法(-)。

例如，`clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1`。然而，这些运算仍然使用通常的算术运算顺序：我们在任何加、减步骤之前执行所有的乘法和除法步骤，并且按从左到右处理乘法和除法步骤。

另外，我们使用的除法是地板除法（*floor division*），所以 `10 * 9 / 8` 等于 `11`。这保证结果是一个整数。

实现上面定义的笨函数：给定一个整数 `N`，它返回 `N` 的笨阶乘。

**示例 1：**

```
输入：4
输出：7
解释：7 = 4 * 3 / 2 + 1
```

**示例 2：**

```
输入：10
输出：12
解释：12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1
```

**提示：**

1. `1 <= N <= 10000`
2. `-2^31 <= answer <= 2^31 - 1` （答案保证符合 32 位整数。）



函数签名：

```go
func clumsy(n int) int
```

## 分析

### 使用数组模拟

用一个数组来存储中间计算的值，在加减运算时，把当前数字追加进去，注意减法时追加的数字是负数，在乘除运算时用数组中最后一个数字与当前数字运算后更新数组最后一个元素。

最后计算数组所有元素和即可。

```go
func clumsy(n int) int {
	memo := []int{n}
	n--
	index := 0 // 用于控制乘、除、加、减
	for ; n > 0; n, index = n-1, (index+1)%4 {
		switch index {
		case 0:
			memo[len(memo)-1] *= n
		case 1:
			memo[len(memo)-1] /= n
		case 2:
			memo = append(memo, n)
		default:
			memo = append(memo, -n)
		}
	}
	res := 0
	for _, v := range memo {
		res += v
	}
	return res
}
```

时空复杂度都是 `O(n)`，memo 实际上需要大概 `n/2` 的空间。

### 优化模拟

```go
func clumsy(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 6
	}
	res := n*(n-1)/(n-2) + n - 3
	n -= 4
	for ; n >= 4; n -= 4 {
		res = res - n*(n-1)/(n-2) + n - 3
	}
	if n == 3 {
		return res - 3*2
	}
	return res - n
}
```

时间复杂度 `O(n/4)`，空间复杂度`O(1)`。

### 梳理规律

$$
clumsy(n) = \frac{n(n-1)}{n-2} + (n-3) - \frac{(n-4)(n-5)}{n-6}+(n-7)-\cdots
$$

对于里边的分式，可以做一变形：

$$
\frac{n(n-1)}{n-2} = \frac{n^2-2n+n}{n-2}= n + \frac{n}{n-2}= n + \frac{n-2+2}{n-2}= n+1 + \frac{2}{n-2}
$$

当 `n > 2` 时，$\frac{2}{n-2} = 0$, 所以：$\frac{n(n-1)}{n-2} =  n+1 + \frac{2}{n-2} = n+1$

根据这一发现，可以看到 (1) 中有许多项是可以消去的，比如 $(n-3) - \frac{(n-4)(n-5)}{n-6} = 0$

现在要考虑  (1) 式最后几项的情况，这可以由 n 对 4 的余数来分类：

`n%4 == 0`: $clumsy(n) = \frac{n(n-1)}{n-2} + \cdots+5 -4\times3\div2 + 1 = n+1 $

`n%4 == 1`: $clumsy(n) = \frac{n(n-1)}{n-2} + \cdots +2-1 = n+1+2-1=n+2$

`n%4 == 2`: $clumsy(n) = \frac{n(n-1)}{n-2} + \cdots +3-2\times1=n+1+1=n+2$

`n%4 == 3`: $clumsy(n) = \frac{n(n-1)}{n-2} + \cdots +4-3\times2\div1=n+1+4-6=n-1$

在 `n` 较大时直接采用上边的公式，`n >= 4` 时单独计算即可。

```go
func clumsy(n int) int {
	switch {
	case n < 3:
		return n
	case n == 3:
		return 6
	case n == 4:
		return 7
	case n%4 == 0:
		return n + 1
	case n%4 == 3:
		return n - 1
	default:
		return n + 2
	}
}
```