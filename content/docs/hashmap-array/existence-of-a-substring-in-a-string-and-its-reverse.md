---
title: 3083. 字符串及其反转中是否存在同一子字符串
date: 2024-12-26T17:51:21+08:00
---

## [3083. 字符串及其反转中是否存在同一子字符串](https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse) (Easy)

给你一个字符串 `s` ，请你判断字符串 `s` 是否存在一个长度为 `2` 的子字符串，在其反转后的字符串中也出现。

如果存在这样的子字符串，返回 `true`；如果不存在，返回 `false` 。

**示例 1：**

**输入：** s = "leetcode"

**输出：** true

**解释：** 子字符串 `"ee"` 的长度为 `2`，它也出现在 `reverse(s) == "edocteel"` 中。

**示例 2：**

**输入：** s = "abcba"

**输出：** true

**解释：** 所有长度为 `2` 的子字符串 `"ab"`、 `"bc"`、 `"cb"`、 `"ba"` 也都出现在 `reverse(s) == "abcba"` 中。

**示例 3：**

**输入：** s = "abcd"

**输出：** false

**解释：** 字符串 `s` 中不存在满足「在其反转后的字符串中也出现」且长度为 `2` 的子字符串。

**提示：**

- `1 <= s.length <= 100`
- 字符串 `s` 仅由小写英文字母组成。

## 分析

因为仅含小写英文字母，可以用一个 [26]int 数组记录长度为 2 的子串是否存在与 s 中，这样仅需要遍历一遍，且在遍历过程中判断逆序子串是否在 s 中。

```go
func isSubstringPresent(s string) bool {
	// memo[i] 代表以字母 i+'a' 开头的长度为 2 的子串是否在 s 中，如 memo[0] = b00000000000000000000000101， 表示 “aa" 和 “ac” 在 s 中
	memo := [26]int{}
	for i := 0; i < len(s)-1; i++ {
		x, y := s[i]-'a', s[i+1]-'a'
		memo[x] |= 1 << y
		if memo[y]&(1<<x) != 0 {
			return true
		}
	}
	return false
}

```
