---
title: "424. 替换后的最长重复字符"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [滑动窗口]
---

## [424. 替换后的最长重复字符](https://leetcode-cn.com/problems/longest-repeating-character-replacement/)

难度中等

给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 *k* 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

**注意：**字符串长度 和 *k* 不会超过 10^4。

 **示例 1：**

```
输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
```

**示例 2：**

```
输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
```

函数签名：

```go
func characterReplacement(s string, k int) int
```

## 分析

先尝试朴素解。穷举每个子串，看看每个子串在允许替换 k 个字符的情况下能否变成全部字符相同的串，可以通过统计得到最多字符的个数，这个个数加上 k 如果大于等于对应子串的长度，就可以用子串的长度更新答案。

```go
func characterReplacement(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < n-1; i++ {
		for j := n-1; j >= i; j-- {
			if check(s[i:j+1], k) {
				res = max(res, j-i+1)
				break
			}			
		}
	}
	return res
}

func check(sub string, k int) bool {
	memo := [26]int{}
	maxCnt := 0
	for _, v := range sub {
		memo[v-'A']++
		maxCnt = max(maxCnt, memo[v-'A'])
	}
	return maxCnt + k >= len(sub)
}
```

时间复杂度是`O(n^3)`。

### 滑动窗口

朴素解法超时，必须考虑一个复杂度更低的实现。

继承一部分朴素解法的思想，用两个指针来维护子串，而不是用内外循环的方法得到子串。

两个指针形成一个窗口，每次统计窗口中最多的元素个数，如果其加上 k 的值不小于窗口长度，则窗口长度就可以更新结果。

```go
func characterReplacement(s string, k int) int {
	n := len(s)
	res := 0
	cnt := make([]int, 26)
	for left, right := 0, 0; right < n; right++ {
		cnt[s[right]-'A']++
		maxCnt := calMax(cnt)
		if maxCnt+k >= right-left+1 {
			res = max(res, right-left+1)
		} else {
			cnt[s[left]-'A']--
			left++
		}
	}
	return res
}

func calMax(cnt []int) int {
	res := 0
	for _, v := range cnt {
		res = max(res, v)
	}
	return res
}
```

注意到用了一个 cnt 数组来记录窗口里边各个字母的个数；借助这个数组来遍历计算窗口中最多的字母的个数；因为字母只有 26 个，这么做的复杂度可以接受。

实际上这个地方还可以优化。

只需要加一个变量 maxCnt 来维护窗口里的最多字符数量，calMax 函数可以省去。在窗口右边界右移的时候更新 maxCnt， 但是在窗口左边界向右移动的地方，**不用更新** maxCnt —— 这样并不影响结果：

左边界向右移动一个位置的时候，maxCnt 或者不变，或者值减 1。
当左边界向右移动之前，如果有多种字符长度相等，左边界向右移动不改变 maxCnt 的值。例如 s = [AAABBB]、k = 2，左边界 A 移除以后，窗口内字符出现次数不变，依然为 3；
当左边界移除以后，使得此时 maxCnt 的值变小，接下来继续让右边界向右移动一格，有两种情况：① 右边界如果读到了刚才移出左边界的字符，恰好 maxCnt 的值被正确维护；② 由于最终要找的只是最长替换 k 次以后重复子串的长度，右边界如果读到了不是刚才移出左边界的字符，新的子串要想在符合题意的条件下变得更长，maxCnt 一定要比之前的值还更大，因此不会错过更优的解。

```go
func characterReplacement(s string, k int) int {
	n := len(s)
	res := 0
	cnt := make([]int, 26)
	maxCnt := 0
	for left, right := 0, 0; right < n; right++ {
		cnt[s[right]-'A']++
		maxCnt = max(maxCnt, cnt[s[right]-'A'])
		if maxCnt+k >= right-left+1 {
			res = max(res, right-left+1)
		} else {
			cnt[s[left]-'A']--
			// 这里无需更新 maxCnt，不影响结果
			left++
		}
	}
	return res
}
```

时间复杂度 `O(n)`，空间复杂度`O(26) = O(1)`。