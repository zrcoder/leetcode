---
title: 863. 二叉树中所有距离为 K 的结点
date: 2023-07-11T09:52:01+08:00
---

## [863. 二叉树中所有距离为 K 的结点](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree) (Medium)

给定一个二叉树（具有根结点 `root`）， 一个目标结点 `target` ，和一个整数值 `k` 。

返回到目标结点 `target` 距离为 `k` 的所有结点的值的列表。 答案可以以 **任何顺序** 返回。

**示例 1：**

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png)

```
输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
输出：[7,4,1]
解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1

```

**示例 2:**

```
输入: root = [1], target = 1, k = 3
输出: []

```

**提示:**

- 节点数在 `[1, 500]` 范围内
- `0 <= Node.val <= 500`
- `Node.val` 中所有值 **不同**
- 目标结点 `target` 是树上的结点。
- `0 <= k <= 1000`

## 分析

先 DFS 获知每个节点的父节点，再 BFS/DFS 得到距离 target 为 k 的节点集合

```go

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	parents := getParents(root)

	var res []int
	var dfs func(pre, cur *TreeNode, dist int)
	action := func(pre, cur, next *TreeNode, dist int) {
		if next != nil && next != pre {
			dfs(cur, next, dist+1)
		}
	}
	dfs = func(pre, cur *TreeNode, dist int) {
		if dist == k {
			res = append(res, cur.Val)
			return
		}
		action(pre, cur, cur.Left, dist)
		action(pre, cur, cur.Right, dist)
		action(pre, cur, parents[cur], dist)
	}
	dfs(nil, target, 0)
	return res
}

func getParents(root *TreeNode) map[*TreeNode]*TreeNode {
	parent := map[*TreeNode]*TreeNode{}
	var dfs func(cur, p *TreeNode)
	dfs = func(cur, p *TreeNode) {
		if cur == nil {
			return
		}
		parent[cur] = p
		dfs(cur.Left, cur)
		dfs(cur.Right, cur)
	}
	dfs(root, nil)
	return parent
}

```
