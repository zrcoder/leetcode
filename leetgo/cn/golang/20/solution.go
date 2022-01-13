package solution

/*
## [20. Valid Parentheses](https://leetcode.cn/problems/valid-parentheses) (Easy)

给定一个只包括 `'('`， `')'`， `'{'`， `'}'`， `'['`， `']'` 的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

**示例 1：**

```
输入：s = "()"
输出：true

```

**示例 2：**

```
输入：s = "()[]{}"
输出：true

```

**示例 3：**

```
输入：s = "(]"
输出：false

```

**提示：**

- `1 <= s.length <= 10⁴`
- `s` 仅由括号 `'()[]{}'` 组成


*/

// [start] don't modify
func isValid(s string) bool {
    if len(s)%2 == 1 { 
        return false
    }   
    dic := map[rune]rune{
        ']': '[',
        '}': '{',
        ')': '(',
    }   
    stack := make([]rune, 0, len(s))
    for _, v := range s { 
        switch v { 
        case ']', '}', ')':
            if len(stack) == 0 || stack[len(stack)-1] != dic[v] {
                return false
            }   
            stack = stack[:len(stack)-1]
        default:
            stack = append(stack, v)
        }   
    }   
    return len(stack) == 0
}
// [end] don't modify
