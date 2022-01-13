---
title: "131. 分割回文串"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

难度中等

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是 **回文串** 。返回 `s` 所有可能的分割方案。

**回文串** 是正着读和反着读都一样的字符串。

**示例 1：**

```
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
```

**示例 2：**

```
输入：s = "a"
输出：[["a"]]
```

**提示：**

- `1 <= s.length <= 16`
- `s` 仅由小写英文字母组成

函数签名：

```go
func partition(s string) [][]string
```

## 分析

### 穷举回溯

假设 s 长度为n，则需要在 n 个字符的空档中做分割，需要保证分割后的所有子串为回文串。

代码框架如下：

```go
func partition(s string) [][]string {
	var res [][]string
	var cur []string
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for end := start+1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				cur = append(cur, s[start:end])
				backtrack(end)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return res
}
```

isPalindrome 函数不难实现，可以用双指针的方法在线性时间复杂度得出结果。但是在递归里边做这个线性操作，复杂度会比较高。可以事先单独计算出每个子串是否为回文串（用区间 dp 的方法会比较高效），在递归回溯里边就可以在常数时间得到结果。

```go
func partition(s string) [][]string {
	var res [][]string
	var cur []string
	isPalindrome := getPalindromeMemo(s)
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for end := start+1; end <= len(s); end++ {
			if isPalindrome[start][end-1] { // isPalindrome(s[start:end])
				cur = append(cur, s[start:end])
				backtrack(end)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func getPalindromeMemo(s string) [][]bool {
	n := len(s)
	// res[i][j] 表示 s[i:j+1] 是不是一个回文串
	res := make([][]bool, n)
	for i := range res {
		res[i] = make([]bool, n)
		for j := range res[i] {
			res[i][j] = true
		}
	}
	// 区间 dp
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			res[i][j] = s[i] == s[j] && res[i+1][j-1]
		}
	}
	return res
}
```

时间复杂度：O(n * 2^n)，其中 n 是字符串 s 的长度。在最坏情况下，s 包含 n 个完全相同的字符，因此它的任意一种划分方法都满足要求。而长度为 n 的字符串的划分方案数为 2^(n-1)，每一种划分方法需要 O(n) 的时间求出对应的划分结果并放入答案，因此总时间复杂度为 O(n * 2^n)。尽管动态规划预处理需要 O(n^2) 的时间，但在渐进意义下小于 O(n * 2^n)，因此可以忽略。

空间复杂度：O(n^2)，这里不计算返回答案占用的空间。动态规划需要使用的空间为 O(n^2)，而在回溯的过程中，使用 O(n) 的栈空间以及 O(n)的用来存储当前字符串分割方法的空间。由于 O(n) 在渐进意义下小于 O(n^2)，因此空间复杂度为 O(n^2)。

## [132. 分割回文串 II](https://leetcode-cn.com/problems/palindrome-partitioning-ii/)

难度困难

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是回文。

返回符合要求的 **最少分割次数** 。

**示例 1：**

```
输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
```

**示例 2：**

```
输入：s = "a"
输出：0
```

**示例 3：**

```
输入：s = "ab"
输出：1
```

**提示：**

- `1 <= s.length <= 2000`
- `s` 仅由小写英文字母组成

函数签名：

```go
func minCut(s string) int
```

## 分析

### 动态规划

定义长度 为 n 的动态规划数组 dp，dp[i] 代表将 s[:i+1] 切割成若干回文子串的最小次数，dp[i] 可以由比 i 小的 j 的值 dp [j] 推出：

```
[xxxjxxxxxi]

如果 xxxxxi 是个回文串，那么 dp[i] = dp[j]+1
当然 需要从 0 到 i-1 枚举 j，枚举过程中更新 dp[i]
```
同样可以用上个问题区间动态规划的方法，事先求出 isPalindrome 数组以减少时间复杂度。

```go
func minCut(s string) int {
	n := len(s)
	isPalindrome := getPalindromeMemo(s)
	// dp[i] 代表将 s[:i+1] 切割成若干回文子串的最小次数
	dp := make([]int, n)
	for i := range dp {
		if isPalindrome[0][i] { // s[:i+1] 已经是个回文串，最小切割次数为 0
			continue
		}
		dp[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if isPalindrome[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}
```

时空复杂度都是 O(n^2)。

