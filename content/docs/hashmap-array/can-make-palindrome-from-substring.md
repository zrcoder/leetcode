---
title: 1177. 构建回文串检测
---

## [1177. 构建回文串检测](https://leetcode.cn/problems/can-make-palindrome-from-substring) (Medium)

给你一个字符串 `s`，请你对 `s` 的子串进行检测。

每次检测，待检子串都可以表示为 `queries[i] = [left, right, k]`。我们可以 **重新排列** 子串 `s[left], ..., s[right]`，并从中选择 **最多** `k` 项替换成任何小写英文字母。

如果在上述检测过程中，子串可以变成回文形式的字符串，那么检测结果为 `true`，否则结果为 `false`。

返回答案数组 `answer[]`，其中 `answer[i]` 是第 `i` 个待检子串 `queries[i]` 的检测结果。

注意：在替换时，子串中的每个字母都必须作为 **独立的** 项进行计数，也就是说，如果 `s[left..right] = "aaa"` 且 `k = 2`，我们只能替换其中的两个字母。（另外，任何检测都不会修改原始字符串 `s`，可以认为每次检测都是独立的）

**示例：**

```
输入：s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
输出：[true,false,false,true,true]
解释：
queries[0] : 子串 = "d"，回文。
queries[1] : 子串 = "bc"，不是回文。
queries[2] : 子串 = "abcd"，只替换 1 个字符是变不成回文串的。
queries[3] : 子串 = "abcd"，可以变成回文的 "abba"。 也可以变成 "baab"，先重新排序变成 "bacd"，然后把 "cd" 替换为 "ab"。
queries[4] : 子串 = "abcda"，可以变成回文的 "abcba"。

```

**提示：**

- `1 <= s.length, queries.length <= 10^5`
- `0 <= queries[i][0] <= queries[i][1] < s.length`
- `0 <= queries[i][2] <= s.length`
- `s` 中只有小写英文字母

## 分析

因为可以排序，子串是否能形成回文，只需要统计各个字母个数，只需要考虑个数是奇数的情况有多少次。

假设对于一个子串，统计出个数为奇数的字母共 odds 个，如果 odds < 2 那么无须替换，子串吧本身重排后就能变成回文。

odds 比较大时，我们可以做 最多 k 次替换，每次替换选择个数为奇数的一个字母替换成另一个个数为奇数的字母，这样会消去两个奇数字母，k 次能消去 2*k 个，所以只需考虑 odds-2*k 是否小于2 即可确定子串在替换排序后变成回文。

为了效率，可以用计算字母个数的前缀和。

```go
func canMakePaliQueries(s string, queries [][]int) []bool {
	const charsLimit = 26
	cnt := make([][charsLimit]int, len(s)+1)
	for i, c := range s {
		for j := 0; j < charsLimit; j++ {
			cnt[i+1][j] = cnt[i][j]
			if j == int(c-'a') {
				cnt[i+1][j]++
			}
		}
	}
	n := len(queries)
	answer := make([]bool, n)
	for i, q := range queries {
		lo, hi, k := q[0], q[1], q[2]
		odds := 0
		for j := 0; j < charsLimit; j++ {
			if (cnt[hi+1][j]-cnt[lo][j])%2 == 1 {
				odds++
			}
		}
		answer[i] = odds-2*k < 2
	}
	return answer
```

时间复杂度是 O((n+m)*c), 其中 n 是 s 长度, m 是 queries 长度， c 是字符集大小（这里是26）。

空间复杂度：O(nc)。

### 优化

实际上，仅需知道子串中字母个数的奇偶性而不需要知道个数本身。这样可以用一个 int 来代替上边的 [26]int 。

时间复杂度： O(n+mlogc)， 空间复杂度：O(n)

```go
func canMakePaliQueries(s string, queries [][]int) []bool {
	cnt := make([]int, len(s)+1)
	for i, c := range s {
		cnt[i+1] = cnt[i] ^ (1 << (c - 'a'))
	}
	answer := make([]bool, len(queries))
	for i, q := range queries {
		lo, hi, k := q[0], q[1], q[2]
		odds := bits.OnesCount(uint(cnt[hi+1] ^ cnt[lo]))
		answer[i] = odds-2*k < 2
	}
	return answer
}

```

Local tests:

```go

func Test_canMakePaliQueries(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				s: "abcda",
				queries: [][]int{
					{3, 3, 0},
					{1, 2, 0},
					{0, 3, 1},
					{0, 3, 2},
					{0, 4, 1},
				},
			},
			want: []bool{true, false, false, true, true},
		},
	}
	for _, tt := range tests {
		if got := canMakePaliQueries(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. canMakePaliQueries() = %v, want %v", tt.args.s, got, tt.want)
		}
	}
}

```
