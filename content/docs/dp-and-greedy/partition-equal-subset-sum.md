---

title: "特定和的子集"
date: 2022-04-07T11:30:20+08:00
math: true

---

在一个集合中找到一个子集，使其和为特定值，可以用一般的回溯穷举方法，还可以结合数据约束，选择复杂度优秀得多的背包解法。

## [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/description/ "https://leetcode-cn.com/problems/partition-equal-subset-sum/description/")

| Category   | Difficulty      | Likes | Dislikes |
| ---------- | --------------- | ----- | -------- |
| algorithms | Medium (51.40%) | 1240  | -        |

给你一个 **只包含正整数** 的 **非空** 数组 `nums` 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

**示例 1：**

```
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
```

**示例 2：**

```
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
```

**提示：**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

函数签名：

```go
func canPartition(nums []int) bool
```

## 分析

假设所有元素和为 sum，问题等价于从数组中挑出一些数字，使其和为 sum/2。

### 回溯穷举

最容易想到回溯解法。

```go
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%2 == 1 {
        return false
    }
    target := sum / 2

    var help func(i, sum int) bool
    help = func(i, sum int) bool {
        if i == len(nums) {
            return sum == target
        }
        if help(i+1, sum) {
            return true
        }
        return help(i+1, sum+nums[i])
    }
    return help(0, 0)
}
```

时间复杂度$O(2^n)$，太高，超时。空间复杂度$O(n)$

### 0-1 背包

实际上是一个0-1背包应用。

```go
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%2 == 1 {
        return false
    }
    target := sum / 2

    dp := make([]bool, target+1)
    dp[0] = true
    for _, v := range nums {
        for j := target; j >= v; j-- {
            dp[j] = dp[j] || dp[j-v]
        }
    }
    return dp[target]
}
```

时间复杂度 `O(n*target)`，空间复杂度 `O(target)`。

## 扩展

### [494. 目标和](https://leetcode-cn.com/problems/target-sum/description/ "https://leetcode-cn.com/problems/target-sum/description/")

给你一个整数数组 `nums` 和一个整数 `target` 。

向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：

- 例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

**示例 1：**

```
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
```

**示例 2：**

```
输入：nums = [1], target = 1
输出：1
```

**提示：**

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 1000`

同样最容易写出回溯解法，实际上还可以稍作转化，成为背包问题。

假设一部分元素最终保持原值，即前边加"+"，这些元素的和为$x$，剩下的一部分（前边加“-”）和就是 $sum-x$，显然 $x-(sum-x) = target$，所以：$x = \frac{target+sum}{2}$，转化成了0-1背包问题。

```go
func findTargetSumWays(nums []int, target int) int {
    // 问题转化为：在nums里找一个子序列，其和为 target
    sum := 0
    for _, v := range nums {
        sum += v
    }
    sum += target
    if sum < 0 || sum%2 != 0 {
        return 0
    }

    target = sum / 2
    dp := make([]int, target+1)
    dp[0] = 1
    for _, v := range nums {
        for j := target; j >= v; j-- {
            dp[j] += dp[j-v]
        }
    }
    return dp[target]
}
```

## 小结

以上两个问题，都可以用普适的穷举解法，复杂度是$O(2^n)$;还可以转化成基于动态规划的背包问题，复杂度是$O(n*target)$。在 target 为正数且值较小时，用第二种方法会非常优秀。
