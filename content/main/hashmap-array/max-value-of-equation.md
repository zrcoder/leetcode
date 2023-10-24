---
title: 1499. 满足不等式的最大值
date: 2023-10-24T22:31:55+08:00
---

## [1499. 满足不等式的最大值](https://leetcode.cn/problems/max-value-of-equation) (Hard)

给你一个数组 `points` 和一个整数 `k` 。数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序。也就是说 `points[i] = [xᵢ, yᵢ]` ，并且在 `1 <= i < j <= points.length` 的前提下， `xᵢ < xⱼ` 总成立。

请你找出 `yᵢ + yⱼ + |xᵢ - xⱼ|` 的 **最大值**，其中 `|xᵢ - xⱼ| <= k` 且 `1 <= i < j <= points.length`。

题目测试数据保证至少存在一对能够满足 `|xᵢ - xⱼ| <= k` 的点。

**示例 1：**

```
输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
输出：4
解释：前两个点满足 |xᵢ - xⱼ| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。第三个和第四个点也满足条件，得到值 10 + -10 + |5 - 6| = 1 。
没有其他满足条件的点，所以返回 4 和 1 中最大的那个。
```

**示例 2：**

```
输入：points = [[0,0],[3,0],[9,2]], k = 3
输出：3
解释：只有前两个点满足 |xᵢ - xⱼ| <= 3 ，代入方程后得到值 0 + 0 + |0 - 3| = 3 。

```

**提示：**

- `2 <= points.length <= 10^5`
- `points[i].length == 2`
- `-10^8 <= points[i][0], points[i][1] <= 10^8`
- `0 <= k <= 2 * 10^8`
- 对于所有的 `1 <= i < j <= points.length` ， `points[i][0] < points[j][0]` 都成立。也就是说， `xᵢ` 是严格递增的。

## 分析

### 单调队列

在 j > i 时, xj > xi，则 yi + yj + |xi - xj| = yi-xi + yj+xj

维护一个 双端队列，元素为坐标对，保持队列中 y-x 的值单调递减，这样当一个新的点要加入队列时，可以仅用队头元素更新结果。

注意队列头的元素如果 x 过小，要加入的新点的横坐标减去其值超过了k，则队头的元素要出队。

```go
func findMaxValueOfEquation(points [][]int, k int) int {
	res := math.MinInt32
	q := [][]int{}
	for _, p := range points {
		for len(q) > 0 && p[0]-q[0][0] > k {
			q = q[1:]
		}
		if len(q) > 0 {
			head := q[0]
			cur := head[1] - head[0] + p[0] + p[1]
			if cur > res {
				res = cur
			}
		}
		for len(q) > 0 && q[len(q)-1][1]-q[len(q)-1][0] <= p[1]-p[0] {
			q = q[:len(q)-1]
		}
		q = append(q, []int{p[0], p[1]})
	}
	return res
}

```
