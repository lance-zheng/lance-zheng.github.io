<!-- customize-category:MySQL -->

# SQL 执行流程

- [SQL 执行流程](#sql-执行流程)
  - [查询缓存](#查询缓存)
    - [通过 Profiling 验证查询缓存的存在](#通过-profiling-验证查询缓存的存在)
  - [解析器](#解析器)
  - [优化器](#优化器)
  - [执行器](#执行器)
  - [存储引擎](#存储引擎)

一条 SQL 的大致执行流程图  
<img width=500 src='/assets/image/1682320478.png'/>

## 查询缓存

由于缓存命中并率不高，MySQL 8 中被移除。

```txt
select name from example;
下面这条 sql 没法使用上面的缓存，因为 sql 中多了空格
select name   from example;

同样有函数调用也没办法被缓存
select name, now()   from example;
```

MySQL 5+ 中默认关闭，需要修改配置开启 `query_cache_type`

`query_cache_type` 的取值有：

- `0`: 关闭
- `1`：开启
- `2`：按需使用`SELECT SQL_CACHE ...`

### 通过 Profiling 验证查询缓存的存在

开启 Profiling

```sql
-- 获取当前 profiling 开启状态
select @@profiling;

-- 在当前会话开启 profiling
set profiling=1;
```

获取被记录的 profile `show profiles`

```sql
show profiles;
+----------+------------+-------------------------+
| Query_ID | Duration   | Query                   |
+----------+------------+-------------------------+
|       28 | 0.00077600 | select * from example   |
|       29 | 0.00077600 | select * from example   |
|       30 | 0.00367200 | select *   from example |
+----------+------------+-------------------------+
```

获取某条 SQL 的执行细节 `show profile for query [query_id]`

```sql
-- 查询 Query_ID 为 28 的 SQL 执行情况
show profile for query 28;
+--------------------------------+----------+
| Status                         | Duration |
+--------------------------------+----------+
| starting                       | 0.000206 |
| Waiting for query cache lock   | 0.000031 |
| starting                       | 0.000011 |
| checking query cache for query | 0.000322 |
| checking permissions           | 0.000051 |
| Opening tables                 | 0.001024 |
| init                           | 0.000187 |
| System lock                    | 0.000514 |
| Waiting for query cache lock   | 0.000021 |
| System lock                    | 0.001964 |
| optimizing                     | 0.000033 |
| statistics                     | 0.000125 |
| preparing                      | 0.000106 |
| executing                      | 0.000016 |
| Sending data                   | 0.002173 |
| end                            | 0.000042 |
| query end                      | 0.000099 |
| closing tables                 | 0.000072 |
| freeing items                  | 0.000078 |
| Waiting for query cache lock   | 0.000016 |
| freeing items                  | 0.001018 |
| Waiting for query cache lock   | 0.000041 |
| freeing items                  | 0.000047 |
| storing result in query cache  | 0.000486 |
| cleaning up                    | 0.000152 |
+--------------------------------+----------+
-- 查询 Query_ID 为 29 的 SQL 执行情况(命中缓存)
show profile for query 29;
+--------------------------------+----------+
| Status                         | Duration |
+--------------------------------+----------+
| starting                       | 0.000140 |
| Waiting for query cache lock   | 0.000020 |
| starting                       | 0.000009 |
| checking query cache for query | 0.000428 |
| checking privileges on cached  | 0.000184 |
| checking permissions           | 0.000234 |
| sending cached result to clien | 0.000445 |
| cleaning up                    | 0.000032 |
+--------------------------------+----------+
```

`show profile` 还可以查看 CPU ,磁盘 IO 等多种资源使用情况

```txt
SHOW PROFILE [type [, type] ... ]
    [FOR QUERY n]
    [LIMIT row_count [OFFSET offset]]

type: {
    ALL
  | BLOCK IO
  | CONTEXT SWITCHES
  | CPU
  | IPC
  | MEMORY
  | PAGE FAULTS
  | SOURCE
  | SWAPS
}
```

> <https://dev.mysql.com/doc/refman/8.0/en/show-profile.html?spm=a2c4g.204734.0.0.2d7629b1MdS2Uk>

## 解析器

包括**词法分析**和**语法分析**

## 优化器

选择合适的索引、决定联表的顺序、找出较优的查询方案。

## 执行器

开始执行的时候，要先判断一下你对这个表 T 有没有执行查询的**权限**，如果没有，就会返回没有权限的错误。

有权限，会根据优化后的 SQL，**向存储引擎发起查询操作**，并且返回查询的结果。

## 存储引擎
