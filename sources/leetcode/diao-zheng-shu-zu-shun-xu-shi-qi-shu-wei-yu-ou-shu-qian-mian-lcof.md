<!-- customize-tags:双指针 -->

# 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

> [题目链接](https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof)

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

**示例：**

```text
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
```

**提示：**

```text
0 <= nums.length <= 50000
0 <= nums[i] <= 10000
```

**代码实现：**

```java
class Solution {
    public int[] exchange(int[] nums) {
        int p1 = 0;
        int p2 = nums.length - 1;

        while (p1 < p2) {
            if (nums[p1] % 2 == 0) {
                // swap
                nums[p1] = nums[p1] ^ nums[p2] ^ (nums[p2] = nums[p1]);
                p2--;
            } else {
                p1++;
            }
        }

        return nums;
    }
}
```
