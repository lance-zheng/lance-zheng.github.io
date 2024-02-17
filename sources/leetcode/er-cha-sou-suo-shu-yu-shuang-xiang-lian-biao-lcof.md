<!---
categories:二叉树
-->
<!-- customize-tags:数组 -->

# 剑指 Offer 36. 二叉搜索树与双向链表

> [题目链接](https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof)

输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

为了让您更好地理解问题，以下面的二叉搜索树为例：
![img](/assets/image/bstdlloriginalbst.png)

我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
![img](/assets/image/bstdllreturndll.png)  
特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

**解题思路：**  
二叉搜索树：

- 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值。
- 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值。

二叉搜索树的中序遍历是一个递增的序列，因此只需要进行中序遍历二叉树，同时将当前节点的值与前一个节点相互指向即可。
遍历过程需要两个引用来记录第一个节点和前一个节点。

**代码实现：**

```java
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
    Node first;
    Node pre;
    public Node treeToDoublyList(Node root) {
        if (root == null) {
            return null;
        }
        rec(root);
        first.left = pre;
        pre.right = first;

        return first;
    }

    private void rec(Node root) {
        if (root == null) {
            return;
        }
        rec(root.left);

        if (first == null) {
            first = root;
            pre = first;
        } else {
            pre.right = root;
            root.left = pre;
            pre = root;
        }

        rec(root.right);
    }
}
```
