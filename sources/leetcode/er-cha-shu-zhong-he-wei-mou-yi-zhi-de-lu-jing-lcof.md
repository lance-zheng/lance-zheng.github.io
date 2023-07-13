<!-- customize-tags:二叉树 -->

# 剑指 Offer 34. 二叉树中和为某一值的路径

> [题目链接](https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof)

给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。  
叶子节点 是指没有子节点的节点。

**示例 1：**  
![img](/assets/image/pathsumii1.jpg)

```text
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
```

**示例 2：**  
![img](/assets/image/pathsum2.jpg)

```text
输入：root = [1,2,3], targetSum = 5
输出：[]
```

**示例 3：**

```text
输入：root = [1,2], targetSum = 0
输出：[]
```

**提示：**

```text
树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
```

**解题思路：**

```text
深度优先搜索，同时记录走过的路径和路径和，在达到末尾时检查路径和是否和目标值相等。
```

**代码实现：**

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int target) {
        List<List<Integer>> ans = new ArrayList();
        rec(root, target, ans, new ArrayList(), 0);
        return ans;
    }

    private void rec(TreeNode root, int target, List<List<Integer>> ans, List<Integer> path, int curSum) {
        if (root == null) {
            return;
        }
        path.add(root.val);
        curSum += root.val;
        if (root.left == null && root.right == null && target == curSum) {
            ans.add(new ArrayList(path));
        } else {
            rec(root.left, target, ans, path, curSum);
            rec(root.right, target, ans, path, curSum);
        }
        path.remove(path.size() - 1);
    }
}
```
