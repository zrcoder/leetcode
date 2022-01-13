---
title: "1383. 最大的团队表现值"
date: 2021-06-05T17:47:43+08:00
weight: 50
tags: [贪心,堆]
---

## [1383. 最大的团队表现值](https://leetcode-cn.com/problems/maximum-performance-of-a-team/)

难度困难

公司有编号为 `1` 到 `n` 的 `n` 个工程师，给你两个数组 `speed` 和 `efficiency` ，其中 `speed[i]` 和 `efficiency[i]` 分别代表第 `i` 位工程师的速度和效率。请你返回由最多 `k` 个工程师组成的 **最大团队表现值** ，由于答案可能很大，请你返回结果对 `10^9 + 7` 取余后的结果。

**团队表现值** 的定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」。

**示例 1：**

```
输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
输出：60
解释：
我们选择工程师 2（speed=10 且 efficiency=4）和工程师 5（speed=5 且 efficiency=7）。他们的团队表现值为 performance = (10 + 5) * min(4, 7) = 60 。
```

**示例 2：**

```
输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
输出：68
解释：
此示例与第一个示例相同，除了 k = 3 。我们可以选择工程师 1 ，工程师 2 和工程师 5 得到最大的团队表现值。表现值为 performance = (2 + 10 + 5) * min(5, 4, 7) = 68 。
```

**示例 3：**

```
输入：n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
输出：72
```

**提示：**

- `1 <= n <= 10^5`
- `speed.length == n`
- `efficiency.length == n`
- `1 <= speed[i] <= 10^5`
- `1 <= efficiency[i] <= 10^8`
- `1 <= k <= n`

函数签名：

```go
func maxPerformance(n int, speed []int, efficiency []int, k int) int
```

## 分析

幸好团队表现值定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」，而不是「所有工程师速度的和」乘以他们「效率值的和」(或效率值的平均值等)，如果是后边的定义，复杂度一下就上去了，在给定数据范围内必然超时。

既然是`速度和`和`效率最小值`的乘积，就好办了。从`效率最小值`来突破，一个团队的表现值显然因为效率最低的人呈现出木桶短板效应。可以先粗略地先选出效率最高的 k 个员工，计算出团队表现值，这个值被效率最低的那位界定。再看未入选的员工中的效率最高的那位，有没有可能替换已选队伍里的人呢？有可能！最有可能替换的是速度最小的那位。

这样就有了一个贪心的解决方案：

1. 将所有员工按照效率降序排列
2. 遍历，计算表现值，注意选择的员工不能超过 k
3. 对于当前员工，每次尝试用该员工替换已选员工中速度最小的员工，使表现值变大

> 为了迅速找到已选员工中速度最小的，使用小顶堆。

```go
const limit = 1e9 + 7

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return efficiency[indices[i]] > efficiency[indices[j]]
	})
	minHeap := &Heap{}
	var speedSum, res int
	for _, index := range indices {
		speedSum += speed[index]
		heap.Push(minHeap, speed[index])
		if minHeap.Len() == k+1 {
			speedSum -= heap.Pop(minHeap).(int)
		}
		res = max(res, speedSum*efficiency[index])
	}
	return res % limit
}

type Heap []int

func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Len() int            { return len(h) }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return res
}
```

时间复杂度： `O(nlogn)`，空间复杂度: `O(n)`。