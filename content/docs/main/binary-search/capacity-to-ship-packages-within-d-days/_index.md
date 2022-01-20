---
title: "1011. 在 D 天内送达包裹的能力"
date: 2021-04-27T08:08:17+08:00
weight: 50
tags: [贪心, 二分搜素]
---

## [1011. 在 D 天内送达包裹的能力](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)

传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。

传送带上的第 `i` 个包裹的重量为 `weights[i]`。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 `D` 天内将传送带上的所有包裹送达的船的最低运载能力。

**示例 1：**

```
输入：weights = [1,2,3,4,5,6,7,8,9,10], D = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。 
```

**示例 2：**

```
输入：weights = [3,2,2,4,1,4], D = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4
```

**示例 3：**

```
输入：weights = [1,2,3,1,1], D = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1
```

**提示：**

1. `1 <= D <= weights.length <= 50000`
2. `1 <= weights[i] <= 500`

函数签名：

```go
func shipWithinDays(weights []int, D int) int
```

## 分析

### 回溯

可以尝试在数组里插入 `D-1` 个隔板，在分隔开的 `D` 组内得到总和最大的和作为一个参考结果。这需要尝试所有排列，最后取每种尝试的最小值。复杂度为 `O(A(n-1, D))`，即在 `n-1`个空里插入 `D` 个隔板的排列数，复杂度太高。

### 二分搜索+贪心

首先可以确定答案的大体范围，在 `[max, sum]`区间内，即数组最大元素值和所有元素总和之间。这样可以枚举区间里每一个值 `limit`，看看载重量为 `limit` 时能否在 `D` 天内运完货物，这可以用贪心策略在线性时间内完成。

更近一步，不用从 `max` 向 `sum`一一枚举，而是用二分法枚举。

```go
func shipWithinDays(weights []int, D int) int {
    lo, hi := 0, 0
    for _, v := range weights {
        lo = max(lo, v)
        hi += v
    }
    hi++
    for lo < hi {
        mid := lo + (hi-lo)/2
        if !check(mid, D, weights) {
            lo = mid+1
        } else {
            hi = mid
        }
    }
    return lo
}
func check(limit, D int, weights []int) bool {
    cnt, sum := 0, 0
        for _, v := range weights {
            if sum+v > limit {
                cnt++
                sum = v
            } else {
                sum += v
            }
            if cnt > D {
                break
            }
        }
        cnt++
        return cnt <= D
}
```

时间复杂度 `O(nlogs)`，其中 `n` 是数组长度，`s`是数组总和与最大元素的差值。

空间复杂度 `O(1)`。