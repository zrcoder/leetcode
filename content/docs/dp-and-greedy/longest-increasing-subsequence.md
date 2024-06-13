---
title: 300. 最长递增子序列
date: 2024-02-10T12:39:50+08:00
---

## [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence) (Medium)

给你一个整数数组 `nums` ，找到其中最长严格递增子序列的长度。

**子序列** 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如， `[3,6,2,7]` 是数组 `[0,3,1,6,2,2,7]` 的子序列。

**示例 1：**

```
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

```

**示例 2：**

```
输入：nums = [0,1,0,3,2,3]
输出：4

```

**示例 3：**

```
输入：nums = [7,7,7,7,7,7,7]
输出：1

```

**提示：**

- `1 <= nums.length <= 2500`
- `-10⁴ <= nums[i] <= 10⁴`

**进阶：**

- 你能将算法的时间复杂度降低到 `O(n log(n))` 吗?

## 分析

### 动态规划

最容易想到的做法：

```go
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := math.MinInt32
	for j, v := range nums {
		dp[j] = 1
		for i, u := range nums[:j] {
			if u < v && dp[i]+1 > dp[j] {
				dp[j] = dp[i] + 1
			}
		}
		res = max(res, dp[j])
	}
	return res
}
```

时间复杂度 O(n^2)，空间复杂度 O(n)。

基于该解法，也能容易地得到这个子序列：只需记住得到最长递增子序列的位置，然后向前遍历，依次得到前一个元素：

```go
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := math.MinInt32
    mi := 0 // 记录最长上升子序列的结束位置
	for j, v := range nums {
		dp[j] = 1
		for i, u := range nums[:j] {
			if u < v && dp[i]+1 > dp[j] {
				dp[j] = dp[i] + 1
			}
		}
        if dp[j] > res {
            res = dp[j]
            mi = j // 更新 mi
        }
		res = max(res, dp[j])
	}

    // 根据 mi 获取最终结果
    lis := make([]int, res)
    j := len(lis)-1
    lis[j] = nums[mi]
    pre := mi
    j--
    for i := mi-1; i >= 0 && j >= 0; i-- {
        if nums[i] < nums[pre] && dp[i] == dp[pre]-1 {
            lis[j] = nums[i]
            j--
            pre = i
        }
    }
    fmt.Println(lis)

	return res
}
```

### 贪心

这是今天的主角。尝试贪心地构建结果：

如果要使上升子序列尽可能长，则需要让序列上升得尽可能慢，因此在构建结果的时候，每次在上升子序列最后加上的那个数需要尽可能小。 建立 memo 数组，memo[i]代表长度为 i+1 的递增子序列末尾数字 遍历 nums，对于当前元素： 如果大于结果数组最后元素，直接追加到结果数组最后； 否则，在结果数组里找到第一个不小于当前元素的元素，并将其更新为当前元素。 这里可以用二分法降低复杂度。

以 [2,1,5,3,4,8,9,7] 为例，可以得到 memo 数组为 [1,3,4,7,9]，这表示： 长度为 1 的递增子序列，最佳末尾数字是 1 长度为 2 的递增子序列，最佳末尾数字是 3 长度为 3 的递增子序列，最佳末尾数字是 4 长度为 4 的递增子序列，最佳末尾数字是 7 长度为 5 的递增子序列，最佳末尾数字是 9

可见，memo 数组的长度就是最长递增子序列的长度。

> 实际上，以上做法是一个不完全的耐心排序(patience sorting)。没有完全排序所有元素，而是借助耐心排序的第一部分，得到了最长递增子序列的长度。


```go
func lengthOfLIS(nums []int) int {
	memo := make([]int, len(nums))
	length := 0
	for _, v := range nums {
		j := sort.SearchInts(memo[:length], v)
		memo[j] = v
		if j == length {
			length++
		}
	}
	return length
}

```
