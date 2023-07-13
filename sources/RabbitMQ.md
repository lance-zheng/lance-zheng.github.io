# RabbitMQ

- [RabbitMQ](#rabbitmq)
  - [AMQP](#amqp)
  - [Installation Guides](#installation-guides)
    - [Docker](#docker)
    - [Yum](#yum)
  - [CLI](#cli)
    - [rabbitmqctl](#rabbitmqctl)
    - [rabbitmq-plugins](#rabbitmq-plugins)
  - [消息模型](#消息模型)
    - ["Hello World!"](#hello-world)
    - [Work Queues](#work-queues)
    - [Publish/Subscribe](#publishsubscribe)
    - [Direct](#direct)
    - [Topics](#topics)
    - [RPC](#rpc)
  - [SpringBoot 整合](#springboot-整合)
  - [集群](#集群)
    - [普通集群](#普通集群)
    - [镜像集群](#镜像集群)

[RabbitMQ](https://www.rabbitmq.com/) is the most widely deployed open source message broker.

## AMQP

AMQP 的全称为：Advanced Message Queuing Protocol（高级消息队列协议），它是一个开放标准，可以在各种语言和平台之间使用。

## Installation Guides

### Docker

`rabbitmq:management` 镜像默认开启了 web 管理页面。
如果使用的是 `rabbitmq` 镜像后续可以通过 `rabbitmq-plugins enable rabbitmq_management` 指令启动 web 管理页面。

```sh
docker run -id --hostname my-rabbitmq --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:management
```

安装完成后浏览器访问 <http://localhost:15672> 可以进入 web 管理页面，默认账号密码是 `guest`

### Yum

**安装 Erlang：**

```sh
yum -y install epel-release

yum -y update

yum -y install erlang socat

erl -version
```

erlang: <https://packagecloud.io/rabbitmq/erlang/install#bash-rpm>
rabbitmq: <https://github.com/rabbitmq/rabbitmq-server/releases/tag/v3.8.30>

## CLI

### rabbitmqctl

<https://www.rabbitmq.com/rabbitmqctl.8.html>
...

```sh
# 创建虚拟主机（虚拟主机类似于 MySQL 中的数据库）
rabbitmqctl add_vhost spm
# 用户管理
rabbitmqctl add_user username password
rabbitmqctl delete_user username
# 配置用户的权限
rabbitmqctl set_permissions -p vhost username ".*" ".*" ".*"
```

### rabbitmq-plugins

RabbitMQ 命令行插件管理器
<https://www.rabbitmq.com/rabbitmq-plugins.8.html>

**插件列表：**

```sh
rabbitmq-plugins list [-v] [-m] [-E] [-e] [pattern]

-v 显示所有插件的详情（详细）
-m 仅仅只显示插件的名称 (简约)
-E 仅仅只显示显式启用的插件
-e 仅仅只显示显式、隐式启用的插件
```

**开启插件 & 禁用插件：**

```sh
# 启用插件及其依赖
rabbitmq-plugins enable {plugin ...}
# 关闭插件及其依赖
rabbitmq-plugins disable {plugin ...}

# 开启指定插件同时禁用其他插件
rabbitmq-plugins set [plugin ...]
```

## 消息模型

### "Hello World!"

**依赖:**

```xml
<dependency>
    <groupId>com.rabbitmq</groupId>
    <artifactId>amqp-client</artifactId>
    <version>5.16.0</version>
</dependency>
```

一对个生产者和一个消费者，直接将消息发送到指定队列。

![img](https://www.rabbitmq.com/img/tutorials/python-one.png)

```java
public class HelloWorld {
    static ConnectionFactory connectionFactory;

    static {
        // 创建连接工厂
        connectionFactory = new ConnectionFactory();
        connectionFactory.setHost("127.0.0.1");
        connectionFactory.setPort(5672);
        connectionFactory.setVirtualHost("spm_qa");
        connectionFactory.setUsername("spm");
        connectionFactory.setPassword("p");
    }

    public static void main(String[] args) throws Exception {
        customer();
        provider();
    }

    private static void provider() throws IOException, TimeoutException {
        Connection connection = getConnection();
        Channel channel = connection.createChannel();
        // 声明一个队列，如不存在会自动创建
        channel.queueDeclare("spm.user", false, false, false, null);
        channel.basicPublish("", "spm.user", null, "Hello RabbitMQ!".getBytes(StandardCharsets.UTF_8));
        channel.close();
        connection.close();
    }

    private static void customer() throws IOException, TimeoutException {
        Connection connection = getConnection();
        Channel channel = connection.createChannel();
        // 声明一个队列，如不存在会自动创建
        channel.queueDeclare("spm.user", false, false, false, null);
        channel.basicConsume("spm.user", true, new DefaultConsumer(channel) {
            @Override
            public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                System.out.println("收到消息：" + new String(body));
            }
        });
    }

    private static Connection getConnection() throws IOException, TimeoutException {
        return connectionFactory.newConnection();
    }
}
```

**queueDeclare 函数参数说明：**

```java
queueDeclare(
    String queue,                   // 队列名称
    boolean durable,                // 队列是否持久化，只是持久化这个队列，并不持久化里面的消息
    boolean exclusive,              // 是否独占，只有当前连接可以使用此队列
    boolean autoDelete,             // 是否自动删除，当队列中的消息被消费完同时没有 Customer 时会自动删除
    Map<String, Object> arguments
);
```

### Work Queues

多个消费者处理同个 Queue

![imgs](https://www.rabbitmq.com/img/tutorials/python-two.png)
下面代码创建了两个工人同时处理 `send_doc` 队列中的消息

```java
public class WorkQueues {
    static ConnectionFactory connectionFactory;
    static String QUEUE_NAME = "send_doc";
    static Runnable WORKER = () -> {
        String name = Thread.currentThread().getName();
        try {
            System.out.println(name + " started");
            Connection connection = getConnection();
            Channel channel = connection.createChannel();
            channel.queueDeclare(QUEUE_NAME, true, false, false, null);
            boolean autoAck = true
            channel.basicConsume(QUEUE_NAME, autoAck, new DefaultConsumer(channel) {
                @Override
                public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                    System.out.println(name + " received message: " + new String(body));
                }

            });
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    };

    static {
        // 创建连接工厂
        connectionFactory = new ConnectionFactory();
        connectionFactory.setHost("127.0.0.1");
        connectionFactory.setPort(5672);
        connectionFactory.setVirtualHost("spm_qa");
        connectionFactory.setUsername("spm");
        connectionFactory.setPassword("p");
    }

    public static void main(String[] args) throws Exception {
        // Customer
        new Thread(WORKER, "Worker 01").start();
        new Thread(WORKER, "Worker 02").start();

        // Provider
        Connection connection = getConnection();
        Channel channel = connection.createChannel();
        channel.queueDeclare(QUEUE_NAME, true, false, false, null);

        for (int i = 0; i < 10; i++) {
            String msg = i + " Hello RabbitMQ!";
            channel.basicPublish("", QUEUE_NAME, MessageProperties.PERSISTENT_TEXT_PLAIN, msg.getBytes(StandardCharsets.UTF_8));
        }
    }


    public static Connection getConnection() throws IOException, TimeoutException {
        return connectionFactory.newConnection();
    }
}
```

**输出：**

```log
Worker 02 received message: 1 Hello RabbitMQ!
Worker 02 received message: 3 Hello RabbitMQ!
Worker 02 received message: 5 Hello RabbitMQ!
Worker 02 received message: 7 Hello RabbitMQ!
Worker 02 received message: 9 Hello RabbitMQ!
Worker 01 received message: 0 Hello RabbitMQ!
Worker 01 received message: 2 Hello RabbitMQ!
Worker 01 received message: 4 Hello RabbitMQ!
Worker 01 received message: 6 Hello RabbitMQ!
Worker 01 received message: 8 Hello RabbitMQ!
```

上述代码存在一些问题：

- 默认情况下 RabbitMQ 会将消息按顺序发送给消费者，每个消费者平均收到的消息数量是一样的。但是每个消费者的处理速度可能不一样。
- 上面代码使用了 `autoAck=true` 进行自动确认消息，被确认的消息将会从队列中移除。消息被确认但不一定已经被消费了，如果 server 在此时宕机会有可能造成数据丢失。

上面的两个问题可以通过配置下面两个参数解决：

- `channel.basicQos(number)`: 告诉 RabbitMQ 在消息没被确认前最多能接受多少条消息
- `channel.basicAck(deliveryTag)`: 手动确认消息

```java
static Runnable WORKER = () -> {
    String name = Thread.currentThread().getName();
    try {
        System.out.println(name + " started");
        Connection connection = getConnection();
        Channel channel = connection.createChannel();
        channel.queueDeclare(QUEUE_NAME, true, false, false, null);
        // 告诉 mq 在消息没被确认前最多接收多少消息
        channel.basicQos(1);
        boolean autoAck = false;
        channel.basicConsume(QUEUE_NAME, autoAck, new DefaultConsumer(channel) {
            @Override
            public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                System.out.println(name + " received message: " + new String(body));
                // 手动确认
                channel.basicAck(envelope.getDeliveryTag(), false);
            }

        });
    } catch (Exception e) {
        throw new RuntimeException(e);
    }
};
```

### Publish/Subscribe

`Fanout` 交换机，广播类型的交换机将会把消息发送给绑定到此交换机上的所有队列  
![img](https://www.rabbitmq.com/img/tutorials/exchanges.png)  
下面代码创建了一个临时队列来接受消息`channel.queueDeclare().getQueue()`，同时使用 `channel.queueBind(queue, EXCHANGE, "")` 将此队列绑定到了交换机上。(在`fanout`类型的交换机上 `routingkey` 没有意义)

```java
public class Fanout {
    static ConnectionFactory connectionFactory;
    static String EXCHANGE = "outbound_msg";
    static Runnable WORKER = () -> {
        try {
            Connection connection = connectionFactory.newConnection();
            Channel channel = connection.createChannel();

            // 创建一个临时的队列
            String queue = channel.queueDeclare().getQueue();
            channel.queueBind(queue, EXCHANGE, "");

            channel.basicConsume(queue, true, new DefaultConsumer(channel) {
                @Override
                public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                    System.out.println(queue + " 收到消息: " + new String(body));
                }
            });
        } catch (IOException | TimeoutException e) {
            throw new RuntimeException(e);
        }

    };

    static {
        connectionFactory = new ConnectionFactory();
        connectionFactory.setHost("127.0.0.1");
        connectionFactory.setPort(5672);
        connectionFactory.setVirtualHost("spm_qa");
        connectionFactory.setUsername("spm");
        connectionFactory.setPassword("p");
    }

    public static void main(String[] args) throws Exception {
        new Thread(WORKER, "Worker 01").start();
        new Thread(WORKER, "Worker 02").start();

        Thread.sleep(500);
        Connection connection = connectionFactory.newConnection();
        Channel channel = connection.createChannel();
        channel.exchangeDeclare(EXCHANGE, "fanout");
        // 广播消息
        // Fanout 下 routingKey 没有作用
        channel.basicPublish(EXCHANGE, "", null, "hello".getBytes(StandardCharsets.UTF_8));
    }
}
```

### Direct

直连交换机 `(type=direct)`：routingKey 必须完全匹配。
在前面的 HelloWord 中就是使用的这种，因为每个 Queue 都会隐式的绑定到一个默认的交换机`""`，同时他们的 `routingKey`就是队列名称。

```java

public class Direct {
    static ConnectionFactory connectionFactory;
    static String EXCHANGE_NAME = "logs";

    static {
        connectionFactory = new ConnectionFactory();
        connectionFactory.setHost("127.0.0.1");
        connectionFactory.setPort(5672);
        connectionFactory.setVirtualHost("spm_qa");
        connectionFactory.setUsername("spm");
        connectionFactory.setPassword("p");
    }

    public static void main(String[] args) throws IOException, TimeoutException, InterruptedException {
        new Thread(() -> {
            try {
                Connection connection = connectionFactory.newConnection();
                Channel channel = connection.createChannel();
                channel.exchangeDeclare(EXCHANGE_NAME, "direct");
                String queue = channel.queueDeclare().getQueue();
                channel.queueBind(queue, EXCHANGE_NAME, "error");
                channel.basicConsume(queue, true, new DefaultConsumer(channel) {
                    @Override
                    public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                        System.out.println("info " + new String(body));
                    }
                });
            } catch (IOException | TimeoutException e) {
                throw new RuntimeException(e);
            }
        }).start();

        new Thread(() -> {
            try {
                Connection connection = connectionFactory.newConnection();
                Channel channel = connection.createChannel();
                channel.exchangeDeclare(EXCHANGE_NAME, "direct");
                String queue = channel.queueDeclare().getQueue();
                channel.queueBind(queue, EXCHANGE_NAME, "info");
                channel.queueBind(queue, EXCHANGE_NAME, "error");
                channel.basicConsume(queue, true, new DefaultConsumer(channel) {
                    @Override
                    public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                        System.out.println("info | error " + new String(body));
                    }
                });
            } catch (IOException | TimeoutException e) {
                throw new RuntimeException(e);
            }
        }).start();

        Thread.sleep(500);

        Connection connection = connectionFactory.newConnection();
        Channel channel = connection.createChannel();
        channel.exchangeDeclare(EXCHANGE_NAME, "direct");
        channel.basicPublish(EXCHANGE_NAME, "info", null, "info message".getBytes(StandardCharsets.UTF_8));
        channel.basicPublish(EXCHANGE_NAME, "error", null, "error message".getBytes(StandardCharsets.UTF_8));
        channel.close();
        connection.close();
    }
}

```

### Topics

Topic 类型的交换机`(type=topic)`允许 `routingKey` 通过一些通配符模糊匹配  
![img](https://www.rabbitmq.com/img/tutorials/python-five.png)
**支持的通配符：**

- `#` 匹配 0 ~ n 个单词
- `*` 匹配 1 个单词

RabbitMQ 推荐用 `.` 分隔单词 例如：`user.*.bbb` `*.aaa.bbb` `user.aaa.#`  
**代码实例：**

```java
public class Topic {
    static ConnectionFactory connectionFactory;
    static String EXCHANGE_NAME = "category";

    static {
        connectionFactory = new ConnectionFactory();
        connectionFactory.setHost("127.0.0.1");
        connectionFactory.setPort(5672);
        connectionFactory.setVirtualHost("spm_qa");
        connectionFactory.setUsername("spm");
        connectionFactory.setPassword("p");
    }

    public static void main(String[] args) throws IOException, TimeoutException, InterruptedException {
        List.of("a.b", "a.#", "*.b", "*.b.*", "a.*").forEach((routingKey) -> {
            new Thread(() -> {
                try {
                    Connection connection = connectionFactory.newConnection();
                    Channel channel = connection.createChannel();
                    channel.exchangeDeclare(EXCHANGE_NAME, "topic");
                    String queue = channel.queueDeclare().getQueue();
                    channel.queueBind(queue, EXCHANGE_NAME, routingKey);
                    channel.basicConsume(queue, true, new DefaultConsumer(channel) {
                        @Override
                        public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                            System.out.println("routingKey:\t" + routingKey + " messageRoutingKey: " + new String(body));
                        }
                    });
                } catch (IOException | TimeoutException e) {
                    throw new RuntimeException(e);
                }
            }).start();
        });

        Thread.sleep(500);
        Connection connection = connectionFactory.newConnection();
        Channel channel = connection.createChannel();
        channel.exchangeDeclare(EXCHANGE_NAME, "topic");
        channel.basicPublish(EXCHANGE_NAME, "a.b", null, "a.b".getBytes(StandardCharsets.UTF_8));
        channel.basicPublish(EXCHANGE_NAME, "a.b.c", null, "a.b.c".getBytes(StandardCharsets.UTF_8));
        channel.basicPublish(EXCHANGE_NAME, "a", null, "a".getBytes(StandardCharsets.UTF_8));
    }
}
```

**输出：**

```log
routingKey: *.b     messageRoutingKey: a.b
routingKey: a.b     messageRoutingKey: a.b
routingKey: a.#     messageRoutingKey: a.b
routingKey: a.*     messageRoutingKey: a.b
routingKey: a.#     messageRoutingKey: a.b.c
routingKey: *.b.*   messageRoutingKey: a.b.c
routingKey: a.#     messageRoutingKey: a
```

### RPC

## SpringBoot 整合

**依赖:**

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-amqp</artifactId>
</dependency>
```

**配置：**

```yml
spring:
  rabbitmq:
    host: localhost
    port: 5672
    username: spm
    password: p
    virtual-host: spm_qa
```

**生产者：**
SpringBoot 启动时会自动创建 `RabbitTemplate` 实例，我们只需要使用 `convertAndSend` 方法就可以实现发送消息到 mq。

```java
@Autowired
RabbitTemplate rabbitTemplate;
void sendMsg() {
    rabbitTemplate.convertAndSend("hello-world", "hello world msg");
    for (int i = 1; i <= 10; i++) {
        rabbitTemplate.convertAndSend("worker", "worker msg" + i);
    }
    rabbitTemplate.convertAndSend("outbound_msg", "", "fanout msg");
    rabbitTemplate.convertAndSend("topic_ex", "a.a", "topic msg");
}
```

**定义消费者：**
常用注解

- `@Queue` 定义队列相关，在没有指定名称时创建的是一个临时队列
- `@Exchange` 定义交换机相关，通过 `type` 可以指定交换机类型
- `@RabbitListener` 定义在 customer 函数上面，可以绑定队列与交换机
- `@QueueBinding` 绑定队列与交换机

**Hello world：**

```java
@RabbitListener(queuesToDeclare = @Queue("hello-world"))
public void helloWorld(String msg) {
    System.out.println("收到消息：" + msg);
}
```

**fanout:**
定义一个`临时队列`和 `outbound_msg 交换机`，并将其绑定在一起。

```java
@RabbitListener(bindings = @QueueBinding(
        value = @Queue,
        exchange = @Exchange(type = ExchangeTypes.FANOUT, name = "outbound_msg"))
)
public void fanout1(String msg) {
    System.out.println("fanout1 收到消息：" + msg);
}

@RabbitListener(bindings = @QueueBinding(
        value = @Queue,
        exchange = @Exchange(type = ExchangeTypes.FANOUT, name = "outbound_msg"))
)
public void fanout2(String msg) {
    System.out.println("fanout2 收到消息：" + msg);
}
```

**Topics:**

```java
@RabbitListener(bindings = @QueueBinding(
        value = @Queue,
        exchange = @Exchange(type = ExchangeTypes.TOPIC, name = "topic_ex"), key = "*.a.#")
)
public void topic(String msg) {
    System.out.println("topic 收到消息：" + msg);
}
```

## 集群

### 普通集群

Slave 节点不能同步 Master 节点队列中的数据

### 镜像集群
