---
title: "341. 扁平化嵌套列表迭代器"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [栈, 递归]
---

## [341. 扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)

难度中等

给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。

列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

**示例 1:**

```
输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
```

**示例 2:**

```
输入: [1,[4,[6]]]
输出: [1,4,6]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
```

已有的 NestedInteger 相关 API（实现不给出）：

```go
type NestedInteger struct {}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {}
```

待实现的迭代器：

```go
type NestedIterator struct {
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
}

func (it *NestedIterator) Next() int {
}

func (it *NestedIterator) HasNext() bool {
}

```

## 分析

这是个非常有意思的问题。首先 NestedInteger 就很有趣，可以想象一下它有什么应用场景，题目只是给出了这个类型的 API，但没有给出其实现；需要根据给定的 API 来完成迭代器相关实现。

先写迭代器，后边尝试把  NestedInteger 的所有 API 实现一把。

### 初始化压扁列表

给定一个 []*NestedInteger 类型，直接压扁成 []int，可以用递归，比较容易：

```go
type NestedIterator struct {
	nums []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var nums []int
	var flat func(nestedList []*NestedInteger)
	flat = func(nestedList []*NestedInteger) {
		for _, v := range nestedList {
			if v.IsInteger() {
				nums = append(nums, v.GetInteger())
			} else {
				flat(v.GetList())
			}
		}
	}
	flat(nestedList)
	return &NestedIterator{nums: nums}
}

func (it *NestedIterator) Next() int {
	res := it.nums[0]
	it.nums = it.nums[1:]
	return res
}

func (it *NestedIterator) HasNext() bool {
	return len(it.nums) > 0
}
```

初始化的时间复杂度是 O(n)， n 为所有数组个数。Next 和 HasNext 都是常数级复杂度。

空间复杂度是 O(n)。

### 借助栈

上边在初始化时就整个压扁嵌套列表的做法，其实不太符合迭代器的约束：

```
迭代器不应该直接存储所有数字，而应提供访问途径。
迭代有条件终止(如键值查找)时初始化方法的全局开销非必要。
初始化迭代器后，迭代过程中无法处理List中某数字值改变的场景。
```

考虑初始化不做复杂工作，在 Next 和 HasNext 方法里做。
```go
type NestedIterator struct {
	items []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{items: append([]*NestedInteger{}, nestedList...)}
}

func (it *NestedIterator) Next() int {
	// 保证调用 Next 之前会调用 HasNext
	res := it.items[0].GetInteger()
	it.items = it.items[1:]
	return res
}

func (it *NestedIterator) HasNext() bool {
	for len(it.items) > 0 && !it.items[0].IsInteger() {
		// 展开第一个元素，再追加后边的——比较耗费性能
		it.items = append(it.items[0].GetList(), it.items[1:]...)
	}
	return len(it.items) > 0
}
```

注意上边 `HasNext` 函数的注释，可以借助一个栈来优化，直接看代码：
```go
type NestedIterator struct {
	// 将列表视作一个队列，栈中直接存储该队列
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
	queue := it.stack[len(it.stack)-1]
	val := queue[0].GetInteger()
	it.stack[len(it.stack)-1] = queue[1:]
	return val
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		queue := it.stack[len(it.stack)-1]
		if len(queue) == 0 { // 当前队列为空，出栈
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		// 若队首元素为列表，则将其弹出队列并入栈
		it.stack[len(it.stack)-1] = queue[1:]
		it.stack = append(it.stack, nest.GetList())
	}
	return false
}
```

初始化、Next 方法时间复杂度都是常数级，HasNext 方法均摊复杂度也是常数级。

空间复杂度，主要在栈的大小，最坏 O(n)。

至此，迭代器写完，问题解决。

### 实现 NestedInteger

我还是很好奇 NestedInteger 各个 Api 的实现，尝试写了一下：

```go
type Any interface{}

type NestedInteger struct {
	// int 或 []*NestedInteger
	val Any
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {
	_, ok := ni.val.(int)
	return ok
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {
	if v, ok := ni.val.(int); ok {
		return v
	}
	return 0
}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {
	ni.val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {
	if ni.IsInteger() {
		ni.val = []*NestedInteger{ni, &elem}
	} else {
		list := ni.val.([]*NestedInteger)
		list = append(list, &elem)
		ni.val = list
	}
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {
	if ni.IsInteger() {
		return nil
	}
	return ni.val.([]*NestedInteger)
}
```

