---
title: "等差数列划分"
date: 2022-11-27T16:53:22+08:00
---

## [413. 等差数列划分](https://leetcode.cn/problems/arithmetic-slices/)

难度中等

如果一个数列 **至少有三个元素** ，并且任意两个相邻元素之差相同，则称该数列为等差数列。

- 例如，`[1,3,5,7,9]`、`[7,7,7,7]` 和 `[3,-1,-5,-9]` 都是等差数列。

给你一个整数数组 `nums` ，返回数组 `nums` 中所有为等差数组的 **子数组** 个数。

**子数组** 是数组中的一个连续序列。

**示例 1：**

**输入：** nums = [1,2,3,4]
**输出：** 3
**解释：** nums 中有三个子等差数组：[1, 2, 3]、[2, 3, 4] 和 [1,2,3,4] 自身。

**示例 2：**

**输入：** nums = [1]
**输出：** 0

**提示：**

- `1 <= nums.length <= 5000`
- `-1000 <= nums[i] <= 1000`

函数签名：

```go
func numberOfArithmeticSlices(nums []int) int
```

### 分析

### 朴素解法

两层循环枚举所有子数组的解法比较容易：

```go
func numberOfArithmeticSlices(nums []int) int {
    res := 0
    for i := 0; i < len(nums)-2; i++ {
        d := nums[i+1]-nums[i]
        for j := i+2; j < len(nums); j++ {
            if nums[j]-nums[j-1] != d {
                break
            }
            res++
        }
    }
    return res
}
```

时间复杂度：`O(n^2)`，空间复杂度：`O(1)`。

### 差分数组

如果事先计算出原数组的差分数组，可以在线性时间内解决问题。

```go
func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }

    diff := make([]int, n-1)
    for i := 1; i < n; i++ {
        diff[i-1] = nums[i]-nums[i-1]
    }

    res := 0
    d := diff[0]
    cur := 0
    for i := 1; i < len(diff); i++ {
        if diff[i] == d {
            cur++
        } else {
            cur = 0
            d = diff[i]
        }
        res += cur
    }
    return res
}
```

实际也可以不用事先计算出差分数组，只需在遍历 nums 时维护当前差值即可，空间复杂度降到常数级。

```go
func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }

    res := 0
    d := nums[1]-nums[0]
    cur := 0
    for i := 2; i < n; i++ {
        if nums[i]-nums[i-1] == d {
            cur++
        } else {
            cur = 0
            d = nums[i]-nums[i-1]
        }
        res += cur
    }
    return res
}
```

## [446. 等差数列划分 II - 子序列](https://leetcode.cn/problems/arithmetic-slices-ii-subsequence/)

难度困难

给你一个整数数组 `nums` ，返回 `nums` 中所有 **等差子序列** 的数目。

如果一个序列中 **至少有三个元素** ，并且任意两个相邻元素之差相同，则称该序列为等差序列。

- 例如，`[1, 3, 5, 7, 9]`、`[7, 7, 7, 7]` 和 `[3, -1, -5, -9]` 都是等差序列。
- 再例如，`[1, 1, 2, 5, 7]` 不是等差序列。

数组中的子序列是从数组中删除一些元素（也可能不删除）得到的一个序列。

- 例如，`[2,5,10]` 是 `[1,2,1,***2***,4,1,***5***,***10***]` 的一个子序列。

题目数据保证答案是一个 **32-bit** 整数。

**示例 1：**

**输入：** nums = [2,4,6,8,10]
**输出：** 7
**解释：** 所有的等差子序列为：
[2,4,6]
[4,6,8]
[6,8,10]
[2,4,6,8]
[4,6,8,10]
[2,4,6,8,10]
[2,6,10]

**示例 2：**

**输入：** nums = [7,7,7,7,7]
**输出：** 16
**解释：** 数组中的任意子序列都是等差子序列。

**提示：**

- `1  <= nums.length <= 1000`
- `-2^31 <= nums[i] <= 2^31 - 1`

 函数签名：

```go
func numberOfArithmeticSlices(nums []int) int
```

## 分析

### 回溯（超时）

可以用回溯的方式枚举出所有子序列，并判断每个子序列是否为等差数列。

```go
func numberOfArithmeticSlices(nums []int) int {
    res := 0
    var cur []int
    var help func(i int)
    help = func(i int) {
        if i == len(nums) {
            if len(cur) < 3 {
                return
            }
            d := cur[1]-cur[0]
            for i := 2; i < len(cur); i++ {
                if cur[i]-cur[i-1] != d {
                    return
                }
            }
            res++
            return
        }
        // 不使用当前元素
        help(i+1)

        // 使用当前元素
        cur = append(cur, nums[i])
        help(i+1)
        // 回溯
        cur = cur[:len(cur)-1]
    }
    help(0)
    return res
}
```

时间复杂度：`O(2^n)`，空间复杂度：O(n)。

### 动态规划

假设 dp[i][d] 表示仅考虑 nums 中前 i 个元素，以 nums[i-1] 结尾、公差为d的等差数列的数量，显然可以用二重循环方便地用动态规划的方式求解。

需要注意的是，我们可以先削弱等差数列的定义，即有两个元素的数组页看成等差数列，但在统计结果时考虑个数至少3。

注意公差有可能为负数，且其范围较大，dp 数组的第二维可以用哈希表。

```go
func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }

    res := 0
    dp := make([]map[int]int, n)
    for i := 1; i < n; i++ {
        dp[i] = map[int]int{}
        for j := i-1; j >= 0; j-- {
            d := nums[i]-nums[j]
            res += dp[j][d] // 不考虑仅含 nums[j], nums[i] 两个元素的情况
            dp[i][d] += dp[j][d]+1
        }
    }
    return res
}
```

时间复杂度：`O(n^2)`，空间复杂度：`O(n^2)`。


