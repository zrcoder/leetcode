---
title: "石子游戏"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

石子游戏非常有意思，可以和小朋友玩。以下几个石子游戏问题，用程序来做计算也非常有意思。

## [877. 石子游戏](https://leetcode-cn.com/problems/stone-game/)

难度中等

亚历克斯和李用几堆石子在做游戏。偶数堆石子**排成一行**，每堆都有正整数颗石子 `piles[i]` 。

游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。

亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。

假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 `true` ，当李赢得比赛时返回 `false` 。

 

**示例：**

```
输入：[5,3,4,5]
输出：true
解释：
亚历克斯先开始，只能拿前 5 颗或后 5 颗石子 。
假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
如果李拿走前 3 颗，那么剩下的是 [4,5]，亚历克斯拿走后 5 颗赢得 10 分。
如果李拿走后 5 颗，那么剩下的是 [3,4]，亚历克斯拿走后 4 颗赢得 9 分。
这表明，取前 5 颗石子对亚历克斯来说是一个胜利的举动，所以我们返回 true 。
```

 

**提示：**

1. `2 <= piles.length <= 500`
2. `piles.length` 是偶数。
3. `1 <= piles[i] <= 500`
4. `sum(piles)` 是奇数。

函数签名：

```go
func stoneGame(piles []int) bool
```

## 分析

### 区间 dp

开辟二维 `n*n` 数组 `dp`，`dp[start][end]` 代表在区间 `[start, end]` 上取石子，先手所得减去后手所得的最大值。

为了先计算出子区间的结果，遍历时起点从后向前枚举。这是区间 `dp` 的常规做法。

```go
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for start := n - 2; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			// dp[start][end] 代表在区间 [start, end] 上取石子，先手所得减去后手所得的最大值
			dp[start][end] = max(piles[start]-dp[start+1][end], piles[end]-dp[start][end-1])
		}
	}
	return dp[0][n-1] > 0
}
```

空间可以优化：

```go
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = piles[i]
	}
	for start := n - 2; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			dp[end] = max(piles[start]-dp[end], piles[end]-dp[end-1])
		}
	}
	return dp[n-1] > 0
}
```

时间复杂度 `O(n^2)`，空间复杂度 `O(n)`。

### 贪心

可以把所有的石堆分成两类，索引为奇数的石堆和索引为偶数的石堆，显然这两类石堆数量各是总石堆数的一半。

又根据题目限制，这两类石堆石子总数肯定不相等。

先手能不能控制只拿奇数石堆或偶数石堆呢？完全可以。

比如认死了就要拿偶数石堆，一开始就拿石堆 0， 对方怎么拿都只能拿到奇数石堆，且不管对方拿的是前边一堆还是后边一堆，下一次只要拿对方刚才拿过的石堆的邻居石堆就又是偶数石堆了。

正因为先手完全可以控制自己拿所有的奇数石堆还是所有的偶数石堆，对手一点办法也没有。所以先手可以在游戏开始前计算一下奇数石堆和偶数石堆的总石子数，然后自己一直拿总数大的那一类石堆就行了。

无论如何，先手都能赢。

```go
func stoneGame(piles []int) bool {
    return true
}
```

时空复杂度都是 `O(1)`。

## [1140. 石子游戏 II](https://leetcode-cn.com/problems/stone-game-ii/)

难度中等

亚历克斯和李继续他们的石子游戏。许多堆石子 **排成一行**，每堆都有正整数颗石子 `piles[i]`。游戏以谁手中的石子最多来决出胜负。

亚历克斯和李轮流进行，亚历克斯先开始。最初，`M = 1`。

在每个玩家的回合中，该玩家可以拿走剩下的 **前** `X` 堆的所有石子，其中 `1 <= X <= 2M`。然后，令 `M = max(M, X)`。

游戏一直持续到所有石子都被拿走。

假设亚历克斯和李都发挥出最佳水平，返回亚历克斯可以得到的最大数量的石头。

 

**示例：**

```
输入：piles = [2,7,9,4,4]
输出：10
解释：
如果亚历克斯在开始时拿走一堆石子，李拿走两堆，接着亚历克斯也拿走两堆。在这种情况下，亚历克斯可以拿到 2 + 4 + 4 = 10 颗石子。 
如果亚历克斯在开始时拿走两堆石子，那么李就可以拿走剩下全部三堆石子。在这种情况下，亚历克斯可以拿到 2 + 7 = 9 颗石子。
所以我们返回更大的 10。 
```

 

