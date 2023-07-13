<!-- customize-category:OS -->

# Docker

- [Docker](#docker)
  - [快速创建实例](#快速创建实例)
    - [Nginx](#nginx)
    - [MySQL](#mysql)
  - [其他](#其他)

## 快速创建实例

### Nginx

使用 Nginx 快速创建一个文件服务器

将下面的配置文件放入 `~/Public/Share/nginx.conf`

```txt
user root;
worker_processes auto;

error_log /var/log/nginx/error.log notice;
pid /var/run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include /etc/nginx/mime.types;
    default_type application/octet-stream;

    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
    '$status $body_bytes_sent "$http_referer" '
    '"$http_user_agent" "$http_x_forwarded_for"';

    access_log /var/log/nginx/access.log main;

    sendfile on;

    keepalive_timeout 65;
    gzip on;
    include /etc/nginx/conf.d/*.conf;
    server {
        listen 8080;
        server_name localhost;
        location / {
            autoindex on;
            root /var/share/data ;
            mp4;
            charset utf-8;
        }
    }
}
```

```shell
docker run --name nginx-share \
-p 8090:8080 \
-v ~/Public/Share:/var/share/data \
-v ~/Public/Share/nginx.conf:/etc/nginx/nginx.conf:ro \
-d nginx
```

### MySQL

latest version

```shell
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=p -d mysql/mysql-server
```

MySQL 5.7

```shell
docker run -d --name mysql5.7 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=p mysql/mysql-server:5.7
```

配置允许远程 ip 访问

```shell

docker exec -it mysql bash
mysql -u root -p
use mysql;
update user set Host='%' where User='root';
flush privileges;`
```

## 其他

- 使用 [Colima](https://github.com/abiosoft/colima) 替换 Docker Desktop
