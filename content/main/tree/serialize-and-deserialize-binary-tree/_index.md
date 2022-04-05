---
title: "297. 二叉树的序列化与反序列化"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [297. 二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree)
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

```
示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
```
说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

二叉树的定义、 Codec 类及序列化、反序列化函数定义如下：
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string

func (c *Codec) deserialize(data string) *TreeNode
```
## 分析
序列化：使用一种方式遍历二叉树，将节点值存入结果字符串，为了后序反序列化，要有分隔符，不妨用逗号。

反序列化，先使用分隔符切割字符串得到所有节点的值，再按照序列化时采取的遍历顺序构造出树。

> 因为二叉树可能不均衡，以上两步需要额外记录并处理空节点的情况。
> 正如示例的例子，必须记录 "[1,2,3,null,null,4,5]"，而不能把两个 null 丢弃，这样反序列化会因信息不足得到错误的树结构。

遍历方式既可以用 dfs， 也可以用 bfs。

### DFS 先序遍历

```go
type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}
// 序列化的递归写法非常简单，如下。
// 注意字符串拼接性能不佳，可以改成用 strings.Builder 或bytes.Buffer优化
func (c *Codec) serialize(root *TreeNode) string {
    var pre func(*TreeNode) string
    pre = func(root *TreeNode) string {
        if root == nil {
            return "#"
        }
        return strconv.Itoa(root.Val)+","+pre(root.Left)+","+pre(root.Right)
    }
    return pre(root)
}
// 使用 strings.Builder{} 优化字符串操作
func (c *Codec) serialize(root *TreeNode) string {
    buf := strings.Builder{}
    var pre func(*TreeNode)
    pre = func(root *TreeNode) {
        if root == nil {
            buf.WriteString("#,")
            return
        }
        buf.WriteString(strconv.Itoa(root.Val))
        buf.WriteString(",")
        pre(root.Left)
        pre(root.Right)
    }
    pre(root)
    return buf.String()
}

func (c *Codec) deserialize(data string) *TreeNode {
    arr := strings.Split(data, ",")
    index := 0
    var pre func() *TreeNode
    pre = func() *TreeNode {
        if index >= len(arr) || arr[index] == "#" {
            return nil
        }
        v, _ := strconv.Atoi(arr[index])
        root := &TreeNode{Val: v}
        index++
        root.Left = pre()
        index++
        root.Right = pre()
        return root
    }
    return pre()
}

```

这里可能有个疑问：反序列化，仅凭先序遍历结果，怎么能确定树结构呢？

实际上这里`额外存储了空节点信息`，是可以唯一确定树结构的。

### BFS
```go
func (c *Codec) serializeBfs(root *TreeNode) string {
	queue := list.New()
	queue.PushBack(root)
	b := strings.Builder{}
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node == nil {
			b.WriteString("#,")
			continue
		}
		b.WriteString(strconv.Itoa(node.Val))
		b.WriteByte(',')
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
	return b.String()
}

func (c *Codec) deserializeBfs(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0
	val, err := strconv.Atoi(values[index])
	if err != nil { // values[0] == "#"
		return nil
	}
	root := &TreeNode{Val: val}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		index++
		if values[index] != "#" {
			val, _ = strconv.Atoi(values[index])
			left := &TreeNode{Val: val}
			node.Left = left
			queue.PushBack(left)
		}
		index++
		if values[index] != "#" {
			val, _ = strconv.Atoi(values[index])
			right := &TreeNode{Val: val}
			node.Right = right
			queue.PushBack(right)
		}
	}
	return root
}
```

> 有个简化代码的思路：
> 也可以借助标准库 json 包，将数组做序列化和反序列化。
> 那么剩下的工作只需要在树和数组间做转换。
> 当然这么做的效率差一点。
