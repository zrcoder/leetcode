---
title: "399. 除法求值"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [399. 除法求值](https://leetcode-cn.com/problems/evaluate-division/)

难度中等

给你一个变量对数组 `equations` 和一个实数值数组 `values` 作为已知条件，其中 `equations[i] = [Ai, Bi]` 和 `values[i]` 共同表示等式 `Ai / Bi = values[i]` 。每个 `Ai` 或 `Bi` 是一个表示单个变量的字符串。

另有一些以数组 `queries` 表示的问题，其中 `queries[j] = [Cj, Dj]` 表示第 `j` 个问题，请你根据已知条件找出 `Cj / Dj = ?` 的结果作为答案。

返回 **所有问题的答案** 。如果存在某个无法确定的答案，则用 `-1.0` 替代这个答案。

 

**注意：**输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

 

**示例 1：**

```
输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
```

**示例 2：**

```
输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]
```

**示例 3：**

```
输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]
```

 

**提示：**

- `1 <= equations.length <= 20`
- `equations[i].length == 2`
- `1 <= Ai.length, Bi.length <= 5`
- `values.length == equations.length`
- `0.0 < values[i] <= 20.0`
- `1 <= queries.length <= 20`
- `queries[i].length == 2`
- `1 <= Cj.length, Dj.length <= 5`
- `Ai, Bi, Cj, Dj` 由小写英文字母与数字组成

函数签名：

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64
```

## 分析

构建图做搜索

如果这个问题里的除法变成加法、减法、乘法，基本就无解了。比如知道 a+b 和 b+c 的值，并不能知道 a+c 的值， 乘法、减法类似。

而除法就有传递性，知道 a/b, b/c, c/d 的值，一定能知道 a/c 和 a/d 的值，比如 a/c = a/b * b/c。

可以在纸上画一张图，每个变量作为图的顶点，根据 equations 和 values，可以画出从某个顶点 x 到另一个顶点 y 的箭头，也就是边，其权值为 x/y 的值，当然也可以再画出从 y 到 x 的箭头， 其值为 y/x 的值。

这就是一张有向图。

如果要计算两个变量 x、y 的比值，实际就是在图里从 x 点出发，一直找到 y 点，将途径的权值都乘起来即可。

### BFS
BFS 方便求最短路径问题，这里再合适不过

```go
type Node = string

var graph map[Node]map[Node]float64

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	if n != len(values) {
		return nil
	}
	graph = make(map[Node]map[Node]float64, 2*n) // graph[i][j]代表图中顶点 i 到 j 的权值， 即 i/j 的值
	for i, v := range equations {
		if graph[v[0]] == nil {
			graph[v[0]] = make(map[Node]float64, 1)
		}
		if graph[v[1]] == nil {
			graph[v[1]] = make(map[Node]float64, 1)
		}
		graph[v[0]][v[1]] = values[i]
		graph[v[1]][v[0]] = 1 / values[i]
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		if graph[q[0]] == nil || graph[q[1]] == nil { // 图中没有要查询的点
			res[i] = -1
		} else {
			res[i] = bfs(q[0], q[1])
		}
	}
	return res
}

func bfs(start, end Node) float64 {
	queue := list.New()
	queue.PushBack(start)
	// visited[i] 既记录 i 点是否访问过，也记录了起点到 i 点的权值
	visited := make(map[Node]float64, len(graph))
	visited[start] = 1
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(Node)
		if cur == end {
			return visited[cur]
		}
		for next, w := range graph[cur] {
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = visited[cur] * w
			queue.PushBack(next)
		}
	}
	return -1
}
```

由于题目限定变量字符串长度 <= 5, 可以忽略其影响。

时间复杂度：O(M+Q⋅(M)) = O(QM)，其中 M 为边的数量，Q 为询问的数量。

空间复杂度：O(N+M) = O(2M + M) = O(M)，其中 N 为点的数量，最大为 2M。

### DFS
也可以用 DFS 方法：

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// ... 第一部分代码同上，略
	res := make([]float64, len(queries))
	for i, q := range queries {
		if graph[q[0]] == nil || graph[q[1]] == nil { // 图中没有要查询的点
			res[i] = -1
		} else {            
			visited := make(map[Node]float64, len(graph)) 
			visited[q[0]] = 1.0
			res[i] = dfs(q[0], q[1], visited)
		}
	}
	return res
}

func dfs(start, end Node, visited map[Node]float64) float64 {
	if start == end {
		return 1
	}
	for next, w := range graph[start] {
		if _, ok := visited[next]; ok {
			continue
		}
		visited[next] = visited[start] * w
		dfs(next, end, visited)
		if res, ok := visited[end]; ok {
			return res
		}
	}
	return -1
}
```

