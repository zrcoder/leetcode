package solution

/*
## [2439. Minimize Maximum of Array](https://leetcode.cn/problems/minimize-maximum-of-array) (Medium)

给你一个下标从 **0** 开始的数组 `nums` ，它含有 `n` 个非负整数。

每一步操作中，你需要：

- 选择一个满足 `1 <= i < n` 的整数 `i` ，且 `nums[i] > 0` 。
- 将 `nums[i]` 减 1 。
- 将 `nums[i - 1]` 加 1 。

你可以对数组执行 **任意** 次上述操作，请你返回可以得到的 `nums` 数组中 **最大值** **最小** 为多少。

**示例 1：**

```
输入：nums = [3,7,1,6]
输出：5
解释：
一串最优操作是：
1. 选择 i = 1 ，nums 变为 [4,6,1,6] 。
2. 选择 i = 3 ，nums 变为 [4,6,2,5] 。
3. 选择 i = 1 ，nums 变为 [5,5,2,5] 。
nums 中最大值为 5 。无法得到比 5 更小的最大值。
所以我们返回 5 。

```

**示例 2：**

```
输入：nums = [10,1]
输出：10
解释：
最优解是不改动 nums ，10 是最大值，所以返回 10 。

```

**提示：**

- `n == nums.length`
- `2 <= n <= 10⁵`
- `0 <= nums[i] <= 10⁹`


*/

// [start] don't modify
func minimizeArrayValue(nums []int) int {
    var res, sum int
    for i, v := range nums {
        sum += v
        res = max(res, (sum+i)/(i+1)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
// [end] don't modify
