---
title: "1268. 搜索推荐系统"
date: 2021-05-31T19:07:26+08:00
weight: 50
tags: [Trie树,二分搜索]
---

## [1268. 搜索推荐系统](https://leetcode-cn.com/problems/search-suggestions-system/)

难度中等

给你一个产品数组 `products` 和一个字符串 `searchWord` ，`products` 数组中每个产品都是一个字符串。

请你设计一个推荐系统，在依次输入单词 `searchWord` 的每一个字母后，推荐 `products` 数组中前缀与 `searchWord` 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。

请你以二维列表的形式，返回在输入 `searchWord` 每个字母后相应的推荐产品的列表。

**示例 1：**

```
输入：products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
输出：[
["mobile","moneypot","monitor"],
["mobile","moneypot","monitor"],
["mouse","mousepad"],
["mouse","mousepad"],
["mouse","mousepad"]
]
解释：按字典序排序后的产品列表是 ["mobile","moneypot","monitor","mouse","mousepad"]
输入 m 和 mo，由于所有产品的前缀都相同，所以系统返回字典序最小的三个产品 ["mobile","moneypot","monitor"]
输入 mou， mous 和 mouse 后系统都返回 ["mouse","mousepad"]
```

**示例 2：**

```
输入：products = ["havana"], searchWord = "havana"
输出：[["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
```

**示例 3：**

```
输入：products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
输出：[["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
```

**示例 4：**

```
输入：products = ["havana"], searchWord = "tatiana"
输出：[[],[],[],[],[],[],[]]
```

**提示：**

- `1 <= products.length <= 1000`
- `1 <= Σ products[i].length <= 2 * 10^4`
- `products[i]` 中所有的字符都是小写英文字母。
- `1 <= searchWord.length <= 1000`
- `searchWord` 中所有字符都是小写英文字母。

函数签名：

```go
func suggestedProducts(products []string, searchWord string) [][]string
```

## 分析

### Trie 树
Trie 树存储所有单词，遍历 searchWord  前缀做查询。

> 为了在树上迅速获得单词，可以给节点加一个 word 属性，插入单词的时候，最后一个节点的 word 属性即为该单词。

```go
const (
	letters = 26
	
	resultLimit = 3
)

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &Trie{links: make([]*Trie, letters)}
	for _, product := range products {
		trie.insert(product)
	}
	res := make([][]string, len(searchWord))
	cur := trie
	for i := range searchWord {
		index := searchWord[i] - 'a'
		if cur.links[index] == nil {
			break
		}
		cur = cur.links[index]
		res[i] = cur.getAllWords(resultLimit)
	}
	return res
}

type Trie struct {
	links []*Trie
	word  string
}

func (t *Trie) insert(word string) {
	cur := t
	for _, char := range word {
		index := char - 'a'
		if cur.links[index] == nil {
			cur.links[index] = &Trie{links: make([]*Trie, letters)}
		}
		cur = cur.links[index]
	}
	cur.word = word
}

func (t *Trie) getAllWords(limit int) []string {
	var res []string
	var dfs func(cur *Trie)
	dfs = func(cur *Trie) {
		if cur.word != "" {
			res = append(res, cur.word)
		}
		if len(res) == limit {
			return
		}
		for _, next := range cur.links {
			if len(res) == limit {
				return
			}
			if next == nil {
				continue
			}
			dfs(next)
		}
	}
	dfs(t)
	return res
}
```

时间复杂度：`O(C+S)`，其中 `C` 指所有单词长度和，`S` 指 `searchWord` 的长度。

空间复杂度：`O(C)`。

### 二分搜索
可以先按照字典序将所有产品排序，对于 searchWord的每个前缀，可以用二分搜索找到第一个不小于该前缀的产品，然后最多检查三个产品得到这个前缀对应的结果。
```go
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		prefix := searchWord[:i+1]
		index := sort.Search(len(products), func(i int) bool {
			return products[i] >= prefix
		})
		if index == len(products) {
			break
		}
		end := min(index+3, len(products))
		for k := index; k < end && strings.HasPrefix(products[k], prefix); k++ {
			res[i] = append(res[i], products[k])
		}
		products = products[index:] // 前边的部分确定可以排除，后续不再需要
	}
	return res
}
```
实测二分搜索要比 Trie 树的解法更快且更省内存~