// Created by Bob at 2023/06/02 14:40
// https://leetcode.cn/problems/detonate-the-maximum-bombs/

/*
2101. 引爆最多的炸弹 (Medium)
给你一个炸弹列表。一个炸弹的 **爆炸范围** 定义为以炸弹为圆心的一个圆。

炸弹用一个下标从 **0** 开始的二维整数数组 `bombs` 表示，其中 `bombs[i] = [xᵢ, yᵢ,
rᵢ]` 。 `xᵢ` 和 `yᵢ` 表示第 `i` 个炸弹的 X 和 Y 坐标， `rᵢ` 表示爆炸范围的
**半径** 。

你需要选择引爆 **一个** 炸弹。当这个炸弹被引爆时， **所有**
在它爆炸范围内的炸弹都会被引爆，这些炸弹会进一步将它们爆炸范围内的其他炸弹引爆。

给你数组 `bombs` ，请你返回在引爆 **一个** 炸弹的前提下， **最多** 能引爆的炸弹数目。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-3.png)

```
输入：bombs = [[2,1,3],[6,1,4]]
输出：2
解释：
上图展示了 2 个炸弹的位置和爆炸范围。
如果我们引爆左边的炸弹，右边的炸弹不会被影响。
但如果我们引爆右边的炸弹，两个炸弹都会爆炸。
所以最多能引爆的炸弹数目是 max(1, 2) = 2 。

```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/11/06/desmos-eg-2.png)

```
输入：bombs = [[1,1,5],[10,10,5]]
输出：1
解释：
引爆任意一个炸弹都不会引爆另一个炸弹。所以最多能引爆的炸弹数目为 1 。

```

**示例 3：**

![](https://assets.leetcode.com/uploads/2021/11/07/desmos-eg1.png)

```
输入：bombs = [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]
输出：5
解释：
最佳引爆炸弹为炸弹 0 ，因为：
- 炸弹 0 引爆炸弹 1 和 2 。红色圆表示炸弹 0 的爆炸范围。
- 炸弹 2 引爆炸弹 3 。蓝色圆表示炸弹 2 的爆炸范围。
- 炸弹 3 引爆炸弹 4 。绿色圆表示炸弹 3 的爆炸范围。
所以总共有 5 个炸弹被引爆。

```

**提示：**

- `1 <= bombs.length <= 100`
- `bombs[i].length == 3`
- `1 <= xᵢ, yᵢ, rᵢ <= 10⁵`
*/

package main

// @lc code=begin

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	next := make([][]int, n)
	for j := range bombs {
		for i := range bombs[:j] {
			a, b := bombs[i], bombs[j]
			dx := a[0] - b[0]
			dy := a[1] - b[1]
			ds := dx*dx + dy*dy
			if ds <= a[2]*a[2] {
				next[i] = append(next[i], j)
			}
			if ds <= b[2]*b[2] {
				next[j] = append(next[j], i)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		// cur := cal(i, next)
		cur := calDfs(i, next)
		if cur > res {
			res = cur
		}
	}
	return res
}

// calculate with bfs
func cal(x int, next [][]int) int {
	vis := make([]bool, len(next))
	vis[x] = true
	q := []int{x}
	cnt := 1
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range next[x] {
			if vis[y] {
				continue
			}
			vis[y] = true
			cnt++
			q = append(q, y)
		}
	}
	return cnt
}

// caculate with dfs
func calDfs(x int, next [][]int) int {
	cnt := 1
	vis := make([]bool, len(next))
	vis[x] = true
	var dfs func(x int)
	dfs = func(x int) {
		for _, y := range next[x] {
			if vis[y] {
				continue
			}
			vis[y] = true
			cnt++
			dfs(y)
		}
	}
	dfs(x)
	return cnt
}

// @lc code=end
