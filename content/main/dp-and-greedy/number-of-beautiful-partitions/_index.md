---
title: "2478. 完美分割的方案数"
date: 2022-11-23T17:36:14+08:00
---

## [2478. 完美分割的方案数](https://leetcode.cn/problems/number-of-beautiful-partitions/)

难度困难

给你一个字符串 `s` ，每个字符是数字 `'1'` 到 `'9'` ，再给你两个整数 `k` 和 `minLength` 。

如果对 `s` 的分割满足以下条件，那么我们认为它是一个 **完美** 分割：

- `s` 被分成 `k` 段互不相交的子字符串。
- 每个子字符串长度都 **至少** 为 `minLength` 。
- 每个子字符串的第一个字符都是一个 **质数** 数字，最后一个字符都是一个 **非质数** 数字。质数数字为 `'2'` ，`'3'` ，`'5'` 和 `'7'` ，剩下的都是非质数数字。

请你返回 `s` 的 **完美** 分割数目。由于答案可能很大，请返回答案对 `109 + 7` **取余** 后的结果。

一个 **子字符串** 是字符串中一段连续字符串序列。

**示例 1：**

**输入：** s = "23542185131", k = 3, minLength = 2
**输出：** 3
**解释：** 存在 3 种完美分割方案：
"2354 | 218 | 5131"
"2354 | 21851 | 31"
"2354218 | 51 | 31"

**示例 2：**

**输入：** s = "23542185131", k = 3, minLength = 3
**输出：** 1
**解释：** 存在一种完美分割方案："2354 | 218 | 5131" 。

**示例 3：**

**输入：** s = "3312958", k = 3, minLength = 1
**输出：** 1
**解释：** 存在一种完美分割方案："331 | 29 | 58" 。

**提示：**

- `1 <= k, minLength <= s.length <= 1000`
- `s` 每个字符都为数字 `'1'` 到 `'9'` 之一。

函数签名：

```go
func beautifulPartitions(s string, k int, minLength int) int
```

## 分析

### 动态规划

问题规模可以缩小。划分为 k 段子串，如果去除最后一段，算出剩余的字符串划分为 k-1 段的数量，那么可以推出将整个字符串划分为 k 个子串的数量。

假设用`dp[i][k]`表示将字符串前 i 个字符分割成k段，满足题意的解法的数量。

要求`dp[i][k]`就需要先找到一个更早的状态`dp[j][k-1]`，当然字符串`s[j+1:i]`需满足题意。这样`dp[i][k] += dp[j][k-1]`。遍历所有可能的j，就能求出答案。

```go
func beautifulPartitions(s string, k int, minLength int) int {
    const mod = 1e9+7
    n := len(s)

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
    }
    dp[0][0] = 1

    for i := minLength; i <= n; i++ {
        if isPrime(s[i-1]) {
            continue
        }
        for j := 0; j <= i-minLength; j++ {
            if !isPrime(s[j]) {
                continue
            }
            for k1 := 1; k1 <= k; k1++ {
                dp[i][k1] = (dp[i][k1]+dp[j][k1-1]) % mod
            }
        }
    }
    return dp[n][k]
}

func isPrime(b byte) bool {
    return b == '2' || b == '3' || b == '5' || b == '7'
}
```

时间复杂度：`O(k*n^2)`， 空间复杂度 `O(n*k)`。

### 优化后的动态规划

TODO



