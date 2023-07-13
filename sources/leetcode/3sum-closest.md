<!-- markdownlint-disable -->
<!-- customize-tags:数组,双指针,排序 -->

# 16. 最接近的三数之和

> [题目链接](https://leetcode.cn/problems/3sum-closest/)

给你一个长度为 `n` 的整数数组  `nums` 和 一个目标值  `target`。请你从 `nums` 中选出三个整数，使它们的和与  `target`  最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

**示例 1：**

```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
```

**示例 2：**

```
输入：nums = [0,0,0], target = 1
输出：0
```

**提示：**

- `3 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-104 <= target <= 104`

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**代码实现：**

```go

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    length := len(nums)

    ans := nums[0] + nums[1] + nums[2]
    for i := 0; i < length - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] { // 避免重复计算
            continue
        }
        // 优化
        sum := nums[i] + nums[i + 1] + nums[i + 2]
        if sum > target {
            if Abs(target - sum) < Abs(target - ans) {
                ans = sum
            }
            break
        }
        // 优化
        sum = nums[i] + nums[length - 1] + nums[length - 2]
        if sum < target {
            if Abs(target - sum) < Abs(target - ans) {
                ans = sum
            }
            continue
        }

        L, R := i + 1, length - 1

        for L < R {
            sum := nums[i] + nums[L] + nums[R]
            if sum == target {
                return sum
            }

            if Abs(target - sum) < Abs(target - ans) {
                ans = sum
            }
            if sum < target {
                L++
                for L < R && nums[L] == nums[L - 1] {
                    L++
                }
            } else {
                R--
                for L < R && nums[R] == nums[R + 1] {
                    R--
                }
            }
        }
    }

    return ans
}

func Abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}
```
