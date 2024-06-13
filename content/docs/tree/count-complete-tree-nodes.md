---
title: "222. 完全二叉树的节点个数"
date: 2022-04-03T16:26:56+08:00
math: true
---

## [222. 完全二叉树的节点个数](https://leetcode-cn.com/problems/count-complete-tree-nodes/description/ "https://leetcode-cn.com/problems/count-complete-tree-nodes/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (79.15%) | 653   | -        |

给你一棵 **完全二叉树** 的根节点 `root` ，求出该树的节点个数。

[完全二叉树](https://baike.baidu.com/item/%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91/7773232?fr=aladdin "https://baike.baidu.com/item/%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91/7773232?fr=aladdin") 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 `h` 层，则该层包含 `1~ 2h` 个节点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/14/complete.jpg)

```
输入：root = [1,2,3,4,5,6]
输出：6
```

**示例 2：**

```
输入：root = []
输出：0
```

**示例 3：**

```
输入：root = [1]
输出：1
```

**提示：**

- 树中节点的数目范围是`[0, 5 * 104]`
- `0 <= Node.val <= 5 * 104`
- 题目数据保证输入的树是 **完全二叉树**

**进阶：** 遍历树来统计节点是一种时间复杂度为 `O(n)` 的简单解决方案。你可以设计一个更快的算法吗？

## 分析

遍历树的解决方案会非常简单，如：

```go
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```

以上代码适用于任何二叉树，完全没利用完全二叉树的性质，需要遍历所有节点，时间复杂度不理想，为 $O(n)$。

如果是一棵满二叉树，可以根据其高度直接计算出节点数。假设高度是$h$，节点总数是：$\sum_{i=0}^{h-1}2^{i}$，根据等比数列前n项和公式，即为：$2^h-1$ . 而满二叉树的高度可以在 $O(log_2n)$ 复杂度内求出。

完全二叉树可以结合普通二叉树和完全二叉树的解法：

```go
func countNodes(root *TreeNode) int {
    // 分别向左、向右记录计算高度 lh 和 rh
    // 对于完全二叉树 rh == lh 或 rh == lh-1
    lh := 0
    for p := root; p != nil; p = p.Left {
        lh++
    }
    rh := 0
    for p := root; p != nil; p = p.Right {
        rh++
    }

    if lh == rh { // 这是一棵满二叉树
        return (1 << lh) - 1
    }
    // root 不是满二叉树，按照普通树的方式计算
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```

关键在于时间复杂度的分析。直观地看好像和普通树的解法复杂度差不多，实则不然，关键在于最后一行的递归调用，`root.Left` 和 `root.Right` 必有一棵是满二叉树，满二叉树会在求完左右高度后直接返回，不会一直递归下去。

递归深度是树的高度$O(log_2n)$, 每次递归花费的时间是向左向右计算高度的两个循环，耗时同样为$O(log_2n)$, 所有总的复杂度是$O((log_2n)^2)$。
