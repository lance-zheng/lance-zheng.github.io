<!-- markdownlint-disable -->
<!-- customize-tags:数组,哈希表,字符串,排序 -->

# 2418. 按身高排序

> [题目链接](https://leetcode.cn/problems/sort-the-people/)

给你一个字符串数组 `names` ，和一个由 **互不相同** 的正整数组成的数组 `heights` 。两个数组的长度均为 `n` 。

对于每个下标 `i`， `names[i]` 和 `heights[i]` 表示第 `i` 个人的名字和身高。

请按身高 **降序** 顺序返回对应的名字数组 `names` 。

**示例 1：**

```
输入：names = ["Mary","John","Emma"], heights = [180,165,170]
输出：["Mary","Emma","John"]
解释：Mary 最高，接着是 Emma 和 John 。
```

**示例 2：**

```
输入：names = ["Alice","Bob","Bob"], heights = [155,185,150]
输出：["Bob","Alice","Bob"]
解释：第一个 Bob 最高，然后是 Alice 和第二个 Bob 。
```

**提示：**

- `n == names.length == heights.length`
- `1 <= n <= 103`
- `1 <= names[i].length <= 20`
- `1 <= heights[i] <= 105`
- `names[i]` 由大小写英文字母组成
- `heights` 中的所有值互不相同

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**Code:**
选择排序

```go
func sortPeople(names []string, heights []int) []string {
    n := len(heights)

    for i := 0; i < n; i++ {
        target := i
        for j := i + 1; j < n; j++ {
            if heights[target] < heights[j] {
                target = j
            }
        }
        names[i], names[target] = names[target], names[i]
        heights[i], heights[target] =heights[target], heights[i]
    }

    return names
}
```
