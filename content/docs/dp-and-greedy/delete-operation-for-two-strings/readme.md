---
title: "583. 两个字符串的删除操作"
date: 2021-05-13T09:13:33+08:00
weight: 50
tags: [动态规划]
---

## [583. 两个字符串的删除操作](https://leetcode-cn.com/problems/delete-operation-for-two-strings/)

难度中等

给定两个单词 *word1* 和 *word2*，找到使得 *word1* 和 *word2* 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

**示例：**

```
输入: "sea", "eat"
输出: 2
解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
```

**提示：**

1. 给定单词的长度不超过500。
2. 给定单词中的字符只含有小写字母。

函数签名：

```go
func minDistance(word1 string, word2 string) int
```

## 分析

一开始没有太好的思路，先这样考虑：如果知道了两个单词的某两个字串对应的最小操作数，是不是能逐渐推出考了整个 word1 和 word2 的最终结果？这样还是有些复杂，限定字串为两个单词的前缀呢？这样的话就可以逐渐增长这两个前缀，长前缀的结果应该能由短的结果推出来。

具体来说，定义二维动态规划数组 `dp`，`dp[i][j]`表示考虑 `word1[0:i]` 和 `word[2][0:j]`，能使之相等的最小操作步数。会发现这样的状态转移方程：

前缀末尾字符相等即`word1[i-1] == word2[j-1]`时，`dp[i][j] =dp[i-1][j-1]`;

末尾字符不相等时，可以删除两个末尾字符，最小操作数为：`dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1`。

而初始状态也比较好确定：`dp[0][j] = j, dp[i][0] = i`。

这就是一个二维的动态规划。

```go
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		w1 := word1[i-1]
		for j := 1; j <= n; j++ {
			w2 := word2[j-1]
			if w1 == w2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1])+1
			}
		}
	}
	return dp[m][n]
}
```

时空复杂度都是 `O(m*n)`，即两个单词长度乘积。值得注意的是`dp`数组可以优化为一维。