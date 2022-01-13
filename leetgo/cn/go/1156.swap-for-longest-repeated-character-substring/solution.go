// Created by Bob at 2023/06/03 07:18
// https://leetcode.cn/problems/swap-for-longest-repeated-character-substring/

/*
1156. 单字符重复子串的最大长度 (Medium)
如果字符串中的所有字符都相同，那么这个字符串是单字符重复的字符串。

给你一个字符串
`text`，你只能交换其中两个字符一次或者什么都不做，然后得到一些单字符重复的子串。返回其中最长的子串的长度。

**示例 1：**

```
输入：text = "ababa"
输出：3

```

**示例 2：**

```
输入：text = "aaabaaa"
输出：6

```

**示例 3：**

```
输入：text = "aaabbaaa"
输出：4

```

**示例 4：**

```
输入：text = "aaaaa"
输出：5

```

**示例 5：**

```
输入：text = "abcdef"
输出：1

```

**提示：**

- `1 <= text.length <= 20000`
- `text` 仅由小写英文字母组成。
*/

package main

// @lc code=begin

func maxRepOpt1(text string) int {
	// 统计每个字符出现的个数
	cnt := [26]int{}
	for _, c := range text {
		cnt[c-'a']++
	}
	res := 0
	for i := 0; i < len(text); {
		// [i, j-1] 窗口内所有字符相同
		j := i + 1
		for ; j < len(text) && text[j] == text[i]; j++ {
		}
		// 局部更新 res
		if j-i < cnt[text[i]-'a'] && (i > 0 || j < len(text)) {
			res = max(res, j-i+1)
		}
		// j 处字符换成 text[i] 后, 可能和后边的连成一片.
		k := j + 1
		for ; k < len(text) && text[k] == text[i]; k++ {
		}
		res = max(res, min(k-i, cnt[text[i]-'a'])) // 注意 min 的作用, 比如对于 aabaa, 不能算成 5.
		i = j
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
