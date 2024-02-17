<!-- customize-tags:二叉树,深度优先搜索 -->

# 剑指 Offer 55 - I. 二叉树的深度

> [题目链接](https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof)

输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

**例如：**  
 给定二叉树 [3,9,20,null,null,15,7]，

```text
    3
   / \
  9  20
    /  \
   15   7
```

返回它的最大深度  3 。

**提示：**

- 节点总数 <= 10000

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
    public int maxDepth(TreeNode root) {
        return root == null ? 0 : Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
    }
}
```
