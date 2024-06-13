---
title: "1263. 推箱子"
date: 2023-05-08T15:29:51+08:00
math: true
---

#### [1263. 推箱子](https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/)

难度困难

「推箱子」是一款风靡全球的益智小游戏，玩家需要将箱子推到仓库中的目标位置。

游戏地图用大小为 `m x n` 的网格 `grid` 表示，其中每个元素可以是墙、地板或者是箱子。

现在你将作为玩家参与游戏，按规则将箱子 `'B'` 移动到目标位置 `'T'` ：

- 玩家用字符 `'S'` 表示，只要他在地板上，就可以在网格中向上、下、左、右四个方向移动。
- 地板用字符 `'.'` 表示，意味着可以自由行走。
- 墙用字符 `'#'` 表示，意味着障碍物，不能通行。
- 箱子仅有一个，用字符 `'B'` 表示。相应地，网格上有一个目标位置 `'T'`。
- 玩家需要站在箱子旁边，然后沿着箱子的方向进行移动，此时箱子会被移动到相邻的地板单元格。记作一次「推动」。
- 玩家无法越过箱子。

返回将箱子推到目标位置的最小 **推动** 次数，如果无法做到，请返回 `-1`。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/11/16/sample_1_1620.png)**

**输入：** grid = [["#","#","#","#","#","#"],
             ["#","T","#","#","#","#"],
             ["#",".",".","B",".","#"],
             ["#",".","#","#",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
**输出：** 3
**解释：** 我们只需要返回推箱子的次数。

**示例 2：**

**输入：** grid = [["#","#","#","#","#","#"],
             ["#","T","#","#","#","#"],
             ["#",".",".","B",".","#"],
             ["#","#","#","#",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
**输出：** -1

**示例 3：**

**输入：** grid = [["#","#","#","#","#","#"],
             ["#","T",".",".","#","#"],
             ["#",".","#","B",".","#"],
             ["#",".",".",".",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
**输出：** 5
**解释：** 向下、向左、向左、向上再向上。

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 20`
- `grid` 仅包含字符 `'.'`, `'#'`,  `'S'` , `'T'`, 以及 `'B'`。
- `grid` 中 `'S'`, `'B'` 和 `'T'` 各只能出现一个。

函数签名：

```go
func minPushBox(grid [][]byte) int
```

## 分析

如果不考虑箱子，求玩家达 T 的最短路径，将会变得非常简单，用常规 BFS 做就行了， 即用一个队列维护玩家位置，每次考虑出队的元素，如果是 T 所在的位置则得到结果，直接返回 bfs 的层数，否则，让玩家向上下左右任意方向移动一步，得到下一个位置，如果没有超出 grid 范围也不是墙就入队。

现在要考虑箱子，在玩家移动的时候，可能影响到箱子使得箱子移动（比如玩家在箱子右侧，玩家向左移动导致箱子也向左移动）。
这样 BFS 的方法依然可用，不过状态从仅考虑玩家的位置（一维）到同时考虑玩家和箱子的位置（二维）。

```go
const maxSize = 20
var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	//  先找到玩家和箱子的初始位置
	var sr, sc, br, bc int
	for r, row := range grid {
		for c, v := range row {
			if v == 'S' {
				sr, sc = r, c
			} else if v == 'B' {
				br, bc = r, c
			}
		}
	}
	vis := [maxSize][maxSize][maxSize][maxSize]bool{}
	queue := [][4]int{{sr, sc, br, bc}}
	vis[sr][sc][br][bc] = true
	isValid := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != '#'
	}
	for steps := 0; len(queue) > 0; steps++ {
		// 一直移动玩家知道推动了箱子
		var pushedQueue [][4]int
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			sr, sc, br, bc = cur[0], cur[1], cur[2], cur[3]
			if grid[br][bc] == 'T' { // 箱子所在的位置就是目的地，直接返回 BFS 层数即路径长度
				return steps
			}
			// 玩家尝试走一步
			for _, d := range dirs {
				nr, nc := sr+d[0], sc+d[1]
				if !isValid(nr, nc) {
					continue
				}
				if nr == br && nc == bc { // 玩家与箱子重合，箱子需要向相同的方向移动
					nbr, nbc := br+d[0], bc+d[1]
					if isValid(nbr, nbc) && !vis[nr][nc][nbr][nbc] {
						vis[nr][nc][nbr][nbc] = true
						// 新状态入队到 pushedQueue 里
						pushedQueue = append(pushedQueue, [4]int{nr, nc, nbr, nbc})
					}
				} else {
					if !vis[nr][nc][br][bc] {
						vis[nr][nc][br][bc] = true
						// 新状态入队到 queue 里
						queue = append(queue, [4]int{nr, nc, br, bc})
					}
				}
			}
		}
		// 此时 queue 空了
		queue = pushedQueue
	}
	return -1
}
```

时空复杂度均为 $O(m^2n^2)$ .
