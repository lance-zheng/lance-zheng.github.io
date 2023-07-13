<!-- markdownlint-disable -->
<!-- customize-tags:数组,哈希表,计数 -->

# 1010. 总持续时间可被 60 整除的歌曲

> [题目链接](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/)

在歌曲列表中，第 `i` 首歌曲的持续时间为 `time[i]` 秒。

返回其总持续时间（以秒为单位）可被 `60` 整除的歌曲对的数量。形式上，我们希望下标数字 `i` 和 `j` 满足   `i < j` 且有  `(time[i] + time[j]) % 60 == 0`。

**示例 1：**

```
输入：time = [30,20,150,100,40]
输出：3
解释：这三对的总持续时间可被 60 整除：
(time[0] = 30, time[2] = 150): 总持续时间 180
(time[1] = 20, time[3] = 100): 总持续时间 120
(time[1] = 20, time[4] = 40): 总持续时间 60
```

**示例 2：**

```
输入：time = [60,60,60]
输出：3
解释：所有三对的总持续时间都是 120，可以被 60 整除。
```

**提示：**

- `1 <= time.length <= 6 * 104`
- `1 <= time[i] <= 500`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**Code：**  
枚举所有歌曲播放时间，若当前歌曲播放时间可以被 `60` 整除，则只需要找到之前也可以被 **60** 整除的歌曲即可，若当前不能被 **60** 整除，只需要找到余数想加能够被 **60** 整除的歌曲即可。

```go
func numPairsDivisibleBy60(time []int) int {
    var ans int
    count := [60]int{}

    for i := range time {
        t := time[i] % 60

        if t == 0 {
            ans += count[0]
        } else {
            ans += count[60 - t]
        }

        count[t]++
    }

    return ans
}
```
