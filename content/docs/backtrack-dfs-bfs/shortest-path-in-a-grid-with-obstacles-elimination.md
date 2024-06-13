---
title: "1293. 网格中的最短路径"
date: 2022-12-18T11:21:48+08:00
---

# [1293. 网格中的最短路径](https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/description)

| Category   | Difficulty    | Likes | Dislikes |
| ------------ | --------------- | ------- | ---------- |
| algorithms | Hard (37.97%) | 214   | -        |

给你一个 `m * n` 的网格，其中每个单元格不是 `0`（空）就是 `1`（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。

如果您 **最多** 可以消除 `k` 个障碍物，请找出从左上角 `(0, 0)` 到右下角 `(m-1, n-1)` 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 `-1` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/09/30/short1-grid.jpg)

```
输入： grid = [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]], k = 1
输出：6
解释：
不消除任何障碍的最短路径是 10。
消除位置 (3,2) 处的障碍后，最短路径是 6 。该路径是 (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/09/30/short2-grid.jpg)

```
输入：grid = [[0,1,1],[1,1,1],[1,0,0]], k = 1
输出：-1
解释：我们至少需要消除两个障碍才能找到这样的路径。
```

**提示：**

* `grid.length == m`
* `grid[0].length == n`
* `1 <= m, n <= 40`
* `1 <= k <= m*n`
* `grid[i][j]` 是 `0` 或**** `1`
* `grid[0][0] == grid[m-1][n-1] == 0`

函数签名：

```go
func shortestPath(grid [][]int, k int) int
```

## 分析

### BFS

经典BFS，只是BFS过程记录的节点需要加上k的信息。

```go
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k >= m+n-3 { // 优化，可以直接向右走到头再向下走到终点的情况
		return m + n - 2
	}
	type Node struct{ r, c, k int } // 除记录当前位置，也记录剩余可以消除障碍的数量
	node := Node{0, 0, k}
	q := []Node{node}
	seen := map[Node]bool{node: true}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	inRange := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n
	}
	for step := 0; len(q) > 0; step++ {
		for size := len(q); size > 0; size-- {
			cur := q[0]
			q = q[1:]
			if cur.r == m-1 && cur.c == n-1 {
				return step
			}
			for _, d := range dirs {
				r, c := cur.r+d[0], cur.c+d[1]
				if !inRange(r, c) || cur.k == 0 && grid[r][c] == 1 {
					continue
				}
				node := Node{r, c, cur.k}
				if grid[r][c] == 1 {
					node.k--
				}
				if seen[node] {
					continue
				}
				seen[node] = true
				q = append(q, node)
			}
		}
	}
	return -1
}
```

时空复杂度均为：`O(m*n*min(k, m+n))`。
