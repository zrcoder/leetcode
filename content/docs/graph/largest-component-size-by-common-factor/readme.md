---
title: "952. 按公因数计算最大组件大小"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [952. 按公因数计算最大组件大小](https://leetcode-cn.com/problems/largest-component-size-by-common-factor/)

难度困难

给定一个由不同正整数的组成的非空数组 `A`，考虑下面的图：

- 有 `A.length` 个节点，按从 `A[0]` 到 `A[A.length - 1]` 标记；
- 只有当 `A[i]` 和 `A[j]` 共用一个大于 1 的公因数时，`A[i]` 和 `A[j]` 之间才有一条边。

返回图中最大连通组件的大小。

 



**示例 1：**

```
输入：[4,6,15,35]
输出：4
```

**示例 2：**

```
输入：[20,50,9,63]
输出：2
```

**示例 3：**

```
输入：[2,3,6,7,4,12,21,39]
输出：8
```

 

**提示：**

1. `1 <= A.length <= 20000`
2. `1 <= A[i] <= 100000`

函数签名：

```go
func largestComponentSize(A []int) int
```

## 分析

用并查集将极大简化问题。一个自然的思路是两重循环让数组里的元素两两相遇，再求下最大公约数 gcd，如果 gcd 大于 1，就把两个数字在并查集合并，最终统计并查集里每个连通分量的大小就行。

不过这样做的复杂度较大，假设数组长度为 n，最大数字为 m， 并查集操作看成常数级别，那么复杂度将是 `O(n^2*logm)`。

实际测试也是超时，没有通过所有用例，需要再找突破。

根据问题的限制，数组长度的平方远远大于元素的最大值。可以避免元素两两相遇的做法。

只遍历一次数组，对每一个元素分解质因数。之后将所有因数和元素本身合并，这样最终 gcd 大于 1 的数字会合并，且所有的因数和元素自己合并并不会导致联通分量增多，最终并查集里的分量数是准确的。

综上，需要从元素值的角度重新考虑。

```go
var uf []int

func find(x int) int {
	for uf[x] != x {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}

func union(x, y int) {
	uf[find(x)] = find(y)
}

func largestComponentSize(A []int) int {
	max := 0
	for i := 0; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
	}
	uf = make([]int, max+1)
	for i := 0; i < max; i++ {
		uf[i] = i
	}
	for _, num := range A {
		for j := 2; j*j <= num; j++ {
			if num%j == 0 {
				union(num, j)
				union(num, num/j)
			}
		}
	}
	res := 0
	cnt := make([]int, max+1)
	for _, v := range A {
		root := find(v)
		cnt[root]++
		if res < cnt[root] {
			res = cnt[root]
		}
	}
	return res
}
```
分解质因数的复杂度是 sqrt(m)。

时间复杂度 O(n*sqrt(m)), 空间复杂度 O(m)。