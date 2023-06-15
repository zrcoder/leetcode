---
title: 1161. 最大层内元素和
---

## [1161. 最大层内元素和](https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree) (Medium)

给你一个二叉树的根节点 `root`。设根节点位于二叉树的第 `1` 层，而根节点的子节点位于第 `2` 层，依此类推。

请返回层内元素之和 **最大** 的那几层（可能只有一层）的层号，并返回其中 **最小** 的那个。

**示例 1：**

**![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/08/17/capture.jpeg)**

```
输入：root = [1,7,0,7,-8,null,null]
输出：2
解释：
第 1 层各元素之和为 1，
第 2 层各元素之和为 7 + 0 = 7，
第 3 层各元素之和为 7 + -8 = -1，
所以我们返回第 2 层的层号，它的层内元素之和最大。

```

**示例 2：**

```
输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
输出：2

```

**提示：**

- 树中的节点数在 `[1, 10⁴]` 范围内
- `-10⁵ <= Node.val <= 10⁵`

## 分析

### BFS
```go
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
	maxSum := math.MinInt64
	for lvl := 1; len(q) > 0; lvl++ {
		next := make([]*TreeNode, 0, len(q))
		sum := 0
		for _, v := range q {
			sum += v.Val
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		if sum > maxSum {
			res = lvl
			maxSum = sum
		}
		q = next
	}
	return res
}
```
### DFS

```go
func maxLevelSum(root *TreeNode) int {
	sum := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, lvl int) {
		if root == nil {
			return
		}
		if lvl == len(sum) {
			sum = append(sum, root.Val)
		} else {
			sum[lvl] += root.Val
		}
		dfs(root.Left, lvl+1)
		dfs(root.Right, lvl+1)
	}
	dfs(root, 0)
	res := 0
	maxSum := math.MinInt
	for i, v := range sum {
		if v > maxSum {
			maxSum = v
			res = i + 1
		}
	}
	return res
}

```
