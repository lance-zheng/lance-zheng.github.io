<!-- customize-category: CI/CD -->

# 安装 GitLab

> <https://about.gitlab.com/>
>
> <https://gitlab.cn/install/>

<!-- ## Installation

1. 安装相关依赖

   ```shell
   sudo yum install -y curl policycoreutils-python openssh-server perl
   sudo systemctl enable sshd
   sudo systemctl start sshd
   ```

2. 添加 GitLab 源镜像

   ```shell
   curl -fsSL https://packages.gitlab.cn/repository/raw/scripts/setup.sh | /bin/bash
   ```

3. 安装 GitLab

```shell
# 将https://gitlab.example.com 换成对应 ip
sudo EXTERNAL_URL="https://gitlab.example.com" yum install -y gitlab-jh

```

...没有 arm64 版本，后续在补充 -->

## Docker 安装 (M1 Mac)

安装 Docker

```shell
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

安装 GitLab (ARM64 版本)

```shell
docker run \
  --detach \
  --name gitlab-ce \
  --privileged \
  --memory 4096M \
  --publish 8822:22 \
  --publish 8880:80 \
  --publish 8443:443 \
  --hostname lance.gitlab.com \
  --env GITLAB_OMNIBUS_CONFIG=" \
    nginx['redirect_http_to_https'] = true; "\
  yrzr/gitlab-ce-arm64v8:latest
```

修改密码

```shell
docker exec -it  b0 /bin/bash
gitlab-rails console
user = User.where(id: 1).first
user.password = 'core@123'
user.password_confirmation = 'core@123'
user.save
```
