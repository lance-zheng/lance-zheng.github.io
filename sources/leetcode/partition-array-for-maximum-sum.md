<!-- markdownlint-disable -->
<!-- customize-tags:数组,动态规划 -->

# 1043. 分隔数组以得到最大和

> [题目链接](https://leetcode.cn/problems/partition-array-for-maximum-sum/)

给你一个整数数组 `arr`，请你将该数组分隔为长度 **最多** 为 k 的一些（连续）子数组。分隔完成后，每个子数组的中的所有值都会变为该子数组中的最大值。

返回将数组分隔变换后能够得到的元素最大和。本题所用到的测试用例会确保答案是一个 32 位整数。

**示例 1：**

```
输入：arr = [1,15,7,9,2,5,10], k = 3
输出：84
解释：数组变为 [15,15,15,9,10,10,10]
```

**示例 2：**

```
输入：arr = [1,4,1,5,7,3,6,1,9,9,3], k = 4
输出：83
```

**示例 3：**

```
输入：arr = [1], k = 1
输出：1
```

**提示：**

- `1 <= arr.length <= 500`
- `0 <= arr[i] <= 109`
- `1 <= k <= arr.length`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：动态规划**  
采用动态规划解法

```markdown
# n < k

dp[n] = max \* n

# n > k

dp[n] =
arr[n] + dp[n - 1]
2 - max(arr[n], arr[n - 1]) + dp[n - 2]
3 - max(arr[n], arr[n - 1], arr[n - 2]) + dp[n - 3]
...
就是不断扩大当前数组的长度，然后取最优解
```

```go
func maxSumAfterPartitioning(arr []int, k int) int {
    dp := make([]int, len(arr) + 1)

    for i := range dp {
        var maxVal int
        for j := i - 1; j >= i - k && j >= 0; j-- {
            maxVal = max(maxVal, arr[j])
            dp[i] = max(dp[i], dp[j] + maxVal * (i - j))
        }
    }
    return dp[len(arr)]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
