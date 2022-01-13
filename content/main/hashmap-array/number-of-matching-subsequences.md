---
title: "792. 匹配子序列的单词数"
date: 2022-12-14T12:22:14+08:00
---

## [792. 匹配子序列的单词数](https://leetcode.cn/problems/number-of-matching-subsequences/)

难度中等

给定字符串 `s` 和字符串数组 `words`, 返回  *`words[i]` 中是`s`的子序列的单词个数* 。

字符串的 **子序列** 是从原始字符串中生成的新字符串，可以从中删去一些字符(可以是none)，而不改变其余字符的相对顺序。

- 例如， `“ace”` 是 `“abcde”` 的子序列。

**示例 1:**

**输入:** s = "abcde", words = ["a","bb","acd","ace"]
**输出:** 3
**解释:** 有三个是 s 的子序列的单词: "a", "acd", "ace"。

**Example 2:**

**输入:** s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
**输出:** 2

**提示:**

- `1 <= s.length <= 5 * 104`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 50`
- `words[i]`和 s 都只由小写字母组成。

函数签名：

```go
func numMatchingSubseq(s string, words []string) int
```

## 分析

### 朴素解法（超时）

遍历每个单词，用双指针的方法确定是否为 s 的子序列。

```go
func numMatchingSubseq(s string, words []string) int {
    res := 0
    for _, w := range words {
        if isSubSeq(s, w) {
            res++
        }
    }
    return res
}

func isSubSeq(s, w string) bool {
    i := 0
    for j := range w {
        for i < len(s) && s[i] != w[j] {
            i++
        }
        if i == len(s) {
            return false
        }
        i++
    }
    return true
}
```

时间复杂度：`O(n*m)`，空间复杂度：`O(1)`。其中 n 是 s 的长度，m 是words中所有字符的总数。

### 二分搜索改进 isSubSeq 函数

可以预处理 s， 得到每种字符的索引列表，如字母 a 出现的索引，该索引有序。这样判断某个单词是否是 s 的子序列时可以用二分法。

```go
func numMatchingSubseq(s string, words []string) int {
    pos := make([][]int, 26) // 记录每个字符在 s 中的索引列表
    for i, c := range s {
        pos[c-'a'] = append(pos[c-'a'], i)
    }
    res := 0
    for _, w := range words {
        if isSubSeq(pos, w) {
            res++
        }
    }
    return res
}

func isSubSeq(pos [][]int, w string) bool {
    k := -1 // 记录 s 中已经被匹配过的索引
    for _, c := range w {
        idx := pos[c-'a']
        i := sort.SearchInts(idx, k+1) // 在 idx 中二分搜索找到第一个比 k 大的值
        if i == len(idx) { // 不存在
            return false
        }
        k = idx[i]
    }
    fmt.Println(w)
    return true
}
```

时间复杂度优化为：`O(m*logn)`，空间复杂度增加为：`O(n)`。

### 多指针

可以仅遍历一遍 s，对于s当前的字符，一次性标记 words 中 以其开头的单词，然后单词去除该前缀，如果某个单词变成空，说明匹配完了该单词。

如果用遍历 words 数组的方式来做一次性标记，那么复杂度和朴素解法一样。有没有办法优化？

因为都是小写字母，总共26个，可以维护每个字符开头的待匹配单词列表。

```go
func numMatchingSubseq(s string, words []string) int {
    memo := [26][]string{} // memo[x] 代表以字符 x 开头的待匹配的单词列表
    for _, w := range words {
        c := int(w[0]-'a')
        memo[c] = append(memo[c], w)
    }

    res := 0
    for _, c := range s {
        ws := memo[c-'a']
        memo[c-'a'] = nil
        for _, w := range ws {
            if len(w) == 1 {
                res++
            } else {
                w = w[1:]
                memo[w[0]-'a'] = append(memo[w[0]-'a'], w)
            }
        }
    }
    return res
}
```

复杂度 TODO
