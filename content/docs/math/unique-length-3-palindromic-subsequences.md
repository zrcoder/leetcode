---
title: 1930. 长度为 3 的不同回文子序列
date: 2023-11-14T18:56:47+08:00
---

## [1930. 长度为 3 的不同回文子序列](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences) (Medium)

给你一个字符串 `s` ，返回 `s` 中 **长度为 3** 的 **不同回文子序列** 的个数。

即便存在多种方法来构建相同的子序列，但相同的子序列只计数一次。

**回文** 是正着读和反着读一样的字符串。

**子序列** 是由原字符串删除其中部分字符（也可以不删除）且不改变剩余字符之间相对顺序形成的一个新字符串。

- 例如， `"ace"` 是 `"abcde"` 的一个子序列。

**示例 1：**

```
输入：s = "aabca"
输出：3
解释：长度为 3 的 3 个回文子序列分别是：
- "aba" ("aabca" 的子序列)
- "aaa" ("aabca" 的子序列)
- "aca" ("aabca" 的子序列)

```

**示例 2：**

```
输入：s = "adc"
输出：0
解释："adc" 不存在长度为 3 的回文子序列。

```

**示例 3：**

```
输入：s = "bbcbaba"
输出：4
解释：长度为 3 的 4 个回文子序列分别是：
- "bbb" ("bbcbaba" 的子序列)
- "bcb" ("bbcbaba" 的子序列)
- "bab" ("bbcbaba" 的子序列)
- "aba" ("bbcbaba" 的子序列)

```

**提示：**

- `3 <= s.length <= 10⁵`
- `s` 仅由小写英文字母组成

## 分析

仅需要统计每个字母在s中第一次出现和最后一次出现的位置，这样两个位置之间字母种类数即3字符回文子序列的个数。

比如对于字母 a，假设在 s 中是这样： xaxxxaxx （其中x代表任意字母），第一个和最后一个a出现的位置分别是1和5，
那么仅需要统计 s[2:5] 里不同字母的个数，就得到了以a开头结尾的3字符回文子序列的个数。

> 即使这两个 a 中间的字母又出现了a，这个统计策略也没有问题。

最后累加所有字母开头结尾的3字符回文子序列个数即可。

这样时间复杂是 O(nE)，其中E是字符种类数，这里是26，n 是 s 的长度；空间复杂度 O(E)。

```go
func countPalindromicSubsequence(s string) int {
	const letters = 26

	first, last := [letters]int{}, [letters]int{}
	for i := len(s) - 1; i >= 0; i-- {
		first[s[i]-'a'] = i
	}
	for i, c := range s {
		last[c-'a'] = i
	}

	countUniq := func(s string) int {
		set := 0
		for _, c := range s {
			set |= 1 << (c - 'a')
		}
		return bits.OnesCount(uint(set))
	}

	sum := 0
	for i := 0; i < letters; i++ {
		f, l := first[i], last[i]
		if f+1 < l {
			sum += countUniq(s[f+1 : l])
		}
	}
	return sum
}

```
