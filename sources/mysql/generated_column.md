<!-- customize-category:MySQL -->

# Generated Column (计算列)

## Generated Column

```txt
col_name data_type [GENERATED ALWAYS] AS (expr)
  [VIRTUAL | STORED] [NOT NULL | NULL]
  [UNIQUE [KEY]] [[PRIMARY] KEY]
  [COMMENT 'string']
```

```sql
CREATE TABLE `user`(
first_name varchar(100),
last_name varchar(100),
-- 计算列
full_name varchar(200) AS (concat(first_name, last_name))
);


mysql> select * from user;
+------------+-----------+------------+------------------+
| first_name | last_name | full_name  | full_name_stored |
+------------+-----------+------------+------------------+
| Lance      | Zheng     | LanceZheng | LanceZheng       |
+------------+-----------+------------+------------------+
```

> <http://mingxinglai.com/cn/2015/12/mysql5.7-virtal-column/>
