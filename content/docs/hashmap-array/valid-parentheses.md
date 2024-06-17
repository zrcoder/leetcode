---
title: "有效的括号"
date: 2024-06-17T14:16:02+08:00
---

从较简单的问题开始：

## [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/)(简单)

给定一个只包括 `'('`，`')'`，`'{'`，`'}'`，`'['`，`']'` 的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

**示例 1：**

**输入：** s = "()"
**输出：** true

**示例 2：**

**输入：** s = "()[]{}"
**输出：** true

**示例 3：**

**输入：** s = "(]"
**输出：** false

**提示：**

- `1 <= s.length <= 104`
- `s` 仅由括号 `'()[]{}'` 组成

## 分析

从前向后遍历，如果遇到一个右括号，其前一个位置一定要是一个匹配的左括号，这样整个字符串才可能有效。

```go
func isValid(s string) bool {
    stk := make([]rune, 0, len(s)) // 用这个栈维护临时的左括号
    
    pair := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, ch := range s {
        switch ch {
        case ')', ']', '}': // 预期栈顶是匹配的左括号
            if len(stk) == 0 || stk[len(stk)-1] != pair[ch] {
                return false
            }
            stk = stk[:len(stk)-1] // 栈顶的左括号出栈
        default:
            stk = append(stk, ch)
        }
    }
    return len(stk) == 0
}
```

时空复杂度均为O(n)，其中 n 为字符串长度。

## [32. 最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/)（困难）

给你一个只包含 `'('` 和 `')'` 的字符串，找出最长有效（格式正确且连续）括号

子串的长度。

**示例 1：**

**输入：** s = "(()"
**输出：** 2
**解释：** 最长有效括号子串是 "()"

**示例 2：**

**输入：** s = ")()())"
**输出：** 4
**解释：** 最长有效括号子串是 "()()"

**示例 3：**

**输入：** s = ""
**输出：** 0

**提示：**

- `0 <= s.length <= 3 * 10^4`
- `s[i]` 为 `'('` 或 `')'`

## 分析

### 方法一：基于栈

继承上一个问题基于栈道解法。这次栈中存放左括号在原字符串中的索引，这样在遍历到一个右括号时，如果栈不空，则栈顶的左括号和这个右括号形成一个合法的括号对，我们记录这对括号的索引——这可以引入一个额外的bool数组 mark，记录 mark[i], mark[j] 为 true。

最后遍历这个bool数组，找到最长连续为true的子数组即可。

```go
func longestValidParentheses(s string) int {
	mark := make([]bool, len(s))    // 记录哪些位置是有效括号
	stk := make([]int, 0, len(s))
	for i, ch := range s {
		if ch == '(' {
			stk = append(stk, i)
		} else if len(stk) > 0 {
			j := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			// i, j 处是有效的括号
			mark[i] = true
			mark[j] = true
		}
	}

	// 以mark 为依据，统计连续的true出现的最长长度
	res := 0
	for i := 0; i < len(mark); {
		if !mark[i] {
			i++
			continue
		}
        // i 时连续段段开始，j是其末尾的下一个位置
		j := i+1
		for ; j < len(mark) && mark[j]; j++ {
		}
		res = max(res, j-i)
		i = j
	}
	return res
}
```

另一个解法：能不能不要mark数组，仅引入一个 stk 来解决呢？

可以这样做：始终保持栈底元素为当前已经遍历过的元素中`最后一个没有被匹配的右括号的下标`，这样的做法主要是考虑了边界条件的处理，栈里其他元素维护左括号的下标。

对于遇到的每个左括号 ，将它的下标入栈，
对于遇到的每个右括号，我们先弹出栈顶元素表示匹配了当前右括号：
> 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的`最后一个没有被匹配的右括号的下标`
>
> 如果栈不为空，当前右括号的下标减去栈顶元素即为`以该右括号为结尾的最长有效括号的长度`

从前往后遍历字符串并更新答案即可。

需要注意的是，为了保持统一，在一开始的时候往栈中放入一个值为 −1 的元素，来代表一开始`最后一个没有被匹配的右括号的下标` 。

```go
func longestValidParentheses(s string) int {
	stk := make([]int, 1, len(s)+1)
	stk[0] = -1
	res := 0
	for i, ch := range s {
		if ch == '(' {
			stk = append(stk, i)
		} else {
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				stk = append(stk, i)
			} else {
				j := stk[len(stk)-1]
				res = max(res, i-j)
			}
		}
	}
	return res
}
```

两种方法的时空复杂度均为O(n)，其中 n 为字符串长度。不过解法一开辟了两个数组，且除了遍历 s 本身，又遍历了 mark，比解法二稍微逊色一点。

### 动态规划

假设用一个dp数组，dp[i] 表示在 s 中以 s[i] 结尾的最长有效子字符串的长度。最终返回 max(dp)即可。

显然 s[i] == '(' 时， dp[i] = 0;

s[i] == ')' 时，考虑 s[i-1] 是左括号还是右括号，根据不同情况来更新 dp[i]。

```go
func longestValidParentheses(s string) int {
    dp := make([]int, len(s))
    res := 0
    for i, ch := range s {
        if ch == '(' { // dp[i] = 0
            continue
        }
        if i == 0 {
            continue
        }
        if s[i-1] == '(' { // ...()
            dp[i] = 2
            if i-2 >=0 {
                dp[i] += dp[i-2]
            }
        } else { // ...))
            j := i-dp[i-1]-1
            if j < 0 || s[j] != '(' {
                continue
            }
            dp[i] = 2+dp[i-1]
            if j-1 >= 0 {
                dp[i] += dp[j-1]
            }
        }
        res = max(res, dp[i])
    }
    return res
}
```

时空复杂度均为O(n)。

### 计数，常数空间复杂度解法

遍历的时候仅记录左括号和右括号的数量行不行？

可以这样做：从前向后遍历，遇到左括号则 left++，遇到右括号则 right++；之后如果right>left，则将 left 和 right 重置为 0 。

```go
func longestValidParentheses(s string) int {
    left, right := 0, 0
    res := 0
    for _, ch := range s {
        if ch == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            res = max(res, left*2)
        } else if right > left {
            left, right = 0, 0
        }
        
    }
    return res
}
```

但是这样还有漏洞，比如对于这个输入："(()"，用该算法跑出来是0，但显然应该是2才对；针对这个漏洞，可以按照类似的思路，从后向前再遍历一次。

最终取两次遍历得到的最大值即为结果。

```go
func longestValidParentheses(s string) int {
	return max(cal1(s), cal2(s))
}

func cal1(s string) int {
	left, right := 0, 0
	res := 0
	for _, ch := range s {
		if ch == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, left*2)
		} else if right > left {
			left, right = 0, 0
		}
	}
	return res
}

func cal2(s string) int {
	left, right := 0, 0
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, left*2)
		} else if left > right {
			left, right = 0, 0
		}

	}
	return res
}
```

时间复杂度 O(n)，需要遍历字符串两遍；空间复杂度 O(1)。
