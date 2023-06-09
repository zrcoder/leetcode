---
title: "最大子序和"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray)

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。  

```
示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
```

## 动态规划一

定义 dp[i] 表示以 nums[i] 结尾且包含 nums[i] 的连续子数组最大和  
则 dp[i] = max(0, dp[i-1]) + nums[i]  
最终的结果即遍历dp后的最大元素  
首先可以在确定dp数组的每个元素时更新最终结果，不一定要完全确定了dp数组再遍历获取最终结果  
其次，每次dp只跟上次的dp值有关，dp数组可以优化为一个变量  
时间复杂度O(n), 空间复杂度O(1)  

## 动态规划二

dp也可定义为当前为止连续子数组的和，如果 dp 小于 0了，说明之前的元素对后边新元素的和没有正向贡献，可以重新开始计算和，dp 置为0  

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := 0, nums[0]
	for _, v := range nums {
		dp = max(0, dp) + v
		res = max(res, dp)
	}
	return res
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := 0, nums[0]
	for _, v := range nums {
		dp += v
		res = max(res, dp)
		if dp < 0 {
			dp = 0
		}
	}
	return res}
```

## [1191. K 次串联后最大子数组之和](https://leetcode-cn.com/problems/k-concatenation-maximum-sum)

给你一个整数数组 arr 和一个整数 k。  
首先，我们要对该数组进行修改，即把原数组 arr 重复 k 次。  
举个例子，如果 arr = [1, 2] 且 k = 3，那么修改后的数组就是 [1, 2, 1, 2, 1, 2]。  
然后，请你返回修改后的数组中的最大的子数组之和。  
注意，子数组长度可以是 0，在这种情况下它的总和也是 0。  
由于 结果可能会很大，所以需要 模（mod） 10^9 + 7 后再返回。

```
示例 1：
输入：arr = [1,2], k = 3
输出：9

示例 2：
输入：arr = [1,-2,1], k = 5
输出：2

示例 3：
输入：arr = [-1,-2], k = 7
输出：0
 
提示：
1 <= arr.length <= 10^5
1 <= k <= 10^5
-10^4 <= arr[i] <= 10^4
```

类似53的解法：

```go
func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
	total := n * k
	if total == 0 {
		return 0
	}
	var res, dp int
	for i := 0; i < total; i++ {
		dp = (max(0, dp) + arr[i%n]) % 1000000007
		res = max(res, dp)
	}
	return res
}
```

不过这样会超时，实际上当 k 很大的时候，没有必要遍历 n*k 次，  
只需要遍历 n*2 次,最后的结果加上 max(0,sum) * (k-2)即可  

```go
func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
	total := min(2, k) * n
	if total == 0 {
		return 0
	}

	dp, res, sum := 0, arr[0], 0
	for i := 0; i < total; i++ {
		v := arr[i%n]
		dp = (max(dp, 0) + v) % 1000000007
		res = max(res, dp)
		if i < len(arr) {
			sum = (sum + v) % 1000000007
		}
	}
	if k <= 2 {
		return res
	}
	// 题目没有描述清楚，实际在结果为负时，需要返回 0
	// return (max(sum, 0)*(k-2) + res) % 1000000007
	return max(0, (max(sum, 0)*(k-2)+res)%1000000007)
}
```
