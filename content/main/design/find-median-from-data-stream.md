---
title: "数据流/滑动窗口中位数"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [堆, 有序集合]
---

## [295. 数据流的中位数](https://leetcode-cn.com/problems/find-median-from-data-stream/)

难度困难

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

- void addNum(int num) - 从数据流中添加一个整数到数据结构中。
- double findMedian() - 返回目前所有元素的中位数。

**示例：**

```
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3) 
findMedian() -> 2
```

**进阶:**

1. 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
2. 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

## 分析

### 朴素实现

维护一个有序的切片，这样可以在常数时间找到中位数，但是添加元素需要 O(n) 的复杂度，n 是已有元素的总数。

添加元素的复杂度不理想，实现代码略。

### BST

Treap 是较平衡的 BST, AVL或红黑树平衡性更理想。在平衡的 BST 里增删查元素是对数级的复杂度，同样，查找第 k 小的数字也是对数时间复杂度。

不过可惜的是标准库里没有实现。且 AVL、红黑树等实现较复杂。倒是可以手写 Treap，略。

### 两个堆优化

可以用一个大顶堆来保存所有元素中较小的一半，再用一个小顶堆来保存较大的另一半。

假设这两个堆分别名为 small 和 large，只需要在添加元素的时候保持两个堆的大小相当（相等或 small 比 large 多一个），这样查找中位数还是常数级的复杂度，添加元素的复杂度优化成了 O(logn) 。

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
// local functions
func (h *Heap) push(x int) { heap.Push(h, x) }
func (h *Heap) pop() int   { return heap.Pop(h).(int) }
func (h *Heap) peek() int  { return h.slice[0] }

type MedianFinder struct {
	small, large *Heap
}

func Constructor() MedianFinder {
	res := MedianFinder{&Heap{}, &Heap{}}
	res.small.cmp = func(a, b int) bool {
		return a > b
	}
	res.large.cmp = func(a, b int) bool {
		return a < b
	}
	return res
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.small.Len() == 0 || num <= mf.small.peek() {
		mf.small.push(num)
	} else {
		mf.large.push(num)
	}
	mf.makeBalance()
}

func (mf *MedianFinder) makeBalance() {
	if mf.small.Len() > mf.large.Len()+1 {
		mf.large.push(mf.small.pop())
	} else if mf.small.Len() < mf.large.Len() {
		mf.small.push(mf.large.pop())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.small.Len() > mf.large.Len() {
		return float64(mf.small.peek())
	}
	if mf.large.Len() == 0 && mf.small.Len() == 0 {
		return 0
	}
	return float64(mf.small.peek()+mf.large.peek()) * 0.5
}
```

## 拓展

一、两个思考问题：

1. 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？

   > 只需要维护每个数字的数量，这样插入元素是常数级复杂度，然后可以用类似计数排序的方式，找到中位数，注意不用真的排序，只需要统计下比指定元素小的有多少个。这个复杂度是 O(100)，常数级。

2. 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

   > 非常像计数的思路，不过这里是将数据划分到一些桶里，可以迅速定位到中位数在哪个桶里，之后在那个桶里做简单排序就找到了中位数

二、如果要支持删除元素呢？

   注意上边两个堆的解法能较好地解决元素一直追加的中位数问题，但是如果要支持能删除元素并求中位数，就不是这么容易了。比如下边的问题。

## [480. 滑动窗口中位数](https://leetcode-cn.com/problems/sliding-window-median/)

难度困难

中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

- `[2,3,4]`，中位数是 `3`
- `[2,3]`，中位数是 `(2 + 3) / 2 = 2.5`

给你一个数组 *nums*，有一个大小为 *k* 的窗口从最左端滑动到最右端。窗口中有 *k* 个数，每次窗口向右移动 *1* 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

 **示例：**

给出 *nums* = `[1,3,-1,-3,5,3,6,7]`，以及 *k* = 3。

```
窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
```

 因此，返回该滑动窗口的中位数数组 `[1,-1,-1,3,5,6]`。

 **提示：**

- 你可以假设 `k` 始终有效，即：`k` 始终小于输入的非空数组的元素个数。
- 与真实值误差在 `10 ^ -5` 以内的答案将被视作正确答案。

函数签名：

```go
func medianSlidingWindow(nums []int, k int) []float64
```

## 分析

### 朴素解法

维护一个有序数组，插入、删除的复杂度较高，但在 LeetCode 实测效果挺好，时间空间都打败了 96% 左右的提交~应该是测试用例的问题。
代码略。

时间复杂度 O(n*k)，空间复杂度 O(k)。

怎么降低时间复杂度呢？可以用平衡二叉搜索树，如红黑树、AVL树等，不过这些数据结构手写还是很复杂的。

### 两个堆+延迟删除

如上边《数据流的中位数》问题中两个堆的解法应用到这个问题会比较困难，难在从堆里删除元素的复杂度高，需要遍历一遍先找到那个元素才行（实际是找到元素的索引，再调用标准库的 Remove方法）。

但是确实也有办法，可以延迟删除元素: 删除一个元素的时候先不从堆里删除，而是将元素记录下来；在 pop 后、push 前、remove 后再做一个操作：循环检查堆顶元素，如果在待删除的缓存中则删除。

可以用一个哈希表维护待删除元素，键为元素，因可能有重复元素，值为待删除的个数。

这样实际支持了一个remove、push 和 pop 都是对数级复杂度的堆。

```go
func medianSlidingWindow(nums []int, k int) []float64 {
	mf := NewMedianFinder()
	res := make([]float64, 0, len(nums)-k+1)
	for i, v := range nums {
		mf.Put(v)
		if i >= k-1 {
			res = append(res, mf.FindMedian())
			mf.Remove(nums[i-k+1])
		}
	}
	return res
}

type MedianFinder struct {
	left, right *Heap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		left:  &Heap{cmp: func(x, y int) bool { return x > y }, memo: map[int]int{}},
		right: &Heap{cmp: func(x, y int) bool { return x < y }, memo: map[int]int{}},
	}
}

