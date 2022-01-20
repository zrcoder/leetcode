---
title: "834. 树中距离之和"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [834. 树中距离之和](https://leetcode-cn.com/problems/sum-of-distances-in-tree/)

难度困难

给定一个无向、连通的树。树中有 `N` 个标记为 `0...N-1` 的节点以及 `N-1` 条边 。

第 `i` 条边连接节点 `edges[i][0]` 和 `edges[i][1]` 。

返回一个表示节点 `i` 与其他所有节点距离之和的列表 `ans`。

**示例 1:**

```
输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
输出: [8,12,6,10,10,10]
解释: 
如下为给定的树的示意图：
  0
 / \
1   2
   /|\
  3 4 5

我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5) 
也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
```

**说明:** `1 <= N <= 10000`

函数签名：

```go
func sumOfDistancesInTree(n int, edges [][]int) []int
```

## 分析

以题目中所给的树为例，看节点2，其他节点和它的举例和，可以分成两部分：

```
1. 节点2 为根的子树中的节点，与它的距离和。
2. 子树之外的节点与它的距离和。
```

先来攻克 1

使用树形动态规划，定义 `dp[u] ` 表示子树 `u` 的所有子节点到根节点 `u` 距离之和。可以看到这个值和该子树的节点数量有关。同时定义 `sz[u]` 表示子树 `u` 的节点数量，则容易得到状态转移方程：

$$
dp[u] = \sum_{v \in son[u]}dp[v]+sz[v]
$$

其中 `son[u]` 表示 `u` 的所有儿子节点集合。

再来攻克 2  

可以直接枚举所有的节点，对每一个节点，使用上边的方法来求以其为根的结果即可，不过这样复杂度比较高。

实际可以借助已有的`dp`结果。比如例子中，求节点 `2` 为根的结果，可以从原始根节点 `0` 推出来：

从节点 `0` 换到节点 `2`，要少走 `sz[2]` 步，同时要多走 `n-sz[2]` 步，所以有如下推导公式：

`ans(2) = (dp[0] - sz[2]) + (n-sz[2])`

如果用先序遍历，将不会影响后续节点。

```go
func sumOfDistancesInTree(n int, edges [][]int) []int {
	// 为方便迅速得到某个节点的相邻节点，将输入处理成邻接表
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	// 当前记录的是每个节点到它所在子树的节点的距离和
	dp := make([]int, n)
	// 记录每个节点作为根节点的子树的节点总数
	sz := make([]int, n)
	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		sz[node] = 1 // 节点自身个数需计算
		for _, v := range graph[node] {
			if v == parent {
				continue
			}
			dfs(v, node)
			sz[node] += sz[v]
			dp[node] += dp[v] + sz[v]
		}
	}
	dfs(0, -1)

	// 做换根操作，之后的 dp[u] 表示节点 `u` 到其他节点的距离和。
	var dfs1 func(node, parent int)
	dfs1 = func(node, parent int) {
		for _, v := range graph[node] {
			if v == parent {
				continue
			}
			dp[v] = dp[node] - sz[v] + (n - sz[v])
			dfs1(v, node)
		}
	}
	dfs1(0, -1) // 第一个参数要和上边 dfs 调用时一致
	return dp
}
```

