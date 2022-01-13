package solution

/*
## [1138. Alphabet Board Path](https://leetcode.cn/problems/alphabet-board-path) (Medium)

我们从一块字母板上的位置 `(0, 0)` 出发，该坐标对应的字符为 `board[0][0]`。

在本题里，字母板为 `board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]`，如下所示。

![](https://assets.leetcode.com/uploads/2019/07/28/azboard.png)

我们可以按下面的指令规则行动：

- 如果方格存在， `'U'` 意味着将我们的位置上移一行；
- 如果方格存在， `'D'` 意味着将我们的位置下移一行；
- 如果方格存在， `'L'` 意味着将我们的位置左移一列；
- 如果方格存在， `'R'` 意味着将我们的位置右移一列；
- `'!'` 会把在我们当前位置 `(r, c)` 的字符 `board[r][c]` 添加到答案中。

（注意，字母板上只存在有字母的位置。）

返回指令序列，用最小的行动次数让答案和目标 `target` 相同。你可以返回任何达成目标的路径。

**示例 1：**

```
输入：target = "leet"
输出："DDR!UURRR!!DDD!"

```

**示例 2：**

```
输入：target = "code"
输出："RR!DDRR!UUL!R!"

```

**提示：**

- `1 <= target.length <= 100`
- `target` 仅含有小写英文字母。


*/

// [start] don't modify
func alphabetBoardPath(target string) string {
    x, y := 0, 0
    buf := strings.Builder{}
    for _, t := range target {
        tx, ty := int(t-'a')/5, int(t-'a')%5
        yDiff := ty-y
        xDiff := tx-x
        // 因字母z位置的特殊性，先向上向左，后向下向右
        if xDiff < 0 {
            buf.WriteString(strings.Repeat("U", -xDiff))
        }
        if yDiff < 0 {
            buf.WriteString(strings.Repeat("L", -yDiff))
        }
        if xDiff > 0 {
            buf.WriteString(strings.Repeat("D", xDiff))
        }
        if yDiff > 0 {
            buf.WriteString(strings.Repeat("R", yDiff))
        }
        buf.WriteByte('!')
        x, y = tx, ty
    }
    return buf.String()
}
// [end] don't modify
