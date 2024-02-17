<!-- customize-category:LeetCode-->

# Notes

- [Notes](#notes)
  - [判断一个数是否是偶数](#判断一个数是否是偶数)
  - [交换两个数](#交换两个数)
  - [排序算法](#排序算法)
    - [快速排序](#快速排序)

## 判断一个数是否是偶数

```java
num & 1 // 如果返回值是 1 则说明是奇数 0 是偶数
num >> 1 << 1 // 去除奇数部分
```

## 交换两个数

不使用第三个变量交换两个数

```java
// java
public class Test {
    public static void main(String args[]) {
        int a = 1;
        int b = 2;
        // 方法一 数学运算
        a = a + b;
        b = a - b;
        a = a - b;
        // 简写
        a = a + b - (b = a);

        // 方法二 位运算
        a = a ^ b;
        b = a ^ b;
        a = a ^ b;
        // 简写
        a = a ^ b ^ (b = a);
    }
}
```

```go
// golang
func main() {
    a,b := 1, 2
    a,b = b,a
}
```

## 排序算法

### 快速排序

快速排序的思路就是先在数组内选择一个基准值 `pivot`，让后将小于 `pivot` 的值放到 `pivot` 的左边，将大于 `pivot` 的值放到右边，然后在递归的对 `pivot` 的左区间和右区间做同样的操作。

- 时间复杂度为：O(nlogn)
- 空间复杂度为：O(logn)

**代码实现：**

```java
public void quickSort(int[] arr, int left, int right) {
    if (left >= right) {
        return;
    }
    int pivot = arr[left + (right - left >> 1)];

    int i = left, j = right;

    while (i <= j) {
        while (arr[i] < pivot) { // 从左边找一个大于基准值的数
            i++;
        }
        while (arr[j] > pivot) { // 从右边找一个小于或等于基准值的数
            j--;
        }
        if (i <= j) {
            arr[i] = arr[i] ^ arr[j] ^ (arr[j] = arr[i]); // 交换两个数
            i++;
            j--;
        }
    }
    quickSort(arr, left, j); // 排序左区间
    quickSort(arr, i, right); // 排序右区间
}
```
