---
title: "二维排序问题"
date: 2021-04-19T22:04:56+08:00
weight: 1

tags: [排序, 动态规划, 贪心, 二维排序]
---

让我们从经典的 `lis` 即`最长递增子序列`问题开始：

## [300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)

难度中等

给你一个整数数组 `nums` ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，`[3,6,2,7]` 是数组 `[0,3,1,6,2,2,7]` 的子序列。

**示例 1：**

```
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
```

**示例 2：**

```
输入：nums = [0,1,0,3,2,3]
输出：4
```

**示例 3：**

```
输入：nums = [7,7,7,7,7,7,7]
输出：1
```

函数签名：

```go
func lengthOfLIS(nums []int) int
```

## 分析

### 动态规划

这个问题用动态规划是经典解法。如下：

```go
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) // dp[i]代表以nums[i]结尾的递增子序列长度
	maxLen := 0
	for i, v := range nums {
		dp[i] = 1 // 一个元素算递增长度为1
		for j := 0; j < i; j++ {
			if nums[j] < v {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
```

时间复杂度 `O(n^2)`，空间复杂度 `O(n)`。

时间复杂度并不是最优秀的。

### 贪心

这是今天的主角。尝试贪心地构建结果：

```
如果要使上升子序列尽可能长，则需要让序列上升得尽可能慢，因此在构建结果的时候，每次在上升子序列最后加上的那个数需要尽可能小。
建立 memo 数组，memo[i]代表长度为 i+1 的递增子序列末尾数字
遍历 nums，对于当前元素：
如果大于结果数组最后元素，直接追加到结果数组最后；
否则，在结果数组里找到第一个不小于当前元素的元素，并将其更新为当前元素。
这里可以用二分法降低复杂度。

以 [2,1,5,3,4,8,9,7] 为例，可以得到 memo 数组为 [1,3,4,7,9]，这表示：
长度为 1 的递增子序列，最佳末尾数字是 1
长度为 2 的递增子序列，最佳末尾数字是 3
长度为 3 的递增子序列，最佳末尾数字是 4
长度为 4 的递增子序列，最佳末尾数字是 7
长度为 5 的递增子序列，最佳末尾数字是 9

可见，memo 数组的长度就是最长递增子序列的长度。
```
> 实际上，以上做法是一个不完全的耐心排序(patience sorting)。没有完全排序所有元素，而是借助耐心排序的第一部分，得到了最长递增子序列的长度。

```go
func lengthOfLIS(nums []int) int {
	memo := make([]int, len(nums))
	length := 0
	for _, v := range nums {
		j := sort.Search(length, func(i int) bool {
			return memo[i] >= v
		})
		memo[j] = v
		if j == length {
			length++
		}
	}
	return length
}
```

时间复杂度`O(nlogn)`, 空间复杂度`O(n)`。

> 如果允许修改原数组，也可以直接在 `nums` 原地构建结果数组，而不用新建。

**“在结果数组里找到第一个不小于当前元素的元素，并将其更新为当前元素”** 这个贪心策略使二分成为可能，从而降低了时间复杂度。

类似的贪心策略可以扩展应用到一类二维排序问题：

## [354. 俄罗斯套娃信封问题](https://leetcode-cn.com/problems/russian-doll-envelopes/)

难度困难

给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 `(w, h)` 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

**说明:**
不允许旋转信封。

**示例:**

```
输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3 
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
```

函数签名：
```go
func maxEnvelopes(envelopes [][]int) int
```

## 分析

### 贪心

可以先粗略排序，显然大信封在前，小信封在后，因为可能存在宽度相等或高度相等的信封，粗排的时候显然只能先按照一个维度去排，在另一个维度相等时，应该怎么排序？先看后边要怎么构建结果，再回过头来看这个问题就行。

然后在这个有序数组基础上构建结果。构建的时候可以采用类似上边 `lis` 问题的贪心策略：先在结果中从左到右找到第一个不比当前信封大的信封，然后用当前信封替换它即可，当然如果没找到，直接追加到结果数组末尾。

假设粗排的时候是按照宽度降序的，那么为了使上边的贪心策略无误，在宽度相等的时候应该怎么排呢？是的，高度小的排前边。

```go
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] > envelopes[j][0]
	})
	length := 0
	for _, v := range envelopes {
		j := sort.Search(length, func(i int) bool {
			c := envelopes[i]
			return c[0] <= v[0] || c[1] <= v[1]
		})
		envelopes[j] = v
		if j == length {
			length++
		}
	}
	return length
}
```

