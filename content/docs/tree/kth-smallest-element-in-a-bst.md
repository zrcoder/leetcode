---
title: "230. 二叉搜索树中第K小的元素"
date: 2022-12-20T17:33:27+08:00
---

##  [230. 二叉搜索树中第K小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description)

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (75.81%) | 691 | -   |

给定一个二叉搜索树的根节点 `root` ，和一个整数 `k` ，请你设计一个算法查找其中第 `k` 个最小元素（从 1 开始计数）。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)

```
输入：root = [3,1,4,null,2], k = 1
输出：1
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)

```
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
```

**提示：**

- 树中的节点数为 `n` 。
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

**进阶：** 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 `k` 小的值，你将如何优化算法？

函数签名：

```go
func kthSmallest(root *TreeNode, k int)
```

## 分析

### 中序遍历

因为BST的中序遍历是有序的，所以只需要做中序遍历即可。

```go
func kthSmallest(root *TreeNode, k int) int {
	res := -1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	return res
}
```

写成迭代版：

```go
func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
        return -1
    }

    stk := []*TreeNode{root}
    for len(stk) > 0 || root != nil {
        for root != nil {
            stk = append(stk, root)
            root = root.Left
        }
        n := len(stk)
        root = stk[n-1]
        stk = stk[:n-1]
        k--
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
    
    return -1
}
```

时间复杂度：`O(h+k)`，空间复杂度：`O(h)`。其中 h 指树的高度。

### 进阶问题

如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 `k` 小的值，你将如何优化算法？

需要记录每个节点作为根的子树的节点数，在增删过程中除了维护 BST 的特性，还需维护每个子树的节点数。

参考如下：

```go
type MyBst struct {
	root *TreeNode
	size map[*TreeNode]int // 维护以每个结点为根结点的子树的结点数
}

// 统计以 node 为根结点的子树的结点数
func (t *MyBst) count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	t.size[node] = 1 + t.count(node.Left) + t.count(node.Right)
	return t.size[node]
}

// 返回二叉搜索树中第 k 小的元素
func (t *MyBst) kthSmallest(k int) int {
	for node := t.root; node != nil && k > 0; {
		leftSize := t.size[node.Left]
		if leftSize == k-1 {
			return node.Val
		}

		if leftSize < k-1 {
			node = node.Right
			k -= leftSize + 1
		} else {
			node = node.Left
		}
	}

	return -1
}

func kthSmallest(root *TreeNode, k int) int {
	tree := &MyBst{root, map[*TreeNode]int{}}
	tree.count(root)
	return tree.kthSmallest(k)
}
```

这样除了预处理，每次查询的时间复杂度是 `O(h)`。

当然，如果树非常不平衡，h 会接近 n。所以这个优化还不够，需要用平衡的 BST 来代替一般的BST。略。
