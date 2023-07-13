<!-- customize-tags:递归 -->

# 剑指 Offer 64. 求 1+2+…+n

> [题目链接](https://leetcode.cn/problems/qiu-12n-lcof/)

求 `1+2+...+n` ，要求不能使用乘除法、for、while、if、else、switch、case 等关键字及条件判断语句（A?B:C）。

**示例 1：**

```text
输入: n = 3
输出: 6
```

**示例 2：**

```text
输入: n = 9
输出: 45
```

**限制：**

- 1 <= n <= 10000

**代码实现：**

```java
class Solution {
    public int sumNums(int n) {
        int sum = n;
        // 利用逻辑运算符的短路效果作为函数递归的出口
        boolean t = n > 0 && (sum += sumNums(n - 1)) > 0;
        return sum;
    }
}
```
