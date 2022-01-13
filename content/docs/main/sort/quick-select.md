---
title: "快速选择算法"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

这个算法适用于计算数组中第 `k` 大或第 `k` 小或第 `k` 个满足某种规则的元素。

实际可以看作部分完成的快速排序，其平均复杂度非常优秀，是线性的复杂度。

## [215. 数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

难度中等

在未排序的数组中找到第 **k** 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

**示例 1:**

```
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
```

**示例 2:**

```
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
```

**说明:**

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

函数签名：

```go
func findKthLargest(nums []int, k int) int
```

## 分析

### 排序

最自然的思路，先把所有元素排序，然后就很容易得出最终的结果了：

```go
func findKthLargest0(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
```

时间复杂度 `O(nlgn)`，空间复杂度 `O(1)`。

### 借助堆

还可以维持一个大小为 `k`的小顶堆，遍历并把所有元素塞到堆里，注意当堆大小超过 `k` 时要把堆顶元素出堆，让堆大小为 `k`。最后，堆顶元素就是所求。

```go
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			_ = heap.Pop(h)
		}
	}
	return (*h)[0]
}
```

时间复杂度 `O(nlgk)`,空间复杂度 `O(k)`。

### 快速选择

实际就是快速排序的思想，但是不需要把所有元素都排序，只需要在某次选完基准元素并把所有大于等于它的元素放到它左边，小于它的元素放它右边后发现基准元素最终落在了第 `k` 的位置就结束排序。直接返回这个元素。

```go
func findKthLargest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return 0
	}
	quickSelect(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func quickSelect(nums []int, left, right, k int) {
	if left == right { // 递归结束条件：区间里仅有一个元素
		return
	}
	pivotIndex := partition(nums, left, right)
	if pivotIndex+1 == k {
		return
	}
	if pivotIndex+1 > k {
		quickSelect(nums, left, pivotIndex-1, k)
	} else {
		quickSelect(nums, pivotIndex+1, right, k)
	}
}

// 以pivotIndex处元素做划分，不妨称这个元素为基准元素，大于基准的放在左侧，小于基准的放在右侧
// 返回最终基准元素的索引
func partition(nums []int, left, right int) int {
	// 0. 在区间[left, right]里随机选一个索引
	pivotIndex := left + rand.Intn(right-left+1)
	pivot := nums[pivotIndex]
	// 1. 先把基准元素放到最后
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	storeIndex := left
	// 2. 把所有大于等于基准元素的元素放到左侧
	for i := left; i < right; i++ {
		if nums[i] >= pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	// 3. 基准元素放到最终位置
	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}
```

也可以一开始随机打乱数组，后边每次选择 right 位置的元素为基准元素，代码更简洁：

```go
func findKthLargest(nums []int, k int) int {
    if k < 1 || k > len(nums) {
	    return 0    
    }
    // 一开始随机打乱所有元素，减少特定输入对性能的影响
    rand.Seed(time.Now().UnixNano())    rand.Shuffle(len(nums), func(i, j int) {
    	nums[i], nums[j] = nums[j], nums[i]
    })
    quickSelect(nums, 0, len(nums)-1, k)
    return nums[k-1]
}

func quickSelect(nums []int, left, right, k int) {
	if left == right { // 递归结束条件：区间里仅有一个元素
		return
	}
	pivotIndex := partition(nums, left, right)
	if pivotIndex+1 == k {
		return
	}
	if pivotIndex+1 > k {
		quickSelect(nums, left, pivotIndex-1, k)
	} else {
		quickSelect(nums, pivotIndex+1, right, k)
	}
}

// 以 right 处元素为基准元素，大于等于基准的放在左侧，小于基准的放在右侧
// 返回最终基准元素的索引
func partition(nums []int, left, right int) int {
	pivot := nums[right]
	storeIndex := left
	// 把所有大于等于基准元素的元素放到左侧
	for i := left; i < right; i++ {
		if nums[i] >= pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	// right 处的元素交换到 storeIndex 处
	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}
```

时间复杂度 : 平均情况 `O(N)`，最坏情况 `O(N^2)`。

空间复杂度 :  `O(1)`。

优秀的是平均时间复杂度，详细证明可以看《算法导论》快速选择算法相关章节。

**这个算法适用于计算数组中 `第k` 个满足某种规则的元素，但是不能计算 `前k` 个满足规则的元素。**