# xrpc
## RPC 功能概述
> 测试用例分开执行
```shell
go test -v ./test/example_rpc_server_and_client_test.go
go test -v ./test/example_rpc_server_test.go
go test -v ./test/example_rpc_client_test.go
```
### 1、注册服务
服务配置目录在`app/rpc/service/service.go`中
```go
// 设置监听端口
conf := config.NewService()
addr := fmt.Sprintf(":%d", conf.Port)

// 设置中间件
middlewares := []server.Handler{middleware2.F1, middleware2.F2}

// 注册服务
srv := server.NewXRpc()
srv.Register("hello", services.Hello, middlewares...)

```

### 监听请求端口
```shell
# 该操作会阻塞，可以通过测试用例检测
go run cmd/rpc/main.go
```


### 3、服务调用
```go
// 设置监听端口 默认为：7777
conf := config.NewService()
addr := fmt.Sprintf(":%d", conf.Port)

// 请求链接
conn, err := net.Dial("tcp", addr)
if err != nil {
    panic(err)
}

// 数据编码和解码
cli := server.NewClient(conn)

// 设置请求超时
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// 调用服务
var hello func(name string) (string, error)
cli.Call(ctx, "hello", &hello)
u, err := hello("张三")
if err != nil {
    panic(err)
}

// 输出响应
fmt.Println(u)
```

## HTTP 功能概述
### 启动Http服务
> http 默认端口为：7000
```shell
go run cmd/web/main.go
```

### 测试
```shell
curl http://127.0.0.1:7000/?name=zhangsan
```


 - [x] RPC服务端
 - [x] RPC客户端
 - [x] 路由管理
 - [x] 中间件实现
 - [x] 多协议支持,同时支持http和rpc(多端口实现)
 - [x] 序列化，反序列化
 - [ ] 通过server options实现超时配置
 - [ ] 安全退出，信号监听
 - [ ] 配置统一管理、待完善
