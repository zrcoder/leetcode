---
title: "341. 扁平化嵌套列表迭代器"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [栈, 递归]
---

## [341. 扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator/)

难度中等

给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。

列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

**示例 1:**

```
输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
```

**示例 2:**

```
输入: [1,[4,[6]]]
输出: [1,4,6]
解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
```

已有的 NestedInteger 相关 API（实现不给出）：

```go
type NestedInteger struct {}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {}
```

待实现的迭代器：

```go
type NestedIterator struct {
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
}

func (it *NestedIterator) Next() int {
}

func (it *NestedIterator) HasNext() bool {
}
```

## 分析

这是个非常有意思的问题。首先 NestedInteger 就很有趣，可以想象一下它有什么应用场景，题目只是给出了这个类型的 API，但没有给出其实现；需要根据给定的 API 来完成迭代器相关实现。

先写迭代器，后边尝试把  NestedInteger 的所有 API 实现一把。

### 初始化压扁列表

给定一个 []*NestedInteger 类型，直接压扁成 []int，可以用递归，比较容易：

```go
type NestedIterator struct {
    nums []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var nums []int
    var flat func(nestedList []*NestedInteger)
    flat = func(nestedList []*NestedInteger) {
        for _, v := range nestedList {
            if v.IsInteger() {
                nums = append(nums, v.GetInteger())
            } else {
                flat(v.GetList())
            }
        }
    }
    flat(nestedList)
    return &NestedIterator{nums: nums}
}

func (it *NestedIterator) Next() int {
    res := it.nums[0]
    it.nums = it.nums[1:]
    return res
}

func (it *NestedIterator) HasNext() bool {
    return len(it.nums) > 0
}
```

初始化的时间复杂度是 O(n)， n 为所有数组个数。Next 和 HasNext 都是常数级复杂度。

空间复杂度是 O(n)。

### 借助栈

上边在初始化时就整个压扁嵌套列表的做法，其实不太符合迭代器的约束：

```
迭代器不应该直接存储所有数字，而应提供访问途径。
迭代有条件终止(如键值查找)时初始化方法的全局开销非必要。
初始化迭代器后，迭代过程中无法处理List中某数字值改变的场景。
```

考虑初始化不做复杂工作，在 Next 和 HasNext 方法里做。

```go
type NestedIterator struct {
    items []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{items: nestedList}
}

func (it *NestedIterator) Next() int {
    // 保证调用 Next 之前会调用 HasNext
    res := it.items[0].GetInteger()
    it.items = it.items[1:]
    return res
}

func (it *NestedIterator) HasNext() bool {
    for len(it.items) > 0 && !it.items[0].IsInteger() {
        // 展开第一个元素，再追加后边的——比较耗费性能
        it.items = append(it.items[0].GetList(), it.items[1:]...)
    }
    return len(it.items) > 0
}
```

注意上边 `HasNext` 函数的注释，可以借助一个栈来优化，直接看代码：

```go
type NestedIterator struct {
    // 将列表视作一个队列，栈中直接存储该队列
    stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
    queue := it.stack[len(it.stack)-1]
    val := queue[0].GetInteger()
    it.stack[len(it.stack)-1] = queue[1:]
    return val
}

func (it *NestedIterator) HasNext() bool {
    for len(it.stack) > 0 {
        queue := it.stack[len(it.stack)-1]
        if len(queue) == 0 { // 当前队列为空，出栈
            it.stack = it.stack[:len(it.stack)-1]
            continue
        }
        nest := queue[0]
        if nest.IsInteger() {
            return true
        }
        // 若队首元素为列表，则将其弹出队列并打平入栈
        it.stack[len(it.stack)-1] = queue[1:]
        it.stack = append(it.stack, nest.GetList())
    }
    return false
}
```

初始化、Next 方法时间复杂度都是常数级，HasNext 方法均摊复杂度也是常数级。

空间复杂度，主要在栈的大小，最坏 O(n)。

至此，迭代器写完，问题解决。

### 实现 NestedInteger

我还是很好奇 NestedInteger 各个 Api 的实现，尝试写了一下：

