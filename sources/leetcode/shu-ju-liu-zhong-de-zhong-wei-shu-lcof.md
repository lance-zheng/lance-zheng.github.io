<!-- customize-tags:堆 -->

# 剑指 Offer 41. 数据流中的中位数

> [题目链接](https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/)

如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如：

- [2,3,4]  的中位数是 3
- [2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

- void addNum(int num) - 从数据流中添加一个整数到数据结构中。
- double findMedian() - 返回目前所有元素的中位数。

**示例 1：**

```text
输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
```

**示例 2：**

```text
输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]
```

**限制：**

- 最多会对  addNum、findMedian 进行  50000  次调用。

**解题思路：**

```text
采用两个堆的数据结构来存储，其中一个是大顶堆，另一个是小顶对。
在存储元素时两个堆的数量差不可以超过 1，也就是说在放置元素时要保证两个堆的平衡，当数量为奇数时我们统一存放到小顶堆中。
这样存储的话中位数都是在堆的顶部，当数量为奇数时就去小顶堆的数为结果，数量为偶数时就是两个堆顶的和除2.
```

**代码实现：**

```java
class MedianFinder {
    PriorityQueue<Integer> a;
    PriorityQueue<Integer> b;

    public MedianFinder() {
        a = new PriorityQueue<Integer>(Collections.reverseOrder());
        b = new PriorityQueue<Integer>();
    }

    public void addNum(int num) {
        if (a.isEmpty() || num <= a.peek()) {
            a.offer(num);
            if (a.size() > b.size() + 1) {
                b.offer(a.poll());
            }
        } else {
            b.offer(num);
            if (b.size() > a.size()) {
                a.offer(b.poll());
            }
        }
    }

    public double findMedian() {
        int sizeA = a.size();
        int sizeB = b.size();
        if (sizeA > sizeB) {
            return a.peek();
        } else {
            return (sizeA == sizeB) ? (a.peek() + b.peek()) / 2.0 : b.peek();
        }
    }
}
```
