---
title: "721. 账户合并"
date: 2021-04-19T22:04:56+08:00
weight: 1
---

## [721. 账户合并](https://leetcode-cn.com/problems/accounts-merge/)

难度中等

给定一个列表 `accounts`，每个元素 `accounts[i]` 是一个字符串列表，其中第一个元素 `accounts[i][0]` 是 *名称 (name)*，其余元素是 *emails* 表示该账户的邮箱地址。

现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。账户本身可以以任意顺序返回。

 

**示例 1：**

```
输入：
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
输出：
[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
解释：
第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。 
第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。
```

 

**提示：**

- `accounts`的长度将在`[1，1000]`的范围内。
- `accounts[i]`的长度将在`[1，10]`的范围内。
- `accounts[i][j]`的长度将在`[1，30]`的范围内。

函数签名：

```go
func accountsMerge(accounts [][]string) [][]string
```

## 分析

每一行输入数据里的几个邮箱必定属于同一个人，不同行同一个邮箱一定属于同一个个人。这样从邮箱的维度统计就行。使用并查集比较方便。

```go
func accountsMerge(accounts [][]string) [][]string {
	emailId, emailPeople := wrapEmail(accounts)

	uf = make([]int, len(emailId))
	for i := range uf {
		uf[i] = i
	}

	// 合并邮件
	for _, v := range accounts {
		if len(v) < 2 {
			continue
		}
		x := emailId[v[1]]
		for _, email := range v[2:] {
			y := emailId[email]
			union(y, x)
		}
	}
	// 构造结果
	res := make([][]string, len(emailId))
	for email, id := range emailId {
		id = find(id)
		people := emailPeople[email]
		if len(res[id]) == 0 {
			res[id] = append(res[id], people)
		}
		res[id] = append(res[id], email)
	}
	return removeAndSort(res)
}

var uf []int

func find(x int) int {
	for x != uf[x] {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
func union(x, y int) {
	x, y = find(x), find(y)
	uf[x] = y
}

func wrapEmail(accounts [][]string) (map[string]int, map[string]string) {
	// 给邮箱编号，方便构建并查集
	emailId := make(map[string]int, 0)
	// 需要能迅速获知每个邮箱对应的人
	emailPeople := make(map[string]string, 0)
	for _, v := range accounts {
		for _, email := range v[1:] {
			if _, ok := emailId[email]; !ok {
				emailId[email] = len(emailId)
				emailPeople[email] = v[0]
			}
		}
	}
	return emailId, emailPeople
}

func removeAndSort(s [][]string) [][]string {
	res := make([][]string, 0, len(s))
	for _, v := range s {
		if len(v) > 0 {
			sort.Strings(v)
			res = append(res, v)
		}
	}
	return res
}
```

时间复杂度 `O(nlogn)`，空间复杂度 `O(n)`。其中 `n` 指邮箱数量。

