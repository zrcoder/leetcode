---
title: 16. 最接近的三数之和
date: 2023-07-10T14:04:29+08:00
---

## [16. 最接近的三数之和](https://leetcode.cn/problems/3sum-closest) (Medium)

给你一个长度为 `n` 的整数数组 `nums` 和 一个目标值 `target`。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

**示例 1：**

```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

```

**示例 2：**

```
输入：nums = [0,0,0], target = 1
输出：0

```

**提示：**

- `3 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-10⁴ <= target <= 10⁴`

## 分析



```go
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MinInt32
	update := func(cur int) {
		if abs(res-target) > abs(cur-target) {
			res = cur
		}
	}
	for i, v := range nums {
		t := target - v
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == t {
				return target
			}
			if sum < t {
				update(sum + v)
				j++
			} else {
				update(sum + v)
				k--
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

```
