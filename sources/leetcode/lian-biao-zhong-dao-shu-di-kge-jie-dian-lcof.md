<!-- customize-tags:双指针 -->

# 剑指 Offer 22. 链表中倒数第 k 个节点

> [题目链接](https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

输入一个链表，输出该链表中倒数第 k 个节点。为了符合大多数人的习惯，本题从 1 开始计数，即链表的尾节点是倒数第 1 个节点。  
例如，一个链表有 `6` 个节点，从头节点开始，它们的值依次是 `1、2、3、4、5、6`。这个链表的倒数第 `3` 个节点是值为 `4` 的节点。

**示例：**

```text
给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
```

**方法一：**

```markdown
# 双指针

定义两个指针 p1 p2, 先让 p1 p2 节点间隔 k 个节点。
然后两个指针同步向后移动, 当 p2 到尾部时则 p1 刚好是倒数第 k 个节点。
```

**代码实现：**

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode getKthFromEnd(ListNode head, int k) {
        ListNode p1, p2;
        p1 = p2 = head;

        while (p2 != null) {
            p2 = p2.next;
            if (k > 0) {
                --k;
            } else {
                p1 = p1.next;
            }
        }

        return p1;
    }
}
```
