---
title: 93. 复原 IP 地址
---

## [93. 复原 IP 地址](https://leetcode.cn/problems/restore-ip-addresses) (Medium)

**有效 IP 地址** 正好由四个整数（每个整数位于 `0` 到 `255` 之间组成，且不能含有前导 `0`），整数之间用 `'.'` 分隔。

- 例如： `"0.1.2.201"` 和 ` "192.168.1.1"` 是 **有效** IP 地址，但是 `"0.011.255.245"`、 `"192.168.1.312"` 和 `"192.168@1.1"` 是 **无效** IP 地址。

给定一个只包含数字的字符串 `s` ，用以表示一个 IP 地址，返回所有可能的 **有效 IP 地址**，这些地址可以通过在 `s` 中插入 `'.'` 来形成。你 **不能** 重新排序或删除 `s` 中的任何数字。你可以按 **任何** 顺序返回答案。

**示例 1：**

```
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

```

**示例 2：**

```
输入：s = "0000"
输出：["0.0.0.0"]

```

**示例 3：**

```
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

```

**提示：**

- `1 <= s.length <= 20`
- `s` 仅由数字组成

## 分析

可以用回溯，时间复杂度是指数级。

枚举 3 个分割位置即可，三重循环搞定。时间复杂度是 O(n^3)

```go
func restoreIpAddresses(s string) []string {
	var res []string
	check := func(i, j int) bool {
		if j > i+1 && s[i] == '0' {
			return false
		}
		v, _ := strconv.Atoi(s[i:j])
		return v >= 0 && v <= 255
	}
	// 枚举分割点 i，j，k
	n := len(s)
	for i := 1; i < 4 && i < n; i++ {
		if !check(0, i) {
			continue
		}
		for j := i + 1; j < i+4 && j < n; j++ {
			if !check(i, j) {
				continue
			}
			for k := j + 1; k < j+4 && k < n; k++ {
				if !check(j, k) || !check(k, n) {
					continue
				}
				arr := []string{s[:i], s[i:j], s[j:k], s[k:]}
				ip := strings.Join(arr, ".")
				res = append(res, ip)
			}
		}
	}
	return res
}

```
