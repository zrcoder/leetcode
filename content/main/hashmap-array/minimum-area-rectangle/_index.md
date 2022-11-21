---
title: "939. 最小面积矩形"
date: 2022-11-21T17:19:57+08:00
math: true
---

#### [939. 最小面积矩形](https://leetcode.cn/problems/minimum-area-rectangle/)

难度中等

给定在 xy 平面上的一组点，确定由这些点组成的矩形的最小面积，其中矩形的边平行于 x 轴和 y 轴。

如果没有任何矩形，就返回 0。

**示例 1：**

**输入：**[[1,1],[1,3],[3,1],[3,3],[2,2]]
**输出**4

**示例 2：**

**输入：**[[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
**输出**2

**提示：**

1. `1 <= points.length <= 500`
2. `0 <= points[i][0] <= 40000`
3. `0 <= points[i][1] <= 40000`
4. 所有的点都是不同的。

函数签名：

```go
func minAreaRect(points [][]int) int
```

## 分析

可以用两层循环枚举任意两个点，如果这两个点的连线不平行于坐标轴，那么可以将这两个点看作一个矩形的对角线端点，然后判断矩形其他两个顶点（它们的坐标可以很容易从已有两个点得出）是否在points中，如果在，找到了一个矩形，计算其面积并更新结果即可。



为了迅速判断一个点是否在输入点数组里，可以用 set。

> 1. 用一个 40001*40001 的 bool 数组，这样比较浪费空间
> 
> 2. 用哈希表，需要将二维的点映射为一维的数字，如 (x, y), 可以映射为 x*40001+y



```go
func minAreaRect(points [][]int) int {
	set := map[int]bool{}
	for _, v := range points {
		set[hash(v)] = true
	}

	n := len(points)
	res := math.MaxInt32

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a, b := points[i], points[j]
			if isSameLine(a, b) {
				continue
			}
			x, y := []int{a[0], b[1]}, []int{b[0], a[1]}
			if !set[hash(x)] || !set[hash(y)] {
				continue
			}
			res = min(res, abs(a[0]-b[0])*abs(a[1]-b[1]))
		}
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func hash(v []int) int {
	return v[0]*40001 + v[1]
}

func isSameLine(a, b []int) bool {
	return a[0] == b[0] || a[1] == b[1]
}
```

时间复杂度 $O(n^2)$， 空间复杂度 $O(n)$ 。