func (mf *MedianFinder) Put(num int) {
	mf.left.push(num)
	mf.right.push(mf.left.pop())
	mf.makeBalance()
}
func (mf *MedianFinder) Remove(x int) {
	if mf.left.size > 0 && mf.left.peek() >= x {
		mf.left.remove(x)
	} else {
		mf.right.remove(x)
	}
	mf.makeBalance()
}

func (mf *MedianFinder) makeBalance() {
	// 注意这里用size而不是Len()
	if mf.left.size > mf.right.size+1 {
		mf.right.push(mf.left.pop())
	} else if mf.left.size < mf.right.size {
		mf.left.push(mf.right.pop())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.size > mf.right.size {
		return float64(mf.left.peek())
	}
	return float64(mf.left.peek()+mf.right.peek()) / 2
}

type Heap struct {
	items []int
	cmp   func(int, int) bool
	memo  map[int]int // 记录应该删除的元素
	size  int         // 记录堆预期的大小，因为删除延迟的原因，堆的实际大小(items的大小)并非预期大小
}

func (h *Heap) Len() int           { return len(h.items) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.items[i], h.items[j]) }
func (h *Heap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *Heap) Push(x interface{}) { h.items = append(h.items, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(h.items)
	res := h.items[n-1]
	h.items = h.items[:n-1]
	return res
}

// push 向堆里添加一个元素
func (h *Heap) push(x int) {
	h.clean()
	h.size++
	heap.Push(h, x)
}

// pop 删除堆顶元素并返回其值
func (h *Heap) pop() int {
	h.clean()
	h.size--
	return heap.Pop(h).(int)
}

// peek 获取堆顶元素
func (h *Heap) peek() int {
	h.clean()
	return h.items[0]
}

// remove 在堆里删除元素 num
func (h *Heap) remove(x int) {
	h.clean()
	h.size--
	h.memo[x]++
}

// 循环检查堆顶元素，如果在 memo 缓存中则删除
func (h *Heap) clean() {
	for len(h.items) > 0 && h.memo[h.items[0]] > 0 {
		h.memo[h.items[0]]--
		heap.Pop(h)
	}
}
```

## 拓展

除了延迟删除技巧，还有一个思路，实际比延迟删除更好。
维护每个元素在堆里的索引，在删除时用标准库的Remove方法。
详见注释。

```go
type Heap struct {
	cmp  Comparator
	data []int
	idx  map[int]int // 维护每个元素在 data 数组中的索引
	cnt  map[int]int // 可能有多个相同元素入堆，我们仅在 data 中维护去重后的元素，cnt 维护每个元素的个数
	size int         // size 维护堆的总大小，这个值大于等于 len(data)
}

type Comparator func(a, b int) bool

func New(cmp Comparator) *Heap {
	return &Heap{
		data: make([]int, 0),
		idx:  make(map[int]int),
		cnt:  make(map[int]int),
		cmp:  cmp,
	}
}

// for stantard heap package ====
func (h *Heap) Len() int           { return len(h.data) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.data[i], h.data[j]) }
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.idx[h.data[i]] = i
	h.idx[h.data[j]] = j
}
func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(h.data)
	res := h.data[n-1]
	h.data = h.data[:n-1]
	return res
}

// === end

func (h *Heap) push(x int) {
	if h.cnt[x] == 0 {
		h.idx[x] = len(h.data)
		heap.Push(h, x)
	}
	h.cnt[x]++
	h.size++
}
func (h *Heap) pop() int {
	res := h.data[0]
	if h.cnt[res] == 1 {
		heap.Pop(h)
		h.idx[res] = -1
	}
	h.cnt[res]--
	h.size--
	return res
}
func (h *Heap) peek() int {
	return h.data[0]
}
func (h *Heap) remove(x int) int {
	if h.cnt[x] == 1 {
		heap.Remove(h, h.idx[x])
		h.idx[x] = -1
	}
	h.size--
	h.cnt[x]--
	return x
}
```
