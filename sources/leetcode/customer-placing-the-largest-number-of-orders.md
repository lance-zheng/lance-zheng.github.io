<!-- markdownlint-disable -->
<!-- customize-tags:数据库 -->

# 586. 订单最多的客户

> [题目链接](https://leetcode.cn/problems/customer-placing-the-largest-number-of-orders/)

表: `Orders`

```
+-----------------+----------+
| Column Name     | Type     |
+-----------------+----------+
| order_number    | int      |
| customer_number | int      |
+-----------------+----------+
Order_number是该表的主键。
此表包含关于订单ID和客户ID的信息。
```

编写一个 SQL 查询，为下了 **最多订单** 的客户查找 `customer_number` 。

测试用例生成后， **恰好有一个客户** 比任何其他客户下了更多的订单。

查询结果格式如下所示。

**示例 1:**

```
输入:
Orders 表:
+--------------+-----------------+
| order_number | customer_number |
+--------------+-----------------+
| 1            | 1               |
| 2            | 2               |
| 3            | 3               |
| 4            | 3               |
+--------------+-----------------+
输出:
+-----------------+
| customer_number |
+-----------------+
| 3               |
+-----------------+
解释:
customer_number 为 '3' 的顾客有两个订单，比顾客 '1' 或者 '2' 都要多，因为他们只有一个订单。
所以结果是该顾客的 customer_number ，也就是 3 。
```

**进阶：** 如果有多位顾客订单数并列最多，你能找到他们所有的 `customer_number` 吗？

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->

**SQL:**

```sql
# Write your MySQL query statement below

select customer_number from Orders group by customer_number order by count(1) desc limit 1;

# 进阶利用 having
select customer_number from Orders group by customer_number having count(1) = (
    select count(1) from Orders o2 group by customer_number order by count(1) desc limit 1
)
```
