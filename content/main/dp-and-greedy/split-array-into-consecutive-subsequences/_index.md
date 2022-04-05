---
title: "659. 分割数组为连续子序列"
date: 2022-04-05T12:04:56+08:00
tags: ["贪心"]
---

## [659. 分割数组为连续子序列](https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/description/ "https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (54.26%) | 366   | -        |

给你一个按升序排序的整数数组 `num`（可能包含重复数字），请你将它们分割成一个或多个长度至少为 3 的子序列，其中每个子序列都由连续整数组成。

如果可以完成上述分割，则返回 `true` ；否则，返回 `false` 。

**示例 1：**

```
输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3
3, 4, 5
```

**示例 2：**

```
输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 : 
1, 2, 3, 4, 5
3, 4, 5
```

**示例 3：**

```
输入: [1,2,3,4,4,5]
输出: False
```

**提示：**

- `1 <= nums.length <= 10000`

函数签名：

```go
func isPossible(nums []int) bool
```

## 分析

题目简述为：有一手扑克牌，需要全部凑成顺子，每个顺子最少3张牌，能恰好用完所有牌吗？

直观地感觉先要排序，题目已经保证排好序了，这步略。

看两个个例子：

- `1, 2, 3, 4, 5, 6`
  
  可以凑成成 `1, 2, 3` 和 `4, 5, 6`; 也可以凑成`1, 2, 3, 4, 4, 6`。

- `1, 2, 3, 4, 5`
  
  只能凑成 `1, 2, 3, 4, 5`

扑克牌 4 究竟另起炉灶以自己为一个顺子的开头好，还是接续到已有顺子后边好？从上边两个例子看，后者好。

另起炉灶有无法构成顺子的可能，而接续已有能尽可能已有增加顺子的长度。因为顺子至少3张牌，所以接续总是优于另起炉灶。

那么对于有重复数字的情况呢？比如：`1, 2, 3, 4, 4, 5, 6`，显然，第一个 4 接续已有顺子 `1, 2, 3`就好，而第二个 4 需要另起炉灶炉灶，实际上它也没有可以接续的顺子可用。

基于以上分析得到一个贪心策略：遍历，对于当前数字，优先接续到已有顺子后边，其次考虑自成一派。

用一个哈希表 `cnt` 维护各个数字的个数；用另一个哈希表 `endCnt` 来维护已经构造出的以某个数字结尾的顺子的个数。

遍历到数字 `v` 时，先判断有没有已经构造的以 `v-1` 结尾的顺子

    有的话（`endCnt[v-1] > 0`），将 `v` 接续到其中一个后边，这只需要让 `cnt[v]--` ，同时` endCnt[v-1]--` 而 `endCnt[v]++`;

    否则，只能让 `v`另起炉灶作为一个顺子的头了。这要求 `v+1`和 `v+2` 需要有剩余，否则直接返回 `false`。

```go
func isPossible(nums []int) bool {
    cnt := make(map[int]int, len(nums))
    for _, v := range nums {
        cnt[v]++
    }
    endCnt := make(map[int]int, len(nums))
    for _, v := range nums {
        if cnt[v] == 0 { // v 已经用完了
            continue
        }
        if endCnt[v-1] > 0 { // 优先接续到其他顺子后边
            cnt[v]--
            endCnt[v-1]--
            endCnt[v]++
        } else if cnt[v+1] > 0 && cnt[v+2] > 0 { // 另起炉灶
            cnt[v]--
            cnt[v+1]--
            cnt[v+2]--
            endCnt[v+2]++
        } else { // 无法另起炉灶
            return false
        }
    }
    return true
}
```

时空复杂度都是`O(n)`。
