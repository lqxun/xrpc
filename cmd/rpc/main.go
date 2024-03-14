package main

import (
	"fmt"
	"github.com/lqxun/xrpc/app/rpc/service"
	"github.com/lqxun/xrpc/config"
)

func main() {
	conf := config.NewService()
	addr := fmt.Sprintf(":%d", conf.Port)

	// 服务端
	service.Service().ListenAndServe(addr)
}
