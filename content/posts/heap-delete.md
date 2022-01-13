---
title: "支持删除任意元素的堆"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [Go, 数据结构]
---

堆可以在常数时间获知最值（即堆顶元素），对数时间插入元素、删除堆顶元素。

但有时候需要能迅速删除任意一个元素，比如 [480. 滑动窗口中位数](https://leetcode-cn.com/problems/sliding-window-median/) 这样的问题。

有没有办法让删除操作也在对数复杂度呢？

标准库提供了 Remove 方法，这个方法本身是对数级复杂度的，不过方法接受的是元素在堆里的索引，而不是元素本身（类似地还有个 Fix 方法）。要能迅速获知元素在堆里的索引，需要额外的空间做记录。详细可参考标准库优先队列相关示例。

下边介绍一个延迟删除的技巧。

## 延迟删除

在要删除一个元素时，先不急着从堆里真正删除，而是先把它记录下来。后边因各种操作如果待删除元素到达了堆顶，就可以在对数时间内把它真正从堆里删除了。

可以用一个哈希表或者另一个堆来存储所有待删除的元素。下边写一个使用哈希表存储待删除元素的实现，简单起见，假设堆里存储的都是 int 类型的元素：

```go
type Cmp func(int, int) bool

type Heap struct {
	slice []int
	cmp   Cmp
	// 缓存应该删除的元素，键为元素，值为应该删除的个数
	delMemo map[int]int
	// 因为待删除元素缓存，用一个额外的属性维护堆的真实大小
	size int
}

func NewHeap(cmp Cmp) *Heap { return &Heap{cmp: cmp, delMemo: map[int]int{}} }

func (h *Heap) Len() int           { return len(h.slice) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.slice[i], h.slice[j]) }
func (h *Heap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.slice[h.Len()-1]
	h.slice = h.slice[:h.Len()-1]
	return x
}

// prune 循环检查堆顶元素，如果在 delMemo 缓存中则删除
func (h *Heap) prune() {
	for len(h.slice) > 0 && h.delMemo[h.slice[0]] > 0 {
		h.delMemo[h.slice[0]]--
		heap.Pop(h)
	}
}
// push 向堆里添加一个元素
func (h *Heap) push(x int) {
	h.prune()
	h.size++
	heap.Push(h, x)
}
// pop 删除堆顶元素并返回其值
func (h *Heap) pop() int {
	h.size--
	res := heap.Pop(h).(int)
	h.prune()
	return res
}
// remove 在堆里删除元素 num
func (h *Heap) remove(num int) {
	h.delMemo[num]++
	h.size--
	h.prune()
}
// peek 获取堆顶元素，这里也可以调 prune 函数，不过因为 push、pop、和 remove 都调了，这里可以不调。
func (h *Heap) peek() int {
	return h.slice[0]
}
```