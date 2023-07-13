<!-- customize-category:Java -->

# Java 9 到 Java 17 新特性

`**` 表示表示实际开发中比较有用的功能

- [Java 9 到 Java 17 新特性](#java-9-到-java-17-新特性)
  - [Java 9](#java-9)
    - [\*\*集合工厂方法 List.of(...)](#集合工厂方法-listof)
    - [\*\*新的 Stream API](#新的-stream-api)
    - [\*\*接口私有方法](#接口私有方法)
    - [\*\* 改进 try-with-resources](#-改进-try-with-resources)
    - [模块化](#模块化)
    - [JShell](#jshell)
  - [Java 10](#java-10)
    - [局部类型推导](#局部类型推导)
    - [\*\*API 更新](#api-更新)
  - [Java 11](#java-11)
    - [\*\*String API 增强](#string-api-增强)
    - [\*\*File API](#file-api)
  - [Java 12](#java-12)
    - [\*\*Switch 语法糖](#switch-语法糖)
    - [\*\* Compact Number](#-compact-number)
    - [文件对比 Files.mismatch](#文件对比-filesmismatch)

## Java 9

> <https://www.runoob.com/java/java9-inner-class-diamond-operator.html>  
> <https://www.wdbyte.com/2020/02/jdk/jdk9-feature/>

- **模块系统**：模块是一个包的容器，Java 9 最大的变化之一是引入了模块系统（Jigsaw 项目）。
- **REPL (JShell)**：交互式编程环境。
- **HTTP 2 客户端**：HTTP/2 标准是 HTTP 协议的最新版本，新的 HTTPClient API 支持 WebSocket 和 HTTP2 流以及服务器推送特性。
- **改进的 Javadoc**：Javadoc 现在支持在 API 文档中的进行搜索。另外，Javadoc 的输出现在符合兼容 HTML5 标准。
- **多版本兼容 JAR 包**：多版本兼容 JAR 功能能让你创建仅在特定版本的 Java 环境中运行库程序时选择使用的 class 版本。
- **集合工厂方法**：List，Set 和 Map 接口中，新的静态工厂方法可以创建这些集合的不可变实例。
- **私有接口方法**：在接口中使用 private 私有方法。我们可以使用 private 访问修饰符在接口中编写私有方法。
- **进程 API**: 改进的 API 来控制和管理操作系统进程。引进 java.lang.ProcessHandle 及其嵌套接口 Info 来让开发者逃离时常因为要获取一个本地进程的 PID 而不得不使用本地代码的窘境。
- **改进的 Stream API**：改进的 Stream API 添加了一些便利的方法，使流处理更容易，并使用收集器编写复杂的查询。
- **改进 try-with-resources**：如果你已经有一个资源是 final 或等效于 final 变量,您可以在 try-with-resources 语句中使用该变量，而无需在 try-with-resources 语句中声明一个新变量。
- **改进的弃用注解 @Deprecated**：注解 @Deprecated 可以标记 Java API 状态，可以表示被标记的 API 将会被移除，或者已经破坏。
- **改进钻石操作符(Diamond Operator)**：匿名类可以使用钻石操作符(Diamond Operator)。 -**改进 Optional 类**：java.util.Optional 添加了很多新的有用方法，Optional 可以直接转为 stream。
- **多分辨率图像 API**：定义多分辨率图像 API，开发者可以很容易的操作和展示不同分辨率的图像了。
- **改进的 CompletableFuture API**： CompletableFuture 类的异步机制可以在 ProcessHandle.onExit 方法退出时执行操作。
- **轻量级的 JSON API**：内置了一个轻量级的 JSON API
- **响应式流（Reactive Streams) API**: Java 9 中引入了新的响应式流 API 来支持 Java 9 中的响应式编程。

### \*\*集合工厂方法 List.of(...)

在 Java 9 中新添加了静态工厂方法，用于创建**只读集合**，里面的对象**不可改变**，并在不能存在 **null** 值，对与 `Set` 和 `Map` 中**key 也不能重复**。

`List.of()`、`Set.of()`、`Map.of()`、`Map.ofEntries()`

```java
List.of("a", "b", null); // NullPointerExceptio
Set.of("1", "1"); // IllegalArgumentException: duplicate element
Map.of("key1", "v1", "key1", "v2"); // IllegalArgumentException: duplicate key
```

创建的集合为不可变集合，不能添加和修改元素。

```java
List<String> list = List.of("a");
list.add("a");    // java.lang.UnsupportedOperationException
list.set(0, "b"); // java.lang.UnsupportedOperationException
```

**List.of vs Arrays.asList：**

1. **List.of** 创建的是不可变的集合，**Arrays.asList** 是可以修改的

   ```java
    String[] data = {"a", "b"};
    List<String> of = List.of(data);
    List<String> asList = Arrays.asList(data);
    data[0] = "c"; // 尝试修改源数据
    System.out.println(of);     // [a, b]
    System.out.println(asList); // [c, b]

    of.set(0, "c"); // Fails with UnsupportedOperationException
    asList.set(0, "c");
   ```

2. **List.of** 不能存储空值，**Arrays.asList** 可以

   ```java
    List.of("a", null);      // NullPointerException
    Arrays.asList("a", null);
   ```

3. **contains** 函数区别

   ```java
    List.of("a").contains(null); // NullPointerException
    Arrays.asList("a").contains(null);
   ```

在 Java8 中也可以通过 `Collections.unmodifiableList()` 创建不可变集合

### \*\*新的 Stream API

在 Java9 中对 Stream API 进行了增强添加了 4 个新的 API

`takeWhile`、`dropWhile`、`ofNullable`、`iterate`

1. **takeWhile**：直到不满足条件

   ```java
   List<Integer> list = Stream.of(1, 2, 3, 2, 5)
           .takeWhile(n -> n < 3).collect(Collectors.toList());
   System.out.println(list); // [1, 2]
   ```

2. **dropWhile**: 删除元素直达不满足条件

   ```java
     List<Integer> list = Stream.of(1, 2, 3, 2, 5)
             .dropWhile(n -> n < 3).collect(Collectors.toList());
     System.out.println(list); // [3, 2, 5]
   ```

3. **ofNullable**：创建支持 null 值的 stream

   ```java
   Stream.of(null).forEach(System.out::println); // java.lang.NullPointerException
   Stream.ofNullable(null).forEach(System.out::println);
   ```

4. **iterate**：它支持谓词（条件）作为第二个参数，并且如果谓词为 false， stream.iterate 将停止。

   ```java
   Stream.iterate(1, n -> n < 20, n -> n * 2).forEach(System.out::print);
   // 124816
   ```

   这个方法倒是从来没用过...

### \*\*接口私有方法

在 Java8 中接口方法可以定义默认实现，Java9 中支持了定义私有方法

```java
interface Foo {
    default void foo() {
        defaultImpl();
    }

    default void baz() {
        defaultImpl();
    }

    private void defaultImpl() {
        System.out.println();
    }
}
```

能想到的场景就是多个方法的默认实现是一样时，可以通私有方法将其抽取出来。

### \*\* 改进 try-with-resources

在 Java 8 若对象实现了 `AutoCloseable` 接口则可以使用 try-with-resources 进行自动关闭资源

但只有在 `try(...)` 中定义的资源才支持

```java
FileInputStream is = new FileInputStream("");
try (FileInputStream a = is) {
    a.read();
}
```

在 Java9 中解除了这个限制，但引用必须是 final 的才可以

```java
// 不写 final 也行，只要没有别的地方在修改
final FileInputStream is = new FileInputStream("");
try (is) {
    is.read();
}
```

### 模块化

JAVA 9 支持模块化，最明显的就是在 JDK8 中的 jre 目录消失了，改为 jmods 目录，不过感觉在开发中没啥用。

### JShell

交互式的编程环境，向 Chorme 中的 Console 那样。

```shell
jshell> String a = "aaa"
a ==> "aaa"

jshell> System.out.println(a)
aaa

jshell>
```

## Java 10

> <https://www.wdbyte.com/2020/02/jdk/jdk10-feature/>

### 局部类型推导

Java 10 增加了局部变量推导功能，定义局部变量时可以不指定类型，让编译器自动推断数据类型。

`var` 关键字使用有很多限制：

1. 只能用于带有初始化值的局部变量中
   `var a = null` 这个是无法通过编译的
2. 用于 `for` 循环或增强 `for` 中变量声明

```java
public static void main(String[] args) {
   // var a = null; compile error

   var str = "";
   var map = new HashMap<String, Object>();

   for (var i = 0; i < args.length; i++) {
      System.out.println(args[i]);
   }

   for (var c : args) {
      System.out.println(c);
   }
}
```

### \*\*API 更新

`List.copyOf`、`Collectors.toUnmodifiableList`

```java
public static void main(String[] args) {
   List<String> data = new ArrayList<>();
   data.add("1");
   data.add("2");

   List<String> copyOfList = List.copyOf(data);
   List<String> collect = data.stream().collect(Collectors.toUnmodifiableList());
   data.add("3");

   // [1, 2]
   System.out.println(copyOfList);
   System.out.println(collect);

   // java.lang.UnsupportedOperationException
   collect.add("3");
   copyOfList.add("3");
}
```

## Java 11

> <https://www.wdbyte.com/2020/03/jdk/jdk11-feature/>

### \*\*String API 增强

- `String.isBlank()`：判断字符串是否为空
- `String.repeat(int)`：重复字符串
- `String.trim()`：去除前后半角空格
- `String.strip()`：去除前后空格
- `String.lines()`：以换行符分割返回一个 Stream

```java
System.out.println(" ".isBlank());     // true
System.out.println("a".repeat(3));     // aaa
System.out.println("      a ".trim()); // a
Stream<String> lines = "a\nb".lines();
```

### \*\*File API

读写文件变得更加方便 `Files.writeString`、`Files.readString`

```java
Path path = Files.writeString(Files.createTempFile("test", ".txt"), "java 11");
System.out.println(Files.readString(path));
```

## Java 12

> <https://www.wdbyte.com/2020/02/jdk/jdk12-feature/>

### \*\*Switch 语法糖

使用新的语法糖 `->` 可以不需要写 `break` 语句，同时新的 `switch` 还支持返回。值

之前的写法

```java
public static void main(String[] args) throws IOException {
   int month = 1;
   switch (month) {
      case 3:
      case 4:
      case 5:
            System.out.println("春季");
            break;
      case 6:
      case 7:
      case 8:
            System.out.println("夏季");
            break;
      default:
            System.out.println("error");
   }
}
```

Java 12 写法

```java
public static void main(String[] args) throws IOException {
   String month = "九月";
   String season = switch (month) { // 带返回值
      case "三月", "四月", "五月" -> "春季";
      case "六月", "七月", "八月" -> "夏季";
      case "九月", "十月", "十一月" -> "秋季";
      case "十二月", "一月", "二月" -> "秋季";
      default -> "错误";
   };

   System.out.println(season);
}
```

### \*\* Compact Number

简化数字格式，例如：`1000 -> 1千`

```java
public static void main(String[] args) throws IOException {
   NumberFormat compactNumberInstance = NumberFormat.getCompactNumberInstance(Locale.CHINA, NumberFormat.Style.LONG);
   System.out.println(compactNumberInstance.format(1000000));

   compactNumberInstance.setMaximumFractionDigits(2);
   System.out.println(compactNumberInstance.format(12345678));
}
```

输出如下

```txt
100万
1234.57万
```

### 文件对比 Files.mismatch

`Files.mismatch` 可以返回两个文件不同字符的起始位置。若两个文件相同则返回 `-1`

```java
public static void main(String[] args) throws IOException {
   Path file1 = Files.createTempFile("file1", ".txt");
   Path file2 = Files.createTempFile("file2", ".txt");
   Files.writeString(file1, "abc", StandardOpenOption.WRITE);
   Files.writeString(file2, "abc", StandardOpenOption.WRITE);

   System.out.println(Files.mismatch(file1, file2));// -1

   // 追加字符
   Files.writeString(file1, "d", StandardOpenOption.APPEND);
   Files.writeString(file2, "e", StandardOpenOption.APPEND);
   System.out.println(Files.mismatch(file1, file2)); // 3
}
```
