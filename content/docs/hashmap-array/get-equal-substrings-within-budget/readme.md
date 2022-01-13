---
title: "1208. 尽可能使字符串相等"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [滑动窗口, 前缀和]
---

## [1208. 尽可能使字符串相等](https://leetcode-cn.com/problems/get-equal-substrings-within-budget/)

难度中等

给你两个长度相同的字符串，`s` 和 `t`。

将 `s` 中的第 `i` 个字符变到 `t` 中的第 `i` 个字符需要 `|s[i] - t[i]|` 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是 `maxCost`。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 `s` 的子字符串转化为它在 `t` 中对应的子字符串，则返回可以转化的最大长度。

如果 `s` 中没有子字符串可以转化成 `t` 中对应的子字符串，则返回 `0`。

**示例 1：**

```
输入：s = "abcd", t = "bcdf", cost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。
```

**示例 2：**

```
输入：s = "abcd", t = "cdef", cost = 3
输出：1
解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。
```

**示例 3：**

```
输入：s = "abcd", t = "acde", cost = 0
输出：1
解释：你无法作出任何改动，所以最大长度为 1。
```

 **提示：**

- `1 <= s.length, t.length <= 10^5`
- `0 <= maxCost <= 10^6`
- `s` 和 `t` 都只含小写英文字母。

函数签名：

```go
func equalSubstring(s string, t string, maxCost int) int
```

## 分析

首先可以有一个 O(n^2) 的朴素解法，穷举所有子串，查看每个子串是否能在预算之内完成修改。复杂度较高，实现代码略。

如果使用区间 DP，时间复杂度也没有优化，反而增大了空间复杂度。

### 滑动窗口

怎么优化？可以用滑动窗口来枚举子串：每次查看窗口里的耗费是否在预算之内即可。

滑动窗口要么右边界右移要么左边界右移，就像一条尺蠖毛毛虫，一会头向前走一段，一会尾向前走一段。这条毛毛虫能在线性时间复杂度搞定问题，因为每个元素最多被毛毛虫头尾指针各访问一次~

```go
func equalSubstring(s string, t string, maxCost int) int {
	i, j := 0, 0 // 窗口的左右边界
	cost := 0    // 窗口的消耗
	res := 0
	for j < len(s) {
		curCost := abs(int(s[j]) - int(t[j]))
		if curCost > maxCost { // 包含当前位置的子串都不可能修改，窗口直接跑到下一个位置
			cost = 0
			j++
			i = j
			continue
		}
		for cost+curCost > maxCost { // 缩短左边界使消耗在预算之内
			cost -= abs(int(s[i]) - int(t[i]))
			i++
		}
		cost += curCost
		res = max(res, j-i+1)
		j++
	}
	return res
}
```

时间复杂度 `O(n)`，空间复杂度 `O(1)`。

### 前缀和+二分

还有一个思路，复杂度没有上边滑动窗口优秀，但是也值得一提。

先用一个数组统计每个位置的花费，并求出这个数组的前缀和。那么只需要在这个前缀和数组里截取一段，用末尾的值减去开始位置前一位的值就能迅速得到这一段的花费，和预算相比就可以更新结果。

不过在前缀和数组里取一段，用两层循环还是太耗时了。可以只用一层循环，每次固定结尾，用二分法找到开始位置即可。

```go
func equalSubstring(s string, t string, maxCost int) int {
	cost := make([]int, len(s)+1)
	for i := range s {
		cost[i+1] = cost[i] + abs(int(s[i])-int(t[i]))
	}

	res := 0
	for end := range s {
		start := sort.SearchInts(cost[:end+1], cost[end+1]-maxCost)
		res = max(res, end-start+1)
	}
	return res
}
```

时间复杂度 `O(nlogn)`，空间复杂度 `O(n)`。