```go
type Any interface{}

type NestedInteger struct {
    // `int` or `[]*NestedInteger`
    Val Any
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (ni NestedInteger) IsInteger() bool {
    _, ok := ni.Val.(int)
    return ok
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (ni NestedInteger) GetInteger() int {
    if v, ok := ni.Val.(int); ok {
        return v
    }
    return 0
}

// Set this NestedInteger to hold a single integer.
func (ni *NestedInteger) SetInteger(value int) {
    ni.Val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (ni *NestedInteger) Add(elem NestedInteger) {
    if ni.IsInteger() {
        ni.Val = []*NestedInteger{{ni.Val}, &elem}
    } else {
        list := ni.Val.([]*NestedInteger)
        list = append(list, &elem)
        ni.Val = list
    }
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (ni NestedInteger) GetList() []*NestedInteger {
    if ni.IsInteger() {
        return nil
    }
    return ni.Val.([]*NestedInteger)
}
```

## 扩展

考虑 NestedInteger 的序列化及反序列化应该如何实现。

可以直接借助标准库 `ecoding/json`来做，代码会非常简洁，不过基于上边的实现，序列化后的格式类似：`{"Val":[{"Val":123},{"Val":[{"Val":456},{"Val":789}]}]}`，稍微有点冗余。反序列化的结果也不尽如预期。考虑序列化结果形如：`[123,[456,[789]]]`，可以借助递归，参考实现如下：

```go
func Marshal(ni *NestedInteger) string {
    if ni == nil {
        return ""
    }
    if ni.IsInteger() {
        return strconv.Itoa(ni.GetInteger())
    }

    list := ni.GetList()
    buf := strings.Builder{}
    buf.WriteByte('[')
    for i, v := range list {
        buf.WriteString(Marshal(v))
        if i < len(list)-1 {
            buf.WriteByte(',')
        }
    }
    buf.WriteByte(']')
    return buf.String()
}
```

反序列化稍微复杂一点，Leetcode 有个现成的题目，如下：

#### [385. 迷你语法分析器](https://leetcode-cn.com/problems/mini-parser/description/ "https://leetcode-cn.com/problems/mini-parser/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (42.16%) | 114   | -        |

给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 `NestedInteger` 。

列表中的每个元素只可能是整数或整数嵌套列表

**示例 1：**

```
输入：s = "324",
输出：324
解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
```

**示例 2：**

```
输入：s = "[123,[456,[789]]]",
输出：[123,[456,[789]]]
解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
1. 一个 integer 包含值 123
2. 一个包含两个元素的嵌套列表：
    i.  一个 integer 包含值 456
    ii. 一个包含一个元素的嵌套列表
         a. 一个 integer 包含值 789
```

**提示：**

- `1 <= s.length <= 5 * 104`
- `s` 由数字、方括号 `"[]"`、负号 `'-'` 、逗号 `','`组成
- 用例保证 `s` 是可解析的 `NestedInteger`
- 输入中的所有值的范围是 `[-106, 106]`

递归或者借助栈来解决：

```go
func deserialize(s string) *NestedInteger {
    if s == "" {
        return nil
    }
    if val, err := strconv.Atoi(s); err == nil {
        return genNestedInteger(val)
    }
    sign := 1
    num := 0
    stk := []*NestedInteger{}
    for i, v := range s {
        switch v {
        case '[':
            stk = append(stk, &NestedInteger{})
        case ',', ']':
            if unicode.IsDigit(rune(s[i-1])) {
                stk[len(stk)-1].Add(*genNestedInteger(num * sign))
                num, sign = 0, 1
            }
            if v == ']' && len(stk) > 1 {
                n := len(stk)
                stk[n-2].Add(*stk[n-1])
                stk = stk[:n-1]
            }
        case '-':
            sign = -1
        default: // digit
            num = num*10 + int(v-'0')
        }
    }
    return stk[0] // stk must have only one element
}

func genNestedInteger(val int) *NestedInteger {
    res := &NestedInteger{}
    res.SetInteger(val)
    return res
}
```
