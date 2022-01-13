package solution

/*
## [1042. Flower Planting With No Adjacent](https://leetcode.cn/problems/flower-planting-with-no-adjacent) (Medium)

有 `n` 个花园，按从 `1` 到 `n` 标记。另有数组 `paths` ，其中 `paths[i] = [xᵢ, yᵢ]` 描述了花园 `xᵢ` 到花园 `yᵢ` 的双向路径。在每个花园中，你打算种下四种花之一。

另外，所有花园 **最多** 有 **3** 条路径可以进入或离开.

你需要为每个花园选择一种花，使得通过路径相连的任何两个花园中的花的种类互不相同。

以数组形式返回 **任一** 可行的方案作为答案 `answer`，其中 `answer[i]` 为在第 `(i+1)` 个花园中种植的花的种类。花的种类用 1、2、3、4 表示。保证存在答案。

**示例 1：**

```
输入：n = 3, paths = [[1,2],[2,3],[3,1]]
输出：[1,2,3]
解释：
花园 1 和 2 花的种类不同。
花园 2 和 3 花的种类不同。
花园 3 和 1 花的种类不同。
因此，[1,2,3] 是一个满足题意的答案。其他满足题意的答案有 [1,2,4]、[1,4,2] 和 [3,2,1]

```

**示例 2：**

```
输入：n = 4, paths = [[1,2],[3,4]]
输出：[1,2,1,2]

```

**示例 3：**

```
输入：n = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
输出：[1,2,3,4]

```

**提示：**

- `1 <= n <= 10⁴`
- `0 <= paths.length <= 2 * 10⁴`
- `paths[i].length == 2`
- `1 <= xᵢ, yᵢ <= n`
- `xᵢ != yᵢ`
- 每个花园 **最多** 有 **3** 条路径可以进入或离开


*/

// [start] don't modify
func gardenNoAdj(n int, paths [][]int) []int {
	graph := make([][]int, n+1)
	for _, p := range paths {
		graph[p[0]] = append(graph[p[0]], p[1])
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	colors := make([]int, n+1)
	var mark func(int)
	mark = func(i int) {
		if colors[i] > 0 {
			return
		}
		nextColors := [5]bool{}
		for _, nei := range graph[i] {
			nextColors[colors[nei]] = true
		}
		for c := 1; c <= 4; c++ {
			if !nextColors[c] {
				colors[i] = c
				break
			}
		}
		for _, nei := range graph[i] {
			mark(nei)
		}
	}
	for i := 1; i <= n; i++ {
		mark(i)
	}
	return colors[1:]
}

// [end] don't modify
