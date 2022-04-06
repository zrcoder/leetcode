---
title: "380. O(1) 时间插入、删除和获取随机元素"
date: 2022-04-06T17:49:48+08:00
draft: true
---

## [380. O(1) 时间插入、删除和获取随机元素](https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/ "https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (50.67%) | 456   | -        |

实现`RandomizedSet` 类：

- `RandomizedSet()` 初始化 `RandomizedSet` 对象
- `bool insert(int val)` 当元素 `val` 不存在时，向集合中插入该项，并返回 `true` ；否则，返回 `false` 。
- `bool remove(int val)` 当元素 `val` 存在时，从集合中移除该项，并返回 `true` ；否则，返回 `false` 。
- `int getRandom()` 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 **相同的概率** 被返回。

你必须实现类的所有函数，并满足每个函数的 **平均** 时间复杂度为 `O(1)` 。

**示例：**

```
输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
```

**提示：**

- `-231 <= val <= 231 - 1`
- 最多调用 `insert`、`remove` 和 `getRandom` 函数 `2 *` `105` 次
- 在调用 `getRandom` 方法时，数据结构中 **至少存在一个** 元素。

## 分析

能不能直接用一个map呢？如下：

```go
type RandomizedSet struct {
    Set map[int]bool
}

func Constructor() RandomizedSet {
    return RandomizedSet{Set: map[int]bool{}}
}

func (s *RandomizedSet) Insert(val int) bool {
    if s.Set[val] {
        return false
    }
    s.Set[val] = true
    return true
}

func (s *RandomizedSet) Remove(val int) bool {
    if !s.Set[val] {
        return false
    }
    return true
}

func (s *RandomizedSet) GetRandom() int {
    for k := range s.Set {
        return k
    }
    return -1
}
```

所有操作都是常数级复杂度，`GwtRandom` 刚进循环就`return`，结合哈希表遍历是随机返回键值，这样做好像是对的。但是，并没有做到每个元素应该有 **相同的概率** 被访问，这是因为哈希表底层是数组+链表实现，尤其脸表部分，不可能等概率访问每个元素。

要能等概率地返回一个随机元素，必须将元素存到数组/切片里才行。

但是数组增删元素怎么做到复杂度为常数级呢？

实际上，这里并不关心元素在数组里的顺序，所以插入元素就可以简单地在末尾追加；删除元素怎么办？可以把带删除元素和最后一个元素交换，再把数组长度减1即可！

当然，为了能迅速获取某个元素在数组中的索引，需要额外维护一个哈希表。

```go
type RandomizedSet struct {
    index map[int]int
    slice []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{index: map[int]int{}}
}

func (s *RandomizedSet) Insert(val int) bool {
    if _, ok := s.index[val]; ok {
        return false
    }
    s.index[val] = len(s.slice)
    s.slice = append(s.slice, val)
    return true
}

func (s *RandomizedSet) Remove(val int) bool {
    i, ok := s.index[val]
    if !ok {
        return false
    }

    delete(s.index, val)

    n := len(s.slice)
    if i == n-1 { // 要删的元素就是最后一个
        s.slice = s.slice[:n-1]
        return true
    }

    last := s.slice[n-1]
    s.index[last] = i // 更新last索引记录
    s.slice[i] = last
    s.slice = s.slice[:n-1]
    return true
}

func (s *RandomizedSet) GetRandom() int {
    i := rand.Intn(len(s.slice))
    return s.slice[i]
}
```
