# gogrpc
用Golang写个grpc的demo

## 更新于2019年10月10日

## 安装说明
1. 由于大陆网络无法下载google的包，使用七牛公司的代理，下载包前执行命令
```
> go env -w GOPROXY=https://goproxy.cn,direct
```
2. 下载 grpc 包
```
> go get -u google.golang.org/grpc
> go get -u github.com/golang/protobuf/protoc-gen-go
```
3. 安装 protoc
```
这个我工具从 nuget 上获取 (.net core)
```
4. 执行 generate_hello.bat 生成 helloworld.pb.go
```
> generate_hello.bat
```