<!-- markdownlint-disable -->
<!-- customize-tags:栈,字符串 -->

# 20. 有效的括号

> [题目链接](https://leetcode.cn/problems/valid-parentheses/)

给定一个只包括 `'('`， `')'`， `'{'`， `'}'`， `'['`， `']'`  的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

**示例 1：**

```
输入：s = "()"
输出：true
```

**示例  2：**

```
输入：s = "()[]{}"
输出：true
```

**示例  3：**

```
输入：s = "(]"
输出：false

```

**提示：**

- `1 <= s.length <= 104`
- `s` 仅由括号 `'()[]{}'` 组成

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

```go
func isValid(s string) bool {
    if len(s)&1 != 0 {
        return false
    }
    stack := []rune{}

    for _, c := range s {
        switch c {
        case '(':
            stack = append(stack, ')')
        case '[':
            stack = append(stack, ']')
        case '{':
            stack = append(stack, '}')
        default:
            if len(stack) == 0 || stack[len(stack)-1] != c {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
```
