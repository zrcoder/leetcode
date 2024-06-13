---
title: "JZ36 二叉搜索树与双向链表"
date: 2022-12-20T16:58:28+08:00
---

##  [JZ36 二叉搜索树与双向链表_牛客题霸_牛客网](https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5)

中等  通过率：31.85%

## 描述

输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示

![](https://uploadfiles.nowcoder.com/images/20210605/557336_1622886924427/E1F1270919D292C9F48F51975FD07CE2)

数据范围：输入二叉树的节点数 0≤n≤1000，二叉树中每个节点的值 0≤val≤1000  
要求：空间复杂度O(1)（即在原树上操作），时间复杂度 O(n)

注意:

1.要求不能创建任何新的结点，只能调整树中结点指针的指向。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继  
2.返回链表中的第一个节点的指针  
3.函数返回的TreeNode，有左右指针，其实可以看成一个双向链表的数据结构

4.你不用输出双向链表，程序会根据你的返回值自动打印输出

### 输入描述

二叉树的根节点

### 返回值描述

双向链表的其中一个头节点。

## 示例1

输入：

{10,6,14,4,8,12,16}

返回值：

From left to right are:4,6,8,10,12,14,16;From right to left are:16,14,12,10,8,6,4;

说明：

输入题面图中二叉树，输出的时候将双向链表的头节点返回即可。

## 示例2

输入：

{5,4,#,3,#,2,#,1}

返回值：

From left to right are:1,2,3,4,5;From right to left are:5,4,3,2,1;

说明：

```
                5
              /
            4
          /
        3
      /
    2
  /
1
```

树的形状如上图

预定义类型及函数签名：

```go
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param pRootOfTree TreeNode类
 * @return TreeNode类
 */
func Convert(pRootOfTree *TreeNode) *TreeNode
```

## 分析

BST 的中序遍历结果就是有序的。只需要中序遍历即可。

```go
func Convert(pRootOfTree *TreeNode) *TreeNode {
	var pre, res *TreeNode
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != nil {
			pre.Right = root
			root.Left = pre
		} else {
			res = root
		}
		pre = root
		dfs(root.Right)
	}
	dfs(pRootOfTree)
	return res
}
```

时间复杂度：`O(n)`，空间复杂度：`O(h)`。其中 n 为节点总数，h 为树的高度，对应递归栈的大小。
