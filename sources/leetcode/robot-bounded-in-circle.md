<!-- markdownlint-disable -->
<!-- customize-tags:数学,字符串,模拟 -->

# 1041. 困于环中的机器人

> [题目链接](https://leetcode.cn/problems/robot-bounded-in-circle/)

在无限的平面上，机器人最初位于  `(0, 0)`  处，面朝北方。注意:

- **北方向** 是 y 轴的正方向。
- **南方向** 是 y 轴的负方向。
- **东方向** 是 x 轴的正方向。
- **西方向** 是 x 轴的负方向。

机器人可以接受下列三条指令之一：

- `"G"`：直走 1 个单位
- `"L"`：左转 90 度
- `"R"`：右转 90 度

机器人按顺序执行指令  `instructions`，并一直重复它们。

只有在平面中存在环使得机器人永远无法离开时，返回  `true`。否则，返回 `false`。

**示例 1：**

```
输入：instructions = "GGLLGG"
输出：true
解释：机器人最初在(0,0)处，面向北方。
“G”:移动一步。位置:(0,1)方向:北。
“G”:移动一步。位置:(0,2).方向:北。
“L”:逆时针旋转90度。位置:(0,2).方向:西。
“L”:逆时针旋转90度。位置:(0,2)方向:南。
“G”:移动一步。位置:(0,1)方向:南。
“G”:移动一步。位置:(0,0)方向:南。
重复指令，机器人进入循环:(0,0)——>(0,1)——>(0,2)——>(0,1)——>(0,0)。
在此基础上，我们返回true。
```

**示例 2：**

```
输入：instructions = "GG"
输出：false
解释：机器人最初在(0,0)处，面向北方。
“G”:移动一步。位置:(0,1)方向:北。
“G”:移动一步。位置:(0,2).方向:北。
重复这些指示，继续朝北前进，不会进入循环。
在此基础上，返回false。
```

**示例 3：**

```
输入：instructions = "GL"
输出：true
解释：机器人最初在(0,0)处，面向北方。
“G”:移动一步。位置:(0,1)方向:北。
“L”:逆时针旋转90度。位置:(0,1).方向:西。
“G”:移动一步。位置:(- 1,1)方向:西。
“L”:逆时针旋转90度。位置:(- 1,1)方向:南。
“G”:移动一步。位置:(- 1,0)方向:南。
“L”:逆时针旋转90度。位置:(- 1,0)方向:东方。
“G”:移动一步。位置:(0,0)方向:东方。
“L”:逆时针旋转90度。位置:(0,0)方向:北。
重复指令，机器人进入循环:(0,0)——>(0,1)——>(- 1,1)——>(- 1,0)——>(0,0)。
在此基础上，我们返回true。
```

**提示：**

- `1 <= instructions.length <= 100`
- `instructions[i]`  仅包含  `'G', 'L', 'R'`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**  
如果存在环的话每次执行完指令后方向都必须发生改变，最多执行 4 次指令，方向就会变成初始的方向；若执行第 4 次时或者次数更少，还没有回到起点就说明不存在环。

```go
func isRobotBounded(instructions string) bool {
    xy := [2]int{0, 0}
    // 北:0 东:1 南:2 西:3
    var fangxiang int
    offset := [4][2]int{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
    }
    for i := 0; i < 4; i++ {
        startFangxiang := fangxiang
        for _, c := range instructions {
            switch c {
            case 'G':
                xOffset := offset[fangxiang][0]
                yOffset := offset[fangxiang][1]
                xy[0] += xOffset
                xy[1] += yOffset
            case 'L':
                fangxiang--
                if fangxiang < 0 {
                    fangxiang = 3
                }
            case 'R':
                fangxiang++
                if fangxiang > 3 {
                    fangxiang = 0
                }
            }
        }
        if xy[0] == 0 && xy[1] == 0 && fangxiang == 0 { // 回到最初起点
            return true
        }
        if startFangxiang == fangxiang { // 执行一次指令前后 朝向 一致
            return false
        }
    }
    return false
}
```
