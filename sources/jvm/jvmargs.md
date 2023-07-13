<!-- customize-category:JVM -->

# 常见 JVM 调优参数

JVM 参数分为三种：

- `标准参数（-）`：每个 JVM 的版本都可以使用。例如 `-cp`
- `-X`：不保证所有 JVM 都有实现，不保证向后兼用。`-Xmx`
- `-XX`：各个 JVM 实现会有所不同，将来可能会随时取消。

> <https://www.oracle.com/java/technologies/javase/vmoptions-jsp.html#Options>  
> <https://docs.oracle.com/en/java/javase/17/docs/specs/man/java.html>

**常见参数：**

- `-Xmx`：最大堆容量
- `-Xms`：初始堆容量
- `-Xss`：堆栈大小
- `-Xmn`：新生代出生大小
- `-XX:MetaspaceSize`：元空间出生大小
- `-XX:MaxMetaspaceSize`：元空间最大值 `-XX:MaxMetaspaceSize=500M`
- `-XX:NewRatio=2`：新生代与老年代的比例，2 表示新生代占 1/3，老年代占 2/3。
- `-XX:SurvivorRatio=8`：设置年轻代中 Eden 区与 Survivor 区的大小比值

- `-XX:+UseStringDeduplication`: 启用字符串去重，减少内存占用。

  > <https://mp.weixin.qq.com/s/uPnoDctc_IJRbYxDkmE1-A>

- `-XX:+UseCompressedOops`: 启用普通对象指针压缩，以减少对象头的大小。
- `-XX:+UseCompressedClassPointers`: 启用压缩类指针，减少元数据的大小。
