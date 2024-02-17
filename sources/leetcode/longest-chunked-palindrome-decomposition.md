<!-- markdownlint-disable -->
<!-- customize-tags:贪心,双指针,字符串,动态规划,哈希函数,滚动哈希 -->

# 1147. 段式回文

> [题目链接](https://leetcode.cn/problems/longest-chunked-palindrome-decomposition/)

你会得到一个字符串  `text` 。你应该把它分成 `k`  个子字符串  `(subtext1, subtext2，…， subtextk)` ，要求满足:

- `subtexti` 是 **非空** 字符串
- 所有子字符串的连接等于 `text` ( 即 `subtext1 + subtext2 + ... + subtextk == text` )
- 对于所有 i  的有效值( 即  `1 <= i <= k` ) ， `subtexti == subtextk - i + 1` 均成立

返回 `k` 可能最大值。

**示例 1：**

```
输入：text = "ghiabcdefhelloadamhelloabcdefghi"
输出：7
解释：我们可以把字符串拆分成 "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)"。
```

**示例 2：**

```
输入：text = "merchant"
输出：1
解释：我们可以把字符串拆分成 "(merchant)"。
```

**示例 3：**

```
输入：text = "antaprezatepzapreanta"
输出：11
解释：我们可以把字符串拆分成 "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)"。
```

**提示：**

- `1 <= text.length <= 1000`
- `text`  仅由小写英文字符组成

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**  
每次选取一个最短的段式回文串，然后对子区间重复操作。

```go
func longestDecomposition(text string) int {
    var ans int
    start, end := 0, len(text)-1
    for start <= end {
        length := foo(text, start, end)
        if length < (end - start + 1) {
            ans += 2
        } else {
            ans += 1
        }
        start += length
        end -= length
    }
    return ans
}

// 返回最短的段式回文串长度
func foo(text string, start, end int) int {
    length := 1
    all := end - start + 1

    for length < (all << 1) {
        L := start
        R := end - length + 1

        for R <= end && text[L] == text[R] {
            L++
            R++
        }
        if R == end+1 {
            break
        }
        length++
    }
    return length
}
```
