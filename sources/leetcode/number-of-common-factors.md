<!-- markdownlint-disable -->
<!-- customize-tags:数学,枚举,数论 -->

# 2427. 公因子的数目

> [题目链接](https://leetcode.cn/problems/number-of-common-factors/)

给你两个正整数 `a` 和 `b` ，返回 `a` 和 `b` 的 **公** 因子的数目。

如果 `x` 可以同时整除 `a` 和 `b` ，则认为 `x` 是 `a` 和 `b` 的一个 **公因子** 。

**示例 1：**

```
输入：a = 12, b = 6
输出：4
解释：12 和 6 的公因子是 1、2、3、6 。
```

**示例 2：**

```
输入：a = 25, b = 30
输出：2
解释：25 和 30 的公因子是 1、5 。
```

**提示：**

- `1 <= a, b <= 1000`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

```go
func commonFactors(a int, b int) int {
    var ans int
    for i:=1; i <= a && i <= b; i++ {
        if a % i == 0 && b % i == 0 {
            ans++
        }
    }

    return ans
}
```