**提示：**

- `1 <= piles.length <= 100`
- `1 <= piles[i] <= 10 ^ 4`

函数签名：

```go
func stoneGameII(piles []int) int
```

## 分析

两个人不断拿走前边的石堆，逆向考虑下，接近最后的情况就是剩余一两堆石头，这样的情况比较好处理。

可以定义 help(i, m) 函数，代表只剩下 piles[i:] 的石堆，且一次性可拿走石头堆数 x 满足 1 <= x <= 2m 时首先行动的人能得到的最大石子数。

如果 i + 2*m >= n (其中 n 代表 piles 长度)，表示一次性可以拿光，结果显然是 sum(piles[i:]);

否则， 需要遍历 x，求剩余的最小 help(i + x, max(x, M))，也就是自己拿 x 的时候，对手拿的石子最少。

注意这里递推，`i+x` 可能超过 `n` ，这也是为什么要判断 `i + 2*m >= n` 的情况。

### 朴素递归
```go
func stoneGameII(piles []int) int {
	n := len(piles)
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i + 2*m >= n {
			return sum(piles[i:])
		}
		res := 0
        for x := 1; x <= 2*m; x++ {
			res = max(res, sum(piles[i:]) - dfs(i+x, max(x, m)))
		}
		return res
	}
	return dfs(0, 1)
}
```
### 记忆化搜索
上边会超时，可以有两方面的优化：

加上备忘录，优化重复的递归子问题；

sum(piles[i:]) 求和可以用后缀和数组优化。

```go
func stoneGameII(piles []int) int {
	n := len(piles)
	suffixSum := make([]int, n)
	suffixSum[n-1] = piles[n-1]
	for i := n-2; i >= 0; i-- {
		suffixSum[i] = piles[i] + suffixSum[i+1]
	}
	memo := make([][]int, n)
	for i := range  memo {
		memo[i] = make([]int, n+1)
	}
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i + 2*m >= n {
			return suffixSum[i]
		}
		if memo[i][m] > 0 {
			return memo[i][m]
		}
		res := 0
		for x := 1; x <= 2*m; x++ {
			res = max(res, suffixSum[i] - dfs(i+x, max(x, m)))
		}
		memo[i][m] = res
		return res
	}
	return dfs(0, 1)
}
```
### 区间 dp
从以上记忆化搜索解法，也比较容易想到正向动态规划。注意 i 要从右向左枚举，这时为了先把所有的子区间 dp 值确定：

```go
func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	suffixSum := 0
	for i := n - 1; i >= 0; i-- {
		suffixSum += piles[i]
		for m := 1; m <= n; m++ {
			if i+2*m >= n {
				dp[i][m] = suffixSum
				continue
			}
			for x := 1; x <= 2*m; x++ {
				dp[i][m] = max(dp[i][m], suffixSum-dp[i+x][max(m, x)])
			}
		}
	}
	return dp[0][1]
}
```

## [1406. 石子游戏 III](https://leetcode-cn.com/problems/stone-game-iii/)

难度困难

Alice 和 Bob 用几堆石子在做游戏。几堆石子排成一行，每堆石子都对应一个得分，由数组 `stoneValue` 给出。

Alice 和 Bob 轮流取石子，**Alice** 总是先开始。在每个玩家的回合中，该玩家可以拿走剩下石子中的的前 **1、2 或 3 堆石子** 。比赛一直持续到所有石头都被拿走。

每个玩家的最终得分为他所拿到的每堆石子的对应得分之和。每个玩家的初始分数都是 **0** 。比赛的目标是决出最高分，得分最高的选手将会赢得比赛，比赛也可能会出现平局。

假设 Alice 和 Bob 都采取 **最优策略** 。如果 Alice 赢了就返回 *"Alice"* *，*Bob 赢了就返回 *"Bob"，*平局（分数相同）返回 *"Tie"* 。

**示例 1：**

```
输入：values = [1,2,3,7]
输出："Bob"
解释：Alice 总是会输，她的最佳选择是拿走前三堆，得分变成 6 。但是 Bob 的得分为 7，Bob 获胜。
```

**示例 2：**

