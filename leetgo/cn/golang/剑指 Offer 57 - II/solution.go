package solution

/*
## [剑指 Offer 57 - II. 和为s的连续正数序列 LCOF](https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof) (Easy)

输入一个正整数 `target` ，输出所有和为 `target` 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

**示例 1：**

```
输入：target = 9
输出：[[2,3,4],[4,5]]

```

**示例 2：**

```
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

```

**限制：**

- `1 <= target <= 10^5`


*/

// [start] don't modify
func findContinuousSequence(target int) [][]int {
    res := [][]int{}
    start, end := 1, 2
    for start < end {
        sum := (start+end)*(end-start+1)/2
        if sum == target {
            res = append(res, getSeq(start, end))
            start++ // end++ is also right
        } else if sum < target {
            end++
        } else {
            start++
        }
    }
    return res
}

func getSeq(start, end int) []int {
    res := make([]int, end-start+1)
    for i := start; i <= end; i++ {
        res[i-start] = i
    }
    return res
}

// [end] don't modify
