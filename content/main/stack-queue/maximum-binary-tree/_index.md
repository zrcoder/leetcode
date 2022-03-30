---
title: "最大二叉树"
date: 2022-03-30T20:01:13+08:00
---
## [最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/description/ "https://leetcode-cn.com/problems/maximum-binary-tree/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (80.99%) | 394   | -        |

给定一个不重复的整数数组 `nums` 。 **最大二叉树** 可以用下面的算法从 `nums` 递归地构建:

1. 创建一个根节点，其值为 `nums` 中的最大值。
2. 递归地在最大值 **左边** 的 **子数组前缀上** 构建左子树。
3. 递归地在最大值 **右边** 的 **子数组后缀上** 构建右子树。

返回 *`nums` 构建的* ***最大二叉树*** 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg)

```
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
        - 空数组，无子节点。
        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
            - 空数组，无子节点。
            - 只有一个元素，所以子节点是一个值为 1 的节点。
    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
        - 只有一个元素，所以子节点是一个值为 0 的节点。
        - 空数组，无子节点。
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg)

```
输入：nums = [3,2,1]
输出：[3,null,2,null,1]
```

**提示：**

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`
- `nums` 中的所有整数 **互不相同**

## 分析

### 递归

递归是最容易实现的解法。

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    m, i := math.MinInt64, -1
    for j, v := range nums {
        if v > m {
            m = v
            i = j
        }
    }
    return &TreeNode{
        Val:   m,
        Left:  constructMaximumBinaryTree(nums[:i]),
        Right: constructMaximumBinaryTree(nums[i+1:]),
    }
}
```

时间复杂度在：平均`O(nlogn)`，最坏`O(n^2)`；

空间复杂度：平均`O(logn)`, 最坏`O(n)`。

### 借助两个单调栈+一个数组

每个数字的父节点是哪个？
找到左边第一个比它大的数a, 右边第一个比它大的数b, 答案是a、b中较小的那个。

> 如果a, b 都不存在，说明当前数字是树根；
>
> 如果a，b有一个不存在，这个情况也好解决。

1. 用一个数组存储所有节点；

2. 事先用单调栈的方式获得每个节点左边/右边第一个比其大的节点；

3. 遍历1中的数组，根据2得到的两个记录构建树。

直接看代码：

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
    nodes := make([]*TreeNode, len(nums))
    for i, v := range nums {
        nodes[i] = &TreeNode{Val: v}
    }
    leftBigger := getLeftBigger(nodes)
    rightBigger := getRightBigger(nodes)
    var root *TreeNode
    for i, node := range nodes {
        left := leftBigger[i]
        right := rightBigger[i]
        if left == nil && right == nil {
            root = node
        } else if left == nil {
            right.Left = node
        } else if right == nil {
            left.Right = node
        } else if left.Val > right.Val {
            right.Left = node
        } else {
            left.Right = node
        }
    }
    return root
}

func getLeftBigger(nodes []*TreeNode) []*TreeNode {
    return getBigger(nodes, true)
}
func getRightBigger(nodes []*TreeNode) []*TreeNode {
    return getBigger(nodes, false)
}

func getBigger(nodes []*TreeNode, isLow2High bool) []*TreeNode {
    res := make([]*TreeNode, len(nodes))
    stack := []*TreeNode{}
    from, to, step := 0, len(nodes)-1, 1
    if !isLow2High {
        from, to, step = len(nodes)-1, 0, -1
    }
    for i := from; isLow2High && i <= to || !isLow2High && i >= to; i += step {
        for len(stack) > 0 && stack[len(stack)-1].Val <= nodes[i].Val {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 {
            res[i] = stack[len(stack)-1]
        }
        stack = append(stack, nodes[i])
    }
    return res
}
```

时间复杂度降到了`O(n)`。

可以想一想这样为什么能正确构建，能不能证明正确性？这部分略。

### 仅借助一个单调栈

上边的思路和代码稍嫌复杂，能不能简化呢？

借助一个辅助递减栈，栈中存储节点。

1. 遍历所有数字，对于当前数字和当前栈，分情况操作：

   - 如果栈空或栈顶节点值大于当前数字，当前数字对应的节点直接入栈；

   - 如果栈不空且栈顶节点值小于当前数字，记录栈顶的节点为cur，cur出栈，且要确定cur的父节点，这个父节点就是新节点和新栈顶中较小的那个，cur作为左孩子还是右孩子也是显而易见的。

2. 在这之后，栈不空，需要将栈顶一一出栈且确定父节点。

```go
func constructMaximumBinaryTree(nums []int) *TreeNode {
    stack := []*TreeNode{}
    var cur *TreeNode
    for _, v := range nums {
        node := &TreeNode{Val: v}
        for len(stack) > 0 && stack[len(stack)-1].Val <= v {
            cur = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 || stack[len(stack)-1].Val > v {
                node.Left = cur
            } else {
                stack[len(stack)-1].Right = cur
            }
        }
        stack = append(stack, node)
    }
    for len(stack) > 0 {
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if len(stack) > 0 {
            stack[len(stack)-1].Right = cur
        }
    }
    return cur
}
```


