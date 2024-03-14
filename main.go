package main

import (
	"context"
	"fmt"
	"github.com/lqxun/xrpc/config"
	"github.com/lqxun/xrpc/core/server"
	"net"
	"time"
)

func main() {

	conf := config.NewService()
	addr := fmt.Sprintf(":%d", conf.Port)

	// 客户端
	conn, err := net.Dial("tcp", addr)
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
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
}
