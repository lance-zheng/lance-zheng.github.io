<!-- customize-category:Go -->

# Go

<https://go.dev/>  
<https://github.com/kevincobain2000/gobrew>

---

常用命令：

- `go install` github.com/gin-gonic/gin@v1.9.0 安装可执行文件
- `go get` github.com/gin-gonic/gin@v1.9.0
  添加依赖同时更新 `go.mod` 文件
- `go get -u` github.com/gin-gonic/gin@v1.9.0
  更新依赖
- `go clean -modcache` 删除本地下载的依赖 `GOPATH/pkg/mod`

## Go Module

模块初始化命令 `go mod init github.com/lance-zheng/xxx`，此命令会在当前目录生成 `go.mod` 文件

常用指令：

- `go mod help`
- `go mod init` 初始化模块
- `go mod tidy` 下载依赖同时删除未使用的依赖
- `go mod edit xxx` 修改 `go.mod` 文件
- `go mod vendor` 将现有的依赖复制到 vendor 目录下
- `go mod verify` 验证 `go.mod` 文件
- `go mod why xxxxx` 查询为什么会依赖这个文件

```txt
// 模块名
module github.com/lance-zheng/go-learning

// go sdk version
go 1.20

// 依赖
require (
)

// 排除的依赖
exculde (
)

// 修改依赖包的路径
replace (
)

// 撤回某些版本
retract (
)
```

## 变量

## 泛型
