---
title: "找到中间态，避免穷举"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [869. 重新排序得到 2 的幂](https://leetcode-cn.com/problems/reordered-power-of-2/)

难度中等

给定正整数 `N` ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。

如果我们可以通过上述方式得到 2 的幂，返回 `true`；否则，返回 `false`。

**示例 1：**

```
输入：1
输出：true
```

**示例 2：**

```
输入：10
输出：false
```

**示例 3：**

```
输入：16
输出：true
```

**示例 4：**

```
输入：24
输出：false
```

**示例 5：**

```
输入：46
输出：true
```



**提示：**

1. `1 <= N <= 10^9`

函数签名：

```go
func reorderedPowerOf2(N int) bool
```

## 分析

如果真的去穷举 N 重新排列各位数组的情况，时间复杂度会比较高。

十进制数字 N 的每一位由 0~9 这些数字组成，可以统计所有位中，0~9这些数字的个数，用一个长度为 10 的数组记录结果即可。

其次，2 的幂从二进制看，就是 1 左移若干位得到的数字，从输入限制看，最大不会超过 32 整数。这样可以查看 1 << 0 到 1 << 31 这32个数字，对每个数字，同样计算十进制角度看所有位中 0~9这些数字的个数。如果发现统计出的数组和对 N 统计得到的数组相等，就意味着 N 重排各位后能得到 2 的幂。

```go
const base = 10

func reorderedPowerOf2(N int) bool {
	cnt := getCnt(N)
	for i := 0; i < 32; i++ {
		if cnt == getCnt(1<<i) {
			return true
		}
	}
	return false
}

func getCnt(n int) [base]int {
	res := [base]int{}
	for n > 0 {
		res[n%base]++
		n /= base
	}
	return res
}
```

时空复杂度都可以看作常数级。getCnt函数的复杂度和数字 N 的十进制位数线性相关，实际不会超过 10（最大的32位整数 2147483647 共10位）

## [966. 元音拼写检查器](https://leetcode-cn.com/problems/vowel-spellchecker/)

难度中等

在给定单词列表 `wordlist` 的情况下，我们希望实现一个拼写检查器，将查询单词转换为正确的单词。

对于给定的查询单词 `query`，拼写检查器将会处理两类拼写错误：

- 大小写：如果查询匹配单词列表中的某个单词（

  不区分大小写

  ），则返回的正确单词与单词列表中的大小写相同。

    - 例如：`wordlist = ["yellow"]`, `query = "YellOw"`: `correct = "yellow"`
    - 例如：`wordlist = ["Yellow"]`, `query = "yellow"`: `correct = "Yellow"`
    - 例如：`wordlist = ["yellow"]`, `query = "yellow"`: `correct = "yellow"`

- 元音错误：如果在将查询单词中的元音（‘a’、‘e’、‘i’、‘o’、‘u’）分别替换为任何元音后，能与单词列表中的单词匹配（

  不区分大小写

  ），则返回的正确单词与单词列表中的匹配项大小写相同。

    - 例如：`wordlist = ["YellOw"]`, `query = "yollow"`: `correct = "YellOw"`
    - 例如：`wordlist = ["YellOw"]`, `query = "yeellow"`: `correct = ""` （无匹配项）
    - 例如：`wordlist = ["YellOw"]`, `query = "yllw"`: `correct = ""` （无匹配项）

此外，拼写检查器还按照以下优先级规则操作：

- 当查询完全匹配单词列表中的某个单词（**区分大小写**）时，应返回相同的单词。
- 当查询匹配到大小写问题的单词时，您应该返回单词列表中的第一个这样的匹配项。
- 当查询匹配到元音错误的单词时，您应该返回单词列表中的第一个这样的匹配项。
- 如果该查询在单词列表中没有匹配项，则应返回空字符串。

给出一些查询 `queries`，返回一个单词列表 `answer`，其中 `answer[i]` 是由查询 `query = queries[i]` 得到的正确单词。

**示例：**

```
输入：wordlist = ["KiTe","kite","hare","Hare"], queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
输出：["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
```

**提示：**

1. `1 <= wordlist.length <= 5000`
2. `1 <= queries.length <= 5000`
3. `1 <= wordlist[i].length <= 7`
4. `1 <= queries[i].length <= 7`
5. `wordlist` 和 `queries` 中的所有字符串仅由**英文**字母组成。

函数签名：

```go
func spellchecker(wordlist []string, queries []string) []string
```

## 分析

首先，如果不理会大小写和元音的条件，只看每次查询的单词是否在给定的单词列表中，这非常简单，可以先把给定的单词列表处理成一个集合，这样对于每次查询都能在常数时间给出结果。

其次，考虑不区分大小写的条件。对于每次查询的单词，先将所有大写字母转成小写字母，然后看给定的单词列表中所有字母都转成小写后是否存在相同的单词，如果存在即返回，如果有有多个，返回列表第一个符合条件的单词。这样的话可以用一个哈希表记录给定单词列表忽略大小写的情况：键为单词全部字母转成小写后的单词，值为原始单词，为了满足后边查询时返回“第一个”，在确定这个哈希表的时候，每次发现键如果不存在就添加键值对，否则忽略。

最后，考虑元音字母修改的情况。如果穷举所有元音字母替换后的情况会非常麻烦。同样可以借鉴上边对大小写条件的处理。另用一个哈希表，对于一个单词，先将全部字母转成小写，再把所有原因字母用一个特定符号如 `*` 表示，这样得到的带星号字符串可以作为哈希表的键，而值记录给定单词列表里第一个能得到这样键的单词的原始内容。在查找时的逻辑与建立这个哈希表类似。

```go
var dic map[string]bool
var dicLow map[string]string
var dicWow map[string]string

func spellchecker(wordlist []string, queries []string) []string {
	dic = make(map[string]bool, len(wordlist))
	dicLow = make(map[string]string, len(wordlist))
	dicWow = make(map[string]string, len(wordlist))
	for _, v := range wordlist {
		dic[v] = true
		low := strings.ToLower(v)
		if _, ok := dicLow[low]; !ok {
			dicLow[low] = v
		}
		wow := ignoreVowels(low)
		if _, ok := dicWow[wow]; !ok {
			dicWow[wow] = v
		}
	}
	res := make([]string, len(queries))
	for i := range res {
		res[i] = help(queries[i])
	}
	return res
}

func help(s string) string {
	if _, ok := dic[s]; ok {
		return s
	}
	low := strings.ToLower(s)
	if origin, ok := dicLow[low]; ok {
		return origin
	}
	wow := ignoreVowels(low)
	if origin, ok := dicWow[wow]; ok {
		return origin
	}
	return ""
}

func ignoreVowels(s string) string {
	res := []byte(s)
	for i := range res {
		if isVowel(res[i]) {
			res[i] = '*'
		}
	}
	return string(res)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
```

时空复杂度都是 O(m+n)，m 和 n 分别是给定单词列表和查询列表的大小。注意每个单词的长度限定在 7 以内，所以忽略大小写及元音字母的操作可以看作常数级复杂度。