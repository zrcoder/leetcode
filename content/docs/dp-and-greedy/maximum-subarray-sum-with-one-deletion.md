---
title: 1186. 删除一次得到子数组最大和
---

## [1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion) (Medium)

给你一个整数数组，返回它的某个 **非空** 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。

注意，删除一个元素后，子数组 **不能为空**。

**示例 1：**

```
输入：arr = [1,-2,0,3]
输出：4
解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
```

**示例 2：**

```
输入：arr = [1,-2,-2,3]
输出：3
解释：我们直接选出 [3]，这就是最大和。

```

**示例 3：**

```
输入：arr = [-1,-1,-1,-1]
输出：-1
解释：最后得到的子数组不能为空，所以我们不能选择 [-1] 并从中删去 -1 来得到 0。
     我们应该直接选择 [-1]，或者选择 [-1, -1] 再从中删去一个 -1。

```

**提示：**

- `1 <= arr.length <= 10⁵`
- `-10⁴ <= arr[i] <= 10⁴`

## 分析

### 朴素解法

两层循环枚举所有子数组和子数组的最小值。时间复杂度 O(n^2)，过高。

```go
func maximumSum(arr []int) int {
	res := math.MinInt
	for i := 0; i < len(arr); i++ {
		sum := 0
		min := math.MaxInt
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if arr[j] < min {
				min = arr[j]
			}
			if min < 0 && j != i {
				tmp := sum - min
				if tmp > res {
					res = tmp
				}
			} else if sum > res {
				res = sum
			}
		}
	}
	return res
}
```

### 动态规划

维护两个数组 f 和 g，f[i] 代表不删除元素且以第i个元素结尾的子串的最大和；而 g[i] 代表删除元素且以第i个元素结尾的子串的最大和.

那么状态转移方程为：

```text
f[i] = max(f[i-1]+arr[i], arr[i])
g[i] = max(g[i-1]+arr[i], f[i-1])
```

初始状态：

```text
f[0] = arr[0]
g[0] = -inf
```

时空复杂度均为 O(n)。

```go
func maximumSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	f, g := make([]int, n), make([]int, n)
	f[0] = arr[0]
	g[0] = math.MinInt32
	res := arr[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1]+arr[i], arr[i])
		g[i] = max(g[i-1]+arr[i], f[i-1])
		res = max(res, f[i], g[i])
	}
	return res
}
```

很容易把空间复杂度降低到 O(1):

```go
func maximumSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	f, g := arr[0], math.MinInt32
	res := arr[0]
	for i := 1; i < n; i++ {
		f, g = max(f+arr[i], arr[i]), max(g+arr[i], f)
		res = max(res, f, g)
	}
	return res
}

func max(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

```
