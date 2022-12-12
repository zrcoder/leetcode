---
title: "LRU/LFU 缓存"
date: 2022-02-24T11:11:27+08:00
---

## [146. LRU 缓存](https://leetcode-cn.com/problems/lru-cache/)

难度中等

请你设计并实现一个满足 [LRU (最近最少使用) 缓存](https://baike.baidu.com/item/LRU) 约束的数据结构。

实现 `LRUCache` 类：

- `LRUCache(int capacity)` 以 **正整数** 作为容量 `capacity` 初始化 LRU 缓存
- `int get(int key)` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1` 。
- `void put(int key, int value)` 如果关键字 `key` 已经存在，则变更其数据值 `value` ；如果不存在，则向缓存中插入该组 `key-value` 。如果插入操作导致关键字数量超过 `capacity` ，则应该 **逐出** 最久未使用的关键字。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。

 

**示例：**

```
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
```

 

**提示：**

- `1 <= capacity <= 3000`
- `0 <= key <= 10000`
- `0 <= value <= 105`
- 最多调用 `2 * 105` 次 `get` 和 `put`

## 分析

需要这样一个数据结构来存储数据：要能维护数据读写的时间顺序，具体来说，插入的时候需要插入到最前边，因为新元素是`最近`使用的，获取或修改元素的时候，需要把元素移动到最前边。

考虑到时间复杂度，显然用链表比数组更合适。不过链表查询元素需要遍历，这可以借助一个哈希表来改进。

```go
type Pair struct {
	Key, Val int
}

type LRUCache struct {
	lis      *list.List
	dic      map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lis:      list.New(),
		dic:      make(map[int]*list.Element, 0),
		capacity: capacity,
	}
}

func (s *LRUCache) Get(key int) int {
	e, ok := s.dic[key]
	if !ok {
		return -1
	}
	s.lis.MoveToFront(e)
	return e.Value.(*Pair).Val
}

func (s *LRUCache) Put(key int, value int) {
	pair := &Pair{Key: key, Val: value}
	if e, ok := s.dic[key]; ok {
		_ = s.lis.Remove(e)
	} else if s.lis.Len() == s.capacity {
		back := s.lis.Remove(s.lis.Back()).(*Pair)
		delete(s.dic, back.Key)
	}
	s.dic[key] = s.lis.PushFront(pair)
}
```

曾经遇到亚麻的面试官，要求不能用标准库的 list，就商量了下，先用标准库写了，再自己手动实现了个双向链表。

标准库的实现比较奇怪，可以简单实现：用 head 和 tail 两个哨兵节点辅助，它们不实际存储内容，仅来维护链表结构，这样会简化代码。

参考如下：
```go
type Pair struct {
	Key, Val int
}

type LRUCache struct {
	lis      *List
	dic      map[int]*Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lis:      NewList(),
		dic:      make(map[int]*Element, 0),
		capacity: capacity,
	}
}

func (s *LRUCache) Get(key int) int {
	e, ok := s.dic[key]
	if !ok {
		return -1
	}
	s.lis.MoveToFront(e)
	return e.val
}

func (s *LRUCache) Put(key int, value int) {
	if e, ok := s.dic[key]; ok {
		s.lis.Remove(e)
	} else if len(s.dic) == s.capacity {
		back := s.lis.Back()
        s.lis.Remove(back)
		delete(s.dic, back.key)
	}
	s.dic[key] = s.lis.PushFront(key, value)
}

type Element struct {
	key, val  int
	pre, next *Element
}

type List struct {
	head, tail *Element
}

func NewList() *List {
	head, tail := new(Element), new(Element)
	head.next = tail
	tail.pre = head
	return &List{head, tail}
}

func (l *List) PushFront(key, val int) *Element {
	res := &Element{
		key: key,
		val: val,
	}
	next := l.head.next
	l.head.next = res
	res.pre = l.head
	res.next = next
	next.pre = res
	return res
}

func (l *List) Back() *Element {
	return l.tail.pre
}

func (l *List) RemoveBack() {
	l.Remove(l.Back())
}

func (l *List) Remove(e *Element) {
	pre, next := e.pre, e.next
	pre.next = next
	next.pre = pre
	e.pre = nil
	e.next = nil
}

func (l *List) MoveToFront(e *Element) {
	l.Remove(e)
	next := l.head.next
	l.head.next = e
	e.pre = l.head
	next.pre = e
	e.next = next
}
```

## [设计LFU缓存结构_牛客题霸_牛客网](https://www.nowcoder.com/practice/93aacb4a887b46d897b00823f30bfea1?tpId=295&tqId=1006014&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)

LFU 相比 LRU 要复杂一些。要删除旧元素时，

参考 LRU，同样用一个哈希表来记录每个key对应的信息；另需一个哈希表，以每个key出现的频率为键，值是一个双向链表，表示对应频率下出现的元素列表。

参考实现：
```go

func LFU(operators [][]int, k int) []int {
	var res []int
	cache := NewCache(k)
	for _, oper := range operators {
		if oper[0] == 1
			cache.Put(oper[1], oper[2])
		} else {
			res = append(res, cache.Get(oper[1]))
		}
	}
	return res
}

type LFUCache struct {
	kElements map[int]*list.Element
	fLists    map[int]*list.List
	capacity  int
	minFre    int
}

type entry struct {
	key, value, freq int
}

func NewCache(capacity int) *LFUCache {
	lfu := &LFUCache{
		kElements: make(map[int]*list.Element),
		fLists:    make(map[int]*list.List),
		capacity:  capacity,
	}
	return lfu
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.kElements[key]
	if !ok {
		return -1
	}
	res := node.Value.(*entry).value
	c.addFreq(node)
	return res
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}

	node, ok := c.kElements[key]
	if ok { //该键值已经存在
		node.Value.(*entry).value = value
		c.addFreq(node)
		return
	}
    
	//该键值不存在
	if len(c.kElements) == c.capacity {
		c.remove()
	}
	kv := &entry{key: key, value: value, freq: 1}
    if c.fLists[1] == nil {
        c.fLists[1] = list.New()
    }
	node = c.fLists[1].PushFront(kv)
	c.kElements[key] = node
	c.minFre = 1
}

func (c *LFUCache) remove() {
	l := c.fLists[c.minFre]
	node := l.Back()
	l.Remove(node)
	delete(c.kElements, node.Value.(*entry).key)
}

func (c *LFUCache) addFreq(node *list.Element) {
	//原频率中删除
	kv := node.Value.(*entry)
	oldList := c.fLists[kv.freq]
	oldList.Remove(node)

	//更新minfreq
	if oldList.Len() == 0 && c.minFre == kv.freq {
		c.minFre++
	}

	//放入新的频率链表
	kv.freq++
	if _, ok := c.fLists[kv.freq]; !ok {
		c.fLists[kv.freq] = list.New()
	}
	newList := c.fLists[kv.freq]
	node = newList.PushFront(kv)
	c.kElements[kv.key] = node
}

```