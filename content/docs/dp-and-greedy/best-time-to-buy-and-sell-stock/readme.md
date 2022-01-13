---
title: "股票收益问题"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

LeetCode 上有六道股票收益类问题，给定一只股票每天的价格，计算在一定限制下的最大收益。

问题共性：

仅有一支股票，且每天手上最多持有一股。

不同之处：

每个问题对交易次数有限制，或者附加了交易冷冻期，或增加了交易费用。

让我们从最具一般性的问题开始

## [188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv)

这个问题给出了交易次数的上限 k

## 分析
类似背包问题，可以分析天数、完成的交易数和是否持有股票这三个状态，找到其转移方程。

定义dp[i][k][s]表示第i天，交易了k次，持有或不持有股票所能得到的最大收益；为了方便处理边界, dp 大小为 (n+1) * (k+1) * 2

假设共n天，最大交易次数为k，则所有的状态共（n * k * 2）种

```
for 1 <= i <= n {
    for 1 <= j <= k {
        for 0 <= s <= 1 {
            dp[i][j][s] = max(buy, sell, rest) // 买入、卖出或啥也不做
```
最终的答案就是dp[n][k][0]，即最后一天且不持有股票所获得的收益，注意k取最大值  
因为一直都是同一支股票，dp的递推关系为：
```
第i天不持有股票的情况有两种：前一天没有股票，今天不买；或前一天有股票，今天卖了；选择收益最大的做法即可

dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + price[i])

第i天持有股票的情况，与上边类似: 前一天有股票， 今天不卖，或前一天没有股票，今天买入； 买入则收益减少 prices[i]

dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])   
```
注意买入我们会将j减去1，卖出则不会；改成卖出时j-1，而买入时不变也是可以的，只是初始情况略有不同

那么初始情况如何确定？

考虑还未开始交易的情况确定初始值  
```
// i从1开始，0代表还没开始
dp[0][...][0] = 0 
dp[0][...][1] = -infinity
```
k==0时会怎么样？意味着不允许交易
```
dp[...][0][0] = 0
dp[...][0][1] = -infinity
``` 
最后要注意 k 的取值，实际上最多进行 n/2 次交易，如果 k 比这个值大，相当于可以有无限次操作，状态转移方程里 k 的维度可以忽略，或者重新赋值 k 为 n/2 即可

参考实现1：
```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		k = n/2
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			if i == 0 || j == 0 {
				dp[i][j][1] = math.MinInt32
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}
```

如果定义买入后卖出算一笔交易：状态转移方程与上边稍有差别，初始条件要注意j为0的情况：
    
dp[i][0][1] = max(dp[i-1][0][1], -prices[i-1]) (特别地，dp[-1][0][1] = -prices[0])

参考实现2：
```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		k = n / 2
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			if i == 0 {
				dp[i][j][1] = math.MinInt32
			}
			if j == 0 {
				if i > 0 {
                    dp[i][j][1] = max(dp[i-1][j][1], -prices[i-1])					
				} else {
					dp[i][j][1] = -prices[0]
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
    		for j := 1; j <= k; j++ {
    			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i-1])
    			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i-1])
    		}
    	}
    	return dp[n][k][0]
}
```

以上解法的时间、空间复杂度都是 O(n*k)

空间优化：

首先可以用两个数组拆分 dp 数组，分别记录当前持有股票和不持有股票所能获得的最大收益

其次，因状态只跟前一天的状态有关，天数那个维度可以取消，最终将空间复杂度优化为 O(k)

参考实现3：
```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		k = n / 2
	}
	hold := make([]int, k+1)
	release := make([]int, k+1)
	// 买入记为完成一次交易
	for i := range hold {
		hold[i] = math.MinInt32
	}
	for _, price := range prices {
		for j := 1; j <= k; j++ {
			hold[j], release[j] = max(hold[j], release[j-1]-price), max(release[j], hold[j]+price)
		}
	}
	return release[k]
}
```

如果将卖出看成是完成一笔交易，则代码调整为参考实现4：
```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		k = n / 2
	}
	hold := make([]int, k+1)
	release := make([]int, k+1)
	// 卖出记为完成一次交易
	for j := range hold {
		hold[j] = math.MinInt32
	}
	for i, price := range prices {
		if i == 0 {
			hold[0] = -price
		} else {
			hold[0] = max(hold[0], -prices[i-1])
		}
		for j := 1; j <= k; j++ {
			release[j], hold[j] = max(release[j], hold[j]+price), max(hold[j], release[j-1]-price)
		}
	}
	return release[k]
}
```
## [123. 买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii)
这里限定 k 为 2，其他限制同上一个问题，直接套用上一问题的解法即可，如参考实现3：
```go
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	const k = 2
	hold := make([]int, k+1)
	release := make([]int, k+1)
	// 买入记为完成一次交易
	for i := range hold {
		hold[i] = math.MinInt32
	}
	for _, price := range prices {
		for j := 1; j <= k; j++ {
			hold[j], release[j] = max(hold[j], release[j-1]-price), max(release[j], hold[j]+price)
		}
	}
	return release[k]
}
```
## [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock)
这里限定 k 为 1，可直接套用上边的解法

```go
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	release, hold := 0, -prices[0]
	for _, price := range prices {
		release = max(release, hold+price)
		hold = max(hold, -price)
	}
	return release
}
```

## [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii)
这里对交易次数没有限定

```go
func maxProfit(prices []int) int {
	release, hold := 0, math.MinInt32
	for _, v := range prices {
		release, hold = max(release, hold+v), max(hold, release-v)
	}
	return release
}
```

## [309. 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown)
同样对交易次数没有限定，但是加了一个限定：

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

只需要增加一个变量记录上一次不持有股票的最大收益：

```go
func maxProfit(prices []int) int {
	release, hold := 0, math.MinInt32
	lastRelease := release
	for _, price := range prices {
		release, hold, lastRelease =
			max(release, hold+price),
			max(hold, lastRelease-price),
			release
	}
	return release
}
```

## [714. 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)
同样对交易次数没有限定，但是加了一个限定: 完成一笔交易需要付出额外费用。

```go
func maxProfit(prices []int, fee int) int {
	release, hold := 0, math.MinInt32
	for _, price := range prices {
		release, hold = max(release, hold+price), max(hold, release-price-fee)
	}
	return release
}
```

## 参考
[一个通用的方法团灭6道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-wen/)
