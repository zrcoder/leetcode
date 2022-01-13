---
title: "堆的使用"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [Go, 数据结构]
---

关于堆和优先队列，标准库已经实现了核心部分，详见container/heap  
基本自定义集合（一般为一个切片）实现 heap.Interface 的 5 个函数，就可以用了。  
详见标准库两个 example 开头的测试文件。  
## 需求
假设我们要同时用到大顶堆和小顶堆，怎么办？简单起见，假设元素都是int。  
## 初步实现
参考标准库example_intheap_test.go,很容易写出以下代码：  
```go
type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
```
## 改进实现
重复代码很多，大顶堆和小顶堆只有比较逻辑不同，即那个 Less 函数，其他实现没有差别。  

有没有办法减少代码呢？

还真想到一个，在 Go 里，函数是一等公民，可以当一般变量传递使用，我们不妨给自定义 Heap 增加一个属性 cmp，类型是 func(int, int) bool，也就是上边的 Less 函数的类型  
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
代码一下子减少一半!使用时是这样：  
```go
	minHeap := &Heap{}
	minHeap.cmp = func(a, b int) bool {
		return a < b
	}
	maxHeap := &Heap{}
	maxHeap.cmp = func(a, b int) bool {
		return a > b
	}
```
要注意几点：  
首先在刚刚初始化完要立即赋予其 cmp，cmp 为 nil的话后边程序会崩溃给我们看~  
其次 cmp 后续不能被修改，不然堆的逻辑会混乱  

基本解决了问题，虽然不够完美~  

## 应用
- [数据流、滑动窗口的中位数](../../design/find-median-from-data-stream/readme.md)
  > 用两个堆来解决中位数的问题，非常妙，尤其可以看看滑动窗口那个问题的解法，实现了一个 remove、push 和 pop 都是对数级复杂度的堆。

## 更进一步
可以重构标准库的 heap 包，参考 https://github.com/zrcoder/dsGo/tree/master/heap
