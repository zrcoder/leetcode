---
title: "树的路径和"
weight: 1
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: false
# bookComments: false
# bookSearchExclude: false
---

## [112. 路径总和](https://leetcode-cn.com/problems/path-sum/)

难度简单

给你二叉树的根节点 `root` 和一个表示目标和的整数 `targetSum` 。判断该树中是否存在 **根节点到叶子节点** 的路径，这条路径上所有节点值相加等于目标和 `targetSum` 。如果存在，返回 `true` ；否则，返回 `false` 。

**叶子节点** 是指没有子节点的节点。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。
```

**示例 3：**

```
输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。
```

 

**提示：**

- 树中节点的数目在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`


### 分析

常规dfs，更一般的回溯思路见下边**路径总和 II**

```go
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
```
## [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)

难度中等

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum` ，找出所有 **从根节点到叶子节点** 路径总和等于给定目标和的路径。

**叶子节点** 是指没有子节点的节点。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg)

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
输入：root = [1,2,3], targetSum = 5
输出：[]
```

**示例 3：**

```
输入：root = [1,2], targetSum = 0
输出：[]
```

 

**提示：**

- 树中节点总数在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

### 分析

用一个切片path记录遍历的路径，到达叶子节点发现path内元素和为sum则将当期path添加到结果里，注意回溯。
```go
func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	var path []int
	prefixSum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		prefixSum += node.Val
		if node.Left == nil && node.Right == nil && prefixSum == sum {
			// 切片底层是同一个数组，添加到结果时要深拷贝一份
			tmp := make([]int, len(path))
			_ = copy(tmp, path)
			result = append(result, tmp)
		}
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		path = path[:len(path)-1]
		prefixSum -= node.Val
	}
	dfs(root)
	return result
}
```
### 扩展
> 类似前缀树Trie的实现，用数组表示树，且树是多叉树时，应该怎么解？
问题具体描述如下：
```
假设有k个节点，每个节点从0到k-1编号，编号即为其id,caps数组表示每个节点的值。
哈希表relations，是个邻接表，键为节点id，值为节点的孩子节点组成的数组。
给定sum，返回每条从根节点（id为0）出发到叶子节点，值相加和为sum的路径组成的集合。
路径处理成字符串，前最终结果按照字符串非递增排序。
```
```go
func getPath(caps []int, relations map[int][]int, sum int) []string {
	var result []string
	var path []int
	prefixSum := 0
	var dfs func(nodeId int)

	dfs = func(nodeId int) {
		path = append(path, caps[nodeId])
		prefixSum += caps[nodeId]
		if len(relations[nodeId]) == 0 && prefixSum == sum {
			result = append(result, parsePath(path))
		}
		for _, c := range relations[nodeId] {
			dfs(c)
		}
		path = path[:len(path)-1]
		prefixSum -= caps[nodeId]
	}
	dfs(0)

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}
func parsePath(path []int) string {
	buf := bytes.NewBuffer(nil)
	for _, v := range path {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(" ")
	}
	result := buf.String()
	return result[:len(result)-1]
}
```

## [437. 路径总和 III](https://leetcode-cn.com/problems/path-sum-iii/)

难度中等

给定一个二叉树的根节点 `root` ，和一个整数 `targetSum` ，求该二叉树里节点值之和等于 `targetSum` 的 **路径** 的数目。

**路径** 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg)

```
输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
```

**示例 2：**

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3
```

 

**提示:**

- 二叉树的节点个数的范围是 `[0,1000]`
- `-109 <= Node.val <= 109` 
- `-1000 <= targetSum <= 1000` 

### 分析
可以稍作转化：求出以每个节点作为根的树的前缀和为sum的个数（不一定到叶子节点），相加就是结果。

先尝试下直观的递归解法
```go
// 时间复杂度较高，有比较多的重复计算
func pathSumCount(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := countPrefix(root, sum)
	result += pathSumCount(root.Left, sum)
	result += pathSumCount(root.Right, sum)
	return result
}

// 返回前缀和为sum的路径个数， 递归版
func countPrefix(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == sum {
		result = 1
	}
	result += countPrefix(root.Left, sum-root.Val)
	result += countPrefix(root.Right, sum-root.Val)
	return result
}
```
时间复杂度为`O(n^2)`，其中 n 是树中节点总数。

countPrefix 按照回溯的思路稍微修改下

```go
func countPrefix1(root *TreeNode, sum int) int {
	prefixSum := 0
	result := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		prefixSum += node.Val
		if prefixSum == sum {
			result++
		}
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		prefixSum -= node.Val
	}
	dfs(root)
	return result
}
```

对于同一路径上的两个节点x，y， 假设 prefixSum(x) = prefixSum(y)-sum ，说明 x 到 y 这条路径的和是 sum

借助一个哈希表记录每条路径上，每个前缀和出现的次数，可以减少重复计算

```go
func pathSumCount0(root *TreeNode, sum int) int {
	// 记录当前路径上的前缀和，key为前缀和，value为前缀和的个数
	counts := make(map[int]int, 0) 
	// 前缀和为0的一条路径，方便边界处理，即节点值就是sum这种情况
	counts[0] = 1                  
	res := 0
	prefixSum := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 当前节点node的前缀和（即从root到当前节点这条路径的和）
		prefixSum += node.Val      
		// 如果当前节点之前已经有前缀和为 prefixSum-sum 的节点，说明那些节点到当前节点的和就是sum  
		res += counts[prefixSum-sum] 
		counts[prefixSum]++
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		counts[prefixSum]--
		prefixSum -= node.Val
	}

	dfs(root)
	return res
}
```
时间复杂度将为`O(n)`。