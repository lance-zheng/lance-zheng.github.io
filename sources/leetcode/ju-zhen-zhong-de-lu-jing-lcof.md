<!-- customize-tags:回溯 -->

# 剑指 Offer 12. 矩阵中的路径

> [题目链接](https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/)

给定一个  `m x n` 二维字符网格  `board` 和一个字符串单词  `word` 。如果  `word` 存在于网格中，返回 `true` ；否则，返回 `false` 。  
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 `ABCCED`（单词中的字母已标出）。

![img](/assets/image/word2.jpg)

**示例 1：**

```text
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
```

**示例 2：**

```text
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
```

**提示：**

```text
m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
```

**方法一：回溯算法**  
通过递归的方式向分别上下左右四个方向查询是否符合条件，同时将已经使用过的元素标记为已使用的状态。

**代码实现：**

```java
class Solution {
    public boolean exist(char[][] board, String word) {

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (rec(board, i, j, word, 0)) {
                    return true;
                }
            }
        }
        return false;
    }

    private boolean rec(char[][] board, int i, int j, String word, int k) {
        if (k >= word.length()) {
            return true;
        }
        if (i >= board.length || i < 0 || j >= board[i].length || j < 0) {
            return false;
        }
        if (board[i][j] != word.charAt(k)) {
            return false;
        }

        board[i][j] *= -1; // 标记为已被使用
        boolean res =  rec(board, i + 1, j, word, k + 1) ||
        rec(board, i - 1, j, word, k + 1) ||
        rec(board, i, j + 1, word, k + 1) ||
        rec(board, i, j - 1, word, k + 1);
        board[i][j] *= -1;

        return res;
    }
}
```
