---
title: "1482. 制作 m 束花所需的最少天数"
date: 2021-05-09T09:18:12+08:00
weight: 50
tags: [二分搜索]
---

## [1482. 制作 m 束花所需的最少天数](https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/)

难度中等

给你一个整数数组 `bloomDay`，以及两个整数 `m` 和 `k` 。

现需要制作 `m` 束花。制作花束时，需要使用花园中 **相邻的 `k` 朵花** 。

花园中有 `n` 朵花，第 `i` 朵花会在 `bloomDay[i]` 时盛开，**恰好** 可以用于 **一束** 花中。

请你返回从花园中摘 `m` 束花需要等待的最少的天数。如果不能摘到 `m`束花则返回 **-1** 。

**示例 1：**

```
输入：bloomDay = [1,10,3,10,2], m = 3, k = 1
输出：3
解释：让我们一起观察这三天的花开过程，x 表示花开，而 _ 表示花还未开。
现在需要制作 3 束花，每束只需要 1 朵。
1 天后：[x, _, _, _, _]   // 只能制作 1 束花
2 天后：[x, _, _, _, x]   // 只能制作 2 束花
3 天后：[x, _, x, _, x]   // 可以制作 3 束花，答案为 3
```

**示例 2：**

```
输入：bloomDay = [1,10,3,10,2], m = 3, k = 2
输出：-1
解释：要制作 3 束花，每束需要 2 朵花，也就是一共需要 6 朵花。而花园中只有 5 朵花，无法满足制作要求，返回 -1 。
```

**示例 3：**

```
输入：bloomDay = [7,7,7,7,12,7,7], m = 2, k = 3
输出：12
解释：要制作 2 束花，每束需要 3 朵。
花园在 7 天后和 12 天后的情况如下：
7 天后：[x, x, x, x, _, x, x]
可以用前 3 朵盛开的花制作第一束花。但不能使用后 3 朵盛开的花，因为它们不相邻。
12 天后：[x, x, x, x, x, x, x]
显然，我们可以用不同的方式制作两束花。
```

**示例 4：**

```
输入：bloomDay = [1000000000,1000000000], m = 1, k = 1
输出：1000000000
解释：需要等 1000000000 天才能采到花来制作花束
```

**示例 5：**

```
输入：bloomDay = [1,10,2,9,3,8,4,7,5,6], m = 4, k = 2
输出：9
```

**提示：**

- `bloomDay.length == n`
- `1 <= n <= 10^5`
- `1 <= bloomDay[i] <= 10^9`
- `1 <= m <= 10^6`
- `1 <= k <= n`

函数签名：

```go
func minDays(bloomDay []int, m int, k int) int
```

## 分析

### 二分搜索

首先，如果给定一个天数 day，可以遍历一次数组判断是否能在这天长出满足题意的花。

其次，结果显然在数组中的最小天数和最大天数之间，即闭区间`[minDay, maxDay]`，且在这个区间从小到大依次取一个数字作为上边的 day，这样肯定在某个取值能得到需要的花。因为满足单调性，用二分法。

> 花园里最多能长出的花数就是数组长度 `n`，如果需要的花 `m*k` 比能提供的总数 `n` 还多可以直接返回 `-1`。

```go
func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)
    if m*k > n {
        return -1
    }
    // 检查day天后是否能开出 m 束花 
    check := func(day int) bool {
        curM := 0
        curK := 0
        for _, v := range bloomDay {
            if v > day {
                curK = 0
                continue
            }
            curK++
            if curK == k {
                curM++
                curK = 0
            }
            if curM == m {
                return true
            }
        }
        return false
    }
    maxDay, minDay := 0, math.MaxInt32
    for _, v := range bloomDay {
        maxDay = max(maxDay, v)
        minDay = min(minDay, v)
    }
    lo, hi := minDay, maxDay+1
    for lo < hi {
        mid := lo+(hi-lo)/2
        if !check(mid) {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    return lo
}
```

时间复杂度 `O(nlogn)`，空间复杂度 `O(1)`。