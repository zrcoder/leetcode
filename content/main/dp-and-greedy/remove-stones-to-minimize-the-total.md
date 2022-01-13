---
title: "1962. 移除石子使总数最小"
date: 2022-12-28T10:28:38+08:00
---
## [1962. 移除石子使总数最小](https://leetcode.cn/problems/remove-stones-to-minimize-the-total/)

难度中等

给你一个整数数组 `piles` ，数组 **下标从 0 开始** ，其中 `piles[i]` 表示第 `i` 堆石子中的石子数量。另给你一个整数 `k` ，请你执行下述操作 **恰好** `k` 次：

* 选出任一石子堆 `piles[i]` ，并从中 **移除** `floor(piles[i] / 2)` 颗石子。

**注意：** 你可以对 **同一堆** 石子多次执行此操作。

返回执行 `k` 次操作后，剩下石子的 **最小** 总数。

`floor(x)` 为 **小于** 或 **等于** `x` 的 **最大** 整数。（即，对 `x` 向下取整）。

**示例 1：**

**输入：** piles = [5,4,9], k = 2
**输出：** 12
**解释：** 可能的执行情景如下：

* 对第 2 堆石子执行移除操作，石子分布情况变成 [5,4, ***5*** ] 。
* 对第 0 堆石子执行移除操作，石子分布情况变成 [ ***3*** ,4,5] 。
  剩下石子的总数为 12 。

**示例 2：**

**输入：** piles = [4,3,6,7], k = 3
**输出：** 12
**解释：** 可能的执行情景如下：

* 对第 2 堆石子执行移除操作，石子分布情况变成 [4,3, ***3*** ,7] 。
* 对第 3 堆石子执行移除操作，石子分布情况变成 [4,3,3, ***4*** ] 。
* 对第 0 堆石子执行移除操作，石子分布情况变成 [ ***2*** ,3,3,4] 。
  剩下石子的总数为 12 。

**提示：**

* `1 <= piles.length <= 105`
* `1 <= piles[i] <= 104`
* `1 <= k <= 105`

函数签名：

```go
func minStoneSum(piles []int, k int) int
```

## 分析

### 贪心+大顶堆

每次消除最大一堆石头的一半即可。

```go
func minStoneSum(piles []int, k int) int {
    h := &Heap{piles}
    h.init()
    for ; k > 0 && h.peek() > 1; k-- {
        cur := h.pop()
        h.push(cur - cur/2)
    }
    sum := 0
    for _, v := range h.data {
        sum += v
    }
    return sum
}

type Heap struct {
    data []int
}

func (h *Heap) Len() int           { return len(h.data) }
func (h *Heap) Less(i, j int) bool { return h.data[i] > h.data[j] }
func (h *Heap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(int)) }
func (h *Heap) Pop() interface{} {
    n := len(h.data)
    x := h.data[n-1]
    h.data = h.data[:n-1]
    return x
}
func (h *Heap) init() {heap.Init(h)}
func (h *Heap) push(x int) { heap.Push(h, x) }
func (h *Heap) pop() int   { return heap.Pop(h).(int) }
func (h *Heap) peek() int  { return h.data[0] }
```

时间复杂度：`O(n+klogn)`，空间复杂度：`O(n)`。
