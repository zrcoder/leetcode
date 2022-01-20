---
title: "946. 验证栈序列"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [946. 验证栈序列](https://leetcode-cn.com/problems/validate-stack-sequences/)

难度中等

给定 `pushed` 和 `popped` 两个序列，每个序列中的 **值都不重复**，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 `true`；否则，返回 `false` 。

 

**示例 1：**

```
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
```

**示例 2：**

```
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。
```

 

**提示：**

1. `0 <= pushed.length == popped.length <= 1000`
2. `0 <= pushed[i], popped[i] < 1000`
3. `pushed` 是 `popped` 的排列。

函数签名：

```go
func validateStackSequences(pushed []int, popped []int) bool
```

## 分析

### 模拟

很难总结出一个规律，可以用一个栈实际模拟。

```go
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	for _, v := range pushed {
		stack = append(stack, v)
		// 尽可能地把 popped 里的元素从 stack 出栈
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(popped) == 0
}
```

时间复杂度、空间复杂度都是 O(n)。

### 构造二叉树

这是一个非常特别的思路。

可以把 pushed、popped 看作一棵二叉树的前序、中序遍历序列，问题转化为是否能由这两个序列还原构造出一棵二叉树。

```go
func validateStackSequences(pushed []int, popped []int) bool {
	return canBuildBinaryTree(pushed, popped)
}

func canBuildBinaryTree(preorder []int, inorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	root := preorder[0]
	i := search(inorder, root)
	if i == -1 {
		return false
	}
	canBuildLeft := canBuildBinaryTree(preorder[1:i+1], inorder[:i])
	canBuildRight := canBuildBinaryTree(preorder[i+1:], inorder[i+1:])
	return canBuildLeft && canBuildRight
}

func search(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}
```

复杂度同模拟法。

### 小结

对比两个解法，实际上一个问题的两种解决方案。一个是迭代写法，另一个是递归写法。