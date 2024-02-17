<!-- markdownlint-disable -->
<!-- customize-tags:数组,模拟 -->

# 874. 模拟行走机器人

> [题目链接](https://leetcode.cn/problems/walking-robot-simulation/)

机器人在一个无限大小的 XY 网格平面上行走，从点  `(0, 0)` 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 `commands` ：

- `-2` ：向左转  `90` 度
- `-1` ：向右转 `90` 度
- `1 <= x <= 9` ：向前移动  `x`  个单位长度

在网格上有一些格子被视为障碍物  `obstacles` 。第 `i`  个障碍物位于网格点  `obstacles[i] = (xi, yi)` 。

机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。

返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 `5` ，则返回 `25` ）

**注意：**

- 北表示 `+Y` 方向。
- 东表示 `+X` 方向。
- 南表示 `-Y` 方向。
- 西表示 `-X` 方向。

**示例 1：**

```
输入：commands = [4,-1,3], obstacles = []
输出：25
解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 3 个单位，到达 (3, 4)
距离原点最远的是 (3, 4) ，距离为 32 + 42 = 25
```

**示例  2：**

```
输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出：65
解释：机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
4. 左转
5. 向北走 4 个单位，到达 (1, 8)
距离原点最远的是 (1, 8) ，距离为 12 + 82 = 65
```

**提示：**

- `1 <= commands.length <= 104`
- `commands[i]` is one of the values in the list `[-2,-1,1,2,3,4,5,6,7,8,9]`.
- `0 <= obstacles.length <= 104`
- `-3 * 104 <= xi, yi <= 3 * 104`
- 答案保证小于 `231`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**Code：**

```java
class Solution {
    public int robotSim(int[] commands, int[][] obstacles) {
        Map<Integer,Set<Integer>> xy = new HashMap();
        for (int i = 0; i < obstacles.length; i++) {
            int x = obstacles[i][0];
            int y = obstacles[i][1];
            Set<Integer> ySet = xy.getOrDefault(x, new HashSet<>());
            ySet.add(y);
            xy.put(x, ySet);
        }

        int answer = -1;
        // 0:北 1:西 2:南 3:东
        int[] state = new int[]{0,0,0};
        int[] dx = new int[]{0,-1,0,1};
        int[] dy = new int[]{1,0,-1,0};
        for (int i = 0; i < commands.length; i++) {
            int c = commands[i];
            if (c == -1) {
                state[2] = (state[2] + 3) % 4;
            } else if (c == -2) {
                state[2] = (state[2] + 1) % 4;
            } else {
                int a = dx[state[2]];
                int b = dy[state[2]];
                while (c > 0) {
                    int nextX = state[0] + a;
                    int nextY = state[1] + b;

                    if (xy.getOrDefault(nextX, Collections.emptySet()).contains(nextY)) {
                        break;
                    }

                    state[0] = nextX;
                    state[1] = nextY;
                    c--;
                }
            }

            answer = Math.max(state[0]*state[0] + state[1]*state[1], answer);
        }

        return answer;
    }
}
```
