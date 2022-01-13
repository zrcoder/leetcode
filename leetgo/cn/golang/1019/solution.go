package solution

/*
## [1019. Next Greater Node In Linked List](https://leetcode.cn/problems/next-greater-node-in-linked-list) (Medium)

给定一个长度为 `n` 的链表 `head`

对于列表中的每个节点，查找下一个 **更大节点** 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 **严格大于** 它的值。

返回一个整数数组 `answer` ，其中 `answer[i]` 是第 `i` 个节点( **从1开始** )的下一个更大的节点的值。如果第 `i` 个节点没有下一个更大的节点，设置 `answer[i] = 0` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext1.jpg)

```
输入：head = [2,1,5]
输出：[5,5,0]

```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext2.jpg)

```
输入：head = [2,7,4,3,5]
输出：[7,0,5,5,0]

```

**提示：**

- 链表中节点数为 `n`
- `1 <= n <= 10⁴`
- `1 <= Node.val <= 10⁹`


*/

// [start] don't modify
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// “下一个最大”是一个可用单调栈解决的经典问题，只需要把链表先化成数组
func nextLargerNodes(head *ListNode) []int {
	vals := []int{}
	for p := head; p != nil; p = p.Next {
		vals = append(vals, p.Val)
	}
	ans := make([]int, len(vals))
	stack := make([]int, 0, len(vals))
	for i, v := range vals {
		for len(stack) > 0 && vals[stack[len(stack)-1]] < v {
			ans[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 小优化， 也可以省掉 vals 数组，直接遍历链表的同时维护单调栈和答案，有不少细节要注意。
/*
func nextLargerNodes(head *ListNode) []int {
	type pair struct {
        i, v int
    }
    i := 0
    result := make([]int, 0)
	stack := []pair{}
    for p := head; p != nil; p = p.Next {
        v := p.Val
        result = append(result, 0)
        for len(stack) > 0 && stack[len(stack)-1].v < v {
            result[stack[len(stack)-1].i] = v
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, pair{i, v})
        i++
    }
	return result
}
*/
// 扩展：”下一个更小“ 问题怎么解？ 仅需改动单调栈出栈时的比较逻辑
// [end] don't modify
