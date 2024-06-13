---
title: "不同路径"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [62. 不同路径](https://leetcode-cn.com/problems/unique-paths)

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。  
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。  
问总共有多少条不同的路径？  
例如，上图是一个7 x 3 的网格。有多少可能的路径？  

```
示例 1:
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28

提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 10 ^ 9
```

## 分析

### 常规动态规划  

dp[r][c]代表到达位置（r，c）共有几种方法，  
显然dp[0][...]和dp[...][0]都是1  
其他位置是到达上面一格的方法数+到达左面一格的方法数  
dp数组可以优化为一维。  

```go
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || c == 0 {
				dp[c] = 1
			} else {
				dp[c] += dp[c-1]
			}
		}
	}
	return dp[n-1]
}
```

时间复杂度 O (m*n), 空间复杂度 O(n)

### 排列组合  

无论怎么走，最终需要向右走n-1步，向下走m-1步，所以答案是C(m+n-2, n-1) ， 其中第二个参数也可以是 m-1。

```go
func uniquePaths(m int, n int) int {
	total := m+n-2
	min := m-1
	if n < m {
		min = n-1
	}
	// 求 C(total, min)
    // 可根据 C(x, y) = C(x, y-1) * (x-y+1) / y 递推得到结果
	res := 1
	for i := 1; i <= min; i++ {
		res = res*(total-i+1) / i
	}
	return res
}
```

时间复杂度 O（min), 空间复杂度 O(1)

## [63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii)

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？  
网格中的障碍物和空位置分别用 1 和 0 来表示。  
说明：m 和 n 的值均不超过 100。  

## 分析

同上一个问题，动态规划。  

```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 初始化dp第0行
	for c := 1; c < n && obstacleGrid[0][c] == 0; c++ {
		dp[0][c] = 1
	}
	// 初始化dp第0列
	for r := 1; r < m && obstacleGrid[r][0] == 0; r++ {
		dp[r][0] = 1
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if obstacleGrid[r][c] == 0 {
				dp[r][c] = dp[r-1][c] + dp[r][c-1]
			}
		}
	}
	return dp[m-1][n-1]
}
```

## [980. 不同路径 III](https://leetcode-cn.com/problems/unique-paths-iii/)

难度困难

在二维网格 `grid` 上，有 4 种类型的方格：

- `1` 表示起始方格。且只有一个起始方格。
- `2` 表示结束方格，且只有一个结束方格。
- `0` 表示我们可以走过的空方格。
- `-1` 表示我们无法跨越的障碍。

返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目**。**

**每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格**。

**示例 1：**

```
输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
输出：2
解释：我们有以下两条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
```

**示例 2：**

```
输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
输出：4
解释：我们有以下四条路径： 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
```

**示例 3：**

```
输入：[[0,1],[2,0]]
输出：0
解释：
没有一条路能完全穿过每一个空的方格一次。
请注意，起始和结束方格可以位于网格中的任意位置。
```

**提示：**

- `1 <= grid.length * grid[0].length <= 20`

## 分析

dfs回溯。  

```go
func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	startR, startC, steps := prepair(grid)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	result := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n ||
			visited[r][c] || grid[r][c] == -1 {
			return
		}
		if grid[r][c] == 2 {
			if steps == 0 {
				result++
			}
			return
		}
		steps--
		visited[r][c] = true
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
		steps++
		visited[r][c] = false
	}
	dfs(startR, startC)
	return result
}

func prepair(grid [][]int) (int, int, int) {
	r, c, steps := 0, 0, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				steps++
			}
			if grid[i][j] == 1 {
				r, c = i, j
			}
		}
	}
	return r, c, steps
}
```