复杂度同 BFS 的解法。

### 动态规划

如果查询次数很多，每次查询都要做一次搜索，效率会比较低。

可以用备忘录提前记录图中任意两点之间的权值，可以使用动态规划，这就是 Floyd 算法。

Floyd 算法的核心代码非常简短，只有三层循环，是一个比较精巧的动态规划。需要花点时间理解其原理。这里不再展开。

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// ... 第一部分代码同上，略
	
	// 使用动态规划，即 Floyd 算法提前计算出图中任意两点间的权值
	for k := range graph {
        for i := range graph {
            for j := range graph {
                if graph[i][k] > 0 && graph[k][j] > 0 {
                    graph[i][j] = graph[i][k] * graph[k][j]
                }
            }
        }
    }

	res := make([]float64, len(queries))
	 for i, q := range queries {
        if graph[q[0]] == nil || graph[q[1]] == nil || graph[q[0]][q[1]] == 0 {
            res[i] = -1
        } else {
            res[i] = graph[q[0]][q[1]]
        }
    }
	return res
}
```

注意 N 最大为 2M

时间复杂度为 O(M + N^3 + Q) = O(M^3 + Q)

空间复杂度为 O(N + N^2) = O(M^2)

## 并查集

带权并查集，做好路径压缩，使并查集最多有 3 层，并在查询两个节点的时候 Find 函数保证这两个节点都在第二层。

Union 和 Find 函数都需要更新节点的权值。

官方有一个非常详细的题解，声图文并茂，[链接在这里](https://leetcode-cn.com/problems/evaluate-division/solution/399-chu-fa-qiu-zhi-nan-du-zhong-deng-286-w45d) 。

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	if n != len(values) {
		return nil
	}
	//  给所有节点编号
	id := make(map[string]int, 2*n)
	for _, v := range equations {
		if _, ok := id[v[0]]; !ok {
			id[v[0]] = len(id)
		}
		if _, ok := id[v[1]]; !ok {
			id[v[1]] = len(id)
		}
	}

	uf := NewUnionFind(len(id))
	for i, v := range equations {
		x, y := id[v[0]], id[v[1]]
		uf.Union(x, y, values[i])
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		x, hasX := id[q[0]]
		y, hasY := id[q[1]]
		if !hasX || !hasY {
			res[i] = -1
			continue
		}
		rootX, rootY := uf.Find(x), uf.Find(y)
		if rootX != rootY {
			res[i] = -1
		} else {
			res[i] = uf.weight[x] / uf.weight[y]
		}
	}
	return res
}

type UnionFind struct {
	parent []int
	weight []float64
}

func NewUnionFind(n int) UnionFind {
	p := make([]int, n)
	w := make([]float64, n)
	for i := range p {
		p[i] = i
		w[i] = 1.0
	}
	return UnionFind{parent: p, weight: w}
}

func (uf UnionFind) Union(x, y int, w float64) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf.parent[rootX] = rootY
	uf.weight[rootX] = w * uf.weight[y] / uf.weight[x]
}

// 寻找 x 的根节点，在寻找的过程中做路径压缩
func (uf UnionFind) Find(x int) int {
	// 迭代法不太好更新权值，这里用递归法
	if x != uf.parent[x] {
		p := uf.parent[x]
		uf.parent[x] = uf.Find(uf.parent[x])
		uf.weight[x] *= uf.weight[p]
	}
	return uf.parent[x]
}
```