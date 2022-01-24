---
title: "剑指 Offer 20. 表示数值的字符串"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [剑指 Offer 20. 表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

难度中等

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

 函数签名：

```go
func isNumber(s string) bool
```

## 分析

题目描述非常模糊，到底什么样的字符串能表示数值只举了几个例子，并没有给出精确规范。查看题解，精确规范如下：

```
在 C++ 文档中，描述了一个合法的数值字符串应当具有的格式。具体而言，它包含以下部分：

符号位，即 +、− 两种符号
整数部分，即由若干字符 0−9 组成的字符串
小数点
小数部分，其构成与整数部分相同
指数部分，其中包含开头的字符 e（大写小写均可）、可选的符号位，和整数部分

相比于 C++ 文档而言，本题还有一点额外的不同，即允许字符串首末两端有一些额外的空格。
在上面描述的五个部分中，每个部分都不是必需的，但也受一些额外规则的制约，如：

如果符号位存在，其后面必须跟着数字或小数点。
小数点的前后两侧，至少有一侧是数字。
```

### 朴素实现

首尾空格可以事先去除。一个数字字符串包含符号、数字、e、小数点四种元素，可以用 4 个 bool 变量维护这四种元素是否已经出现过，在遍历过程中发现违反规则的情况直接返回 false。

```go
func isNumber(s string) bool {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return false
    }    
    var signSeen, digitSeen, eSeen, dotSeen bool
    for _, v := range s {
        switch {
        case v == '+' || v == '-':
            // 全局最多一个符号、且必须在数字和小数点之前。（注意允许在 e 或 E 之后，e 或 E 比较特殊）
            if signSeen || digitSeen || dotSeen {
                return false
            }
            signSeen = true
        case v >= '0' && v <= '9':
            // 任意位置都可以出现数字
            digitSeen = true
        case v == 'e' || v == 'E':
            // e/E 最多只能出现一次，且之前必须有数字
            if eSeen || !digitSeen {
                return false
            }
            // 注意这里把符号、小数点、数字标识重置为 false
            signSeen, dotSeen, digitSeen = false, false, false
            eSeen = true
        case v == '.':
            // 小数点最多只能出现一次，且不能在 e/E 之后
            if dotSeen || eSeen {
                return false
            }
            dotSeen = true
        default:
            // 其他非法字符
            return false
        }
    }
    // 至少要有一个数字
    return digitSeen
}
```

时间复杂度 O(n)， 空间复杂度 O(1)

### 有限状态机

这是一个更一般的思路。

```go
type CharType int

const (
	CharIllegal CharType = iota
	CharExp
	CharPoint
	CharSign
	CharNumber
)

const charTypeCnt = 5

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	}
	return CharIllegal
}

type State int

const (
	StateIllegal State = iota
	StateInitial
	StateIntSign // 整数部分的符号
	StateInteger // 正数部分的数字
	StatePoint
	StatePointWithoutInt // 没有正数部分的小数点
	StateFraction        // 小数部分
	StateExp             // 指数部分
	StateExpSign         // 指数部分的符号
	StateExpNumber       // 指数部分的数字
)

const stateCnt = 10

var transfer = [stateCnt][charTypeCnt]State{
	StateInitial: { // 可以以数字、小数点、符号开始
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
		CharSign:   StateIntSign,
	},
	StateInteger: {
		CharNumber: StateInteger,
		CharExp:    StateExp,
		CharPoint:  StatePoint,
	},
	StatePoint: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
	},
	StateIntSign: {
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
	},
	StatePointWithoutInt: {
		CharNumber: StateFraction,
	},
	StateFraction: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
	},
	StateExp: {
		CharNumber: StateExpNumber,
		CharSign:   StateExpSign,
	},
	StateExpSign: {
		CharNumber: StateExpNumber,
	},
	StateExpNumber: {
		CharNumber: StateExpNumber,
	},
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	state := StateInitial
	for i := 0; i < len(s); i++ {
		charType := toCharType(s[i])
		if transfer[state][charType] == StateIllegal {
			return false
		}
		state = transfer[state][charType]
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber
}
```

时间复杂度 O(n)， 空间复杂度 O(1)

