---
title: "220. 存在重复元素 III"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [滑动窗口, 桶排序, 哈希表]
---

## [220. 存在重复元素 III](https://leetcode-cn.com/problems/contains-duplicate-iii/)

难度中等

给你一个整数数组 `nums` 和两个整数 `k` 和 `t` 。请你判断是否存在 **两个不同下标** `i` 和 `j`，使得 `abs(nums[i] - nums[j]) <= t` ，同时又满足 `abs(i - j) <= k` 。

如果存在则返回 `true`，不存在返回 `false`。

**示例 1：**

```
输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
```

**示例 2：**

```
输入：nums = [1,0,1,1], k = 1, t = 2
输出：true
```

**示例 3：**

```
输入：nums = [1,5,9,1,5,9], k = 2, t = 3
输出：false
```

**提示：**

- `0 <= nums.length <= 2 * 104`
- `-231 <= nums[i] <= 231 - 1`
- `0 <= k <= 104`
- `0 <= t <= 231 - 1`

## 分析

### 滑动窗口

可以维护一个长度固定为 `k+1` 的滑动窗口，在滑动过程中计算窗口中最小差值。

用什么数据结构能迅速计算出最小差值呢？可以用 BST，查找大于等于 `x - t` 的最小的元素 `y`，如果 `y` 存在，且 `y <= x + t` 则找到了满足约束的一对元素；这个查找复杂的是对数级的，即 `O(logk)`，这样整个解决方案的复杂度是 `O(nlogk)`。

> 如果题目约束的是窗口中元素的最大差值，可以使用两个单调栈分别维护窗口里的最大值和最小值，计算最大差值的复杂的是常数级。

不过 Go 标准库没有好用的 API，手写红黑树等又非常麻烦，倒是可以手写个 Treap 树，但也不简单。代码略。

### 桶排序

受桶排序思路启发，可以这么做：

遍历所有元素，对于每个元素，尝试将其放进桶里。

> 定义桶的大小为 `t+1`，这样每个数字都能映射到一个桶里。
>
> 比如桶的大小为 3， 数字和桶的对应关系如下：
>
> ...|-3, -2, -1|0, 1, 2| 3, 4, 5|...
>
> ------ -1 ------ 0 ----- 1 ---------
>
> 代码为：
>
> ```go
> func getBucketId(num, wide int) int {
> 	if num >= 0 {
> 		return num / wide
> 	}
> 	return (num+1)/wide - 1
> }
> ```

如果发现对应的桶里已经有元素，则找到满足约束的一对元素；如果发现对应的桶前一个或后一个桶里也有元素，则可能找到满足约束的元素对。除了这三个桶，其他桶里的元素与当前元素的差值肯定大于 t。

实际不用真的排序，只需要在遍历时记录桶里的元素，且每个桶只要记录一个元素即可，用一个哈希表做这个记录就行。

```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	buckets := make(map[int]int, len(nums)) // 键为桶的 id，值为真实元素
	for i, x := range nums {
		id := getBucketId(x, t+1)
		if _, ok := buckets[id]; ok {
			return true
		}
		if y, ok := buckets[id-1]; ok && abs(x-y) <= t {
			return true
		}
		if y, ok := buckets[id+1]; ok && abs(x-y) <= t {
			return true
		}
		buckets[id] = x
		if i >= k {
			delete(buckets, getBucketId(nums[i-k], t+1))
		}
	}
	return false
}
```

时空复杂度都是 `O(n)`，哈希表中至多包含 `n` 个元素。
