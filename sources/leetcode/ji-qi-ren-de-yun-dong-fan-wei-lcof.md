<!-- customize-tags:双指针 -->

# 面试题 13. 机器人的运动范围

> [题目链接](https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

地上有一个 m 行 n 列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于 k 的格子。例如，当 k 为 18 时，机器人能够进入方格 [35, 37] ，因为 3+5+3+7=18。但它不能进入方格 [35, 38]，因为 3+5+3+8=19。请问该机器人能够到达多少个格子？

**示例 1：**

```text
输入：m = 2, n = 3, k = 1
输出：3
```

**示例 2：**

```text
输入：m = 3, n = 1, k = 0
输出：1
```

**提示：**

- `1 <= n,m <= 100`
- `0 <= k <= 20`

**代码实现：**

```java
class Solution {
    public int movingCount(int m, int n, int k) {
        boolean[][] used = new boolean[m][n];
        return foo(0, 0, m, n, k, used);
    }

    private int foo(int i, int j, int m, int n, int k, boolean[][] used) {
        if (i >= m || j >= n || i < 0 || j < 0) {
            return 0;
        }
        if (used[i][j]) {
            return 0;
        }
        int sum = 0;
        int t = i;
        while (t != 0) {
            sum += (t % 10);
            t = t / 10;
        }
        t = j;
        while (t != 0) {
            sum += (t % 10);
            t = t / 10;
        }
        if (sum > k) {
            return 0;
        }

        used[i][j] = true;
        return 1 +
        foo(i - 1, j, m, n, k, used) +
        foo(i + 1, j, m, n, k, used) +
        foo(i, j - 1, m, n, k, used) +
        foo(i, j + 1, m, n, k, used);
    }
}
```
