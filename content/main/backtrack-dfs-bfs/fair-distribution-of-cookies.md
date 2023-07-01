---
title: 2305. 公平分发饼干
date: 2023-07-01T21:14:14+08:00
---

## [2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies) (Medium)

给你一个整数数组 `cookies` ，其中 `cookies[i]` 表示在第 `i` 个零食包中的饼干数量。另给你一个整数 `k` 表示等待分发零食包的孩子数量， **所有** 零食包都需要分发。在同一个零食包中的所有饼干都必须分发给同一个孩子，不能分开。

分发的 **不公平程度** 定义为单个孩子在分发过程中能够获得饼干的最大总数。

返回所有分发的最小不公平程度。

**示例 1：**

```
输入：cookies = [8,15,10,20,8], k = 2
输出：31
解释：一种最优方案是 [8,15,8] 和 [10,20] 。
- 第 1 个孩子分到 [8,15,8] ，总计 8 + 15 + 8 = 31 块饼干。
- 第 2 个孩子分到 [10,20] ，总计 10 + 20 = 30 块饼干。
分发的不公平程度为 max(31,30) = 31 。
可以证明不存在不公平程度小于 31 的分发方案。

```

**示例 2：**

```
输入：cookies = [6,1,3,2,2,4,1,2], k = 3
输出：7
解释：一种最优方案是 [6,1]、[3,2,2] 和 [4,1,2] 。
- 第 1 个孩子分到 [6,1] ，总计 6 + 1 = 7 块饼干。
- 第 2 个孩子分到 [3,2,2] ，总计 3 + 2 + 2 = 7 块饼干。
- 第 3 个孩子分到 [4,1,2] ，总计 4 + 1 + 2 = 7 块饼干。
分发的不公平程度为 max(7,7,7) = 7 。
可以证明不存在不公平程度小于 7 的分发方案。

```

**提示：**

- `2 <= cookies.length <= 8`
- `1 <= cookies[i] <= 10⁵`
- `2 <= k <= cookies.length`

## 分析


k 个人，分 n 袋饼干，要求尽量公平，使得每个人得到的饼干数量最接近。

回溯，数据量很小，数组长度限制为 8。

用长度为 k 的 数组 记录划分的结果。

时间复杂度是 O(k*2^n).


```go
func distributeCookies(cookies []int, k int) int {
	res := math.MaxInt
	// 记录k个人获得的饼干数量，可以为0
	tmp := make([]int, k)

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(cookies) {
			res = min(res, max(tmp))
			return
		}
		for j := range tmp {
			tmp[j] += cookies[i]
			backtrack(i + 1)
			tmp[j] -= cookies[i]
		}
	}

	backtrack(0)
	return res
}

func max(s []int) int {
	res := math.MinInt
	for _, v := range s {
		if v > res {
			res = v
		}
	}
	return res
}
func min(s ...int) int {
	res := math.MaxInt
	for _, v := range s {
		if v < res {
			res = v
		}
	}
	return res
}

```
