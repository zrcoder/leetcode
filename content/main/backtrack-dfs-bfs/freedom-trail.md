---
title: "自由之路"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [514. 自由之路](https://leetcode-cn.com/problems/freedom-trail)

`难度困难`

视频游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。

给定一个字符串 **ring**，表示刻在外环上的编码；给定另一个字符串 **key**，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的**最少**步数。

最初，**ring** 的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使 **key** 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 **key** 中的所有字符。

旋转 **ring** 拼出 key 字符 **key[i]** 的阶段中：

1. 您可以将 **ring** 顺时针或逆时针旋转**一个位置**，计为1步。旋转的最终目的是将字符串 **ring** 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 **key[i] 。**

2. 如果字符 **key[i]** 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 **1 步**。按完之后，您可以开始拼写 **key** 的下一个字符（下一阶段）, 直至完成所有拼写。

**示例：**

```
输入: ring = "godding", key = "gd"
输出: 4
解释:
 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 当然, 我们还需要1步进行拼写。
 因此最终的输出是 4。
```

**提示：**

1. **ring** 和 **key** 的字符串长度取值范围均为 1 至 100；
2. 两个字符串中都只有小写字符，并且均可能存在重复字符；
3. 字符串 **key** 一定可以由字符串 **ring** 旋转拼出。

## 分析

如果 ring 中没有重复字母，这个问题将变得非常简单。只需要事先统计出 `ring` 中每个字母的索引，用一个 26 大小的数组 `indices`。用一个变量 `cur` 指向 `ring` 中当前位置，遍历 `key`，对于当前字母 `ch`，计算 `cur` 到达 `indices[ch-'a']` 所需要的最小步数，即 `dist = abs(cur -  indices[ch-'a'])`， 因为时环状，这个距离需要更新为 `min(dist, n - dist)`，其中 `n` 是 `ring` 的长度。

> 问题 [1974. 使用特殊打字机键入单词的最少时间](https://leetcode-cn.com/problems/minimum-time-to-type-word-using-special-typewriter/) 即是这种简单情况。参考解答代码如下：
>
> ```go
> func minTimeToType(word string) int {
>     var c rune = 'a' // 当前字符
>     res := len(word) // 每个字符要打印出来，需要选定后再按一次键
>     for _, v := range word {
>         dist := abs(int(v-c))
>         res += min(dist, 26-dist)
>         c = v
>     }
>     return res
> }
> ```
>
> `abs`, `min` 是求绝对值和最小值的函数，实现在后边。

但是 ring 中有重复字母，看看贪心思路是否可行：
首先 indices 将记录每个字母出现的所有位置
用一个变量 cur 维护 ring 的字母位置，遍历 key 里的每一个字母：
对于当前字母 ch，查找与 cur 距离最近的 ch 的索引，决定 cur 下一次应该指向哪个位置。
因 indices 严格升序，可以用二分法。

这样整个时间复杂度是 O(n * log(m))， n、m 分别是 key 和 ring 的长度。

但是这个策略不完全正确。
比如 ring = “adbcbae", key = "cba"，
当 cur 更新到 'c' 的位置后，现在需要找 'b', 发现有两个位置的 'b' 距离 cur 一样远，这时候就没法确定到底是选哪个位置会比较好，
显然不能随便选一个，如上边的例子，选左边的 'b' 就会导致错误结果

可以看到，只有每次找到的 ch 的位置距离 cur 仅有一个的时候才能使用上边的贪心策略
如果有左右两个都满足，只能穷举这两种选择，比较最终的结果

### 递归穷举，超时

定义 `var dfs func(i, j int) int`，表示从 ring 的索引 i、 key 的索引 j 开始，直到 j 到达 key 末尾，所需要旋转转盘的步数，结果即  `dfs(0, 0) + len(key)`。

> 因 key  里每个字母选定后要按一次键，最后总步数要加上 key 的长度

```go
func findRotateSteps1(ring string, key string) int {
    indices := calIndices(ring)
    // 返回从 ring 的索引 i、 key 的索引 j 开始，直到 j 到达 key 末尾，所需要旋转转盘的步数
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if j == len(key) {
            return 0
        }
        res := math.MaxInt32
        for _, index := range indices[key[j]-'a'] {
            dist := abs(i - index)
            dist = min(dist, len(ring)-dist)
            res = min(res, dist+dfs(index, j+1))
        }
        return res
    }
    return dfs(0, 0) + len(key) // + len(key) ：key中每个字母需要按一下按钮
}
```

需要的辅助函数如下：

```go
func calIndices(s string) [][]int {
    indices := make([][]int, 26)
    for i, v := range s {
        indices[v-'a'] = append(indices[v-'a'], i)
    }
    return indices
}
```

```go
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

```go
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

可以使用备忘录优化以上过程：

```go
func findRotateSteps2(ring string, key string) int {
    indices := calIndices(ring)
    memo := genMemo(len(ring), len(key))
    // dfs 返回从 ring 的索引 i 开始， key 的索引 j 开始，直到 j 到达 key 末尾，所需要旋转转盘的步数
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if j == len(key) {
            return 0
        }
        if memo[i][j] != math.MaxInt32 {
            return memo[i][j]
        }
        for _, index := range indices[key[j]-'a'] {
            dist := abs(i - index)
            dist = min(dist, len(ring)-dist)
            memo[i][j] = min(memo[i][j], dist+dfs(index, j+1))
        }
        return memo[i][j]
    }

    return dfs(0, 0) + len(key) // + len(key) ：key中每个字母需要按一下按钮
}
```

```go
func genMemo(m, n int) [][]int {
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
        for j := range res[i] {
            res[i][j] = math.MaxInt32
        }
    }
    return res
}
```

时空复杂度均为 `O(m*n)`, 其中 `m`、`n` 分别为 `ring` 和 `key` 的大小。

## 拓展

看到以上带备忘录的递归（也称记忆化搜索）解法，可以联想到反过来自底向上动态规划解决。
定义 `dp[i][j]`表示：对于 `key` 中位置 `i` 处的字母 `ch`, `ring` 中从位置 `j` 最短旋转几次到达字母 `ch`。
则 `dp[i]` 可从 `dp[i-1]` 推导出来：

`dp[i][j] = min(dp[i-1][k] + min(abs(j-k), n-abs(j-k)))`, 其中 `j ∈ key[i] 在 ring 中的位置集合`，`k ∈ key[i-1] 在 ring 中的位置集合`, `n 是 ring 的长度`。
注意`i == 0` 时边界问题， 可在循环前先单独确定。

空间复杂度同同记忆化搜索。当然少了递归栈空间开销。
时间复杂度：最坏情况下是 `O(m^2*n)`，即 ring 里含有 m 个相同的 字母。因输入规模小，这个差异不明显。实际还可以想办法优化，这里不再展开。

```go
func findRotateSteps3(ring string, key string) int {
    indices := calIndices(ring)
    dp := genMemo(len(key), len(ring))
    for _, j := range indices[key[0]-'a'] {
        dp[0][j] = min(j, len(ring)-j)
    }
    for i := 1; i < len(key); i++ {
        for _, j := range indices[key[i]-'a'] {
            for _, k := range indices[key[i-1]-'a'] {
                dist := abs(j - k)
                dp[i][j] = min(dp[i][j], dp[i-1][k]+min(dist, len(ring)-dist))
            }
        }
    }
    res := math.MaxInt32
    for _, v := range dp[len(key)-1] {
        res = min(res, v)
    }
    return res + len(key)
}
```

注意到 `dp[i]` 只和`dp[i-1]` 有关，可以把 `dp` 数组减成两行轮换使用，空间复杂度降低为 `O(n)`。

```go
func findRotateSteps(ring string, key string) int {
    indices := calIndices(ring)
    dp := genMemo(2, len(ring))
    for _, j := range indices[key[0]-'a'] {
        dp[0][j] = min(j, len(ring)-j)
    }
    lastDp, curDp := dp[0], dp[1]
    for i := 1; i < len(key); i++ {
        for _, j := range indices[key[i]-'a'] {
            for _, k := range indices[key[i-1]-'a'] {
                dist := abs(j - k)
                curDp[j] = min(curDp[j], lastDp[k]+min(dist, len(ring)-dist))
            }
        }
        lastDp, curDp = curDp, lastDp
        for i := range curDp {
            curDp[i] = math.MaxInt32
        }
    }
    res := math.MaxInt32
    for _, v := range lastDp {
        res = min(res, v)
    }
    return res + len(key)
}
```
