<!-- customize-category:MySQL -->

# View

## 什么是视图

视图是一个虚拟的表，视图本身不储存数据，他的数据是来自于一条查询语句。

视图创建成功之后和和普通表用法差不多，但是修改视图数据时会有些限制。

## 为什么要用视图

使用视图有以下优点：

1. 重用 SQL
2. 简化复杂查询：使用者无需关心它的查询细节。
3. 权限控制：可以通过视图给用户授予部分数据的访问权限，而不是整个表的数据。

## 使用视图

创建视图使用 `CREATE VIEW view_name AS 查询语句`。
创建或替换视图使用 `CREATE OR REPLACE VIEW view_name AS 查询语句`。
_视图嵌套使用：创建的视图同样可以被其他视图使用_

删除视图 `DROP VIEW view_name` **删除视图并不会删除基表中的数据**。
查询视图创建语句 `SHOW CREATE VIEW view_name`可以查询视图的创建语句。

示例：

```sql
-- 创建视图 查询语句可以写的很复杂
CREATE VIEW `test_view` AS
SELECT 1;
-- 创建或更新视图
CREATE OR REPLACE VIEW `test_view` AS
SELECT 1;
-- 查询创建语句
SHOW CREATE VIEW `test_view`;
-- 删除视图
DROP VIEW `test_view`;
```

计算字段：

```sql
CREATE TABLE `user` (
    first_name varchar(255),
    last_name varchar(255)
);
INSERT INTO `user` value('Lance', 'Zheng');

CREATE VIEW `user_view` AS
SELECT first_name,last_name, CONCAT(first_name,' ',last_name) AS full_name FROM `user`;

SELECT * FROM user_view;

|first_name|last_name|full_name  |
|----------+---------+-----------+
|Lance     |Zheng    |Lance Zheng|
```

通过视图可用用来**过滤数据**，**格式转换**，**计算字段**等

## 更新视图中的数据

不推荐

## 总结

视图可以简化复杂查询，使用视图时同时应要注意查询性能。

表结构比较简单的项目不推荐使用，一般用于表结构复杂的项目。

视图作为一个虚拟的表修，被改时会直接去修改基表的数据，若 MySQL 无法定位到具体的数据则无法更新。**视图应该只用于查询，不推荐通过视图修改基表中数据**。
