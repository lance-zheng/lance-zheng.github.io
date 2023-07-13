<!-- markdownlint-disable -->
<!-- customize-tags:贪心,双指针,字符串 -->

# 680. 验证回文串 II

> [题目链接](https://leetcode.cn/problems/valid-palindrome-ii/)

给你一个字符串  `s`， **最多** 可以从中删除一个字符。

请你判断 `s` 是否能成为回文字符串：如果能，返回 `true` ；否则，返回 `false` 。

**示例 1：**

```
输入：s = "aba"
输出：true
```

**示例 2：**

```
输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。
```

**示例 3：**

```
输入：s = "abc"
输出：false
```

**提示：**

- `1 <= s.length <= 105`
- `s` 由小写英文字母组成

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**  
使用双指针判断是否为回文串，题目只允许删除`一`个字符当遇到不想等的字符时这时需要判断删除左边还是右边的字符，因为后面没有删除机会了，所以我们只需要验证删除左边和删除右边这两个区间内是不是回文串。

```go
func validPalindrome(s string) bool {
    low, high := 0, len(s) - 1

    for low < high {
        if s[low] != s[high] {
            return isPalindrome(s, low + 1, high) || isPalindrome(s, low, high - 1)
        }
        low++
        high--
    }

    return true
}

func isPalindrome(s string, low, high int) bool{
    for low < high {
        if s[low] != s[high] {
            return false
        }
        low++
        high--
    }

    return true
}
```
