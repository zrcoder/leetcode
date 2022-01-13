---
title: "870. 优势洗牌"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [870. 优势洗牌](https://leetcode-cn.com/problems/advantage-shuffle)

难度中等

给定两个大小相等的数组 `A` 和 `B`，A 相对于 B 的*优势*可以用满足 `A[i] > B[i]` 的索引 `i` 的数目来描述。

返回 `A` 的**任意**排列，使其相对于 `B` 的优势最大化。

**示例 1：**

```
输入：A = [2,7,11,15], B = [1,10,4,11]
输出：[2,11,7,15]
```

**示例 2：**

```
输入：A = [12,24,8,32], B = [13,25,32,11]
输出：[24,32,8,12]
```

**提示：**

1. `1 <= A.length = B.length <= 10000`
2. `0 <= A[i] <= 10^9`
3. `0 <= B[i] <= 10^9`

函数签名：

```go
func advantageCount(A []int, B []int) []int
```

## 分析

想到两种贪心策略。

## 借助 BST

对于 B 里的元素 x，需要在 A 中查找 y，比 x 大且要最小，如果 y 不存在（即没有比 x 大的元素），那么需要查找 A 中最小的元素 z，y/z 就是 x 对应的答案，当然遍历处理 B 中 x 找到答案后要在 A 中删除对应的元素。

在一棵平衡的 BST 中添加、删除、查找元素的复杂度都是对数级，在 BST 中查找比给定元素大的最小值比较好实现，也是对数级的复杂度。

一般的 BST 实现可能因为插入特定数据变得非常不平衡，甚至退化成一条链表，增删查的复杂度会是线性的。平衡的BST 有红黑树、AVL树等，标准库并没有提供，实现较复杂，手写困难。

可以手写 Treap，这是一棵较为平衡的 BST，实现起来相对简单。

```go
func advantageCount(A []int, B []int) []int {
	n := len(A)
	if n == 0 || len(B) != n {
		return nil
	}
	bst := &Treap{}
	for _, v := range A {
		bst.Put(v)
	}
	res := make([]int, n)
	for i, v := range B {
		node := bst.UpperBound(v)
		if node == nil {
			min := bst.GetMin()
			bst.Delete(min)
			res[i] = min
		} else {
			bst.Delete(node.val)
			res[i] = node.val
		}
	}
	return res
}
```

时间复杂度 O(nlogn)，空间复杂度(O(n))，其中 n 是给定数组大小。

Treap 的原理和实现详见 [Treap](../../go/data-structural/treap.md)

### 排序+双指针

可以从大到小遍历 B 中元素，对于当前元素 x，在 A 中查找比它大的元素 y，可以贪心地找当前 A 中剩余元素里最大的那个，不过有可能不存在，这时候只需要找到 A 中剩余元素里最小的那个 z即可，每次找到结果后别忘了从 A 中删除。

实际可以先把 B 降序排序，为了后边生成结果，数据要记录原始索引。

再把 A 升序排序，这样可以用双指针技巧迅速得到其中的最大最小值。起初 lo 指针指向头、hi 指针指向尾，每次找到结果后移动指针达到“删除”元素的目的。

```go
func advantageCount(A []int, B []int) []int {
	n := len(A)
	if n == 0 || len(B) != n {
		return nil
	}
	// A 升序、B降序，B 排序后保留原始索引
	sort.Ints(A)
	sortedB := make([][]int, n)
	for i, v := range B {
		sortedB[i] = []int{i, v}
	}
	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i][1] > sortedB[j][1]
	})
	res := make([]int, n)
	lo, hi := 0, n-1
	for _, v := range sortedB {
		index, val := v[0], v[1]
		if val >= A[hi] {
			res[index] = A[lo]
			lo++
		} else {
			res[index] = A[hi]
			hi--
		}
	}
	return res
}
```
时间复杂度 O(nlogn)，空间复杂度(O(n))，其中 n 是给定数组大小。
复杂度同上边用 BST 的解法。
