<!-- customize-category:Testing in Java -->

# JUnit5

- [JUnit5](#junit5)
  - [JUnit5 Lifecyle](#junit5-lifecyle)
  - [Custom Display Name](#custom-display-name)
  - [Ordering](#ordering)
  - [Testing Report \& Code Coverage](#testing-report--code-coverage)
  - [Conditional Tests](#conditional-tests)
  - [Parameterized Test](#parameterized-test)

> <https://junit.org/junit5/docs/current/user-guide/>

## JUnit5 Lifecyle

1. `@BeforeAll`: 所以测试函数运行前执行，只会执行一次
2. `@BefareEach`: 每个测试函数运行前执行
3. `@Test`: 测试函数
4. `@AfterEach`: 每个测试函数运行后执行
5. `@AfterAll`: 所以测试函数运行后执行，只会执行一次

## Custom Display Name

默认情况下每个测试的名称就是该函数的名称，不过这个名称可以通过一些注解实现自定义

**方法一：**`@DisplayName` 这个注解可以直接指定名称

```java
@DisplayName("Test Class")
class SpringSecurityJwtApplicationTests {
    @Test
    @DisplayName("Test Equal")
    void testEqual() {
        Assertions.assertEquals(1, 1);
    }
}
```

**方法二：** `@DisplayNameGeneration()` 这个注解可以指定生成规则

例如：`DisplayNameGenerator.ReplaceUnderscores.class` 会将下划线换成空格

```java
@DisplayNameGeneration(DisplayNameGenerator.ReplaceUnderscores.class)
class SpringSecurityJwtApplicationTests {
    @Test
    void test_Equal() {
        Assertions.assertEquals(1, 1);
    }
}

//name: test Equal
```

JUnit 默认提供了下面这些生成规则

- `DisplayNameGenerator.ReplaceUnderscores.class`
- `DisplayNameGenerator.Simple.class`
- `DisplayNameGenerator.Standard.class)`
- `DisplayNameGenerator.ReplaceUnderscores.class`

当然也可以根据自己的需求实现一个 DisplayNameGenerator

> <https://leeturner.me/posts/building-a-camel-case-junit5-displaynamegenerator/>

## Ordering

通常情况下 JUnit 的**测试结果不应该依赖于执行顺序**，每个测试函数之间不应该存在依赖关系。

不过 JUnit 也提供了一些方式修改执行顺序

通过 `@TestMethodOrder` 注解可以指定方法执行的顺序

```java
// 根据名称排序执行
@TestMethodOrder(MethodOrderer.DisplayName.class)
// 每次都随机顺序
@TestMethodOrder(MethodOrderer.Random.class)
// 根据方法名
@TestMethodOrder(MethodOrderer.MethodName.class)

// 需要和 @Order 配合使用
// 根据 @Order 的值，其中值越小的越先执行
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
```

## Testing Report & Code Coverage

代码覆盖率，是一种通过计算测试过程中被执行的源代码占全部源代码的比例

IDEA 中测试报告与代码覆盖率报告导出

<img width=400 src='/assets/image/1684508647.png'/>
<img width=400 src='/assets/image/1684508787.png'/>

> 具体可以参考 <https://www.jetbrains.com/help/idea/code-coverage.html>

---

除了使用 IDE 的功能之外还可以使用 Maven 插件完成，通过执行 `mvn test` 即可生成测试报表，这种方式比较适合 CICD。

默认情况下 maven 是不会去执行 JUnit 编写的测试代码。

需要使用这个 [maven-surefire-plugin](https://maven.apache.org/surefire/maven-surefire-plugin/) 插件，它最后会在 `${basedir}/target/surefire-reports/` 目录中生成一份测试报告，不过是 xml 格式的。这个 [maven-surefire-report-plugin](https://maven.apache.org/surefire/maven-surefire-report-plugin/) 插件可以实现生成 HTML 格式的报告

最后配置如下：

```xml
<build>
    <plugins>
        <!--
        此插件用于生成 html 格式报告
        生成文件位于：target/site/surefire-report.html
        -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-surefire-report-plugin</artifactId>
            <version>3.1.0</version>
            <executions>
                <execution>
                    <phase>test</phase>
                    <goals>
                        <goal>report</goal>
                    </goals>
                </execution>
            </executions>
        </plugin>
        <!--
        默认情况下 maven 的 test 阶段并不会运行 JUnit 的测试代码
        此插件用于运行 JUnit5 的测试用例
        -->
        <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-surefire-plugin</artifactId>
            <version>3.1.0</version>
            <configuration>
                <!--
                maven-surefire-report-plugin 在所有测试通过时才会生成 report
                加上此配置可以在出现失败 case 时也生成 report
                -->
                <testFailureIgnore>true</testFailureIgnore>

                <!-- 使用 JUnit5 中的 @DisplayName  -->
                <statelessTestsetReporter
                        implementation="org.apache.maven.plugin.surefire.extensions.junit5.JUnit5Xml30StatelessReporter">
                    <usePhrasedTestCaseMethodName>true</usePhrasedTestCaseMethodName>
                </statelessTestsetReporter>
            </configuration>
        </plugin>
    </plugins>
</build>
```

JaCoCo 是一个生成代码覆盖率报告的库

具体配置如下：

```xml
<!-- target/site/jacoco/index.html -->
<plugin>
    <groupId>org.jacoco</groupId>
    <artifactId>jacoco-maven-plugin</artifactId>
    <version>0.8.7</version>
    <executions>
        <execution>
            <id>jacoco-prepare</id>
            <goals>
                <goal>prepare-agent</goal>
            </goals>
        </execution>
        <execution>
            <id>jacoco-report</id>
            <phase>test</phase>
            <goals>
                <goal>report</goal>
            </goals>
        </execution>
    </executions>
</plugin>
```

> <https://www.baeldung.com/jacoco>

## Conditional Tests

在某些情况下跳过某些测试，或者在指定版本下运行，在指定操作系上运行，基于环境变量运行不同的测试...

**常见注解：**

- `@Disabled`
- `@DisabledOnOs`: 在指定操作系统中禁用
- `@EnabledOnJre` `@EnabledForJreRange` 在指定 Java 版本中运行
- `@EnabledIfEnvironmentVariable(named = "p1", matches = "123")`：系统环境变量
- `@EnabledIfSystemProperty(named = "p1", matches = "123")`  
  通过 JVM 参数 `-D` 指定。 e.g. `java -Dxxx=123 -jar xxx.jar`

这里只列举了一部分，还有需要以 `@Disabled` 和 `@Enabled` 开头的注解

## Parameterized Test

可以自定义一系列测试 case 然后通过参数传进来

下面这个例子是使用 Csv 格式的参数列表

```java
@ParameterizedTest(name = "num1:{0}, num2{1}, excepted:{2}")
@CsvSource({
        "1,2,3",
        "5,5,10",
        "100,-1,99",
        "-1,-1,-2",
})
void testAdd(int n1, int n2, int excepted) {
    assertEquals(excepted, demoUtils.add(n1, n2));
}
```

使用 `@ParameterizedTest` 替换 `@Test`，其中 name 是用来定制显示的名称，`@CsvSource` 是用来指定测试用例，还可以使用 `@CsvFileSource(resources = "/two-column.csv", numLinesToSkip = 1)` 从文件中读取

> 其他数据源，以及更详细的用法参考
> <https://junit.org/junit5/docs/current/user-guide/#writing-tests-parameterized-tests>

<!-- ## Unit Testing support in Spring Boot

依赖：

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-test</artifactId>
    <scope>test</scope>
</dependency>
```

然后在测试类上添加 `@SpringBootTest` 注解，这个注解会去寻找主启动类，然后还初始化 **ApplicationContext**

```java
@SpringBootTest
class SpringSecurityJwtApplicationTests {
    @Value("${server.port:8080}")
    int port;

    @Test
    void contextLoads() {
        System.out.println(port);
    }

}
```

使用 `@SpringBootTest`，后就可以在类中使用 `@Autowired`、`@Value` 等注解

## Mockito

@Mock + @InjectMocks

@MockBean + @Autoware

ReflectionTestUtils

@Sql

测试 Spring MVC Web Controller

1. Autoconfigure @AoutConfigureMockMvc
2. Inject MockMvc
3. Perform web requests
4. Define exceptations
5. Assert results

ModelAndViewAssert
MockHttpServletRequest();
new MockHttpServletResponse();

Testing restapi -->
