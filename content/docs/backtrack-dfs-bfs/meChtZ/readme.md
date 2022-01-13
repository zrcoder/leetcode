---
title: "LCP 20. 快速公交"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [LCP 20. 快速公交](https://leetcode-cn.com/problems/meChtZ/)

难度困难

小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：

- 小扣从 `x` 号站点移动至 `x + 1` 号站点需要花费的时间为 `inc`；
- 小扣从 `x` 号站点移动至 `x - 1` 号站点需要花费的时间为 `dec`。

现有 `m` 辆公交车，编号为 `0` 到 `m-1`。小扣也可以通过搭乘编号为 `i` 的公交车，从 `x` 号站点移动至 `jump[i]*x` 号站点，耗时仅为 `cost[i]`。小扣可以搭乘任意编号的公交车且搭乘公交次数不限。

假定小扣起始站点记作 `0`，秋日市集站点记作 `target`，请返回小扣抵达秋日市集最少需要花费多少时间。由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。

注意：小扣可在移动过程中到达编号大于 `target` 的站点。

**示例 1：**

> 输入：`target = 31, inc = 5, dec = 3, jump = [6], cost = [10]`
>
> 输出：`33`
>
> 解释：
> 小扣步行至 1 号站点，花费时间为 5；
> 小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10；
> 小扣从 6 号站台步行至 5 号站台，花费时间为 3；
> 小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10；
> 小扣从 30 号站台步行至 31 号站台，花费时间为 5；
> 最终小扣花费总时间为 33。

**示例 2：**

> 输入：`target = 612, inc = 4, dec = 5, jump = [3,6,8,11,5,10,4], cost = [4,7,6,3,7,6,4]`
>
> 输出：`26`
>
> 解释：
> 小扣步行至 1 号站点，花费时间为 4；
> 小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4；
> 小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3；
> 小扣从 33 号站台步行至 34 站台，花费时间为 4；
> 小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4；
> 小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7；
> 最终小扣花费总时间为 26。

**提示：**

- `1 <= target <= 10^9`
- `1 <= jump.length, cost.length <= 10`
- `2 <= jump[i] <= 10^6`
- `1 <= inc, dec, cost[i] <= 10^6`

函数签名：

```go
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int
```

## 分析

### 正向 dfs

穷举模拟所有可能走法，先 dfs 尝试一把，直接看代码：

```go
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
    const mod = int(1e9) + 7
	memo := map[int]int{}
	// 返回从 start 站到达 target 站所需的最小时间
	var help func(start int) int
	help = func(start int) int {
		if start == target {
			return 0
		}
		if start > target { // 坐公交会越走越远，只能步行返回
			return dec * (start - target)
		}
		if res, ok := memo[start]; ok {
			return res
		}
		// 第一步尝试步行到下一站
		memo[start] = inc+help(start+1)
		// 第一步尝试步行到上一站
		if start > 1 {
			memo[start] = min(memo[start], dec+help(start-1))
		}
		// 第一步尝试坐公交
		for i, v := range jump {
			memo[start] = min(memo[start], cost[i]+help(start*v))
		}
		return memo[start]
	}
	// 在第 0 站坐公交只能回到原地，是无用功，必须先步行到下一站
	return (inc + help(1)) % mod
}
```

发现只过了 示例里的两个测试用例，到了其他用例，发生了栈溢出错误。状态太多，递归栈溢出。

### 逆向 dfs

逆向考虑这个问题，从 0 到 target 跟从 target 到 0 所花费的最小时间应该是一样的。

help 函数修改，help(end int)返回从站 0 到站 end所需的最小时间。

```go
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	const mod = int(1e9) + 7
	memo := map[int]int{}
	// 返回从起点 0 到达 end 站点所需最小时间
	var help func(end int) int
	help = func(end int) int {
		if end == 0 {
			return 0
		}
		// 从 0 站坐公交会回到原点，是无用功，肯定要步行
		if end == 1 {
			return inc
		}
		if res, ok := memo[end]; ok {
			return res
		}
		res := end * inc // 先假设全靠步行
		// 最后一步尝试坐每一辆公交
		for i, v := range jump {
			x := end / v
			// 从 x 站点坐公交达到的站点
			dest := x * v
			if dest == end { // end 可以整除 v
				res = min(res, cost[i]+help(x))
			} else {
				// 即 end 不能整除 v
				res = min(res, cost[i]+help(x)+inc*(end-dest))
				// 尝试从 x+1 坐公交之后步行返回的方案
				dest = (x + 1) * v
				res = min(res, cost[i]+help(x+1)+dec*(dest-end))
			}
		}
		memo[end] = res
		return res
	}
	return help(target) % mod
}
```

### 逆向 bfs

根据上边逆向 dfs 解法，也可以改写成借助堆的 bfs 写法。

```go
type Pos struct {
	//从当前站点到达 target 站点所需的最小时间为 cost
	id, cost int
}

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	const mod = int(1e9) + 7
	h := &Heap{slice: []Pos{ {id: target, cost: 0} } }
	for h.Len() > 0 {
		cur := h.pop()
		if cur.id == 0 {
			return cur.cost % mod
		}
		h.push(Pos{id: 0, cost: cur.cost + inc*cur.id})
		for i, v := range jump {
			x := cur.id / v
			dest := x * v
			h.push(Pos{id: x, cost: cost[i] + cur.cost + inc*(cur.id-dest)})
			if dest < cur.id {
				dest = (x + 1) * v
				h.push(Pos{id: x + 1, cost: cost[i] + cur.cost + dec*(dest-cur.id)})
			}
		}
	}
	return 0
}

type Heap struct {
	slice []Pos
}

func (h *Heap) Len() int           { return len(h.slice) }
func (h *Heap) Less(i, j int) bool { return h.slice[i].cost < h.slice[j].cost }
func (h *Heap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(Pos)) }
func (h *Heap) Pop() interface{} {
	res := h.slice[h.Len()-1]
	h.slice = h.slice[:h.Len()-1]
	return res
}
func (h *Heap) push(p Pos) { heap.Push(h, p) }
func (h *Heap) pop() Pos   { return heap.Pop(h).(Pos) }
```



