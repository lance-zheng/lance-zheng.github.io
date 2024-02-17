<!-- customize-category:MySQL -->

# 窗口函数

- [窗口函数](#窗口函数)
  - [ROW_NUMBER()](#row_number)
    - [找到每组的前 N 行](#找到每组的前-n-行)

## ROW_NUMBER()

`ROW_NUMBER()` 是 MySQL 8.0 引入的功能，它从 `1` 开始为每一行生成一个行号

```text
ROW_NUMBER() OVER (<partition_definition> <order_definition>)
```

**`PARTITION BY`**：

```txt
PARTITION BY <expression>,[{,<expression>}...]

对结果集进行分组，然后组内序号从 `1` 开启累加，支持多个字段。当您使用 `PARTITION BY` 子句时，每个分区也可以被视为一个窗口。
```

**`ORDER BY`**：

```txt
ORDER BY <expression> [ASC|DESC],[{,<expression>}...]

ORDER BY 子句的目的是设置行的顺序。此 ORDER BY 子句独立 ORDER BY 于查询的子句。
```

**测试数据：**

```sql
CREATE TABLE `supplier_kpi` (
  `name` varchar(255) NOT NULL,
  `score` INT DEFAULT NULL,
  `year`  INT NOT NULL,
  `month` INT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

INSERT INTO `supplier_kpi` VALUES
('Tom', 99, 2023, 2),
('Tom', 80, 2023, 1),
('Tom', null, 2022, 12),
('Tom', 77, 2022, 11),
('Tom', 100, 2022, 10),
('Tom', 80, 2022, 9),
('Lance', 99, 2023, 2),
('Lance', 80, 2023, 1),
('Lance', 80, 2022, 12),
('Lance', 77, 2022, 11),
('Lance', 100, 2022, 10),
('Lance', 80, 2022, 9);
```

**案例：**

1. 获取每一行的编号 `ROW_NUMBER() OVER() AS row_num`

   ```sql
   SELECT ROW_NUMBER() OVER() AS row_num, `name`, `score`,`year`,`month` FROM supplier_kpi;
   ```

   输出

   ```txt
   +---------+-------+-------+------+-------+
   | row_num | name  | score | year | month |
   +---------+-------+-------+------+-------+
   |       1 | Tom   |    99 | 2023 |     2 |
   |       2 | Tom   |    80 | 2023 |     1 |
   |       3 | Tom   |  NULL | 2022 |    12 |
   |       4 | Tom   |    77 | 2022 |    11 |
   |       5 | Tom   |   100 | 2022 |    10 |
   |       6 | Tom   |    80 | 2022 |     9 |
   |       7 | Lance |    99 | 2023 |     2 |
   |       8 | Lance |    80 | 2023 |     1 |
   |       9 | Lance |    80 | 2022 |    12 |
   |      10 | Lance |    77 | 2022 |    11 |
   |      11 | Lance |   100 | 2022 |    10 |
   |      12 | Lance |    80 | 2022 |     9 |
   +---------+-------+-------+------+-------+
   ```

2. 根据名字分组然后获取序号  
   可以使用 `OVER(PARTITION BY name)`进行分组，这里会根据 name 分组然后组内从 1 开始产生序号

   ```sql
   SELECT ROW_NUMBER() OVER(PARTITION BY `name`) AS row_num, `name`, `score`,`year`,`month` FROM supplier_kpi;
   ```

   输出

   ```txt
   +---------+-------+-------+------+-------+
   | row_num | name  | score | year | month |
   +---------+-------+-------+------+-------+
   |       1 | Lance |    99 | 2023 |     2 |
   |       2 | Lance |    80 | 2023 |     1 |
   |       3 | Lance |    80 | 2022 |    12 |
   |       4 | Lance |    77 | 2022 |    11 |
   |       5 | Lance |   100 | 2022 |    10 |
   |       6 | Lance |    80 | 2022 |     9 |
   |       1 | Tom   |    99 | 2023 |     2 |
   |       2 | Tom   |    80 | 2023 |     1 |
   |       3 | Tom   |  NULL | 2022 |    12 |
   |       4 | Tom   |    77 | 2022 |    11 |
   |       5 | Tom   |   100 | 2022 |    10 |
   |       6 | Tom   |    80 | 2022 |     9 |
   +---------+-------+-------+------+-------+
   ```

3. 对数据排序后再生成序号
   可以使用 `OVER(ORDER BY column, column2 DESC)` 对数据排序然后再生成序号

   ```sql
   SELECT ROW_NUMBER() OVER(ORDER BY `score` DESC) AS row_num, `name`, `score`,`year`,`month` FROM supplier_kpi;
   ```

   输出：这里根据了 `score` 降序排列然后生成序号

   ```txt
   +---------+-------+-------+------+-------+
   | row_num | name  | score | year | month |
   +---------+-------+-------+------+-------+
   |       1 | Tom   |   100 | 2022 |    10 |
   |       2 | Lance |   100 | 2022 |    10 |
   |       3 | Tom   |    99 | 2023 |     2 |
   |       4 | Lance |    99 | 2023 |     2 |
   |       5 | Tom   |    80 | 2023 |     1 |
   |       6 | Tom   |    80 | 2022 |     9 |
   |       7 | Lance |    80 | 2023 |     1 |
   |       8 | Lance |    80 | 2022 |    12 |
   |       9 | Lance |    80 | 2022 |     9 |
   |      10 | Tom   |    77 | 2022 |    11 |
   |      11 | Lance |    77 | 2022 |    11 |
   |      12 | Tom   |  NULL | 2022 |    12 |
   +---------+-------+-------+------+-------+
   ```

---

### 找到每组的前 N 行

查询某个月份及其前 `3` 次的 `score`，不包含 `null`。

这里先通过 `where` 先将之前时间的数据筛选出来，然后使用 `(PARTITION BY name ORDER BY year DESC, month DESC)`，对名字进行分组然后根据时间降序排序生成行号，最后在筛选出行号是 4 之前的数据。

```sql
WITH t AS (
 SELECT ROW_NUMBER() OVER(PARTITION BY `name` ORDER BY `year` DESC, `month` DESC) AS seq, `name`, `score`,`year`, `month` FROM supplier_kpi
 WHERE
 `score` IS NOT NULL AND
 ((`year` = 2023 AND `month` <= 1) OR `year` < 2023)
)
SELECT * FROM t WHERE t.seq < 4;
```

输出

```txt
+-----+-------+-------+------+-------+
| seq | name  | score | year | month |
+-----+-------+-------+------+-------+
|   1 | Lance |    80 | 2023 |     1 |
|   2 | Lance |    80 | 2022 |    12 |
|   3 | Lance |    77 | 2022 |    11 |
|   1 | Tom   |    80 | 2023 |     1 |
|   2 | Tom   |    77 | 2022 |    11 |
|   3 | Tom   |   100 | 2022 |    10 |
+-----+-------+-------+------+-------+
```