```
输入：values = [1,2,3,-9]
输出："Alice"
解释：Alice 要想获胜就必须在第一个回合拿走前三堆石子，给 Bob 留下负分。
如果 Alice 只拿走第一堆，那么她的得分为 1，接下来 Bob 拿走第二、三堆，得分为 5 。之后 Alice 只能拿到分数 -9 的石子堆，输掉比赛。
如果 Alice 拿走前两堆，那么她的得分为 3，接下来 Bob 拿走第三堆，得分为 3 。之后 Alice 只能拿到分数 -9 的石子堆，同样会输掉比赛。
注意，他们都应该采取 最优策略 ，所以在这里 Alice 将选择能够使她获胜的方案。
```

**示例 3：**

```
输入：values = [1,2,3,6]
输出："Tie"
解释：Alice 无法赢得比赛。如果她决定选择前三堆，她可以以平局结束比赛，否则她就会输。
```

**示例 4：**

```
输入：values = [1,2,3,-1,-2,-3,7]
输出："Alice"
```

**示例 5：**

```
输入：values = [-1,-2,-3]
输出："Tie"
```

**提示：**

- `1 <= values.length <= 50000`
- `-1000 <= values[i] <= 1000`

函数签名：

```go
func stoneGameIII(stoneValue []int) string
```

## 分析

思路同上一问题，更简单：

```go
func stoneGameIII(stoneValue []int) string {
    n := len(stoneValue)
    if n == 0 {
        return "Tie"
    }
    dp := make([]int, n+1)
    sum := 0
    for i := n-1; i >= 0; i-- {
        val := stoneValue[i]
        sum += val
        // 对方的得分
        point := math.MaxInt32
        for j := i+1; j <= i+3 && j <= n; j++ {
            point = min(point, dp[j])
        }
        dp[i] = sum - point
    }
    alice := dp[0]
    bob := sum - alice
    if alice > bob {
        return "Alice"
    }
    if alice < bob {
        return "Bob"
    }
    return "Tie"
}
```

## [1510. 石子游戏 IV](https://leetcode-cn.com/problems/stone-game-iv/)

难度困难

Alice 和 Bob 两个人轮流玩一个游戏，Alice 先手。

一开始，有 `n` 个石子堆在一起。每个人轮流操作，正在操作的玩家可以从石子堆里拿走 **任意** 非零 **平方数** 个石子。

如果石子堆里没有石子了，则无法操作的玩家输掉游戏。

给你正整数 `n` ，且已知两个人都采取最优策略。如果 Alice 会赢得比赛，那么返回 `True` ，否则返回 `False` 。

 

**示例 1：**

```
输入：n = 1
输出：true
解释：Alice 拿走 1 个石子并赢得胜利，因为 Bob 无法进行任何操作。
```

**示例 2：**

```
输入：n = 2
输出：false
解释：Alice 只能拿走 1 个石子，然后 Bob 拿走最后一个石子并赢得胜利（2 -> 1 -> 0）。
```

**示例 3：**

```
输入：n = 4
输出：true
解释：n 已经是一个平方数，Alice 可以一次全拿掉 4 个石子并赢得胜利（4 -> 0）。
```

**示例 4：**

```
输入：n = 7
输出：false
解释：当 Bob 采取最优策略时，Alice 无法赢得比赛。
如果 Alice 一开始拿走 4 个石子， Bob 会拿走 1 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7 -> 3 -> 2 -> 1 -> 0）。
如果 Alice 一开始拿走 1 个石子， Bob 会拿走 4 个石子，然后 Alice 只能拿走 1 个石子，Bob 拿走最后一个石子并赢得胜利（7 -> 6 -> 2 -> 1 -> 0）。
```

**示例 5：**

```
输入：n = 17
输出：false
解释：如果 Bob 采取最优策略，Alice 无法赢得胜利。
```

 

**提示：**

- `1 <= n <= 10^5`

函数签名：

```go
func winnerSquareGame(n int) bool
```

##  分析

### 朴素实现

```go
func winnerSquareGame(n int) bool {
	if n <= 0 {
		return true
	}
	return help(n)
}

func help(n int) bool {
	if n == 0 {
		return true
	}
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
		if !help(n - i * i) {
			return true
		}
	}
	return false
}
```

### 记忆化搜索

朴素实现重复计算较多，会超时，可以加入备忘录，得到记忆化搜索方法

