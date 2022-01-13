package solution

/*
## [1210. Minimum Moves to Reach Target with Rotations](https://leetcode.cn/problems/minimum-moves-to-reach-target-with-rotations) (Hard)

你还记得那条风靡全球的贪吃蛇吗？

我们在一个 `n*n` 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（ `(0, 0)` 和 `(0, 1)`）开始移动。我们用 `0` 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（ `(n-1, n-2)` 和 `(n-1, n-1)`）。

每次移动，蛇可以这样走：

- 如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
- 如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。
- 如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（ `(r, c)`、 `(r, c+1)`）移动到 （ `(r, c)`、 `(r+1, c)`）。

  ![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/28/image-2.png)
- 如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（ `(r, c)`、 `(r+1, c)`）移动到（ `(r, c)`、 `(r, c+1)`）。

  ![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/28/image-1.png)

返回蛇抵达目的地所需的最少移动次数。

如果无法到达目的地，请返回 `-1`。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/28/image.png)**

```
输入：grid = [[0,0,0,0,0,1],
               [1,1,0,0,1,0],
               [0,0,0,0,1,1],
               [0,0,1,0,1,0],
               [0,1,1,0,0,0],
               [0,1,1,0,0,0]]
输出：11
解释：
一种可能的解决方案是 [右, 右, 顺时针旋转, 右, 下, 下, 下, 下, 逆时针旋转, 右, 下]。

```

**示例 2：**

```
输入：grid = [[0,0,1,1,1,1],
               [0,0,0,0,1,1],
               [1,1,0,0,0,1],
               [1,1,1,0,0,1],
               [1,1,1,0,0,1],
               [1,1,1,0,0,0]]
输出：9

```

**提示：**

- `2 <= n <= 100`
- `0 <= grid[i][j] <= 1`
- 蛇保证从空单元格开始出发。


*/

// [start] don't modify
type State struct {
	tailX, tailY, dir int
}

var dirs = []State{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func minimumMoves(grid [][]int) int {
	n := len(grid)
	vis := make(map[State]bool, 0)
	s := State{0, 0, 0}
	vis[s] = true
	q := []State{s}
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, t := range tmp {
			if t.tailX == n-1 && t.tailY == n-2 {
				return step
			}
			for _, d := range dirs {
				s := State{t.tailX + d.tailX, t.tailY + d.tailY, t.dir ^ d.dir}
				headX, headY := s.tailX+s.dir, s.tailY+(s.dir^1)
				if s.tailX < n && s.tailY < n && headX < n && headY < n && !vis[s] &&
					grid[headX][headY] == 0 && grid[s.tailX][s.tailY] == 0 && (d.dir == 0 || grid[s.tailX+1][s.tailY+1] == 0) {
					vis[s] = true
					q = append(q, s)
				}
			}
		}
	}
	return -1
}
// [end] don't modify
