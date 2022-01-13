---
title: "567. 字符串的排列"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [滑动窗口]
---

## [567. 字符串的排列](https://leetcode-cn.com/problems/permutation-in-string/)

难度中等

给定两个字符串 **s1** 和 **s2**，写一个函数来判断 **s2** 是否包含 **s1** 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

**示例1:**

```
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
```



**示例2:**

```
输入: s1= "ab" s2 = "eidboaoo"
输出: False
```



**注意：**

1. 输入的字符串只包含小写字母
2. 两个字符串的长度都在 [1, 10,000] 之间

函数签名：

```go
func checkInclusion(s1 string, s2 string) bool
```

## 分析

### 定长滑动窗口

只需统计 s1 里各个字符的个数，再用一个长度固定为 s1 的长度的滑动窗口，统计窗口里的字符及个数，与之前对 s1 的统计比较，完全相等即返回 true，字符串里只有 26 个小写字母，统计、比较字符个数用一个数据就行。

```go
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	memo1 := [26]int{}
	memo2 := [26]int{}
	for i, v := range s1 {
		memo1[v-'a']++
		memo2[s2[i]-'a']++
	}
	if memo2 == memo1 {
		return true
	}
	for i := m; i < len(s2); i++ {
		memo2[s2[i-m]-'a']--
		memo2[s2[i]-'a']++
		if memo2 == memo1 {
			return true
		}
	}
	return false
}
```

时间复杂度就是 O(Σ+m+(n-m)*Σ)。空间复杂度常数级 O(Σ) 。其中 Σ 是字符集大小，这里是 26。

如果字符集变大，复杂度就比较高了。

### 优化后的变长滑动窗口

使用一个变长的滑动窗口，首先要保证窗口里的字符都在 s1 中，其次要尽量保证窗口里的字符及其个数与 s1 中相同，即窗口里的字符是 s1 所有字符的一个子集。

算法：

还是先用一个数组 memo 统计一下 s1 里各个字符及个数。

在 s2 里用左右双指针维护一个窗口。

起初，左右边界都在 s2 开始，每次右边界右移一步，对于新字符 ch ：

1）如果不存在于 s1 中，那么直接将左右指针都移动到下一个位置；

2）如果存在，那么让 **ch 在 memo 里的个数减去 1**，这样可能导致个数变成负数，那么需要**向右移动左指针，且在 memo 里将左指针对应的字符个数加 1，直到 ch 在 memo 里的个数变成 0**。

首先可以发现，2) 中右移左指针会满足 1)的要求。

其次，2）的规则保证了窗口里边的字符是 s1 的一个子集。且可以考虑在满足子集的前提下，如果窗口的长度正好是 s1 的长度，那么这个窗口里的字符集实际上就和 s1 的字符集相等。

```go
func checkInclusion(s1, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	memo := [26]int{}
	for _, ch := range s1 {
		memo[ch-'a']++
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		memo[x]--
		for memo[x] < 0 {
			memo[s2[left]-'a']++
			left++
		}
		if right-left+1 == m {
			return true
		}
	}
	return false
}
```

时间复杂度 O(Σ+m+n)，基本消除了字符集大小的影响（在一开始初始化 memo 数组的时候也是 Σ 的复杂度，但是这个可以忽略，比解法一里循环里边每次比较一次要优秀很多）。

空间复杂度 O(Σ)。

