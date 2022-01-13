package solution

/*
## [1470. Shuffle the Array](https://leetcode.cn/problems/shuffle-the-array) (Easy)

给你一个数组 `nums` ，数组中有 `2n` 个元素，按 `[x₁,x₂,...,xₙ,y₁,y₂,...,yₙ]` 的格式排列。

请你将数组按 `[x₁,y₁,x₂,y₂,...,xₙ,yₙ]` 格式重新排列，返回重排后的数组。

**示例 1：**

```
输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x₁=2, x₂=5, x₃=1, y₁=3, y₂=4, y₃=7 ，所以答案为 [2,3,5,4,1,7]

```

**示例 2：**

```
输入：nums = [1,2,3,4,4,3,2,1], n = 4
输出：[1,4,2,3,3,2,4,1]

```

**示例 3：**

```
输入：nums = [1,1,2,2], n = 2
输出：[1,2,1,2]

```

**提示：**

- `1 <= n <= 500`
- `nums.length == 2n`
- `1 <= nums[i] <= 10^3`


*/

// [start] don't modify
func shuffle(nums []int, n int) []int {
    res := make([]int, 2*n)
    i, j := 0, n
    for k := 0; k < len(res); k += 2 {
        res[k] = nums[i]
        i++
        res[k+1] = nums[j]
        j++
    }
    return res
}
// [end] don't modify
