---
title: "面试题 04.09. 二叉搜索树序列"
date: 2022-12-14T14:08:35+08:00
---

#### [面试题 04.09. 二叉搜索树序列](https://leetcode.cn/problems/bst-sequences-lcci/)

难度困难

从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。

给定一个由**不同节点**组成的二叉搜索树 `root`，输出所有可能生成此树的数组。

**示例 1:**

**输入:** root = [2,1,3]
**输出:** [[2,1,3],[2,3,1]]
解释: 数组 [2,1,3]、[2,3,1] 均可以通过从左向右遍历元素插入树中形成以下二叉搜索树
   2
  / \
  1 3

**示例** **2:**

**输入:** root = [4,1,null,null,3,2]
**输出:** [[4,1,3,2]]

**提示：**

- 二叉搜索树中的节点数在 `[0, 1000]` 的范围内
- `1 <= 节点值 <= 10^6`
- 用例保证符合要求的数组数量不超过 `5000`

函数签名：

```go
func BSTSequences(root *TreeNode) [][]int
```

## 分析

### 经典回溯

画一画尝试，会发现实际就是输出树的所有遍历方式。

用一个集合来存储当前待遍历的节点

1. 任意选取一个节点
  
2. 该节点加入遍历路径
  
3. 把它的子节点加入集合
  
4. 递归直到集合为空，将路径加入结果
  
5. 回溯，尝试选取其他节点
  
集合用队列，保证当前节点在队列前部。可以用切片模拟队列，回溯时，之前删除的队首元素不用重新放回队首，只需要放到队尾即可，顺序不重要，这样实际只需要在队首删除元素，队尾增加或删除元素。

```go
func BSTSequences(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{{}}
    }

    res := [][]int{}
    nodes := []*TreeNode{root}
    var path []int

    var backtrack func()
    backtrack = func() {
        if len(nodes) == 0 {
            res = append(res, append([]int{}, path...))
            return
        }

        // 对当前size个节点，任意选一个加入path
        for size := len(nodes); size > 0; size-- {
            // 选第一个节点，并从nodes中删除
            cur := nodes[0]
            nodes = nodes[1:]
            path = append(path, cur.Val)
            if cur.Left != nil {
                nodes = append(nodes, cur.Left)
            }
            if cur.Right != nil {
                nodes = append(nodes, cur.Right)
            }

            backtrack()

            // 回溯，不一定（也不必要）保持节点原来的顺序
            // 删掉原来新增的左右节点
            if cur.Left != nil {
                nodes = nodes[:len(nodes)-1]
            }
            if cur.Right != nil {
                nodes = nodes[:len(nodes)-1]
            }
            // 原来的第一个节点加回来，现在在nodes数组最后了，这里无所谓顺序
            nodes = append(nodes, cur) 
            path = path[:len(path) - 1]
        }
    }

    backtrack()

    return res
}
```

复杂度，不太好分析，因题目保证结果数组数量不超过5000，复杂度不会超过这个数量级。