```go
func winnerSquareGame(n int) bool {
	if n <= 0 {
		return true
	}
	memo := make([]int, n+1)
	return help(n, memo)
}

func help(n int, memo []int) bool {
	if n == 0 {
		return true
	}
	if memo[n] != 0 {
		return memo[n] == 1
	}
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			memo[n] = 1
			return true
		}
		if !help(n - i * i, memo) {
			memo[n] = 1
			return true
		}
	}
	memo[n] = 2
	return false
}
```

时间复杂度 O(n*sqrt(n))，空间复杂度 O(n)

### 动态规划

上边是自顶向下的递归解法，也可以自底向上动态规划解决。

```go
func winnerSquareGame(n int) bool {
	if n <= 0 {
		return true
	}
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp[i-j*j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
```

时空复杂度同记忆化搜索解法。

## [1563. 石子游戏 V](https://leetcode-cn.com/problems/stone-game-v/)

难度困难

几块石子 **排成一行** ，每块石子都有一个关联值，关联值为整数，由数组 `stoneValue` 给出。

游戏中的每一轮：Alice 会将这行石子分成两个 **非空行**（即，左侧行和右侧行）；Bob 负责计算每一行的值，即此行中所有石子的值的总和。Bob 会丢弃值最大的行，Alice 的得分为剩下那行的值（每轮累加）。如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。

只 **剩下一块石子** 时，游戏结束。Alice 的分数最初为 **`0`** 。

返回 **Alice 能够获得的最大分数** *。*

 

**示例 1：**

```
输入：stoneValue = [6,2,3,4,5,5]
输出：18
解释：在第一轮中，Alice 将行划分为 [6，2，3]，[4，5，5] 。左行的值是 11 ，右行的值是 14 。Bob 丢弃了右行，Alice 的分数现在是 11 。
在第二轮中，Alice 将行分成 [6]，[2，3] 。这一次 Bob 扔掉了左行，Alice 的分数变成了 16（11 + 5）。
最后一轮 Alice 只能将行分成 [2]，[3] 。Bob 扔掉右行，Alice 的分数现在是 18（16 + 2）。游戏结束，因为这行只剩下一块石头了。
```

**示例 2：**

```
输入：stoneValue = [7,7,7,7,7,7,7]
输出：28
```

**示例 3：**

```
输入：stoneValue = [4]
输出：0
```

 

**提示：**

- `1 <= stoneValue.length <= 500`
- `1 <= stoneValue[i] <= 10^6`

函数签名：

```go
func stoneGameV(stoneValue []int) int
```

## 分析

### 朴素实现

先写一下朴素递归实现：

```go
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	var cal func(start, end int) int
	cal = func(start, end int) int {
		if start == end {
			return 0
		}
		sum := 0
		for i := start; i <= end; i++ {
			sum += stoneValue[i]
		}
		sumLeft := 0
		res := 0
		// i 代表分组的位置，索引 i 之前包括索引 i 处为左半部分，索引 i 之后为右半部分
		for i := start; i < end; i++ {
			sumLeft += stoneValue[i]
			sumRight := sum - sumLeft
			if sumLeft < sumRight {
				res = max(res, sumLeft + cal(start, i))
			} else if sumLeft > sumRight {
				res = max(res, sumRight + cal(i+1, end))
			} else {
				res = max(res, sumLeft + max(cal(start, i), cal(i+1, end)))
			}
		}
		return res
	}
	return cal(0, n-1)
}
```

重复递归太多，超时了。

### 记忆化搜索

加上备忘录， 同时用前缀和技巧优化计算子序列和的复杂度：

```go
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	prefixSum := make([]int, n+1)
	for i, v := range stoneValue {
		prefixSum[i+1] = prefixSum[i] + v
	}

	var cal func(start, end int) int
	cal = func(start, end int) int {
		if start == end {
			return 0
		}
		if dp[start][end] > 0 {
			return dp[start][end]
		}
		sum := prefixSum[end+1] - prefixSum[start]
		sumLeft := 0
		// i 代表分组的位置，索引 i 之前包括索引 i 处为左半部分，索引 i 之后为右半部分
		for i := start; i < end; i++ {
			sumLeft += stoneValue[i]
			sumRight := sum - sumLeft
			if sumLeft < sumRight {
				dp[start][end] = max(dp[start][end], sumLeft + cal(start, i))
			} else if sumLeft > sumRight {
				dp[start][end] = max(dp[start][end], sumRight + cal(i+1, end))
			} else {
				dp[start][end] = max(dp[start][end], sumLeft + max(cal(start, i), cal(i+1, end)))
			}
		}
		return dp[start][end]
	}
	return cal(0, n-1)
}
```

