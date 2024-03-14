package test

import (
	"fmt"
	middleware2 "github.com/lqxun/xrpc/app/rpc/middleware"
	"github.com/lqxun/xrpc/app/services"
	"github.com/lqxun/xrpc/config"
	"github.com/lqxun/xrpc/core/server"
)

func ExampleRegisterServiceTest() {
	conf := config.NewService()
	addr := fmt.Sprintf(":%d", conf.Port)

	// 设置中间件
	middlewares := []server.Handler{middleware2.F1, middleware2.F2}

	// 服务端
	srv := server.NewXRpc()
	srv.Register("hello", services.Hello, middlewares...)
	go srv.ListenAndServe(addr)

	// 阻塞等待链接成功
	srv.Success()

	// Output:
	//Listening on :7777
}
