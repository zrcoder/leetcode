---
title: "搜索有序二维数组"
date: 2021-04-19T22:04:56+08:00
weight: 1
tags: [二维数组]
---

编写一个高效的算法来判断 m x n 数组中，是否存在一个目标值。数组按照一定规则排序。
## [74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix)
逐行遍历，整体单调递增
```
示例 1:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true

示例 2:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
```
## 分析
### 二分法
如果把这个矩阵压平成一个一维数组，则是一个递增的一维数组，可以用二分法来做  
实际并不用真正转换一个一维数组，矩阵本身可以看做一行行拼接的一维数组  
从`0`到`M*N-1`遍历， 对于`i`，可以容易得到在矩阵里对应的行和列：`r = i/N, c = i%N`

时间复杂度`O(lg(M*N))=O(lgM + lgN)`, 空间复杂度`O(1)`
```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid/n][mid%n] == target {
			return true
		}
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
```

使用标准库的一个二分法版本：
```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	return i != m*n && matrix[i/n][i%n] == target
}
```
### 线性时间复杂度的解法
也可以从右上角或左下角开始搜索  
假设从右上角开始。如果元素大于`target`，则向左走一格；如果元素小于`target`，则向下走一格  
时间复杂度`O(M+N)`，比二分法的复杂度高一些， 空间复杂度`O(1)`
```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1
	for r < m && c >= 0 {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] < target {
			r++
		} else {
			c--
		}
	}
	return false
}
```
## [240. 搜索二维矩阵 II](https://leetcode-cn.com/problems/search-a-2d-matrix-ii)
从左到右，从上到下递增
```
示例:
现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
```

算法、代码同上个问题线性时间复杂度解法

> 如果稍微改变下排列特征，比如矩阵从左到右递减而从上到下递增，都可以用这种技巧，即从左上、右上、左下、右下某个顶点开始遍历即可。

