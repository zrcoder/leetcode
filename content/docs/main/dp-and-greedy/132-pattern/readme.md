---
title: "456. 132 模式"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [456. 132 模式](https://leetcode-cn.com/problems/132-pattern/)

难度中等

给你一个整数数组 `nums` ，数组中共有 `n` 个整数。**132 模式的子序列** 由三个整数 `nums[i]`、`nums[j]` 和 `nums[k]` 组成，并同时满足：`i < j < k` 和 `nums[i] < nums[k] < nums[j]` 。

如果 `nums` 中存在 **132 模式的子序列** ，返回 `true` ；否则，返回 `false` 。

**进阶：**很容易想到时间复杂度为 `O(n^2)` 的解决方案，你可以设计一个时间复杂度为 `O(n logn)` 或 `O(n)` 的解决方案吗？

**示例 1：**

```
输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。
```

**示例 2：**

```
输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
```

**示例 3：**

```
输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
```

**提示：**

- `n == nums.length`
- `1 <= n <= 104`
- `-109 <= nums[i] <= 109`

函数签名：

```go
func find132pattern(nums []int) bool
```

## 分析

最容易想到的是 O(n^3) 的解法，即对 i，j，k 做三层循环。时间复杂度较高。

### 两层循环

实际复杂度可以降低一维：从左向右枚举 j，对于当前 j，左半部分中最小的值 minLeft 可以作为 i，这个 minLeft 可以在枚举 j 的过程中更新，当 j 处的数字比 minLeft 至少大 2 的时候，固定 j，遍历其右侧找到 k，使得 k 处数字介于 minLeft 和 nums[j] 之间时返回 true。

```go
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMin := nums[0]
	for j := 1; j < len(nums)-1; j++ {
		leftMin = min(leftMin, nums[j])
		if nums[j] <= leftMin+1 {
			continue
		}
		for k := j + 1; k < len(nums); k++ {
			if nums[k] > leftMin && nums[k] < nums[j] {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
```

时间复杂度 O(n^2)，空间复杂度 O(1)。

## 使用 BST 优化

枚举 j 的过程中，对于当前的 j，怎么能迅速判读其右侧是否存在介于 leftMin 和 nums[j] 之间的数字？有没有这样的数据结构？

这个数据结构需要在一开始存储除去前两个之外的所有元素，随着 j 不断增大，每次需要在这个数据结构里删除 j 处元素。

C++里有 multiset，其 insert、multiset 和 erase 方法都是对数级的复杂度。Jave 里也有类似的数据结构如 TreeMap。

Go 标准库没有这样的数据结构。需要手写一个。

首先这个数据结构要是一个二叉搜索树（BST），如果树足够平衡，增删元素和查找比当前数字大的最小元素这样的操作都是对数级复杂度。平衡的 BST 有 b 树、红黑树等，但是编码较复杂，一个比较简单的实现是 Treap，即树堆，除了各个节点的值满足 BST 特性，每个节点额外附加一个权重属性。新建节点时权重值随机，在增删元素时通过旋转等操作保证整棵树所有节点的权重满足堆的性质。这样保证了树基本平衡。

```go
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	leftMin := nums[0]
	rights := &Treap{}
	for _, v := range nums[2:] {
		rights.Put(v)
	}

	for j := 1; j < n-1; j++ {
		if nums[j] > leftMin+1 {
			ub := rights.UpperBound(leftMin)
			if ub != nil && ub.val < nums[j] {
				return true
			}
		} else if nums[j] < leftMin {
			leftMin = nums[j]
		}
		rights.Delete(nums[j+1])
	}
	return false
}
```

时间复杂度 O(nlogn)，空间复杂度 O(n)。

Treap 的原理和实现详见 [Treap](../../go/data-structural/treap.md)
### 使用单调栈、前缀最小值优化

遍历所有元素，根据 132 模式，对于当前元素 x，需要判断是否能作为 2 类型的元素。

这就需要在左侧找到一个比其大的元素 y， 再在 y 的左侧找到一个比  x、y 都小的元素 z。

```
---z---y---x---
其中 z < x < y
```

贪心策略：**找 y 只需要找 x 左侧距离最近的，找 z 只需要找 y 左侧中最小的**。

为了迅速找 y，可以用单调栈；为了迅速找到 z，可以用前缀最小值数组。

```go
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	// prefixMin(i) 记录 i 左侧最小的元素
	prefixMin := getPrefixMin(nums)
	// leftMax(i) 记录 i 左侧第一个大于 nums[i] 的元素索引
	leftMax := getLeftMax(nums)
	for k, v := range nums {
		j := leftMax[k]
		if j == -1 {
			continue
		}
		iValue := prefixMin[j]
		if iValue < v {
			return true
		}
	}
	return false
}

// 返回每个元素左侧比其大的最近元素索引
func getLeftMax(nums []int) []int {
	res := make([]int, len(nums))
	var stack []int
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

// 返回每个元素左侧最小元素
func getPrefixMin(nums []int) []int {
	res := make([]int, len(nums))
	min := math.MaxInt64
	for i, v := range nums {
		res[i] = min
		if v < min {
			min = v
		}
	}
	return res
}
```

时空复杂度都是 O(n)。

## 扩展

上边使用单调栈和前缀最小值的解法遍历了数组三次，为了得到 prefixMin 和 leftMax 分别单独做了一次遍历，最优又遍历一次得到结果。实际上这三次遍历合并，最终只需要一次遍历。

```go
func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	prefixMin := make([]int, n)
	min := math.MaxInt64
	stack := make([]int, 0, n)
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= v {
			stack = stack[:len(stack)-1]
		}
		// stack[len(stack)-1] 是距离 v 最近的元素的索引,假设值为 j, prefixMin(j) 是 j 左侧最小元素
		if len(stack) > 0 && prefixMin[stack[len(stack)-1]] < v {
			return true
		}
		stack = append(stack, i)
		if v < min {
			min = v
		}
		prefixMin[i] = min
	}
	return false
}
```

时空复杂度都是 O(n)。