<!-- markdownlint-disable -->
<!-- customize-tags:字典树,双指针,字符串,字符串匹配 -->

# 1023. 驼峰式匹配

> [题目链接](https://leetcode.cn/problems/camelcase-matching/)

如果我们可以将 **小写字母** 插入模式串  `pattern`  得到待查询项  `query`，那么待查询项与给定模式串匹配。（我们可以在任何位置插入每个字符，也可以插入 0 个字符。）

给定待查询列表  `queries`，和模式串  `pattern`，返回由布尔值组成的答案列表  `answer`。只有在待查项  `queries[i]` 与模式串  `pattern` 匹配时， `answer[i]`  才为 `true`，否则为 `false`。

**示例 1：**

```
输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
输出：[true,false,true,true,false]
示例：
"FooBar" 可以这样生成："F" + "oo" + "B" + "ar"。
"FootBall" 可以这样生成："F" + "oot" + "B" + "all".
"FrameBuffer" 可以这样生成："F" + "rame" + "B" + "uffer".
```

**示例 2：**

```
输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
输出：[true,false,true,false,false]
解释：
"FooBar" 可以这样生成："Fo" + "o" + "Ba" + "r".
"FootBall" 可以这样生成："Fo" + "ot" + "Ba" + "ll".
```

**示例 3：**

```
输出：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
输入：[false,true,false,false,false]
解释：
"FooBarTest" 可以这样生成："Fo" + "o" + "Ba" + "r" + "T" + "est".
```

**提示：**

1. `1 <= queries.length <= 100`
2. `1 <= queries[i].length <= 100`
3. `1 <= pattern.length <= 100`
4. 所有字符串都仅由大写和小写英文字母组成。

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

代码实现：

```go
func camelMatch(queries []string, pattern string) []bool {
    ans := make([]bool, len(queries))
    for i := range queries {
        ans[i] = isMatch(queries[i], pattern)
    }

    return ans
}

func isMatch(query, pattern string) bool {
    var k int
    for i := range query {
        if k < len(pattern) && query[i] == pattern[k] {
            k++
        } else if unicode.IsUpper(rune(query[i])) {
            return false
        }
    }

    return k == len(pattern)
}
```
