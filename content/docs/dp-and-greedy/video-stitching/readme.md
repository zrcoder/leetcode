---
title: "1024. 视频剪辑"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [1024. 视频剪辑](https://leetcode-cn.com/problems/video-stitching)
你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。

视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。

我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。

 
```
示例 1：

输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
输出：3
解释：
我们选中 [0,2], [8,10], [1,9] 这三个片段。
然后，按下面的方案重制比赛片段：
将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。
示例 2：

输入：clips = [[0,1],[1,2]], T = 5
输出：-1
解释：
我们无法只用 [0,1] 和 [1,2] 覆盖 [0,5] 的整个过程。
示例 3：

输入：clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], T = 9
输出：3
解释： 
我们选取片段 [0,4], [4,7] 和 [6,9] 。
示例 4：

输入：clips = [[0,4],[2,8]], T = 5
输出：2
解释：
注意，你可能录制超过比赛结束时间的视频。
 

提示：

1 <= clips.length <= 100
0 <= clips[i][0] <= clips[i][1] <= 100
0 <= T <= 100
```
## 分析
- 首先想到一个有疏漏贪心的策略
  ```
  先排序，所有片段按照开始时间升序，若开始时间相同则按结束时间降序排列
  再遍历，用一个变量 last 记录已选区间所能达到的右边界：
    若当前视频段开始时间比 last 还大，可以确定无解；
    若当前视频段结束时间不大于 last，则当前视频段不入选，继续遍历后边的
    否则，把当前视频段加入结果
    最后返回结果，注意如果 last 最终小于 T，则无解，返回-1
  ```

  ```go
  func videoStitching(clips [][]int, T int) int {
    n := len(clips)
    if n == 0 {return -1}
    sort.Slice(clips, func(i, j int) bool {
        if clips[i][0] == clips[j][0] {
            return clips[i][1] > clips[j][1]
        }
        return clips[i][0] < clips[j][0]
    })
    res := 0
    last := 0
    for _, v := range clips {
        if last < v[0] {
            return -1
        }
        if last >= v[1] {
            continue
        }
        res++
        last = v[1]
    }
    if last < T {
        return -1
    }
    return res
  }
  ```
  不过还有疏漏，比如这样的用例：   
  ```
  [0 4],[2 6],[4 7],[6 9] 9
  ```
  显然，[2,6]或[4,7]只能选一个，但上边的策略把两个都选了；这个问题还没想到修改都办法。

- 另一个贪心策略
  ```
  可以先把 clips 降为长度为 T 的一维数组，索引代表开始时间，值代表结束时间。
  在有相同开始时间的视频片段时，只保留结束时间大的片段。

  接下来遍历这个降维后的数组，用 maxEnd 维护能到达的最大右边界，pre 维护上一次选择的区间的右边界
    首先，如果当前视频片段结束时间大于 maxEnd，更新 maxEnd
    如果当前视频片段开始时间 == maxEnd，则无解，直接返回 -1
    如果当前视频片段开始时间 == pre， 则当前片段应该加入结果，pre 更新为 maxEnd

  ```

  ```go
  func videoStitching(clips [][]int, T int) int {
      if T < 1 {
          return 0
      }
      ends := make([]int, T)
      for _, v := range clips {
          if v[0] < T && ends[v[0]] < v[1] {
              ends[v[0]] = v[1]
          }
      }
      maxEnd, pre, res := 0, 0, 0
      for start, end := range ends {
          if maxEnd < end {
              maxEnd = end
          }
          if start == maxEnd {
              return -1
          }
          if start == pre {
              res++
              pre = maxEnd
          }
      }
      return res
  }
  ```

- 动态规划
  
  dp[i] 表示将区间 [0,i) 覆盖所需的最少子区间的数量

  ```go
  func videoStitching(clips [][]int, t int) int {
    const inf = 200
    dp := make([]int, t+1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0
    for i := 1; i <= t; i++ {
        for _, c := range clips {
            l, r := c[0], c[1]
            // 若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
            if l < i && i <= r && dp[l]+1 < dp[i] {
                dp[i] = dp[l] + 1
            }
        }
    }
    if dp[t] == inf {
        return -1
    }
    return dp[t]
  }
  ```