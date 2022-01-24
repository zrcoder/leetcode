---
title: "数组中的第K个元素"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [215. 数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

`难度中等`

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

函数签名如下：
```go
func findKthLargest(nums []int, k int) int
```

## 分析
比较直观易解的问题。值得注意的是基于快排思想的快速选择解法

1.朴素实现，时间复杂度O(nlgn)，空间复杂度O(1)
```go
func findKthLargest0(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
```

2.使用堆

借助一个小顶堆，将nums里的元素一一入堆，但需要保持堆的大小最多为k，如果超出k，需要把堆顶元素出堆
最后堆顶元素就是结果。

时间复杂度O(nlgk),空间复杂度O(k)；实际测试，并不比朴素实现快

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
func findKthLargest1(nums []int, k int) int {
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

3.快速选择

时间复杂度 : 平均情况O(N)，最坏情况O(N^2)。空间复杂度 : O(1)
```go
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right { // 递归结束条件：区间里仅有一个元素
		return nums[left]
	}
	// 在区间[left, right]里随机选一个索引
	pivotIndex := left + rand.Intn(right-left+1)
	pivotIndex = partition(nums, left, right, pivotIndex)
	if pivotIndex+1 == k {
		return nums[pivotIndex]
	}
	if pivotIndex+1 > k {
		return quickSelect(nums, left, pivotIndex-1, k)
	}
	return quickSelect(nums, pivotIndex+1, right, k)
}

// 以pivotIndex处元素做划分，不妨称这个元素为基准元素，大于基准的放在左侧，小于基准的放在右侧
// 返回最终基准元素的索引
func partition(nums []int, left, right, pivotIndex int) int {
	// 1. 先把基准元素放到最后
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]

	pivot := nums[right]
	storeIndex := left
	// 2. 把所有大于基准元素的元素放到左侧
	for i := left; i < right; i++ {
		if nums[i] > pivot {
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
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
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
实际测试发现比上边每次都随机选基准元素的实现稍微耗时。

## 扩展： [973. 最接近原点的 K 个点](https://leetcode-cn.com/problems/k-closest-points-to-origin)
`难度中等`

我们有一个由平面上的点组成的列表 `points`。需要从中找出 `K` 个距离原点 `(0, 0)` 最近的点。

（这里，平面上两点之间的距离是欧几里德距离。）

你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

 

**示例 1：**

```
输入：points = [[1,3],[-2,2]], K = 1
输出：[[-2,2]]
解释： 
(1, 3) 和原点之间的距离为 sqrt(10)，
(-2, 2) 和原点之间的距离为 sqrt(8)，
由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
```

**示例 2：**

```
输入：points = [[3,3],[5,-1],[-2,4]], K = 2
输出：[[3,3],[-2,4]]
（答案 [[-2,4],[3,3]] 也会被接受。）
```

 

**提示：**

1. `1 <= K <= points.length <= 10000`
2. `-10000 < points[i][0] < 10000`
3. `-10000 < points[i][1] < 10000`

## 参考代码
```go
var k int

func kClosest(points [][]int, K int) [][]int {
    k = K
    quickSelect(points, 0, len(points)-1)
    return points[:k]
}

func quickSelect(points [][]int, lo, hi int) {
    if lo == hi {
        return
    }
    
    index := partion(points, lo, hi)
    if index == k-1 {
        return
    }
    if index < k-1 {
        quickSelect(points, index, hi)
    } else {
        quickSelect(points, lo, index)
    }
}

func partion(points [][]int, lo, hi int) int {
    index := lo + rand.Intn(hi-lo+1)
    piv := points[index]
    val := piv[0]*piv[0] + piv[1]*piv[1]
    points[index], points[hi] = points[hi], points[index]
    storedIndex := lo
    for i := lo; i < hi; i++ {
        p := points[i]        
        if p[0]*p[0]+p[1]*p[1] <= val {
            points[i], points[storedIndex] = points[storedIndex], points[i]
            storedIndex++
        }
    }
    points[storedIndex], points[hi] = points[hi], points[storedIndex]
    return storedIndex
}
```
