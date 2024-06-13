---
title: 815. 公交路线
date: 2023-11-12T16:25:32+08:00
---

## [815. 公交路线](https://leetcode.cn/problems/bus-routes) (Hard)

给你一个数组 `routes` ，表示一系列公交线路，其中每个 `routes[i]` 表示一条公交线路，第 `i` 辆公交车将会在上面循环行驶。

- 例如，路线 `routes[0] = [1, 5, 7]` 表示第 `0` 辆公交车会一直按序列 `1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ...` 这样的车站路线行驶。

现在从 `source` 车站出发（初始时不在公交车上），要前往 `target` 车站。 期间仅可乘坐公交车。

求出 **最少乘坐的公交车数量** 。如果不可能到达终点车站，返回 `-1` 。

**示例 1：**

```
输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
输出：2
解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。

```

**示例 2：**

```
输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
输出：-1

```

**提示：**

- `1 <= routes.length <= 500`.
- `1 <= routes[i].length <= 10⁵`
- `routes[i]` 中的所有值 **互不相同**
- `sum(routes[i].length) <= 10⁵`
- `0 <= routes[i][j] < 10⁶`
- `0 <= source, target < 10⁶`

## 分析

以公交线路为节点建图，再BFS即可。

> 可以借助一个哈希表（代码中的 buses）来建图。

```go
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	n := len(routes)
	graph := make([][]int, n) // grap[i]  代表能从 i 路公交换乘的公交线路列表
	seen := make([]bool, n)
	q := []int{}
	buses := map[int][]int{} // buses[x] 代表经过 x 站点的公交线路列表
	for i, route := range routes {
		for _, stop := range route {
			for _, j := range buses[stop] {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
			buses[stop] = append(buses[stop], i)
			if stop == source {
				q = append(q, i)
				seen[i] = true
			}
		}
	}
	for step := 1; len(q) > 0; step++ {
		var tmp []int
		for _, bus := range q {
			if slices.Contains(routes[bus], target) {
				return step
			}
			for _, next := range graph[bus] {
				if seen[next] {
					continue
				}
				seen[next] = true
				tmp = append(tmp, next)
			}
		}
		q = tmp
	}
	return -1
}

```
