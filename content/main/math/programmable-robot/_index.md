---
title: "LCP 03. 机器人大冒险"
date: 2022-03-05T17:38:32+08:00
---

## [LCP 03. 机器人大冒险](https://leetcode-cn.com/problems/programmable-robot/)

难度中等

力扣团队买了一个可编程机器人，机器人初始位置在原点`(0, 0)`。小伙伴事先给机器人输入一串指令`command`，机器人就会**无限循环**这条指令的步骤进行移动。指令有两种：

1. `U`: 向`y`轴正方向移动一格
2. `R`: 向`x`轴正方向移动一格。

不幸的是，在 xy 平面上还有一些障碍物，他们的坐标用`obstacles`表示。机器人一旦碰到障碍物就会被**损毁**。

给定终点坐标`(x, y)`，返回机器人能否**完好**地到达终点。如果能，返回`true`；否则返回`false`。

 

**示例 1：**

```
输入：command = "URR", obstacles = [], x = 3, y = 2
输出：true
解释：U(0, 1) -> R(1, 1) -> R(2, 1) -> U(2, 2) -> R(3, 2)。
```

**示例 2：**

```
输入：command = "URR", obstacles = [[2, 2]], x = 3, y = 2
输出：false
解释：机器人在到达终点前会碰到(2, 2)的障碍物。
```

**示例 3：**

```
输入：command = "URR", obstacles = [[4, 2]], x = 3, y = 2
输出：true
解释：到达终点后，再碰到障碍物也不影响返回结果。
```

 

**限制：**

1. `2 <= command的长度 <= 1000`
2. `command`由`U，R`构成，且至少有一个`U`，至少有一个`R`
3. `0 <= x <= 1e9, 0 <= y <= 1e9`
4. `0 <= obstacles的长度 <= 1000`
5. `obstacles[i]`不为原点或者终点

函数签名：

```go
func robot(command string, obstacles [][]int, x int, y int) bool
```

## 分析

### 模拟

最容易想到的解法就是模拟。

机器人每次走一步，向上或向下，每一步判断：

- 是不是遇到了障碍（可以借助哈希表），是则返回false

- 是不是到达了终点，是则返回 true
- 是不是已经超过了x或y的值， 是则返回 false

参考解答如下：

```go
type P struct {
  x, y int
}

func robot(command string, obstacles [][]int, x int, y int) bool {
  memo := make(map[P]bool, len(obstacles))
  for _, v := range obstacles {
    memo[P{v[0], v[1]}] = true
  }
  cx, cy := 0, 0
  for i := 0; ; i = (i+1)%len(command) {
    if command[i] == 'U' {
      cy++
    } else {
      cx++
    }
    if cx > x || cy > y || memo[P{cx, cy}] {
      return false
    }
    if cx == x && cy == y {
      return true
    }
  }
}
```

时间复杂度：O(x+y)，空间复杂度: O(m)，m是 obstacles的大小，超时了。

### 数学

对于一个给定的 command 和一个点 (x, y)， 其实是可以在常数时间复杂度内确定这个点是否会落在 command 不断循环形成的路径上的。具体做法见代码。

```go
func robot(command string, obstacles [][]int, x int, y int) bool {
  if !isOnPath(command, x, y) {
    return false
  }
  for _, v := range obstacles {
    if v[0] <= x && v[1] <= y && isOnPath(command, v[0], v[1]) {
      return false
    }
  }
  return true
}

func isOnPath(c string, x, y int) bool {
  steps := x+y // 要从原点走到(x, y)点，恰好走 x+y 步
  // 计算用 c 来走 steps 步能贡献多少个 R 和 U
  rs := strings.Count(c, "R")*(steps/len(c)) + strings.Count(c[:steps%len(c)], "R")
  us := strings.Count(c, "U")*(steps/len(c)) + strings.Count(c[:steps%len(c)], "U")
  return rs == x && us == y
}
```

时间复杂度：O(m)，m 是 obstacles的大小，空间复杂度常数级。

对比两种解法的时间复杂度，结合题目的限制，会发现后者远优于前者。
