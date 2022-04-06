---
title: "710. 黑名单中的随机数"
date: 2022-04-06T18:55:24+08:00
---

## [710. 黑名单中的随机数](https://leetcode-cn.com/problems/random-pick-with-blacklist/description/ "https://leetcode-cn.com/problems/random-pick-with-blacklist/description/")

| Category   | Difficulty    | Likes | Dislikes |
| ---------- | ------------- | ----- | -------- |
| algorithms | Hard (37.18%) | 90    | -        |

给定一个整数 `n` 和一个 **无重复** 黑名单整数数组 `blacklist` 。设计一种算法，从 `[0, n - 1]` 范围内的任意整数中选取一个 **未加入** 黑名单 `blacklist` 的整数。任何在上述范围内且不在黑名单 `blacklist` 中的整数都应该有 **同等的可能性** 被返回。

优化你的算法，使它最小化调用语言 **内置** 随机函数的次数。

实现 `Solution` 类:

- `Solution(int n, int[] blacklist)` 初始化整数 `n` 和被加入黑名单 `blacklist` 的整数
- `int pick()` 返回一个范围为 `[0, n - 1]` 且不在黑名单 `blacklist` 中的随机整数

**示例 1：**

```
输入
["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
[[7, [2, 3, 5]], [], [], [], [], [], [], []]
输出
[null, 0, 4, 1, 6, 1, 0, 4]

解释
Solution solution = new Solution(7, [2, 3, 5]);
solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
                 // 0、1、4和6的返回概率必须相等(即概率为1/4)。
solution.pick(); // 返回 4
solution.pick(); // 返回 1
solution.pick(); // 返回 6
solution.pick(); // 返回 1
solution.pick(); // 返回 0
solution.pick(); // 返回 4
```

**提示:**

- `1 <= n <= 10^9`
- `0 <= blacklist.length <- min(10^5, n - 1)`
- `0 <= blacklist[i] < n`
- `blacklist` 中所有值都 **不同**
-  `pick` 最多被调用 `2 * 104` 次

## 分析

比较容易想到的解法是在初始化时生成白名单，那么获取百名单里但随机值会非常简单，但这样容易超内存——从题目约束能看出黑名单的大小远小于白名单的大小。

假设黑名单里共有m个数字，那么每次求随机值，可以生成一个小于n-m的随机值，如果在黑名单里，就映射到大于等于 n-m（当然同时保证小于n）的白名单里的一个数。这要怎么做呢？在初始化里搞定。对于黑名单里的数字，不小于 n-m的不管，小于n-m的映射到大于等于n-m且不在黑名单里的一个数字。

```go
type Solution struct {
	mmapping map[int]int
	limit    int
}

func Constructor(n int, blacklist []int) Solution {
	limit := n - len(blacklist)

	set := make(map[int]bool, len(blacklist))
	for _, v := range blacklist {
		set[v] = true
	}
	m := map[int]int{}
	k := limit
	for _, v := range blacklist {
		if v >= limit {
			continue
		}
		for set[k] {
			k++
		}
		m[v] = k
		k++
	}

	return Solution{
		limit:    limit,
		mmapping: m,
	}

}

func (s *Solution) Pick() int {
	x := rand.Intn(s.limit)
	y, ok := s.mmapping[x]
	if ok {
		return y
	}
	return x
}
```

时间复杂度：主要是初始化的耗时，为`O(m)`；空间复杂度：`O(m)`，`m` 是黑名单的大小。


