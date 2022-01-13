package solution

/*
## [652. Find Duplicate Subtrees](https://leetcode.cn/problems/find-duplicate-subtrees) (Medium)

给你一棵二叉树的根节点 `root` ，返回所有 **重复的子树**。

对于同一类的重复子树，你只需要返回其中任意 **一棵** 的根结点即可。

如果两棵树具有 **相同的结构** 和 **相同的结点值**，则认为二者是 **重复** 的。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/08/16/e1.jpg)

```
输入：root = [1,2,3,4,null,2,4,null,null,4]
输出：[[2,4],[4]]
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/08/16/e2.jpg)

```
输入：root = [2,1,1]
输出：[[1]]
```

**示例 3：**

**![](https://assets.leetcode.com/uploads/2020/08/16/e33.jpg)**

```
输入：root = [2,2,2,3,null,3,null]
输出：[[2,3],[3]]
```

**提示：**

- 树中的结点数在 `[1, 5000]` 范围内。
- `-200 <= Node.val <= 200`


*/

// [start] don't modify
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution 1, O(n^2)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    cnt := map[string]int{}
    var dfs func(*TreeNode) string
    var res []*TreeNode
    dfs = func(tn *TreeNode) string {
        if tn == nil {return "#"}
        s := strconv.Itoa(tn.Val)+","+dfs(tn.Left)+","+dfs(tn.Right)
        cnt[s]++
        if cnt[s] == 2 {
            res = append(res, tn)
        }
        return s
    }
    dfs(root)
    return res
}
// solution 2, O(n)
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    cnt := map[string]int{}
    idx := map[string]int{}
    var dfs func(*TreeNode) int
    var res []*TreeNode
    curId := 0
    dfs = func(tn *TreeNode) int {
        if tn == nil {return 0}
        leftId := dfs(tn.Left)
        rightId := dfs(tn.Right)
        s := strconv.Itoa(tn.Val)+","+strconv.Itoa(leftId)+","+strconv.Itoa(rightId)
        cnt[s]++
        if cnt[s] == 2 {
            res = append(res, tn)
        }
        _, ok := idx[s]
        if !ok {
            curId++
            idx[s] = curId
        }
        return  idx[s]
    }
    dfs(root)
    return res
}
// [end] don't modify
