---
title: "1168. 水资源分配优化"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1168. 水资源分配优化](https://leetcode-cn.com/problems/optimize-water-distribution-in-a-village)

村里面一共有 n 栋房子。我们希望通过建造水井和铺设管道来为所有房子供水。  
对于每个房子 i，我们有两种可选的供水方案：  
一种是直接在房子内建造水井，成本为 wells[i]；  
另一种是从另一口井铺设管道引水，数组 pipes 给出了在房子间铺设管道的成本，  
其中每个 pipes[i] = [house1, house2, cost] 代表用管道将 house1 和 house2 连接在一起的成本。  
当然，连接是双向的。  
请你帮忙计算为所有房子都供水的最低总成本。

```
示例：
输入：n = 3, wells = [1,2,2], pipes = [[1,2,1],[2,3,1]]
输出：3
解释：
上图展示了铺设管道连接房屋的成本。
最好的策略是在第一个房子里建造水井（成本为 1），然后将其他房子铺设管道连起来（成本为 2），所以总成本为 3。
提示：
1 <= n <= 10000
wells.length == n
0 <= wells[i] <= 10^5
1 <= pipes.length <= 10000
1 <= pipes[i][0], pipes[i][1] <= n
0 <= pipes[i][2] <= 10^5
pipes[i][0] != pipes[i][1]
```

## 解析

增加一个房子，其他房子直接在其内部建井相当于和这个新增的房子联接；  
问题转化成了[最低成本联通所有城市](/main/graph/connecting-cities-with-minimum-cost)

```go
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	for i, cost := range wells {
		house := i + 1
		pipes = append(pipes, []int{0, house, cost})
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})
	connected, i, result := 0, 0, 0
	uf := NewUnionFind(n + 1)
	for connected < n {
		pipe := pipes[i]
		i++
		house1, house2, cost := pipe[0], pipe[1], pipe[2]
		house1, house2 = uf.Find(house1), uf.Find(house2)
		if house1 != house2 {
			uf.Union(house1, house2)
			result += cost
			connected++
		}
	}
	return result
}
```
