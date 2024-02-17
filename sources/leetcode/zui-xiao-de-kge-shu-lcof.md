<!-- customize-tags:排序 -->

# 剑指 Offer 40. 最小的 k 个数

> [题目链接](https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof)

输入整数数组 `arr` ，找出其中最小的 `k` 个数。例如，输入 4、5、1、6、2、7、3、8 这 8 个数字，则最小的 4 个数字是 1、2、3、4。

**示例 1：**

```text
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
```

**示例 2：**

```text
输入：arr = [0,1,2,1], k = 1
输出：[0]
```

**限制：**

```text
0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
```

**方法一：**

```markdown
# 快速排序

使用快速排序，题目只需要查找前 k 个最小的数，所以不需要对整个数组排序
```

**代码实现：**

```java
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        quickSort(arr, 0, arr.length - 1, k);
        return Arrays.copyOf(arr, k);
    }

    private void quickSort(int[] arr, int left, int right, int k) {
        if (left >= right) {
            return;
        }
        int pivot = arr[left + (right - left >> 1)];

        int i = left, j = right;

        while (i <= j) {
            while (arr[i] < pivot) {
                i++;
            }
            while (arr[j] > pivot) {
                j--;
            }
            if (i <= j) {
                arr[i] = arr[i] ^ arr[j] ^(arr[j] = arr[i]);
                i++;
                j--;
            }
        }
        // 优化 只排前 k 个数所在的区间
        if (k <= j) {
            quickSort(arr, left, j, k);
        } else if (k >= i) {
            quickSort(arr, i, right, k);
        }
    }
}
```
