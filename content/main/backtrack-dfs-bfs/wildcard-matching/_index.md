---
title: "通配符/模式匹配"
date: 2022-04-08T08:57:01+08:00
math: true
---

从最简单的44题开始。

## [44. 通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/description/ "https://leetcode-cn.com/problems/wildcard-matching/description/")

| Category   | Difficulty    | Likes | Dislikes |
| ---------- | ------------- | ----- | -------- |
| algorithms | Hard (33.10%) | 854   | -        |

给定一个字符串 (`s`) 和一个字符模式 (`p`) ，实现一个支持 `'?'` 和 `'*'` 的通配符匹配。

```
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
```

两个字符串**完全匹配**才算匹配成功。

**说明:**

- `s` 可能为空，且只包含从 `a-z` 的小写字母。
- `p` 可能为空，且只包含从 `a-z` 的小写字母，以及字符 `?` 和 `*`。

**示例 1:**

```
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
```

**示例 2:**

```
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
```

**示例 3:**

```
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
```

**示例 4:**

```
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
```

**示例 5:**

```
输入:
s = "acdcb"
p = "a*c?b"
输出: false
```

## 分析

非常容易写出递归解法：

```go
func isMatch(s string, p string) bool {
    if len(p) == 0 {
        return len(s) == 0
    }
    if len(s) == 0 {
        return strings.Count(p, "*") == len(p)
    }
    if p[0] == '?' || p[0] == s[0] {
        return isMatch(s[1:], p[1:])
    }
    if p[0] == '*' {
        return isMatch(s[1:], p) || isMatch(s, p[1:])
    }
    return false
}
```

但因为有重复计算，复杂度太高，会超时。加上备忘录改造下。

首先用一个内部函数改造上边代码：

```go
func isMatch(s string, p string) bool {
    var help func(s, p string) bool
    help = func(s, p string) bool {
        if len(p) == 0 {
            return len(s) == 0
        }
        if len(s) == 0 {
            return strings.Count(p, "*") == len(p)
        }
        if p[0] == '?' || p[0] == s[0] {
            return help(s[1:], p[1:])
        }
        if p[0] == '*' {
            return help(s[1:], p) || help(s, p[1:])
        }
        return false
    }
    return help(s, p)
}
```

加上备忘录：

```go
func isMatch(s string, p string) bool {
    type Set = map[string]bool
    memo := map[string]Set{}

    var help func(s, p string) bool
    help = func(s, p string) bool {
        if len(p) == 0 {
            return len(s) == 0
        }
        if len(s) == 0 {
            return strings.Count(p, "*") == len(p)
        }
        if memo[s] == nil {
            memo[s] = Set{}
        }
        set := memo[s]
        if res, ok := set[p]; ok {
            return res
        }
        if p[0] == '?' || p[0] == s[0] {
            res := help(s[1:], p[1:])
            memo[s][p] = res
            return res
        }
        if p[0] == '*' {
            res := help(s[1:], p) || help(s, p[1:])
            memo[s][p] = res
            return res
        }
        memo[s][p] = false
        return false
    }
    return help(s, p)
}
```

用map嵌套map做备忘录有点低效，可以修改help函数入参为两个整数i， j，分别代表当前检查到了 s 的 索引 i处， p 的索引 j 处，这样备忘录可以用二维数组。

```go
func isMatch(s string, p string) bool {
    memo := make([][]int, len(s))
    for i := range memo {
        memo[i] = make([]int, len(p))
    }

    markMemo := func(i, j int, ok bool) {
        memo[i][j] = 1
        if !ok {
            memo[i][j] = -1
        }
    }

    var help func(i, j int) bool
    help = func(i, j int) bool {
        if j == len(p) {
            return i == len(s)
        }
        if i == len(s) {
            return strings.Count(p[j:], "*") == len(p[j:])
        }
        if memo[i][j] != 0 {
            return memo[i][j] == 1
        }
        if p[j] == '?' || p[j] == s[i] {
            res := help(i+1, j+1)
            markMemo(i, j, res)
            return res
        }
        if p[j] == '*' {
            res := help(i+1, j) || help(i, j+1)
            markMemo(i, j, res)
            return res
        }
        markMemo(i, j, false)
        return false
    }
    return help(0, 0)
}
```

