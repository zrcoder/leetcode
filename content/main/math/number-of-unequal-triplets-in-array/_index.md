---
title: "2475. 数组中不等三元组的数目"
date: 2022-11-23T16:02:19+08:00
---

## [2475. 数组中不等三元组的数目](https://leetcode.cn/problems/number-of-unequal-triplets-in-array/)

难度简单

给你一个下标从 **0** 开始的正整数数组 `nums` 。请你找出并统计满足下述条件的三元组 `(i, j, k)` 的数目：

- `0 <= i < j < k < nums.length`
- `nums[i]`、`nums[j]` 和 `nums[k]` **两两不同** 。
  - 换句话说：`nums[i] != nums[j]`、`nums[i] != nums[k]` 且 `nums[j] != nums[k]` 。

返回满足上述条件三元组的数目*。*

**示例 1：**

**输入：** nums = [4,4,2,4,3]
**输出：** 3
**解释：** 下面列出的三元组均满足题目条件：

- (0, 2, 4) 因为 4 != 2 != 3
- (1, 2, 4) 因为 4 != 2 != 3
- (2, 3, 4) 因为 2 != 4 != 3
  共计 3 个三元组，返回 3 。
  注意 (2, 0, 4) 不是有效的三元组，因为 2 > 0 。

**示例 2：**

**输入：** nums = [1,1,1,1,1]
**输出：** 0
**解释：** 不存在满足条件的三元组，所以返回 0 。

**提示：**

- `3 <= nums.length <= 100`
- `1 <= nums[i] <= 1000`

函数签名：

```go
func unequalTriplets(nums []int) int
```

## 分析

这个问题标为简单，实际很不简单。

### 朴素解

因为数据规模很小，可以直接三重循环枚举三个元素可能的情况。

代码略。

时间复杂度`O(n^3)`， 空间复杂度 `O(1)`。

### 容斥原理-排序

如果对数组排序，那么相同的数字会聚在一起`。

在排序的基础上考虑每个数字对结果的贡献。如对于数字 x，小于 x 的元素个数为 a， 等于 x 的元素个数为 b，大于 x 的元素个数为 c，可以从小于x、等于x、大于x的元素中分别拿出一个形成一个满足题意的三元组，所以对数字 x 来说，它对结果的贡献是`a*b*c`。

```go
func unequalTriplets(nums []int) int {
    sort.Ints(nums)

    n := len(nums)
    less := 0 // 记录比当前元素小的其他元素的个数
    res := 0

    for i := 0; i < n-1; i++ {
        if nums[i] != nums[i+1] {
            cur := i+1-less
            more := n-i-1
            res += less*cur*more
            less = i+1
        }
    }

    return res
}
```

时间复杂度`O(nlogn)`，主要花在排序上； 空间复杂度 `O(1)`。

### 容斥原理-哈希表

结果不受元素顺序影响，只跟每个元素的个数有关。

可以直接用哈希表统计每个元素出现的个数，再遍历这些个数来计算结果。

```go
func unequalTriplets(nums []int) int {
    cnt := map[int]int{}
    for _, v := range nums {
        cnt[v]++
    }

    less := 0
    more := len(nums)
    res := 0

    for _, cur := range cnt {
        more -= cur
        res += less*cur*more
        less += cur
    }

    return res
}
```

时空复杂度都是 `O(n)`。
