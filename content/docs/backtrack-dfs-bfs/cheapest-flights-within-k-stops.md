---
title: "787. K 站中转内最便宜的航班s"
date: 2022-04-07T17:08:12+08:00
---

## [787. K 站中转内最便宜的航班](https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/description/ "https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (38.84%) | 469   | -        |

有 `n` 个城市通过一些航班连接。给你一个数组 `flights` ，其中 `flights[i] = [fromi, toi, pricei]` ，表示该航班都从城市 `fromi` 开始，以价格 `pricei` 抵达 `toi`。

现在给定所有的城市和航班，以及出发城市 `src` 和目的地 `dst`，你的任务是找到出一条最多经过 `k` 站中转的路线，使得从 `src` 到 `dst` 的 **价格最便宜** ，并返回该价格。 如果不存在这样的路线，则输出 `-1`。

**示例 1：**

```
输入: 
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释: 
城市航班图如下


从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
```

**示例 2：**

```
输入: 
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释: 
城市航班图如下


从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
```

**提示：**

- `1 <= n <= 100`
- `0 <= flights.length <= (n * (n - 1) / 2)`
- `flights[i].length == 3`
- `0 <= fromi, toi < n`
- `fromi != toi`
- `1 <= pricei <= 104`
- 航班没有重复，且不存在自环
- `0 <= src, dst, k < n`
- `src != dst`

函数签名：

```go
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int
```

## 分析

最容易想到的是BFS和DFS，根据DFS也可以反推出DP解法。

首先为了能迅速获知某个节点的后续节点，需要把 flights  转化成邻接表。

```go
type Pair struct {
    ID, Price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    nexts := make([][]Pair, n)
    for _, v := range flights {
        nexts[v[0]] = append(nexts[v[0]], Pair{ID: v[1], Price: v[2]})
    }
// ...
}
```

### BFS

如果没有k的限制，BFS 会很容易，有 k 的限制实际上也并不难，只需要维护BFS的层数并与k判断即可。限制 k 个中转站，即限制 k+1 条边，即限制BFS的层数为 k+1。

```go
type Pair struct {
    ID, Price int
}

// inf 既要能表示无限大，又要能在累加的过程中不越界
// 根据数据约束，k的限制为 101，每次航班花费最大 10000, 从起点到终点的最大花费就是 101*10000
const inf = 1010001

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    nexts := make([][]Pair, n)
    for _, v := range flights {
        nexts[v[0]] = append(nexts[v[0]], Pair{ID: v[1], Price: v[2]})
    }

    // memo[x] 表示从src走到x节点花费的最小金额
    memo := make([]int, n)
    for i := range memo {
        memo[i] = inf
    }

    queue := []Pair{{src, 0}}
    memo[src] = 0
    res := inf
    // steps 是 BFS 的层数，限制最多 k+1 次
    for steps := 0; steps <= k && len(queue) > 0; steps++ {
        for size := len(queue); size > 0; size-- {
            cur := queue[0]
            queue = queue[1:]
            for _, v := range nexts[cur.ID] {
                price := cur.Price + v.Price
                if v.ID == dst {
                    // 不要直接返回，可能有多条路径到达 dst 节点
                    res = min(res, price)
                    continue
                }
                if price > res || memo[v.ID] <= price {
                    continue
                }
                queue = append(queue, Pair{v.ID, price})
                memo[v.ID] = price
            }
        }
    }

    if res == inf {
        return -1
    }
    return res
}
```

### DFS

实际上，可以定义一个递归函数 help 来计算从一个指定节点 start 到达最终 dst 节点的最小花费，除了 start，还需有一个参数 k 来携带步数限制信息。

因为有较多重复计算，带上备忘录优化。

```go
type Pair struct {
    ID, Price int
}

const inf = 1010001

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    nexts := make([][]Pair, n)
    for _, v := range flights {
        nexts[v[0]] = append(nexts[v[0]], Pair{ID: v[1], Price: v[2]})
    }
    // help 函数的备忘录
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, k+2)
    } 
    // 返回从 start 节点到 dst 节点，限制步数为 k 时的最少花费
    var help func(start, k int) int
    help = func(start, k int) int {
        if k < 0 {
            return inf
        }
        if start == dst {
            return 0
        }
        if memo[start][k] != 0 {
            return memo[start][k]
        }
        res := inf
        for _, v := range nexts[start] {
            res = min(res, v.Price+help(v.ID, k-1))
        }
        memo[start][k] = res
        return res
    }
    res := help(src, k+1)
    if res == inf {
        return -1
    }
    return res
}
```

### 动态规划

从自顶向下的 dfs 解法，容易想到自底向上的动态规划解法。

定义 dp, `dp[k][end]` 表示经过 k 次航行，到达城市 end 需要的最小花费。

```go
const inf = 1010001

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    dp := make([][]int, k+2)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][src] = 0 // dp[*][src] 都应为0，不过这里只需初始化dp[0][src]
    // 注意 k 的遍历在外层，航班信息遍历在内层，不能反——为什么？
    for t := 1; t <= k+1; t++ {
        for _, v := range flights {
            i, j, cost := v[0], v[1], v[2]
            dp[t][j] = min(dp[t][j], dp[t-1][i]+cost)
        }
    }
    res := inf
    for t := 1; t <= k+1; t++ {
        res = min(res, dp[t][dst])
    }
    if res == inf {
        res = -1
    }
    return res
}
```

比起BFS 和 DFS，代码最简洁，无须预处理 flights。另外 dp 数组可以优化为两个一维数组。

实际这就是 `Bellman Ford` 算法。
