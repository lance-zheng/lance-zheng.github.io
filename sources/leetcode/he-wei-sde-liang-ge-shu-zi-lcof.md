<!-- customize-tags:双指针 -->

# 剑指 Offer 57. 和为 s 的两个数字

> 链接：<https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof>

输入一个递增排序的数组和一个数字 s，在数组中查找两个数，使得它们的和正好是 s。如果有多对数字的和等于 s，则输出任意一对即可。

**示例 1：**

```text
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
```

**示例 2：**

```text
输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]
```

**限制：**

```text
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
```

**代码实现：**

```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        int p1 = 0;
        int p2 = nums.length - 1;

        while (p1 < p2) {
            // 加法需注意可能会溢出, 但本体指定了nums[i]的范围所以不会
            int sum = nums[p1] + nums[p2];
            if (sum == target) {
                return new int[]{nums[p1], nums[p2]};
            }
            if (sum > target) {
                p2--;
            } else {
                p1++;
            }
        }

        return new int[]{};

    }
}
```
