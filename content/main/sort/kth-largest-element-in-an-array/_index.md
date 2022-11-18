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
           heap.Pop(h)
        }
    }
    return (*h)[0]
}
```

3.快速选择

时间复杂度 : 平均情况`O(N)`，最坏情况`O(N^2)`。空间复杂度 : `O(1)`

```go
func findKthLargest(nums []int, k int) int {
    var quickSelect func(lo, hi int)
    quickSelect = func(lo, hi int) {
        if lo == hi {
            return
        }
        pivotIndex := lo+rand.Intn(hi-lo) // 在[lo, hi]闭区间选择一个随机元素作为基准元素
        nums[pivotIndex], nums[hi] = nums[hi], nums[pivotIndex] // 先把基准元素放到最后
        pivot := nums[hi]
        j := lo
        for i := lo; i < hi; i++ {
            if nums[i] >= pivot { // 不小于基准元的数字放到前半部分，小于基准元的数字放到后半部分
                nums[i], nums[j] = nums[j], nums[i]
                j++
            }
        }
        nums[j], nums[hi] = nums[hi], nums[j] // 基准元素放到左右部分之间
        if j+1 == k {
            return
        }
        if j+1 > k {
            quickSelect(lo, j-1)
        } else {
            quickSelect(j+1, hi)
        }
    }

    quickSelect(0, len(nums)-1)
    return nums[k-1]
}
```

也可以一开始随机打乱数组，后边每次选择 right 位置的元素为基准元素：

```go
func findKthLargest(nums []int, k int) int {
    rand.Shuffle(len(nums), func(i,j int) {
        nums[i], nums[j] = nums[j], nums[i]
    })

    var help func(lo, hi int)
    help = func(lo, hi int) {
        if lo == hi {
            return
        }
        pivot := nums[hi]
        j := lo
        for i := lo; i < hi; i++ {
            if nums[i] >= pivot {
                nums[i], nums[j] = nums[j], nums[i]
                j++
            }
        }
        nums[j], nums[hi] = nums[hi], nums[j]
        if j+1 == k {
            return
        }
        if j+1 > k {
            help(lo, j-1)
        } else {
            help(j+1, hi)
        }
    }

    help(0, len(nums)-1)
    return nums[k-1]
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

### 参考解答

```go
func kClosest(points [][]int, k int) [][]int {
    rand.Shuffle(len(points), func(i, j int) {
        points[i], points[j] = points[j], points[i]
    })

    var quickSelect func(lo, hi int)
    quickSelect = func(lo, hi int) {
        if lo == hi {
            return
        }
        pivot := points[hi]
        d := dist(pivot)
        j := lo
        for i := lo; i < hi; i++ {
            if dist(points[i]) <= d {
                points[j], points[i] = points[i], points[j]
                j++
            }
        }
        points[j], points[hi] = points[hi], points[j]
        if j+1 == k {
            return
        }
        if j+1 < k {
            quickSelect(j+1, hi)
        } else {
            quickSelect(lo, j-1)
        }
    }

    quickSelect(0, len(points)-1)
    return points[:k]
}

func dist(p []int) int {
    return p[0]*p[0]+p[1]*p[1]
}
```
