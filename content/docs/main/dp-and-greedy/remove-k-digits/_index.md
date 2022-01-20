---
title: "去掉 k 个字符使结果保持一定顺序"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

以下几个问题类似，都可以使用贪心策略解决，具体实施时使用类似单调栈的思路降低时间复杂度和代码复杂度。


## [402. 移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits)

`难度中等`

给定一个以字符串表示的非负整数 *num*，移除这个数中的 *k* 位数字，使得剩下的数字最小。

**注意:**

- *num* 的长度小于 10002 且 ≥ *k。*
- *num* 不会包含任何前导零。

**示例 1 :**

```
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
```

**示例 2 :**

```
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
```

示例 **3 :**

```
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
```

函数定义如下：

````go
func removeKdigits(num string, k int) string
````

## 分析

先考虑简单情况，k == 1。对于一个数字字符串，去掉哪一个数字会使结果字符串代表点数字最小呢？

如题目里的 1432219、10200，显然应该分别去掉 4 和 1，再比如 235419，应该去掉 5。

很自然地就解决了 k == 1 的情况，则个自然思路用的策略到底是什么？

```
从左到右找到第一个比后一位大的数字，就是要去除的目标。
```

> 如果原数字各位已经升序排列，只需要删除最后一位

只需要循环 k 次，每次按照如上贪心策略去除一位数字即可。

现在问题已经解决了。只是时间复杂度较高，是 `O(nk)`，其中 n 代表字符串长度，并且代码写起来比较繁琐。

可以稍作转化，`去掉 k 个 <=> 保留 n-k 个`。

遍历 num，不断将当前位数字加入结果数组，在加入前判断当前数字是不是小于结果数组最后几个元素，如果小则删掉结果数组最后那几个元素。注意总共删除的数字最多 k 个。

如果遍历完成后还没有删够 k 个，还差 x 个要删除，只需要去掉结果数组的最后 x 个元素。

最后将结果数组转化成 string 返回，注意删除前导零，以及 “” 要返回 “0” 的情况。

```go
func removeKdigits(num string, k int) string {
  res := make([]byte, 0, len(num))
  for i := range num {
    for k > 0 && len(res) > 0 && res[len(res)-1] > num[i] {
      res = res[:len(res)-1]
      k--
    }
    res = append(res, num[i])
  }
  res = res[:len(res)-k] // 还有 k 个没有删除，k 可能为 0
  s := string(res)
  s = strings.TrimLeft(s, "0")
  if s == "" {
    return "0"
  }
  return s
}
```

时间复杂度 `O(n)`，每个元素最多进出结果数组 res 一次。可以看到 res 实际类似一个单调栈。

空间复杂度 `O(n)`，结果数组和结果字符串占用的空间。

## [316. 去除重复字母](https://leetcode-cn.com/problems/remove-duplicate-letters/)

`难度中等`

给你一个字符串 `s` ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 **返回结果的字典序最小**（要求不能打乱其他字符的相对位置）。

**注意：**该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同

**示例 1：**

```
输入：s = "bcabc"
输出："abc"
```

**示例 2：**

```
输入：s = "cbacdcbc"
输出："acdb"
```

**提示：**

- `1 <= s.length <= 104`
- `s` 由小写英文字母组成

函数定义如下：

```go
func removeDuplicateLetters(s string) string
```

## 分析

类似上一问题，最终都是保持原相对位置且字典序最小。只是上一问题要保留 n- k 个元素，元素可以有重复，当前问题是保留所有元素各一个。

可以用 大小 26 的 cnt 数组事先统计每个字母出现的次数，之后遍历字符串，如果结果中没有当前字母则加入结果，在加入之前判断结果字符串最后几位是否比当前字母大，如果大且后边还有就可以删除。

怎么判断一个字母是否已经在结果中？只需要一个 大小 26 的 bool 数组来记录即可。

怎么判断后边还有没有？只需要更新 cnt 数组即可，用完则对应值为 0。

