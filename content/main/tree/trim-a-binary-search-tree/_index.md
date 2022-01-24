---
title: "669. 修剪二叉搜索树"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [669. 修剪二叉搜索树](https://leetcode-cn.com/problems/trim-a-binary-search-tree/)

难度中等

给你二叉搜索树的根节点 `root` ，同时给定最小边界`low` 和最大边界 `high`。通过修剪二叉搜索树，使得所有节点的值在`[low, high]`中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg)

```
输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg)

```
输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]
```

**示例 3：**

```
输入：root = [1], low = 1, high = 2
输出：[1]
```

**示例 4：**

```
输入：root = [1,null,2], low = 1, high = 3
输出：[1,null,2]
```

**示例 5：**

```
输入：root = [1,null,2], low = 2, high = 4
输出：[2]
```

 

**提示：**

- 树中节点数在范围 `[1, 104]` 内
- `0 <= Node.val <= 104`
- 树中每个节点的值都是唯一的
- 题目数据保证输入是一棵有效的二叉搜索树
- `0 <= low <= high <= 104`

BST 树定义：

```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
```



函数签名：

```go
func trimBST(root *TreeNode, low int, high int) *TreeNode
```

## 分析

对于当前节点，如果其值小于 low，那么应该在其右子树寻找答案，其本身及左子树应该抛弃；如果其值大于 high，应该在其左子树寻求答案，其本身及右子树应该抛弃；剩下的情况就是其值在 [low, hight] 的闭区间里，这时候只需要修建左右子树就行，最后返回当前节点本身。

```go
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		tmp := root.Right
		root.Right = nil
		return trimBST(tmp, low, high)
	}
	if root.Val > high {
		tmp := root.Left
		root.Left = nil
		return trimBST(tmp, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
```