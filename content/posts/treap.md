---
title: "Treap"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

简单实现的二叉搜索树（BST），在特定的插入和删除操作后容易退化成一条链表，性能大幅度降低。

如果能维持相对平衡，则增删查元素的复杂度都会是对数级。像 bTree、AVL树、红黑树等都是比较平衡的 BST，但是实现较复杂。

Treap 即树堆，除了各个节点的值满足 BST 特性，每个节点额外附加一个权重属性。

新建节点时权重取随机值，在增删元素时通过旋转等操作保证整棵树所有节点的权重满足堆的性质。这样整体保证了树基本平衡。
## API
```go
type Treap struct {
	root *Node
}

func (t *Treap) Put(val int) {
	t.root = t.put(t.root, val)
}

func (t *Treap) Delete(val int) {
	t.root = t.delete(t.root, val)
}

func (t *Treap) Search(val int) bool {
	return t.root.search(val)
}
```
其中节点定义如下：
```go
type Node struct {
	// 左右孩子
	ch [2]*Node
	// 权重、值、个数
	priority, val, cnt int
}
```
## search 的实现
相比而言，search 方法最容易理解。
```go
func (o *Node) search(val int) bool {
	for cur := o; cur != nil; {
		if cur.val == val {
			return true
		}
		if cur.val < val {
			cur = cur.ch[1]
		} else {
			cur = cur.ch[0]
		}
	}
	return false
}
```
## put
先根据 BST 的特性将元素插入为一个叶子节点，新建节点时赋予随机权重值，如果新插入的节点权重比父节点权重小，则不断上浮该节点直到权重满足堆的性质。

还要注意不能破坏 BST 的性质，这样上浮操作可以通过旋转来完成。
```go
func (t *Treap) put(root *Node, val int) *Node {
	if root == nil {
		return &Node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := root.cmp(val); d == -1 {
		root.cnt++
	} else {
		root.ch[d] = t.put(root.ch[d], val)
		if root.ch[d].priority > root.priority {
			// 为了达到上浮的效果，如果是左孩子就右旋，如果是右孩子就左旋
			root = root.rotate(d ^ 1)
		}
	}
	return root
}
```
cmp 方法如下：
```go
func (o *Node) cmp(b int) int {
	switch {
	case o.val > b:
		// b 应该在左子树中
		return 0
	case o.val < b:
		// b 应该在右子树中
		return 1
	default:
		// b 就是当前节点的值
		return -1
	}
}
```
rotate 方法如下：
```go
func (o *Node) rotate(d int) *Node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}
```
## delete
可以用堆的方式删除，只需要把要删除的节点旋转到叶节点上，然后直接删除就可以了。

也可以用 BST 的方式删除，如下：
```go
func (t *Treap) delete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if d := root.cmp(val); d != -1 {
		root.ch[d] = t.delete(root.ch[d], val)
		return root
	}
	if root.cnt > 1 {
		root.cnt--
		return root
	}
	if root.ch[1] == nil {
		res := root.ch[0]
		root.ch[0] = nil
		return res
	}
	if root.ch[0] == nil {
		res := root.ch[1]
		root.ch[1] = nil
		return res
	}
	d := 0
	if root.ch[0].priority > root.ch[1].priority {
		d = 1
	}
	root = root.rotate(d)
	root.ch[d] = t.delete(root.ch[d], val)
	return root
}
```
## 扩展 API
可以很方便地写出查找树中最小值、最大值的方法，还可以在树中查找比给定值大的最小元素或比给定值小的最大元素。
## 完整代码
```go
type Treap struct {
	root *Node
}

func (t *Treap) Put(val int) {
	t.root = t.put(t.root, val)
}

func (t *Treap) Delete(val int) {
	t.root = t.delete(t.root, val)
}

func (t *Treap) Search(val int) bool {
	return t.root.search(val)
}

func (t *Treap) GetMin() int {
	return t.getMinMax(0)
}

func (t *Treap) GetMax() int {
	return t.getMinMax(1)
}

// 在树中查找比 val 大的最小元素
func (t *Treap) UpperBound(val int) *Node {
	var res *Node
	for cur := t.root; cur != nil; {
		if cur.cmp(val) == 0 { // 在左子树中查找
			res = cur
			cur = cur.ch[0]
		} else { // 在右子树中查找
			cur = cur.ch[1]
		}
	}
	return res
}

func (t *Treap) put(root *Node, val int) *Node {
	if root == nil {
		return &Node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := root.cmp(val); d == -1 {
		root.cnt++
	} else {
		root.ch[d] = t.put(root.ch[d], val)
		if root.ch[d].priority > root.priority {
			// 为了达到上浮的效果，如果是左孩子就右旋，如果是右孩子就左旋
			root = root.rotate(d ^ 1)
		}
	}
	return root
}

func (t *Treap) delete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if d := root.cmp(val); d != -1 {
		root.ch[d] = t.delete(root.ch[d], val)
		return root
	}
	if root.cnt > 1 {
		root.cnt--
		return root
	}
	if root.ch[1] == nil {
		res := root.ch[0]
		root.ch[0] = nil
		return res
	}
	if root.ch[0] == nil {
		res := root.ch[1]
		root.ch[1] = nil
		return res
	}
	d := 0
	if root.ch[0].priority > root.ch[1].priority {
		d = 1
	}
	root = root.rotate(d)
	root.ch[d] = t.delete(root.ch[d], val)
	return root
}

type Node struct {
	// 左右孩子
	ch [2]*Node
	// 权重、值、个数
	priority, val, cnt int
}

func (o *Node) cmp(b int) int {
	switch {
	case o.val > b:
		// b 应该在左子树中
		return 0
	case o.val < b:
		// b 应该在右子树中
		return 1
	default:
		// b 就是当前节点的值
		return -1
	}
}

func (o *Node) search(val int) bool {
	for cur := o; cur != nil; {
		if cur.val == val {
			return true
		}
		if cur.val < val {
			cur = cur.ch[1]
		} else {
			cur = cur.ch[0]
		}
	}
	return false
}

func (o *Node) rotate(d int) *Node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

func (t *Treap) getMinMax(d int) int {
	var pre *Node
	cur := t.root
	for cur != nil {
		pre = cur
		cur = cur.ch[d]
	}
	if pre == nil {
		return 0
	}
	return pre.val
}
```
## 应用
- [132模式](../../dp-and-greedy/132-pattern/readme.md)
- [优势洗牌](../../dp-and-greedy/advantage-shuffle/readme.md)
## 参考
[图文详解Treap](https://blog.csdn.net/yang_yulei/article/details/46005845)