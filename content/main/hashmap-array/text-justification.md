---
title: "Text Justification"
date: 2023-01-05T10:48:25+08:00
---
## [68. Text Justification](https://leetcode.cn/problems/text-justification) (Hard)

给定一个单词数组 `words` 和一个长度 `maxWidth` ，重新排版单词，使其成为每行恰好有 `maxWidth` 个字符，且左右两端对齐的文本。

你应该使用 “ **贪心算法**” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 `' '` 填充，使得每行恰好有 _maxWidth_ 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入 **额外的** 空格。

**注意:**

- 单词是指由非空格字符组成的字符序列。
- 每个单词的长度大于 0，小于等于 _maxWidth_。
- 输入单词数组 `words` 至少包含一个单词。

**示例 1:**

```
输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

```

**示例 2:**

```
输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
输出:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
     因为最后一行应为左对齐，而不是左右两端对齐。
     第二行同样为左对齐，这是因为这行只包含一个单词。

```

**示例 3:**

```
输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
输出:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]

```

**提示:**

- `1 <= words.length <= 300`
- `1 <= words[i].length <= 20`
- `words[i]` 由小写英文字母和符号组成
- `1 <= maxWidth <= 100`
- `words[i].length <= maxWidth

函数签名：

```go
func fullJustify(words []string, maxWidth int) []string
```

## 分析
直接模拟就可以，不过编码细节较多，我们一步步来。

首先写出框架。需要用类似滑动窗口的双指针方式来遍历单词，窗口中维护结果的每一行。

```go
func fullJustify(words []string, maxWidth int) []string {
	lines := make([]string, 0)
	left := 0
	charsWidth := 0
	for right, word := range words {
		charsWidth += len(word)
		spaceWidth := right - left
		if charsWidth+spaceWidth > maxWidth {
            // words[left:right] 即新的一行，需要两端对齐
			lines = append(lines, justify(words[left:right], charsWidth-len(word), maxWidth))
			charsWidth, left = len(word), right
		}
	}
    // 处理最后一行，左对齐，比较简单
	last := strings.Join(words[left:], " ")
	lines = append(lines, last+strings.Repeat(" ", maxWidth-len(last)))
	return lines
}
```

接下来需要攻克那个用于两端对齐的 `justify` 函数：

```go
// justify 对于给定的几个单词，需要按照两端对齐的方式拼成总长度为 maxWidth 的句子
func justify(words []string, charsWidth, maxWidth int) string
```

仅一个或两个单词的情况比较简单，对于多于两个单词的情况，需要尽可能平均地插入空格。

需要插入的空格总数为：`maxWidth-charsWidth`, 分成 `n-1` 部分，其中n是 words 长度。

如果能整除，也比较简单，不能整除的话怎么办？假设余数为mod，需要给前边的mod 个分割部分各多加一个空格。

```go
func justify(words []string, charsWidth, maxWidth int) string {
	n := len(words)
	if n == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-charsWidth)
	}
	if n == 2 {
		return strings.Join(words, strings.Repeat(" ", maxWidth-charsWidth))
	}
	spaceWidth := (maxWidth - charsWidth) / (n - 1)
	mod := (maxWidth - charsWidth) % (n - 1)
	sep := strings.Repeat(" ", spaceWidth)
	if mod == 0 {
		return strings.Join(words, sep)
	}
	preSep := sep + " "
	return strings.Join(words[:mod+1], preSep) + sep + strings.Join(words[mod+1:], sep)
}
```

代码还可以精简。 `n == 2` 和 `mod == 0` 的判断可以直接删除.

```go
func justify(words []string, charsWidth, maxWidth int) string {
	n := len(words)
	if n == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-charsWidth)
	}
	spaceWidth := (maxWidth - charsWidth) / (n - 1)
	mod := (maxWidth - charsWidth) % (n - 1)
	sep := strings.Repeat(" ", spaceWidth)
	preSep := sep + " "
	return strings.Join(words[:mod+1], preSep) + sep + strings.Join(words[mod+1:], sep)
}
```
