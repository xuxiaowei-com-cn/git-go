# Git Go

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go get -u github.com/urfave/cli/v2
go get -u gopkg.in/yaml.v3
```

### mod

```shell
go mod tidy
```

### run

```shell
go run main.go
```

### run help

```shell
go run main.go help
```

```shell
$ go run main.go help
NAME:
   git-go - 基于 Go 语言的 Git 命令行工具

USAGE:
   git-go [global options] command [command options] [arguments...]

VERSION:
   dev

AUTHOR:
   徐晓伟 <xuxiaowei@xuxiaowei.com.cn>

COMMANDS:
   commit-author, commitAuthor, ca         Name <email> 格式的提交作者
   commit-branch, commitBranch, cb         提交分支名称
   commit-sha, commitSha, cs               项目为其构建的提交修订
   commit-short-sha, commitShortSha, css   项目为其构建的提交修订的前八个字符
   commit-tag, commitTag, ct               提交标签名称
   commit-timestamp, commitTimestamp, cts  ISO 8601 格式的提交时间戳，如：2023-10-02T00:29:17+08:00
   now                                     ISO 8601 格式的当前时间戳，如：2023-10-02T02:35:20+08:00
   go-version, goVersion, gv               go version 命令返回值，如：go version go1.21.1 windows/amd64
   go-short-version, goShortVersion, gsv   go version 命令返回值截取版本号，如：go1.21.1
   go-platform, goPlatform, gp             go version 命令返回值截取平台信息，如：windows/amd64
   help, h                                 Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>
```

### test

```shell
go test ./... -v
```

### build

```shell
go build
# GOOS=：设置构建的目标操作系统（darwin | freebsd | linux | windows）
# GOARCH=：设置构建的目标操作系统（386 | amd64 | arm | arm64）
# -v：打印编译过程中的详细信息
# -ldflags：设置在编译时传递给链接器的参数
# -ldflags "-s -w -buildid="
#                           -s: 删除符号表信息，减小可执行文件的大小。
#                           -w: 删除调试信息，使可执行文件在运行时不会打印调试信息。
#                           -buildid=: 删除构建ID，使可执行文件在运行时不会打印构建ID。
# -trimpath：去掉所有包含 go path 的路径
# -o：指定构建后输出的文件名
```

- Windows
    - amd64
        ```shell
        go build -o buildinfo/buildinfo.exe buildinfo/buildinfo.go
        GOOS=windows GOARCH=amd64 go build -v -ldflags "-s -w -buildid= -X main.GitCommit=$(buildinfo/buildinfo)" -trimpath -o git-go-windows-amd64.exe .
        ```
    - arm64
        ```shell
        GOOS=windows GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o git-go-windows-arm64.exe .
        ```

- Linux
    - amd64
        ```shell
        GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w -buildid=" -trimpath -o git-go-linux-amd64 .
        ```
    - arm64
        ```shell
        GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o git-go-linux-arm64 .
        ```

- Darwin
    - amd64
        ```shell
        GOOS=darwin GOARCH=amd64 go build -v -ldflags "-s -w -buildid=" -trimpath -o git-go-darwin-amd64 .
        ```
    - arm64
        ```shell
        GOOS=darwin GOARCH=arm64 go build -v -ldflags "-s -w -buildid=" -trimpath -o git-go-darwin-arm64 .
        ```
