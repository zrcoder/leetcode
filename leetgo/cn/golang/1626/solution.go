package solution

/*
## [1626. Best Team With No Conflicts](https://leetcode.cn/problems/best-team-with-no-conflicts) (Medium)

假设你是球队的经理。对于即将到来的锦标赛，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数 **总和** 。

然而，球队中的矛盾会限制球员的发挥，所以必须选出一支 **没有矛盾** 的球队。如果一名年龄较小球员的分数 **严格大于** 一名年龄较大的球员，则存在矛盾。同龄球员之间不会发生矛盾。

给你两个列表 `scores` 和 `ages`，其中每组 `scores[i]` 和 `ages[i]` 表示第 `i` 名球员的分数和年龄。请你返回 **所有可能的无矛盾球队中得分最高那支的分数** 。

**示例 1：**

```
输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
输出：34
解释：你可以选中所有球员。
```

**示例 2：**

```
输入：scores = [4,5,6,5], ages = [2,1,2,1]
输出：16
解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。

```

**示例 3：**

```
输入：scores = [1,2,3,5], ages = [8,9,10,1]
输出：6
解释：最佳的选择是前 3 名球员。

```

**提示：**

- `1 <= scores.length, ages.length <= 1000`
- `scores.length == ages.length`
- `1 <= scores[i] <= 10⁶`
- `1 <= ages[i] <= 1000`


*/

// [start] don't modify
func bestTeamScore(scores []int, ages []int) int {
	type Person struct {
		Age, Score int
	}
	persons := make([]Person, len(ages))
	for i := range ages {
		persons[i].Age = ages[i]
		persons[i].Score = scores[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		a, b := persons[i], persons[j]
		return a.Age < b.Age || a.Age == b.Age && a.Score < b.Score
	})
	dp := make([]int, len(ages))
	res := 0
	for i, v := range persons {
		max := 0
		for j := range dp[:i] {
			if persons[j].Score <= v.Score && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = v.Score + max
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// [end] don't modify