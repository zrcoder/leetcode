---
title: "全排列"
date: 2021-04-19T22:04:56+08:00
weight: 1
tags: ["回溯"]
---

## [46. 全排列](https://leetcode-cn.com/problems/permutations)

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]

输出:

```
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## 分析

### 分治

自然的思路：如果已经得到前n-1个元素的全排列，再加一个元素的全排列也就不难得到：只需要遍历前n-1个元素全排列，对每个排列，在每个空隙插入新元素。

```go
func permute(nums []int) [][]int {
    res := [][]int{nil}
    for _, v := range nums {
        var tmp [][]int
        for _, s := range res {
            for i := 0; i <= len(s); i++ {
                // 在s的空位i处插入v
                ss := append(append(s[:i:i], v), s[i:]...) 
                tmp = append(tmp, ss)
            }
        }
        res = tmp
    }
    return res
}
```

也可以写成递归：

```go
func permute(nums []int) [][]int {
    if len(nums) < 2 {
        return [][]int{nums}
    }
    var res [][]int
    v := nums[len(nums)-1]
    for _, s := range permute(nums[:len(nums)-1]) {
        for i := 0; i <= len(s); i++ {
            ss := append(append(s[:i:i], v), s[i:]...)
            res = append(res, ss)
        }
    }
    return res
}
```

### DFS: 填空

将这个问题看作有 n个排列成一行的空格，需要从左往右依次填入题目给定的 n个数，每个数只能使用一次。

那么很直接的可以想到一种穷举的算法，即从左往右每一个位置都依此尝试填入一个数，看能不能填完这 n 个空格，在程序中我们可以用「回溯法」来模拟这个过程。

```go
func permute(nums []int) [][]int {
    var res [][]int
    var path []int
    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            res = append(res, append([]int{}, path...))
            return
        }
        for _, v := range  nums {
            if contains(path, v) {
                continue
            }
            path = append(path, v)
            backtrack()
            path = path[:len(path)-1]
        }
    }
    backtrack()
    return res
}

func contains(s []int, x int) bool {
    for _, v := range s {
        if v == x {
            return true
        }
    }
    return false
}
```

递归函数里边用了一个contains方法，非常低效，可以通过备忘来优化：

```go
func permute(nums []int) [][]int {
    var res [][]int
    var path []int
    // 备忘录，记录某次回溯过程中元素是否被访问
    seen := make([]bool, len(nums)) 
    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            res = append(res, append([]int{}, path...))
            return
        }
        for i, v := range  nums {
            if seen[i] {
                continue
            }
            seen[i] = true
            path = append(path, v)
            backtrack()
            seen[i] = false
            path = path[:len(path)-1]
        }
    }
    backtrack()
    return res
}
```

### DFS: 指定递归起始位置

深度优先搜索，先固定前边几个元素，然后开始尝试排列后边的，这样能逐步降低问题规模。

排列可以通过交换元素实现，参见dfs函数:

```go
func permute(nums []int) [][]int {
    n := len(nums)
    var result [][]int
    // 保持start之前的元素固定不变，将其及其之后的元素全排列
    var dfs func(start int)
    dfs = func(start int) {
        if start == n {
            result = append(result, append([]int{}, nums...))
            return
        }
        for i := start; i < n; i++ { // 将i及其i之后的元素全排列，注意不能漏了i
            nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
            dfs(start + 1)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }
    dfs(0)
    return result
}
```

## [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii)

```
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

## 分析

问题与46相似，只是加了元素可能重复的情况，结果不能有重复；

如果用46的第一种解法，需要对结果去重。而46的后两种解法可以事先去重。

### DFS: 填空

用上一问题的填空法，可以事先对nums排序，递归过程中去重。

```go
func permuteUnique(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums) // 事先排序
    n := len(nums)
    cur := []int{}
    seen := make([]bool, n)
    var dfs func()
    dfs = func() {
        if len(cur) == n {
            res = append(res, append([]int{}, cur...))
            return
        }
        for i, v := range nums {
            if seen[i] || i > 0 && v == nums[i-1] && !seen[i-1] { // 注意这里的 !seen[i-1]
                continue
            }
            seen[i] = true
            cur = append(cur, v)
            dfs()
            seen[i] = false
            cur = cur[:len(cur)-1]
        }
    }
    dfs()
    return res
}
```

### DFS: 指定递归起始位置

递归时用set去重, 具体在交换 start 处元素与后边元素的时候，看看是否已有相同的元素参与过交换，已经参与过的跳过。

```go
func permuteUnique(nums []int) [][]int {
    n := len(nums)
    var res [][]int
    // 保持start之前的元素固定不变，将其及其之后的元素全排列
    var dfs func(int)
    dfs = func(start int) {
        if start == n {
            res = append(res, append([]int{}, nums...))
            return
        }
        visited := make(map[int]bool, n-start)
        for i := start; i < n; i++ { // 将start及其之后的元素全排列，注意不能漏了start
            if visited[nums[i]] {
                continue
            }
            visited[nums[i]] = true
            nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
            dfs(start + 1)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }
    dfs(0)
    return res
}
```