```go
func removeDuplicateLetters(s string) string {
  cnt := [26]int{}
  for _, v := range s {
    cnt[v-'a']++
  }
  inRes := [26]bool{}
  res := make([]byte, 0, len(s))
  for i := range s {
    ch := s[i]
    cnt[ch-'a']--
    if inRes[ch-'a'] {
      continue
    }
    for len(res) > 0 &&  res[len(res)-1] > ch && cnt[res[len(res)-1]-'a'] > 0  {
      inRes[res[len(res)-1] - 'a'] = false
      res = res[:len(res)-1]
    }
    res = append(res, ch)
    inRes[ch-'a'] = true
  }
  return string(res)
}
```

时空复杂度同上个问题，都是 `O(n)`。

## [321. 拼接最大数](https://leetcode-cn.com/problems/create-maximum-number/)

`难度困难`

给定长度分别为 `m` 和 `n` 的两个数组，其元素由 `0-9` 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 `k (k <= m + n)` 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。

求满足该条件的最大数。结果返回一个表示该最大数的长度为 `k` 的数组。

**说明:** 请尽可能地优化你算法的时间和空间复杂度。

**示例 1:**

```
输入:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
输出:
[9, 8, 6, 5, 3]
```

**示例 2:**

```
输入:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
输出:
[6, 7, 6, 0, 4]
```

**示例 3:**

```
输入:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
输出:
[9, 8, 9]
```

函数定义：

```go
func maxNumber(nums1 []int, nums2 []int, k int) []int
```

## 分析

我们已经可以解决一个数组的问题，当前问题扩展到了两个数组，考虑转化成一个数组的问题。

可以从一个数组找到 x 个，另一个数组中找到剩下的 k - x 个，最后合并成结果即可。

从一个数组找到 x 个，是我们已经解决过的问题！

不过还有个问题要解决：怎么合并结果？因为找到的两组数不一定有序，不能按照常规的合并方式合并。

如 [6，3， 2] 和 [6，7]，第一次应该选第二个数组里的 6 而不是第一个数组里的6， 这样会得到正确结果 [6，7， 6， 3，2]，否则得到结果 [6，6 ，7 ，3，2]，错了。实际应该看两个数组代表的整个数字哪个大，每次选大的那个数字首数字，且要注意长度不同时补齐长度。如这里的 670 > 632 所以选应该选第二个数组里的 6.

```go
func maxNumber(nums1 []int, nums2 []int, k int) []int {
  var res []int
  // 需要从 nums1 中找到 x 个数字
  // x 的最小值
  from := max(k-len(nums2), 0)
  // x 的最大值
  end := min(k, len(nums1))
  for x := from; x <= end; x++ { // O(k)
    // 从 nums1 中找到 x 个数字，从 nums2 中找到剩余的 k-x 个数字，最后合并
    sorted := merge(find(nums1, x), find(nums2, k-x)) // O(m + n + k*(m+n))
    if len(res) == 0 || larger(sorted, res) { // O(k)
      res = sorted
    }
  }
  return res
}
```

```go
// 在 nums 中找出 x 个元素使得结果字典序最大，这些元素保持在 nums 里的相对顺序
func find(nums []int, x int) []int {
  res := make([]int, 0, len(nums))
  k := len(nums)-x
  for _, v := range nums {
    for k > 0 && len(res) > 0 && res[len(res)-1] < v {
      res = res[:len(res)-1]
      k--
    }
    res = append(res, v)
  }
  return res[:len(res)-k]
}
```

```go
// 合并两个数组，使得结果代表的数字最大
func merge(s1, s2 []int) []int {
  m, n := len(s1), len(s2)
  res := make([]int, 0, m+n)
  for i, j := 0, 0; i < m || j < n; {
    if i < m && (j == n || larger(s1[i:], s2[j:])) {
      res = append(res, s1[i])
      i++
    } else {
      res = append(res, s2[j])
      j++
    }
  }
  return res
}
```

```go
// 返回哪个数组代表的数字更大，如果长度不同，短的数组后边补齐 0
// 即只看 min(len(s1), len(s2)) 位，从头比较
func larger(s1, s2 []int) bool {
  n := min(len(s1), len(s2))
  for i := 0; i < n; i++ {
    if s1[i] == s2[i] {
      continue
    }
    return s1[i] > s2[i]
  }
  return len(s1) > len(s2)
}
```

时间复杂度：`O(k^2*(m+n))`，其中 m 和 n 分别是两个数组的大小。循环内外的复杂度见注释。

空间复杂度：`O(k*(m+n))`。

> min 和 max 函数略