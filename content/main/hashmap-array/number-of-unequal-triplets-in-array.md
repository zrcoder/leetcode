---
title: 2475. 数组中不等三元组的数目
---

## [2475. 数组中不等三元组的数目](https://leetcode.cn/problems/number-of-unequal-triplets-in-array) (Easy)

给你一个下标从 **0** 开始的正整数数组 `nums` 。请你找出并统计满足下述条件的三元组 `(i, j, k)` 的数目：

- `0 <= i < j < k < nums.length`
- `nums[i]`、 `nums[j]` 和 `nums[k]` **两两不同** 。

    - 换句话说： `nums[i] != nums[j]`、 `nums[i] != nums[k]` 且 `nums[j] != nums[k]` 。

返回满足上述条件三元组的数目。

**示例 1：**

```
输入：nums = [4,4,2,4,3]
输出：3
解释：下面列出的三元组均满足题目条件：
- (0, 2, 4) 因为 4 != 2 != 3
- (1, 2, 4) 因为 4 != 2 != 3
- (2, 3, 4) 因为 2 != 4 != 3
共计 3 个三元组，返回 3 。
注意 (2, 0, 4) 不是有效的三元组，因为 2 > 0 。

```

**示例 2：**

```
输入：nums = [1,1,1,1,1]
输出：0
解释：不存在满足条件的三元组，所以返回 0 。

```

**提示：**

- `3 <= nums.length <= 100`
- `1 <= nums[i] <= 1000`

## 分析

### 朴素解法

时间复杂度 O(n^3), 空间复杂度 O(1).

```go
func unequalTriplets(nums []int) int {
	res := 0
	for i, vi := range nums {
		for j, vj := range nums[:i] {
			for _, vk := range nums[:j] {
				if vi != vj && vi != vk && vj != vk {
					res++
				}
			}
		}
	}
	return res
}
```

### 排序

时间复杂度 O(nlogn), 空间复杂度，O(1).
因为结果与元素顺序无关，可以排序让相同元素聚在一起，对于当前一堆相同的元素 [i, j), 其对结果的贡献为 left *（j-i）* right

```go
func unequalTriplets(nums []int) int {
	res := 0
	sort.Ints(nums)
	n := len(nums)
	for i, j := 0, 0; i < n; i = j {
		for j < n && nums[j] == nums[i] {
			j++
		}
		res += i * (j - i) * (n - j)
	}
	return res
}
```

### 哈希表统计元素数量

类似排序解法. 时空复杂度均为 O(n).

```go
func unequalTriplets(nums []int) int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	n := len(nums)
	pre := 0
	res := 0
	for _, v := range cnt {
		res += pre * v * (n - pre - v)
		pre += v
	}
	return res
}

```
