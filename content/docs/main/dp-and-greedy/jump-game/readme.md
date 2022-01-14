---
title: "跳跃游戏"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game)
给定一个非负整数数组，你最初位于数组的第一个位置。  
数组中的每个元素代表你在该位置可以跳跃的最大长度。  
判断你是否能够到达最后一个位置。  
```
示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
```

## 分析
直观贪心、自底向上dp及自顶向下dp。 

{{<tabs "I">}}
{{<tab "直观贪心">}}
```
贪心，也是最朴素的实现方法；时间复杂度O(n), 空间复杂度O(1)
```
```go
func canJump(nums []int) bool {
	farthest := 0
	n := len(nums)
	for i, v := range nums {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+v)
		if farthest >= n-1 {
			return true
		}
	}
	return false
}
```
{{</tab>}}
{{<tab "自底向上动态规划">}}
```
时间复杂度O(n^2),空间复杂度O(n)
```
```go
func canJump1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		// j的初值需要防止索引越界; 这里是从右向左遍历，也可以从左向右
		for j := min(len(nums)-1, i+nums[i]); j > i; j-- { 
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
```
{{</tab>}}
{{<tab "自顶向下动态规划">}}
```
自顶向下动态规划，或者理解为带备忘的回溯
时间复杂度O(n^2),空间复杂度O(2n)=O(n)，第一个n是栈空间开销，第二个是dp数组开销
```
```go
func canJump2(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	const ok, nok = 1, 2 // unknown = 0
	memo := make([]int, len(nums))
	memo[len(nums)-1] = ok
	var canJumpFrom func(pos int) bool
	canJumpFrom = func(pos int) bool {
		if memo[pos] == ok {
			return true
		}
		if memo[pos] == nok {
			return false
		}
		// i的初值需要防止索引越界；这里是从右向左遍历，也可以从左向右
		for i := min(pos+nums[pos], len(nums)-1); i > pos; i-- { 
			if canJumpFrom(i) {
				memo[pos] = ok
				return true
			}
		}
		memo[pos] = nok
		return false
	}
	return canJumpFrom(0)
}
```
{{</tab>}}
{{</tabs>}}

## [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii)
给定一个非负整数数组，你最初位于数组的第一个位置。  
数组中的每个元素代表你在该位置可以跳跃的最大长度。  
假设你总是可以到达数组的最后一个位置。你的目标是使用最少的跳跃次数到达数组的最后一个位置。  
```
示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```

## 分析
动态规划的解法会在最后一个用例超出时间限制；这里需要贪心策略。  
值得一提的是另一个最坏情况时间复杂度O(n^2)的实现，思路直观，且并不超时。

{{<tabs "II">}}
{{<tab "自顶向下动态规划">}}
```
自顶向下动态规划，或者理解为带备忘的回溯
时间复杂度O(n^2),空间复杂度O(n)
```
```go
func jump1(nums []int) int {
	n := len(nums)
     // memo[i]表示从位置i跳到最后位置的最小步数
	memo := make([]int, n) 
	memo[n-1] = 0
	for i := 0; i < n-1; i++ {
		memo[i] = n // n相当于max
	}
	var helper func(pos int) int
	helper = func(pos int) int {
		if memo[pos] < n {
			return memo[pos]
		}
		end := min(pos+nums[pos], n-1)
		for i := end; i > pos; i-- {
			memo[pos] = min(memo[pos], helper(i)+1)
		}
		return memo[pos]
	}
	return helper(0)
}
```
{{</tab>}}
{{<tab "自底向上动态规划">}}
```
时间复杂度O(n^2),空间复杂度O(n)
```
```go
func jump2(nums []int) int {
	n := len(nums)
     // dp[i]表示从位置i跳到最后位置的最小步数
	dp := make([]int, n) 
	dp[n-1] = 0
	for i := 0; i < n-1; i++ {
		dp[i] = n // n相当于max
	}
	for i := n - 2; i >= 0; i-- {
		end := min(nums[i]+i, n-1)
		for j := end; j > i; j-- {
			dp[i] = min(dp[i], dp[j]+1)
		}
	}
	return dp[0]
}
```
{{</tab>}}
{{<tab "逆向思考">}}
```
要到达最后一个位置，前一个位置在哪？找到后，再继续寻找上上个位置，直到找到第0个位置；
为了使最终步数最少，每次需要找到距离当前位置最远的距离，从左到右遍历数组，第一个满足的位置就是了。
时间复杂度， 最坏情况下是O(n^2), 空间复杂度O(1)
```
```go
func jump3(nums []int) int {
	pos := len(nums) - 1
	result := 0
	for pos > 0 {
		i := 0
		// 从左到右找到第一个能跳到pos的位置i即为最优的i
		for i < pos && i+nums[i] < pos { 
			i++
		}
		pos = i
		result++
	}
	return result
}
```
{{</tab>}}
{{<tab "贪心策略">}}
```
每次在可跳范围内选择可以跳得更远的位置
遍历时对i+nums[i]使用贪心策略做选择
例如，对于数组 [2,3,1,2,4,2,3]，初始位置是下标 0，从下标 0 出发，最远可到达下标 2。
下标 0 可到达的位置中，下标 1 的值是 3，从下标 1 出发可以达到更远的位置，因此第一步到达下标 1。
从下标 1 出发，最远可到达下标 4。
下标 1 可到达的位置中，下标 4 的值是 4 ，从下标 4 出发可以达到更远的位置，因此第二步到达下标 4。
...
依次类推
时间复杂度O(n), 空间复杂度O(1)
```
```go
func jump(nums []int) int {
	end := 0 // 表示当前能跳的边界；如以上举例，end分别是0,2,4,8
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
```
{{</tab>}}
{{</tabs>}}