时间复杂度不太好分析，理论上与下边动态规划的复杂度一致，为 O(n^3)，空间复杂度 O(n^2)

### 动态规划

自底向上写出动态规划解法，需要注意枚举顺序；

起点从后向前枚举，终点反方向枚举，这样的顺序可以保证父问题依赖的子问题已经解决。

```go
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	prefixSum := make([]int, n+1)
	for i, v := range stoneValue {
		prefixSum[i+1] = prefixSum[i] + v
	}
	for start := n - 2; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			// i 代表分组的位置，索引 i 之前包括索引 i 处为左半部分，索引 i 之后为右半部分
			for i := start; i < end; i++ {
				sumLeft := prefixSum[i+1] - prefixSum[start]
				sumRight := prefixSum[end+1] - prefixSum[i+1]
				if sumLeft < sumRight {
					dp[start][end] = max(dp[start][end], sumLeft+dp[start][i])
				} else if sumLeft > sumRight {
					dp[start][end] = max(dp[start][end], sumRight+dp[i+1][end])
				} else {
					dp[start][end] = max(dp[start][end], sumLeft+max(dp[start][i], dp[i+1][end]))
				}
			}
		}
	}
	return dp[0][n-1]
}
```

时间复杂度 O(n^3)，空间复杂度 O(n^2)

实际测试发现动态规划耗费的实际是上边记忆化搜索的两倍多，应该是动态规划计算了不少没用的子状态，而记忆化搜索则计算的少一些。

也可以最外层枚举长度，代码几乎和上边一致，只是三层循环这样写：

```go
for length := 2; length <= n; length++ {
		for start := 0; start+length <= len(stoneValue); start++ {
			end := start + length - 1
			for i := start; i < end; i++ {
```

实际测试比原来的循环约快 1/3， 这应该是跟测试用例有关了。

#### 时间复杂度优化到 O(n^2) 

对于上面的状态转移，看下这里：

```go
if sumLeft < sumRight { // sum(start, i) < sum(i+1, end)
	dp[start][end] = max(dp[start][end], sumLeft+dp[start][i]) // ... 1
}
```

在这种情况下，可以一直增加 end， start 和 i 不动，因为石子值总是正的，sum(start, i) < sum(i+1, end + k) 将一直满足，其中 k ∈ [1, n-end-1]，则

```go
dp[start][end+k] = max(dp[start][end], sumLeft+dp[start][i]) // ... 2
```

1 和 2 的 **右侧完全一样**。

可根据这一特点来优化。

因为石头的值都是正数，在区间 `[start-1, end)`中一定存在一个分界点 `i'` 满足：
```
sum(start, i') <= sum(i'+1, end)
sum(start, i'+1) > sum(i'+2, end)
```
> 注意 `i'` 可以小于 `start`，此时的 `sum(start, i')` 为 `0`。

那么  当 `i <= i'` 时，总有 `dp[start][end] = max{dp[start][i'] + sum(start, i')} ... 3` 

当 `i > i' ` 时，总有 `dp[start][end] = max{dp[i'+2][end]+sum(i'+2, end)} ... 4`

特别地，如果 `sum(start, i') == sum(i'+1, end)`， 则 `dp[start][end] = max(max{dp[start][i'] + sum(start, i')}, max{dp[i'+1][end] + sum(i'+1, end)})`

维护两个辅助数组 `maxLeft` 和 `maxRight` 来记录 3 和 4 两个等式右边的值。

当 `start` 固定时，这个分界点 `i'` 不会随着 `end` 的增大而递减，所以可以用 `O(n^2)` 的均摊时间复杂度预处理得到所有区间的临界点。之后再枚举区间。

怎么迅速计算 `maxLeft[start][end], maxRight[start][end]` 的值呢？

观察发现，对于  `maxLeft` 和 `maxRight` ，可以递推计算：

```
maxLeft[start][end] = max(maxLeft[start][end-1], sum(start, end)+dp[start][end])
maxRight[start][end] = max(maxRight[start+1][end], sum(start, end)+dp[start][end])
```

