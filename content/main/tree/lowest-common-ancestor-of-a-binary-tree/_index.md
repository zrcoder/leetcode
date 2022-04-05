---
title: "236. 二叉树的最近公共祖先"
date: 2022-04-03T15:37:56+08:00
---

## [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/description/ "https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (68.72%) | 1659  | -        |

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

[百度百科](https://baike.baidu.com/item/%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88/8918834?fr=aladdin "https://baike.baidu.com/item/%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88/8918834?fr=aladdin")中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（**一个节点也可以是它自己的祖先**）。”

**示例 1：**

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
```

**示例 3：**

```
输入：root = [1,2], p = 1, q = 2
输出：1
```

**提示：**

- 树中节点数目在范围 `[2, 105]` 内。
- `-109 <= Node.val <= 109`
- 所有 `Node.val` `互不相同` 。
- `p != q`
- `p` 和 `q` 均存在于给定的二叉树中。

函数签名：

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode
```

## 分析

### 递归

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p == root || q == root {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
}
```

时空复杂度均为O (n)

### 寻找父节点

可以遍历一遍树，借助一个哈希表记录每个节点的父节点；之后从节点p开始向上直到root，用一个set记录过程中经过的节点；最后从q开始向上，如果途经p向上的路径上的节点即位所求。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent := map[*TreeNode]*TreeNode{}
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left != nil {
            parent[node.Left] = node
            dfs(node.Left)
        }
        if node.Right != nil {
            parent[node.Right] = node
            dfs(node.Right)
        }
    }
    dfs(root)

    path := map[*TreeNode]bool{}
    for p != nil {
        path[p] = true
        p = parent[p]
    }
    for q != nil {
        if path[q] {
            return q
        }
        q = parent[q]
    }
    return nil
}
```

时空复杂度均为O (n)

## 扩展

如果树是BST呢？除了上边对于一般二叉树的解法，也可以利用BST树节点值的特性来做，如  [235. 二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root
}
```

时空复杂度O(h)，h为树高，最好情况下是树退化成链表，h == n，最优情况下是BST相对平衡，h == logn.
