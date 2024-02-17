<!-- markdownlint-disable -->
<!-- customize-tags:贪心,哈希表,字符串 -->

# 409. 最长回文串

> [题目链接](https://leetcode.cn/problems/longest-palindrome/)

给定一个包含大写字母和小写字母的字符串 `s` ，返回  \*通过这些字母构造成的 **最长的回文串\*** 。

在构造过程中，请注意 **区分大小写** 。比如  `"Aa"`  不能当做一个回文字符串。

**示例 1:**

```
输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
```

**示例 2:**

```
输入:s = "a"
输出:1
```

**示例 3：**

```
输入:s = "aaaaaccc"
输出:7
```

**提示:**

- `1 <= s.length <= 2000`
- `s`  只由小写 **和/或** 大写英文字母组成

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**
位运算`num&1` 若 `num` 是偶数则返回 `0` 奇数返回 `1`

```go
func longestPalindrome(s string) int {
    m := make(map[rune]int)
    for _, v := range s {
        count := m[v] + 1
        m[v] = count
    }
    var count int
    for _, v := range m {
        count += v >> 1
    }
    count = count << 1
    if count < len(s) {
        count++
    }

    return count
}
```
