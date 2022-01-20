---
title: "树的遍历"
date: 2021-04-23T10:20:48+08:00
weight: -1
---

## 一般树
不妨假设每个结点的值是一个字符串
```
              A
            / | \
           B  C  D
          /|     |
         E F     G
        /|\
       H I J
```
```go
type TreeNode struct {
	Children []*TreeNode
	Val string
}
```

也可以用数组或哈希表来代表结点，如[Trie前缀树的一个实现](../../design/implement-trie-prefix-tree)及[树的路径和问题探讨](../solutions/path-sum/d.go)中getPath函数的入参relations
* [BFS层次遍历](tree-traversal-levelorder)
* [DFS深度遍历](tree-traversal-dfs)

## 二叉树
对于二叉树，DFS又可细分为前序、中序、后序遍历
* [层次遍历](../binary-tree-level-order-traversal)
* [之字形层次遍历](../binary-tree-zigzag-level-order-traversal)
* [先序遍历](../binary-tree-preorder-traversal)
* [中序遍历](../binary-tree-inorder-traversal)
* [后序遍历](../binary-tree-postorder-traversal)
* [二叉树 morris 遍历](binary-tree-morris)
## 限定条件遍历
在遍历时，可以有一些限定条件，比如统计从根结点到叶子结点路径和为定值的路径；  
这里可以增加额外辅助数据结构如栈来记录路径，同时做好回溯。  
参考[树的路径和问题探讨](../solutions/path-sum/d.go)


