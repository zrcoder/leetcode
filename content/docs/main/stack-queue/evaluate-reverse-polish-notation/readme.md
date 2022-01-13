---
title: "逆波兰表达式求值"
date: 2021-05-26T20:52:46+08:00
weight: 50
tags: [栈]
---

许多问题，用递归解决比较自然直观，也可以改成用栈模拟，这样甚至会更加直观。

## [150. 逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

难度中等

根据[ 逆波兰表示法](https://baike.baidu.com/item/逆波兰式/128437)，求表达式的值。

有效的算符包括 `+`、`-`、`*`、`/` 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

**说明：**

- 整数除法只保留整数部分。
- 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况 

**示例 1：**

```
输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
```

**示例 2：**

```
输入：tokens = ["4","13","5","/","+"]
输出：6
解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
```

**示例 3：**

```
输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
输出：22
解释：
该算式转化为常见的中缀算术表达式为：
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

**提示：**

- `1 <= tokens.length <= 104`
- `tokens[i]` 要么是一个算符（`"+"`、`"-"`、`"*"` 或 `"/"`），要么是一个在范围 `[-200, 200]` 内的整数

**逆波兰表达式：**

逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。

- 平常使用的算式则是一种中缀表达式，如 `( 1 + 2 ) * ( 3 + 4 )` 。
- 该算式的逆波兰表达式写法为 `( ( 1 2 + ) ( 3 4 + ) * )`。

逆波兰表达式主要有以下两个优点：

- 去掉括号后表达式无歧义，上式即便写成 `1 2 + 3 4 + * `也可以依据次序计算出正确结果。
- 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

函数签名：

```go
func evalRPN(tokens []string) int
```

## 分析

逆波兰表达式，对人来说不直观，但是对计算机来说非常直观～顺序读取数字，遇到运算法就计算之前读到的最后两个数字的对应运算，两个数字合为一个数字。用一个栈来做非常自然。

```go
func evalRPN(tokens []string) int {
    var stack []int
    for _, v := range tokens {
        num, err := strconv.Atoi(v)
        if err == nil {
            stack = append(stack, num)
            continue
        }
        n := len(stack)
        switch v {
        case "+":
            stack[n-2] += stack[n-1]
        case "-":
            stack[n-2] -= stack[n-1]
        case "*":
            stack[n-2] *= stack[n-1]
        case "/":
            stack[n-2] /= stack[n-1]
        }
        stack = stack[:n-1]
    }
    return stack[0]    
}
```

时空复杂度都是`O(n)`。

## [1190. 反转每对括号间的子串](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/)

难度中等

给出一个字符串 `s`（仅含有小写英文字母和括号）。

请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

注意，您的结果中 **不应** 包含任何括号。

**示例 1：**

```
输入：s = "(abcd)"
输出："dcba"
```

**示例 2：**

```
输入：s = "(u(love)i)"
输出："iloveu"
```

**示例 3：**

```
输入：s = "(ed(et(oc))el)"
输出："leetcode"
```

**示例 4：**

```
输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"
```

**提示：**

- `0 <= s.length <= 2000`
- `s` 中只有小写英文字母和括号
- 我们确保所有括号都是成对出现的

函数签名：

```go
func reverseParentheses(s string) string
```

## 分析

同样用栈解决，直接上代码：

```go
func reverseParentheses(s string) string {
	var stack [][]byte
	var cur []byte
	for i := range s {
		switch s[i] {
		case '(':
			stack = append(stack, cur)
			cur = nil
		case ')':
			if len(stack) == 0 { // invalid input
				return ""
			}
			last := stack[len(stack)-1]
			for j := len(cur) - 1; j >= 0; j-- {
				last = append(last, cur[j])
			}
			cur = last
			stack = stack[:len(stack)-1]
		default:
			cur = append(cur, s[i])
		}
	}
	return string(cur)
}
```

时空复杂度都是`O(n)`。