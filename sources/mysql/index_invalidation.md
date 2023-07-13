<!-- customize-category:MySQL -->

# MySQL 索引失效场景

- [MySQL 索引失效场景](#mysql-索引失效场景)
  - [联合索引不符合最左匹配原则](#联合索引不符合最左匹配原则)
  - [ORDER BY 导致索引失效](#order-by-导致索引失效)
  - [函数调用](#函数调用)
  - [运算符](#运算符)
  - [LIKE column '%xxx'](#like-column-xxx)
  - [OR 中有字段没有索引](#or-中有字段没有索引)
  - [查询数据是表中大部分数据](#查询数据是表中大部分数据)

测试数据：

```sql
CREATE TABLE t_user (
  id INT(11) NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  age INT(11) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(20),
  address VARCHAR(100),
  city VARCHAR(50),
  state VARCHAR(50),
  country VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE t_user ADD INDEX idx_city_state_country(city, state, country);
ALTER TABLE t_user ADD INDEX idx_t_user_age(age);


INSERT INTO t_user (first_name, last_name, age, email, phone, address, city, state, country) VALUES
('John', 'Smith', 25, 'john@example.com', '555-1234', '123 Main St', 'Anytown', 'CA', 'USA'),
('Jane', 'Doe', 30, 'jane@example.com', '555-5678', '456 Elm St', 'Otherville', 'NY', 'USA'),
('David', 'Johnson', 42, 'david@example.com', '555-9012', '789 Oak St', 'Somecity', 'TX', 'USA'),
('Emily', 'Jones', 27, 'emily@example.com', '555-3456', '234 Cedar St', 'Othercity', 'FL', 'USA'),
('Michael', 'Brown', 35, 'michael@example.com', '555-7890', '567 Maple St', 'Smallville', 'OH', 'USA'),
('Sophia', 'Wilson', 28, 'sophia@example.com', '555-2345', '901 Pine St', 'Bigtown', 'NC', 'USA'),
('William', 'Davis', 33, 'william@example.com', '555-6789', '345 Birch St', 'Somewhere', 'MI', 'USA'),
('Emma', 'Taylor', 29, 'emma@example.com', '555-0123', '678 Spruce St', 'Anyplace', 'WA', 'USA'),
('Benjamin', 'Miller', 36, 'benjamin@example.com', '555-4567', '890 Oak St', 'Anothercity', 'IL', 'USA'),
('Olivia', 'Anderson', 31, 'olivia@example.com', '555-8901', '123 Pine St', 'Nowhere', 'KS', 'USA');
```

## 联合索引不符合最左匹配原则

最左匹配的意思是假设目前有个联合索引是 `(city, state, country)` 要想使用的这个索引那么查询条件就必须从 `city` 开始。如果中间有字段没有被使用则后面的字段也无法使用索引。  
以下的例子是可以用到索引的：

- `where city = '' and state = '' and country = ''`
- `where city = '' and state = ''`
- `where state = '' and city = ''`：前后位置没关系
- `where city = '' and country = ''`：`city` 可以使用索引，但是 `country` 不能使用索引

创建索引:

下面可以使用到索引

```sql
EXPLAIN SELECT * FROM t_user WHERE country = '' AND state = '' AND city = '';
/*
Name         |Value                 |
-------------+----------------------+
id           |1                     |
select_type  |SIMPLE                |
table        |t_user                |
partitions   |                      |
type         |ref                   |
possible_keys|idx_city_state_country|
key          |idx_city_state_country|
key_len      |609                   |
ref          |const,const,const     |
rows         |1                     |
filtered     |100.0                 |
Extra        |                      |
*/
```

下面这种情况直接跳过了 `city` 不符合最左匹配原则，所以无法使用到索引

```sql
EXPLAIN SELECT * FROM t_user WHERE country = '' AND state = '';

/*
Name         |Value      |
-------------+-----------+
id           |1          |
select_type  |SIMPLE     |
table        |t_user     |
partitions   |           |
type         |ALL        |
possible_keys|           |
key          |           |
key_len      |           |
ref          |           |
rows         |10         |
filtered     |10.0       |
Extra        |Using where|
*/
```

## ORDER BY 导致索引失效

order by 在某些情况下也会导致索引失效
下面这种情况我们在 `age` 和 `city, state, country` 都有索引但是最后却没有使用到索引。

```sql
EXPLAIN SELECT age, city, state, country FROM t_user tu ORDER BY age DESC;
/*
Name         |Value         |
-------------+--------------+
id           |1             |
select_type  |SIMPLE        |
table        |tu            |
partitions   |              |
type         |ALL           |
possible_keys|              |
key          |              |
key_len      |              |
ref          |              |
rows         |10            |
filtered     |100.0         |
Extra        |Using filesort|
*/
```

解决方法可以为 `city, state, country, age` 创建联合索引。

## 函数调用

索引列进行函数调用导致索引失效

```sql
EXPLAIN SELECT * FROM t_user tu WHERE IFNULL(city, 'Somewhere') = 'Anytown';
/*
Name         |Value      |
-------------+-----------+
id           |1          |
select_type  |SIMPLE     |
table        |tu         |
partitions   |           |
type         |ALL        |
possible_keys|           |
key          |           |
key_len      |           |
ref          |           |
rows         |10         |
filtered     |100.0      |
Extra        |Using where|
*/
```

## 运算符

下面对索引列进行了计算导致索引失效

```sql
EXPLAIN SELECT * FROM t_user WHERE age + 1  = 25;
/*
Name         |Value      |
-------------+-----------+
id           |1          |
select_type  |SIMPLE     |
table        |t_user     |
partitions   |           |
type         |ALL        |
possible_keys|           |
key          |           |
key_len      |           |
ref          |           |
rows         |10         |
filtered     |100.0      |
Extra        |Using where|
*/
```

## LIKE column '%xxx'

`%` 最前面也会导致索引失效，如果放到后面或者中间的某个位置还是可以用到索引的。

```sql
EXPLAIN SELECT * FROM t_user tu WHERE city LIKE '%Anytown';
/*
Name         |Value      |
-------------+-----------+
id           |1          |
select_type  |SIMPLE     |
table        |tu         |
partitions   |           |
type         |ALL        |
possible_keys|           |
key          |           |
key_len      |           |
ref          |           |
rows         |10         |
filtered     |11.11      |
Extra        |Using where|
*/
```

## OR 中有字段没有索引

下面的例子中 `age` 和 `city` 都有索引

```sql
EXPLAIN SELECT * FROM t_user tu WHERE age = 25 OR city = ''
/*
Name         |Value                                                               |
-------------+--------------------------------------------------------------------+
id           |1                                                                   |
select_type  |SIMPLE                                                              |
table        |tu                                                                  |
partitions   |                                                                    |
type         |index_merge                                                         |
possible_keys|idx_t_user_age,idx_city_state_country                               |
key          |idx_t_user_age,idx_city_state_country                               |
key_len      |4,203                                                               |
ref          |                                                                    |
rows         |2                                                                   |
filtered     |100.0                                                               |
Extra        |Using sort_union(idx_t_user_age,idx_city_state_country); Using where|
*/
```

`first_name` 没有索引导致索引失效

```sql
EXPLAIN SELECT * FROM t_user tu WHERE age = 25 OR first_name = ''
```

## 查询数据是表中大部分数据

若查询到的数据是表中大部分的数据，也不会使用索引
