package solution

/*
## [1514. Path with Maximum Probability](https://leetcode.cn/problems/path-with-maximum-probability) (Medium)

给你一个由 `n` 个节点（下标从 0 开始）组成的无向加权图，该图由一个描述边的列表组成，其中 `edges[i] = [a, b]` 表示连接节点 a 和 b 的一条无向边，且该边遍历成功的概率为 `succProb[i]` 。

指定两个节点分别作为起点 `start` 和终点 `end` ，请你找出从起点到终点成功概率最大的路径，并返回其成功概率。

如果不存在从 `start` 到 `end` 的路径，请 **返回 0** 。只要答案与标准答案的误差不超过 **1e-5**，就会被视作正确答案。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/07/12/1558_ex1.png)**

```
输入：n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
输出：0.25000
解释：从起点到终点有两条路径，其中一条的成功概率为 0.2 ，而另一条为 0.5 * 0.5 = 0.25

```

**示例 2：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/07/12/1558_ex2.png)**

```
输入：n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
输出：0.30000

```

**示例 3：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/07/12/1558_ex3.png)**

```
输入：n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
输出：0.00000
解释：节点 0 和 节点 2 之间不存在路径

```

**提示：**

- `2 <= n <= 10^4`
- `0 <= start, end < n`
- `start != end`
- `0 <= a, b < n`
- `a != b`
- `0 <= succProb.length == edges.length <= 2*10^4`
- `0 <= succProb[i] <= 1`
- 每两个节点之间最多有一条边


*/

// [start] don't modify
/*
    先构建图，再dfs/bfs遍历即可。用一个备忘录数组 memo 记录从 start 到达每个点的最大概率，最后返回 memo[end] 即可
    以下是 dfs 解法：
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type pair struct {
		id   int
		rate float64
	}
	graph := make([][]pair, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], pair{id: e[1], rate: succProb[i]})
		graph[e[1]] = append(graph[e[1]], pair{id: e[0], rate: succProb[i]})
	}
	memo := make([]float64, n)
	memo[start] = 1.0
	var dfs func(node int)
	dfs = func(node int) {
		rate := memo[node]
		for _, p := range graph[node] {
			if rate*p.rate > memo[p.id] {
				memo[p.id] = rate * p.rate
				dfs(p.id)
			}
		}
	}
	dfs(start)
	return memo[end]
}

    实测超时，改用 bfs, 时间复杂度是O(mn),其中m、n 分别是边和节点的数量
*/
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type pair struct {
		id   int
		rate float64
	}
	graph := make([][]pair, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], pair{id: e[1], rate: succProb[i]})
		graph[e[1]] = append(graph[e[1]], pair{id: e[0], rate: succProb[i]})
	}
	memo := make([]float64, n)
	memo[start] = 1.0
	queue := []pair{{id: start, rate: 1.0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, p := range graph[cur.id] {
			if cur.rate*p.rate > memo[p.id] {
				memo[p.id] = cur.rate * p.rate
				queue = append(queue, pair{id: p.id, rate: cur.rate * p.rate})
			}
		}
	}
	return memo[end]
}
/*
   也可以用优先队列替换queue，并在遇到end时提前返回，理论上是个优化，时间复杂度会下降到O(m*log(mn))实测效果不明显

import "container/heap"

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]pair, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], pair{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], pair{e[0], succProb[i]})
	}
	memo := make([]float64, n)
	memo[start] = 1
	h := &hp{{start, 1}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		if cur.id == end { 
			return cur.rate
		}
		for _, next := range graph[cur.id] {
			if d := cur.rate * next.rate; d > memo[next.id] {
				memo[next.id] = d
				heap.Push(h, pair{next.id, d})
			}
		}
	}

	return memo[end]
}

type pair struct {
	id   int
	rate float64
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].rate > h[j].rate }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
*/

// [end] don't modify
