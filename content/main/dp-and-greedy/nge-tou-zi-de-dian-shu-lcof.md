---
title: "剑指 Offer 60. n个骰子的点数  LCOF"
date: 2023-01-03T23:32:12+08:00
---

## [剑指 Offer 60. n个骰子的点数  LCOF](https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof) (Medium)

把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

**示例 1:**

```
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

```

**示例 2:**

```
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
```

**限制：**

`1 <= n <= 11`

函数签名：

```go
func dicesProbability(n int) []float64
```

## 分析

李白没有为黄鹤楼赋诗，只因已经有人写得非常好了：
> 眼前有景道不得，崔颢题诗在上头。

这里有一篇[很棒的题解](https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/solution/jian-zhi-offer-60-n-ge-tou-zi-de-dian-sh-z36d)，可以直接移步～

不过题解里给出的代码不够完美，数组空间开辟多了。这里直接给出优化后的代码。

### 动态规划

```go
func dicesProbability(n int) []float64 {
    arr1 := make([]float64, 5*n+1)
    arr2 := make([]float64, 5*n+1)
	dp := arr1[:6]
    r := 1.0/6.0
    for i := range dp {
        dp[i] = r
    }
    for i := 2; i <= n; i++ {
        next := arr2[:5*i+1]
        reset(next)
        for j, v := range dp {
            for k := 0; k < 6; k++ {
                next[j+k] += v/6.0
            }
        }
        dp = next
        arr1, arr2 = arr2, arr1
    }
    return dp
}

func reset(arr []float64) {
    for i := range arr {
        arr[i] = 0
    }
}
```

时间复杂度：`O(n^2)`，空间复杂度：`O(n)`。

附崔颢《黄鹤楼》：

```text
昔人已乘黄鹤去，此地空余黄鹤楼。
黄鹤一去不复返，白云千载空悠悠。
晴川历历汉阳树，芳草萋萋鹦鹉洲。
日暮乡关何处是，烟波江上使人愁。
```

