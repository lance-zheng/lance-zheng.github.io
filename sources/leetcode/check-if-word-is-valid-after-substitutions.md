<!-- markdownlint-disable -->
<!-- customize-tags:栈,字符串 -->

# 1003. 检查替换后的词是否有效

> [题目链接](https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions/)

给你一个字符串 `s` ，请你判断它是否 **有效** 。

字符串 `s` **有效** 需要满足：假设开始有一个空字符串 `t = ""` ，你可以执行 **任意次** 下述操作将 `t` **转换为** `s` ：

- 将字符串 `"abc"` 插入到 `t` 中的任意位置。形式上， `t` 变为 `tleft + "abc" + tright`，其中 `t == tleft + tright` 。注意， `tleft` 和 `tright` 可能为 **空** 。

如果字符串 `s` 有效，则返回 `true`；否则，返回 `false`。

**示例 1：**

```
输入：s = "aabcbc"
输出：true
解释：
"" -> "abc" -> "aabcbc"
因此，"aabcbc" 有效。
```

**示例 2：**

```
输入：s = "abcabcababcc"
输出：true
解释：
"" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
因此，"abcabcababcc" 有效。
```

**示例 3：**

```
输入：s = "abccba"
输出：false
解释：执行操作无法得到 "abccba" 。
```

**提示：**

- `1 <= s.length <= 2 * 104`
- `s` 由字母 `'a'`、 `'b'` 和 `'c'` 组成

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**Code：**
使用栈数据结构

```go
func isValid(s string) bool {
    var stack []byte

    for i := range s {
        stack = append(stack, s[i])

        length := len(stack)
        if length >= 3 &&
           stack[length - 1] == 'c' &&
           stack[length - 2] == 'b' &&
           stack[length - 3] == 'a' {
               stack = stack[:length - 3]
        }
    }

    return len(stack) == 0
}
```
