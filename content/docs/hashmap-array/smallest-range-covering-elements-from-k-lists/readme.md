---
title: "632. 最小区间"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [堆, 贪心, 哈希表, 滑动窗口]
---

## [632. 最小区间](https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/)

难度困难

你有 `k` 个 **非递减排列** 的整数列表。找到一个 **最小** 区间，使得 `k` 个列表中的每个列表至少有一个数包含在其中。

我们定义如果 `b-a < d-c` 或者在 `b-a == d-c` 时 `a < c`，则区间 `[a,b]` 比 `[c,d]` 小。


**示例 1：**

```
输入：nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
输出：[20,24]
解释： 
列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
```

**示例 2：**

```
输入：nums = [[1,2,3],[1,2,3],[1,2,3]]
输出：[1,1]
```

**示例 3：**

```
输入：nums = [[10,10],[11,11]]
输出：[10,11]
```

**示例 4：**

```
输入：nums = [[10],[11]]
输出：[10,11]
```

**示例 5：**

```
输入：nums = [[1],[2],[3],[4],[5],[6],[7]]
输出：[1,7]
```

**提示：**

- `nums.length == k`
- `1 <= k <= 3500`
- `1 <= nums[i].length <= 50`
- `-105 <= nums[i][j] <= 105`
- `nums[i]` 按非递减顺序排列

函数签名：
```go
func smallestRange(nums [][]int) []int
```

## 分析
贪心策略。参考如下借助堆的方法和借助哈希表的滑动窗口解法。
## 解法一： 借助堆的贪心解法
问题可以转化为：从 k 个列表中各取一个数，使得这 k 个数中的最大值与最小值的差最小。  
每个列表都是有序的，可以用贪心策略来解决。以题目中给的例子为例，前边带点表示选中的数字：

```
.4  10  15  24  26
.0  9   12  20
.5  18  22  30
```

> max：5， min：0， abs：5-0=5

```
.4  10  15  24  26
0   .9  12  20
.5  18  22  30
```

> max：9， min：4， abs：9-4=5

```
4   .10 15  24  26
0   .9  12  20
.5  18  22  30
```

> max：10， min：5， abs：10-5=5

```
4   .10 15  24  26
0   .9  12  20
5   .18 22  30
```

> max：18， min：9， abs：18-9=9

```
4   .10 15  24  26
0   9   .12 20
5   .18 22  30
```

> max：18， min：10， abs：18-10=8

```
4   10  .15 24  26
0   9   .12 20
5   .18 22  30
```

> max：18， min：12， abs：18-12=6

```
4   10  .15 24  26
0   9   12  .20
5   .18 22  30
```

> max：20， min：15， abs：20-15=5

```
4   10  15  .24  26
0   9   12  .20
5   .18 22  30
```

> max：24， min：18， abs：24-18=6

```
4   10  15  .24  26
0   9   12  .20
5   18  .22 30
```

> max：24， min：20， abs：24-20=4

以上所有步骤中，abs最小的min和max即为所求

> 如果每次遍历k次找到k个元素的最小值和最大值，时间复杂度会比较大。
> 
> 可以借助一个大小为k的小顶堆h来记录每次尝试的k个数字，一个变量 curMax 记录k个数字里最大的元素。
> 
> 每次计算 curMax 和堆顶元素的差得到 abs，之后堆顶元素修改成其所在列表的下一个元素，直到某一个列表元素尝试完。

以上所有步骤中，abs最小的min和max即为所求

```go
// 记录每个列表当前尝试元素的索引
var pos []int

func smallestRange(nums [][]int) []int {
	size := len(nums)
	pos = make([]int, size)

	h := &Heap{s: make([]Item, 0, size)}
	curMax := math.MinInt32
	// 先处理每个列表首元素
	for i, v := range nums {
		heap.Push(h, Item{val: v[0], listIndex: i})
		curMax = max(curMax, v[0])
	}
	lo, hi, minAbs := math.MinInt32, math.MaxInt32, math.MaxInt32
	for {
		peek := h.s[0] // 堆顶
		abs := curMax - peek.val
		if abs < minAbs {
			minAbs = abs
			lo = peek.val
			hi = curMax
		}
		index := peek.listIndex
		curPos := pos[index] + 1
		if curPos == len(nums[index]) {
			break
		}
		pos[index] = curPos
		curNum := nums[index][curPos]
		h.s[0].val = curNum
		heap.Fix(h, 0)
		curMax = max(curMax, curNum)
	}
	return []int{lo, hi}
}

type Item struct {
	val, listIndex int // 记录每个数字及其所在列表的索引
}

type Heap struct {
	s []Item
}

func (h *Heap) Len() int           { return len(h.s) }
func (h *Heap) Swap(i, j int)      { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap) Less(i, j int) bool { return h.s[i].val < h.s[j].val }
func (h *Heap) Push(x interface{}) { h.s = append(h.s, x.(Item)) }
func (h *Heap) Pop() interface{} {
	x := h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]
	return x
}
```

假设所有数字共`n个`，则时间复杂度为`O(n*lgk)`。

空间复杂度`O(k)`，堆的大小。

## 解法二：借助哈希表的滑动窗口解法
先统计出所有数字里的最小值 xMin 和最大值 xMax， [xMin, xMax]区间包含了所有数字，答案不会比这个区间更大  
接下来可以用滑动窗口的方式遍历[xMin, xMax]区间来尝试  
窗口左右边界一开始都为xMin，增加right指针，使得对于窗口[left, right]， k 个列表中的每个列表至少有一个数包含在其中；则left， right可能为一个答案  
但这时候可以缩小窗口，增加左边界， 一直维持k 个列表中的每个列表至少有一个数包含在窗口中的性质，直到不再满足这个限制后，停止增加左边界，开始增加右边界

> 为了迅速判断是否 k 个列表中的每个列表至少有一个数包含在窗口中，可以事先用一个哈希表统计每个数字都在哪些列表出现

```go
func smallestRange1(nums [][]int) []int {
	size := len(nums)
	indices := map[int][]int{}                 // 记录每个数字所在列表的索引
	xMin, xMax := math.MaxInt32, math.MinInt32 // 分别记录所有数字中的最小和最大元素
	for i, v := range nums {
		for _, x := range v {
			indices[x] = append(indices[x], i)
			xMin = min(xMin, x)
			xMax = max(xMax, x)
		}
	}
	freq := make([]int, size)
	inside := 0
	start, end := xMin, xMax
	for left, right := xMin, xMin; right < xMax; right++ {
		if len(indices[right]) == 0 { // k个列表里都不含right这个数字
			continue
		}
		for _, index := range indices[right] {
			freq[index]++
			if freq[index] == 1 { // right这个数字在列表index里出现一次了，即列表index至少有一个元素包含在窗口里了
				inside++
			}
		}
		for ; inside == size; left++ {
			if right-left < end-start {
				start, end = left, right
			}
			for _, index := range indices[left] {
				freq[index]--
				if freq[index] == 0 {
					inside--
				}
			}
		}
	}
	return []int{start, end}
}
```

假设所有数字共`n`个， 最大数字与最小数字差为`abs`，时间复杂度`O(max(n, k*abs))`。