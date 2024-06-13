---
title: "1834. 单线程 CPU"
date: 2022-12-29T11:02:40+08:00
---
## [1834. 单线程 CPU](https://leetcode.cn/problems/single-threaded-cpu/)

难度中等

给你一个二维数组 `tasks` ，用于表示 `n` 项从 `0` 到 `n - 1` 编号的任务。其中 `tasks[i] = [enqueueTimei, processingTimei]` 意味着第 `i` 项任务将会于 `enqueueTimei` 时进入任务队列，需要 `processingTimei` 的时长完成执行。

现有一个单线程 CPU ，同一时间只能执行 **最多一项** 任务，该 CPU 将会按照下述方式运行：

* 如果 CPU 空闲，且任务队列中没有需要执行的任务，则 CPU 保持空闲状态。
* 如果 CPU 空闲，但任务队列中有需要执行的任务，则 CPU 将会选择 **执行时间最短** 的任务开始执行。如果多个任务具有同样的最短执行时间，则选择下标最小的任务开始执行。
* 一旦某项任务开始执行，CPU 在 **执行完整个任务** 前都不会停止。
* CPU 可以在完成一项任务后，立即开始执行一项新任务。

返回 CPU 处理任务的顺序。

**示例 1：**

**输入：** tasks = [[1,2],[2,4],[3,2],[4,1]]
**输出：** [0,2,3,1]
**解释：** 事件按下述流程运行：

* time = 1 ，任务 0 进入任务队列，可执行任务项 = {0}
* 同样在 time = 1 ，空闲状态的 CPU 开始执行任务 0 ，可执行任务项 = {}
* time = 2 ，任务 1 进入任务队列，可执行任务项 = {1}
* time = 3 ，任务 2 进入任务队列，可执行任务项 = {1, 2}
* 同样在 time = 3 ，CPU 完成任务 0 并开始执行队列中用时最短的任务 2 ，可执行任务项 = {1}
* time = 4 ，任务 3 进入任务队列，可执行任务项 = {1, 3}
* time = 5 ，CPU 完成任务 2 并开始执行队列中用时最短的任务 3 ，可执行任务项 = {1}
* time = 6 ，CPU 完成任务 3 并开始执行任务 1 ，可执行任务项 = {}
* time = 10 ，CPU 完成任务 1 并进入空闲状态

**示例 2：**

**输入：** tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
**输出：** [4,3,2,0,1]
**解释：** 事件按下述流程运行：

* time = 7 ，所有任务同时进入任务队列，可执行任务项 = {0,1,2,3,4}
* 同样在 time = 7 ，空闲状态的 CPU 开始执行任务 4 ，可执行任务项 = {0,1,2,3}
* time = 9 ，CPU 完成任务 4 并开始执行任务 3 ，可执行任务项 = {0,1,2}
* time = 13 ，CPU 完成任务 3 并开始执行任务 2 ，可执行任务项 = {0,1}
* time = 18 ，CPU 完成任务 2 并开始执行任务 0 ，可执行任务项 = {1}
* time = 28 ，CPU 完成任务 0 并开始执行任务 1 ，可执行任务项 = {}
* time = 40 ，CPU 完成任务 1 并进入空闲状态

**提示：**

* `tasks.length == n`
* `1 <= n <= 10^5`
* `1 <= enqueueTimei, processingTimei <= 10^9`

函数签名：

```go
func getOrder(tasks [][]int) []int
```

## 分析

模拟即可。先将所有任务安装开始时间排序，对于当前时间，挑选开始时间不晚于当前时间的任务中耗时最短的来执行，需要借助小顶堆来做。

```go
type Task struct {
	Index, Start, Cost int
}

type TaskHeap struct {
	ts []*Task
}

func (h *TaskHeap) Len() int { return len(h.ts) }
func (h *TaskHeap) Less(i, j int) bool {
	return h.ts[i].Cost < h.ts[j].Cost ||
		h.ts[i].Cost == h.ts[j].Cost && h.ts[i].Index < h.ts[j].Index
}
func (h *TaskHeap) Swap(i, j int)      { h.ts[i], h.ts[j] = h.ts[j], h.ts[i] }
func (h *TaskHeap) Push(x interface{}) { h.ts = append(h.ts, x.(*Task)) }
func (h *TaskHeap) Pop() interface{} {
	n := len(h.ts)
	x := h.ts[n-1]
	h.ts = h.ts[:n-1]
	return x
}
func (h *TaskHeap) push(t *Task) { heap.Push(h, t) }
func (h *TaskHeap) pop() *Task   { return heap.Pop(h).(*Task) }

func getOrder(tasks [][]int) []int {
	ts := make([]*Task, len(tasks))
	for i, v := range tasks {
		ts[i] = &Task{
			Index: i,
			Start: v[0],
			Cost:  v[1],
		}
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Start < ts[j].Start
	})

	time := 0
	h := &TaskHeap{}
	i := 0
	res := make([]int, 0, len(ts))
	for h.Len() > 0 || i < len(ts) {
		for i < len(ts) && ts[i].Start <= time {
			h.push(ts[i])
			i++
		}
		if h.Len() == 0 && i < len(ts) {
			time = ts[i].Start
			continue
		}
		cur := h.pop()
		res = append(res, cur.Index)
		time += cur.Cost
	}
	return res
}
```

时间复杂度：`O(nlogn)`，空间复杂度：`O(n)`。
