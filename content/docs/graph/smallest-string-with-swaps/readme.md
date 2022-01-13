---
title: "1202. 交换字符串中的元素"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1202. 交换字符串中的元素](https://leetcode-cn.com/problems/smallest-string-with-swaps/)

难度中等

给你一个字符串 `s`，以及该字符串中的一些「索引对」数组 `pairs`，其中 `pairs[i] = [a, b]` 表示字符串中的两个索引（编号从 0 开始）。

你可以 **任意多次交换** 在 `pairs` 中任意一对索引处的字符。

返回在经过若干次交换后，`s` 可以变成的按字典序最小的字符串。

 

**示例 1:**

```
输入：s = "dcab", pairs = [[0,3],[1,2]]
输出："bacd"
解释： 
交换 s[0] 和 s[3], s = "bcad"
交换 s[1] 和 s[2], s = "bacd"
```

**示例 2：**

```
输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
输出："abcd"
解释：
交换 s[0] 和 s[3], s = "bcad"
交换 s[0] 和 s[2], s = "acbd"
交换 s[1] 和 s[2], s = "abcd"
```

**示例 3：**

```
输入：s = "cba", pairs = [[0,1],[1,2]]
输出："abc"
解释：
交换 s[0] 和 s[1], s = "bca"
交换 s[1] 和 s[2], s = "bac"
交换 s[0] 和 s[1], s = "abc"
```

 

**提示：**

- `1 <= s.length <= 10^5`
- `0 <= pairs.length <= 10^5`
- `0 <= pairs[i][0], pairs[i][1] < s.length`
- `s` 中只含有小写英文字母

函数签名：

```go
func smallestStringWithSwaps(s string, pairs [][]int) string
```

## 分析

根据 pairs 可以得到一系列联通的顶点，假设 `i1, i2, ..., ik` 这些位置相邻的点是联通的，可以发现这些点是可以通过交换任意排序的。

比如，要让 `i1` 和 `ik` 交换，而中间的点不变，只需要先将 `i1` 一直换到 最后，类似冒泡的方式，再类似地把现在居于倒数第二位的 `ik` 冒泡交换到最开始。

任意指定两个点，都可以在其他点不动的情况下交换这两个点。

因为这样的原因，要使得最终结果字典序最小，只需要把每个联通分量里的点按照对应字符字典序排列。

用并查集来事先处理联通分量的情况，之后将每个连通分量里的点排序，最后构建结果即可。

```go
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := makeUnionFind(n)
	for _, v := range pairs {
		uf.Union(v[0], v[1])
	}
	m := make(map[int][]byte, n)
	for i := range uf {
		root := uf.Find(i)
		m[root] = append(m[root], s[i])
	}
	for _, b := range m {
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
	}
	res := make([]byte, n)
	for i := range res {
		root := uf.Find(i)
		res[i] = m[root][0]
		m[root] = m[root][1:]
	}
	return string(res)
}

type UnionFind []int

func makeUnionFind(n int) UnionFind {
	uf := make([]int, n)
	for i := range uf {
		uf[i] = i
	}
	return uf
}

func (uf UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY
}

func (uf UnionFind) Find(x int) int {
	for x != uf[x] {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
```

时间复杂度： O(nlog n + m )，其中 n 为字符串长度，m 为索引对数量。并查集的操作可以看作常数级。题解中 Find 用了路径压缩，也可以在 Union 的时候按秩合并。最坏情况下所有点联通，需要把整个字符串排序。

空间复杂度： O(n)。