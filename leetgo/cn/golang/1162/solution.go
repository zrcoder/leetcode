package solution

/*
## [1162. As Far from Land as Possible](https://leetcode.cn/problems/as-far-from-land-as-possible) (Medium)

你现在手里有一份大小为 `n x n` 的 网格 `grid`，上面的每个 单元格 都用 `0` 和 `1` 标记好了。其中 `0` 代表海洋， `1` 代表陆地。

请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 `-1`。

我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）： `(x0, y0)` 和 `(x1, y1)` 这两个单元格之间的距离是 `|x0 - x1| + |y0 - y1|` 。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/08/17/1336_ex1.jpeg)**

```
输入：grid = [[1,0,1],[0,0,0],[1,0,1]]
输出：2
解释：
海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。

```

**示例 2：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/08/17/1336_ex2.jpeg)**

```
输入：grid = [[1,0,0],[0,0,0],[0,0,0]]
输出：4
解释：
海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。

```

**提示：**

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 100`
- `grid[i][j]` 不是 `0` 就是 `1`


*/

// [start] don't modify
func maxDistance(grid [][]int) int {
	n := len(grid)
	var queue [][]int
	for r, row := range grid {
		for c, v := range row {
			if v == 1 {
				queue = append(queue, []int{r, c})
			}
		}
	}
	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}

	var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	inrange := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n
	}
	res := -1
	for ; len(queue) > 0; res++ {
		tmp := queue
		queue = nil
		for _, v := range tmp {
			for _, d := range dirs {
				r, c := v[0]+d[0], v[1]+d[1]
				if inrange(r, c) && grid[r][c] == 0 {
					grid[r][c] = 1
					queue = append(queue, []int{r, c})
				}
			}
		}
	}
	return res
}
// [end] don't modify
