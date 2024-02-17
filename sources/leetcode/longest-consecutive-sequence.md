<!-- customize-tags:哈希表 -->

# 128. 最长连续序列

> [题目链接](https://leetcode.cn/problems/longest-consecutive-sequence/)

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。  
请你设计并实现时间复杂度为  O(n) 的算法解决此问题。

**示例 1：**

```text
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
```

**示例 2：**

```text
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
```

**提示：**

```text
0 <= nums.length <= 105
-109 <= nums[i] <= 109
```

**方法一：**

```markdown
# 哈希表

定义一个哈希表，先将所有数字存入。然后遍历数组中的元素，
如果当前 元素-1 在哈希表中不存在，则说明当前元素是某个序列的边界值。
在以当前元素为起点向后寻找同时统计个数，直到遇到在哈希表中不存在的元素。
```

**代码实现：**

```java
class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int n : nums) {
            set.add(n);
        }

        int maxSeq = 0;

        for (int n:nums) {
            if (!set.contains(n - 1)) {
                int curSeq = 1;
                while (set.contains(++n)) {
                    curSeq++;
                }
                maxSeq = Math.max(curSeq, maxSeq);
            }
        }

        return maxSeq;
    }
}
```
