<!-- customize-category:OS -->

# MacOS

- [MacOS](#macos)
  - [VIM](#vim)
    - [常用快捷键](#常用快捷键)
  - [Homebrew](#homebrew)
    - [Homebrew 常用指令](#homebrew-常用指令)
    - [Homebrew 卸载软件及其依赖包](#homebrew-卸载软件及其依赖包)
    - [Upgrade cask app](#upgrade-cask-app)
    - [Backup and restore](#backup-and-restore)
  - [虚拟机](#虚拟机)
    - [虚拟机中安装 CentOS 7](#虚拟机中安装-centos-7)
  - [Docker](#docker)
  - [其他](#其他)
    - [Dock 栏弹出延迟](#dock-栏弹出延迟)

## VIM

### 常用快捷键

```text
# insert
ctrl h   # 退格
ctrl w   # 删除上一个单词
ctrl u   # 删除行

# normal
gi       # 上次编辑位置
0        # 行首
$        # 行尾
f char   # 搜索移动
cw       # 替换单词

# command
vs       # 垂直分屏
sp       # 水平分屏

# terminal
ctrl e   # 结尾
ctrl a   # 开头
ctrl b   # 前移
ctrl f   # 后移
```

## [Homebrew](https://brew.sh/)

> [Homebrew Formulae](https://formulae.brew.sh/)

### Homebrew 常用指令

```sh
# search
brew search redis

# 列出已安装的软件包
brew list

# install
brew install wget

# 卸载
bew uninstall vim

# zap 可以删除相关的文件夹，如果想要完全卸载的话可以使用这个参数
# brew cat firfox 可以查看到有哪些关联的文件会被同时删除
bew uninstall firefox --zap

# 删除无用依赖包  --dry-run 将自会列出不会删除
brew autoremove --dry-run

# 更新 Homebrew 以及 Formula
brew update

# 查看可以更新的包
brew outdated

# 更新软件
brew upgrade [vim]

# 查看依赖包
brew deps vim

# 查看包信息
brew info vim

# 清理安装包缓存 [-n 只查看不删除]
brew cleanup

# 列出不依赖其他包的包
brew leaves

# remove tap
brew untap sidneys/homebrew
```

### Homebrew 卸载软件及其依赖包

> `brew uninstall xxx` 无法卸载相关依赖包  
> 可以使用 `rmtree` 实现卸载软件包同时卸载相关依赖

```sh
# install
brew tap beeftornado/rmtree
brew install brew-rmtree

# example
brew rmtree maven
```

### Upgrade cask app

> <https://github.com/buo/homebrew-cask-upgrade>

```sh
# install
brew tap buo/cask-upgrade

# upgrade specific app
brew -a cu [CASK]

# upgrade all app
brew cu -a
```

### Backup and restore

```sh
brew tap Homebrew/bundle

# backup
# By default, the Brewfile will be located at ~/Brewfile
brew bundle dump

# restore
brew bundle --file=~/Brewfile
```

## 虚拟机

UTM 和 Vmware Fusion 都是不错的选择

[UTM](https://github.com/utmapp/UTM)

```sh
brew install --cask utm
```

Vmware Fusion

```sh
brew install --cask vmware-fusion

# License
# 4C21U-2KK9Q-M8130-4V2QH-CF810
```

### 虚拟机中安装 CentOS 7

CentOS7 镜像  
<https://mirrors.tuna.tsinghua.edu.cn/centos-altarch/7.9.2009/isos/aarch64/>  
若使用 **Vmware Fusion** 需要使用下面这个镜像，官方的镜像无法正常安装  
<https://www.aliyundrive.com/s/unTcan2mKHX>

> UTM 安装注意事项  
> 安装时需要选择下面这个选项，否则会遇到和 **Vmware Fusion** 一样的问题

<img width=400 src='/assets/image/1682338781.png'/>

**安装成功后续操作：**

1. 网路配置  
   开机自动启动网卡，将下面文件中的 `ONBOOT` 值改为 `yes`

   ```shell
   vi /etc/sysconfig/network-scripts/ifcfg-enp0s1
   ```

   启动网卡

   ```shell
   systemctl start network
   ```

2. 配置 Yum 源

   ```shell
   # 备份 yum 源
   mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo_bak
   # 使用 aliyun yum 源
   curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-altarch-7.repo
   # 重建 yum 缓存
   yum clean all && yum makecache
   ```

3. 其他配置

   ```shell
   # 关闭防火墙
   systemctl stop firewalld
   systemctl disable firewalld
   ```

**可能遇到的问题：**

- SSH 连接时出现下面这个 error

  ```text
  Centos warning: setlocale: LC_CTYPE: cannot change locale (UTF-8): No such file or directory
  ```

  解决方法：`sudo vi /etc/environment`

  ```text
  LANG="en_US.UTF-8"
  LC_ALL="en_US.UTF-8"
  LC_CTYPE="en_US.UTF-8"
  ```

- 重启网卡失败： `systemctl restart network` 出现下面 error，这个问题发生在我克隆原有的实例

  ```txt
  Job for network.service failed because the control process exited with error code. See "systemctl status network.service" and "journalctl -xe" for details.
  ```

  解决方法：删除多余网卡
  通过 `ip addr` 查看当前使用的网卡，通过 `ll /etc/sysconfig/network-scripts/ifcfg*` 查看网卡配置文件  
  <img width=400 src='/assets/image/1682347983.png'/>  
  将 `ifcfg-enp0s1` 删除

## Docker

> 使用 [Colima](https://github.com/abiosoft/colima) 替代 Docker Desktop

```sh
# install
brew install colima docker
# start
colima start
# stop
colima stop
```

**配置国内源：**
可以直接修改 `$HOME/.colima/default/colima.yaml` 然后重启下 Colima 也可以使用 `colima start --edit` 在创建 instance 时修改

```yml
# $HOME/.colima/default/colima.yaml
docker:
  registry-mirrors:
    - https://hub-mirror.c.163.com
```

**配置开机自动运行脚本：**
在用户的启动项中添加此应用程序即可

```sh
eval "$(/opt/homebrew/bin/brew shellenv)"
colima start
```

<img src="./assets/image/colima.png" width = "600"  alt="" align=center />

## 其他

### Dock 栏弹出延迟

```sh
# 延迟 0 秒弹出
defaults write com.apple.Dock autohide-delay -float 0 && killall Dock
# 恢复默认延迟
defaults delete com.apple.Dock autohide-delay && killall Dock
```
