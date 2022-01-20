---
title: "1371. 每个元音包含偶数次的最长子字符串"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1371. 每个元音包含偶数次的最长子字符串](https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)

难度中等

给你一个字符串 `s` ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。

**示例 1：**

```
输入：s = "eleetminicoworoep"
输出：13
解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
```

**示例 2：**

```
输入：s = "leetcodeisgreat"
输出：5
解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。
```

**示例 3：**

```
输入：s = "bcbcbc"
输出：6
解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。
```

**提示：**

- `1 <= s.length <= 5 x 10^5`
- `s` 只包含小写英文字母。

函数签名：

```go
func findTheLongestSubstring(s string) int
```

## 分析
朴素做法是 `O(n^2)`的复杂度。可以考虑前缀和技巧，先遍历一遍得到从开头到每个位置的前缀字串各个元音字母出现的个数，然后枚举所有子串，可以用前缀和数组通过减法迅速计算出子串中元音字母的个数。但是枚举所有子串也是`O(n^2)`的复杂度。

可以进一步考虑，实际上不需要直到每个位置结束的子串字母的个数，只需要知道字母出现的个数奇偶性就行。如果 i 处 字母出现的奇偶性和 j 处的相同，那么它们之间的子串字母出现的次数就是偶数次。

将 5 个元音字母出现次数的奇偶视为一种状态，一共有 从 00000 到 11111 共 1 << 5 即 32 种状态，可以用一个整数变量 status 代表每种状态。最低的 5 位记录 5 个字母出现的奇偶性，0 代表出现偶数次，1 代表出现奇数次。

如果子串 [0，i] 与字串 [0,j] 状态`相同`，那么字串 [i+1,j] 的状态一定是 0，即5个元音字母出现的次数是偶数。

因此可以记录每个状态第一次出现的位置，此后再出现该状态时更新结果即可。
需要注意状态 0 首次出现的位置应该设定为 -1。

```go
func findTheLongestSubstring(s string) int {
	dic := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}
	status := 0
	pos := make([]int, 1<<len(dic))
	const invalidIndex = -2
	for i := range pos {
		pos[i] = invalidIndex
	}
	pos[0] = -1
	res := 0
	for i := range s {
		index, ok := dic[s[i]]
		if ok {
			status ^= 1 << index
		}
		if pos[status] == invalidIndex {
			pos[status] = i
		} else {
			res = max(res, i-pos[status])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

时间复杂度 `O(n)`， 空间复杂度 `O(1<<m)`，n 为 s 长度，m 是要考虑的字母个数，这个问题里是 5。

>  为了通用性，代码定义了一个 dic 字典，表示了需要包含的字母，并为其一一编了序号。这样如果题目改变，比如  s 为 unicode 字符串，但是需要找到所有英文小写字母出现次数为偶数的最长字串，因为英文小写字母只有 26 个，一个整数变量还是可以表示，题目解法相同，只需要修改 dic 内容即可。