假设 s 和 p 的长度分别为 m 和 n，加了备忘录的时空复杂度都是 $O(m\times{n})$ 。

也可以改成自底向上的动态规划解法：

```go
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
	// dp[x][y] 代表前缀 s[:x] 和 p[:y] 是否匹配， x，y分别是前缀长度
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true

    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || i > 0 && dp[i-1][j]
            } else {
                dp[i][j] = i > 0 && (p[j-1] == '?' || p[j-1] == s[i-1]) && dp[i-1][j-1]
             }
        }
    }

    return dp[m][n]
}
```

复杂度同上。

## [10. 正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/description/ "https://leetcode-cn.com/problems/regular-expression-matching/description/")

| Category   | Difficulty    | Likes | Dislikes |
| ---------- | ------------- | ----- | -------- |
| algorithms | Hard (31.59%) | 2887  | -        |

给你一个字符串 `s` 和一个字符规律 `p`，请你来实现一个支持 `'.'` 和 `'*'` 的正则表达式匹配。

- `'.'` 匹配任意单个字符
- `'*'` 匹配零个或多个前面的那一个元素

所谓匹配，是要涵盖 **整个** 字符串 `s`的，而不是部分字符串。

 

**示例 1：**

```
输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
```

**示例 2:**

```
输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
```

**示例 3：**

```
输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
```

**提示：**

- `1 <= s.length <= 20`
- `1 <= p.length <= 30`
- `s` 只包含从 `a-z` 的小写字母。
- `p` 只包含从 `a-z` 的小写字母，以及字符 `.` 和 `*`。
- 保证每次出现字符 `*` 时，前面都匹配到有效的字符

## 分析

和 44 题类似，但是 `*` 的规则更复杂，题目示例没有给出，其实应该补充下：如 `s="abc", p="a*bc"` 是应该匹配上的，p里的星号可以匹配p中其前边的字符。

```go
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }

    var help func(i, j int) bool
    help = func(i, j int) bool {
        if j == len(p) {
            return i == len(s)
        }
        if i == len(s) {
            return canMathEmpty(p[j:])
        }
        if memo[i][j] != 0 {
            return memo[i][j] == 1
        }
        firstMath := s[i] == p[j] || p[j] == '.'
        if j < len(p)-1 && p[j+1] == '*' { // j 后一位是星号
            res := help(i, j+2) || firstMath && help(i+1, j)
            markMemo(memo, i, j, res)
            return res
        }
        res := firstMath && help(i+1, j+1)
        markMemo(memo, i, j, res)
        return res

    }
    return help(0, 0)
}

// p 是否能匹配一个空串
func canMathEmpty(p string) bool {
    if p[len(p)-1] != '*' {
        return false
    }
    for i := 0; i < len(p)-1; i++ {
        if p[i] != '*' && p[i+1] != '*' {
            return false
        }
    }
    return true
}

func markMemo(dp [][]int, i, j int, ok bool) {
    if ok {
        dp[i][j] = 1
    } else {
        dp[i][j] = -1
    }
}
```

复杂度同上。

同样可以修改为自底向上的动态规划解法

```go
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true

    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                dp[i][j] = j >= 2 && dp[i][j-2] ||
                    i > 0 && j >= 2 && (p[j-2] == '.' || p[j-2] == s[i-1]) && dp[i-1][j]
            } else {
                dp[i][j] = i > 0 && (p[j-1] == '.' || p[j-1] == s[i-1]) && dp[i-1][j-1]
            }
        }
    }

    return dp[m][n]
}
```