在递推更新 `dp` 数组的过程中，可以同时更新这两个数组。之前预处理计算  `i'` 的过程也可以一并放在一起。

每次只需要根据 `i'` 和 两个辅助数组来更新 dp，而不用枚举所有的分割点，详见如下实现：

```go
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	dp := makeMatrix(n)
	maxLeft := makeMatrix(n)
	maxRight := makeMatrix(n)
	prefixSum := make([]int, n+1)
	for i, v := range stoneValue {
		prefixSum[i+1] = prefixSum[i] + v
	}

	for start := n - 1; start >= 0; start-- {
		maxLeft[start][start] = stoneValue[start]
		maxRight[start][start] = stoneValue[start]
		i := start // 临界点。end 增加时，i 不用从 start 重新来过，只是会向右移动（也可能不移动），所以循环里边确定临界点的均摊复杂度是 O(1)。
		for end:= start+1; end < n; end++ {
			sum := prefixSum[end+1] - prefixSum[start]
			// 找到临界点
			for i < end && (prefixSum[i+1]-prefixSum[start])*2 <= sum { // sum of left part * 2 <= sum
				i++
			}
			i-- // 多走了一步退回来

			if start <= i {
				dp[start][end] = max(dp[start][end], maxLeft[start][i])
			}
			if i+1 < end {
				dp[start][end] = max(dp[start][end], maxRight[i+2][end])
			}
			sumLeft := prefixSum[i+1] - prefixSum[start]
			if sumLeft*2 == sum {
				dp[start][end] = max(dp[start][end], max(maxLeft[start][i], maxRight[i+1][end]))
			}
			maxLeft[start][end] = max(maxLeft[start][end-1], sum+dp[start][end])
			maxRight[start][end] = max(maxRight[start+1][end], sum+dp[start][end])
		}
	}
	return dp[0][n-1]
}
```

```go
func makeMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
```

## [1686. 石子游戏 VI](https://leetcode-cn.com/problems/stone-game-vi/)

难度中等

Alice 和 Bob 轮流玩一个游戏，Alice 先手。

一堆石子里总共有 `n` 个石子，轮到某个玩家时，他可以 **移出** 一个石子并得到这个石子的价值。Alice 和 Bob 对石子价值有 **不一样的的评判标准** 。双方都知道对方的评判标准。

给你两个长度为 `n` 的整数数组 `aliceValues`和 `bobValues` 。`aliceValues[i]` 和 `bobValues[i]` 分别表示 Alice 和 Bob 认为第 `i` 个石子的价值。

所有石子都被取完后，得分较高的人为胜者。如果两个玩家得分相同，那么为平局。两位玩家都会采用 **最优策略** 进行游戏。

请你推断游戏的结果，用如下的方式表示：

- 如果 Alice 赢，返回 `1` 。
- 如果 Bob 赢，返回 `-1` 。
- 如果游戏平局，返回 `0` 。

 

**示例 1：**

```
输入：aliceValues = [1,3], bobValues = [2,1]
输出：1
解释：
如果 Alice 拿石子 1 （下标从 0开始），那么 Alice 可以得到 3 分。
Bob 只能选择石子 0 ，得到 2 分。
Alice 获胜。
```

**示例 2：**

```
输入：aliceValues = [1,2], bobValues = [3,1]
输出：0
解释：
Alice 拿石子 0 ， Bob 拿石子 1 ，他们得分都为 1 分。
打平。
```

**示例 3：**

```
输入：aliceValues = [2,4,3], bobValues = [1,6,7]
输出：-1
解释：
不管 Alice 怎么操作，Bob 都可以得到比 Alice 更高的得分。
比方说，Alice 拿石子 1 ，Bob 拿石子 2 ， Alice 拿石子 0 ，Alice 会得到 6 分而 Bob 得分为 7 分。
Bob 会获胜。
```

 

**提示：**

- `n == aliceValues.length == bobValues.length`
- `1 <= n <= 105`
- `1 <= aliceValues[i], bobValues[i] <= 100`

函数签名：

```go
func stoneGameVI(aliceValues []int, bobValues []int)
```

## 分析

直接尝试模拟会非常困难，需要先想想有没有贪心策略。

