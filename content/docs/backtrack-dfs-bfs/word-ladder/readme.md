---
title: "单词接龙"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [127. 单词接龙](https://leetcode-cn.com/problems/word-ladder)

难度中等

给定两个单词（*beginWord* 和 *endWord*）和一个字典，找到从 *beginWord* 到 *endWord* 的最短转换序列的长度。转换需遵循如下规则：

1. 每次转换只能改变一个字母。
2. 转换过程中的中间单词必须是字典中的单词。

**说明:**

- 如果不存在这样的转换序列，返回 0。
- 所有单词具有相同的长度。
- 所有单词只由小写字母组成。
- 字典中不存在重复的单词。
- 你可以假设 *beginWord* 和 *endWord* 是非空的，且二者不相同。

**示例 1:**

```
输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
```

**示例 2:**

```
输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。
```

## 分析

常规 BFS

```go
var ( 
    // 记录每个单词改变一个字母能得到的且存在于 wordList 的单词列表
    nexts map[string][]string
    // 记录从 beginWord 变换到当前单词需要的步数
    steps map[string]int
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
    hasBegin, hasEnd := false, false
    for _, v := range wordList {
        if v == beginWord {
            hasBegin = true
        } else if v == endWord {
            hasEnd = true
        }
    }
    if !hasEnd {
        return 0
    }
    if !hasBegin {
        wordList = append(wordList, beginWord)
    }
    
    initNexts(wordList)
    initSteps(wordList)
    steps[beginWord] = 0

    queue := list.New()
    queue.PushBack(beginWord)
    for queue.Len() > 0 {
        word := queue.Remove(queue.Front()).(string)
        if word == endWord {
            return steps[word]+1
        }
        for _, next := range nexts[word] {
            if steps[next] <= steps[word]+1 {
                continue
            }
            steps[next] = steps[word]+1
            queue.PushBack(next)
        }
    }
    return 0
}

func initNexts(wordList []string) {
    n := len(wordList)
    nexts = make(map[string][]string, n)
    for i := 0; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            s, t := wordList[i], wordList[j]
            if diffOneChar(s, t) {
                nexts[s] = append(nexts[s], t)
                nexts[t] = append(nexts[t], s)
            }
        }
    }
}

func diffOneChar(s, t string) bool {
    diffs := 0
    for i := 0; i < len(s); i++ {
        if s[i] != t[i] {
            diffs++
        }
    }
    return diffs == 1
}

func initSteps(wordList []string) {
    steps = make(map[string]int, len(wordList))
    for _, w := range wordList {
        steps[w] = len(wordList)
    }
}
```

## TODO
- 可以优化为双向 BFS
- steps 的构造可以优化

## 变体 [126. 单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii)

约束相同，只是这次要返回所有从 beginWord 到 endWord 的最短转换序列。

解法也同 127，稍作调整即可。
