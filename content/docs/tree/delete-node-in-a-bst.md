---
title: "450. 删除二叉搜索树中的节点"
date: 2022-04-02T15:58:11+08:00
---
## [450. 删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/description/ "https://leetcode-cn.com/problems/delete-node-in-a-bst/description/")

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (50.20%) | 695 | -   |

给定一个二叉搜索树的根节点 **root** 和一个值 **key**，删除二叉搜索树中的 **key** 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

1. 首先找到需要删除的节点；
2. 如果找到了，删除它。

**示例 1:**

![](https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg)

```
输入：root = [5,3,6,2,4,null,7], key = 3
输出：[5,4,6,2,null,null,7]
解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
另一个正确答案是 [5,2,6,null,4,null,7]。
```

**示例 2:**

```
输入: root = [5,3,6,2,4,null,7], key = 0
输出: [5,3,6,2,4,null,7]
解释: 二叉树不包含值为 0 的节点
```

**示例 3:**

```
输入: root = [], key = 0
输出: []
```

**提示:**

- 节点数的范围 `[0, 104]`.
- `-105 <= Node.val <= 105`
- 节点值唯一
- `root` 是合法的二叉搜索树
- `-105 <= key <= 105`

**进阶：** 要求算法时间复杂度为 O(h)，h 为树的高度。

函数签名：

```go
func deleteNode(root *TreeNode, key int) *TreeNode
```

## 分析

BST 中序遍历后是单调递增的序列。这样使得查找元素变成二分查找了。

要删除一个BST里的节点，首先用二分的方式找到它，然后该怎么删除呢？找到左子树中的最大节点（或者找右子树中的最小节点），然后让它来代替要被删除的节点。

怎么找？为什么这么替换后依然是一棵BST？都可以用一句话回答：BST里的任意子树依然是一棵BST。

先写以查找节点为主的代码框架：

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    dummy := &TreeNode{Left: root}
    parent, cur := dummy, root
    isLeft := true
    for cur != nil {
        if cur.Val == key {
            break
        }
        parent = cur
        if cur.Val < key {
            isLeft = false
            cur = cur.Right
        } else {
            isLeft = true
            cur = cur.Left
        }
    }
    if cur != nil {
        if isLeft {
            parent.Left = delete(cur)
        } else {
            parent.Right = delete(cur)
        }
    }
    dummy.Left, cur = nil, dummy.Left
    return cur
}
```

注意到引入了一个dummy节点，这是因为有可能被删除的节点就是root，引入哨兵节点能简化代码。

下边是删除一个节点的函数，返回删除后新的节点：

```go
func delete(node *TreeNode) *TreeNode {
    if node.Left == nil && node.Right == nil {
        return nil
    }
    if node.Left == nil {
        res := node.Right
        node.Right = nil
        return res
    }
    if node.Right == nil {
        res := node.Left
        node.Left = nil
        return res
    }
    // 找到左子树的最大节点来替换当前节点（也可找右子树的最小节点）
    left, right := node.Left, node.Right
    var parent *TreeNode
    p := left
    for p.Right != nil {
        parent = p
        p = p.Right
    }
    if parent != nil { // p != left
        parent.Right = p.Left
        p.Left = left
    }
    p.Right = right
    node.Left = nil
    node.Right = nil
    return p
}
```

时间复杂度为 O(h), h是树的高度。空间复杂度为 O(1)。

## 扩展

怎么给BST插入新值呢？——这个比删除要简单很多，不详细讨论。

可直接尝试解决： [701. 二叉搜索树中的插入操作](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/) 。
