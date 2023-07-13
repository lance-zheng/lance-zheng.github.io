<!-- customize-tags:二叉树 -->

# 剑指 Offer 54. 二叉搜索树的第 k 大节点

> [题目链接](https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)

给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

**示例 1：**

```text
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
```

**示例 2：**

```text
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
```

**限制：**

```text
1 ≤ k ≤ 二叉搜索树元素个数
```

**解题思路：**  
二叉搜索树：

- 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值。
- 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值。

通过 `右` -> `根` -> `左` 的顺序遍历出来的就是一个递减的序列。

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
    public int kthLargest(TreeNode root, int k) {
        int[] ans = new int[2];
        rec(root, k, ans);
        return ans[1];
    }

    private void rec(TreeNode root, int k, int[] ans) {
        if (root == null) {
            return;
        }
        rec(root.right, k, ans);
        ans[0]++;
        if (ans[0] == k) {
            ans[1] = root.val;
            return;
        }
        rec(root.left, k, ans);

    }
}
```
