---
title: "1373. 二叉搜索子树的最大键值和"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1373. 二叉搜索子树的最大键值和](https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/)

难度困难

给你一棵以 `root` 为根的 **二叉树** ，请你返回 **任意** 二叉搜索子树的最大键值和。

二叉搜索树的定义如下：

- 任意节点的左子树中的键值都 **小于** 此节点的键值。
- 任意节点的右子树中的键值都 **大于** 此节点的键值。
- 任意节点的左子树和右子树都是二叉搜索树。

 

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/03/07/sample_1_1709.png)

```
输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
输出：20
解释：键值为 3 的子树是和最大的二叉搜索树。
```

**示例 2：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/03/07/sample_2_1709.png)

```
输入：root = [4,3,null,1,2]
输出：2
解释：键值为 2 的单节点子树是和最大的二叉搜索树。
```

**示例 3：**

```
输入：root = [-4,-2,-5]
输出：0
解释：所有节点键值都为负数，和最大的二叉搜索树为空。
```

**示例 4：**

```
输入：root = [2,1,3]
输出：6
```

**示例 5：**

```
输入：root = [5,4,8,3,null,6,3]
输出：7
```

 

**提示：**

- 每棵树最多有 `40000` 个节点。
- 每个节点的键值在 `[-4 * 10^4 , 4 * 10^4]` 之间。

函数签名：
```go
func maxSumBST(root *TreeNode) int
```
## 分析

### 在二叉树上做动态规划

为了得到结果，需要以某种顺序遍历每个节点。在遍历过程中，对于当前节点，需要以下信息确定是不是 BST：
`左右子树是否为 BST，左子树的最大值、右子树的最小值。`
最终判断当前节点为根的树是 BST 后还要计算所有节点和，这可以通过事先计算左右子树的和，之后加上自身的键值得到。
综上，需要后序遍历，每次返回以当前节点为根的子树最小值、最大值、和及是否为 BST 四个信息。

```go
func maxSumBST(root *TreeNode) int {
	res := 0 // 注意示例3，空子树算bst，其和为0
	var dfs func(root *TreeNode) (int, int, int, bool)
	dfs = func(root *TreeNode) (int, int, int, bool) {
		if root == nil {
			// 单独的叶子节点是 BST，由这一约束确定空节点的各项值。
			return math.MaxInt64, math.MinInt64, 0, true
		}
		minLeft, maxLeft, sumLeft, isLeftBst := dfs(root.Left)
		minRight, maxRight, sumRight, isRightBst := dfs(root.Right)
		var minVal, maxVal, sum int
		var isBst bool
		if isLeftBst && isRightBst && maxLeft < root.Val && minRight > root.Val {
			isBst = true
			minVal = min(minLeft, root.Val)
			maxVal = max(maxRight, root.Val)
			sum = sumLeft + sumRight + root.Val
			res = max(sum, res)
		}
		return minVal, maxVal, sum, isBst
	}
	dfs(root)
	return res
}
```