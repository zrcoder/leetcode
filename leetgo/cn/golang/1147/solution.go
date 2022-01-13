package solution

/*
## [1147. Longest Chunked Palindrome Decomposition](https://leetcode.cn/problems/longest-chunked-palindrome-decomposition) (Hard)

你会得到一个字符串 `text` 。你应该把它分成 `k` 个子字符串 `(subtext1, subtext2，…， subtextk)` ，要求满足:

- `subtextᵢ` 是 **非空** 字符串
- 所有子字符串的连接等于 `text` ( 即 `subtext₁ + subtext₂ + ... + subtextₖ == text` )
- 对于所有 i 的有效值( 即 `1 <= i <= k` ) ， `subtextᵢ == subtextₖ - ᵢ + ₁` 均成立

返回 `k` 可能最大值。

**示例 1：**

```
输入：text = "ghiabcdefhelloadamhelloabcdefghi"
输出：7
解释：我们可以把字符串拆分成 "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)"。

```

**示例 2：**

```
输入：text = "merchant"
输出：1
解释：我们可以把字符串拆分成 "(merchant)"。

```

**示例 3：**

```
输入：text = "antaprezatepzapreanta"
输出：11
解释：我们可以把字符串拆分成 "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)"。

```

**提示：**

- `1 <= text.length <= 1000`
- `text` 仅由小写英文字符组成


*/

// [start] don't modify
func longestDecomposition(text string) int {
	res := 0
	for len(text) > 0 {
		n := len(text)
		preLen := 1
		for ; preLen <= n/2 && text[:preLen] != text[n-preLen:]; preLen++ {
		}
		if preLen > n/2 {
			res++
			break
		}
		res += 2
		text = text[preLen : n-preLen]
	}
	return res
}
// Time complex: O(n^2)

// [end] don't modify
