package solution

/*
## [1144. Decrease Elements To Make Array Zigzag](https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag) (Medium)

给你一个整数数组 `nums`，每次 **操作** 会从中选择一个元素并 **将该元素的值减少 1**。

如果符合下列情况之一，则数组 `A` 就是 **锯齿数组**：

- 每个偶数索引对应的元素都大于相邻的元素，即 `A[0] > A[1] < A[2] > A[3] < A[4] > ...`
- 或者，每个奇数索引对应的元素都大于相邻的元素，即 `A[0] < A[1] > A[2] < A[3] > A[4] < ...`

返回将数组 `nums` 转换为锯齿数组所需的最小操作次数。

**示例 1：**

```
输入：nums = [1,2,3]
输出：2
解释：我们可以把 2 递减到 0，或把 3 递减到 1。

```

**示例 2：**

```
输入：nums = [9,6,1,6,2]
输出：4

```

**提示：**

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 1000`


*/

// [start] don't modify
func movesToMakeZigzag(nums []int) int {
    // 注意只能减少不能增加。只需要所有偶数位元素不变，只改变奇数位，或所有奇数位不变，只改变偶数位
    cal := func(start int) int {
        res := 0
        for i := start; i < len(nums); i += 2 {
            left, right := math.MaxInt32, math.MaxInt32
            if i-1 >= 0 {
                left = nums[i-1]
            }
            if i+1 < len(nums) {
                right = nums[i+1]
            }
            res += max(0, nums[i]-min(left, right)+1) // 如果nums[i] 本来就比左右都小不用修改
        }
        return res
    }
    return min(cal(0), cal(1))
}
// [end] don't modify
