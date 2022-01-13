---
title: "1697. 检查边长度限制的路径是否存在"
date: 2022-12-14T13:04:46+08:00
---

#### [1697. 检查边长度限制的路径是否存在](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/)

难度困难

给你一个 `n` 个点组成的无向图边集 `edgeList` ，其中 `edgeList[i] = [ui, vi, disi]` 表示点 `ui` 和点 `vi` 之间有一条长度为 `disi` 的边。请注意，两个点之间可能有 **超过一条边** 。

给你一个查询数组`queries` ，其中 `queries[j] = [pj, qj, limitj]` ，你的任务是对于每个查询 `queries[j]` ，判断是否存在从 `pj` 到 `qj` 的路径，且这条路径上的每一条边都 **严格小于** `limitj` 。

请你返回一个 **布尔数组** `answer` ，其中 `answer.length == queries.length` ，当 `queries[j]` 的查询结果为 `true` 时， `answer` 第 `j` 个值为 `true` ，否则为 `false` 。

**示例 1：**

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/12/19/h.png)

**输入：** n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
**输出：**[false,true]
**解释：** 上图为给定的输入数据。注意到 0 和 1 之间有两条重边，分别为 2 和 16 。
对于第一个查询，0 和 1 之间没有小于 2 的边，所以我们返回 false 。
对于第二个查询，有一条路径（0 -> 1 -> 2）两条边都小于 5 ，所以这个查询我们返回 true 。

**示例 2：**

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/12/19/q.png)

**输入：** n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
**输出：**[true,false]
**解释：** 上图为给定数据。

**提示：**

- `2 <= n <= 105`
- `1 <= edgeList.length, queries.length <= 105`
- `edgeList[i].length == 3`
- `queries[j].length == 3`
- `0 <= ui, vi, pj, qj <= n - 1`
- `ui != vi`
- `pj != qj`
- `1 <= disi, limitj <= 109`
- 两个点之间可能有 **多条** 边。

函数签名：

```go
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool
```

## 分析

### 建图+BFS（超时）

朴素解法，先根据 edgeList 构建图，然后对每个查询做BFS搜索来确定两点之间是否有符合要求的通路。

```go
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
    graph := make([][][]int, n)
    for i := range graph {
        graph[i] = make([][]int, n)
    }
    for _, e := range edgeList {
        u, v, dis := e[0], e[1], e[2]
        graph[u][v] = append(graph[u][v], dis)
        graph[v][u] = append(graph[v][u], dis)
    }

    res := make([]bool, len(queries))
    for i, v := range queries {
        res[i] = bfs(v, graph)
    }
    return res
}

func bfs(query []int, graph [][][]int) bool {
    u, v, lim := query[0], query[1], query[2]
    n := len(graph)
    seen := make([]bool, n)
    q := []int{u}
    seen[u] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if cur == v {
            return true
        }
        for next, paths := range graph[cur] {
            if seen[next] || !hasLimitedPath(paths, lim) {
                continue
            }
            seen[next] = true
            q = append(q, next)
        }
    }
    return false
}

func hasLimitedPath(paths []int, limit int) bool {
    for _, dis := range paths {
        if dis < limit {
            return true
        }
    }
    return false
}
```

时间复杂度：`O(n*m)`，空间复杂度：`O(n^2)`。

### 排序+并查集

可以用这样一个贪心策略：将edgeList和queries都按照边长/限制边长排序，然后遍历 queries，对于当前限制，用并查集将小于当前限制的点联通，然后用并查集查看查询到两点是否联通即可，到下次查询，可以直接利用这次联通的结果。

```go
var uf []int

func find(x int) int {
    if x != uf[x] {
        uf[x] = find(uf[x])
    }
    return uf[x]
}

func union(x, y int) {
    x, y = find(x), find(y)
    uf[x] = y
}

type Query struct {
    u, v, lim, index int
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
    uf = make([]int, n)
    for i := range uf {
        uf[i] = i
    }

    qs := make([]*Query, len(queries))
    for i, v := range queries {
        qs[i] = &Query{u:v[0], v:v[1], lim:v[2], index:i}
    }
    sort.Slice(qs, func(i, j int) bool {
        return qs[i].lim < qs[j].lim
    })
    sort.Slice(edgeList, func(i, j int) bool {
        return edgeList[i][2] < edgeList[j][2]
    })

    edgeIndex := 0
    res := make([]bool, len(qs))
    for _, v := range qs {
        for edgeIndex < len(edgeList) && edgeList[edgeIndex][2] < v.lim {
            e := edgeList[edgeIndex]
            union(e[0], e[1])
            edgeIndex++
        }
        res[v.index] = find(v.u) == find(v.v)
    }
    return res
}
```

时间复杂度：`O(ElogE+mlogm+(E+m)logn+n)`，其中 E 是 edgeList 的长度，m 是 queries 的长度，n 是点数。

空间复杂度`O(logE+m+n)`。
