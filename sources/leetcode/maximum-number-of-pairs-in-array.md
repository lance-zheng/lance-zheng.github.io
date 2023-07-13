<!-- customize-tags:哈希表 -->

# 2341. 数组能形成多少数对

> [题目链接](https://leetcode.cn/problems/maximum-number-of-pairs-in-array)

给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：

从 `nums` 选出 两个 相等的 整数
从 `nums` 中移除这两个整数，形成一个 数对
请你在 `nums` 上多次执行此操作直到无法继续执行。  
返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，其中 `answer[0]` 是形成的数对数目，`answer[1]` 是对 `nums` 尽可能执行上述操作后剩下的整数数目。

**示例 1：**

```text
输入：nums = [1,3,2,1,3,2,2]
输出：[3,1]
解释：
nums[0] 和 nums[3] 形成一个数对，并从 nums 中移除，nums = [3,2,3,2,2] 。
nums[0] 和 nums[2] 形成一个数对，并从 nums 中移除，nums = [2,2,2] 。
nums[0] 和 nums[1] 形成一个数对，并从 nums 中移除，nums = [2] 。
无法形成更多数对。总共形成 3 个数对，nums 中剩下 1 个数字。
```

**示例 2：**

```text
输入：nums = [1,1]
输出：[1,0]
解释：nums[0] 和 nums[1] 形成一个数对，并从 nums 中移除，nums = [] 。
无法形成更多数对。总共形成 1 个数对，nums 中剩下 0 个数字。
```

**示例 3：**

```text
输入：nums = [0]
输出：[0,1]
解释：无法形成数对，nums 中剩下 1 个数字。
```

**提示：**

```text
1 <= nums.length <= 100
0 <= nums[i] <= 100
```

**代码实现：**

```java
class Solution {
    public int[] numberOfPairs(int[] nums) {
        int[] map =new int[101];
        for (int num:nums) {
            map[num]++;
        }
        int[] result = new int[2];
        for (int i = 0; i < map.length; ++i) {
            if (map[i] != 0) {
                result[0] += map[i] / 2;
                result[1] += map[i] % 2;
            }
        }

        return result;
    }
}
```
