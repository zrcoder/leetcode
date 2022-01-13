---
title: "重排使相同元素至少间隔k"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [621. 任务调度器](https://leetcode-cn.com/problems/task-scheduler)
给定一个用字符数组表示的 CPU 需要执行的任务列表。  
其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。  
任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。  
CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。  
然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，  
因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。  
你需要计算完成所有任务所需要的最短时间。 
```
示例 ：
输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，
而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。  

提示：
任务的总个数为 [1, 10000]。
n 的取值范围为 [0, 100]。
```
## 分析
有一个明显的贪心策略，即优先安排数量大的任务。下边解法都是基于这个贪心策略。
## 模拟
先统计每个任务的数量。  
以n+1个任务为一轮，同一轮中一个任务最多被安排一次。  
每一轮中，将当前任务按照剩余次数降序排列，再选择剩余次数最多的n+1个任务一次执行  
如果任务的种类 t 少于 n + 1 个，就只选择全部的 t 种任务，其余的时间空闲。 
```go
func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(count)))
	result := 0
	for count[0] > 0 {
		for i := 0; i <= n && count[0] > 0; i++ {
			result++
			if i < 26 && count[i] > 0 {
				count[i]--
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(count)))
	}
	return result
}
```
时间复杂度: O(result),给每个任务都安排了时间，因此时间复杂度和最终的答案成正比  
空间复杂度: O(26)=O(1)

也可以用堆代替排序来做模拟：
```go
func leastInterval(tasks []byte, n int) int {
	h := prepareHeap(tasks)
	result := 0
	set := list.New()
	for h.Len() > 0 {
		for i := 0; i <= n; i++ {
			if h.Len() == 0 && set.Len() == 0 {
				return result
			}
			result++
			if h.Len() == 0 { // 需要待命只到i==n
				continue
			}
			t := heap.Pop(h).(int)
			if t > 1 {
				set.PushBack(t - 1)
			}
		}
		for set.Len() > 0 {
			heap.Push(h, set.Remove(set.Front()).(int))
		}
	}
	return result
}

func prepareHeap(tasks []byte) *Heap {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	h := &Heap{}
	for _, v := range count {
		if v > 0 {
			h.Push(v)
		}
	}
	heap.Init(h)
	return h
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
```
复杂度同上。

## 总结规律用公式计算
1.统计数量最大的任务的数量max；假设数组 ["A","A","A","B","B","C"]，n = 2，A的频率最高，记为max = 3，
两个A之间必须间隔2个任务，才能满足题意并且是最短时间（两个A的间隔大于2的总时间必然不是最短），
因此执行顺序为： A->X->X->A->X->X->A，这里的X表示除了A以外其他字母，或者是待命，不用关心具体是什么。
max个A，中间有 max - 1个间隔，每个间隔需要搭配n个X，再加上最后一个A，所以总时间为 (max - 1) * (n + 1) + 1

2.要注意可能会出现多个频率相同且都是最高的任务，比如 ["A","A","A","B","B","B","C","C"]，所以最后会剩下一个A和一个B，
因此最后要加上频率最高的不同任务的个数 maxCount

3.公式算出的值可能会比数组的长度小，如["A","A","B","B","D"]，n=1，遗漏了最后的 D（还可以有 E、F等），这种情况要取数组的长度
```go
func leastInterval3(tasks []byte, n int) int {
	// 统计每个任务的数量
	count := [26]int{}
	// 数量最大的任务的数量及个数
	max, maxCount := 0, 0
	for _, v := range tasks {
		c := count[v-'A'] + 1
		count[v-'A'] = c
		if max < c {
			max = c
			maxCount = 1
		} else if max == c {
			maxCount++
		}
	}
	result := (max-1)*(n+1) + maxCount
	if result < len(tasks) {
		result = len(tasks)
	}
	return result
}
```
时间复杂度：O(M)，其中 M 是任务的总数，即 tasks 数组的长度。

空间复杂度：O(1)。
## [358. K 距离间隔重排字符串](https://leetcode-cn.com/problems/rearrange-string-k-distance-apart/)
给你一个非空的字符串 s 和一个整数 k，你要将这个字符串中的字母进行重新排列，  
使得重排后的字符串中相同字母的位置间隔距离至少为 k。  
所有输入的字符串都由小写字母组成，如果找不到距离至少为 k 的重排结果，请返回一个空字符串 ""。  
```
示例 1：
输入: s = "aabbcc", k = 3
输出: "abcabc" 
解释: 相同的字母在新的字符串中间隔至少 3 个单位距离。

示例 2:
输入: s = "aaabc", k = 3
输出: "" 
解释: 没有办法找到可能的重排结果。

示例 3:
输入: s = "aaadbbcc", k = 2
输出: "abacabcd"
解释: 相同的字母在新的字符串中间隔至少 2 个单位距离。
```
## 分析
同问题621一般思路。
```go
func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}
	result := []byte(s)
	pairs := count(result)
	cmp := func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].char < pairs[j].char
		}
		return pairs[i].count > pairs[j].count
	}
	sort.Slice(pairs, cmp)
	j := 0
	for pairs[0].count > 0 {
		for i := 0; i < k; i++ {
			if pairs[0].count == 0 {
				break
			}
			if i >= len(pairs) || pairs[i].count == 0 {
				return ""
			}
			result[j] = pairs[i].char
			j++
			pairs[i].count--
		}
		sort.Slice(pairs, cmp)
	}
	return string(result)
}

type Pair struct {
	count int
	char  byte
}

func count(s []byte) []Pair {
	pairs := make([]Pair, 26)
	for _, b := range s {
		pairs[b-'a'].char = b
		pairs[b-'a'].count++
	}
	return pairs
}
``` 
## [767. 重构字符串](https://leetcode-cn.com/problems/reorganize-string)
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。  
若可行，输出任意可行的结果。若不可行，返回空字符串。  
```
示例 1:
输入: S = "aab"
输出: "aba"

示例 2:
输入: S = "aaab"
输出: ""

注意:
S 只包含小写字母并且长度在[1, 500]区间内。
```
## 分析
问题358变体，k==2的特例：
```go
func reorganizeString0(s string) string {
	return rearrangeString(s, 2)
}
```

