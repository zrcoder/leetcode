package solution

/*
## [1139. Largest 1-Bordered Square](https://leetcode.cn/problems/largest-1-bordered-square) (Medium)

给你一个由若干 `0` 和 `1` 组成的二维网格 `grid`，请你找出边界全部由 `1` 组成的最大 **正方形** 子网格，并返回该子网格中的元素数量。如果不存在，则返回 `0`。

**示例 1：**

```
输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
输出：9

```

**示例 2：**

```
输入：grid = [[1,1,0,0]]
输出：1

```

**提示：**

- `1 <= grid.length <= 100`
- `1 <= grid[0].length <= 100`
- `grid[i][j]` 为 `0` 或 `1`


*/

// [start] don't modify

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	up := make([][]int, m+1)
	left := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		up[i] = make([]int, n+1)
		left[i] = make([]int, n+1)
	}
	tmp := 0
	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			if grid[r-1][c-1] == 0 {
				continue
			}
			up[r][c] = up[r-1][c] + 1
			left[r][c] = left[r][c-1] + 1
			cur := min(up[r][c], left[r][c])
			for left[r-cur+1][c] < cur || up[r][c-cur+1] < cur {
				cur--
			}
			if cur > tmp {
				tmp = cur
			}
		}
	}

	return tmp * tmp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// [end] don't modify
