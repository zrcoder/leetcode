---
title: "936. 戳印序列"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [936. 戳印序列](https://leetcode-cn.com/problems/stamping-the-sequence/)

难度困难

你想要用**小写字母**组成一个目标字符串 `target`。

开始的时候，序列由 `target.length` 个 `'?'` 记号组成。而你有一个小写字母印章 `stamp`。

在每个回合，你可以将印章放在序列上，并将序列中的每个字母替换为印章上的相应字母。你最多可以进行 `10 * target.length` 个回合。

举个例子，如果初始序列为 "?????"，而你的印章 `stamp` 是 `"abc"`，那么在第一回合，你可以得到 "abc??"、"?abc?"、"??abc"。（请注意，印章必须完全包含在序列的边界内才能盖下去。）

如果可以印出序列，那么返回一个数组，该数组由每个回合中被印下的最左边字母的索引组成。如果不能印出序列，就返回一个空数组。

例如，如果序列是 "ababc"，印章是 `"abc"`，那么我们就可以返回与操作 "?????" -> "abc??" -> "ababc" 相对应的答案 `[0, 2]`；

另外，如果可以印出序列，那么需要保证可以在 `10 * target.length` 个回合内完成。任何超过此数字的答案将不被接受。

**示例 1：**

```
输入：stamp = "abc", target = "ababc"
输出：[0,2]
（[1,0,2] 以及其他一些可能的结果也将作为答案被接受）
```

**示例 2：**

```
输入：stamp = "abca", target = "aabcaca"
输出：[3,0,1]
```

**提示：**

1. `1 <= stamp.length <= target.length <= 1000`
2. `stamp` 和 `target` 只包含小写字母。

函数签名：

```go
func movesToStamp(stamp string, target string) []int
```

## 分析

从“?????...?" 经过不断盖戳得到 target 字符串，正向考虑情况很复杂。

如果反过来考虑从 target 得到“?????...?"呢？这样会简单一点。

显然，target 里必须包含 stamp 子串（正向考虑的最后一步盖戳）。逆向操作的话，第一步可以先把找到的 stamp 子串全部替换成‘？’，后边怎么操作呢？

每次从左边开始找到第一个和 stamp “匹配”的子串，这里的“匹配”包含两种情况：完全相等，部分相等，其他地方是‘？’；找到后将其全部替换成‘？’。

如果最终能得到全为 ‘？’的结果就找到了一条路径。否则没法完成操作。

```go
func movesToStamp(stamp string, target string) []int {
	m, n := len(stamp), len(target)
	cur := []byte(target)
	dest := bytes.Repeat([]byte{'?'}, n)
	var res []int
	replaced := true
	for replaced {
		replaced = false
		for i := 0; i <= n-m; i++ {
			if bytes.Compare(cur[i:i+m], dest[i:i+m]) == 0 {
				continue
			}
			if match(cur[i:i+m], stamp) {
				res = append(res, i)
				copy(cur[i:i+m], dest[i:i+m])
				replaced = true
				break
			}
		}
		if bytes.Compare(cur, dest) == 0 {
			return reverse(res)
		}
	}
	return nil
}

func match(b []byte, s string) bool {
	for i := range b {
		if b[i] != '?' && b[i] != s[i] {
			return false
		}
	}
	return true
}

func reverse(s []int) []int {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
```

时间复杂度 `O(n*(n-m))`，空间复杂度 `O(m+n)`。其中 `n`、`m` 分别是 `target` 和 `stamp` 的长度。