<!-- customize-category:JVM -->

# 一个 Object 占用多少字节？

> 以下所有内容都是基于 JDK17 的 64 位 HotSpot 虚拟机中测试

在 64 位 **HotSpot** 虚拟机中一个对象是由 3 部分组成，**对象头**，**实例数据**，**对齐填充**

<img width=200 src='/assets/image/1681741126.png'/>

可以使用 openjdk 中提供的工具查看对象的大小

```xml
<dependency>
    <groupId>org.openjdk.jol</groupId>
    <artifactId>jol-core</artifactId>
    <version>0.17</version>
</dependency>
```

```java
public class Main {
    public static void main(String[] args) {
        System.out.println(ClassLayout.parseInstance(new Object()).toPrintable());
    }
}
```

## 对象头

在 64 位虚拟机中 mark word 占 8 个字节 class pointer 占 4 个字节总共 12 个字节

若关闭类指针压缩的话 class pointer 将会占用 8 个字节(`-XX:-UseCompressedClassPointers`) 此时对象头占 16 个字节

```txt
java.lang.Object object internals:
OFF  SZ   TYPE DESCRIPTION               VALUE
  0   8        (object header: mark)     0x0000000000000001 (non-biasable; age: 0)
  8   4        (object header: class)    0x00000f68
 12   4        (object alignment gap)
Instance size: 16 bytes


关闭类指针压缩 -XX:-UseCompressedClassPointers

java.lang.Object object internals:
OFF  SZ   TYPE DESCRIPTION               VALUE
  0   8        (object header: mark)     0x0000000000000001 (non-biasable; age: 0)
  8   8        (object header: class)    0x00000001129bdcb0
Instance size: 16 bytes
```

Mark word 在不同锁状态下每个 bit 的含义

```txt
|-----------------------------------------------------------------------------------------------------------------|
|                                             Object Header(128bits)                                              |
|-----------------------------------------------------------------------------------------------------------------|
|                                   Mark Word(64bits)               |  Klass Word(64bits)    |      State         |
|-----------------------------------------------------------------------------------------------------------------|
| unused:25|identity_hashcode:31|unused:1|age:4|biase_lock:1|lock:2 | OOP to metadata object |      Nomal         |
|-----------------------------------------------------------------------------------------------------------------|
| thread:54|      epoch:2       |unused:1|age:4|biase_lock:1|lock:2 | OOP to metadata object |      Biased        |
|-----------------------------------------------------------------------------------------------------------------|
|                     ptr_to_lock_record:62                 |lock:2 | OOP to metadata object | Lightweight Locked |
|-----------------------------------------------------------------------------------------------------------------|
|                    ptr_to_heavyweight_monitor:62          |lock:2 | OOP to metadata object | Heavyweight Locked |
|-----------------------------------------------------------------------------------------------------------------|
|                                                           |lock:2 | OOP to metadata object |    Marked for GC   |
|-----------------------------------------------------------------------------------------------------------------|
```

## 对齐填充

64 位 HotSpot 虚拟机要求对象的大小必须是 8 的倍数。

## 数组的大小

如果是数组对象还需要额外花费 4 字节来记录长度，所以在开启类指针压缩时占用 `8 + 4 + 4 = 16`，关闭时占用 `8 + 8 + 4 + 4 = 24`字节。

```txt
OFF  SZ   TYPE DESCRIPTION               VALUE
  0   8        (object header: mark)     0x0000000000000001 (non-biasable; age: 0)
  8   8        (object header: class)    0x0000000113424b98
 16   4        (array length)            0
 20   4        (alignment/padding gap)
 24   0    int [I.<elements>             N/A
Instance size: 24 bytes
```

## Integer

开启类指针压缩时占 16 字节
其中对象头 12，value 属性占 4 个字节
