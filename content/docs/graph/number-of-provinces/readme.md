---
title: "547. 省份数量"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [547. 省份数量](https://leetcode-cn.com/problems/number-of-provinces/)

难度中等

有 `n` 个城市，其中一些彼此相连，另一些没有相连。如果城市 `a` 与城市 `b` 直接相连，且城市 `b` 与城市 `c` 直接相连，那么城市 `a` 与城市 `c` 间接相连。

**省份** 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 `n x n` 的矩阵 `isConnected` ，其中 `isConnected[i][j] = 1` 表示第 `i` 个城市和第 `j` 个城市直接相连，而 `isConnected[i][j] = 0` 表示二者不直接相连。

返回矩阵中 **省份** 的数量。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg)

```
输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg)

```
输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
输出：3
```

 

**提示：**

- `1 <= n <= 200`
- `n == isConnected.length`
- `n == isConnected[i].length`
- `isConnected[i][j]` 为 `1` 或 `0`
- `isConnected[i][i] == 1`
- `isConnected[i][j] == isConnected[j][i]`



函数签名：

```go
func findCircleNum(isConnected [][]int) int
```

## 分析

可以把 n 个城市和它们之间的相连关系看成一张图，城市就是各个顶点，相连关系就是图中的边，isConnected 就是其邻接矩阵，最终求省份数，就是求联通分量的个数。

### DFS

```go
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	marked := make([]bool, n)
	res := 0
	for i, v := range marked {
		if v {
			continue
		}
		res++
		mark(i, marked, isConnected)
	}
	return res
}

// 将和城市 from 联通的城市都标记, dfs
func mark(from int, marked []bool, isConnected [][]int) {
	marked[from] = true
	for to, conn := range isConnected[from] {
		if conn == 1 && !marked[to] {
			mark(to, marked, isConnected)
		}
	}
}
```

时间复杂度 O(n^2)， 空间复杂度 O(n)

### BFS

框架同 dfs， 只是 mark 方法用 bfs：

```go
// 将和城市 from 联通的城市都标记, bfs
func mark(from int, marked []bool, isConnected [][]int) {
	queue := list.New()
	queue.PushBack(from)
	marked[from] = true
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)
		for next, conn := range isConnected[cur] {
			if conn == 1 && !marked[next] {
				queue.PushBack(next)
				marked[next] = true
			}
		}
	}
}
```

复杂度同上。

### 并查集

```go
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    uf := NewUnionFind(n)
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if isConnected[r][c] == 1 {
                union(uf, r, c)
            }
        }
    }
    return cal(uf)
}

func NewUnionFind(n int) []int {
    uf := make([]int, n)
    for i := range uf {
        uf[i] = i
    }
    return uf
}

func union(uf []int, x, y int) {
    rootX, rootY := find(uf, x), find(uf, y)
    uf[rootX] = rootY
}

func find(uf []int, x int) int {
    for x != uf[x] {
        x, uf[x] = uf[x], uf[uf[x]]
    }
    return x
}

func cal(uf []int) int {
    res := 0
    for i, v := range uf {
        if i == v {
            res++
        }
    }
    return res
}
```

复杂度同上。并查集的 union 和 find 方法复杂度可以看作常数级。实际上这里做了路径压缩，但没有按秩合并，最坏情况下并查集复杂度是 log n， 但是平均复杂度还是常数级。