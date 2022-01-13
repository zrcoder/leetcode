---
title: "最大平均值和的分组"
date: 2022-11-28T10:02:57+08:00
---

## [813. 最大平均值和的分组](https://leetcode.cn/problems/largest-sum-of-averages/)

难度中等

给定数组 `nums` 和一个整数 `k` 。我们将给定的数组 `nums` 分成 **最多** `k` 个相邻的非空子数组 。 **分数** 由每个子数组内的平均值的总和构成。

注意我们必须使用 `nums` 数组中的每一个数进行分组，并且分数不一定需要是整数。

返回我们所能得到的最大 **分数** 是多少。答案误差在 `10-6` 内被视为是正确的。

**示例 1:**

**输入:** nums = [9,1,2,3,9], k = 3
**输出:** 20.00000
**解释:**
nums 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
我们也可以把 nums 分成[9, 1], [2], [3, 9].
这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.

**示例 2:**

**输入:** nums = [1,2,3,4,5,6,7], k = 4
**输出:** 20.50000

**提示:**

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 104`
- `1 <= k <= nums.length`

函数签名：

```go
func largestSumOfAverages(nums []int, k int) float64
```

## 分析

这个问题非常类似背包问题，可以用动态规划解决。

### 动态规划

设 `dp[i][j]` 表示仅考虑前`i`个元素，将其分为j组所能得到的最大分数；这样可以比较简单地用动态规划的方式解决问题。

可以借助前缀和数组迅速得到某个子数组的元素和，进而计算平均值。

```go
func largestSumOfAverages(nums []int, k int) float64 {
    n := len(nums)
    
    preSum := make([]int, n+1)
    dp := make([][]float64, n+1)
    for i, v := range nums {
        preSum[i+1] = preSum[i]+v
        dp[i+1] = make([]float64, k+1)
        dp[i+1][1] = float64(preSum[i+1])/float64(i+1)
    }

    for j := 2; j <= k; j++ {
        for i := j; i <= n; i++ {
            for x := j-1; x < i; x++ {
                avg := float64(preSum[i]-preSum[x])/float64(i-x)
                dp[i][j] = math.Max(dp[i][j], dp[x][j-1]+avg)
            }
        }
    }

    return dp[n][k]
}
```

时间复杂度：`O(k*n^2)`，空间复杂度：`O(n*k)`。

注意到状态转移方程，第二维的 j 仅由 j-1 得到，可以优化dp数组为一维。但须注意枚举 i 需要反向来，以防覆盖部分 j-1状态的值。

```go
func largestSumOfAverages(nums []int, k int) float64 {
    n := len(nums)

    preSum := make([]int, n+1)
    dp := make([]float64, n+1)
    for i, v := range nums {
        preSum[i+1] = preSum[i]+v
        dp[i+1] = float64(preSum[i+1])/float64(i+1)
    }

    for j := 2; j <= k; j++ {
        for i := n; i >= j; i-- {
            for x := j-1; x < i; x++ {
                avg := float64(preSum[i]-preSum[x])/float64(i-x)
                dp[i] = math.Max(dp[i], dp[x]+avg)
            }
        }
    }

    return dp[n]
}
```

空间复杂度降到了`O(n)`。
