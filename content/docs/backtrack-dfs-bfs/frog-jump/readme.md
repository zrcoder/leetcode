---
title: "403. 青蛙过河"
date: 2021-05-01T11:03:06+08:00
weight: 50
tags: [记忆化搜索, 动态规划]
---

## [403. 青蛙过河](https://leetcode-cn.com/problems/frog-jump/)

一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。

给你石子的位置列表 `stones`（用单元格序号 **升序** 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。

开始时， 青蛙默认已站在第一块石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格 1 跳至单元格 2 ）。

如果青蛙上一步跳跃了 `k` 个单位，那么它接下来的跳跃距离只能选择为 `k - 1`、`k` 或 `k + 1` 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。

**示例 1：**

```
输入：stones = [0,1,3,5,6,8,12,17]
输出：true
解释：青蛙可以成功过河，按照如下方案跳跃：跳 1 个单位到第 2 块石子, 然后跳 2 个单位到第 3 块石子, 接着 跳 2 个单位到第 4 块石子, 然后跳 3 个单位到第 6 块石子, 跳 4 个单位到第 7 块石子, 最后，跳 5 个单位到第 8 个石子（即最后一块石子）。
```

**示例 2：**

```
输入：stones = [0,1,2,3,4,8,9,11]
输出：false
解释：这是因为第 5 和第 6 个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。
```

**提示：**

- `2 <= stones.length <= 2000`
- `0 <= stones[i] <= 231 - 1`
- `stones[0] == 0`

函数签名：

```go
func canCross(stones []int) bool
```

## 分析

### DFS 穷举

假设青蛙在某一时刻跳了 `k` 步到了 `i` 处，那么下一步就有三种跳法：跳 `k-1`，`k`, `k+1` 步。这样可以很容易想到 DFS 解法。

```go
func canCross(stones []int) bool {
    n := len(stones)
    isStone := make(map[int]bool, n)
    for _, v := range stones {
        isStone[v] = true
    }

    var dfs func(i, k int) bool
    dfs = func(i, k int) bool {
        if i == stones[n-1] {
            return true
        }
        for j := k-1; j <= k+1; j++ {
            next := i+j
            if next > i && isStone[next] && dfs(next, j) {
                return true
            }
        }
        return false
    }
    
    return dfs(0, 0)
}
```

状态太多，超时。

时间复杂度是 `O(3^n)`，其中 n 是石子的数量，题目上限为 2000，指数级的复杂度。

### 记忆化搜索

朴素 DFS 有重复计算，所以复杂度高。可以加上备忘录优化，得到记忆化搜索解法。

```go
func canCross(stones []int) bool {
    n := len(stones)
    isStone := make(map[int]bool, n)
    dp := make(map[int]map[int]bool, n)
    for _, v := range stones {
        isStone[v] = true
        dp[v] = map[int]bool{}
    }

    var dfs func(i, k int) bool
    dfs = func(i, k int) bool {
        if i == stones[n-1] {
            return true
        }
        if res, ok := dp[i][k]; ok {
            return res
        }
        for j := k-1; j <= k+1; j++ {
            next := i+j
            if next > i && isStone[next] && dfs(next, j) {
                dp[i][k] = true
                return true
            }
        }
        dp[i][k] = false
        return false
    }
    
    return dfs(0, 0)
}
```

时间复杂度 `O(n^2)`，空间复杂度 `O(n^2)` 。

### 动态规划

直接上代码，有些小优化。

```go
func canCross(stones []int) bool {
    n := len(stones)
    for i := 1; i < n; i++ {
        if stones[i]-stones[i-1] > i {
            return false
        }
    }
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    dp[0][0] = true
    for i := 1; i < n; i++ {
        for j := i - 1; j >= 0; j-- {
            k := stones[i] - stones[j]
            if k > j+1 {
                break
            }
            dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
            if i == n-1 && dp[i][k] {
                return true
            }
        }
    }
    return false
}
```

时空复杂度都是 `O(n^2)`，同记忆化搜索。
