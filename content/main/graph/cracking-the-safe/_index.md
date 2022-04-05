---
title: "753. 破解保险箱"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [753. 破解保险箱](https://leetcode-cn.com/problems/cracking-the-safe)
有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位是 k 位序列 0, 1, ..., k-1 中的一个 。  
你可以随意输入密码，保险箱会自动记住最后 n 位输入，如果匹配，则能够打开保险箱。  
举个例子，假设密码是 "345"，你可以输入 "012345" 来打开它，只是你输入了 6 个字符.  
请返回一个能打开保险箱的最短字符串。
```
示例1:
输入: n = 1, k = 2
输出: "01"
说明: "10"也可以打开保险箱。

示例2:
输入: n = 2, k = 2
输出: "00110"
说明: "01100", "10011", "11001" 也能打开保险箱。

提示：
n 的范围是 [1, 4]。
k 的范围是 [1, 10]。
k^n 最大可能为 4096。
```
## 分析
即需要穷举 n 位 k 进制数的排列。

题目描述太抽象，需要再梳理。以题目中两个例子为例：
```
n=1, k=2
密码只有一位，每位只能是0或1
那么密码可能是0，也可能是1
01 或 10 就是能打开保险箱的最短字符串
```
```
n=2， k=2
密码有2位，每一位只能是0或1，
那么密码只能是 00、 01、 10、 11， 其中一个
11001，10011，01100，00110 就是能打开保险箱的最短字符串
```

## DFS

一般化，总共有 n 位， 每位是一个 k 进制数（值在 [0,k) 区间内），要穷举所有密码的可能，最终得到一个最短字符串。

可以想象一个长度为 n 的滑动窗口从这个最短字符串滑过，每次滑动一位，窗口里的内容就是一个个密码。

怎么得到这样的最短字符串呢？

首先可以肯定 “000...000” （共 n 个0）是可能的一个结果，可以从这里开始递归。

每次对于当前密码，`可以去掉最高位，再在最后加一位`来得到下一个密码。为什么用这样的策略？正是上边的滑动窗口提示的，实际上这也是得到结果字符串的过程。

```go
func crackSafe(n int, k int) string {
	total := int(math.Pow(float64(k), float64(n)))
	high := total / k
	seen := make([]bool, total)
	buf := strings.Builder{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if seen[cur] {
			return
		}
		seen[cur] = true
		last := cur % k      // cur 的最后一位，要加入结果
		cur = cur % high * k // cur 删除最高位，最后一位补 0
		for i := 0; i < k; i++ { // 尝试在最低位追加 i
			next := cur + i
			dfs(next)
		}
		// 后序遍历，追加 cur 的最后一位
		buf.WriteByte(byte(last) + '0')
	}
	dfs(0)

    // 一开始的n-1 位 0，因后序遍历的原因，这些 0 也在随后追加
	for i := 1; i < n; i++ { 
		buf.WriteByte('0')
	}
	// 实际应该返回逆序，但是对这个问题，正序、逆序都行
	return buf.String()
}
```

里边取余操作的效果如果不好理解，把 k 看成 10 就很容易明白了。

时间复杂度是 `O(n*k^n)`，空间复杂度是 `O(n^k)`。

### 为什么是“后序”

值得一提的是，一定要是`后序`遍历，如果改成`前序`，结果就不一定对了。

比如对于 `n=2, k=2` 的用例，返回的结果字符串里就没有连续的两个 `1` 。这是为什么？

画个图看下：  
![](https://raw.githubusercontent.com/zrcoder/leetcodeGo/master/solutions/cracking-the-safe/1.png)

实际上，需要从起点 `00` 开始，走完所有的边才行，且必须一笔完成，不能中断再起一笔，这实际上就是欧拉回路问题。

可以对照图仔细考虑为什么前序遍历会导致不一定一笔完成，而后序就可以。这也是欧拉回路问题的一个特色。

一个更简单的例子是这样：
```
  a <- x -> b
       ^    |
        \   v
           c 
```
对于点 `x` ，如果从这个点开始，怎么才能完成这幅一笔画？

先走 `x->a` 这条死胡同，还是先走 `x->b` 这条环路？这里是显而易见的，而后序遍历能保证无论选择了哪条路，都能先把死胡同里的点加入结果，其次才会把环里的点加入。

 
更多关于图的欧拉回路问题，可以看看 [[332] 重新安排行程](../reconstruct-itinerary) 。

实际上上边这个简图，就是在`重新安排行程`题解里“画”的。
