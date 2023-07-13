<!-- customize-tags:字符串 -->

# 面试题 45. 把数组排成最小的数

> [题目链接](https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)

输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

**示例 1：**

```text
输入: [10,2]
输出: "102"
```

**示例 2：**

```text
输入: [3,30,34,5,9]
输出: "3033459"
```

**提示：**

```text
0 < nums.length <= 100
```

**说明：**

- 输出结果可能非常大，所以你需要返回一个字符串而不是整数
- 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

**代码实现：**

```java
class Solution {
    public String minNumber(int[] nums) {
        List<String> ans = new ArrayList();
        for (Integer n:nums) {
            ans.add(n.toString());
        }
        // 这里就是比较字符串 s1 + s2 大还是 s2 + s1大，我们需要选择小的那个在前面
        ans.sort(((s1, s2) -> (s1 + s2).compareTo(s2 + s1)));
        StringBuilder sb = new StringBuilder();
        for (String str:ans) {
            sb.append(str);
        }

        return sb.toString();
    }
}
```
