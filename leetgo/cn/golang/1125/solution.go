package solution

/*
## [1125. Smallest Sufficient Team](https://leetcode.cn/problems/smallest-sufficient-team) (Hard)

作为项目经理，你规划了一份需求的技能清单 `req_skills`，并打算从备选人员名单 `people` 中选出些人组成一个「必要团队」（ 编号为 `i` 的备选人员 `people[i]` 含有一份该备选人员掌握的技能列表）。

所谓「必要团队」，就是在这个团队中，对于所需求的技能列表 `req_skills` 中列出的每项技能，团队中至少有一名成员已经掌握。可以用每个人的编号来表示团队中的成员：

- 例如，团队 `team = [0, 1, 3]` 表示掌握技能分别为 `people[0]`， `people[1]`，和 `people[3]` 的备选人员。

请你返回 **任一** 规模最小的必要团队，团队成员用人员编号表示。你可以按 **任意顺序** 返回答案，题目数据保证答案存在。

**示例 1：**

```
输入：req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
输出：[0,2]

```

**示例 2：**

```
输入：req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [["algorithms","math","java"],["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
输出：[1,2]

```

**提示：**

- `1 <= req_skills.length <= 16`
- `1 <= req_skills[i].length <= 16`
- `req_skills[i]` 由小写英文字母组成
- `req_skills` 中的所有字符串 **互不相同**
- `1 <= people.length <= 60`
- `0 <= people[i].length <= 16`
- `1 <= people[i][j].length <= 16`
- `people[i][j]` 由小写英文字母组成
- `people[i]` 中的所有字符串 **互不相同**
- `people[i]` 中的每个技能是 `req_skills` 中的技能
- 题目数据保证「必要团队」一定存在


*/
/*
   1、有没有贪心策略？比如将人员按照技能数降序排列，然后从前向后选，只到覆盖所有需要的技能。这个策略是错的，如需要技能1、2、3、4、5， 人员技能是 （1，2，3）（1，3，5）（1，4，5），实际只需两人，但按照贪心策略3人都选上了。
   2、注意到数据范围非常有限，不妨来穷举。因req_skills 不会超过16，那么可以用一个整数来表示需要的技能集合，每个人的技能集合同样可用一个整数来表示。这样我们可以穷举技能集合的状态（0 - 1<<(n-1))，实际上能做动态规划。
*/
// [start] don't modify
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	n := len(req_skills)
	idx := make(map[string]int, n)
	for i, v := range req_skills {
		idx[v] = i
	}
	getSkillSet := func(skills []string) int {
		res := 0
		for _, s := range skills {
			i := idx[s]
			res |= 1 << i
		}
		return res
	}
	// the needed skills set is 1<<n -1
	// dp[i] stands for the people list to satisfy the skills set i
	dp := make([][]int, 1<<n)
	dp[0] = []int{} // see ↓
	for i, p := range people {
		cur := getSkillSet(p)
		for pre := 0; pre < 1<<n; pre++ {
			if dp[pre] == nil {
				continue
			} // see ↑
			comb := pre | cur
			if dp[comb] == nil || len(dp[comb]) > len(dp[pre])+1 {
				dp[comb] = append(append([]int{}, dp[pre]...), i)
			}
		}
	}
	return dp[(1<<n)-1]
}

// time complex: O(m^2*2^n), space complex: O(m*2^n)

/* improve
每一个状态都用数组记录了具体的人员编号。这个过程浪费很多空间去储存结果，也消耗了很多时间去生成数组。
实际上只要记录下每个状态的产生来源，就可以按序还原每个状态的具体人员编号的数组。
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    n, m := len(req_skills), len(people)
    idx := make(map[string]int)
    for i, v := range req_skills {
        idx[v] = i
    }
    getSkillSet := func(skills []string) int {
        res := 0
        for _, s := range skills {
            i := idx[s]
            res |= 1 << i
        }
        return res
    }
    dp := make([]int, 1 << n)
    for i := range dp {
        dp[i] = m // m is the max value
    }
    dp[0] = 0
    preSkill := make([]int, 1 << n)
    prePeople := make([]int, 1 << n)
    for i, p := range people {
        cur := getSkillSet(p)
        for pre := 0; pre < (1 << n); pre++ {
            comb := pre | cur
            if dp[comb] > dp[pre]+1 {
                dp[comb] = dp[pre] + 1
                preSkill[comb] = pre
                prePeople[comb] = i
            }
        }
    }
    res := []int{}
    i := (1 << n) - 1
    for i > 0 {
        res = append(res, prePeople[i])
        i = preSkill[i]
    }
    return res
}
// time complex: O(m*2^n), space complex: O(2^n)
*/

// [end] don't modify