时间复杂度 `O(nlogn)`，因为使用了原数组原地操作，空间复杂度为`O(1)`。

>  如果不允许修改原数组，只需新建一个数组来构建结果即可。

### 动态规划

这个问题同样可以用动态规划，思路同样非常像 `lis` 问题的动态规划解法。只需要先从一个维度排序，把问题降维后就好用动态规划了。

```go
func maxEnvelopes(envelopes [][]int) int {
	// 先按一个维度排序（宽度或高度都行）；只是把问题降维
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] > envelopes[j][0]
	})
	result := 0
	dp := make([]int, len(envelopes))
	for i, v := range envelopes {
		dp[i] = 1
		for j := 0; j < i; j++ {
			vj := envelopes[j]
			if vj[0] > v[0] && vj[1] > v[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}
```

时间复杂度 `O(n^2)`，空间复杂度`O(n)`。


> 一个几乎相同的问题是：[面试题 17.08. 马戏团人塔](https://leetcode-cn.com/problems/circus-tower-lcci/),可以练习。

## [406. 根据身高重建队列](https://leetcode-cn.com/problems/queue-reconstruction-by-height/)

难度中等

假设有打乱顺序的一群人站成一个队列，数组 `people` 表示队列中一些人的属性（不一定按顺序）。每个 `people[i] = [hi, ki]` 表示第 `i` 个人的身高为 `hi` ，前面 **正好** 有 `ki` 个身高大于或等于 `hi` 的人。

请你重新构造并返回输入数组 `people` 所表示的队列。返回的队列应该格式化为数组 `queue` ，其中 `queue[j] = [hj, kj]` 是队列中第 `j` 个人的属性（`queue[0]` 是排在队列前面的人）。

**示例 1：**

```
输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
解释：
编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
```

**示例 2：**

```
输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
```

**提示：**

- `1 <= people.length <= 2000`
- `0 <= hi <= 106`
- `0 <= ki < people.length`
- 题目数据确保队列可以被重建

## 分析

题意有点难理解。可以加点背景说明：一开始人们随便站成了一队，然后班长统计了每个人的身高 h 以及排在其前边不比自己矮的人的个数 k。突然这些人一哄而散跑去看美女了。问题是恢复原来的队列。

这个问题总体思路和上边的信封套娃及叠罗汉问题类似。都是排序，经历粗排和细排两轮。

很自然的思路：越高的人k 值理应越小。先按照身高降序，在身高相等的时候怎么排呢？k 小的排前边。

在构建结果数组的时候，如果当前人的 k 不小于结果数组的长度，直接把他追加到队尾，否则，用二分法找到他该插入的位置，当然后边的人要一一后移。

> 这里需要注意，其实用二分法反而有点浪费，先二分再一一后移一些人，不如放弃二分，一开始直接从结果数组后边向前找，类似冒泡的方法，将当前人插入队里，而他后边的人在冒泡的过程中已经移动好了。

```go
func reconstructQueue(people [][]int) [][]int {
	// 高的排前边，一样高的按照k升序排列
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	result := make([][]int, len(people))
	length := 0
	for _, p := range people {
		k := p[1]
		i := length
		for i > k { // 根据前边的排序，实际不会出现 k > length 的情况
			result[i] = result[i-1]
			i--
		}
		result[i] = p
		length++
	}
	return result
}
```

时间复杂度 `O(n^2)`，空间复杂度 `O(n)`。

## [630. 课程表 III](https://leetcode-cn.com/problems/course-schedule-iii/)

难度困难

这里有 `n` 门不同的在线课程，他们按从 `1` 到 `n` 编号。每一门课程有一定的持续上课时间（课程时间）`t` 以及关闭时间第 d 天。一门课要持续学习 `t` 天直到第 d 天时要完成，你将会从第 1 天开始。

给出 `n` 个在线课程用 `(t, d)` 对表示。你的任务是找出最多可以修几门课。

**示例：**

```
输入: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
输出: 3
解释: 
这里一共有 4 门课程, 但是你最多可以修 3 门:
首先, 修第一门课时, 它要耗费 100 天，你会在第 100 天完成, 在第 101 天准备下门课。
第二, 修第三门课时, 它会耗费 1000 天，所以你将在第 1100 天的时候完成它, 以及在第 1101 天开始准备下门课程。
第三, 修第二门课时, 它会耗时 200 天，所以你将会在第 1300 天时完成它。
第四门课现在不能修，因为你将会在第 3300 天完成它，这已经超出了关闭日期。
```

**提示:**

1. 整数 1 <= d, t, n <= 10,000 。
2. 你不能同时修两门课程。

函数签名：

```go
func scheduleCourse(courses [][]int) int
```

## 分析

为了修尽可能多的课程，要优先修那些关闭时间早的课程。所以首先可以按照关闭时间将所有课程排序，接下来遍历这些课程，对于当前课程，如果已经花费的时间加上该课程需要的时间没有超过关闭时间，则暂定修这门课，否则，需要在已经暂定要修的课程里找到话费时间最长的那门课，决定不修那一门，当然有可能其比当前课程需要的时间短，那就不修当前课程。

为了迅速找到耗时最长的课程，可以用大顶堆，另用一个变量 day 维护已经花费的时间。每次先将当前课程的耗时入堆，如果发现已经花费的时间加上当前课程需要的时间超过了当前课程的关闭时间，需要将堆顶元素出堆，且更新 day，即减去堆顶元素。

```go
func scheduleCourse(courses [][]int) int {
    sort.Slice(courses, func(i, j int) bool {
        return courses[i][1] < courses[j][1] 
    })
    h := &Heap{cmp: func(a, b int) bool {
        return a > b
    }}
    day := 0
    for _, v := range courses {
        if v[0] > v[1] {
            continue
        }
        heap.Push(h, v[0])
        day += v[0]
        if day > v[1] {
            day -= heap.Pop(h).(int)
        }
    }
    return h.Len()
}
```

时间复杂度 `O(nlogn)`，空间复杂度 `O(n)`。

附堆的相关实现：

```go
type Cmp func(int, int) bool

type Heap struct {
	slice []int
	cmp   Cmp
}

// implement heap.Interface
func (h *Heap) Len() int           { return len(h.slice) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.slice[i], h.slice[j]) }
func (h *Heap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.slice[h.Len()-1]
	h.slice = h.slice[:h.Len()-1]
	return x
}
```

## 扩展

## [368. 最大整除子集](https://leetcode-cn.com/problems/largest-divisible-subset/)

给你一个由 **无重复** 正整数组成的集合 `nums` ，请你找出并返回其中最大的整除子集 `answer` ，子集中每一元素对 `(answer[i], answer[j])` 都应当满足：

- `answer[i] % answer[j] == 0` ，或
- `answer[j] % answer[i] == 0`

如果存在多个有效解子集，返回其中任何一个均可。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。
```

**示例 2：**

```
输入：nums = [1,2,4,8]
输出：[1,2,4,8]
```

**提示：**

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 2 * 109`
- `nums` 中的所有整数 **互不相同**

## 分析

这个问题用类似的动态规划来解决。

参考代码：
```go
func largestDivisibleSubset(nums []int) (res []int) {
	// 排序预处理
	sort.Ints(nums)
	n := len(nums)
	// 动态规划确定每个位置结尾所能得到的满足约束的子序列最大长度
	dp := make([]int, n)
	dp[0] = 1
	// index和maxSize用来维护满足约束的最长子序列的末尾和长度，方便后边构造出结果
	index, maxSize := 0, 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			index = i
			maxSize = dp[i]
		}
	}
	// 构造结果
	if index == 0 {
		return []int{nums[0]}
	}
	maxVal := nums[index]
	for i := index; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return
}
```

## [面试题 08.13. 堆箱子](https://leetcode-cn.com/problems/pile-box-lcci/)

难度困难

堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。

输入使用数组`[wi, di, hi]`表示每个箱子。

**示例1:**

```
 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
 输出：6
```

**示例2:**

```
 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
 输出：10
```

**提示:**

1. 箱子的数目不大于3000个。

## 分析

类似上面的动态规划，实现需要把箱子在某一维（长/宽/高）排序。

参考实现：

```go
func pileBox(box [][]int) int {
    sort.Slice(box, func(i, j int) bool {
        return box[i][0] > box[j][0]
    })
    dp := make([]int, len(box))
    res := 0
    for i, v := range box {
        dp[i] = v[2]
        for j := 0; j < i; j++ {
            if box[j][0] > v[0] && box[j][1] > v[1] && box[j][2] > v[2] {
                dp[i] = max(dp[i], dp[j]+v[2])
            }
        }
        res = max(res, dp[i])
    }
    return res
}
```