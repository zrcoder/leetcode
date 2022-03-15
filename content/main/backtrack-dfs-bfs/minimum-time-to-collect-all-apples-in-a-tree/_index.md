---
title: "1443. 收集树上所有苹果的最少时间"
date: 2022-03-15T09:27:52+08:00
---

## [1443. 收集树上所有苹果的最少时间](https://leetcode-cn.com/problems/minimum-time-to-collect-all-apples-in-a-tree/)

难度中等

给你一棵有 `n` 个节点的无向树，节点编号为 `0` 到 `n-1` ，它们中有一些节点有苹果。通过树上的一条边，需要花费 1 秒钟。你从 **节点 0** 出发，请你返回最少需要多少秒，可以收集到所有苹果，并回到节点 0 。

无向树的边由 `edges` 给出，其中 `edges[i] = [fromi, toi]` ，表示有一条边连接 `from` 和 `toi` 。除此以外，还有一个布尔数组 `hasApple` ，其中 `hasApple[i] = true` 代表节点 `i` 有一个苹果，否则，节点 `i` 没有苹果。

**示例 1：**

****

**输入：** n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
**输出：** 8 
**解释：** 上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。

**示例 2：**

****

**输入：** n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,false,true,false]
**输出：** 6
**解释：** 上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。

**示例 3：**

**输入：** n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,false,false,false,false,false]
**输出：** 0

**提示：**

- `1 <= n <= 10^5`
- `edges.length == n-1`
- `edges[i].length == 2`
- `0 <= fromi, toi <= n-1`
- `fromi < toi`
- `hasApple.length == n`

函数签名：

```go
func minTime(n int, edges [][]int, hasApple []bool) int
```

## 分析

对于某个节点 node，其代表的子树上如果没有苹果，则无需统计采到苹果的路径；反之需要统计。可以用 dfs 来做，定义 `func dfs(node, cost int) int`，表示以 node 为根的子树采摘完所有苹果并回到node的最小耗时， 其中 cost 表示从祖先节点到达node需要的时间耗费。

另为了能迅速获知每个节点的相邻节点，需要事先根据edges数组得到neibors数组。

```go
func minTime(n int, edges [][]int, hasApple []bool) int {
    neibors := make([][]int, n)
    for _, v := range edges {
        neibors[v[0]] = append(neibors[v[0]], v[1])
        neibors[v[1]] = append(neibors[v[1]], v[0])
    }
    seen := make([]bool, n)
    var dfs func(node, cost int) int
    dfs = func(node, cost int) int {
        if seen[node] {
            return 0
        }
        seen[node] = true
        childrenCost := 0
        for _, v := range neibors[node] {
            childrenCost += dfs(v, 2)
        }
        if childrenCost == 0 && !hasApple[node] {
            return 0
        }
        return cost + childrenCost
    }
    return dfs(0, 0)
}
```

时间复杂度 `O(n)`，空间复杂度 `O(n)`。