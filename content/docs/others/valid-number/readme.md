---
title: "65. 有效数字"
date: 2021-05-15T18:26:04+08:00
weight: 50
tags: [分治, 状态机]
---

## [65. 有效数字](https://leetcode-cn.com/problems/valid-number/)

难度困难

**有效数字**（按顺序）可以分成以下几个部分：

1. 一个 **小数** 或者 **整数**
2. （可选）一个 `'e'` 或 `'E'` ，后面跟着一个 **整数**

**小数**（按顺序）可以分成以下几个部分：

1. （可选）一个符号字符（`'+'` 或 `'-'`）
2. 下述格式之一：
   1. 至少一位数字，后面跟着一个点 `'.'`
   2. 至少一位数字，后面跟着一个点 `'.'` ，后面再跟着至少一位数字
   3. 一个点 `'.'` ，后面跟着至少一位数字

**整数**（按顺序）可以分成以下几个部分：

1. （可选）一个符号字符（`'+'` 或 `'-'`）
2. 至少一位数字

部分有效数字列举如下：

- `["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]`

部分无效数字列举如下：

- `["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]`

给你一个字符串 `s` ，如果 `s` 是一个 **有效数字** ，请返回 `true` 。

 **示例 1：**

```
输入：s = "0"
输出：true
```

**示例 2：**

```
输入：s = "e"
输出：false
```

**示例 3：**

```
输入：s = "."
输出：false
```

**示例 4：**

```
输入：s = ".1"
输出：true
```

**提示：**

- `1 <= s.length <= 20`
- `s` 仅含英文字母（大写和小写），数字（`0-9`），加号 `'+'` ，减号 `'-'` ，或者点 `'.'` 。

函数签名：

```go
func isNumber(s string) bool
```

## 分析

### 分治

这个问题非常繁琐，这里主要探讨怎么明确地拆分处理。

首先可以删除所有空格，并将所有字母小写，这不会影响结果的正确且能简化问题。

其次，根据 'e' 这个字母，来分别处理底数和指数，如果 `e` 存在，则校验其左侧是不是一个有符号的浮点数，右侧是不是一个有符号的整数——这里把指数的问题消解拆分成了两个子问题。

> 当然如果不存在字母 `e` 的话只需要判断整个字符串是否代表一个有符号浮点数（或整数）。

判断字符串是否代表一个有符号的浮点数，先判断第一位，如果是正负号则忽略，判断剩余的部分是否是浮点数。判断字符串是否代表一个有符号整数，同样可以先消去首位的正负号影响。

现在的问题是判断字符是否代表浮点数和整数，比较容易。

```go
func isNumber(s string) bool {
  s = strings.ReplaceAll(s, " ", "")
  s = strings.ToLower(s)
  i := strings.Index(s, "e")
  if i == -1 {
    return isSignedFloat(s)
  }
  return isSignedFloat(s[:i]) && isSignedInt(s[i+1:])
}

func isSignedFloat(s string) bool {
  if len(s) == 0 {
    return false
  }
  if s[0] == '+' || s[0] == '-' {
    s = s[1:]
  }
  return isFloat(s)
}

func isSignedInt(s string) bool {
  if len(s) == 0 {
    return false
  }
  if s[0] == '+' || s[0] == '-' {
    s = s[1:]
  }
  return isInt(s)
}

func isFloat(s string) bool {
  if len(s) == 0 {
    return false
  }
  if strings.Count(s, ".") > 1 {
    return false
  }
  s = strings.ReplaceAll(s, ".", "")
  return isInt(s)
}

func isInt(s string) bool {
  if len(s) == 0 {
    return false
  }
  for _, v := range s {
    if v < '0' || v >'9' {
      return false
    }
  }
  return true
}
```

线性时间复杂度，常数空间复杂度。

### 状态机

这里引入了状态机的内容。考虑各种状态和状态转移还是非常头大，参考参考力扣官方题解[《确定有限状态自动机》](https://leetcode-cn.com/problems/valid-number/solution/you-xiao-shu-zi-by-leetcode-solution-298l)。