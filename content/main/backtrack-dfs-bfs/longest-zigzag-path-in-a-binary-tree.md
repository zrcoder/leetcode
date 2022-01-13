---
title: "1372. 二叉树中的最长交错路径"
date: 2023-05-25T21:44:45+08:00
---

##  [1372. 二叉树中的最长交错路径](https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/)

难度中等

给你一棵以 `root` 为根的二叉树，二叉树中的交错路径定义如下：

- 选择二叉树中 **任意** 节点和一个方向（左或者右）。
- 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
- 改变前进方向：左变右或者右变左。
- 重复第二步和第三步，直到你在树中无法继续移动。

交错路径的长度定义为：**访问过的节点数目 - 1**（单个节点的路径长度为 0 ）。

请你返回给定树中最长 **交错路径** 的长度。

**示例 1：**

****

**输入：** root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
**输出：** 3
**解释：** 蓝色节点为树中最长交错路径（右 -> 左 -> 右）。

**示例 2：**

****

**输入：** root = [1,1,1,null,1,null,null,1,1,null,1]
**输出：** 4
**解释：** 蓝色节点为树中最长交错路径（左 -> 右 -> 左 -> 右）。

**示例 3：**

**输入：** root = [1]
**输出：** 0

**提示：**

- 每棵树最多有 `50000` 个节点。
- 每个节点的值在 `[1, 100]` 之间。

函数签名：

```go
func longestZigZag(root *TreeNode) int
```

## 分析

BFS 或 DFS 都可以，DFS 较简洁。

### DFS

直观解法：

{{<tabs groupid="1">}}

{{%tab name="Go"%}}

```go
func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var dfs func(*TreeNode, bool, int) int
    dfs = func(root *TreeNode, isLeft bool, count int) int {
        if root == nil {
            return count
        }
        if isLeft {
            return max(dfs(root.Left, true, 0), dfs(root.Right, false, count+1))
        }
        return max(dfs(root.Left, true, count+1), dfs(root.Right, false, 0))
    }
    return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}
```

{{%/tab%}}

{{%tab name="Python3"%}}

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        def dfs(root, isleft, count):
            if not root:
                return count
            if isleft:
                return max(dfs(root.left, True, 0), dfs(root.right, False, count+1))
            return max(dfs(root.left, True, count+1), dfs(root.right, False, 0))
        
        return max(dfs(root.left, True, 0), dfs(root.right, False, 0))
```

{{%/tab%}}

{{%/tabs%}}

实际上，在 dfs 时，递归函数设计为多个入参，无论代码可读性还是运行性能，都是劣于设计为多个返回值的。可以参考专题《在树上动态规划》。

{{<tabs groupid="2">}}
{{%tab name="Go"%}}

```go
func longestZigZag(root *TreeNode) int {
    res := 0

    var dfs func(*TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        if root == nil {
            return -1, -1
        }

        _, lr := dfs(root.Left)
        rl, _ := dfs(root.Right)

        res = max(res, lr+1, rl+1)
        return lr+1, rl+1
    }
    
    dfs(root)
    return res
}
```

{{%/tab%}}

{{%tab name="Python3"%}}

```python
class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        self.res = 0

        def dfs(root):
            if not root:
                return (-1, -1)
            
            _, lr = dfs(root.left)
            rl, _ = dfs(root.right)

            self.res = max(self.res, lr+1, rl+1)
            return (lr+1, rl+1)
        
        dfs(root)
        return self.res
```

{{%/tab%}}

{{</tabs>}}
