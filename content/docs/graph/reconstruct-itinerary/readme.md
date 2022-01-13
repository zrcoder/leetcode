---
title: "332. 重新安排行程"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [332. 重新安排行程](https://leetcode-cn.com/problems/reconstruct-itinerary/)

难度中等

给定一个机票的字符串二维数组 `[from, to]`，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。

 

**提示：**

1. 如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
2. 所有的机场都用三个大写字母表示（机场代码）。
3. 假定所有机票至少存在一种合理的行程。
4. 所有的机票必须都用一次 且 只能用一次。

 

**示例 1：**

```
输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
```

**示例 2：**

```
输入：[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
```

函数签名

```go
func findItinerary(tickets [][]string) []string
```

## 分析
这是数学里的欧拉“七桥问题”，即“一笔画”问题。主要难点在于图有环。

先不考虑字符排序，尝试深度优先遍历。

对于当前节点 `x`， 再向下有多于一种走法，需要优先把可能形成环的路径先走完。如下：

```     
  a <- x -> b
       ^    |
        \   v
           c 
```

题目保证了有解，即对于 `x` ，在接下来的选择中可以有多个环路，但最多有一个非环路（死胡同），否则无法走完所有的边。

DFS 是可以比较方便地确定某一条路径是否会回到当前节点的，即除了记录节点是否访问过，对于访问过的节点额外记录状态即可，参考拓扑排序中 DFS 做法。

但这样依次去判读每个选择，复杂度太高，且要考虑对于当前节点，虽然选择某条路能回到当前节点，但是这条路可能中间又叉出去一条死胡同，情况比较复杂。

Hierholzer 算法比较优雅地解决了这个问题。

对于当前节点，循环对所有的选择做 DFS，每做一次选择，就`删除对应的边`防止环路回来后再次走这条边——当然也可以用一个额外的备忘录来记录；

并且保证在所有子 DFS 完成后再将当前节点加入结果，即`后序遍历`。

这样所有的环路都会返回到当前节点，只有死胡同不会返回。使用后序遍历，可以保证死胡同的点最先入栈。

可以对照如上例子模拟基于 `x` 点的递归过程，可以发现，无论先走 `x->a` 这条死胡同，还是先走 `x->b` 这条环路，最终都能保证死胡同先加入结果。

> 注意，`x` 可以在结果中多次出现。

最终只需要将结果逆序即可。

```go
func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string, len(tickets))
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		graph[src] = append(graph[src], dst)
	}
	var res []string
	var dfs func(cur string)
	dfs = func(cur string) {
		for len(graph[cur]) > 0 {
			next := graph[cur][0]
			graph[cur] = graph[cur][1:]
			dfs(next)
		}
		res = append(res, cur)
	}
	dfs("JFK")
	reverse(res)
	return res
}
```

```go
func reverse(s []string) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
```

再考虑下字符排序，只需要在 dfs 前，对邻接表中每个点的邻居节点排下序就行：
```go
	for key := range graph {
		sort.Strings(graph[key])
	}
```

时间复杂度是 `O(mlogm)`, `m` 是边的数量。

空间复杂度是 `O(m)`。