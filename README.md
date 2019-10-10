# gogrpc
用Golang写个grpc的demo

# 安装说明
## 1. 由于大陆网络无法下载google的包，下载包前执行命令
```
> go env -w GOPROXY=https://goproxy.cn,direct
```
## 2. 下载 grpc 包，以及 protoc-gen-go 工具
```
> go get -u google.golang.org/grpc
> go get -u github.com/golang/protobuf/protoc-gen-go
```
如果成功，会在GOPATH/bin下生成protoc-gen-go.exe文件

## 3. 下载 protoc，去到下载页面找到对应包
https://github.com/protocolbuffers/protobuf/releases

windows：下载 protoc-3.10.0-win64.zip

linux：protoc-3.10.0-linux-x86_64.zip

解压可以得到一个protoc文件，下一步需要用到

## 4. 执行 generate_hello.bat 生成 helloworld.pb.go
修改 generate_hello.bat 里面的 protoc 文件的存放目录
```
set TOOLS_PATH=你的protoc文件的存放目录
```

执行前，你可以先删除 src/helloworld/helloworld.pd.go，然后自己生成
```
> .\generate_hello.bat
```