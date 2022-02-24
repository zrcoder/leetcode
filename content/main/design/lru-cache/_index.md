---
title: "LRU 缓存"
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