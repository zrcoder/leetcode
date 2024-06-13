---
title: "593. 有效的正方形"
date: 2022-12-16T18:03:50+08:00
---
## [593. 有效的正方形](https://leetcode.cn/problems/valid-square/)

**难度** **中等**

给定2D空间中四个点的坐标 `p1`, `p2`, `p3` 和 `p4`，如果这四个点构成一个正方形，则返回 `true` 。

点的坐标 `p<sub>i</sub>` 表示为 `[xi, yi]` 。 `输入没有任何顺序` 。

一个 **有效的正方形** 有四条等边和四个等角(90度角)。

**示例 1:**

<pre><strong>输入:</strong> p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
<strong>输出:</strong> True
</pre>

**示例 2:**

<pre><strong>输入：</strong>p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
<b>输出：</b>false
</pre>

**示例 3:**

<pre><b>输入：</b>p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
<b>输出：</b>true
</pre>

**提示:**

* `p1.length == p2.length == p3.length == p4.length == 2`
* `-10<sup>4</sup> <= x<sub>i</sub>, y<sub>i</sub> <= 10<sup>4</sup>`

函数签名：

```go
func validSquare(p1, p2, p3, p4 []int) bool
```

## 分析

可以选定两条对角线来判断：

1. 中点相同，说明是平行四边形
2. 长度相同，说明是矩形
3. 互相垂直，说明是正方形

对角线的选定需要对给定的四个点做排列组合，这看起来有24种情况，实际只需要三种。

```go
func validSquare(p1, p2, p3, p4 []int) bool {
	return check(p1, p2, p3, p4) ||
		check(p1, p3, p2, p4) ||
		check(p1, p4, p2, p3)
}

func check(p1, p2, p3, p4 []int) bool {
	if !checkMid(p1, p2, p3, p4) {
		return false
	}
	v1 := []int{p1[0] - p2[0], p1[1] - p2[1]}
	v2 := []int{p3[0] - p4[0], p3[1] - p4[1]}
	return checkLength(v1, v2) && isVertical(v1, v2)
}

func checkMid(p1, p2, p3, p4 []int) bool {
	return p1[0]+p2[0] == p3[0]+p4[0] && p1[1]+p2[1] == p3[1]+p4[1]
}

func checkLength(v1, v2 []int) bool {
	l1, l2 := length(v1), length(v2)
	return l1 != 0 && l1 == l2
}

func isVertical(v1, v2 []int) bool {
	return v1[0]*v2[0]+v1[1]*v2[1] == 0
}

func length(x []int) int {
	return x[0]*x[0] + x[1]*x[1]
}
```

时空复杂度都是`O(1)`。
