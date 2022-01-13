package solution

/*
## [2444. Count Subarrays With Fixed Bounds](https://leetcode.cn/problems/count-subarrays-with-fixed-bounds) (Hard)

给你一个整数数组 `nums` 和两个整数 `minK` 以及 `maxK` 。

`nums` 的定界子数组是满足下述条件的一个子数组：

- 子数组中的 **最小值** 等于 `minK` 。
- 子数组中的 **最大值** 等于 `maxK` 。

返回定界子数组的数目。

子数组是数组中的一个连续部分。

**示例 1：**

```
输入：nums = [1,3,5,2,7,5], minK = 1, maxK = 5
输出：2
解释：定界子数组是 [1,3,5] 和 [1,3,5,2] 。

```

**示例 2：**

```
输入：nums = [1,1,1,1], minK = 1, maxK = 1
输出：10
解释：nums 的每个子数组都是一个定界子数组。共有 10 个子数组。
```

**提示：**

- `2 <= nums.length <= 10⁵`
- `1 <= nums[i], minK, maxK <= 10⁶`


*/

// [start] don't modify
func countSubarrays(nums []int, minK int, maxK int) int64 {
    iMin, iMax, iOut := -1, -1, -1
    var res int64
    for i, v := range nums {
        if v == minK {
            iMin = i
        }
        if v == maxK {
            iMax = i
        }
        if v < minK || v > maxK {
            iOut = i
        }
        res += int64(max(0, min(iMin, iMax)-iOut))
    }
    return res
}

func min(a, b int) int {if a < b {return a}; return b}
func max(a, b int) int {if a > b {return a}; return b}
// [end] don't modify
