---
title: "841. 钥匙和房间"
date: 2022-12-20T12:04:10+08:00
---

##  [841. 钥匙和房间](https://leetcode.cn/problems/keys-and-rooms/description")

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (67.33%) | 280 | -   |

有 `n` 个房间，房间按从 `0` 到 `n - 1` 编号。最初，除 `0` 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。

当你进入一个房间，你可能会在里面找到一套不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。

给你一个数组 `rooms` 其中 `rooms[i]` 是你进入 `i` 号房间可以获得的钥匙集合。如果能进入 **所有** 房间返回 `true`，否则返回 `false`。

**示例 1：**

```
输入：rooms = [[1],[2],[3],[]]
输出：true
解释：
我们从 0 号房间开始，拿到钥匙 1。
之后我们去 1 号房间，拿到钥匙 2。
然后我们去 2 号房间，拿到钥匙 3。
最后我们去了 3 号房间。
由于我们能够进入每个房间，我们返回 true。
```

**示例 2：**

```
输入：rooms = [[1,3],[3,0,1],[2],[0]]
输出：false
解释：我们不能进入 2 号房间。
```

**提示：**

- `n == rooms.length`
- `2 <= n <= 1000`
- `0 <= rooms[i].length <= 1000`
- `1 <= sum(rooms[i].length) <= 3000`
- `0 <= rooms[i][j] < n`
- 所有 `rooms[i]` 的值 **互不相同**

函数签名：

```go
func canVisitAllRooms(rooms [][]int) bool
```

## 分析

BFS/DFS入门问题。

{{<tabs>}}
{{%tab name="DFS"%}}

```go
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(cur int) {
		if vis[cur] {
			return
		}

		vis[cur] = true
		n--
		for _, v := range rooms[cur] {
			dfs(v)
		}
	}
	dfs(0)
	return n == 0
}
```

{{%/tab%}}
{{%tab name="BFS"%}}

```go
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)
	queue := []int{0}
	vis[0] = true
	n--
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, v := range rooms[cur] {
			if !vis[v] {
				vis[v] = true
				n--
				queue = append(queue, v)
			}
		}
	}
	return n == 0
}
```

{{%/tab%}}
{{</tabs>}}

时空复杂度都是`O(n)`。需要遍历/存储所有的房间。
