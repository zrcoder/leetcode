package solution

/*
## [1615. Maximal Network Rank](https://leetcode.cn/problems/maximal-network-rank) (Medium)

`n` 座城市和一些连接这些城市的道路 `roads` 共同组成一个基础设施网络。每个 `roads[i] = [aᵢ, bᵢ]` 都表示在城市 `aᵢ` 和 `bᵢ` 之间有一条双向道路。

两座不同城市构成的 **城市对** 的 **网络秩** 定义为：与这两座城市 **直接** 相连的道路总数。如果存在一条道路直接连接这两座城市，则这条道路只计算 **一次** 。

整个基础设施网络的 **最大网络秩** 是所有不同城市对中的 **最大网络秩** 。

给你整数 `n` 和数组 `roads`，返回整个基础设施网络的 **最大网络秩** 。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/11/ex1.png)**

```
输入：n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
输出：4
解释：城市 0 和 1 的网络秩是 4，因为共有 4 条道路与城市 0 或 1 相连。位于 0 和 1 之间的道路只计算一次。

```

**示例 2：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/11/ex2.png)**

```
输入：n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
输出：5
解释：共有 5 条道路与城市 1 或 2 相连。

```

**示例 3：**

```
输入：n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
输出：5
解释：2 和 5 的网络秩为 5，注意并非所有的城市都需要连接起来。

```

**提示：**

- `2 <= n <= 100`
- `0 <= roads.length <= n * (n - 1) / 2`
- `roads[i].length == 2`
- `0 <= aᵢ, bᵢ <= n-1`
- `aᵢ != bᵢ`
- 每对城市之间 **最多只有一条** 道路相连


*/

// [start] don't modify
func maximalNetworkRank(n int, roads [][]int) int {
    graph := make(map[int]map[int]bool, n)
    for _, v := range roads {
        if graph[v[0]] == nil {
            graph[v[0]] = map[int]bool{}
        }
        if graph[v[1]] == nil {
            graph[v[1]] = map[int]bool{}
        }
        graph[v[0]][v[1]] = true
        graph[v[1]][v[0]] = true
    }
    res := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            cur := len(graph[i]) + len(graph[j])
            if graph[i][j] {
                cur--
            }
            res = max(res, cur)
        }
    }
    return res
}
// [end] don't modify