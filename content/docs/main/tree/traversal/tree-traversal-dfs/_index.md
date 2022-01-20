---
title: "一般树的 DFS 深度遍历"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

不妨以前序遍历为例，需注意多叉树没有中序遍历

递归版
```go
func preorder(root *TreeNode) []string {
	var result []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, v := range node.Children {
			dfs(v)
		}
	}
	dfs(root)
	return result
}
```

迭代版，借助一个栈实现
```go
func preorder1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var result []string
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}
	return result
}
```