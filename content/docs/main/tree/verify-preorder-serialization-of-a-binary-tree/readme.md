---
title: "331. 验证二叉树的前序序列化"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [331. 验证二叉树的前序序列化](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/)

难度中等

序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 `#`。

```
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
```

例如，上面的二叉树可以被序列化为字符串 `"9,3,4,#,#,1,#,#,2,#,6,#,#"`，其中 `#` 代表一个空节点。

给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。

每个以逗号分隔的字符或为一个整数或为一个表示 `null` 指针的 `'#'` 。

你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 `"1,,3"` 。

**示例 1:**

```
输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
```

**示例 2:**

```
输入: "1,#"
输出: false
```

**示例 3:**

```
输入: "9,#,#,1"
输出: false
```

函数签名：

```go
func isValidSerialization(preorder string) bool
```

## 分析

非常有意思的一个问题。解法较多。

### 模拟建树

可以参考 [297. 二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) 问题中先序遍历解法的 `deserialize` 函数，根据先序遍历字符串尝试模拟构建一棵树。

```go
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	index := 0
	var help func() bool
	help = func() bool {
		if index == len(nodes) {
			return false
		}
		// root
		isLeaf := nodes[index] == "#"
		index++
		if isLeaf {
			return true
		}
		// left and right
		return help() && help()
	}
	return help() && index == len(nodes)
}
```

时空复杂度都是 O(n)， 其中 n 是节点总数。

### 不断消除叶子节点

以示例1为例，观察可以看到，只有叶子节点在序列化后是 x## 的格式，这样可以在序列化字符串中不断把 x## 替换成 #，即不断删除叶子节点，最后能得到一个 “#” 字符串，则是合法的，反之不合法。

```go
func isValidSerialization(preorder string) bool {
	s := replaceNums(preorder)
	for strings.Contains(s, "*##") {
		s = strings.ReplaceAll(s, "*##", "#")
	}
	return s == "#"
}

// 替换所有数字为"*"且删除所有逗号
func replaceNums(s string) string {
	arr := strings.Split(s, ",")
	for i, v := range arr {
		if v != "#" {
			arr[i] = "*"
		}
	}
	return strings.Join(arr, "")
}
```

replaceNums 还可以用上正则，精简代码：

```go
// 替换所有数字为"*"且删除所有逗号
func replaceNums(s string) string {
	reg := regexp.MustCompile("[0-9]+")
	s = reg.ReplaceAllString(s, "*")
	return strings.ReplaceAll(s, ",", "")
}
```

时空复杂度都是 O(n)。

### 维护待填充点

根据 preorder 构建二叉树的过程中，可以关注下每次待填充的点的数量。

如果出现一个真实节点，那么待填充的点减去1同时要加上2（这个点有左右两个孩子）；如果遇到一个空节点，那么待填充点需要减去1。

用 points 维护待填充点的数量。

如果 preorder 合法，那么遍历过程中 points 一直大于0且遍历完成后points等于0。反过来，如果遍历过程中 points 一直大于0且遍历完成后points等于0，则preorder合法。

简而言之：`preorder 合法 <=> 遍历过程中 points 一直大于0且遍历完成后points等于0`。

根据遍历和判断标准，points 初始值应为 1。

```go
func isValidSerialization(preorder string) bool {
	points := 1
	nodes := strings.Split(preorder, ",")
	for _, v := range nodes {
		if points == 0 {
			return false
		}
		if v == "#" {
			points--
		} else {
			points++
		}
	}
	return points == 0
}
```

时空复杂度都是 O(n)，也可以不用预处理 preorder 为一个数组，空间复杂度能降低到常数空间。

> 维护 points 这个解法，也可以从另一个角度去理解：
>
> 将树的每条边看成从父节点指向子节点的箭头，那么每个节点都有了入度和出度。
>
> points 的值即为当前所有节点的出度减去入度的值。
>
> 遍历过程中出度需要大于入度，遍历结束出度需要等于入度。

## 小结

第一种模拟构建树的方法比较通用，后边两种解法比较难想到，实际是先假设合法，观察得出必要条件，反过来发现该必要条件也可以作为充分条件，从而解决问题。

## 扩展

如果要判断后序遍历序列或中序遍历序列或层序遍历序列是否合法呢？

