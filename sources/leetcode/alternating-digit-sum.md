<!-- markdownlint-disable -->
<!-- customize-tags:数学 -->

# 2544. 交替数字和

> [题目链接](https://leetcode.cn/problems/alternating-digit-sum/)

给你一个正整数 `n` 。 `n` 中的每一位数字都会按下述规则分配一个符号：

- **最高有效位** 上的数字分配到 **正** 号。
- 剩余每位上数字的符号都与其相邻数字相反。

返回所有数字及其对应符号的和。

**示例 1：**

```
输入：n = 521
输出：4
解释：(+5) + (-2) + (+1) = 4
```

**示例 2：**

```
输入：n = 111
输出：1
解释：(+1) + (-1) + (+1) = 1
```

**示例 3：**

```
输入：n = 886996
输出：0
解释：(+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0
```

**提示：**

- `1 <= n <= 109`

**代码实现：**

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

```go
func alternateDigitSum(n int) int {
    var ans int
    sing := 1
    for n > 0 {
        ans += sing * (n % 10)
        n /= 10
        sing = -sing
    }
    // 若最后 sing 为 1 则说明最高位是负数，所以这里要取反
    return -sing * ans
}
```
