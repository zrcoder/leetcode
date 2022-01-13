---
title: "803. 打砖块"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [803. 打砖块](https://leetcode-cn.com/problems/bricks-falling-when-hit/)

难度困难

有一个 `m x n` 的二元网格，其中 `1` 表示砖块，`0` 表示空白。砖块 **稳定**（不会掉落）的前提是：

- 一块砖直接连接到网格的顶部，或者
- 至少有一块相邻（4 个方向之一）砖块 **稳定** 不会掉落时

给你一个数组 `hits` ，这是需要依次消除砖块的位置。每当消除 `hits[i] = (rowi, coli)` 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而掉落。一旦砖块掉落，它会立即从网格中消失（即，它不会落在其他稳定的砖块上）。

返回一个数组 `result` ，其中 `result[i]` 表示第 `i` 次消除操作对应掉落的砖块数目。

**注意**，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。

 

**示例 1：**

```
输入：grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
输出：[2]
解释：
网格开始为：
[[1,0,0,0]，
 [1,1,1,0]]
消除 (1,0) 处加粗的砖块，得到网格：
[[1,0,0,0]
 [0,1,1,0]]
两个加粗的砖不再稳定，因为它们不再与顶部相连，也不再与另一个稳定的砖相邻，因此它们将掉落。得到网格：
[[1,0,0,0],
 [0,0,0,0]]
因此，结果为 [2] 。
```

**示例 2：**

```
输入：grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
输出：[0,0]
解释：
网格开始为：
[[1,0,0,0],
 [1,1,0,0]]
消除 (1,1) 处加粗的砖块，得到网格：
[[1,0,0,0],
 [1,0,0,0]]
剩下的砖都很稳定，所以不会掉落。网格保持不变：
[[1,0,0,0], 
 [1,0,0,0]]
接下来消除 (1,0) 处加粗的砖块，得到网格：
[[1,0,0,0],
 [0,0,0,0]]
剩下的砖块仍然是稳定的，所以不会有砖块掉落。
因此，结果为 [0,0] 。
```

 

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 200`
- `grid[i][j]` 为 `0` 或 `1`
- `1 <= hits.length <= 4 * 104`
- `hits[i].length == 2`
- `0 <= xi <= m - 1`
- `0 <= yi <= n - 1`
- 所有 `(xi, yi)` 互不相同

函数签名：

```go
func hitBricks(grid [][]int, hits [][]int) []int
```
## 分析
对于 hits 中的每个坐标，先消除该位置的砖块，然后考察其上下左右四个邻居是否和第一行砖块联通，如果不联通则就是需要消去的砖块；判断是否联通、消除一片砖块可以用 bfs 或 dfs 的方式。不过这样复杂度非常高。

可以借助并查集降低复杂度。但并查集只能不断合并节点，没法一一消除点。这里可以逆向考虑：

先把 hits 中的所有砖块消去。再逆序遍历 hits，将其中指代的砖块一一补上。每次补一块砖，和上下左右的砖块做一次合并操作，计算合并前后和最顶端一行联通的砖块增加了多少，这就是这次的结果。

可以增加一块虚拟砖代表最顶层。

> 根据需要，Union 不能按秩合并，也不能随机指定顺序合并，而是要保证 from 子树 合并到 to 子树。
> Find 可以做路径压缩，因为只关心每棵树根节点的 size，Find 路径压缩时并不影响树根的大小，虽然会使其他节点的大小不准确。


```go
func hitBricks(grid [][]int, hits [][]int) []int {
    m, n := len(grid), len(grid[0])
    // 拷贝 grid
    status := make([][]int, m)
    for r := range grid {
        status[r] = make([]int, n)
        for c := range status[r] {
            status[r][c] = grid[r][c]
        }
    }
    // 遍历 hits 得到最终状态
    for _, p := range hits {
        r, c := p[0], p[1]
        status[r][c] = 0
    }

    // 根据最终状态建立并查集
    uf := NewUnionFind(m*n+1)
    top := m * n // 额外添加的点，代表最顶部
    for r, row := range status {
        for c, v := range row {
            if v == 0 {
                continue
            }
            id := r*n+c
            if r == 0 {
                uf.Union(id, top)
                continue
            }
            if r > 0 && status[r-1][c] == 1 {
                uf.Union(id, (r-1)*n+c)
            }
            if c > 0 && status[r][c-1] == 1 {
                uf.Union(id, r*n+c-1)
            }
        }
    }

    dirs := [][]int{ {-1, 0}, {1, 0}, {0, -1}, {0, 1} }
    res := make([]int, len(hits))
    for i := len(hits) - 1; i >= 0; i-- {
        r, c := hits[i][0], hits[i][1]
        if grid[r][c] == 0 {
            continue
        }
        preSize := uf.GetSize(top)-1
        id := r*n+c
        if r == 0 {
            uf.Union(id, top)
        }
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if isIn(nr, nc, grid) && status[nr][nc] == 1 {
                uf.Union(id, nr*n+nc)
            }
        }
        status[r][c] = 1
        curSize := uf.GetSize(top)-1
        if cnt := curSize - preSize - 1; cnt > 0 {
            res[i] = cnt
        }
    }
    return res
}

func isIn(r, c int, grid [][]int) bool {
    return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

type UnionFind struct {
    parent, size []int
}

func NewUnionFind(n int) *UnionFind {
    p, s := make([]int, n), make([]int, n)
    for i := range p {
        p[i] = i
        s[i] = 1
    }
    return &UnionFind{parent:p, size:s}
}

func (uf *UnionFind) Union(from, to int) {
    from, to = uf.Find(from), uf.Find(to)
    if from == to {
        return
    }
    uf.parent[from] = to
    uf.size[to] += uf.size[from]
}

func (uf *UnionFind) Find(x int) int {
    for x != uf.parent[x] {
        x, uf.parent[x] = uf.parent[x], uf.parent[uf.parent[x]]
    }
    return x
}

func (uf *UnionFind) GetSize(x int) int {
    x = uf.Find(x)
    return uf.size[x]
}
```
