<!-- customize-tags:二叉树,深度优先搜索 -->

# 剑指 Offer 55 - II. 平衡二叉树

> [题目链接](https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof)

输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过 1，那么它就是一棵平衡二叉树。

**示例 1:**  
给定二叉树 [3,9,20,null,null,15,7]

```text
    3
   / \
  9  20
    /  \
   15   7
```

返回 true 。

**示例 2:**  
给定二叉树 [1,2,2,3,3,null,null,4,4]

```text
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```

返回  false 。

**限制：**

- 0 <= 树的结点个数 <= 10000

**代码实现：**

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    boolean flag = true;
    public boolean isBalanced(TreeNode root) {
        if (root == null) {
            return true;
        }
        int a = foo(root.left);
        int b = foo(root.right);
        return flag && Math.abs(a - b) < 2;
    }

    public int foo(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int left = foo(root.left) + 1;
        int right = foo(root.right) + 1;
        if (Math.abs(left - right) > 1) {
            flag = false;
        }

        return Math.max(left, right);
    }
}
```
