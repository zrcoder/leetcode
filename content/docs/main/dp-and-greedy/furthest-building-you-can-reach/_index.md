---
title: "1642. 可以到达的最远建筑"
date: 2021-04-27T20:47:20+08:00
weight: 50
tags: [贪心, 堆]
---

## [1642. 可以到达的最远建筑](https://leetcode-cn.com/problems/furthest-building-you-can-reach/)

难度中等

给你一个整数数组 `heights` ，表示建筑物的高度。另有一些砖块 `bricks` 和梯子 `ladders` 。

你从建筑物 `0` 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。

当从建筑物 `i` 移动到建筑物 `i+1`（下标 **从 0 开始** ）时：

- 如果当前建筑物的高度 **大于或等于** 下一建筑物的高度，则不需要梯子或砖块
- 如果当前建筑的高度 **小于** 下一个建筑的高度，您可以使用 **一架梯子** 或 **`(h[i+1] - h[i])` 个砖块**

如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 **从 0 开始** ）。

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/31/q4.gif)

```
输入：heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
输出：4
解释：从建筑物 0 出发，你可以按此方案完成旅程：
- 不使用砖块或梯子到达建筑物 1 ，因为 4 >= 2
- 使用 5 个砖块到达建筑物 2 。你必须使用砖块或梯子，因为 2 < 7
- 不使用砖块或梯子到达建筑物 3 ，因为 7 >= 6
- 使用唯一的梯子到达建筑物 4 。你必须使用砖块或梯子，因为 6 < 9
无法越过建筑物 4 ，因为没有更多砖块或梯子。
```

**示例 2：**

```
输入：heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
输出：7
```

**示例 3：**

```
输入：heights = [14,3,19,3], bricks = 17, ladders = 0
输出：3
```

**提示：**

- `1 <= heights.length <= 105`
- `1 <= heights[i] <= 106`
- `0 <= bricks <= 109`
- `0 <= ladders <= heights.length`

函数签名：

```go
func furthestBuilding(heights []int, bricks int, ladders int) int
```

## 分析

遍历所有楼，在遇到比当前高的楼时，既可以选择用梯子也可以选择用砖块，可以回溯枚举所有可能来计算最终结果。

上边这样的策略复杂度非常高，其实还可以调整。一个明显的事实是，如果面对的楼非常高，用梯子是优于用砖块的。所以需要把梯子用在刀刃上，即非常大的高度差出现时用梯子，砖块只用来填补较小的高度差。

可以先假设对于前 `ladders` 个高度差都用梯子，这样来到了索引为 `ladders` 的位置。现在每遇到一座新楼，先将其与当前楼的高度差加入之前假设的用了梯子的高度差里集合里，再找到集合中最小的高度差，用砖块填补。这相当于之前假设用梯子的那些高度差，逐渐用换成用砖块解决较低的那些高度差，剩下较高的高度差都是梯子解决。如果当前砖块不够填补当前鸿沟，也就到达了最远的位置。

用一个小顶堆维护访问过的高度差再适合不过，这样能在对数级时间获知最低的高度差。

```go
func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &Heap{cmp: func(a, b int) bool {
		return a < b
	}}
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}
		heap.Push(h, diff)
		if h.Len() <= ladders {
			continue
		}
		bricks -= heap.Pop(h).(int)
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}
```

时间复杂度：`O(nlogx)`，其中 n 是数组长度，x 是梯子数量。

空间复杂度：`O(x)`。

### 附

堆的实现：

```go
type Cmp func(int, int) bool

type Heap struct {
	slice []int
	cmp   Cmp
}

// implement heap.Interface
func (h *Heap) Len() int           { return len(h.slice) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.slice[i], h.slice[j]) }
func (h *Heap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.slice[h.Len()-1]
	h.slice = h.slice[:h.Len()-1]
	return x
}
```



