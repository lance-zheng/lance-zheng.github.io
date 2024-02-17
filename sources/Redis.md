# Redis

- [Redis](#redis)
  - [使用 Docker 搭建主从集群](#使用-docker-搭建主从集群)
  - [布隆过滤器](#布隆过滤器)
    - [Redis 实现](#redis-实现)
    - [Java 内存实现](#java-内存实现)

## 使用 Docker 搭建主从集群

```sh
...
```

## 布隆过滤器

[wikipedia](https://zh.wikipedia.org/zh-cn/%E5%B8%83%E9%9A%86%E8%BF%87%E6%BB%A4%E5%99%A8)

> 它实际上是一个很长的二进制向量和一系列随机映射函数。布隆过滤器可以用于检索一个元素是否在一个集合中。它的优点是空间效率和查询时间都远远超过一般的算法，缺点是有一定的误识别率和删除困难。

优点：

- 查询快，占用空间小
- 不存储元素数据，安全性高

缺点：

- 没办法判断元素是否一定存在，只能判断出元素一定不存在
- 没法判断元素是否一定存在，所以也很难删除

### Redis 实现

Redis 布隆过滤器插件
<https://github.com/RedisBloom/RedisBloom>

```sh
# 创建
# 不创建的话默认 为 0.01 100
BF.RESERVE mykey 0.1 10000000
# 添加
BF.ADD mykey 1
# 判断是否存在
BF.EXISTS mykey 1 # return 1
BF.EXISTS mykey 2 # return 0
# 添加多个 key
BF.MADD mykey 1 2 3
# 判断多个 key 是否存在
BF.MEXISTS mykey 1 2 5
```

Java API 操作

```xml
<dependency>
    <groupId>org.redisson</groupId>
    <artifactId>redisson-spring-boot-starter</artifactId>
    <version>3.16.7</version>
</dependency>
```

```java
package com.github.lance.zheng.redis;

import org.redisson.Redisson;
import org.redisson.api.RBloomFilter;
import org.redisson.api.RedissonClient;
import org.redisson.config.Config;

public class RestApplication {
    public static void main(String[] args) {
        private Integer expectedInsertions = 10000;
        private Double fpp = 0.01;
        Config config = new Config();
        config.useSingleServer().setAddress("redis://localhost:6379");
        RedissonClient client = Redisson.create(config);
        RBloomFilter<Object> bloomFilter = client.getBloomFilter("user");
        bloomFilter.tryInit(expectedInsertions, fpp);

        for (Integer i = 0; i < expectedInsertions; i++) {
            bloomFilter.add(i);
        }

        int count = 0;
        for (int i = expectedInsertions; i < expectedInsertions * 2; i++) {
            if (bloomFilter.contains(i)) {
                count++;
            }
        }

        System.out.println("误判次数" + count);
    }
}
```

### Java 内存实现

```xml
<!-- pom -->
<dependency>
  <groupId>com.google.guava</groupId>
  <artifactId>guava</artifactId>
  <version>29.0-jre</version>
</dependency>
```

```java
package com.github.lance.zheng.redis;
import com.google.common.hash.BloomFilter;
import com.google.common.hash.Funnels;

public class BF {
    public static void main(String[] args) {
        final int expectedInsertions = 10000000;
        final double fpp = 0.01;
        BloomFilter<Integer> bf = BloomFilter.create(Funnels.integerFunnel(), expectedInsertions, fpp);
        for (int i = 0; i < expectedInsertions; i++) {
            bf.put(i);
        }
        int c = 0;
        for (int i = expectedInsertions; i < expectedInsertions * 2; i++) {
            if (bf.mightContain(i)) {
                c++;
            }
        }

        System.out.println("误判率：" + c * 1.0 / expectedInsertions);
    }
}
```

- expectedInsertions：预期元素个数
- fpp：误判率 ( < 1.0)
