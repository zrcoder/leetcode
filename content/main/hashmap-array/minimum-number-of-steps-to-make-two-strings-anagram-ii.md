---
title: 2186. 使两字符串互为字母异位词的最少步骤数
date: 2023-11-12T00:06:34+08:00
---

## [2186. 使两字符串互为字母异位词的最少步骤数](https://leetcode.cn/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii) (Medium)

给你两个字符串 `s` 和 `t` 。在一步操作中，你可以给 `s` 或者 `t` 追加 **任一字符** 。

返回使 `s` 和 `t` 互为 **字母异位词** 所需的最少步骤数。

**字母异位词** 指字母相同但是顺序不同（或者相同）的字符串。

**示例 1：**

```
输入：s = "leetcode", t = "coats"
输出：7
解释：
- 执行 2 步操作，将 "as" 追加到 s = "leetcode" 中，得到 s = "leetcodeas" 。
- 执行 5 步操作，将 "leede" 追加到 t = "coats" 中，得到 t = "coatsleede" 。
"leetcodeas" 和 "coatsleede" 互为字母异位词。
总共用去 2 + 5 = 7 步。
可以证明，无法用少于 7 步操作使这两个字符串互为字母异位词。
```

**示例 2：**

```
输入：s = "night", t = "thing"
输出：0
解释：给出的字符串已经互为字母异位词。因此，不需要任何进一步操作。

```

**提示：**

- `1 <= s.length, t.length <= 2 * 10⁵`
- `s` 和 `t` 由小写英文字符组成

## 分析



```go
func minSteps(s string, t string) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range t {
		cnt[c-'a']--
	}
	res := 0
	for _, v := range cnt {
		res += abs(v)
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
