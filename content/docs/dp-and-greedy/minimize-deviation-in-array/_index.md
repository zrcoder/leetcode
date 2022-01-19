---
title: "1675. 数组的最小偏移量"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1675. 数组的最小偏移量](https://leetcode-cn.com/problems/minimize-deviation-in-array/)

难度困难

给你一个由 `n` 个正整数组成的数组 `nums` 。

你可以对数组的任意元素执行任意次数的两类操作：

- 如果元素是偶数，除以 2

  - 例如，如果数组是 `[1,2,3,4]` ，那么你可以对最后一个元素执行此操作，使其变成 `[1,2,3,**2**]`

- 如果元素是奇数，乘上 2

  - 例如，如果数组是 `[1,2,3,4]` ，那么你可以对第一个元素执行此操作，使其变成 `[**2**,2,3,4]`

数组的 **偏移量** 是数组中任意两个元素之间的 **最大差值** 。

返回数组在执行某些操作之后可以拥有的 **最小偏移量** 。

 

**示例 1：**

```
输入：nums = [1,2,3,4]
输出：1
解释：你可以将数组转换为 [1,2,3,2]，然后转换成 [2,2,3,2]，偏移量是 3 - 2 = 1
```

**示例 2：**

```
输入：nums = [4,1,5,20,3]
输出：3
解释：两次操作后，你可以将数组转换为 [4,2,5,5,3]，偏移量是 5 - 2 = 3
```

**示例 3：**

```
输入：nums = [2,10,8]
输出：3
```

 

**提示：**

- `n == nums.length`
- `2 <= n <= 105`
- `1 <= nums[i] <= 109`

## 分析

如果有一个数据结构，能迅速获知集合中最大值和最小值就好了。这样只需要看最大值是否偶数，是的话考虑除以 2，只要除以2 之后会使结果更小就除，之后更新下结果；同时看最小值是不是奇数，是的话考虑乘以2，类似对最大值的处理。重复执行以上比较并不断更新结果。

很可惜没有这样的数据结构。可以尝试用两个堆， 一个大顶堆一个小顶堆来执行上边的操作。不过在更新大顶堆堆顶后要遍历小顶堆找到对应的元素更新小顶堆，时间复杂度也非常高。

还需要再寻找规律，找其他的方法。

**如果一开始所有数字都是偶数会怎么样？**

显然可以把最大的偶数除以2，再更新结果。只需要一个大顶堆维护所有的数字，一个变量维护最小值，每次堆顶数字如果使偶数，就除以2，更新堆，在堆顶为偶数的条件下一直重复操作。

有两个点需要注意，首先更新的过程中最小值可能改变，其次更新后会出现奇数，但是这些奇数是没有必要乘以2的，这样相当于开了倒车。

最终当堆顶为奇数时，结束操作，返回堆顶和最小值的差值即可。

**回到这个问题，只需要一开始把所有奇数乘以 2 ，就划归成了刚才的问题**，想一想，这样并不影响最终结果的准确性。

```go
type Heap struct {
	s []int
}

func NewHeapWithCap(n int) *Heap {
	return &Heap{s: make([]int, 0, n)}
}
func (h *Heap) Less(i, j int) bool { return h.s[i] > h.s[j] }
func (h *Heap) Len() int           { return len(h.s) }
func (h *Heap) Swap(i, j int)      { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap) Push(x interface{}) { h.s = append(h.s, x.(int)) }
func (h *Heap) Pop() interface{} {
	res := h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	return res
}
func (h *Heap) push(x int) { heap.Push(h, x) }
func (h *Heap) pop() int   { return heap.Pop(h).(int) }
func (h *Heap) peek() int  { return h.s[0] }
func (h *Heap) updatePeek(v int) {
	h.s[0] = v
	heap.Fix(h, 0)
}

func minimumDeviation(nums []int) int {
	h := NewHeapWithCap(len(nums))
	minVal := math.MaxInt64
	for _, v := range nums {
		if v%2 == 1 {
			v *= 2
		}
		minVal = min(minVal, v)
		h.push(v)
	}
	res := h.peek() - minVal
	for h.peek()%2 == 0 {
		tmp := h.peek() / 2
		h.updatePeek(tmp)
		minVal = min(minVal, tmp)
		res = min(res, h.peek()-minVal)
	}
	return res
}
```
