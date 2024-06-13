---
title: "23. 合并K个升序链表"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [23. 合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

难度困难

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

**示例 1：**

```
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
```

**示例 2：**

```
输入：lists = []
输出：[]
```

**示例 3：**

```
输入：lists = [[]]
输出：[]
```

**提示：**

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` 按 **升序** 排列
- `lists[i].length` 的总和不超过 `10^4`

链表定义：

```go
type ListNode struct {
	Val  int
	Next *ListNode
}
```

函数签名：

```go
func mergeKLists(lists []*ListNode) *ListNode
```

## 分析

### 控制步长的两两合并

首先需要实现合并两个有序列表的函数 merge，再一一合并所有的列表。

合并两个链表的函数如下：

```go
func merge(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	for p := dummy; l1 != nil || l2 != nil; p = p.Next {
		if l1 != nil && l2 != nil && l1.Val < l2.Val || l2 == nil {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
	}
	l1, dummy.Next = dummy.Next, nil
	return l1
}
```

现在一一合并所有的列表。如果写成如下这样：

```go
var r *ListNode
for _, v := range lists {
	r = merge(r, v)
}
return r
```

时间复杂度是 O(n*k^2)，其中 n 是原始单个链表的长度，这样的合并很不均衡，可以想象临近最后是一个很长的链表和一个很短的链表合并。如果能保证每次合并的两个链表规模相当，就能优化这个问题了。

```go
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	for interval := 1; interval < k; interval *= 2 {
		for i := 0; i+interval < k; i += interval * 2 {
			lists[i] = merge(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}
```

时间复杂度是 O(nklogk)，空间复杂度是 O(1)。

### 借助堆

```go
func mergeKLists(lists []*ListNode) *ListNode {
    k := len(lists)
    if k == 0 {
        return nil
    }
	h := &Heap{}
    for _, node := range lists {
        if node != nil {
            h.push(node)
        }
    }
    dummy := &ListNode{}
    p := dummy
    for h.Len() > 0 {
        cur := h.pop()
        p.Next = cur
        p = p.Next
        if cur.Next != nil {
            h.push(cur.Next)
        }
    }
    p, dummy.Next = dummy.Next, nil
    return p
}

type Heap struct {
    s []*ListNode
}
func (h *Heap) Len() int {return len(h.s)}
func (h *Heap) Less(i, j int) bool {return h.s[i].Val < h.s[j].Val}
func (h *Heap) Swap(i, j int) {h.s[i], h.s[j] = h.s[j], h.s[i]}
func (h *Heap) Push(x any) {h.s = append(h.s, x.(*ListNode))}
func (h *Heap) Pop() any {
    n := len(h.s)
    x := h.s[n-1]
    h.s = h.s[:n-1]
    return x
}
func (h *Heap) push(x *ListNode) {heap.Push(h, x)}
func (h *Heap) pop() *ListNode{return heap.Pop(h).(*ListNode)}
```

时间复杂度同样为 O(nklogk)，空间复杂度是 O(k)，维护了一个大小为 k 的堆 。
