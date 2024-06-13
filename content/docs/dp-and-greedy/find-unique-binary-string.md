---
title: 1980. 找出不同的二进制字符串
date: 2023-11-16T09:09:59+08:00
---

## [1980. 找出不同的二进制字符串](https://leetcode.cn/problems/find-unique-binary-string) (Medium)

给你一个字符串数组 `nums` ，该数组由 `n` 个 **互不相同** 的二进制字符串组成，且每个字符串长度都是 `n` 。请你找出并返回一个长度为 `n` 且 **没有出现** 在 `nums` 中的二进制字符串。如果存在多种答案，只需返回 **任意一个** 即可。

**示例 1：**

```
输入：nums = ["01","10"]
输出："11"
解释："11" 没有出现在 nums 中。"00" 也是正确答案。

```

**示例 2：**

```
输入：nums = ["00","01"]
输出："11"
解释："11" 没有出现在 nums 中。"10" 也是正确答案。

```

**示例 3：**

```
输入：nums = ["111","011","001"]
输出："101"
解释："101" 没有出现在 nums 中。"000"、"010"、"100"、"110" 也是正确答案。
```

**提示：**

- `n == nums.length`
- `1 <= n <= 16`
- `nums[i].length == n`
- `nums[i] ` 为 `'0'` 或 `'1'`
- `nums` 中的所有字符串 **互不相同**

## 分析


注意到每个字符串的长度正好是所有字符串的个数 n，可以这样构造与所有字符串不同的串 res：

从头遍历res，res[i] 和 nums[i][i] 不同，这样保证最后构造的res与原来的每个串至少有一位是不同的。

```go
func findDifferentBinaryString(nums []string) string {
	res := make([]byte, len(nums))
	for i := range res {
		if nums[i][i] == '0' {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

```
