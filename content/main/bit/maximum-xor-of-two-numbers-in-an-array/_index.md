---
title: "421. 数组中两个数的最大异或值"
date: 2021-05-24T19:53:16+08:00
weight: 50
tags: [位,Trie树]
---

## [421. 数组中两个数的最大异或值](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/)

难度中等

给你一个整数数组 `nums` ，返回 `nums[i] XOR nums[j]` 的最大运算结果，其中 `0 ≤ i ≤ j < n` 。

**进阶：**你可以在 `O(n)` 的时间解决这个问题吗？

**示例 1：**

```
输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
```

**示例 2：**

```
输入：nums = [0]
输出：0
```

**示例 3：**

```
输入：nums = [2,4]
输出：6
```

**示例 4：**

```
输入：nums = [8,10,2]
输出：10
```

**示例 5：**

```
输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
输出：127
```

**提示：**

- `1 <= nums.length <= 2 * 10^4`
- `0 <= nums[i] <= 2^31 - 1`

函数签名：

```go
func findMaximumXOR(nums []int) int
```

## 分析

每次遍历到 `nums[i]` ，可以假设这个数字是构成结果的一个元素，已经遍历过的数字里挑一个和当前数字亦或运算，这样的话是 `O(n^2)` 的复杂度。

两层循环的朴素解法显然不是这个问题的目的，怎么降低复杂度呢？根据提示只需遍历一遍数组就能得出结果，这是怎么做到的？

可以用前缀树的技巧把已经遍历的数字记录起来。具体来说，把每个数字看做二进制，根据题目限制，最多30位，将其插入前缀树里，前缀树的深度最多为30。

用当前数字和之前的数字亦或时，尽量在高位取得 1 即可。

> 这里的前缀树每个节点最多 2 个孩子节点，是一棵简单二叉树。

```go
type Trie struct {
	zero, one *Trie
}

const bitLimt = 30

func (t *Trie) insert(x int) {
	cur := t
	for i := bitLimt; i >= 0; i-- {
		b := (x >> i) & 1
		if b == 0 {
			if cur.zero == nil {
				cur.zero = &Trie{}
			}
			cur = cur.zero
		} else {
			if cur.one == nil {
				cur.one = &Trie{}
			}
			cur = cur.one
		}
	}
}

func (t *Trie) check(x int) int {
	res := 0
	cur := t
	for i := bitLimt; i >= 0; i-- {
		res <<= 1
		b := (x >> i) & 1
		if b == 0 {
			if cur.one != nil {
				cur = cur.one
				res++
			} else {
				cur = cur.zero
			}
		} else {
			if cur.zero != nil {
				cur = cur.zero
				res++
			} else {
				cur = cur.one				
			}
		}
	}
	return res
}

func findMaximumXOR(nums []int) int {
	res := 0
	root := &Trie{}
	for i := 1; i < len(nums); i++ {
		root.insert(nums[i-1])
		res = max(res, root.check(nums[i]))
	}
	return res
}
```

时空复杂度都是 `O(n*30) = O(n)`。