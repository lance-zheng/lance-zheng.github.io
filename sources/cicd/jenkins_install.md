<!-- customize-category: CI/CD -->

# 安装 Jenkins

<https://www.jenkins.io/doc/book/installing/war-file/>

安装 JDK

```shell
yum install java-11-openjdk.aarch64 -y
```

下载 Jenkins

<https://www.jenkins.io/download/>

```shell
wget https://get.jenkins.io/war-stable/2.401.1/jenkins.war
```

启动

```shell
java -jar jenkins.war
```

浏览器打开 <http://192.168.30.140:8080/>，然后输入控制台的密码
