<!-- markdownlint-disable -->
<!-- customize-tags:深度优先搜索,广度优先搜索,并查集,数组,矩阵 -->

# 200. 岛屿数量

> [题目链接](https://leetcode.cn/problems/number-of-islands/)

给你一个由  `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**示例 1：**

```
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
```

**示例 2：**

```
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` 的值为 `'0'` 或 `'1'`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现**：  
深度优先搜索，遍历网格，如果遇到岛屿则将 `ans + 1` 同时移除掉整个岛屿，这样可以确保之后的遍历不会遇到同一个岛屿。

```go
func numIslands(grid [][]byte) int {
    var ans int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                ans++
                dfs(grid, i, j);
            }
        }
    }

    return ans
}

// 将相邻的 '1' 修改为 '0'
func dfs(grid [][]byte, i, j int) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
        return
    }
    if grid[i][j] == '0' {
        return
    }

    grid[i][j] = '0'
    dfs(grid, i - 1, j);
    dfs(grid, i + 1, j);
    dfs(grid, i, j - 1);
    dfs(grid, i, j + 1);
}
```
