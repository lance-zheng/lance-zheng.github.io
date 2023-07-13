<!-- customize-category:MySQL -->

# MySQL 基础

## MySQL 数据类型

### 数值类型

> <https://www.runoob.com/mysql/mysql-data-types.html>
> BIT 类型 <http://www.hushowly.com/articles/1369>

选取类型的时候应该要考虑到

| 类型                        | 字节 | 取值范围                 | 无符号取值范围 |
| --------------------------- | ---- | ------------------------ | -------------- |
| **TINYINT**                 | 1    | -128 ~ 127               | 0 ~ 255        |
| SMALLINT                    | 2    | -32768 ~ 32767           | 0 ~ 65535      |
| MEDIUMINT                   | 3    | -8388608 ~ 8388607       | 0 ~ 4294967295 |
| **INT / INTEGER**           | 4    | -2147483648 ~ 2147483647 | 0 ~ 4294967295 |
| **BIGINT**                  | 8    |                          |                |
| FLOAT                       | 4    |                          |                |
| DOUBLE                      | 8    |                          |                |
| **DEC / DECIMAL / NUMERIC** |      |                          |                |

DECIMAL: `DECIMAL(M,D)` ，如果 M>D，为 M+2 否则为 D+2

**浮点数 vs 定点数：**  
浮点数在长度一定的情况下取值范围会更大，但是会有精度问题，适用于能够接受小误差的场景。定点数没有误差，适合金融场景。

### 字符串

...

### 日期类型

...

### JSON

...

### 空间类型
