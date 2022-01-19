---
title: "745. 前缀和后缀搜索"
date: 2021-05-01T21:36:54+08:00
weight: 50
tags: [哈希表, Trie树]
---

## [745. 前缀和后缀搜索](https://leetcode-cn.com/problems/prefix-and-suffix-search/)

给定多个 `words`，`words[i]` 的权重为 `i` 。

设计一个类 `WordFilter` 实现函数`WordFilter.f(String prefix, String suffix)`。这个函数将返回具有前缀 `prefix` 和后缀`suffix` 的词的最大权重。如果没有这样的词，返回 -1。

**例子:**

```
输入:
WordFilter(["apple"])
WordFilter.f("a", "e") // 返回 0
WordFilter.f("b", "") // 返回 -1
```

**注意:**

1. `words`的长度在`[1, 15000]`之间。
2. 对于每个测试用例，最多会有`words.length`次对`WordFilter.f`的调用。
3. `words[i]`的长度在`[1, 10]`之间。
4. `prefix, suffix`的长度在`[0, 10]`之前。
5. `words[i]`和`prefix, suffix`只包含小写字母。

## 分析

### 朴素实现

一开始没有太好的思路，先写朴素实现，虽然超时了～

```go
type WordFilter struct {
    words []string
}

func Constructor(words []string) WordFilter {
    return WordFilter{words: words}
}


func (f *WordFilter) F(prefix string, suffix string) int {
    for i := len(f.words)-1; i >= 0; i-- {
        v := f.words[i]
        if strings.HasPrefix(v, prefix) && strings.HasSuffix(v, suffix) {
            return i
        }
    } 
    return -1
}
```

### 存储每个单词所有前缀+后缀组合

可以事先枚举每个单词所有`前缀+后缀`，存储起来，在后续查找时直接查找是否已经存储过。

可以用一个哈希表来存储，键就是`前缀+“#”+后缀`，值就是原始索引，这样可以在常数级复杂度做每次查找。

可以把哈希表改成 Trie 树来节省内存，不过在这个问题里，实际测试哈希表无论时间还是空间复杂度都更优。

{{< tabs >}}

{{% tab name="哈希表" %}}

```go
type WordFilter struct {
	memo map[string]int
}

func Constructor(words []string) WordFilter {
    m := map[string]int{}
    for i, v := range words {
        for j := 0; j <= len(v); j++ {
            pre := v[:j]
            for k := 0; k <= len(v); k++ {
                suf := v[k:]
                m[pre+"#"+suf] = i
            }
        }
    }
	return WordFilter{memo: m}
}

func (f *WordFilter) F(prefix string, suffix string) int {
    if index, ok := f.memo[prefix+"#"+suffix]; ok {
        return index
    }
    return -1
}
```

{{% /tab %}}

{{% tab name="Trie 树" %}}

```go
type WordFilter struct {
	tree *Trie
}

func Constructor(words []string) WordFilter {
    tree := &Trie{links: map[byte]*Trie{}}
    for i, v := range words {
        for j := 0; j <= len(v); j++ {
            pre := v[:j]
            for k := 0; k <= len(v); k++ {
                suf := v[k:]
                tree.insert(pre+"#"+suf, i)
            }
        }
    }
	return WordFilter{tree: tree}
}

func (f *WordFilter) F(prefix string, suffix string) int {
    return f.tree.find(prefix+"#"+suffix)
}

```

{{% /tab %}}

{{< /tabs >}}

Trie 树相关实现：
```go
type Trie struct {
    index int
    links map[byte]*Trie
}

func (t *Trie) insert(s string, index int) {
    node := t
    for i := range s {
        if node.links[s[i]] == nil {
            node.links[s[i]] = &Trie{links: map[byte]*Trie{}}
        }
        node = node.links[s[i]]
    }
    node.index = index
} 

func (t *Trie) find(s string) int {
    node := t
    for i := range s {
        if node.links[s[i]] == nil {
            return -1
        }
        node = node.links[s[i]]
    }
    return node.index
}
```