---
title: "210. 课程表 II"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [210. 课程表 II](https://leetcode-cn.com/problems/course-schedule-ii)
现在你总共有 n 门课需要选，记为 0 到 n-1。  
在选修某些课程之前需要一些先修课程。  
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]  
给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。  
可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
```
示例 1:
输入: 2, [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

示例 2:
输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。
     并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

说明:
输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。

提示:
这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，
则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。]	
```

## 分析
经典拓扑排序。

可以用 BFS 或 DFS 两种方法解决。

### BFS
可以先用一个数组记录每门课程还未修的前置课程数目，可称为各个节点的入度数组，记为 `degree`。

之后逐一把入度为 0 的节点加入结果，并及时更新其相邻节点的入度，这样可能有新的节点的入度变成 0。

循环直到所有入度为 0 的节点都加入了结果。

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	degree := make([]int, numCourses)
	nexts := make([][]int, numCourses)
	for _, req := range prerequisites {
		degree[req[0]]++
		nexts[req[1]] = append(nexts[req[1]], req[0])
	}
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue.PushBack(i)
		}
	}
	res := make([]int, 0, numCourses)
	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		res = append(res, course)
		for _, next := range nexts[course] {
			degree[next]--
			if degree[next] == 0 {
				queue.PushBack(next)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return nil
}
```

时间复杂度 `O(n+m)`，n 和 m 分别为节点数量和边的数量。

空间复杂度 `O(n)`。

### DFS

也可以用 dfs 来解决拓扑排序问题。dfs 过程中，除了要判断一个节点是否已经访问过，还需要判断当次 dfs 是否形成了环，
这可以给每个节点增加两个状态来标记：

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	dependency := make([][]int, numCourses)
	for _, req := range prerequisites {
		dependency[req[0]] = append(dependency[req[0]], req[1])
	}
	flags := make([]int, numCourses)
	var result []int
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// 1.节点已经被访问过
		if flags[course] != 0 { 
			return flags[course] == 1 // 1 代表之前 dfs 访问该点无环
		}
		// 2.节点还没有被访问过，先假设有环
		flags[course] = 2
		for _, neighbor := range dependency[course] {
			if !dfs(neighbor) { // 真的有环
				return false
			}
		}
		// 所有相邻节点都 dfs 搜索过了，可以确定当次 dfs 无环，加入结果
		flags[course] = 1
		result = append(result, course)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return nil
		}
	}
	return result
}
```

复杂度同 BFS。