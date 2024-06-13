---
title: "最长公共子序列/子数组"
date: 2022-11-30T18:33:41+08:00
---

## [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)

难度中等

给定两个字符串 `text1` 和 `text2`，返回这两个字符串的最长 **公共子序列** 的长度。如果不存在 **公共子序列** ，返回 `0` 。

一个字符串的 **子序列** 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

- 例如，`"ace"` 是 `"abcde"` 的子序列，但 `"aec"` 不是 `"abcde"` 的子序列。

两个字符串的 **公共子序列** 是这两个字符串所共同拥有的子序列。

**示例 1：**

**输入：** text1 = "abcde", text2 = "ace"
**输出：** 3  
**解释：** 最长公共子序列是 "ace" ，它的长度为 3 。

**示例 2：**

**输入：** text1 = "abc", text2 = "abc"
**输出：** 3
**解释：** 最长公共子序列是 "abc" ，它的长度为 3 。

**示例 3：**

**输入：** text1 = "abc", text2 = "def"
**输出：** 0
**解释：** 两个字符串没有公共子序列，返回 0 。

**提示：**

- `1 <= text1.length, text2.length <= 1000`
- `text1` 和 `text2` 仅由小写英文字符组成。

函数签名：

```go
func longestCommonSubsequence(text1 string, text2 string) int
```

## 分析

#### 动态规划

我们可以从较小规模的问题不断扩展来求解最终的问题。

假设 f(m, n) 表示仅考虑 text1前m个字符、text2前n个字符所得到的结果，显然，如果找到两个索引 `i`, `j` 使得 `text1[i] == text2[j]`，显然 `f(i+1, j+1) = f(i, j)+1`; 如果`text1[i] != text2[j]`, 则 `f(i+1, j+1) = max(f(i, j+1), f(i+1, j)`

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1]+1
                continue
            }
            if dp[i-1][j] > dp[i][j-1] {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[m][n]
}
```

时空复杂度都是`O(m*n)`。

## 扩展

### 扩展1: 怎么构造出那个最长公共子序列

如果要返回那个最长的公共子序列呢？

如[最长公共子序列(二)_牛客题霸_牛客网](https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11?tpId=295&tqId=991075&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)就是这样一个问题。

可以用一个额外的二维数组dir来记录动态规划过程中转移的路径，最后根据这个信息递归地逆向求出结果。参考解答如下：

```go
func LCS(s1 string, s2 string) string {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)  // dp[i][j] 表示仅考虑 s1[:i] 和 s2[:j] 能够得到的最长公共子序列的长度
	dir := make([][]int, m+1) // dir 路径数组，记录了dp过程中转移的方向，为了最后构造出结果
	for i := range dp {
		dp[i] = make([]int, n+1)
		dir[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				dir[i][j] = 1 // 转移路径1，状态(i-1, j-1) -> 状态(i,j)
				continue
			}
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
				dir[i][j] = 2 // 转移路径2, 状态(i-1, j) -> 状态(i,j)
			} else {
				dp[i][j] = dp[i][j-1]
				dir[i][j] = 3 // 转移路径3, 状态(i, j-1) -> 状态(i,j)
			}
		}
	}
	
	if dp[m][n] == 0 {
		return "-1"
	}
    
    res := make([]byte, dp[m][n])
    k := len(res)-1
    i, j := m, n
    for k >= 0 {
        switch dir[i][j] {
        case 1:
            res[k] = s1[i-1]
            k--
            i--
            j--
        case 2:
            i--
        case 3:
            j--
        }
    }
	return string(res)
}
```

### 扩展2: 如果把子序列改成子数组呢？

子序列可以不连续，但是子数组必须连续，如果问题中的子序列改成子数组，问题会变得更简单，只需要对dp数组的定义做一点微调，对应的转移方程页微调即可。

如 [最长公共子串_牛客题霸_牛客网](https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac?tpId=295&tqId=991150&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)

参考解答如下：

```go
func LCS(str1 string, str2 string) string {
    m, n := len(str1), len(str2)
    // dp[i][j] 表示仅考虑 str1[:i] 和 str2[:j]且公共子串 **必须以 str1[i-1] 结尾** 的最长公共子串的长度
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    max, end := 0, 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } // else dp[i][j] = 0，末尾不同，不满足dp定义

            if dp[i][j] > max {
                max = dp[i][j]
                end = i - 1
            }
        }
    }
    return str1[end-max+1 : end+1]
}
```
