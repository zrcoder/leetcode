---
title: "丑数"
date: 2022-11-22T10:04:42+08:00
math: true
---

## [263. 丑数](https://leetcode.cn/problems/ugly-number/)

难度简单

**丑数** 就是只包含质因数 `2`、`3` 和 `5` 的正整数。

给你一个整数 `n` ，请你判断 `n` 是否为 **丑数** 。如果是，返回 `true` ；否则，返回 `false` 。

**示例 1：**

**输入：** n = 6
**输出：** true
**解释：** 6 = 2 × 3

**示例 2：**

**输入：** n = 1
**输出：** true
**解释：** 1 没有质因数，因此它的全部质因数是 {2, 3, 5} 的空集。习惯上将其视作第一个丑数。

**示例 3：**

**输入：** n = 14
**输出：** false
**解释：** 14 不是丑数，因为它包含了另外一个质因数 `7` 。

**提示：**

- `-231 <= n <= 231 - 1`

函数签名：

```go
func isUgly(n int) bool
```

## 分析

**丑数**入门，让我们开启丑数之旅。

先处理边界情况：`n < 1`, 因丑数是**正**整数，所以这种情况直接返回 false；`n==1`,示例2已经解释过。

对于`n>1`的情况，可以不断判断 n 是否被 2，3，5 整除，如果是就除以对应的因子，如果最终 n 变成 1，那么原来的数字就是丑数，否则不是。

```go
func isUgly(n int) bool {
    if n <= 0 {
        return false
    }

    for  n % 2 == 0 {
        n /= 2
    }
    for n % 3 == 0 {
        n /= 3
    }
    for n % 5 == 0 {
        n /= 5
    }
    
    return n == 1
}
```

时间复杂度：`O(logn)`，空间复杂度：`O(1)`。

## [264. 丑数 II](https://leetcode.cn/problems/ugly-number-ii/)

难度中等

给你一个整数 `n` ，请你找出并返回第 `n` 个 **丑数** 。

**丑数** 就是只包含质因数 `2`、`3` 和/或 `5` 的正整数。

**示例 1：**

