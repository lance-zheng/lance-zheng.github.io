<!-- customize-tags:动态规划 -->

# 1139. 最大的以 1 为边界的正方形

> [题目链接](https://leetcode.cn/problems/largest-1-bordered-square)

给你一个由若干 `0` 和 `1` 组成的二维网格  grid，请你找出边界全部由 1 组成的最大 `正方形` 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0。

**示例 1：**

```text
输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
输出：9
```

**示例 2：**

```text
输入：grid = [[1,1,0,0]]
输出：1
```

**提示：**

```text
1 <= grid.length <= 100
1 <= grid[0].length <= 100
grid[i][j] 为 0 或 1
```

**解题思路：**  
先创建一个 3 维数组

- `dp[i][j][0]` 表示 `grid[i][j]` 上方连续 1 的个数(包括 `grid[i][j]` )
- `dp[i][j][1]` 表示 `grid[i][j]` 左方连续 1 的个数(包括 `grid[i][j]` )
  <img width = "400" height = "400" alt="" align="center" src="/assets/image/1630649803-BqicKf-image.png">

然后从数组右下角开始遍历，以当前元素为正方形的右下角顶点。  
因为从 dp 数组中我们已知每个元素上方和左方分别有多少个连续 1。  
我们只需要从左方和上方选取一个最小的边作为正方形的边长，然后验证能否构成正方形，如果不能则将边缩小一个单位继续重复验证，直到能够形成正方形。

**代码实现：**

```java
class Solution {
    public int largest1BorderedSquare(int[][] grid) {
        int h = grid.length, w = grid[0].length;
        int[][][] dp = new int[h + 1][w + 1][2];

        for (int i = 1; i <= h; i++) {
            for (int j = 1; j <= w; ++j) {
                if (grid[i - 1][j - 1] == 1) {
                    dp[i][j][0] = dp[i - 1][j][0] + 1;
                    dp[i][j][1] = dp[i][j - 1][1] + 1;
                }
            }
        }

        int maxWidth = 0;
        for (int i = h; i > 0; i--) {
            if (i <= maxWidth) {
                break;
            }
            for (int j = w; j > 0; j--) {
                if (j <= maxWidth) {
                    break;
                }

                int width = Math.min(dp[i][j][0], dp[i][j][1]);
                while (width > maxWidth) {
                    if (dp[i - width + 1][j][1] < width || dp[i][j - width + 1][0] < width) {
                        width--;
                        continue;
                    }
                    maxWidth = Math.max(width, maxWidth);
                }
            }
        }


        return maxWidth * maxWidth;
    }
}
```
