<!-- markdownlint-disable -->
<!-- customize-tags:栈,数组,链表,单调栈 -->

# 1019. 链表中的下一个更大节点

> [题目链接](https://leetcode.cn/problems/next-greater-node-in-linked-list/)

给定一个长度为  `n`  的链表  `head`

对于列表中的每个节点，查找下一个 **更大节点** 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 **严格大于** 它的值。

返回一个整数数组 `answer` ，其中 `answer[i]` 是第 `i` 个节点( **从 1 开始** )的下一个更大的节点的值。如果第 `i` 个节点没有下一个更大的节点，设置  `answer[i] = 0` 。

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

- 链表中节点数为  `n`
- `1 <= n <= 104`
- `1 <= Node.val <= 109`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

我们从后遍历链表(使用递归)，当我们变量到 `4` 这个元素时后面的这个，后面小于当前元素的值都没有用了，因为当前面小于当前元素时右边最近的就是 `4`，当大于 `4` 时，右侧小于`4`的元素也都不符合条件。

<img width=400 src='/assets/image/1681110326.png'/>

可以使用单调栈解决这个问题，从后遍历链表，如果发现栈顶元素比当前元素小就弹栈，直到寻找到一个大于当前元素的数据，若没有则说明右侧没有大于当前元素的数据 `ans[i] = 0`，若存在则 `ans[i] = 栈顶元素`。最后将当前元素压入栈顶。

```go
func nextLargerNodes(head *ListNode) []int {

    var ans, stack []int

    var dfs func(*ListNode, int)
    dfs = func(node *ListNode, i int) {
        if node == nil {
            ans = make([]int, i)
            return
        }
        dfs(node.Next, i+1)

        if len(stack) > 0 {
            idx := len(stack) - 1
            for ;idx > -1;idx-- {
                if stack[idx] > node.Val {
                    break
                }
            }
            stack = stack[:idx + 1]
        }

        if len(stack) == 0 {
            ans[i] = 0
        } else {
            ans[i] = stack[len(stack) - 1]
        }

        stack = append(stack, node.Val)
    }
    dfs(head, 0)

    return ans
}
```

**优化：**
参考上面的思路其实也可以正向遍历，对于一个数据 `y` 我们可以去寻找它是哪一个数的**更大节点**

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    type pair struct{ K, V int }
    var ans []int
    var stack []pair
    var k int
    for head != nil {
        // 占位
        ans = append(ans, 0)
        n := 0
        for i := len(stack) - 1; i > -1; i-- {
            if head.Val > stack[i].V {
                // 把所有栈中小于当前数的元素都弹出去
                // 并将它的 更大节点 更新为当前节点
                ans[stack[i].K] = head.Val
                n++
            } else {
                break
            }
        }
        stack = append(stack[:len(stack)-n], pair{k, head.Val})
        head = head.Next
        k++
    }

    return ans
}
```
