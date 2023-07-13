<!-- markdownlint-disable -->
<!-- customize-tags:数学,数论 -->

# 2413. 最小偶倍数

> [题目链接](https://leetcode.cn/problems/smallest-even-multiple/)

给你一个正整数 `n` ，返回 `2` 和 `n` 的最小公倍数（正整数）。

**示例 1：**

```
输入：n = 5
输出：10
解释：5 和 2 的最小公倍数是 10 。
```

**示例 2：**

```
输入：n = 6
输出：6
解释：6 和 2 的最小公倍数是 6 。注意数字会是它自身的倍数。
```

**提示：**

- `1 <= n <= 150`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

```go
func smallestEvenMultiple(n int) int {
    return n << (n & 1)
}
```
