package solution

/*
## [1641. Count Sorted Vowel Strings](https://leetcode.cn/problems/count-sorted-vowel-strings) (Medium)

给你一个整数 `n`，请返回长度为 `n` 、仅由元音 ( `a`, `e`, `i`, `o`, `u`) 组成且按 **字典序排列** 的字符串数量。

字符串 `s` 按 **字典序排列** 需要满足：对于所有有效的 `i`， `s[i]` 在字母表中的位置总是与 `s[i+1]` 相同或在 `s[i+1]` 之前。

**示例 1：**

```
输入：n = 1
输出：5
解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]

```

**示例 2：**

```
输入：n = 2
输出：15
解释：仅由元音组成的 15 个字典序字符串为
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后

```

**示例 3：**

```
输入：n = 33
输出：66045

```

**提示：**

- `1 <= n <= 50`


*/

// [start] don't modify
// dp
func countVowelStrings(n int) int {
    a, e, i, o, u := 1, 1, 1, 1, 1 // 分别表示当前以 a, e, i, o, u 字母结尾且满足字典序的字符串的数量
    for x := 2; x <= n; x++ {
        a, e, i, o, u = a, a+e, a+e+i, a+e+i+o, a+e+i+o+u
    }
    return a+e+i+o+u
}
// math 相当于要把 n 个小球分隔成 5 组，但是允许有空组，可以加5个小球转化为把n+5个小球分隔成 5 个非空组，仅需要再 n+4 个空档选4个插入隔板
// 答案即 C_{n+4}^{4}
func countVowelStrings(n int) int {
    return (n+4)*(n+3)*(n+2)*(n+1) / 24
}
// [end] don't modify
