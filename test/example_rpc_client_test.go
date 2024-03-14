package test

import (
	"context"
	"fmt"
	"github.com/lqxun/xrpc/config"
	"github.com/lqxun/xrpc/core/server"
	"net"
	"time"
)

func ExampleRPCClientTest() {
	// 获取配置
	conf := config.NewService()
	addr := fmt.Sprintf(":%d", conf.Port)

	// 服务启动成功后开始建立链接
	fmt.Println("connecting...")
	conn, err := net.Dial("tcp", addr)
	fmt.Println("connected...")
	if err != nil {
		panic(err)
	}
	cli := server.NewClient(conn)
	var hello func(name string) (string, error)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli.Call(ctx, "hello", &hello)

	u, err := hello("张三")
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	// Output:
	//middleware 1 start
	//middleware 2 start
	//handler...
	//middleware 2 end
	//middleware 1 end
	//hello, 张三
}
