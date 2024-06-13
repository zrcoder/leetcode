---
title: "1760. 袋子里最少数目的球"
date: 2022-12-20T10:51:11+08:00
---

##  [1760. 袋子里最少数目的球](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/description)

| Category | Difficulty | Likes | Dislikes |
| --- | --- | --- | --- |
| algorithms | Medium (58.10%) | 121 | -   |

给你一个整数数组 `nums` ，其中 `nums[i]` 表示第 `i` 个袋子里球的数目。同时给你一个整数 `maxOperations` 。

你可以进行如下操作至多 `maxOperations` 次：

- 选择任意一个袋子，并将袋子里的球分到 2 个新的袋子中，每个袋子里都有 **正整数** 个球。
    - 比方说，一个袋子里有 `5` 个球，你可以把它们分到两个新袋子里，分别有 `1` 个和 `4` 个球，或者分别有 `2` 个和 `3` 个球。

你的开销是单个袋子里球数目的 **最大值** ，你想要 **最小化** 开销。

请你返回进行上述操作后的最小开销。

**示例 1：**

```
输入：nums = [9], maxOperations = 2
输出：3
解释：
- 将装有 9 个球的袋子分成装有 6 个和 3 个球的袋子。[9] -> [6,3] 。
- 将装有 6 个球的袋子分成装有 3 个和 3 个球的袋子。[6,3] -> [3,3,3] 。
装有最多球的袋子里装有 3 个球，所以开销为 3 并返回 3 。
```

**示例 2：**

```
输入：nums = [2,4,8,2], maxOperations = 4
输出：2
解释：
- 将装有 8 个球的袋子分成装有 4 个和 4 个球的袋子。[2,4,8,2] -> [2,4,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,4,4,4,2] -> [2,2,2,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,4,4,2] -> [2,2,2,2,2,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2] 。
装有最多球的袋子里装有 2 个球，所以开销为 2 并返回 2 。
```

**示例 3：**

```
输入：nums = [7,17], maxOperations = 2
输出：7
```

**提示：**

- `1 <= nums.length <= 10^5`
- `1 <= maxOperations, nums[i] <= 10^9`

函数签名：

```go
func minimumSize(nums []int, maxOperations int) int
```

## 分析

### 二分搜索

需要转化成判定问题：对于一个给定的成本 x，遍历数组，拆分 > x 的数字，统计总共拆分的次数是否 <= maxOperations 即可。

朴素实现如下：

```go
func minimumSize(nums []int, maxOperations int) int {
    check := func(x int) bool {
        opers := 0
        for _, v := range nums {
            opers += (v-1)/x
        }
        return opers <= maxOperations
    }

    for x := 1; x <= 1e9; x++ { // 1e9 actually should be max(nums)
        if check(x) {
            return x
        }
    }
    return -1
}
```

时间复杂度：`O(n*max)`，空间复杂度`O(1)`,其中 n 指数组长度，max 指数组中的最大元素。

实际上，check 具有单调性，显然 x 较小时一直返回 false，在某个临界点后一直返回 true，实际就是求该临界点，朴素实现改为二分搜索：

```go
func minimumSize(nums []int, maxOperations int) int {
    check := func(x int) bool {
        opers := 0
        for _, v := range nums {
            opers += (v-1)/x
        }
        return opers <= maxOperations
    }

    lo, hi := 1, int(1e9)+1
    for lo < hi {
        mid := (lo+hi)/2
        if check(mid) {
            hi = mid
        } else {
            lo = mid+1
        }
    }
    return lo
}
```

改用标准库，简化代码：

```go
func minimumSize(nums []int, maxOperations int) int {
    return 1 + sort.Search(1e9, func(x int) bool {
        opers := 0
        for _, v := range nums {
            opers += (v-1) / (x+1)
        }
        return opers <= maxOperations
    })
}
```

时间复杂度降为：`O(n*logmax)`，空间复杂度不变，仍然为常数级。