对于一颗石头 `i`，拿这颗石头自己会加分，同时对方会失分，不妨把两人认为的价值和作为权值，即 `weight[i] = AliceValues[i]+BobValues[i]`，在先拿的时候尽量取权值较大的石头即可 。

或者这样看：

假设 Alice 拿来 i1, i2, ..., ik 位置的石头，那么其总得分就是 `AliceValues[i1] + AliceValues[i2] + ... + AliceValues[ik]`，而 Bob 的总得分就是 `sum(BobValues) - BobValues[i1] - BobValues[i2] - ... -  BobValues[ik]`，总得分差值就是 ` AliceValues[i1]+BobValues[i1] + AliceValues[i2]+BobValues[i2]+ ... + AliceValues[ik]+BobValues[id]` = `weight[i1] + weight[i2]+...+weight[ik]`；可见，在先拿的时候尽量取权值较大的石头可以保证最终得分最高。

```go
func stoneGameVI(aliceValues []int, bobValues []int) int {
    n := len(aliceValues)
    if n != len(bobValues) {
        return 0
    }
    indices := make([]int, n)
    for i := range indices {
        indices[i] = i
    }
    sort.Slice(indices, func(i, j int) bool {
        ii, jj := indices[i], indices[j]
        return aliceValues[ii] + bobValues[ii] > aliceValues[jj] + bobValues[jj]
    })
    alice, bob := 0, 0
    for i, index := range indices {
        if i % 2 == 0 {
            alice += aliceValues[index]
        } else {
            bob += bobValues[index]
        }
    }
    if alice > bob {
        return 1
    }
    if alice < bob {
        return -1
    }
    return 0
}
```

## [1690. 石子游戏 VII](https://leetcode-cn.com/problems/stone-game-vii/)

难度中等

石子游戏中，爱丽丝和鲍勃轮流进行自己的回合，**爱丽丝先开始** 。

有 `n` 块石子排成一排。每个玩家的回合中，可以从行中 **移除** 最左边的石头或最右边的石头，并获得与该行中剩余石头值之 **和** 相等的得分。当没有石头可移除时，得分较高者获胜。

鲍勃发现他总是输掉游戏（可怜的鲍勃，他总是输），所以他决定尽力 **减小得分的差值** 。爱丽丝的目标是最大限度地 **扩大得分的差值** 。

给你一个整数数组 `stones` ，其中 `stones[i]` 表示 **从左边开始** 的第 `i` 个石头的值，如果爱丽丝和鲍勃都 **发挥出最佳水平** ，请返回他们 **得分的差值** 。

 

**示例 1：**

```
输入：stones = [5,3,1,4,2]
输出：6
解释：
- 爱丽丝移除 2 ，得分 5 + 3 + 1 + 4 = 13 。游戏情况：爱丽丝 = 13 ，鲍勃 = 0 ，石子 = [5,3,1,4] 。
- 鲍勃移除 5 ，得分 3 + 1 + 4 = 8 。游戏情况：爱丽丝 = 13 ，鲍勃 = 8 ，石子 = [3,1,4] 。
- 爱丽丝移除 3 ，得分 1 + 4 = 5 。游戏情况：爱丽丝 = 18 ，鲍勃 = 8 ，石子 = [1,4] 。
- 鲍勃移除 1 ，得分 4 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [4] 。
- 爱丽丝移除 4 ，得分 0 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [] 。
得分的差值 18 - 12 = 6 。
```

**示例 2：**

```
输入：stones = [7,90,5,1,100,10,10,2]
输出：122
```

 

**提示：**

- `n == stones.length`
- `2 <= n <= 1000`
- `1 <= stones[i] <= 1000`

函数签名：

```go
func stoneGameVII(stones []int) int
```



## 分析

类似如上石子游戏 II、III、V，区间 dp：

```go
func stoneGameVII(stones []int) int {
    n := len(stones)
    if n < 2 {
        return 0
    }
    prefixSum := make([]int, n+1)
    for i, v := range stones {
        prefixSum[i+1] = prefixSum[i] + v
    }
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for start := n-2; start >= 0; start-- {
        for end := start+1; end < n; end++ {
            dp[start][end] = max(prefixSum[end+1]-prefixSum[start+1] - dp[start+1][end], prefixSum[end]-prefixSum[start] - dp[start][end-1])
        }
    }
    return dp[0][n-1]
}
```
