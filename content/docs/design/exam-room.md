---
title: 855. 考场就座
date: 2024-12-24T22:41:38+08:00
---

## [855. 考场就座](https://leetcode.cn/problems/exam-room) (Medium)

在考场里，有 `n` 个座位排成一行，编号为 `0` 到 `n - 1`。

当学生进入考场后，他必须坐在离最近的人最远的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 `0` 号座位上。)

设计一个模拟所述考场的类。

实现 `ExamRoom` 类：

- `ExamRoom(int n)` 用座位的数量 `n` 初始化考场对象。
- `int seat()` 返回下一个学生将会入座的座位编号。
- `void leave(int p)` 指定坐在座位 `p` 的学生将离开教室。保证座位 `p` 上会有一位学生。

**示例 1：**

```
输入：
["ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"]
[[10], [], [], [], [], [4], []]
输出：
[null, 0, 9, 4, 2, null, 5]
解释：
ExamRoom examRoom = new ExamRoom(10);
examRoom.seat(); // 返回 0，房间里没有人，学生坐在 0 号座位。
examRoom.seat(); // 返回 9，学生最后坐在 9 号座位。
examRoom.seat(); // 返回 4，学生最后坐在 4 号座位。
examRoom.seat(); // 返回 2，学生最后坐在 2 号座位。
examRoom.leave(4);
examRoom.seat(); // 返回 5，学生最后坐在 5 号座位。

```

**提示：**

1. `1 <= n <= 10⁹`
2. 保证有学生正坐在座位 `p` 上。
3. `seat` 和 `leave` 最多被调用 `10⁴` 次。

## 分析

用一个有序集合维护所有的连续空位段，比如 （x, y） 表示 x 和 y 之间的空位， 注意为开区间。

```go
type ExamRoom struct {
	Set    *redblacktree.Tree
	Lo, Hi map[int]int
	N      int
}

func Constructor(n int) ExamRoom {
	dist := func(seg []int) int {
		lo, hi := seg[0], seg[1]
		if lo == -1 || hi == n {
			return hi - lo - 1
		}
		return (hi - lo) / 2
	}
	er := ExamRoom{
		Set: redblacktree.NewWith(func(a, b any) int {
			x, y := a.([]int), b.([]int)
			return cmp.Or(dist(y)-dist(x), x[0]-y[0])
		}),
		Lo: map[int]int{},
		Hi: map[int]int{},
		N:  n,
	}
	er.add([]int{-1, n})
	return er
}

func (er *ExamRoom) Seat() int {
	s := er.Set.Left().Key.([]int)
	p := (s[0] + s[1]) / 2
	if s[0] == -1 {
		p = 0
	} else if s[1] == er.N {
		p = er.N - 1
	}
	er.remove(s)
	er.add([]int{s[0], p})
	er.add([]int{p, s[1]})
	return p
}

func (er *ExamRoom) Leave(p int) {
	lo := er.Lo[p]
	hi := er.Hi[p]
	er.remove([]int{lo, p})
	er.remove([]int{p, hi})
	er.add([]int{lo, hi})
}

func (er *ExamRoom) add(s []int) {
	er.Set.Put(s, struct{}{})
	er.Lo[s[1]] = s[0]
	er.Hi[s[0]] = s[1]
}

func (er *ExamRoom) remove(s []int) {
	er.Set.Remove(s)
	delete(er.Lo, s[1])
	delete(er.Hi, s[0])
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
```
