---
title: "102. 二叉树的层序遍历"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal)
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:  
给定二叉树: [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```
返回其层次遍历结果：
```
[
  [3],
  [9,20],
  [15,7]
]
```
二叉树定义为：
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

有两种通用的遍历树的策略：

深度优先搜索 depth first search（DFS）  
在这个策略中，采用深度作为优先级，以便从根开始一直到达某个确定的叶子，然后再返回根到达另一个分支。  
深度优先搜索策略又可以根据根节点、左孩子和右孩子的相对顺序被细分为先序遍历，中序遍历和后序遍历。  

广度优先搜索breadth first search（BFS）  
按照高度顺序一层一层的访问整棵树，高层次的节点将会比低层次的节点先被访问到。  

本问题就是用广度优先搜索遍历二叉树。

对于 BFS，一般用迭代方式实现，递归方式也可以，但较麻烦，下边只介绍迭代实现

1.使用数组迭代
```go
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	currentLever := []*TreeNode{root}
	for len(currentLever) > 0 {
		var values []int
		var nextLever []*TreeNode
		for _, node := range currentLever {
			values = append(values, node.Val)
			if node.Left != nil {
				nextLever = append(nextLever, node.Left)
			}
			if node.Right != nil {
				nextLever = append(nextLever, node.Right)
			}
		}
		result = append(result, values)
		currentLever = nextLever
	}
	return result
}
```

2.使用队列迭代
```go
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var values []int
		currentLen := queue.Len()
		for i := 0; i < currentLen; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			values = append(values, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, values)
	}
	return result
}
```