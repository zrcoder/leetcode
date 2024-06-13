---
title: "1438. 绝对差不超过限制的最长连续子数组"
date: 2022-03-31T15:50:35+08:00
---

## [1438. 绝对差不超过限制的最长连续子数组](https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/ "https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (48.48%) | 235   | -        |

给你一个整数数组 `nums` ，和一个表示限制的整数 `limit`，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 `limit` *。*

如果不存在满足条件的子数组，则返回 `0` 。

**示例 1：**

```
输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。
```

**示例 2：**

```
输入：nums = [10,1,2,4,7,2], limit = 5
输出：4
解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
```

**示例 3：**

```
输入：nums = [4,2,2,2,4,4,2,2], limit = 0
输出：3
```

**提示：**

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= limit <= 10^9`

函数签名：

```go
func longestSubarray(nums []int, limit int) int
```

## 分析

整体思路是滑动窗口：可以枚举窗口的右端点，尽量找到最靠左的左端点，窗口内部最大值与最小值之差不超过limit；如果没超过，右端点右移，否则，左端点右移。

如果直接使用数组模拟窗口，像窗口添加元素是常数级复杂度，删除元素和查找最大/最小值需要遍历，复杂度不理想。有没有合适的数据结构来模拟窗口呢？平衡的二叉搜索树就是一种满足需求的数据结构，增删元素、查找最大值、最小值的复杂度都是对数级。

像C++里的multiset，Java里的TreeMap，Go语言标准库并没有现成的数据结构，要手写一个红黑树，太麻烦，倒是可以尝试写个Treap。

以下代码示意这个解法，其中的 BST 结构代表平衡二叉搜索树，但不给出实现。

```go
func longestSubarray(nums []int, limit int) int {
    bst := &BST{}
    res := 0
    left := 0
    for right, num := range nums {
        bst.Insert(num)
        for bst.Max()-bst.Min() > limit {
            bst.Delete(nums[left])
            left++
        }
        res = max(res, right-left+1)
    }
    return res
}
```

时间复杂度：O(nlogn), 空间复杂度：O(n)。

复杂度更低的做法：

可以借助两个单调队列来维护窗口中的内容，其添加元素、删除元素、获取最值的复杂度都是常数级。

```go
func longestSubarray(nums []int, limit int) int {
    var minQ, maxQ []int
    res := 0
    left := 0
    for right, num := range nums {
        for len(minQ) > 0 && minQ[len(minQ)-1] > num {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, num)
        for len(maxQ) > 0 && maxQ[len(maxQ)-1] < num {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, num)
        for len(maxQ)> 0 && len(minQ) > 0 && maxQ[0]-minQ[0] > limit {
            if nums[left] == maxQ[0] {
                maxQ = maxQ[1:]
            }
            if nums[left] == minQ[0] {
                minQ = minQ[1:]
            }
            left++
        }
        res = max(res, right-left+1)
    }
    return res
}
```

时间复杂度降到了O(n)。
