---
title: "子集"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [378. 子集](https://leetcode-cn.com/problems/subsets)

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。
```
示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```
## 分析
### 朴素实现
时空复杂度均为 `O(n*2^n)`
```go
func subsets(nums []int) [][]int {
	res := [][]int{{}} // 空集也是子集之一
	for _, num := range nums {
		for _, r := range res {
			res = append(res, append(getCopy(r), num))
		}
	}
	return res
}
```
其中 getCopy 函数用于深拷贝切片：

```go
func getCopy(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}
```

> 也可以这样来拷贝一个切片：
>
> ```go
> func getCopy(s []int) []int {
> 	return append([]int{}, s...)
> }
> ```

### 通用回溯写法

```go
func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			res = append(res, getCopy(cur))
			return
		}
		// 不选择当前元素
		backtrack(i + 1)
		// 选择当前元素
		cur = append(cur, nums[i])
		backtrack(i + 1)
		cur = cur[:len(cur)-1]
	}
	backtrack(0)
	return res
}
```
> 选择、不选择两段内容可以改变顺序，不影响结果

### 固定前段的回溯写法

```go
func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(start int)
	backtrack = func(start int) {
		res = append(res, getCopy(cur))
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return res
}
```
### 二进制枚举
nums 里的每个元素，要么在结果中，要么不在结果中。用一个 n 位的 bitset 来表示各个元素在不在结果中，
如 000...000 表示所有元素都不在结果中，000..011 表示后边两个元素在结果中。这样从 000...000 一直递增到 111...111， 就枚举了所有可能。

> 局限：len(nums)不能大于64， 否则无法用一个int 状态。

时空复杂度均为O(n*2^n)
```go
func subsets(nums []int) [][]int {
	var res [][]int
	max := 1 << len(nums)
	for mask := 0; mask < max; mask++ {
		var cur []int
		for i, v := range nums {
			if (1<<i)&mask != 0 {
				cur = append(cur, v)
			}
		}
		res = append(res, cur)
	}
	return res
}
```

## [90. 子集 II](https://leetcode-cn.com/problems/subsets-ii/)

难度中等

给你一个整数数组 `nums` ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。返回的解集中，子集可以按 **任意顺序** 排列。

**示例 1：**

```
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
```

**示例 2：**

```
输入：nums = [0]
输出：[[],[0]]
```

**提示：**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`

函数签名：

```go
func subsetsWithDup(nums []int) [][]int
```

## 分析

可以用通用回溯写法，最后去重。怎么去重？因为元素规模，可以把每个子集转化成 string 后用哈希表去重。

另一个比较好的做法是：可以事先排序，回溯时在不选择当前元素的时候，就可以循环找到下一个不同的元素，这样得到在子集就不存在重复的可能。

```go
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var cur []int
	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(nums) {
			res = append(res, getCopy(cur))
			return
		}
		// 不选择当前元素
		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		backtrack(j)
		
		// 选择当前元素
		cur = append(cur, nums[i])
		backtrack(i + 1)
		cur = cur[:len(cur)-1]		
	}
	backtrack(0)
	return res
}
```

## [1593. 拆分字符串使唯一子字符串的数目最大](https://leetcode-cn.com/problems/split-a-string-into-the-max-number-of-unique-substrings/)

难度中等

给你一个字符串 `s` ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。

字符串 `s` 拆分后可以得到若干 **非空子字符串** ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 **唯一的** 。

注意：**子字符串** 是字符串中的一个连续字符序列。

**示例 1：**

```
输入：s = "ababccc"
输出：5
解释：一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。
```

**示例 2：**

```
输入：s = "aba"
输出：2
解释：一种最大拆分方法为 ['a', 'ba'] 。
```

**示例 3：**

```
输入：s = "aa"
输出：1
解释：无法进一步拆分字符串。
```

**提示：**

- `1 <= s.length <= 16`
- `s` 仅包含小写英文字母

函数签名：

```go
func maxUniqueSplit(s string) int
```

## 分析

### 回溯

对子集问题加了限制：所有子集不能相同。可以在回溯的过程中用哈希表去重。

> 另需注意，s 本身就是一个子集，但是这里空集不算。如示例3。

```go
func maxUniqueSplit(s string) int {
	type Set = map[string]bool
	var dfs func(i int)
	set := make(Set, 0)
	res := 1
	dfs = func(start int) {
		if start == len(s) {
			res = max(res, len(set))
			return
		}
		for i := start; i < len(s); i++ {
			sub := s[start:i+1]
			if !set[sub] {
				set[sub] = true
				dfs(i+1)
				delete(set, sub)
			}
		}
	}
	dfs(0)
	return res
}
```
## [77. 组合](https://leetcode-cn.com/problems/combinations)
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:
```
输入:  n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

## 分析
子集问题的一个特例，可以直接用上边的方法，加上子集长度固定为 k 的约束：
```go
func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(i int)
	dfs = func(i int) {
		if i > n+1 { // 这个条件还可以收紧，需要优化剪枝
			return
		}
		if len(cur) == k {
			res = append(res, copySlice(cur))
			return
		}
		dfs(i+1)
		cur = append(cur, i)
		dfs(i+1)
		cur = cur[:len(cur)-1]
	}
	dfs(1)
	return res
}
```
注意注释，实际上，如果 cur 等长度加上剩余要枚举的数字的个数小于 k，就已经不需要再向后递归尝试了，注释处的条件改为：
```go
if len(cur)+n-i+1 < k
```
## [39. 组合总和](https://leetcode-cn.com/problems/combination-sum)
给定一个无重复元素的数组 candidates 和一个目标数 target ，  
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。  
解集不能包含重复的组合。

```
示例 1：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

示例 2：
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
```
## 分析

可以采用子集问题的两种写法，注意每个元素可以无限次使用，这里有些特殊。
### 固定前段的递归回溯
```go
func combinationSum1(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, start int)
	dfs = func(t, start int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, getCopy(cur))
			return
		}
		// 从 start 开始，而不是从 0 开始，防止重复的组合出现
		for j := start; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			dfs(t-candidates[j], j)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return res
}
```
### 通用回溯写法
```go
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, i int)
	dfs = func(t, i int) {
		if t < 0 || i == len(candidates) {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}

		// 使用 i 处元素
		cur = append(cur, candidates[i])
		// 元素可无限制重复使用，这里 i 不加1
		dfs(t-candidates[i], i)
		cur = cur[:len(cur)-1]

		// 不使用 i 处元素，这里 i + 1
		dfs(t, i+1)
	}
	dfs(target, 0)
	return res
}
```
## [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii)
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。  
candidates 中的每个数字在每个组合中只能使用一次。

说明：  
所有数字（包括目标数）都是正整数。  
解集不能包含重复的组合。
```
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
```
## 分析

可以参考上边《子集II》问题，采用通用回溯解法：

```go
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var cur []int
	var backtrack func(t, i int)
	backtrack = func(t, i int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, getCopy(cur))
			return
		}
        if i == len(candidates) {
            return
        }
		cur = append(cur, candidates[i])
		backtrack(t-candidates[i], i+1)
		cur = cur[:len(cur)-1]
		j := i + 1
		for j < len(candidates) && candidates[i] == candidates[j] {
			j++
		}
		backtrack(t, j)
	}
	backtrack(target, 0)
	return res
}
```

