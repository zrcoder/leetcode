---
title: 373. 查找和最小的 K 对数字
---

## [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums) (Medium)

给定两个以 **升序排列** 的整数数组 `nums1` 和 `nums2`, 以及一个整数 `k`。

定义一对值 `(u,v)`，其中第一个元素来自 `nums1`，第二个元素来自 `nums2`。

请找到和最小的 `k` 个数对 `(u₁,v₁)`, ` (u₂,v₂)` ... `(uₖ,vₖ)` 。

**示例 1:**

```
输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

```

**示例 2:**

```
输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

```

**示例 3:**

```
输入: nums1 = [1,2], nums2 = [3], k = 3
输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

```

**提示:**

- `1 <= nums1.length, nums2.length <= 10⁵`
- `-10⁹ <= nums1[i], nums2[i] <= 10⁹`
- `nums1` 和 `nums2` 均为升序排列
- `1 <= k <= 1000`

## 分析


首先，无法用双指针，不存在对应的贪心策略

可以用类似 BFS 的方法，首先可以把 nums1[0] 和 nums2[0] 加入集合（BFS中一般为双端队列），然后开始 k 次遍历。

每次从集合中取出和最小的 nums1[i] 和 nums2[j], 再把 (i+1, j) 及 （i， j+1) 放入集合。

为了能迅速得到最小和，集合用小顶堆即可。

```go
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	if m*n == 0 {
		return nil
	}
	if k > m*n {
		k = m * n
	}
	vis := map[Info]bool{}
	h := &Heap{}
	info := Info{0, 0, nums1[0] + nums2[0]}
	h.push(info)
	vis[info] = true
	res := make([][]int, k)
	for i := range res {
		info := h.pop()
		res[i] = []int{nums1[info.i], nums2[info.j]}
		if info.i < m-1 {
			p1 := Info{info.i + 1, info.j, nums1[info.i+1] + nums2[info.j]}
			if !vis[p1] {
				h.push(p1)
				vis[p1] = true
			}
		}
		if info.j < n-1 {
			p2 := Info{info.i, info.j + 1, nums1[info.i] + nums2[info.j+1]}
			if !vis[p2] {
				h.push(p2)
				vis[p2] = true
			}
		}
	}
	return res
}

```

时间复杂度是 O(klogk)，空间复杂度O(m*n*sum)，其中 sum 指可能的元素对的和。空间复杂度主要是哈希表 vis，实际可以优化，不需要 vis。

空间复杂度降为小顶堆占用的空间，即 O(k)。

```go

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	if m*n == 0 {
		return nil
	}
	if k > m*n {
		k = m * n
	}
	h := &Heap{}
	h.push(Info{0, 0, nums1[0] + nums2[0]})
	res := make([][]int, k)
	for i := range res {
		info := h.pop()
		res[i] = []int{nums1[info.i], nums2[info.j]}
		if info.j == 0 && info.i < m-1 {
			next := Info{info.i + 1, info.j, nums1[info.i+1] + nums2[info.j]}
			h.push(next)
		}
		if info.j < n-1 {
			next := Info{info.i, info.j + 1, nums1[info.i] + nums2[info.j+1]}
			h.push(next)
		}
	}
	return res
}

type Info struct {
	i, j, val int
}

type Heap struct {
	data []Info
}

func (h *Heap) Len() int           { return len(h.data) }
func (h *Heap) Less(i, j int) bool { return h.data[i].val < h.data[j].val }
func (h *Heap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Push(x any)         { h.data = append(h.data, x.(Info)) }
func (h *Heap) Pop() any {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}
func (h *Heap) push(i Info) { heap.Push(h, i) }
func (h *Heap) pop() Info   { return heap.Pop(h).(Info) }

```
