---
title: "子数组/子序列宽度和"
date: 2022-11-18T21:57:54+08:00
---

## [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)

难度中等

给你一个整数数组 `nums` 。`nums` 中，子数组的 **范围** 是子数组中最大元素和最小元素的差值。

返回 `nums` 中 **所有** 子数组范围的 **和** *。*

子数组是数组中一个连续 **非空** 的元素序列。

**示例 1：**

**输入：** nums = [1,2,3]
**输出：** 4
**解释：** nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0 
[2]，范围 = 2 - 2 = 0
[3]，范围 = 3 - 3 = 0
[1,2]，范围 = 2 - 1 = 1
[2,3]，范围 = 3 - 2 = 1
[1,2,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4

**示例 2：**

**输入：** nums = [1,3,3]
**输出：** 4
**解释：** nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[3]，范围 = 3 - 3 = 0
[3]，范围 = 3 - 3 = 0
[1,3]，范围 = 3 - 1 = 2
[3,3]，范围 = 3 - 3 = 0
[1,3,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 2 + 0 + 2 = 4

**示例 3：**

**输入：** nums = [4,-2,-3,4,1]
**输出：** 59
**解释：** nums 中所有子数组范围的和是 59

**提示：**

- `1 <= nums.length <= 1000`
- `-109 <= nums[i] <= 109`

**进阶：** 你可以设计一种时间复杂度为 `O(n)` 的解决方案吗？

## 分析
### 枚举子数组

可以用两层循环来枚举所有的子数组，并且在枚举过程中维护字数组最大值和最小值。

```go
func subArrayRanges(nums []int) int64 {
    var res int64
    for i := range nums { // i 作为子数组左边界
        min, max := nums[i], nums[i]
        for j := i+1; j < len(nums); j++ { // j 作为子数组右边界
            if nums[j] < min {
                min = nums[j]
            }
            if nums[j] > max {
                max = nums[j]
            }
            res += int64(max)-int64(min)
        }
    }
    return res
}
```

显然时间复杂度为`O(n^2)`,其中 n 是数组长度。
空间复杂度为 `O(1)`

### 单调栈
可以将问题作一转化，对于数组中每一个元素，可以向前后扩展得到一个子数组，这个子数组的最小值/最大值恰好是该元素，那么可以计算出这个元素对最终结果的贡献值。

具体说，对于元素 nums[i]， 假设向左找到第一个大于它的元素 nums[j]，向右找到第一个大于它的元素nums[k], 显然在 [j+1:k] 这个子数组中，nums[i] 是最小值。在这个子数组里，继续划分，共有多少个子数组包含 nums[i]呢？只需要保证左边界在[j+1, i] 闭区间内，右边界在 [i, k-1] 闭区间内即可，即这样的子数组总数为 (i-j)*(k-i)，则 nums[i] 作为最小值对最终结果的贡献为 (i-j)*(k-i)*(-nums[i])。
> 如果某个元素左边没有比它大的元素，那么可以将其左边第一个比它大的位置看成 -1， 如果右边没有则看成 n；这样也能得到正确的子数组个数。

同样可以计算出 nums[i] 作为最大值对结果的贡献；最终将所有贡献加起来即可。

找到一个元素左边/右边第一个大于/小于该元素的元素，可以用`单调栈`在线性复杂度内求得。

```go
func subArrayRanges(nums []int) int64 {
    n := len(nums)
    // leftLess[i] 表示 i 向左查询到的第一个 < nums[i] 的元素位置
    leftLess := make([]int, n)
    // leftMore[i] 表示 i 向左查询到的第一个 > nums[i] 的元素位置
    leftMore := make([]int, n)
    // rightLess[i] 表示 i 向右查询到的第一个 <= nums[i] 的元素位置
    rightLess := make([]int, n)
    // rightMore[i] 表示 i 向左查询到的第一个 >= nums[i] 的元素位置
    rightMore := make([]int, n)
    stk := make([]int, 0, n) // 辅助用的单调栈
    cal := func(s []int, isReverse bool, cmp func(x, y int) bool) {
        stk = stk[:0]
        from, to := 0, n-1
        delta := 1
        if isReverse {
            from, to = to, from
            delta = -1
        }
        for i := from; !isReverse && i <= to || isReverse && i >= to; i += delta {
            for len(stk) > 0 && cmp(nums[i], nums[stk[len(stk)-1]]) {
                stk = stk[:len(stk)-1]
            }
            if len(stk) == 0 {
                s[i] = from-delta
            } else {
                s[i] = stk[len(stk)-1]
            }
            stk = append(stk, i)
        }
    }
    // 注意比较逻辑里，有的有等号有的没有，是为了枚举所有子数组时不重不漏
    cal(leftLess, false, func(x, y int) bool {return x >= y})
    cal(leftMore, false, func(x, y int) bool {return x <= y})
    cal(rightLess, true, func(x, y int) bool {return x > y})
    cal(rightMore, true, func(x, y int) bool {return x < y})

    var res int64 = 0
    for i, v := range nums {
        res -= int64((i-leftMore[i])*(rightMore[i]-i)*v)
        res += int64((i-leftLess[i])*(rightLess[i]-i)*v)
    }
    return res
}
```

时间复杂度可空间复杂度都是 `O(n)`。关于单调栈的时间复杂度，可以这样考虑：每个元素最多进出单调栈各一次，所以虽然看起来两层循环，实际是线性复杂度。


### 进一步优化
上边单调栈的解法中用了4个额外的数组来存储信息，实际可以优化掉。

```go
func subArrayRanges(nums []int) int64 {
	return minMaxSum(nums, func(a, b int)bool{return a < b}) -
        minMaxSum(nums, func(a, b int) bool {return a > b})
}

func minMaxSum(nums []int, cmp func(int, int) bool) int64 {
    res := int64(0)
    stk := []int{}
    for i := 0; i <= len(nums); i++ {
        for len(stk) > 0 && (i == len(nums) || cmp(nums[stk[len(stk)-1]], nums[i])) {
            j := stk[len(stk)-1]
            stk = stk[:len(stk)-1]
            k := -1
            if len(stk) > 0 {
                k = stk[len(stk)-1]
            }
            // k，i处元素同时大于/小于 j 处元素
            res += int64(nums[j]*(i-j)*(j-k))
        }
        stk = append(stk, i)
    }
    return res
}
```

时空复杂度同上，还是 `O(n)`。逻辑稍有不同，代码更简洁。

## [891. 子序列宽度之和](https://leetcode.cn/problems/sum-of-subsequence-widths/)

难度困难

一个序列的 **宽度** 定义为该序列中最大元素和最小元素的差值。

给你一个整数数组 `nums` ，返回 `nums` 的所有非空 **子序列** 的 **宽度之和** 。由于答案可能非常大，请返回对 `109 + 7` **取余** 后的结果。

**子序列** 定义为从一个数组里删除一些（或者不删除）元素，但不改变剩下元素的顺序得到的数组。例如，`[3,6,2,7]` 就是数组 `[0,3,1,6,2,2,7]` 的一个子序列。

**示例 1：**

**输入：** nums = [2,1,3]
**输出：** 6
**解释：** 子序列为 [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3] 。
相应的宽度是 0, 0, 0, 1, 1, 2, 2 。
宽度之和是 6 。

**示例 2：** 

**输入：** nums = [2]
**输出：** 0

**提示：**

- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 105`

## 分析
这次是子序列而不是子数组。

对于每个元素，需要统计比它小/大的元素个数。这样就能计算得到最终结果。

比如对于元素 x，比它小的有 a 个，比它大的有 b 个，那么 x 作为最大值，可以放到多少个子序列里呢？首先 x 必选，其次只能在比它大的元素里选，比它大的有 a 个，每一个都有选或不选两种情况，所以共有 `2^a` 个子序列以 x 为最小值；同理可知，有 `2^b` 个子序列以 x 为最大值，那么 x 对最终结果的贡献是 `x*(2^b-2^a)。

可以对数组排序，这样能迅速知道对于每个元素，有多少比它大/小的元素。

```go

func sumSubseqWidths(nums []int) int {
    const mod int = 1e9 + 7
    sort.Ints(nums)
    n := len(nums)
    pow2 := make([]int, n)
    pow2[0] = 1
    for i := 1; i < n; i++ {
        pow2[i] = pow2[i-1] * 2 % mod // 预处理 2 的幂次
    }

    res := 0
    for i, v := range nums {
        res = (res + (pow2[i] - pow2[n-1-i]) * v) % mod
    }
    return (res + mod) % mod // 注意上面有减法，res 可能为负数
}
```
时间复杂度 `O(nlogn)`，空间复杂度 `O(n)`。

还可以换个角度来计算，将空间复杂度降到常数级：

```go
func sumSubseqWidths(nums []int) int {
    sort.Ints(nums)
    const mod = 1e9+7
    n := len(nums)
    res := 0
    pow2 := 1
    for i := range nums {
        res = (res + (nums[i]-nums[n-1-i])*pow) % mod
        pow2 = pow2*2%mod
    }
    return (res+mod)%mod
}
```