<!-- markdownlint-disable -->
<!-- customize-tags:数组,哈希表,计数 -->

# 2404. 出现最频繁的偶数元素

> [题目链接](https://leetcode.cn/problems/most-frequent-even-element/)

给你一个整数数组 `nums` ，返回出现最频繁的偶数元素。

如果存在多个满足条件的元素，只需要返回 **最小** 的一个。如果不存在这样的元素，返回 `-1` 。

**示例 1：**

```
输入：nums = [0,1,2,2,4,4,1]
输出：2
解释：
数组中的偶数元素为 0、2 和 4 ，在这些元素中，2 和 4 出现次数最多。
返回最小的那个，即返回 2 。
```

**示例 2：**

```
输入：nums = [4,4,4,9,2,4]
输出：4
解释：4 是出现最频繁的偶数元素。
```

**示例 3：**

```
输入：nums = [29,47,21,41,13,37,25,7]
输出：-1
解释：不存在偶数元素。
```

**提示：**

- `1 <= nums.length <= 2000`
- `0 <= nums[i] <= 105`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

```go
func mostFrequentEven(nums []int) int {
    m := make(map[int]int)
    ans, count := -1, 0
    for _, num := range nums {
        if num & 1 == 1 {
            continue
        }

        m[num]++
        if m[num] > count {
            ans = num
            count = m[num]
        } else if m[num] == count && num < ans {
            ans = num
        }
    }

    return ans
}
```