**输入：** n = 10
**输出：** 12
**解释：**[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。

**示例 2：**

**输入：** n = 1
**输出：** 1
**解释：** 1 通常被视为丑数。

**提示：**

- `1 <= n <= 1690`

函数签名：

```go
func nthUglyNumber(n int) int
```

## 分析

### 小顶堆

从1开始构造丑数，每次从已经构造的丑数里找出最小的那个，分别与2、3、5相乘得到新的丑数，用小顶堆维护这些丑数比较合适。

```go
func nthUglyNumber(n int) int {
    h := &Heap{}
    h.push(1)
    res := 0
    for i := 1; i <= n; i++ {
        res = h.pop()
        // 去除重复;如丑数6，可以是2*3， 也可以是3*2
        for h.Len() > 0 && h.peek() == res {
            h.pop()
        }
        h.push(res*2)
        h.push(res*3)
        h.push(res*5)
    }
    return res
}
```

小顶堆相关辅助代码：

```go
type Heap struct {
    s []int
}

func(h *Heap) Len() int {return len(h.s)}
func(h *Heap) Less(i, j int) bool {return h.s[i] < h.s[j]}
func(h *Heap) Swap(i, j int) {h.s[i], h.s[j] = h.s[j], h.s[i]}
func(h *Heap) Push(x interface{}) {h.s = append(h.s, x.(int))}
func(h *Heap) Pop() interface{} {
    n := len(h.s)
    res := h.s[n-1]
    h.s = h.s[:n-1]
    return res
}
func(h *Heap) peek() int {return h.s[0]}
func(h *Heap) push(x int) {heap.Push(h, x)}
func(h *Heap) pop() int {return heap.Pop(h).(int)}
```

时间复杂度：`O(nlogn)`，空间复杂度`O(n)`。

### 动态规划

假设我们用一个数组dp维护丑数。一开始有一个元素 1，不断乘以2、3、5 构造出新的丑数。

可以用三个指针p2, p3, p5，分别记录上一次乘以2、3、5的元素索引，每次在 dp[p2]*2、dp[p3]*3、dp[p5]*5 中选最小值加入数组，同时对应移动三个指针的值。

```go
func nthUglyNumber(n int) int {
    dp := make([]int, n)
    dp[0] = 1
    var p2, p3, p5 int
    for i := 1; i < n; i++ {
        a, b, c := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(a, b, c)
        // 注意 a、b、c 可能有相等的情况
        if dp[i] == a {
            p2++
        }
        if dp[i] == b {
            p3++
        }
        if dp[i] == c {
            p5++
        }
    }
    return dp[n-1]
}

func min(nums ...int) int {
    res := nums[0]
    if nums[1] < res {
        res = nums[1]
    }
    if nums[2] < res {
        res = nums[2]
    }
    return res
}
```

时空复杂度都是`O(n)`。

## [1201. 丑数 III](https://leetcode.cn/problems/ugly-number-iii/)

难度中等

给你四个整数：`n` 、`a` 、`b` 、`c` ，请你设计一个算法来找出第 `n` 个丑数。

丑数是可以被 `a` **或** `b` **或** `c` 整除的 **正整数** 。

**示例 1：**

**输入：** n = 3, a = 2, b = 3, c = 5
**输出：** 4
**解释：** 丑数序列为 2, 3, 4, 5, 6, 8, 9, 10... 其中第 3 个是 4。

**示例 2：**

**输入：** n = 4, a = 2, b = 3, c = 4
**输出：** 6
**解释：** 丑数序列为 2, 3, 4, 6, 8, 9, 10, 12... 其中第 4 个是 6。

**示例 3：**

**输入：** n = 5, a = 2, b = 11, c = 13
**输出：** 10
**解释：** 丑数序列为 2, 4, 6, 8, 10, 11, 12, 13... 其中第 5 个是 10。

**示例 4：**

**输入：** n = 1000000000, a = 2, b = 217983653, c = 336916467
**输出：** 1999999984

**提示：**

- `1 <= n, a, b, c <= 10^9`
- `1 <= a * b * c <= 10^18`
- 本题结果在 `[1, 2 * 10^9]` 的范围内

## 分析

注意这里对丑数的定义是和上边两题不同的。上边两题说的是`仅包含xxxx因子`，而这个问题是可以被 `a` **或** `b` **或** `c` 整除，比如示例2，因子2， 3， 4， 但是丑数里有10，10 的因子5虽然不在给出的三个因子里但另一个因子2在。

需要另找思路，可以先看看下个问题 **878. 第 N 个神奇数字**，非常有趣，用到了容斥原理和二分搜索。

## [878. 第 N 个神奇数字](https://leetcode.cn/problems/nth-magical-number/)

难度困难

一个正整数如果能被 `a` 或 `b` 整除，那么它是神奇的。

给定三个整数 `n` , `a` , `b` ，返回第 `n` 个神奇的数字。因为答案可能很大，所以返回答案 **对** $10^9+7$ **取模** 后的值。

**示例 1：**

**输入：** n = 1, a = 2, b = 3
**输出：** 2

**示例 2：**

**输入：** n = 4, a = 2, b = 3
**输出：** 6

**提示：**

- `1 <= n <= 109`
- `2 <= a, b <= 4 * 104`

## 分析

假设 $m = min(a, b)$，首先可以确定，答案在$[m, n*m]$ 区间内。

遍历区间内每个数字 $x$，**计算不大于 x 的丑数有多少个**，如果恰好等于 $n$，$x$ 就是所求。

对于给定的数字 $x$，设 $f(x)$ 代表不大于$x$的丑数的个数，$A$、$B$分别表示不大于$x$的数字中能被 $a$、$b$整除的数字集合，则 $f(x)$ 应为 $A\cup B$ 的个数，则根据容斥原理，$f(x)=count(A)+count(B)-count(A\cap B)=\lfloor\frac{x}{a}\rfloor +\lfloor\frac{x}{b}\rfloor − \lfloor\frac{x}{c}\rfloor$，其中$c$为$a$和$b$的最小公倍数。

这是一个关于 $x$ 的单调非递减函数，可以把遍历改成二分搜索。

```go
const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
    lo := min(a, b)
    hi := lo*n+1
    c := a / gcd(a, b) * b
    return (lo + sort.Search(hi-lo, func(x int) bool {
        x += lo
        return x/a+x/b-x/c >= n
    })) % mod
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
```

时间复杂度：$O(log(n \times min(a, b))$，空间复杂度：$O(1)$。

对于 **1201. 丑数 III**，因子变成了三个，$f(x)=count(A)+count(B)+count(C)-count(A\cap B)-count(A\cap C)-count(B\cap C)+count(A\cap B\cap C)$。

```go
func nthUglyNumber(n int, a int, b int, c int) int {
    ab := a/gcd(a, b)*b
    ac := a/gcd(a, c)*c
    bc := b/gcd(b, c)*c
    abc := ab/gcd(ab, c)*c

    lo := min(a, b, c)
    hi := lo*n+1
    return lo + sort.Search(hi-lo, func(x int) bool {
        x += lo
        return x/a + x/b + x/c - x/ab - x/ac - x/bc + x/abc >= n
    })
}
```
