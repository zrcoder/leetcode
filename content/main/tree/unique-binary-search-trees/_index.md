---
title: "不同的二叉搜索树"
date: 2022-04-02T16:37:48+08:00
---
## [96. 不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/description/ "https://leetcode-cn.com/problems/unique-binary-search-trees/description/")

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (70.04%) | 1665 | -   |

给你一个整数 `n` ，求恰由 `n` 个节点组成且节点值从 `1` 到 `n` 互不相同的 **二叉搜索树** 有多少种？返回满足题意的二叉搜索树的种数。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)

```
输入：n = 3
输出：5
```

**示例 2：**

```
输入：n = 1
输出：1
```

**提示：**

- `1 <= n <= 19`

函数签名：

```go
func numTrees(n int) int
```

## 分析

可以递归地解决这个问题。

一个明显的结论是，枚举每个数字，让其作为根节点来构造一棵BST，则构成的树必是唯一的。

下面考虑当枚举到数字 i 时，应该怎么计算以i为根的可能的BST共有多少棵。很明显，左子树用[1, i-1]区间里的数字构造，而右子树用[i+1, n]区间里的数字构造，左子树共 i-1个节点，右子树共 n-i个节点。容易发现构造出的树的个数只和节点个数有关，而和节点的内容无关。

所以构造以i为根的BST，共有的可能方案数是 `numTrees(i-1)*numTrees(n-i)`。

```go
func numTrees(n int) int {
    if n < 2 { // 注意 n==0 的情况，对应 i 为1 或 n 即左子树或右子树可为nil的情况
        return 1
    }
    res := 0
    for i := 1; i <= n; i++ {
        res += numTrees(i-1)*numTrees(n-i)
    }
    return res
}
```

有很多重复计算，可以加上备忘录来优化。

```go
func numTrees(n int) int {
    memo := make([]int, n+1)
    var help func(int) int
    help = func(n int) int {
        if n < 2 { // 注意 n==0 的情况，对应 i 为1 或 n 即左子树或右子树可为nil的情况
            return 1
        }
        if memo[n] > 0 {
            return memo[n]
        }
        res := 0
        for i := 1; i <= n; i++ {
            res += help(i-1) * help(n-i)
        }
        memo[n] = res
        return res
    }
    return help(n)
}
```

可以进一步改写为自底向上的动态规划写法。

```go
func numTrees(n int) int {
    f := make([]int, n+1)
    f[0], f[1] = 1, 1
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            f[i] += f[j-1] * f[i-j]
        }
    }
    return f[n]
}
```

时间复杂度 O(n^2)， 空间复杂度O(n)。

## [95. 不同的二叉搜索树 II](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/ "https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/")

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (71.48%) | 1178 | -   |

给你一个整数 `n` ，请你生成并返回所有由 `n` 个节点组成且节点值从 `1` 到 `n` 互不相同的不同 **二叉搜索树** 。可以按 **任意顺序** 返回答案。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)

```
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
```

**示例 2：**

```
输入：n = 1
输出：[[1]]
```

**提示：**

- `1 <= n <= 8`

函数签名：

```go
func generateTrees(n int) []*TreeNode
```

## 分析

这次需要真正构造出所有可能的BST。总体思路和上边是类似的，即枚举所有数字作为根节点来构造。但需要注意这次构造子树的时候必须考虑子树各个节点的值，至少需要知道子树中节点的最大值和最小值。

```go
func generateTrees(n int) []*TreeNode {
    var help func(lo, hi int) []*TreeNode
    help = func(lo, hi int) []*TreeNode {
        if lo > hi {
            return []*TreeNode{nil}
        }
        var res []*TreeNode
        for i := lo; i <= hi; i++ {
            left, right := help(lo, i-1), help(i+1, hi)
            for _, lv := range left {
                for _, rv := range right {
                    res = append(res, &TreeNode{
                        Val:   i,
                        Left:  lv,
                        Right: rv,
                    })
                }
            }
        }
        return res
    }
    return help(1, n)
}
```

假设总共生成的BST有 X 棵，则时空复杂度都是O(n*X)，X到底是多少？不太好分析，实际就是上一问题求的的结果，前人已经分析过了，称卡特兰数，这里不作讨论。

